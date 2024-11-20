// tbl.go
package tbl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// Constants for file handling
const (
	ColorHeader    = "ÿc"
	NewTblFileName = "!newstring!.tbl"

	HeaderSize          = 0x15 // 21 bytes
	HashNodeSize        = 0x11 // 17 bytes
	DefaultVersion      = 1
	DefaultHashMaxTries = 1
)

// TblHeader represents the complete TBL file header (21 bytes total)
type TblHeader struct {
	CRC             uint16 // +0x00 - CRC value (2 bytes)
	NodesNumber     uint16 // +0x02 - size of Indices array (2 bytes)
	HashTableSize   uint32 // +0x04 - size of TblHashNode array (4 bytes)
	Version         uint8  // +0x08 - file version, 0 or 1 (1 byte)
	DataStartOffset uint32 // +0x09 - string table start offset (4 bytes)
	HashMaxTries    uint32 // +0x0D - max collisions for string key search (4 bytes)
	FileSize        uint32 // +0x11 - size of the file (4 bytes)
}

// TblHashNode represents a hash table entry (17 bytes total)
type TblHashNode struct {
	Active          uint8  // +0x00 - shows if entry is used (1 byte)
	Index           uint16 // +0x01 - index in Indices array (2 bytes)
	HashValue       uint32 // +0x03 - hash value of current string key (4 bytes)
	StringKeyOffset uint32 // +0x07 - offset of current string key (4 bytes)
	StringValOffset uint32 // +0x0B - offset of current string value (4 bytes)
	StringValLength uint16 // +0x0F - length of current string value (2 bytes)
}

// TblEntry represents a processed table entry
type TblEntry struct {
	ID    int
	Key   string
	Value string
}

// Color code handling
type ColorCode struct {
	Code  rune
	Name  string
	Color string
}

// TblParser handles parsing and writing of .tbl files
type TblParser struct {
	File     *os.File
	header   TblHeader
	Encoder  *charmap.Charmap
	Entries  map[string][]TblEntry
	ByID     map[int]TblEntry
	ByValue  map[string][]int
	NextID   int
	modified bool
	Filename string
}

// // TblStructure related enhancements
// type TblStructure struct {
// 	header TblHeader
// 	data   []DataNode
// }

type DataNode struct {
	Index     int
	KeyOffset uint32
	Key       string
	Value     string
}

