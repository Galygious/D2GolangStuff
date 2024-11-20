// main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"tbl-to-json/tbl"
)

// GenerateGoFile generates a Go source file with the parsed data
func generateGoFile(parser *tbl.TblParser, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer f.Close()

	// Write package header and imports
	if _, err := fmt.Fprintf(f, `package main

// Code generated from TBL file. DO NOT EDIT.

type TblEntry struct {
	ID    int
	Key   string
	Value string
}

`); err != nil {
		return fmt.Errorf("error writing header: %w", err)
	}

	// Write the three map structures
	if err := writeByKeyMap(f, parser.Entries); err != nil {
		return err
	}

	if err := writeByIDMap(f, parser.ByID); err != nil {
		return err
	}

	if err := writeByValueMap(f, parser.ByValue); err != nil {
		return err
	}

	return nil
}

func writeByKeyMap(f *os.File, data map[string][]tbl.TblEntry) error {
	if _, err := fmt.Fprintf(f, "var TblDataByKey = map[string][]TblEntry{\n"); err != nil {
		return err
	}

	entries := make([]string, 0, len(data))
	for key := range data {
		entries = append(entries, key)
	}
	sort.Strings(entries)

	for _, key := range entries {
		if _, err := fmt.Fprintf(f, "\t%q: {\n", key); err != nil {
			return err
		}
		for _, entry := range data[key] {
			if _, err := fmt.Fprintf(f, "\t\t{ID: %d, Key: %q, Value: %s},\n",
				entry.ID, entry.Key, formatString(entry.Value)); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(f, "\t},\n"); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(f, "}\n\n"); err != nil {
		return err
	}
	return nil
}

func writeByIDMap(f *os.File, data map[int]tbl.TblEntry) error {
	if _, err := fmt.Fprintf(f, "var TblDataByID = map[int]TblEntry{\n"); err != nil {
		return err
	}

	// Sort IDs for consistent output
	ids := make([]int, 0, len(data))
	for id := range data {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	for _, id := range ids {
		entry := data[id]
		if _, err := fmt.Fprintf(f, "\t%d: {ID: %d, Key: %q, Value: %s},\n",
			id, entry.ID, entry.Key, formatString(entry.Value)); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(f, "}\n\n"); err != nil {
		return err
	}
	return nil
}

func writeByValueMap(f *os.File, data map[string][]int) error {
	if _, err := fmt.Fprintf(f, "var TblDataByValue = map[string][]int{\n"); err != nil {
		return err
	}

	// Sort values for consistent output
	values := make([]string, 0, len(data))
	for value := range data {
		values = append(values, value)
	}
	sort.Strings(values)

	for _, value := range values {
		ids := data[value]
		// Sort IDs within each value array
		sort.Ints(ids)
		// Remove the redundant []int type
		if _, err := fmt.Fprintf(f, "\t%s: %#v[1:],\n",
			formatString(value), append([]int{0}, ids...)); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(f, "}\n"); err != nil {
		return err
	}
	return nil
}

// formatString prepares a string value for Go source code
func formatString(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "`", "`+\"`\"+`")
	return "`" + value + "`"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <input.tbl> [output.go]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := "string_tbl_data.go"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	parser, err := tbl.NewTblParser(inputFile)
	if err != nil {
		fmt.Printf("Error creating parser: %v\n", err)
		return
	}

	if err := parser.Parse(); err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}

	if err := generateGoFile(parser, outputFile); err != nil {
		fmt.Printf("Error generating Go file: %v\n", err)
		return
	}

	parser.CloseAndCleanup()
	fmt.Printf("Successfully generated %s\n", outputFile)
}