// CRC calculation table - matches C++ implementation
var crcTable = [256]uint16{
	0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50A5, 0x60C6, 0x70E7, 0x8108, 0x9129, 0xA14A, 0xB16B, 0xC18C, 0xD1AD, 0xE1CE, 0xF1EF,
	0x1231, 0x0210, 0x3273, 0x2252, 0x52B5, 0x4294, 0x72F7, 0x62D6, 0x9339, 0x8318, 0xB37B, 0xA35A, 0xD3BD, 0xC39C, 0xF3FF, 0xE3DE,
	0x2462, 0x3443, 0x0420, 0x1401, 0x64E6, 0x74C7, 0x44A4, 0x5485, 0xA56A, 0xB54B, 0x8528, 0x9509, 0xE5EE, 0xF5CF, 0xC5AC, 0xD58D,
	0x3653, 0x2672, 0x1611, 0x0630, 0x76D7, 0x66F6, 0x5695, 0x46B4, 0xB75B, 0xA77A, 0x9719, 0x8738, 0xF7DF, 0xE7FE, 0xD79D, 0xC7BC,
	0x48C4, 0x58E5, 0x6886, 0x78A7, 0x0840, 0x1861, 0x2802, 0x3823, 0xC9CC, 0xD9ED, 0xE98E, 0xF9AF, 0x8948, 0x9969, 0xA90A, 0xB92B,
	0x5AF5, 0x4AD4, 0x7AB7, 0x6A96, 0x1A71, 0x0A50, 0x3A33, 0x2A12, 0xDBFD, 0xCBDC, 0xFBBF, 0xEB9E, 0x9B79, 0x8B58, 0xBB3B, 0xAB1A,
	0x6CA6, 0x7C87, 0x4CE4, 0x5CC5, 0x2C22, 0x3C03, 0x0C60, 0x1C41, 0xEDAE, 0xFD8F, 0xCDEC, 0xDDCD, 0xAD2A, 0xBD0B, 0x8D68, 0x9D49,
	0x7E97, 0x6EB6, 0x5ED5, 0x4EF4, 0x3E13, 0x2E32, 0x1E51, 0x0E70, 0xFF9F, 0xEFBE, 0xDFDD, 0xCFFC, 0xBF1B, 0xAF3A, 0x9F59, 0x8F78,
	0x9188, 0x81A9, 0xB1CA, 0xA1EB, 0xD10C, 0xC12D, 0xF14E, 0xE16F, 0x1080, 0x00A1, 0x30C2, 0x20E3, 0x5004, 0x4025, 0x7046, 0x6067,
	0x83B9, 0x9398, 0xA3FB, 0xB3DA, 0xC33D, 0xD31C, 0xE37F, 0xF35E, 0x02B1, 0x1290, 0x22F3, 0x32D2, 0x4235, 0x5214, 0x6277, 0x7256,
	0xB5EA, 0xA5CB, 0x95A8, 0x8589, 0xF56E, 0xE54F, 0xD52C, 0xC50D, 0x34E2, 0x24C3, 0x14A0, 0x0481, 0x7466, 0x6447, 0x5424, 0x4405,
	0xA7DB, 0xB7FA, 0x8799, 0x97B8, 0xE75F, 0xF77E, 0xC71D, 0xD73C, 0x26D3, 0x36F2, 0x0691, 0x16B0, 0x6657, 0x7676, 0x4615, 0x5634,
	0xD94C, 0xC96D, 0xF90E, 0xE92F, 0x99C8, 0x89E9, 0xB98A, 0xA9AB, 0x5844, 0x4865, 0x7806, 0x6827, 0x18C0, 0x08E1, 0x3882, 0x28A3,
	0xCB7D, 0xDB5C, 0xEB3F, 0xFB1E, 0x8BF9, 0x9BD8, 0xABBB, 0xBB9A, 0x4A75, 0x5A54, 0x6A37, 0x7A16, 0x0AF1, 0x1AD0, 0x2AB3, 0x3A92,
	0xFD2E, 0xED0F, 0xDD6C, 0xCD4D, 0xBDAA, 0xAD8B, 0x9DE8, 0x8DC9, 0x7C26, 0x6C07, 0x5C64, 0x4C45, 0x3CA2, 0x2C83, 0x1CE0, 0x0CC1,
	0xEF1F, 0xFF3E, 0xCF5D, 0xDF7C, 0xAF9B, 0xBFBA, 0x8FD9, 0x9FF8, 0x6E17, 0x7E36, 0x4E55, 0x5E74, 0x2E93, 0x3EB2, 0x0ED1, 0x1EF0,
}

// NewTblParser creates a new parser instance
func NewTblParser(filePath string) (*TblParser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	parser := &TblParser{
		File:     file,
		Encoder:  charmap.Windows1252,
		Entries:  make(map[string][]TblEntry),
		ByID:     make(map[int]TblEntry),
		ByValue:  make(map[string][]int),
		NextID:   1,
		Filename: filePath,
	}

	return parser, nil
}

// Close releases resources
func (p *TblParser) Close() error {
	if p.File != nil {
		return p.File.Close()
	}
	return nil
}

// Parse determines file type and processes accordingly
func (p *TblParser) Parse() error {
	// Read first few bytes to determine file type
	header := make([]byte, 2)
	if _, err := p.File.Read(header); err != nil {
		return fmt.Errorf("error reading file header: %w", err)
	}
	p.File.Seek(0, 0)

	ext := strings.ToLower(filepath.Ext(p.Filename))

	// Determine file type
	if ext == ".tbl" || (header[0] == 0 && header[1] == 0) {
		return p.parseTBL()
	} else if ext == ".txt" || ext == ".csv" {
		return p.parseText()
	}

	return fmt.Errorf("unsupported file format")
}

// parseTBL handles binary TBL format
func (p *TblParser) parseTBL() error {
	defer p.Close()

	// Read header
	header, err := readTblHeader(p.File)
	if err != nil {
		return fmt.Errorf("reading header: %w", err)
	}
	p.header = header

	// Get file size for validation
	fileInfo, err := p.File.Stat()
	if err != nil {
		return fmt.Errorf("getting file info: %w", err)
	}
	actualFileSize := fileInfo.Size()

	// Debug header info
	fmt.Printf("File Size: %d bytes\n", actualFileSize)
	fmt.Printf("Header Info:\n")
	fmt.Printf("  CRC: 0x%04X\n", p.header.CRC)
	fmt.Printf("  NodesNumber: %d\n", p.header.NodesNumber)
	fmt.Printf("  HashTableSize: %d\n", p.header.HashTableSize)
	fmt.Printf("  Version: %d\n", p.header.Version)
	fmt.Printf("  DataStartOffset: %d\n", p.header.DataStartOffset)
	fmt.Printf("  HashMaxTries: %d\n", p.header.HashMaxTries)
	fmt.Printf("  FileSize: %d\n", p.header.FileSize)

	// Skip index table (uint16 array)
	indexTableSize := int64(p.header.NodesNumber) * 2
	if _, err := p.File.Seek(indexTableSize, io.SeekCurrent); err != nil {
		return fmt.Errorf("skipping index table: %w", err)
	}

	// Read hash nodes
	fmt.Printf("Reading %d hash nodes...\n", p.header.HashTableSize)
	hashNodes := make([]TblHashNode, p.header.HashTableSize)
	for i := uint32(0); i < p.header.HashTableSize; i++ {
		node, err := readTblHashNode(p.File)
		if err != nil {
			return fmt.Errorf("reading hash node %d: %w", i, err)
		}
		hashNodes[i] = node
	}

	// Read string data
	stringData, err := io.ReadAll(p.File)
	if err != nil {
		return fmt.Errorf("reading string data: %w", err)
	}

	// Process entries
	activeNodes := 0
	for i, node := range hashNodes {
		if node.Active == 0 {
			continue
		}
		activeNodes++

		key, err := p.extractString(stringData, node.StringKeyOffset-p.header.DataStartOffset)
		if err != nil {
			return fmt.Errorf("extracting key from node %d: %w", i, err)
		}

		value, err := p.extractString(stringData, node.StringValOffset-p.header.DataStartOffset)
		if err != nil {
			return fmt.Errorf("extracting value from node %d: %w", i, err)
		}

		entry := TblEntry{
			ID:    p.NextID,
			Key:   key,
			Value: value,
		}

		p.Entries[key] = append(p.Entries[key], entry)
		p.ByID[p.NextID] = entry
		p.ByValue[value] = append(p.ByValue[value], p.NextID)
		p.NextID++
	}

	fmt.Printf("Processed %d active nodes out of %d total nodes\n", activeNodes, len(hashNodes))
	return nil
}

// parseText handles CSV/TSV formats
func (p *TblParser) parseText() error {
	defer p.Close()

	content, err := io.ReadAll(p.File)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	lines := bytes.Split(content, []byte{'\n'})
	if len(lines) == 0 {
		return nil
	}

	// Determine separator and wrapping
	var separator byte = '\t'
	hasWrapping := false

	if bytes.Contains(lines[0], []byte{'"'}) {
		hasWrapping = true
		if bytes.Contains(lines[0], []byte{','}) {
			separator = ','
		} else if bytes.Contains(lines[0], []byte{';'}) {
			separator = ';'
		}
	}

	// Skip header if present
	start := 0
	if bytes.Contains(lines[0], []byte("Key")) || bytes.Contains(lines[0], []byte("String Index")) {
		start = 1
	}

	// Process entries
	for i := start; i < len(lines); i++ {
		line := bytes.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}

		key, value := p.splitTextLine(line, separator, hasWrapping)

		entry := TblEntry{
			ID:    p.NextID,
			Key:   key,
			Value: p.unfoldNewlines(value),
		}

		p.Entries[key] = append(p.Entries[key], entry)
		p.ByID[p.NextID] = entry
		p.ByValue[value] = append(p.ByValue[value], p.NextID)
		p.NextID++
	}

	return nil
}

// WriteFile writes the current data back to a file
// func (p *TblParser) WriteFile(filename string) error {
// 	ext := strings.ToLower(filepath.Ext(filename))

// 	var err error
// 	if ext == ".tbl" {
// 		err = p.writeTBL(filename)
// 	} else {
// 		err = p.writeText(filename)
// 	}

// 	if err != nil {
// 		return fmt.Errorf("error writing file: %w", err)
// 	}

// 	return nil
// }

// Helper functions
func (p *TblParser) extractString(data []byte, offset uint32) (string, error) {
	if int(offset) >= len(data) {
		return "", fmt.Errorf("invalid string offset: %d (data length: %d)", offset, len(data))
	}

	// Find the null terminator
	end := offset
	for end < uint32(len(data)) && data[end] != 0 {
		end++
	}
	if end >= uint32(len(data)) {
		return "", fmt.Errorf("no null terminator found for string at offset %d", offset)
	}

	decoded, err := p.Encoder.NewDecoder().Bytes(data[offset:end])
	if err != nil {
		return "", fmt.Errorf("error decoding string at offset %d: %w", offset, err)
	}

	return string(decoded), nil
}

// func (p *TblParser) processColorCodes(value string) string {
// 	result := value
// 	colorHeader := "ÿc"

// 	for strings.Contains(result, colorHeader) {
// 		idx := strings.Index(result, colorHeader)
// 		if idx+3 <= len(result) {
// 			colorCode := result[idx+2]
// 			result = result[:idx] + "\\c" + string(colorCode) + result[idx+3:]
// 		}
// 	}

// 	return result
// }

func (p *TblParser) foldNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}

func (p *TblParser) unfoldNewlines(s string) string {
	return strings.ReplaceAll(s, "\\n", "\n")
}

func (p *TblParser) splitTextLine(line []byte, separator byte, hasWrapping bool) (string, string) {
	parts := bytes.Split(line, []byte{separator})
	if len(parts) != 2 {
		return "", ""
	}

	key := string(bytes.TrimSpace(parts[0]))
	value := string(bytes.TrimSpace(parts[1]))

	if hasWrapping {
		key = strings.Trim(key, "\"")
		value = strings.Trim(value, "\"")
	}

	// Handle key decoding
	decodedKey, err := p.Encoder.NewDecoder().String(key)
	if err != nil {
		return "", ""
	}

	return decodedKey, value
}

// // writeTBL writes data in binary TBL format
// func (p *TblParser) writeTBL(filename string) error {
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	// Calculate offsets and prepare nodes
// 	dataStartOffset := uint32(binary.Size(TblHeader{}) + len(p.Entries)*2 + len(p.Entries)*binary.Size(TblHashNode{}))
// 	nodes := make([]TblHashNode, len(p.Entries))
// 	stringData := new(bytes.Buffer)

// 	currentOffset := dataStartOffset
// 	i := 0

// 	// Get sorted keys for consistent output
// 	sortedKeys := make([]string, 0, len(p.Entries))
// 	for k := range p.Entries {
// 		sortedKeys = append(sortedKeys, k)
// 	}
// 	sort.Strings(sortedKeys)

// 	// Process entries in sorted order
// 	for _, key := range sortedKeys {
// 		entries := p.Entries[key]
// 		for _, entry := range entries {
// 			keyBytes, err := p.encodeKey(entry.Key)
// 			if err != nil {
// 				return fmt.Errorf("error encoding key: %w", err)
// 			}
// 			valueBytes := []byte(entry.Value)

// 			nodes[i] = TblHashNode{
// 				Active:          1,
// 				Index:           uint32(i),
// 				HashValue:       p.calculateHash(keyBytes),
// 				StringKeyOffset: currentOffset,
// 				StringValOffset: currentOffset + uint32(len(keyBytes)) + 1,
// 				StringValLength: uint32(len(valueBytes)) + 1,
// 			}

// 			stringData.Write(keyBytes)
// 			stringData.WriteByte(0)
// 			stringData.Write(valueBytes)
// 			stringData.WriteByte(0)

// 			currentOffset += uint32(len(keyBytes)) + uint32(len(valueBytes)) + 2
// 			i++
// 		}
// 	}

// 	// Write header
// 	header := TblHeader{
// 		NodesNumber:     uint32(len(p.Entries)),
// 		HashTableSize:   uint32(len(p.Entries)),
// 		Version:         1,
// 		DataStartOffset: dataStartOffset,
// 		HashMaxTries:    1,
// 		FileSize:        currentOffset,
// 	}
// 	header.CRC = p.calculateCRC(stringData.Bytes())

// 	if err := binary.Write(f, binary.LittleEndian, &header); err != nil {
// 		return fmt.Errorf("error writing header: %w", err)
// 	}

// 	// Write indices
// 	for i := uint16(0); i < uint16(len(p.Entries)); i++ {
// 		if err := binary.Write(f, binary.LittleEndian, i); err != nil {
// 			return fmt.Errorf("error writing index: %w", err)
// 		}
// 	}

// 	// Write nodes
// 	for _, node := range nodes {
// 		if err := binary.Write(f, binary.LittleEndian, node); err != nil {
// 			return fmt.Errorf("error writing node: %w", err)
// 		}
// 	}

// 	// Write string data
// 	if _, err := f.Write(stringData.Bytes()); err != nil {
// 		return fmt.Errorf("error writing string data: %w", err)
// 	}

// 	return nil
// }

// writeText writes data in CSV/TSV format
func (p *TblParser) writeText(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(filename))
	isCsv := ext == ".csv"

	separator := byte('\t')
	if isCsv {
		separator = ','
	}

	// Write header
	if isCsv {
		f.WriteString("\"Key\",\"Value\"\n")
	} else {
		f.WriteString("Key\tValue\n")
	}

	// Write entries
	for _, entries := range p.Entries {
		for _, entry := range entries {
			key := entry.Key
			value := p.foldNewlines(entry.Value)

			if isCsv {
				fmt.Fprintf(f, "\"%s\"%c\"%s\"\n", key, separator, value)
			} else {
				fmt.Fprintf(f, "%s%c%s\n", key, separator, value)
			}
		}
	}

	return nil
}

// Enhanced hash calculation - matches C++ implementation exactly
func (p *TblParser) calculateHash(key []byte) uint32 {
	var hashValue uint32

	for _, c := range key {
		hashValue *= 0x10
		hashValue += uint32(c)

		if (hashValue & 0xF0000000) != 0 {
			tempValue := (hashValue & 0xF0000000) >> 24
			hashValue &= 0x0FFFFFFF
			hashValue ^= tempValue
		}
	}

	if len(p.Entries) == 0 {
		return hashValue
	}
	return hashValue % uint32(len(p.Entries))
}

// Enhanced CRC calculation - matches C++ implementation
func (p *TblParser) calculateCRC(data []byte) uint16 {
	crcValue := uint16(0xFFFF)

	for _, b := range data {
		charValue := uint16(b) ^ (crcValue >> 8)
		temp := (crcValue & 0xFF) << 8
		crcValue = crcTable[charValue] ^ temp
	}

	return crcValue
}

// Enhanced string encoding/decoding
func (p *TblParser) encodeKey(key string) ([]byte, error) {
	encoder := charmap.Windows1252.NewEncoder()
	return encoder.Bytes([]byte(key))
}

func (p *TblParser) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})
	stats["total_entries"] = len(p.ByID)
	stats["unique_keys"] = len(p.Entries)
	stats["unique_values"] = len(p.ByValue)
	stats["file_size"] = p.header.FileSize
	stats["version"] = p.header.Version
	stats["hash_max_tries"] = p.header.HashMaxTries
	return stats
}

func (p *TblParser) CloseAndCleanup() error {
	err := p.Close()
	p.Entries = nil
	p.ByID = nil
	p.ByValue = nil
	p.NextID = 0
	p.modified = false
	return err
}

// Helper function to read packed TblHeader
func readTblHeader(r io.Reader) (TblHeader, error) {
	var header TblHeader

	// Read the fields in order
	if err := binary.Read(r, binary.LittleEndian, &header.CRC); err != nil {
		return header, fmt.Errorf("reading CRC: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &header.NodesNumber); err != nil {
		return header, fmt.Errorf("reading NodesNumber: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &header.HashTableSize); err != nil {
		return header, fmt.Errorf("reading HashTableSize: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &header.Version); err != nil {
		return header, fmt.Errorf("reading Version: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &header.DataStartOffset); err != nil {
		return header, fmt.Errorf("reading DataStartOffset: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &header.HashMaxTries); err != nil {
		return header, fmt.Errorf("reading HashMaxTries: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &header.FileSize); err != nil {
		return header, fmt.Errorf("reading FileSize: %w", err)
	}

	return header, nil
}

// Helper function to read packed TblHashNode
func readTblHashNode(r io.Reader) (TblHashNode, error) {
	var node TblHashNode

	// Read the fields in order
	if err := binary.Read(r, binary.LittleEndian, &node.Active); err != nil {
		return node, fmt.Errorf("reading Active: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &node.Index); err != nil {
		return node, fmt.Errorf("reading Index: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &node.HashValue); err != nil {
		return node, fmt.Errorf("reading HashValue: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &node.StringKeyOffset); err != nil {
		return node, fmt.Errorf("reading StringKeyOffset: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &node.StringValOffset); err != nil {
		return node, fmt.Errorf("reading StringValOffset: %w", err)
	}
	if err := binary.Read(r, binary.LittleEndian, &node.StringValLength); err != nil {
		return node, fmt.Errorf("reading StringValLength: %w", err)
	}

	return node, nil
}

// func (p *TblParser) findHashSlot(hashValue uint32, key string) uint32 {
// 	slot := hashValue % uint32(len(p.entries))
// 	originalSlot := slot
// 	tries := uint32(0)

// 	// Linear probing
// 	for {
// 		if tries >= p.header.HashMaxTries {
// 			// Reset to original slot if we've tried too many times
// 			return originalSlot
// 		}

// 		// Check if slot is empty or matches our key
// 		if entry, exists := p.entries[key]; !exists || len(entry) == 0 {
// 			return slot
// 		}

// 		slot = (slot + 1) % uint32(len(p.entries))
// 		tries++
// 	}
// }

// func (p *TblParser) decodeKey(key []byte) (string, error) {
// 	decoder := charmap.Windows1252.NewDecoder()
// 	return decoder.String(string(key))
// }

// func (p *TblParser) validateHeader() error {
// 	if p.header.Version != DefaultVersion {
// 		return fmt.Errorf("unsupported TBL version: %d", p.header.Version)
// 	}
// 	if p.header.NodesNumber != p.header.HashTableSize {
// 		return fmt.Errorf("nodes number (%d) doesn't match hash table size (%d)",
// 			p.header.NodesNumber, p.header.HashTableSize)
// 	}
// 	if p.header.DataStartOffset < uint32(HeaderSize+p.header.NodesNumber*2+p.header.HashTableSize*HashNodeSize) {
// 		return fmt.Errorf("invalid data start offset: %d", p.header.DataStartOffset)
// 	}
// 	return nil
// }

// Add method to get sorted entries for consistent output
// func (p *TblParser) getSortedEntries() []TblEntry {
// 	entries := make([]TblEntry, 0, len(p.byID))
// 	for _, entry := range p.byID {
// 		entries = append(entries, entry)
// 	}
// 	sort.Slice(entries, func(i, j int) bool {
// 		return entries[i].ID < entries[j].ID
// 	})
// 	return entries
// }
