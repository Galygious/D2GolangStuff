package main

import (
	"tbl-to-json/tbl"
)

var expansionstring_TblDataByKey = map[string][]tbl.TblEntry{
	"0sc": {
		{ID: 136, Key: "0sc", Value: `Scroll of Knowledge`},
	},
	"6bs": {
		{ID: 1412, Key: "6bs", Value: `Shillelagh`},
	},
	"6cb": {
		{ID: 1411, Key: "6cb", Value: `Great Bow`},
	},
	"6cs": {
		{ID: 1428, Key: "6cs", Value: `Elder Staff`},
	},
	"6hb": {
		{ID: 1491, Key: "6hb", Value: `Blade Bow`},
	},
	"6hx": {
		{ID: 1513, Key: "6hx", Value: `Colossus Crossbow`},
	},
	"6l7": {
		{ID: 1512, Key: "6l7", Value: `Crusader Bow`},
	},
	"6lb": {
		{ID: 1555, Key: "6lb", Value: `Shadow Bow`},
	},
	"6ls": {
		{ID: 1572, Key: "6ls", Value: `Stalagmite`},
	},
	"6lw": {
		{ID: 1576, Key: "6lw", Value: `Hydra Bow`},
	},
	"6lx": {
		{ID: 1577, Key: "6lx", Value: `Pellet Bow`},
	},
	"6mx": {
		{ID: 1593, Key: "6mx", Value: `Gorgon Crossbow`},
	},
	"6rx": {
		{ID: 1673, Key: "6rx", Value: `Demon Crossbow`},
	},
	"6s7": {
		{ID: 1630, Key: "6s7", Value: `Diamond Bow`},
	},
	"6sb": {
		{ID: 1670, Key: "6sb", Value: `Spider Bow`},
	},
	"6ss": {
		{ID: 1685, Key: "6ss", Value: `Walking Stick`},
	},
	"6sw": {
		{ID: 1689, Key: "6sw", Value: `Ward Bow`},
	},
	"6ws": {
		{ID: 1749, Key: "6ws", Value: `Archon Staff`},
	},
	"72a": {
		{ID: 883, Key: "72a", Value: `Ettin Axe`},
	},
	"72h": {
		{ID: 890, Key: "72h", Value: `Legend Sword`},
	},
	"7ar": {
		{ID: 1651, Key: "7ar", Value: `Suwayyah`},
	},
	"7ax": {
		{ID: 1657, Key: "7ax", Value: `Small Crescent`},
	},
	"7b7": {
		{ID: 1608, Key: "7b7", Value: `Champion Sword`},
	},
	"7b8": {
		{ID: 1609, Key: "7b8", Value: `Winged Axe`},
	},
	"7ba": {
		{ID: 1650, Key: "7ba", Value: `Silver-edged Axe`},
	},
	"7bk": {
		{ID: 1660, Key: "7bk", Value: `Winged Knife`},
	},
	"7bl": {
		{ID: 1661, Key: "7bl", Value: `Legend Spike`},
	},
	"7br": {
		{ID: 1667, Key: "7br", Value: `Mancatcher`},
	},
	"7bs": {
		{ID: 1668, Key: "7bs", Value: `Conquest Sword`},
	},
	"7bt": {
		{ID: 1669, Key: "7bt", Value: `Decapitator`},
	},
	"7bw": {
		{ID: 1672, Key: "7bw", Value: `Lich Wand`},
	},
	"7cl": {
		{ID: 1677, Key: "7cl", Value: `Truncheon`},
	},
	"7cm": {
		{ID: 1678, Key: "7cm", Value: `Highland Blade`},
	},
	"7cr": {
		{ID: 1683, Key: "7cr", Value: `Phase Blade`},
	},
	"7cs": {
		{ID: 1684, Key: "7cs", Value: `Battle Cestus`},
	},
	"7dg": {
		{ID: 1688, Key: "7dg", Value: `Bone Knife`},
	},
	"7di": {
		{ID: 1690, Key: "7di", Value: `Mithril Point`},
	},
	"7fb": {
		{ID: 1715, Key: "7fb", Value: `Colossus Sword`},
	},
	"7fc": {
		{ID: 1716, Key: "7fc", Value: `Hydra Edge`},
	},
	"7fl": {
		{ID: 1725, Key: "7fl", Value: `Scourge`},
	},
	"7ga": {
		{ID: 1730, Key: "7ga", Value: `Champion Axe`},
	},
	"7gd": {
		{ID: 1733, Key: "7gd", Value: `Colossus Blade`},
	},
	"7gi": {
		{ID: 1738, Key: "7gi", Value: `Glorious Axe`},
	},
	"7gl": {
		{ID: 1743, Key: "7gl", Value: `Ghost Glaive`},
	},
	"7gm": {
		{ID: 1742, Key: "7gm", Value: `Thunder Maul`},
	},
	"7gs": {
		{ID: 1748, Key: "7gs", Value: `Balrog Blade`},
	},
	"7gw": {
		{ID: 1752, Key: "7gw", Value: `Unearthed Wand`},
	},
	"7h7": {
		{ID: 1706, Key: "7h7", Value: `Great Poleaxe`},
	},
	"7ha": {
		{ID: 1746, Key: "7ha", Value: `Tomahawk`},
	},
	"7ja": {
		{ID: 1778, Key: "7ja", Value: `Hyperion Javelin`},
	},
	"7kr": {
		{ID: 1811, Key: "7kr", Value: `Fanged Knife`},
	},
	"7la": {
		{ID: 1810, Key: "7la", Value: `Feral Axe`},
	},
	"7ls": {
		{ID: 1828, Key: "7ls", Value: `Cryptic Sword`},
	},
	"7lw": {
		{ID: 1832, Key: "7lw", Value: `Feral Claws`},
	},
	"7m7": {
		{ID: 1785, Key: "7m7", Value: `Ogre Maul`},
	},
	"7ma": {
		{ID: 1826, Key: "7ma", Value: `Reinforced Mace`},
	},
	"7mp": {
		{ID: 1841, Key: "7mp", Value: `War Spike`},
	},
	"7mt": {
		{ID: 1845, Key: "7mt", Value: `Devil Star`},
	},
	"7o7": {
		{ID: 1816, Key: "7o7", Value: `Ogre Axe`},
	},
	"7p7": {
		{ID: 1833, Key: "7p7", Value: `War Pike`},
	},
	"7pa": {
		{ID: 1886, Key: "7pa", Value: `Cryptic Axe`},
	},
	"7pi": {
		{ID: 1883, Key: "7pi", Value: `Stygian Pilum`},
	},
	"7qr": {
		{ID: 1907, Key: "7qr", Value: `Scissors Suwayyah`},
	},
	"7qs": {
		{ID: 1908, Key: "7qs", Value: `Seraph Rod`},
	},
	"7s7": {
		{ID: 1884, Key: "7s7", Value: `Balrog Spear`},
	},
	"7s8": {
		{ID: 1885, Key: "7s8", Value: `Thresher`},
	},
	"7sb": {
		{ID: 1923, Key: "7sb", Value: `Elegant Blade`},
	},
	"7sc": {
		{ID: 1924, Key: "7sc", Value: `Mighty Scepter`},
	},
	"7sm": {
		{ID: 1934, Key: "7sm", Value: `Ataghan`},
	},
	"7sp": {
		{ID: 1938, Key: "7sp", Value: `Tyrant Club`},
	},
	"7sr": {
		{ID: 1942, Key: "7sr", Value: `Hyperion Spear`},
	},
	"7ss": {
		{ID: 1940, Key: "7ss", Value: `Falcata`},
	},
	"7st": {
		{ID: 1943, Key: "7st", Value: `Ghost Spear`},
	},
	"7ta": {
		{ID: 1941, Key: "7ta", Value: `Flying Axe`},
	},
	"7tk": {
		{ID: 1954, Key: "7tk", Value: `Flying Knife`},
	},
	"7tr": {
		{ID: 1955, Key: "7tr", Value: `Stygian Pike`},
	},
	"7ts": {
		{ID: 1956, Key: "7ts", Value: `Winged Harpoon`},
	},
	"7tw": {
		{ID: 1960, Key: "7tw", Value: `Runic Talons`},
	},
	"7vo": {
		{ID: 1984, Key: "7vo", Value: `Colossus Voulge`},
	},
	"7wa": {
		{ID: 1986, Key: "7wa", Value: `Beserker Axe`},
	},
	"7wb": {
		{ID: 1987, Key: "7wb", Value: `Wrist Sword`},
	},
	"7wc": {
		{ID: 1988, Key: "7wc", Value: `Giant Thresher`},
	},
	"7wd": {
		{ID: 1989, Key: "7wd", Value: `Mythical Sword`},
	},
	"7wh": {
		{ID: 1993, Key: "7wh", Value: `Legendary Mallet`},
	},
	"7wn": {
		{ID: 1999, Key: "7wn", Value: `Polished Wand`},
	},
	"7ws": {
		{ID: 2004, Key: "7ws", Value: `Caduceus`},
	},
	"7xf": {
		{ID: 2007, Key: "7xf", Value: `War Fist`},
	},
	"7yw": {
		{ID: 2040, Key: "7yw", Value: `Ghost Wand`},
	},
	"9ar": {
		{ID: 2163, Key: "9ar", Value: `Quhab`},
	},
	"9cs": {
		{ID: 2196, Key: "9cs", Value: `Hand Scythe`},
	},
	"9lw": {
		{ID: 2344, Key: "9lw", Value: `Greater Claws`},
	},
	"9qr": {
		{ID: 2419, Key: "9qr", Value: `Scissors Quhab`},
	},
	"9tw": {
		{ID: 2472, Key: "9tw", Value: `Greater Talons`},
	},
	"9wb": {
		{ID: 2499, Key: "9wb", Value: `Wrist Spike`},
	},
	"9xf": {
		{ID: 2519, Key: "9xf", Value: `Fascia`},
	},
	"A4Q2ExpansionSuccessCain": {
		{ID: 104, Key: "A4Q2ExpansionSuccessCain", Value: `40
I knew there was great potential in you, 
my friend. You've done a fantastic job.
 
Though my ancestors often struggled 
against the Three Evils and their 
minions, I've always lived a shut-in, 
scholarly life. I'm glad that my wisdom 
aided you.
 
Now, I wish to leave this place. Though 
Heaven's Gates are a marvel to behold, 
I hope I won't have to see them again 
for many, many years.
 
Please talk to Tyrael about leaving this 
place now!
`},
	},
	"A4Q2ExpansionSuccessTyrael": {
		{ID: 753, Key: "A4Q2ExpansionSuccessTyrael", Value: `40
Praise be to the Light! You have 
accomplished the impossible!
 
Diablo and Mephisto have been 
banished back into the Black Abyss 
that spawned them, and the corrupted 
Soulstones are no more.
 
However, while you were fighting here, 
Baal remained behind in the mortal 
realm, building an army of hellish 
minions. Now, Baal's army is searching 
for the Worldstone, the ancient source 
of all the Soulstones and their power, 
while leaving behind a wake of 
destruction. They have forged deeply 
into the Barbarian homelands, heading 
directly for the summit of Mount 
Arreat!
 
Baal knows, mortal hero! That is the 
very site of the blessed Worldstone!
 
Now, enter the portal I have opened for 
you. It will take you to the Barbarian 
city of Harrogath, the last bastion of 
Order on the slopes of Arreat.
`},
	},
	"A5Q1AfterInitAnya": {
		{ID: 1889, Key: "A5Q1AfterInitAnya", Value: `84
You've proven your skill at rescuing fair 
maidens...but how are you at killing 
vicious beasts?
`},
	},
	"A5Q1AfterInitCain": {
		{ID: 1078, Key: "A5Q1AfterInitCain", Value: `83
I believe that stopping the siege on 
Harrogath is the only way for you to 
earn the trust of these people.
`},
	},
	"A5Q1AfterInitLarzuk": {
		{ID: 1472, Key: "A5Q1AfterInitLarzuk", Value: `150
Uh...Did I mention there were catapults?
`},
	},
	"A5Q1AfterInitMalah": {
		{ID: 2444, Key: "A5Q1AfterInitMalah", Value: `77
Qual-Kehk and his men have been 
fighting to break the siege for some 
time. Where many have failed, you may 
succeed.
`},
	},
	"A5Q1AfterInitNihlathak": {
		{ID: 2702, Key: "A5Q1AfterInitNihlathak", Value: `70
After so many have died, who are you 
to think you can accomplish what our 
proud warriors could not?
`},
	},
	"A5Q1AfterInitQualKehk": {
		{ID: 2350, Key: "A5Q1AfterInitQualKehk", Value: `66
About to face Shenk the Overseer and 
stop the siege, are you? You should 
ask Malah to perform your last rites 
before you go, stranger.
`},
	},
	"A5Q1EarlyReturnAnya": {
		{ID: 1710, Key: "A5Q1EarlyReturnAnya", Value: `66
Those demons have been out there 
since before your arrival. Can you do 
nothing to stop them?
 
Your task is a simple one, warrior. Find 
Shenk and destroy him.
`},
	},
	"A5Q1EarlyReturnCain": {
		{ID: 730, Key: "A5Q1EarlyReturnCain", Value: `105
I understand your reluctance, but now 
is the time to strike.
`},
	},
	"A5Q1EarlyReturnLarzuk": {
		{ID: 1073, Key: "A5Q1EarlyReturnLarzuk", Value: `90
What's the matter, hero? Questioning 
your fortitude? I know we are.
`},
	},
	"A5Q1EarlyReturnMalah": {
		{ID: 219, Key: "A5Q1EarlyReturnMalah", Value: `64
This may not be very encouraging, but 
you have fared better than the others 
who passed through those gates.
 
In regards to Shenk the Overseer, 
remember: a general is nothing without 
his men.
`},
	},
	"A5Q1EarlyReturnNihlathak": {
		{ID: 688, Key: "A5Q1EarlyReturnNihlathak", Value: `78
What are you still doing here? I thought 
you were going off to die.
 
Go...Be quick about it.
`},
	},
	"A5Q1EarlyReturnQualKehk": {
		{ID: 2267, Key: "A5Q1EarlyReturnQualKehk", Value: `93
So, you still live. You're either quick or a 
coward.
`},
	},
	"A5Q1InitLarzuk": {
		{ID: 926, Key: "A5Q1InitLarzuk", Value: `45
If you're here to defeat Baal, you must 
prove it!
 
As we speak, Harrogath is under siege 
by Baal's demons. Catapults rain death 
just outside the town walls.
 
Baal himself travels up the sacred 
mountain, having left in charge here 
one of his most vicious generals, Shenk 
the Overseer. A ruthless taskmaster, 
he lashes his own minions into suicidal 
frenzies on the battlefield.
 
If you wish to prove yourself to us, 
destroy the monster, Shenk, that 
commands those infernal catapults 
outside Harrogath.  If you manage to 
do this, return to me.
`},
	},
	"A5Q1SuccessfulAnya": {
		{ID: 1705, Key: "A5Q1SuccessfulAnya", Value: `108
I was starting to wonder how long it 
would take you to stop those beasts.
 
Good job.
`},
	},
	"A5Q1SuccessfulCain": {
		{ID: 2808, Key: "A5Q1SuccessfulCain", Value: `60
Those catapults were like nothing I have 
ever seen before. You have prevailed 
against Shenk, my friend, but Baal is 
still far ahead of you.
`},
	},
	"A5Q1SuccessfulLarzuk": {
		{ID: 2549, Key: "A5Q1SuccessfulLarzuk", Value: `63
You're an even greater warrior than I 
expected...Sorry for underestimating 
you.
 
As a token of my appreciation, I will 
craft sockets into an item of your 
choosing, and from now on, you'll get 
the best price for all my wares.
`},
	},
	"A5Q1SuccessfulMalah": {
		{ID: 1409, Key: "A5Q1SuccessfulMalah", Value: `94
Oh...At last, the siege is ended. You 
truly are an angel.
 
I thank you.
`},
	},
	"A5Q1SuccessfulNihlathak": {
		{ID: 2704, Key: "A5Q1SuccessfulNihlathak", Value: `65
Ending the siege does not earn 
immediate respect, outsider.
 
Respect only comes with sacrifice -- 
something I'm sure you know nothing 
of.
`},
	},
	"A5Q1SuccessfulQualKehk": {
		{ID: 1534, Key: "A5Q1SuccessfulQualKehk", Value: `68
So...You managed to stop the siege.
 
You are more powerful than I gave you 
credit for. You have rightfully earned 
my respect.
`},
	},
	"A5Q2AfterInitAnya": {
		{ID: 1905, Key: "A5Q2AfterInitAnya", Value: `85
If those men are being treated like I 
was, you must find them. They won't 
survive much longer.
`},
	},
	"A5Q2AfterInitCain": {
		{ID: 1094, Key: "A5Q2AfterInitCain", Value: `74
I know firsthand that captivity is a sad 
fate for a man. Find them quickly.
`},
	},
	"A5Q2AfterInitLarzuk": {
		{ID: 2749, Key: "A5Q2AfterInitLarzuk", Value: `73
I crafted the swords and armor for 
Qual-Kehk's men. They were like 
brothers to me. I can't imagine the pain 
they must be suffering.
 
Save them if you can!
`},
	},
	"A5Q2AfterInitMalah": {
		{ID: 2188, Key: "A5Q2AfterInitMalah", Value: `84
Qual-Kehk's men have been imprisoned 
for some time. They are certain to be 
tired and weak.
 
You must protect them once you free 
them.
`},
	},
	"A5Q2AfterInitNihlathak": {
		{ID: 1355, Key: "A5Q2AfterInitNihlathak", Value: `94
You have proven you can take life, 
warrior, but can you save it as well?
`},
	},
	"A5Q2AfterInitQualKehk": {
		{ID: 2258, Key: "A5Q2AfterInitQualKehk", Value: `85
Those of my men fortunate enough to 
escape on their own told me that they 
were held captive in the highlands and 
plateaus.
`},
	},
	"A5Q2EarlyReturnAnya": {
		{ID: 434, Key: "A5Q2EarlyReturnAnya", Value: `115
Well, you found me on the mountain; I'm 
sure you'll find them too.
`},
	},
	"A5Q2EarlyReturnCain": {
		{ID: 1741, Key: "A5Q2EarlyReturnCain", Value: `63
If you are having trouble finding 
Qual-Kehk's soldiers, you should talk to 
Malah. She healed those who made it 
back before. Perhaps she would have 
some advice.
`},
	},
	"A5Q2EarlyReturnLarzuk": {
		{ID: 1348, Key: "A5Q2EarlyReturnLarzuk", Value: `76
As a kid, I used to play soldier among 
the barricades on the mountain. 
There's no easy way through that maze 
of walls.
`},
	},
	"A5Q2EarlyReturnMalah": {
		{ID: 2338, Key: "A5Q2EarlyReturnMalah", Value: `67
Soldiers who made it back told of a 
system of barricades placed among the 
mountain passes. They said that is 
where the prisoners are held.
`},
	},
	"A5Q2EarlyReturnNihlathak": {
		{ID: 2738, Key: "A5Q2EarlyReturnNihlathak", Value: `52
Did you ever stop to think why these 
demons are capturing Qual-Kehk's 
men? Why they are attacking us? Have 
you considered what Baal wants with 
the mountain?
 
No. You've not. You have no idea what 
you are dealing with.
`},
	},
	"A5Q2EarlyReturnQualKehk": {
		{ID: 2219, Key: "A5Q2EarlyReturnQualKehk", Value: `95
They say that discretion, not 
procrastination, is the better part of 
valor.
`},
	},
	"A5Q2EarlyReturnQualKehkMan": {
		{ID: 1069, Key: "A5Q2EarlyReturnQualKehkMan", Value: `78
More of my men are still alive out there. 
I am certain of it!
 
Find them. Free them from their cages 
and bring them back to me.
`},
	},
	"A5Q2InitQualKehk": {
		{ID: 2750, Key: "A5Q2InitQualKehk", Value: `58
My concerns have turned to my men 
taken prisoner on the battlefield by 
Baal's demons. I hate to think what's 
happened to them.
 
As you journey up the mountain, keep 
your eyes open for my soldiers and 
bring them back to me if you can.
`},
	},
	"A5Q2SuccessfulAnya": {
		{ID: 1449, Key: "A5Q2SuccessfulAnya", Value: `107
I'm sure those men appreciate the 
freedom you gave them...as much as I 
do.
`},
	},
	"A5Q2SuccessfulCain": {
		{ID: 2552, Key: "A5Q2SuccessfulCain", Value: `80
You've become a hero to this town, my 
friend. The shadows have lifted ever 
since you brought the Light to 
Harrogath.
`},
	},
	"A5Q2SuccessfulLarzuk": {
		{ID: 430, Key: "A5Q2SuccessfulLarzuk", Value: `55
Since your arrival, Cain has spoken of 
your deeds in the Southern Kingdoms, 
defeating both Mephisto and Diablo. At 
first, I scoffed at his tales, but I'm 
finding them more believable with every 
passing day.
`},
	},
	"A5Q2SuccessfulMalah": {
		{ID: 133, Key: "A5Q2SuccessfulMalah", Value: `73
You have inspired the people in this 
town -- not only those you rescued, but 
those you've helped in other ways as 
well.
`},
	},
	"A5Q2SuccessfulNihlathak": {
		{ID: 2720, Key: "A5Q2SuccessfulNihlathak", Value: `72
So. You brought the lost sheep home to 
the shepherd. Well done.
`},
	},
	"A5Q2SuccessfulQualKehk": {
		{ID: 188, Key: "A5Q2SuccessfulQualKehk", Value: `51
Thank you for rescuing my men. They 
have spoken well of your bravery in 
battle. Perhaps there is hope for us 
after all. 
 
If you wish, you may hire some of my 
mercenaries that you saved. And 
please...take this set of runes. I had 
been saving them for a socketed shield, 
but I think you'll make better use of 
them.
 
Be sure to set them in the right order 
for their fullest effect.
`},
	},
	"A5Q3AfterInitCain": {
		{ID: 1110, Key: "A5Q3AfterInitCain", Value: `59
I would listen to Malah. Nihlathak 
speaks with a venomous tongue and 
acts as if the entire weight of this town 
rests upon his shoulders.
 
Perhaps there is more going on here 
than we know.
`},
	},
	"A5Q3AfterInitLarzuk": {
		{ID: 1737, Key: "A5Q3AfterInitLarzuk", Value: `57
Anya is an amazing alchemist, 
especially for her young age. As long 
as I've known her, she's never let 
anything stop her from pursuing what 
she believed in. 
 
I wouldn't doubt that Nihlathak is 
involved. Ever since her father died, 
they haven't gotten along.
`},
	},
	"A5Q3AfterInitMalah": {
		{ID: 138, Key: "A5Q3AfterInitMalah", Value: `94
When you talk to Nihlathak, be careful. 
There is no telling what he will say or 
do.
`},
	},
	"A5Q3AfterInitNihlathak": {
		{ID: 2577, Key: "A5Q3AfterInitNihlathak", Value: `41
Anya! Who have you been talking to? 
Likely it was that meddling Malah. 
 
Well, I'll tell you what really happened. 
Anya came to me for guidance, after 
receiving a vision that her mother and 
younger brother were trapped in the 
lands beyond the Ice Caves. She had 
decided to go rescue them. 
 
I told her that her quest was a foolish 
one and that she would be safer 
staying within the city walls. However, 
she is a willful girl and would not listen 
to me. 
 
The next morning, she was gone. No 
one is more distraught than I over 
losing her. 
 
However, if you feel the need to be 
Malah's errand child, I won't try to stop 
you.
`},
	},
	"A5Q3AfterInitQualKehk": {
		{ID: 2166, Key: "A5Q3AfterInitQualKehk", Value: `43
Anya's father, Aust, was our wisest 
Elder. He was killed along with the 
other Elders who erected the ward that 
protects this city. The ward has kept 
Baal's demons out of Harrogath, but at 
a costly sacrifice. 
 
Nihlathak, on the other hand, was the 
only Elder to escape the demons. 
Somehow, he alone managed to find 
sanctuary, while the others died 
around him.
 
Ever since that day, Nihlathak and Anya 
have been at odds.
`},
	},
	"A5Q3EarlyReturnCain": {
		{ID: 465, Key: "A5Q3EarlyReturnCain", Value: `60
Nihlathak's story does sound 
reasonable, considering what I've 
heard about Anya. However, the best 
lies are often hidden within truth.
`},
	},
	"A5Q3EarlyReturnLarzuk": {
		{ID: 1257, Key: "A5Q3EarlyReturnLarzuk", Value: `56
As the daughter of Elder Aust, Anya is 
the only person, besides Nihlathak, 
who has any real knowledge of Mount 
Arreat's secrets. I'd hate to see our 
fate in the hands of Nihlathak alone.
`},
	},
	"A5Q3EarlyReturnMalah": {
		{ID: 1639, Key: "A5Q3EarlyReturnMalah", Value: `52
If it were anyone else, I would assume 
her dead. However, Anya is not so 
easily beaten. I know she must be alive.
`},
	},
	"A5Q3EarlyReturnNihlathak": {
		{ID: 176, Key: "A5Q3EarlyReturnNihlathak", Value: `46
Look, I've told you! She's dead! If you 
knew what was good for you, you'd 
concentrate your efforts on saving 
Harrogath -- not on lost causes like 
Anya.
`},
	},
	"A5Q3EarlyReturnQualKehk": {
		{ID: 2237, Key: "A5Q3EarlyReturnQualKehk", Value: `81
It seems like everyone feels Nihlathak 
played a part in Anya's disappearance.
 
Why would he do such a thing?
`},
	},
	"A5Q3FoundAnyaAnya": {
		{ID: 277, Key: "A5Q3FoundAnyaAnya", Value: `50
Hero. Nihlathak did this to me!
 
If you've come to help me, my only hope 
lies with Malah.
 
Please...Tell her you've found me...
`},
	},
	"A5Q3FoundAnyaCain": {
		{ID: 2608, Key: "A5Q3FoundAnyaCain", Value: `54
Goodness! Anya frozen by that fallen 
Barbarian, Nihlathak...Perhaps Malah 
can help you where I cannot.
`},
	},
	"A5Q3FoundAnyaLarzuk": {
		{ID: 1143, Key: "A5Q3FoundAnyaLarzuk", Value: `80
Poor Anya! I should've known Nihlathak 
was a traitor...
 
If you find him alive, kill him for me!
`},
	},
	"A5Q3FoundAnyaMalah": {
		{ID: 1573, Key: "A5Q3FoundAnyaMalah", Value: `56
So! That snake Nihlathak was behind 
Anya's disappearance...and he trapped 
her with a freezing curse.
 
Here. Take this potion to Anya and give 
it to her. That should release her.
`},
	},
	"A5Q3FoundAnyaQualKehk": {
		{ID: 1316, Key: "A5Q3FoundAnyaQualKehk", Value: `70
The snake has slipped our grasp! While 
you were gone, Nihlathak disappeared.
 
I'll bet Anya knows how to track him 
down.
`},
	},
	"A5Q3InitMalah": {
		{ID: 45, Key: "A5Q3InitMalah", Value: `43
There is a matter which I hesitate to 
share, but I believe you are the only 
one who can help me now.
 
Anya, the young alchemist and 
daughter to one of our slain Elders, 
has been missing for some time. She is 
a strong, crafty woman with a spirit 
like no other.
 
One night, just before your arrival, I 
overheard her and Nihlathak arguing 
about her father's death. The next 
morning she was gone.
 
Nihlathak has his own tale as to where 
she went and why. Don't believe him! I 
fear he is at the root of her 
disappearance.
 
Please, if you can, search for Anya and 
bring her back to us. She'll know what 
to do about Nihlathak.
`},
	},
	"A5Q3SuccessfulAnya": {
		{ID: 1194, Key: "A5Q3SuccessfulAnya", Value: `80
Thank you, hero, for rescuing me.
 
To show my personal gratitude, I give 
you this. I had it custom-made for you 
by Larzuk.
`},
	},
	"A5Q3SuccessfulCain": {
		{ID: 1273, Key: "A5Q3SuccessfulCain", Value: `62
For one so young, Anya commands 
great respect. Now that she is here, I 
will make it a point to talk to her about 
Mount Arreat.
 
You should do the same.
`},
	},
	"A5Q3SuccessfulLarzuk": {
		{ID: 1106, Key: "A5Q3SuccessfulLarzuk", Value: `67
I never liked Nihlathak, but I never 
suspected that he'd betray us! I just 
can't understand how an Elder could do 
such a thing.
`},
	},
	"A5Q3SuccessfulMalah": {
		{ID: 2204, Key: "A5Q3SuccessfulMalah", Value: `48
Thank you so much for bringing Anya 
back to us. I have devised this spell to 
increase your resistances as a token of 
my thanks. I know it isn't much, but I 
hope you find it helpful.
 
Please go talk to Anya. She has urgent 
news concerning Nihlathak.
`},
	},
	"A5Q3SuccessfulQualKehk": {
		{ID: 1410, Key: "A5Q3SuccessfulQualKehk", Value: `140
Your rescue of Anya was quite an 
accomplishment.
`},
	},
	"A5Q4AfterInitAnya": {
		{ID: 1937, Key: "A5Q4AfterInitAnya", Value: `90
As noble as Nihlathak's intentions are, 
nothing can excuse his actions.
`},
	},
	"A5Q4AfterInitCain": {
		{ID: 1126, Key: "A5Q4AfterInitCain", Value: `50
Regretfully, I know very little about this 
Relic. However, if what the others say 
is true, then Baal must not gain 
possession of it.
 
Stop Nihlathak...before all is lost.
`},
	},
	"A5Q4AfterInitLarzuk": {
		{ID: 196, Key: "A5Q4AfterInitLarzuk", Value: `56
Now, rescuing Anya -- that's good. 
Talking to me while Nihlathak hands 
over the Relic to Baal -- uh...that's bad!
`},
	},
	"A5Q4AfterInitMalah": {
		{ID: 2700, Key: "A5Q4AfterInitMalah", Value: `65
My worst fear has come true. Nihlathak 
has gone mad.
 
You must stop him, before he gives the 
Relic to the Lord of Destruction!
`},
	},
	"A5Q4AfterInitQualKehk": {
		{ID: 2074, Key: "A5Q4AfterInitQualKehk", Value: `60
I saw Nihlathak leave town just before 
you found Anya. He must be held 
accountable for his criminal deeds.
 
Find him and bring him back, if you can. 
Likely, he won't come willingly, and 
you'll be forced to kill him.
 
So be it.
`},
	},
	"A5Q4EarlyReturnAnya": {
		{ID: 699, Key: "A5Q4EarlyReturnAnya", Value: `55
Nihlathak is a traitor! If Baal gets the 
Relic, he shall enter the mountain and 
wreak havoc there!
 
I cannot believe that Nihlathak would 
give the Relic to Baal in any kind of 
trade. He's truly insane if he believes 
that he can deal with the Lord of 
Destruction.
`},
	},
	"A5Q4EarlyReturnCain": {
		{ID: 1476, Key: "A5Q4EarlyReturnCain", Value: `58
Ohh...This is a truly horrible turn of 
events.
 
I know it seems you have always been 
one step behind, my friend. But look at 
it this way...You have evil on the run.
`},
	},
	"A5Q4EarlyReturnLarzuk": {
		{ID: 2268, Key: "A5Q4EarlyReturnLarzuk", Value: `160
...What's there to talk about?
 
Kill Nihlathak!
`},
	},
	"A5Q4EarlyReturnMalah": {
		{ID: 987, Key: "A5Q4EarlyReturnMalah", Value: `63
Nihlathak was never the kindest man. 
But for him to betray the whole world...
 
Ahh...Where shall his soul finally rest?
`},
	},
	"A5Q4EarlyReturnQualKehk": {
		{ID: 2187, Key: "A5Q4EarlyReturnQualKehk", Value: `65
My advice is to go in quick and hit hard. 
Nihlathak can't be half as tough as the 
beasts you've faced out there.
`},
	},
	"A5Q4InitAnya": {
		{ID: 2341, Key: "A5Q4InitAnya", Value: `43
Nihlathak told me he struck a deal with 
Baal to protect Harrogath. In exchange 
for the demon's mercy, the misguided 
fool plans to give Baal the Relic of the 
Ancients, our most holy totem!
 
Doing so will allow Baal to enter Mount 
Arreat unchallenged by the Ancients. I 
tried to stop Nihlathak, but he 
imprisoned me in that icy tomb.
 
Nihlathak must be stopped before he 
dooms the whole world. As much as I 
would love to strangle the life out of 
him, I'm afraid I haven't the strength.
 
You must go to his lair through this 
portal I've opened, kill him, and then 
bring back the Relic of the Ancients.
 
Stop Nihlathak from destroying what we 
have striven for eons to protect.
`},
	},
	"A5Q4SuccessfulAnya": {
		{ID: 938, Key: "A5Q4SuccessfulAnya", Value: `55
You have stopped Nihlathak, but he 
didn't have the Relic! He must have 
already given it to Baal. Now, Baal will 
not be tested when he reaches Arreat's 
summit.
 
...Damn Nihlathak!
 
I do thank you for trying, though. 
Please, allow me to honor your courage 
by magically inscribing your name onto 
an item of your choosing. It's the least 
I can do.
`},
	},
	"A5Q4SuccessfulCain": {
		{ID: 1017, Key: "A5Q4SuccessfulCain", Value: `90
Beware! Baal grows stronger with every 
passing moment.
`},
	},
	"A5Q4SuccessfulLarzuk": {
		{ID: 1804, Key: "A5Q4SuccessfulLarzuk", Value: `93
Hmmm...What does Baal plan to do 
inside Mount Arreat?
`},
	},
	"A5Q4SuccessfulMalah": {
		{ID: 928, Key: "A5Q4SuccessfulMalah", Value: `65
So, the Relic is lost. Do not dwell on 
failures past. It is your future that 
matters more.
`},
	},
	"A5Q4SuccessfulQualKehk": {
		{ID: 63, Key: "A5Q4SuccessfulQualKehk", Value: `57
Nihlathak was a vile demon that shall 
find his home among the tortured 
minions of Hell!
 
You battled the Darkness without fear. I 
laud your skill and courage.
`},
	},
	"A5Q5AfterInitAnya": {
		{ID: 1953, Key: "A5Q5AfterInitAnya", Value: `72
You wouldn't have to test yourself 
against the Ancients if it weren't for 
Nihlathak's treachery. He was a fool 
and deserves to suffer for eternity.
`},
	},
	"A5Q5AfterInitCain": {
		{ID: 1142, Key: "A5Q5AfterInitCain", Value: `90
A test of mettle is a fitting rite of 
passage for a Barbarian hero.
`},
	},
	"A5Q5AfterInitLarzuk": {
		{ID: 2701, Key: "A5Q5AfterInitLarzuk", Value: `75
Every night, I've prayed to the Ancients 
to bring us peace...and now you must 
fight them.
 
Huh...Who shall I pray to now, warrior?
`},
	},
	"A5Q5AfterInitMalah": {
		{ID: 1420, Key: "A5Q5AfterInitMalah", Value: `65
By reaching the summit, you cease 
being just a simple warrior. When you 
come back, you will be far more.
`},
	},
	"A5Q5AfterInitQualKehk": {
		{ID: 636, Key: "A5Q5AfterInitQualKehk", Value: `63
The Ancients are not our enemies. 
Remember that! They are our 
ancestors -- our gods.
`},
	},
	"A5Q5EarlyReturnAnya": {
		{ID: 2240, Key: "A5Q5EarlyReturnAnya", Value: `71
Look. I must apologize.
 
I feel responsible for your current 
struggle. If I had only stopped 
Nihlathak from giving Baal the Relic, 
you would not have to fight those 
ghosts.
`},
	},
	"A5Q5EarlyReturnCain": {
		{ID: 200, Key: "A5Q5EarlyReturnCain", Value: `52
We have come too far to be defeated 
now, my friend. I have seen you 
complete many difficult quests. Though 
this may be your greatest trial, it is not 
beyond your reach.
`},
	},
	"A5Q5EarlyReturnLarzuk": {
		{ID: 2176, Key: "A5Q5EarlyReturnLarzuk", Value: `70
You've walked on the burial grounds of 
our greatest ancestors. I'm not sure 
whether I should bow before you or 
revile you for committing sacrilege.
 
Tread lightly when you walk with gods.
`},
	},
	"A5Q5EarlyReturnMalah": {
		{ID: 288, Key: "A5Q5EarlyReturnMalah", Value: `68
Do not doubt yourself. I believe you are 
worthy to pass through the Ancients' 
gates, but you must believe, as well.
`},
	},
	"A5Q5EarlyReturnQualKehk": {
		{ID: 2203, Key: "A5Q5EarlyReturnQualKehk", Value: `114
I warned you!
 
The Ancients are not like the demons 
you're accustomed to fighting.
`},
	},
	"A5Q5InitQualKehk": {
		{ID: 2028, Key: "A5Q5InitQualKehk", Value: `37
Every time I hear of you, warrior, your 
deeds become more legendary.
 
But take heed. You are approaching the 
very summit of Mount Arreat. I have 
never dared venture there. It is sacred 
-- our most holy place. The legends say 
it is guarded by the Ancient Ones, who 
block the path of all who are unworthy.
 
Your reputation here does not 
matter...It will be the Ancients who 
determine your worthiness.
 
Good luck.
`},
	},
	"A5Q5SuccessfulAnya": {
		{ID: 682, Key: "A5Q5SuccessfulAnya", Value: `82
You stand before me a worthy hero -- 
and on you rests the last hope of our 
people.
 
Bear it well, warrior.
`},
	},
	"A5Q5SuccessfulCain": {
		{ID: 1784, Key: "A5Q5SuccessfulCain", Value: `67
You have proven yourself to these 
people. They look to you as their 
warrior, their champion.
`},
	},
	"A5Q5SuccessfulLarzuk": {
		{ID: 2526, Key: "A5Q5SuccessfulLarzuk", Value: `82
The Ancients have honored you, and in 
turn, so do we. I have no remaining 
doubts about you, now.
`},
	},
	"A5Q5SuccessfulMalah": {
		{ID: 1939, Key: "A5Q5SuccessfulMalah", Value: `60
I knew the Ancients would find you 
worthy of Mount Arreat's secrets. Now, 
stop Baal before he destroys all that is 
sacred.
`},
	},
	"A5Q5SuccessfulQualKehk": {
		{ID: 1286, Key: "A5Q5SuccessfulQualKehk", Value: `75
Besting the Ancients in battle is a 
mighty feat indeed. I hope this means 
you're ready to battle Baal.
`},
	},
	"A5Q6EarlyReturnAnya": {
		{ID: 964, Key: "A5Q6EarlyReturnAnya", Value: `60
Baal has blocked Tyrael from entering 
the Worldstone Chamber? This truly 
has become a battle against Hell.
 
Whether or not it was the Heavens' 
decree, this is your fight now -- your 
destiny.
`},
	},
	"A5Q6EarlyReturnCain": {
		{ID: 1212, Key: "A5Q6EarlyReturnCain", Value: `51
Remember this. Baal once possessed Tal 
Rasha, one of the most powerful of the 
ancient Horadrim.
 
Your battles with Mephisto and Diablo 
will pale in comparison to your battle 
with Baal.
 
The Lord of Destruction aided by Tal 
Rasha's knowledge...The mountain 
itself will tremble when you clash.
`},
	},
	"A5Q6EarlyReturnLarzuk": {
		{ID: 2452, Key: "A5Q6EarlyReturnLarzuk", Value: `59
I may be just an armorer, but I know 
this...Baal plans to destroy the world 
with the secrets contained in that 
mountain. It doesn't take a genius to 
know he has to be stopped.
`},
	},
	"A5Q6EarlyReturnMalah": {
		{ID: 2407, Key: "A5Q6EarlyReturnMalah", Value: `70
You knew it would eventually come down 
to this. Kill Baal! Finish the game!
`},
	},
	"A5Q6EarlyReturnQualKehk": {
		{ID: 2155, Key: "A5Q6EarlyReturnQualKehk", Value: `63
You have ventured to a place beyond 
legend. You rush to face an evil few 
can even imagine.
 
Be careful, my friend, and may the Light 
watch over you.
`},
	},
	"A5Q6InitAncients": {
		{ID: 889, Key: "A5Q6InitAncients", Value: `31
You are a worthy hero! We augment 
your skill and grant you entry to the 
interior of Mount Arreat, wherein lies 
the Worldstone.
 
Beware. You will not be alone. Baal the 
Lord of Destruction is already inside. 
 
The Archangel Tyrael has always been 
our benefactor, but even he cannot 
help us now. For Baal blocks Tyrael's 
spiritual presence from entering the 
chamber of the Worldstone. Only you, 
mortal, have the power to defeat Baal 
now.
 
Baal threatens the Worldstone -- and 
through it, the mortal realm, itself. You 
must stop him before he gains full 
control of the sacred stone. With it 
under his control, Baal could shatter 
the boundaries between this world and 
the Burning Hells, thus allowing the 
hordes of the Prime Evils to pour forth 
into the mortal realm like an 
unstoppable tide!
 
If you are weak, the world as you know 
it could be lost forever. You must NOT 
fail!
`},
	},
	"A5Q6SuccessfulAnya": {
		{ID: 426, Key: "A5Q6SuccessfulAnya", Value: `53
You have done the impossible, hero. 
Your defeat of the last of the three 
Prime Evils is a great victory for the 
Light. 
 
Strange that you say that the 
Worldstone must be destroyed. The 
prophecies said nothing about that.
 
Perhaps all we have fought for will be 
lost...or perhaps we'll never need fight 
again!
`},
	},
	"A5Q6SuccessfulCain": {
		{ID: 1528, Key: "A5Q6SuccessfulCain", Value: `40
I knew in time you would defeat Baal. 
You have done everything you set out 
to do, my friend.
 
Ever since you rescued me from 
Tristram, I have believed in you. It has 
been a supreme honor to aid you along 
the way. 
 
So...The Worldstone was corrupted by 
Baal. And now Tyrael must destroy it. 
Worry not. Through whatever lies 
ahead I have faith that the Light will 
guide us both.
 
Go, now, back to the Worldstone 
chamber, and enter the portal Tyrael 
has opened for you.
`},
	},
	"A5Q6SuccessfulLarzuk": {
		{ID: 407, Key: "A5Q6SuccessfulLarzuk", Value: `74
The Ancients themselves will envy our 
songs about you.
 
Please, don't forget about us! Farewell, 
my friend.
`},
	},
	"A5Q6SuccessfulMalah": {
		{ID: 663, Key: "A5Q6SuccessfulMalah", Value: `48
If Tyrael says the Worldstone must be 
destroyed, then it must. We cannot let 
Baal's corruption prevail!
 
The world will change, true -- but who is 
to say it isn't for the better?
`},
	},
	"A5Q6SuccessfulQualKehk": {
		{ID: 2756, Key: "A5Q6SuccessfulQualKehk", Value: `60
The destruction of the Worldstone does 
not bode well for our world. But I'll try 
not to worry...
 
After all, we have warriors like you 
fighting for us and for the Light.
 
Farewell!
`},
	},
	"A5Q6SuccessfulTyrael": {
		{ID: 1642, Key: "A5Q6SuccessfulTyrael", Value: `40
I am impressed, mortal. You have 
overcome the greatest challenge this 
world has ever faced and defeated the 
last of the Prime Evils. However, we are 
too late to save the Worldstone. Baal's 
destructive touch has corrupted it 
completely.
 
Given enough time, the Worldstone's 
energies will drain away and the 
barriers between the worlds will 
shatter -- the powers of Hell will flood 
into this...Sanctuary...and eradicate 
your people and everything you've 
labored to build.
 
Therefore, I must destroy the corrupted 
Worldstone before the powers of Hell 
take root. This act will change your 
world forever -- with consequences 
even I cannot foresee. However, it is 
the only way to ensure mankind's 
survival.
 
Go now, mortal. I have opened a portal 
that will lead you to safety.
 
May the Eternal Light shine upon you 
and your descendants for what you've 
done this day. The continued survival 
of mankind is your legacy! Above all 
else, you have earned a rest from this 
endless battle.
`},
	},
	"Accursed": {
		{ID: 782, Key: "Accursed", Value: `Cursing`},
	},
	"Acrobat's": {
		{ID: 1898, Key: "Acrobat's", Value: `Acrobatic`},
	},
	"Act 5 Combatant": {
		{ID: 265, Key: "Act 5 Combatant", Value: `Combatant`},
	},
	"Act 5 Townguard": {
		{ID: 264, Key: "Act 5 Townguard", Value: `Guard`},
	},
	"Addsocketsui": {
		{ID: 837, Key: "Addsocketsui", Value: `add sockets`},
	},
	"Addsocketsui2": {
		{ID: 838, Key: "Addsocketsui2", Value: `Choose the item to which you wish to add sockets.`},
	},
	"Akarat's Protector": {
		{ID: 2806, Key: "Akarat's Protector", Value: `Akarat's Protector`},
	},
	"Aladdin's Eviserator": {
		{ID: 2627, Key: "Aladdin's Eviserator", Value: `Aladdin's Eviscerator`},
	},
	"Aldur's Advance": {
		{ID: 2073, Key: "Aldur's Advance", Value: `Aldur's Advance`},
	},
	"Aldur's Deception": {
		{ID: 877, Key: "Aldur's Deception", Value: `Aldur's Deception`},
	},
	"Aldur's Gauntlet": {
		{ID: 435, Key: "Aldur's Gauntlet", Value: `Aldur's Rhythm`},
	},
	"Aldur's Guantlet": {
		{ID: 2744, Key: "Aldur's Guantlet", Value: `Aldur's Gauntlet`},
	},
	"Aldur's Stony Gaze": {
		{ID: 1753, Key: "Aldur's Stony Gaze", Value: `Aldur's Stony Gaze`},
	},
	"Aldur's Watchtower": {
		{ID: 2006, Key: "Aldur's Watchtower", Value: `Aldur's Watchtower`},
	},
	"Alma Negra": {
		{ID: 2477, Key: "Alma Negra", Value: `Alma Negra`},
	},
	"Alma's Reflection": {
		{ID: 749, Key: "Alma's Reflection", Value: `Alma's Reflection`},
	},
	"Altar": {
		{ID: 470, Key: "Altar", Value: `Altar`},
	},
	"AmaOnly": {
		{ID: 1206, Key: "AmaOnly", Value: `(Amazon Only)`},
	},
	"Ambergris": {
		{ID: 152, Key: "Ambergris", Value: `Ambergris`},
	},
	"Amodeus's Manipulator": {
		{ID: 2170, Key: "Amodeus's Manipulator", Value: `Amodeus' Manipulator`},
	},
	"Ancient Barbarian 1": {
		{ID: 281, Key: "Ancient Barbarian 1", Value: `Talic`},
	},
	"Ancient Barbarian 2": {
		{ID: 280, Key: "Ancient Barbarian 2", Value: `Madawc`},
	},
	"Ancient Barbarian 3": {
		{ID: 278, Key: "Ancient Barbarian 3", Value: `Korlic`},
	},
	"Ancient Eye": {
		{ID: 1772, Key: "Ancient Eye", Value: `Ancient Eye`},
	},
	"Ancient Statue 1": {
		{ID: 1602, Key: "Ancient Statue 1", Value: `Talic the Defender`},
	},
	"Ancient Statue 2": {
		{ID: 1601, Key: "Ancient Statue 2", Value: `Madawc the Guardian`},
	},
	"Ancient Statue 3": {
		{ID: 1600, Key: "Ancient Statue 3", Value: `Korlic the Protector`},
	},
	"AncientsAct5IntroGossip1": {
		{ID: 1992, Key: "AncientsAct5IntroGossip1", Value: `35
We are the spirits of the Nephalem, the 
Ancient Ones. We have been chosen to 
guard sacred Mount Arreat, wherein 
the Worldstone rests. Few are worthy 
to stand in its presence; fewer still can 
comprehend its true purpose.
 
Before you enter, you must defeat us.
`},
	},
	"Andariel's Visage": {
		{ID: 937, Key: "Andariel's Visage", Value: `Andariel's Visage`},
	},
	"Angel's Song": {
		{ID: 814, Key: "Angel's Song", Value: `Angel's Song`},
	},
	"Annihilus": {
		{ID: 1074, Key: "Annihilus", Value: `Annihilus`},
	},
	"Anodized Elite": {
		{ID: 294, Key: "Anodized Elite", Value: `Hell's Belle`},
	},
	"Antimagic": {
		{ID: 2458, Key: "Antimagic", Value: `Antimagic`},
	},
	"AnyaAct5IntroGossip1": {
		{ID: 1821, Key: "AnyaAct5IntroGossip1", Value: `54
You have proven yourself a true hero to 
me and my people.
 
These are dark times, warrior. I hope 
you can bring an end to Baal's reign of 
destruction. 
 
Our Council of Elders is gone -- my 
father, Aust, among them. The one 
thing that keeps us from total despair 
is the promise of vengeance against 
Baal.
`},
	},
	"AnyaGossip1": {
		{ID: 1145, Key: "AnyaGossip1", Value: `45
Now that the Elders are all dead, I don't 
know who will guide our people through 
this dark time. I was to be next in line 
after my father, but this burden is too 
great for me to shoulder alone. 
 
We are a people of strong blood. I shall 
do what I can and let fate do the rest.
`},
	},
	"AnyaGossip10": {
		{ID: 1439, Key: "AnyaGossip10", Value: `60
Our people believe that the Ancients 
protecting Mount Arreat have the 
power to stop Baal. Unfortunately, the 
Lord of Destruction has shown great 
power to undermine our faith.
`},
	},
	"AnyaGossip2": {
		{ID: 1146, Key: "AnyaGossip2", Value: `82
Baal's minions are not to be trifled with. 
In their bloodlust they will sacrifice 
themselves to destroy you.
`},
	},
	"AnyaGossip3": {
		{ID: 1147, Key: "AnyaGossip3", Value: `60
Many outsiders believe that the 
fantastic stories about our ancestors, 
the Ancients, are but fables. However, I 
believe that the Ancients were more 
than human -- that mankind has fallen 
from what it once was.
`},
	},
	"AnyaGossip4": {
		{ID: 1148, Key: "AnyaGossip4", Value: `58
When I was a child, the Elders told us 
stories about the mountain and its 
power...and how the Barbarian people 
are bound to it as protectors.
 
Baal is not just taking our land -- with 
every step he takes up the mountain, 
he takes a part of who we are as a 
people.
`},
	},
	"AnyaGossip5": {
		{ID: 1149, Key: "AnyaGossip5", Value: `72
I am truly glad you are here, warrior. 
Perhaps things would be different now 
if we had asked for assistance from 
the neighboring kingdoms.
 
Our foolish, foolish pride...
`},
	},
	"AnyaGossip6": {
		{ID: 1150, Key: "AnyaGossip6", Value: `54
My father, Aust, was among those 
Elders wise enough to see that we 
would need outside help to deal with 
Baal's legions. He believed that this 
conflict would affect the entire world, 
not just our homeland. He said that 
Mount Arreat is as necessary to the 
world's survival as food and water is to 
our own.
 
I believe this to be true.
`},
	},
	"AnyaGossip7": {
		{ID: 1151, Key: "AnyaGossip7", Value: `50
Qual-Kehk's confidence in his abilities 
once bordered on arrogance, but in the 
early days of the siege, he was 
humbled by Baal. While he saved us 
from immediate destruction, a third of 
our warriors were lost.
 
None felt those losses more than he. 
Though he may not admit it, your 
presence gives him hope, warrior.
`},
	},
	"AnyaGossip8": {
		{ID: 1152, Key: "AnyaGossip8", Value: `78
If Larzuk could sing or dance half as 
well as he smiths, he'd be married by 
now.
 
...Just look at those shoulders.
`},
	},
	"AnyaGossip9": {
		{ID: 1153, Key: "AnyaGossip9", Value: `57
Nihlathak was the last of our Elders, 
whose charge it was to safeguard the 
mountain. He alone tried to guide us 
through the most desperate time in our 
history -- but he was just as helpless as 
the rest of us.
 
I cannot forgive his betrayal, but I can 
learn from it.
`},
	},
	"Apocrypha": {
		{ID: 2015, Key: "Apocrypha", Value: `Apocrypha`},
	},
	"Apothecary's Tote": {
		{ID: 748, Key: "Apothecary's Tote", Value: `Apothecary's Tote`},
	},
	"Arachnid Mesh": {
		{ID: 2075, Key: "Arachnid Mesh", Value: `Arachnid Mesh`},
	},
	"Aragon's Icy Stare": {
		{ID: 2043, Key: "Aragon's Icy Stare", Value: `Aragon's Icy Stare`},
	},
	"Aragon's Masterpiece": {
		{ID: 2027, Key: "Aragon's Masterpiece", Value: `Aragon's Masterpiece`},
	},
	"Aragon's Storm Cloud": {
		{ID: 2511, Key: "Aragon's Storm Cloud", Value: `Aragon's Storm Cloud`},
	},
	"Aragon's Sunfire": {
		{ID: 2014, Key: "Aragon's Sunfire", Value: `Aragon's Sunfire`},
	},
	"Arcadian": {
		{ID: 2491, Key: "Arcadian", Value: `Arcadian`},
	},
	"Archangel's Deliverance": {
		{ID: 2768, Key: "Archangel's Deliverance", Value: `Arch-angel's Deliverance`},
	},
	"Arcing": {
		{ID: 534, Key: "Arcing", Value: `Arcing`},
	},
	"Argent": {
		{ID: 1825, Key: "Argent", Value: `Argent`},
	},
	"Arioc's Needle": {
		{ID: 1570, Key: "Arioc's Needle", Value: `Arioc's Needle`},
	},
	"Arkaine's Valor": {
		{ID: 1028, Key: "Arkaine's Valor", Value: `Arkaine's Valor`},
	},
	"Arm of King Leoric": {
		{ID: 2158, Key: "Arm of King Leoric", Value: `Arm of King Leoric`},
	},
	"Armor": {
		{ID: 939, Key: "Armor", Value: `Armor`},
	},
	"Arreat Plateau": {
		{ID: 522, Key: "Arreat Plateau", Value: `Arreat Plateau`},
	},
	"Arreat's Face": {
		{ID: 143, Key: "Arreat's Face", Value: `Arreat's Face`},
	},
	"Artificer's": {
		{ID: 469, Key: "Artificer's", Value: `Artisan's`},
	},
	"Ashrera's Wired Frame": {
		{ID: 1445, Key: "Ashrera's Wired Frame", Value: `Asheara's Wired Frame`},
	},
	"AssOnly": {
		{ID: 1959, Key: "AssOnly", Value: `(Assassin Only)`},
	},
	"Assamic": {
		{ID: 1958, Key: "Assamic", Value: `Lunar`},
	},
	"Astral": {
		{ID: 2490, Key: "Astral", Value: `Astral`},
	},
	"Athena's Wrath": {
		{ID: 491, Key: "Athena's Wrath", Value: `Athena's Wrath`},
	},
	"Athlete's": {
		{ID: 2271, Key: "Athlete's", Value: `Athletic`},
	},
	"Atma's Scarab": {
		{ID: 533, Key: "Atma's Scarab", Value: `Atma's Scarab`},
	},
	"Atma's Wail": {
		{ID: 98, Key: "Atma's Wail", Value: `Atma's Wail`},
	},
	"Aureolin": {
		{ID: 2134, Key: "Aureolin", Value: `Aureolic`},
	},
	"Aurora's Guard": {
		{ID: 2017, Key: "Aurora's Guard", Value: `Aurora's Guard`},
	},
	"Axe Dweller": {
		{ID: 989, Key: "Axe Dweller", Value: `Axe Dweller`},
	},
	"Baal Crab": {
		{ID: 628, Key: "Baal Crab", Value: `Baal`},
	},
	"Baal Crab Clone": {
		{ID: 107, Key: "Baal Crab Clone", Value: `Baal`},
	},
	"Baal Crab to Stairs": {
		{ID: 1456, Key: "Baal Crab to Stairs", Value: `Baal`},
	},
	"Baal Subject 1": {
		{ID: 1615, Key: "Baal Subject 1", Value: `Colenzo the Annihilator`},
	},
	"Baal Subject 2": {
		{ID: 1616, Key: "Baal Subject 2", Value: `Achmel the Cursed`},
	},
	"Baal Subject 3": {
		{ID: 1617, Key: "Baal Subject 3", Value: `Bartuc the Bloody`},
	},
	"Baal Subject 4": {
		{ID: 1665, Key: "Baal Subject 4", Value: `Ventar the Unholy`},
	},
	"Baal Subject 5": {
		{ID: 1666, Key: "Baal Subject 5", Value: `Lister the Tormentor`},
	},
	"Baal Subject 6": {
		{ID: 1674, Key: "Baal Subject 6", Value: `The Butcher`},
	},
	"Baal Subject 6a": {
		{ID: 647, Key: "Baal Subject 6a", Value: `The Baker`},
	},
	"Baal Subject 6b": {
		{ID: 648, Key: "Baal Subject 6b", Value: `The Candlestick Maker`},
	},
	"Baal Subject Mummy": {
		{ID: 1050, Key: "Baal Subject Mummy", Value: `Unraveler`},
	},
	"Baal Taunt": {
		{ID: 541, Key: "Baal Taunt", Value: `Baal`},
	},
	"Baal Tentacle": {
		{ID: 1121, Key: "Baal Tentacle", Value: `Festering Appendages`},
	},
	"Baal Throne": {
		{ID: 698, Key: "Baal Throne", Value: `Baal`},
	},
	"BaalColdMage": {
		{ID: 738, Key: "BaalColdMage", Value: `Death Mage`},
	},
	"Baals Minion": {
		{ID: 2740, Key: "Baals Minion", Value: `Minion of Destruction`},
	},
	"Baezil's Vortex": {
		{ID: 456, Key: "Baezil's Vortex", Value: `Baezil's Vortex`},
	},
	"Baezils Vortex": {
		{ID: 2515, Key: "Baezils Vortex", Value: `Baezil's Vortex`},
	},
	"Bahamut's": {
		{ID: 1038, Key: "Bahamut's", Value: `Bahamut's`},
	},
	"BanishedSoul": {
		{ID: 1071, Key: "BanishedSoul", Value: `Banished Soul`},
	},
	"Banshee's Wail": {
		{ID: 2628, Key: "Banshee's Wail", Value: `Banshee's Wail`},
	},
	"BarOnly": {
		{ID: 1444, Key: "BarOnly", Value: `(Barbarian Only)`},
	},
	"Baranar's Star": {
		{ID: 1658, Key: "Baranar's Star", Value: `Baranar's Star`},
	},
	"Barricade Door": {
		{ID: 227, Key: "Barricade Door", Value: `Barricaded Door`},
	},
	"Barricade Tower": {
		{ID: 228, Key: "Barricade Tower", Value: `Barricaded Tower`},
	},
	"Barricade Wall Left": {
		{ID: 553, Key: "Barricade Wall Left", Value: `Barricade`},
	},
	"Barricade Wall Right": {
		{ID: 224, Key: "Barricade Wall Right", Value: `Barricade`},
	},
	"Bear": {
		{ID: 72, Key: "Bear", Value: `Bear`},
	},
	"BeginTaintedSunAss": {
		{ID: 486, Key: "BeginTaintedSunAss", Value: `An eclipse... never a good omen.`},
	},
	"BeginTaintedSunDru": {
		{ID: 216, Key: "BeginTaintedSunDru", Value: `Strange... an unexpected eclipse.`},
	},
	"Berserk": {
		{ID: 255, Key: "Berserk", Value: `Berserker`},
	},
	"BerserkSlayer": {
		{ID: 244, Key: "BerserkSlayer", Value: `Berserker Slayer`},
	},
	"Bing Sz Wang": {
		{ID: 603, Key: "Bing Sz Wang", Value: `Bing Sz Wang`},
	},
	"Black Hades": {
		{ID: 1192, Key: "Black Hades", Value: `Black Hades`},
	},
	"Blackbog's Sharp": {
		{ID: 587, Key: "Blackbog's Sharp", Value: `Blackbog's Sharp`},
	},
	"Blackhand Key": {
		{ID: 147, Key: "Blackhand Key", Value: `Blackhand Key`},
	},
	"Blackhorn": {
		{ID: 1918, Key: "Blackhorn", Value: `Blackhorn`},
	},
	"Blackhorn's Face": {
		{ID: 86, Key: "Blackhorn's Face", Value: `Blackhorn's Face`},
	},
	"Blackleach Blade": {
		{ID: 2057, Key: "Blackleach Blade", Value: `Blackleach Blade`},
	},
	"Blackoak Shield": {
		{ID: 2247, Key: "Blackoak Shield", Value: `Blackoak Shield`},
	},
	"Blade Creeper": {
		{ID: 62, Key: "Blade Creeper", Value: `Blade Sentinel`},
	},
	"Blade of Ali Baba": {
		{ID: 1172, Key: "Blade of Ali Baba", Value: `Blade of Ali Baba`},
	},
	"Blademaster": {
		{ID: 1671, Key: "Blademaster", Value: `Blade Master`},
	},
	"Blanched": {
		{ID: 1929, Key: "Blanched", Value: `Blanched`},
	},
	"Blaze Ripper": {
		{ID: 282, Key: "Blaze Ripper", Value: `Blaze Ripper`},
	},
	"Blazing": {
		{ID: 2600, Key: "Blazing", Value: `Blazing`},
	},
	"Blighting": {
		{ID: 1344, Key: "Blighting", Value: `Blighting`},
	},
	"Bling Bling": {
		{ID: 2000, Key: "Bling Bling", Value: `Bling Bling`},
	},
	"Blood Chalice": {
		{ID: 319, Key: "Blood Chalice", Value: `Blood Chalice`},
	},
	"Blood Comet": {
		{ID: 2244, Key: "Blood Comet", Value: `Blood Comet`},
	},
	"Blood Lord1": {
		{ID: 1011, Key: "Blood Lord1", Value: `Moon Lord`},
	},
	"Blood Lord2": {
		{ID: 1014, Key: "Blood Lord2", Value: `Night Lord`},
	},
	"Blood Lord3": {
		{ID: 1015, Key: "Blood Lord3", Value: `Blood Lord`},
	},
	"Blood Lord4": {
		{ID: 1016, Key: "Blood Lord4", Value: `Hell Lord`},
	},
	"Blood Lord5": {
		{ID: 1019, Key: "Blood Lord5", Value: `Death Lord`},
	},
	"Blood Rain": {
		{ID: 892, Key: "Blood Rain", Value: `Blood Rain`},
	},
	"Blood Temptress": {
		{ID: 114, Key: "Blood Temptress", Value: `Blood Temptress`},
	},
	"Blood Witch": {
		{ID: 117, Key: "Blood Witch", Value: `Blood Witch`},
	},
	"BloodBoss": {
		{ID: 169, Key: "BloodBoss", Value: `Blood Boss`},
	},
	"BloodBringer": {
		{ID: 166, Key: "BloodBringer", Value: `Blood Bringer`},
	},
	"Bloodletter": {
		{ID: 186, Key: "Bloodletter", Value: `Bloodletter`},
	},
	"Bloodmoon": {
		{ID: 777, Key: "Bloodmoon", Value: `Bloodmoon`},
	},
	"Bloodraven's Charge": {
		{ID: 811, Key: "Bloodraven's Charge", Value: `Blood Raven's Charge`},
	},
	"Bloodtree Stump": {
		{ID: 1708, Key: "Bloodtree Stump", Value: `Bloodtree Stump`},
	},
	"Bloody": {
		{ID: 1696, Key: "Bloody", Value: `Bloody`},
	},
	"Bloody Foothills": {
		{ID: 519, Key: "Bloody Foothills", Value: `Bloody Foothills`},
	},
	"Boneflame": {
		{ID: 2509, Key: "Boneflame", Value: `Boneflame`},
	},
	"Bonehew": {
		{ID: 1125, Key: "Bonehew", Value: `Bonehew`},
	},
	"Bonesaw Breaker": {
		{ID: 520, Key: "Bonesaw Breaker", Value: `Bonesaw Breaker`},
	},
	"Bonescapel": {
		{ID: 583, Key: "Bonescapel", Value: `Bonescalpel`},
	},
	"Boneshade": {
		{ID: 1755, Key: "Boneshade", Value: `Boneshade`},
	},
	"Boneslayer Blade": {
		{ID: 2054, Key: "Boneslayer Blade", Value: `Boneslayer Blade`},
	},
	"Boreal": {
		{ID: 2177, Key: "Boreal", Value: `Boreal`},
	},
	"BoundSpirit": {
		{ID: 84, Key: "BoundSpirit", Value: `Bound Spirit`},
	},
	"Bowyer's": {
		{ID: 1709, Key: "Bowyer's", Value: `Bowyer's`},
	},
	"Brown": {
		{ID: 2260, Key: "Brown", Value: `Brown`},
	},
	"Bul Katho's Wedding Band": {
		{ID: 2217, Key: "Bul Katho's Wedding Band", Value: `Bul-Kathos' Wedding Band`},
	},
	"Bul-Kathos": {
		{ID: 2418, Key: "Bul-Kathos", Value: `Bul-Kathos`},
	},
	"Bul-Kathos' Children": {
		{ID: 575, Key: "Bul-Kathos' Children", Value: `Bul-Kathos' Children`},
	},
	"Bul-Kathos' Custodian": {
		{ID: 2058, Key: "Bul-Kathos' Custodian", Value: `Bul-Kathos' Custodian`},
	},
	"Bul-Kathos' Might": {
		{ID: 46, Key: "Bul-Kathos' Might", Value: `Bul-Kathos' Might`},
	},
	"Bul-Kathos' Sacred Charge": {
		{ID: 2354, Key: "Bul-Kathos' Sacred Charge", Value: `Bul-Kathos' Sacred Charge`},
	},
	"Bul-Kathos' Tribal Guardian": {
		{ID: 2609, Key: "Bul-Kathos' Tribal Guardian", Value: `Bul-Kathos' Tribal Guardian`},
	},
	"Bul-Kathos' Warden": {
		{ID: 2736, Key: "Bul-Kathos' Warden", Value: `Bul-Kathos' Warden`},
	},
	"Buriza-Do Kyanon": {
		{ID: 643, Key: "Buriza-Do Kyanon", Value: `Buriza-Do Kyanon`},
	},
	"Burly": {
		{ID: 1057, Key: "Burly", Value: `Burly`},
	},
	"Burning": {
		{ID: 1155, Key: "Burning", Value: `Burning`},
	},
	"Bush Wacker": {
		{ID: 131, Key: "Bush Wacker", Value: `Bush Wacker`},
	},
	"Butcher's Pupil": {
		{ID: 2527, Key: "Butcher's Pupil", Value: `Butcher's Pupil`},
	},
	"Butchers Cleaver": {
		{ID: 2595, Key: "Butchers Cleaver", Value: `Butcher's Cleaver`},
	},
	"Buzzing": {
		{ID: 2333, Key: "Buzzing", Value: `Buzzing`},
	},
	"CainAct5Gossip1": {
		{ID: 1944, Key: "CainAct5Gossip1", Value: `52
With hellspawn, size is no measure of 
their threat. Demons half the size of 
men can kill with a gesture, while 
hellish pack animals trample any who 
stand in their way.
`},
	},
	"CainAct5Gossip10": {
		{ID: 145, Key: "CainAct5Gossip10", Value: `54
It is my belief that the Soulstones are at 
the center of this conflict. If only that 
fool Marius had not intervened, Baal 
would still be imprisoned within Tal 
Rasha.
`},
	},
	"CainAct5Gossip2": {
		{ID: 1945, Key: "CainAct5Gossip2", Value: `55
Though these Barbarians are known 
throughout the kingdoms as ferocious 
fighters, they are also capable of great 
compassion.
 
They have trained throughout history 
for a battle their legends foretell will 
decide the fate of the world.
`},
	},
	"CainAct5Gossip3": {
		{ID: 1946, Key: "CainAct5Gossip3", Value: `55
The angel Tyrael has watched over the 
guardians of Arreat throughout 
history. I do not believe that Baal and 
Tyrael have come to fight over a paltry 
few souls.
 
They are here to settle a conflict as old 
as time itself.
`},
	},
	"CainAct5Gossip4": {
		{ID: 1947, Key: "CainAct5Gossip4", Value: `44
During my time with the Horadrim, we 
often debated the nature of Mount 
Arreat.
 
We knew that the Barbarian clans 
zealously guarded the mountain as 
their sacred duty. However, many 
dismissed their zeal as simple 
superstition...combined with an inborn 
hostility toward outsiders.
 
Those Horadrim who trekked up Arreat 
were never heard from again...Still, I 
do not believe they died at the hands of 
Barbarians.
`},
	},
	"CainAct5Gossip5": {
		{ID: 1948, Key: "CainAct5Gossip5", Value: `41
All users of the magical arts know of 
Mount Arreat, but few understand its 
true nature. It is the nexus of an 
unfathomable magic.
 
It bodes ill that the Lord of Destruction 
races to its summit with such purpose. 
I fear for the whole world should Baal 
gain what he seeks.
`},
	},
	"CainAct5Gossip6": {
		{ID: 1949, Key: "CainAct5Gossip6", Value: `59
I have spent decades trying to 
understand the forces at work in this 
world. But in the face of all that is 
transpiring, I realize how meager my 
knowledge is.
 
I will be of assistance where I can, my 
friend.
`},
	},
	"CainAct5Gossip7": {
		{ID: 1950, Key: "CainAct5Gossip7", Value: `54
Though the Elder Council of Harrogath 
is gone, there are many capable young 
leaders to take their place.
 
Anya certainly has enough courage and 
intelligence to lead them all, if they can 
survive this catastrophe.
`},
	},
	"CainAct5Gossip8": {
		{ID: 1951, Key: "CainAct5Gossip8", Value: `53
Ah, Anya. Such a fine example of 
feminine strength...
 
She reminds me of the Zakarum 
priestesses I knew in my youth. They 
don't take vows of chastity, you know.
`},
	},
	"CainAct5Gossip9": {
		{ID: 1952, Key: "CainAct5Gossip9", Value: `65
It is fortunate that this town has such a 
talented smith.
 
The quality of Larzuk's work surely 
complements your skills. In fact, he 
would have been quite welcome 
amongst the Horadrim.
`},
	},
	"CainAct5IntroGossip1": {
		{ID: 1285, Key: "CainAct5IntroGossip1", Value: `39
I am amazed to find this place so 
untouched. Everything else in the path 
of Baal the Lord of Destruction lies in 
ruin. 
 
These Barbarians must indeed be the 
legendary guardians of Mount Arreat. 
They are a proud, hardy people. Don't 
expect to be greeted warmly -- 
strangers here rarely are.
 
Perhaps I can gain their trust. I'll spend 
some time with the townsfolk and try 
to understand them better. I'll let you 
know what I discover.
`},
	},
	"Cairn Shard": {
		{ID: 1182, Key: "Cairn Shard", Value: `Cairn Shard`},
	},
	"Calling": {
		{ID: 409, Key: "Calling", Value: `Calling`},
	},
	"Camphor": {
		{ID: 349, Key: "Camphor", Value: `Camphor`},
	},
	"Captain's": {
		{ID: 121, Key: "Captain's", Value: `Captain's`},
	},
	"Carbuncle": {
		{ID: 1867, Key: "Carbuncle", Value: `Carbuncle`},
	},
	"Caretaker's": {
		{ID: 1538, Key: "Caretaker's", Value: `Caretaker's`},
	},
	"Carin Shard": {
		{ID: 2320, Key: "Carin Shard", Value: `Carin Shard`},
	},
	"Carmine": {
		{ID: 240, Key: "Carmine", Value: `Carmine`},
	},
	"Carnageleaver": {
		{ID: 193, Key: "Carnageleaver", Value: `Carnage Leaver`},
	},
	"Carrion Wind": {
		{ID: 1010, Key: "Carrion Wind", Value: `Carrion Wind`},
	},
	"Catapult Spotter E": {
		{ID: 209, Key: "Catapult Spotter E", Value: `Catapult`},
	},
	"Catapult Spotter N": {
		{ID: 205, Key: "Catapult Spotter N", Value: `Catapult`},
	},
	"Catapult Spotter S": {
		{ID: 206, Key: "Catapult Spotter S", Value: `Catapult`},
	},
	"Catapult Spotter Siege": {
		{ID: 215, Key: "Catapult Spotter Siege", Value: `Catapult`},
	},
	"Catapult Spotter W": {
		{ID: 213, Key: "Catapult Spotter W", Value: `Catapult`},
	},
	"CatapultE": {
		{ID: 1160, Key: "CatapultE", Value: `Catapult`},
	},
	"CatapultN": {
		{ID: 1185, Key: "CatapultN", Value: `Catapult`},
	},
	"CatapultS": {
		{ID: 1186, Key: "CatapultS", Value: `Catapult`},
	},
	"CatapultSiege": {
		{ID: 218, Key: "CatapultSiege", Value: `Catapult`},
	},
	"CatapultW": {
		{ID: 1187, Key: "CatapultW", Value: `Catapult`},
	},
	"Catgut": {
		{ID: 1331, Key: "Catgut", Value: `Catgut`},
	},
	"Celestial": {
		{ID: 2811, Key: "Celestial", Value: `Celestial`},
	},
	"Cellar of Pity": {
		{ID: 526, Key: "Cellar of Pity", Value: `Frozen River`},
	},
	"Cerebus": {
		{ID: 2172, Key: "Cerebus", Value: `Cerebus' Bite`},
	},
	"CfgSay7": {
		{ID: 2076, Key: "CfgSay7", Value: `Say 'Retreat'`},
	},
	"CfgSkill10": {
		{ID: 695, Key: "CfgSkill10", Value: `Skill 10`},
	},
	"CfgSkill11": {
		{ID: 696, Key: "CfgSkill11", Value: `Skill 11`},
	},
	"CfgSkill12": {
		{ID: 702, Key: "CfgSkill12", Value: `Skill 12`},
	},
	"CfgSkill13": {
		{ID: 703, Key: "CfgSkill13", Value: `Skill 13`},
	},
	"CfgSkill14": {
		{ID: 704, Key: "CfgSkill14", Value: `Skill 14`},
	},
	"CfgSkill15": {
		{ID: 705, Key: "CfgSkill15", Value: `Skill 15`},
	},
	"CfgSkill16": {
		{ID: 710, Key: "CfgSkill16", Value: `Skill 16`},
	},
	"CfgSkill9": {
		{ID: 694, Key: "CfgSkill9", Value: `Skill 9`},
	},
	"CfgToggleminimap": {
		{ID: 712, Key: "CfgToggleminimap", Value: `Toggle MiniMap`},
	},
	"Cfghireling": {
		{ID: 766, Key: "Cfghireling", Value: `Hireling Screen`},
	},
	"Cfgswapweapons": {
		{ID: 715, Key: "Cfgswapweapons", Value: `Swap Weapons`},
	},
	"ChampionFormatX": {
		{ID: 913, Key: "ChampionFormatX", Value: `%0 %1`},
	},
	"Chaotic": {
		{ID: 1457, Key: "Chaotic", Value: `Chaotic`},
	},
	"Charged": {
		{ID: 2025, Key: "Charged", Value: `Charged`},
	},
	"Charged Bolt Sentry": {
		{ID: 59, Key: "Charged Bolt Sentry", Value: `Charged Bolt Sentry`},
	},
	"Charmdes": {
		{ID: 905, Key: "Charmdes", Value: `Keep in Inventory to Gain Bonus`},
	},
	"Charsi's Favor": {
		{ID: 2613, Key: "Charsi's Favor", Value: `Charsi's Favor`},
	},
	"ChestL": {
		{ID: 312, Key: "ChestL", Value: `Chest`},
	},
	"ChestSL": {
		{ID: 530, Key: "ChestSL", Value: `Chest`},
	},
	"ChestSR": {
		{ID: 676, Key: "ChestSR", Value: `Chest`},
	},
	"Chestnut": {
		{ID: 2493, Key: "Chestnut", Value: `Chestnut`},
	},
	"Chilling": {
		{ID: 2772, Key: "Chilling", Value: `Chilling`},
	},
	"Chromatic": {
		{ID: 1586, Key: "Chromatic", Value: `Chromatic`},
	},
	"Chromatic Ire": {
		{ID: 184, Key: "Chromatic Ire", Value: `Chromatic Ire`},
	},
	"Cinnabar": {
		{ID: 287, Key: "Cinnabar", Value: `Cinnabar`},
	},
	"Class Specific": {
		{ID: 1001, Key: "Class Specific", Value: `Class-specific`},
	},
	"Cliffkiller": {
		{ID: 482, Key: "Cliffkiller", Value: `Cliffkiller`},
	},
	"Cloudcrack": {
		{ID: 2159, Key: "Cloudcrack", Value: `Cloudcrack`},
	},
	"Cloudy": {
		{ID: 315, Key: "Cloudy", Value: `Cloudy`},
	},
	"Coldkill": {
		{ID: 344, Key: "Coldkill", Value: `Coldkill`},
	},
	"Coldsteal Eye": {
		{ID: 2169, Key: "Coldsteal Eye", Value: `Coldsteal Eye`},
	},
	"Coldsteel Eye": {
		{ID: 1062, Key: "Coldsteel Eye", Value: `Coldsteel Eye`},
	},
	"Colorful": {
		{ID: 1996, Key: "Colorful", Value: `Colorful`},
	},
	"Commander's": {
		{ID: 1207, Key: "Commander's", Value: `Commander's`},
	},
	"Communal": {
		{ID: 684, Key: "Communal", Value: `Communal`},
	},
	"Compact": {
		{ID: 26, Key: "Compact", Value: `Compact`},
	},
	"CompletingAndarielAss": {
		{ID: 2425, Key: "CompletingAndarielAss", Value: `Death becomes you, Andariel.`},
	},
	"CompletingAndarielDru": {
		{ID: 2220, Key: "CompletingAndarielDru", Value: `Your reign is over, Andariel.`},
	},
	"CompletingBeneathCityAss": {
		{ID: 2039, Key: "CompletingBeneathCityAss", Value: `This is one drain I don't mind cleaning out.`},
	},
	"CompletingBeneathCityDru": {
		{ID: 1610, Key: "CompletingBeneathCityDru", Value: `From trash to treasure...`},
	},
	"CompletingBladeAss": {
		{ID: 2812, Key: "CompletingBladeAss", Value: `What a delicious blade! I should consult Ormus.`},
	},
	"CompletingBladeDru": {
		{ID: 779, Key: "CompletingBladeDru", Value: `Ormus may know something about this unusual blade.`},
	},
	"CompletingBurialAss": {
		{ID: 1219, Key: "CompletingBurialAss", Value: `What I kill stays dead.`},
	},
	"CompletingBurialDru": {
		{ID: 437, Key: "CompletingBurialDru", Value: `Your time is past, Blood Raven.`},
	},
	"CompletingDOEAss": {
		{ID: 1184, Key: "CompletingDOEAss", Value: `The Rogues' test is done.`},
	},
	"CompletingDOEDru": {
		{ID: 914, Key: "CompletingDOEDru", Value: `Bah! Is that all of them?`},
	},
	"CompletingDefeatBaalAct5Ams": {
		{ID: 1295, Key: "CompletingDefeatBaalAct5Ams", Value: `My work here is truly done.`},
	},
	"CompletingDefeatBaalAct5Ass": {
		{ID: 2213, Key: "CompletingDefeatBaalAct5Ass", Value: `The Evil brotherhood is no more.`},
	},
	"CompletingDefeatBaalAct5Bar": {
		{ID: 2212, Key: "CompletingDefeatBaalAct5Bar", Value: `The Prime Evils are no more.`},
	},
	"CompletingDefeatBaalAct5Dru": {
		{ID: 54, Key: "CompletingDefeatBaalAct5Dru", Value: `Baal! Join your brothers in oblivion.`},
	},
	"CompletingDefeatBaalAct5Nec": {
		{ID: 1765, Key: "CompletingDefeatBaalAct5Nec", Value: `Baal, never doubt my skills.`},
	},
	"CompletingDefeatBaalAct5Pal": {
		{ID: 2071, Key: "CompletingDefeatBaalAct5Pal", Value: `Baal, you shall no longer taint this mortal realm.`},
	},
	"CompletingDefeatBaalAct5Sor": {
		{ID: 1054, Key: "CompletingDefeatBaalAct5Sor", Value: `The last of the Three has fallen.`},
	},
	"CompletingForgottenTPal": {
		{ID: 1067, Key: "CompletingForgottenTPal", Value: `So much treasure almost covers the stench.`},
		{ID: 1068, Key: "CompletingForgottenTPal", Value: `This tower has its charms...`},
	},
	"CompletingGuardianTowerAss": {
		{ID: 2047, Key: "CompletingGuardianTowerAss", Value: `Mephisto, you were no match for me.`},
		{ID: 2048, Key: "CompletingGuardianTowerAss", Value: `Mephisto's hatred was a poisonous void.`},
	},
	"CompletingLamEsenAss": {
		{ID: 1292, Key: "CompletingLamEsenAss", Value: `One can't judge a book by its cover.`},
		{ID: 1294, Key: "CompletingLamEsenAss", Value: `Ormus... You have strange taste in books.`},
	},
	"CompletingLamEsenDru": {
		{ID: 2038, Key: "CompletingLamEsenDru", Value: `Ormus... study the book well.`},
	},
	"CompletingNihlathakAct5Ams": {
		{ID: 862, Key: "CompletingNihlathakAct5Ams", Value: `Conspiring with Baal... What a tragic mistake.`},
	},
	"CompletingNihlathakAct5Ass": {
		{ID: 962, Key: "CompletingNihlathakAct5Ass", Value: `You Dark Mages are all alike - obsessed with power.`},
	},
	"CompletingNihlathakAct5Bar": {
		{ID: 961, Key: "CompletingNihlathakAct5Bar", Value: `A fitting death for a traitor.`},
	},
	"CompletingNihlathakAct5Dru": {
		{ID: 1764, Key: "CompletingNihlathakAct5Dru", Value: `Betrayer, you've reaped your reward.`},
	},
	"CompletingNihlathakAct5Nec": {
		{ID: 1227, Key: "CompletingNihlathakAct5Nec", Value: `You were a sad little man, Nihlathak.`},
	},
	"CompletingNihlathakAct5Pal": {
		{ID: 1763, Key: "CompletingNihlathakAct5Pal", Value: `Nihlathak. What led you to this end?`},
	},
	"CompletingNihlathakAct5Sor": {
		{ID: 2681, Key: "CompletingNihlathakAct5Sor", Value: `Your power was no match for mine.`},
	},
	"CompletingRadamentAss": {
		{ID: 1484, Key: "CompletingRadamentAss", Value: `Vengeance... for Atma.`},
	},
	"CompletingRadamentDru": {
		{ID: 1195, Key: "CompletingRadamentDru", Value: `Return to dust, Radament.`},
	},
	"CompletingStopSiegeAms": {
		{ID: 2741, Key: "CompletingStopSiegeAms", Value: `Oops...Did I do that?`},
	},
	"CompletingStopSiegeAss": {
		{ID: 2068, Key: "CompletingStopSiegeAss", Value: `Shenk, your command has ended.`},
	},
	"CompletingStopSiegeBar": {
		{ID: 2050, Key: "CompletingStopSiegeBar", Value: `The siege is broken.`},
	},
	"CompletingStopSiegeDru": {
		{ID: 1265, Key: "CompletingStopSiegeDru", Value: `The catapults have been silenced.`},
	},
	"CompletingStopSiegeNec": {
		{ID: 1540, Key: "CompletingStopSiegeNec", Value: `My, my, what a messy little demon.`},
	},
	"CompletingStopSiegePal": {
		{ID: 1761, Key: "CompletingStopSiegePal", Value: `Harrogath is free of your kind, demon.`},
	},
	"CompletingStopSiegeSor": {
		{ID: 1249, Key: "CompletingStopSiegeSor", Value: `Harrogath can rest easier now.`},
	},
	"CompletingSummonerAss": {
		{ID: 1034, Key: "CompletingSummonerAss", Value: `Horazon. Your decoy is dead.`},
	},
	"CompletingSummonerDru": {
		{ID: 2037, Key: "CompletingSummonerDru", Value: `Now I can leave this twisted nightmare.`},
	},
	"CompletingTaintedSunAss": {
		{ID: 2201, Key: "CompletingTaintedSunAss", Value: `Serpents! I expected worse.`},
	},
	"CompletingTaintedSunDru": {
		{ID: 2036, Key: "CompletingTaintedSunDru", Value: `The sun warms the world once more.`},
	},
	"CompletingTempleAss": {
		{ID: 2734, Key: "CompletingTempleAss", Value: `The dark magic here is dispelled.`},
	},
	"CompletingTempleDru": {
		{ID: 669, Key: "CompletingTempleDru", Value: `There is hope once again.`},
	},
	"CompletingTombAss": {
		{ID: 2202, Key: "CompletingTombAss", Value: `I shall track the Prime Evils to the ends of the world.`},
	},
	"CompletingTombDru": {
		{ID: 102, Key: "CompletingTombDru", Value: `Diablo... I will find you yet.`},
	},
	"Consecrated": {
		{ID: 36, Key: "Consecrated", Value: `Consecrated`},
	},
	"Constricting Ring": {
		{ID: 2596, Key: "Constricting Ring", Value: `Constricting Ring`},
	},
	"ConsumedFireBoar": {
		{ID: 94, Key: "ConsumedFireBoar", Value: `Consumed Fire Boar`},
	},
	"ConsumedIceBoar": {
		{ID: 100, Key: "ConsumedIceBoar", Value: `Consumed Ice Boar`},
	},
	"Continuous": {
		{ID: 1309, Key: "Continuous", Value: `Everlasting`},
	},
	"Corosive": {
		{ID: 2791, Key: "Corosive", Value: `Corrosive`},
	},
	"Corporal": {
		{ID: 550, Key: "Corporal", Value: `Corporal`},
	},
	"Corpsemourn": {
		{ID: 2280, Key: "Corpsemourn", Value: `Corpsemourn`},
	},
	"Cow King's Hide": {
		{ID: 1729, Key: "Cow King's Hide", Value: `Cow King's Hide`},
	},
	"Cow King's Hoofs": {
		{ID: 2771, Key: "Cow King's Hoofs", Value: `Cow King's Hooves`},
	},
	"Cow King's Horns": {
		{ID: 2667, Key: "Cow King's Horns", Value: `Cow King's Horns`},
	},
	"Cow King's Leathers": {
		{ID: 2012, Key: "Cow King's Leathers", Value: `Cow King's Leathers`},
	},
	"Crainte Vomir": {
		{ID: 1300, Key: "Crainte Vomir", Value: `Crainte Vomir`},
	},
	"Cranebeak": {
		{ID: 101, Key: "Cranebeak", Value: `Cranebeak`},
	},
	"Credendum": {
		{ID: 51, Key: "Credendum", Value: `Credendum`},
	},
	"Crescent Moon": {
		{ID: 1161, Key: "Crescent Moon", Value: `Crescent Moon`},
	},
	"Crest of Morn": {
		{ID: 2809, Key: "Crest of Morn", Value: `Crest of Morn`},
	},
	"Crow Caw": {
		{ID: 967, Key: "Crow Caw", Value: `Crow Caw`},
	},
	"Crown of Ages": {
		{ID: 2747, Key: "Crown of Ages", Value: `Crown of Ages`},
	},
	"Crown of Thieves": {
		{ID: 159, Key: "Crown of Thieves", Value: `Crown of Thieves`},
	},
	"Crude": {
		{ID: 1404, Key: "Crude", Value: `Crude`},
	},
	"CrushBiest": {
		{ID: 79, Key: "CrushBiest", Value: `Crush Beast`},
	},
	"Crystalized Cavern Level 1": {
		{ID: 1055, Key: "Crystalized Cavern Level 1", Value: `Crystalline Passage`},
	},
	"Crystalized Cavern Level 2": {
		{ID: 1059, Key: "Crystalized Cavern Level 2", Value: `Glacial Trail`},
	},
	"Cunning": {
		{ID: 2520, Key: "Cunning", Value: `Cunning`},
	},
	"Curly": {
		{ID: 273, Key: "Curly", Value: `Curly`},
	},
	"Cutthroat": {
		{ID: 2606, Key: "Cutthroat", Value: `Bartuc's Chop Chop`},
	},
	"Dac Farren": {
		{ID: 850, Key: "Dac Farren", Value: `Dac Farren`},
	},
	"Dangoon's Teaching": {
		{ID: 2518, Key: "Dangoon's Teaching", Value: `Dangoon's Teaching`},
	},
	"Dark Clan Crusher": {
		{ID: 993, Key: "Dark Clan Crusher", Value: `Dark Clan Crusher`},
	},
	"Darkfear": {
		{ID: 1655, Key: "Darkfear", Value: `Darkfear`},
	},
	"Darkforge Spawn": {
		{ID: 185, Key: "Darkforge Spawn", Value: `Darkforce Spawn`},
	},
	"Darksight Helm": {
		{ID: 1798, Key: "Darksight Helm", Value: `Darksight Helm`},
	},
	"Darksoul": {
		{ID: 2171, Key: "Darksoul", Value: `Darksoul`},
	},
	"Dawnbringer": {
		{ID: 427, Key: "Dawnbringer", Value: `Dawn Bringer`},
	},
	"DeamonSteed": {
		{ID: 1271, Key: "DeamonSteed", Value: `Demon Steed`},
	},
	"Death Mauler1": {
		{ID: 1090, Key: "Death Mauler1", Value: `Death Mauler`},
	},
	"Death Mauler2": {
		{ID: 1092, Key: "Death Mauler2", Value: `Death Brawler`},
	},
	"Death Mauler3": {
		{ID: 1093, Key: "Death Mauler3", Value: `Death Slasher`},
	},
	"Death Mauler4": {
		{ID: 1100, Key: "Death Mauler4", Value: `Death Berserker`},
	},
	"Death Mauler5": {
		{ID: 1102, Key: "Death Mauler5", Value: `Death Brigadier`},
	},
	"Death Sentry": {
		{ID: 242, Key: "Death Sentry", Value: `Death Sentry`},
	},
	"Death's Web": {
		{ID: 2745, Key: "Death's Web", Value: `Death's Web`},
	},
	"Deathbit": {
		{ID: 158, Key: "Deathbit", Value: `Deathbit`},
	},
	"Deathcleaver": {
		{ID: 2778, Key: "Deathcleaver", Value: `Death Cleaver`},
	},
	"Deathexp": {
		{ID: 85, Key: "Deathexp", Value: `Death`},
	},
	"DeathlyVisage": {
		{ID: 83, Key: "DeathlyVisage", Value: `Deathly Visage`},
	},
	"Deaths's Web": {
		{ID: 1035, Key: "Deaths's Web", Value: `Death's Web`},
	},
	"Decaying": {
		{ID: 2573, Key: "Decaying", Value: `Decaying`},
	},
	"DefiledWarrior": {
		{ID: 77, Key: "DefiledWarrior", Value: `Defiled Warrior`},
	},
	"Demon Machine": {
		{ID: 773, Key: "Demon Machine", Value: `Demon Machine`},
	},
	"Demon's Arch": {
		{ID: 97, Key: "Demon's Arch", Value: `Demon's Arch`},
	},
	"Demonhorn's Edge": {
		{ID: 270, Key: "Demonhorn's Edge", Value: `Demonhorn's Edge`},
	},
	"Demonlimb": {
		{ID: 858, Key: "Demonlimb", Value: `Demon Limb`},
	},
	"Demonweb": {
		{ID: 511, Key: "Demonweb", Value: `Demonweb`},
	},
	"Dense": {
		{ID: 824, Key: "Dense", Value: `Dense`},
	},
	"Diamond": {
		{ID: 1779, Key: "Diamond", Value: `Diamond`},
	},
	"Divine": {
		{ID: 1313, Key: "Divine", Value: `Divine`},
	},
	"Djinnslayer": {
		{ID: 1925, Key: "Djinnslayer", Value: `Djinn Slayer`},
	},
	"Dominus": {
		{ID: 115, Key: "Dominus", Value: `Siren`},
	},
	"Doombringer": {
		{ID: 395, Key: "Doombringer", Value: `Doombringer`},
	},
	"Doomseer": {
		{ID: 404, Key: "Doomseer", Value: `Doomseer`},
	},
	"Doppleganger's Shadow": {
		{ID: 1434, Key: "Doppleganger's Shadow", Value: `Doppleganger's Shadow`},
	},
	"Dracul's Grasp": {
		{ID: 2062, Key: "Dracul's Grasp", Value: `Dracul's Grasp`},
	},
	"Dragonscale": {
		{ID: 2664, Key: "Dragonscale", Value: `Dragonscale`},
	},
	"Dragontooth": {
		{ID: 2224, Key: "Dragontooth", Value: `Dragontooth`},
	},
	"Drakeflame": {
		{ID: 2665, Key: "Drakeflame", Value: `Drakeflame`},
	},
	"Drehya": {
		{ID: 261, Key: "Drehya", Value: `Anya`},
	},
	"DruOnly": {
		{ID: 1852, Key: "DruOnly", Value: `(Druid Only)`},
	},
	"Druid Bear": {
		{ID: 70, Key: "Druid Bear", Value: `Grizzly`},
	},
	"Druid Cycle of Life": {
		{ID: 71, Key: "Druid Cycle of Life", Value: `Cycle of Life`},
	},
	"Druid Fenris": {
		{ID: 68, Key: "Druid Fenris", Value: `Dire Wolf`},
	},
	"Druid Hawk": {
		{ID: 1031, Key: "Druid Hawk", Value: `Raven`},
	},
	"Druid Plague Poppy": {
		{ID: 323, Key: "Druid Plague Poppy", Value: `Plague Poppy`},
	},
	"Druid Spirit Wolf": {
		{ID: 388, Key: "Druid Spirit Wolf", Value: `Spirit Wolf`},
	},
	"Druid Totem": {
		{ID: 687, Key: "Druid Totem", Value: `Druid Spirit`},
	},
	"Druid Wolf": {
		{ID: 66, Key: "Druid Wolf", Value: `Wolf`},
	},
	"Drulan's Tongue": {
		{ID: 1478, Key: "Drulan's Tongue", Value: `Drulan's Tongue`},
	},
	"Drulan's Tounge": {
		{ID: 2324, Key: "Drulan's Tounge", Value: `Drulan's Tongue`},
	},
	"Dun": {
		{ID: 2476, Key: "Dun", Value: `Dun`},
	},
	"Duriel's Shell": {
		{ID: 2414, Key: "Duriel's Shell", Value: `Duriel's Shell`},
	},
	"Dust Storm": {
		{ID: 17, Key: "Dust Storm", Value: `Dust Storm`},
	},
	"Dusty": {
		{ID: 20, Key: "Dusty", Value: `Dusty`},
	},
	"Dwarf Star": {
		{ID: 2534, Key: "Dwarf Star", Value: `Dwarf Star`},
	},
	"ESkillHawk": {
		{ID: 2800, Key: "ESkillHawk", Value: `Ravens: `},
	},
	"ESkillPerKick": {
		{ID: 1109, Key: "ESkillPerKick", Value: ` per kick`},
	},
	"ESkillShoots": {
		{ID: 2804, Key: "ESkillShoots", Value: `Shoots `},
	},
	"ESkillSpikes": {
		{ID: 2801, Key: "ESkillSpikes", Value: `Spikes: `},
	},
	"ESkillSpikes2": {
		{ID: 2805, Key: "ESkillSpikes2", Value: ` Spikes`},
	},
	"ESkillStars": {
		{ID: 653, Key: "ESkillStars", Value: `Stars: `},
	},
	"ESkillTimes": {
		{ID: 419, Key: "ESkillTimes", Value: ` Times`},
	},
	"ESkillWolf": {
		{ID: 627, Key: "ESkillWolf", Value: `Wolf: `},
	},
	"ESkillWolves": {
		{ID: 2802, Key: "ESkillWolves", Value: `Wolves: `},
	},
	"Eagleexp": {
		{ID: 328, Key: "Eagleexp", Value: `Eagle`},
	},
	"Eagleeye": {
		{ID: 333, Key: "Eagleeye", Value: `Eagleeye`},
	},
	"Eaglehorn": {
		{ID: 1728, Key: "Eaglehorn", Value: `Eaglehorn`},
	},
	"Eaglewind": {
		{ID: 422, Key: "Eaglewind", Value: `Eaglewind`},
	},
	"Earthshaker": {
		{ID: 2797, Key: "Earthshaker", Value: `Earthshaker`},
	},
	"Earthshifter": {
		{ID: 1732, Key: "Earthshifter", Value: `Earth Shifter`},
	},
	"Eburin": {
		{ID: 105, Key: "Eburin", Value: `Eburine`},
	},
	"Echo Chamber": {
		{ID: 527, Key: "Echo Chamber", Value: `Drifter Cavern`},
	},
	"Echoing": {
		{ID: 1136, Key: "Echoing", Value: `Echoing`},
	},
	"El Espiritu": {
		{ID: 2773, Key: "El Espiritu", Value: `El Espiritu`},
	},
	"El Infierno": {
		{ID: 2395, Key: "El Infierno", Value: `El Infierno`},
	},
	"Elysian": {
		{ID: 183, Key: "Elysian", Value: `Elysian`},
	},
	"Ember": {
		{ID: 2817, Key: "Ember", Value: `Fiery`},
	},
	"Endlesshail": {
		{ID: 606, Key: "Endlesshail", Value: `Endlesshail`},
	},
	"Endlessshail": {
		{ID: 1291, Key: "Endlessshail", Value: `Endlesshail`},
	},
	"EnterBurialAss": {
		{ID: 719, Key: "EnterBurialAss", Value: `Whose handiwork lies buried here?`},
	},
	"EnterBurialDru": {
		{ID: 960, Key: "EnterBurialDru", Value: `Planting the dead... How odd.`},
	},
	"EnterCatacombsAss": {
		{ID: 2326, Key: "EnterCatacombsAss", Value: `I don't like it down here.`},
	},
	"EnterCatacombsDru": {
		{ID: 2633, Key: "EnterCatacombsDru", Value: `The supernatural is strong here.`},
	},
	"EnterDOEAss": {
		{ID: 2297, Key: "EnterDOEAss", Value: `So dark... perfect.`},
	},
	"EnterDOEDru": {
		{ID: 2034, Key: "EnterDOEDru", Value: `So, this is where evil hides.`},
	},
	"EnterForgottenTAss": {
		{ID: 178, Key: "EnterForgottenTAss", Value: `Who would want to remember this place?`},
	},
	"EnterForgottenTDru": {
		{ID: 2692, Key: "EnterForgottenTDru", Value: `I can smell why this tower was abandoned.`},
	},
	"EnterJailAss": {
		{ID: 701, Key: "EnterJailAss", Value: `Try and cage me, demons.`},
	},
	"EnterJailDru": {
		{ID: 1485, Key: "EnterJailDru", Value: `Bars can't hold a force of nature.`},
	},
	"EnterMonasteryAss": {
		{ID: 675, Key: "EnterMonasteryAss", Value: `Such corruption in this place...`},
	},
	"EnterMonasteryDru": {
		{ID: 468, Key: "EnterMonasteryDru", Value: `Evil flows from here.`},
	},
	"EnteringArcaneAss": {
		{ID: 625, Key: "EnteringArcaneAss", Value: `The Sanctuary - Horazon's obsession.`},
	},
	"EnteringArcaneDru": {
		{ID: 356, Key: "EnteringArcaneDru", Value: `This was not designed by nature's Architect.`},
	},
	"EnteringClawViperAss": {
		{ID: 2035, Key: "EnteringClawViperAss", Value: `Dark magic in a darker tomb...`},
	},
	"EnteringClawViperDru": {
		{ID: 2570, Key: "EnteringClawViperDru", Value: `Snakes... I hate snakes.`},
	},
	"EnteringNihlathakAct5Ams": {
		{ID: 2763, Key: "EnteringNihlathakAct5Ams", Value: `...Nihlathak's home away from home.`},
	},
	"EnteringNihlathakAct5Ass": {
		{ID: 2051, Key: "EnteringNihlathakAct5Ass", Value: `I should have known...`},
	},
	"EnteringNihlathakAct5Bar": {
		{ID: 1797, Key: "EnteringNihlathakAct5Bar", Value: `A coward's hiding place.`},
	},
	"EnteringNihlathakAct5Dru": {
		{ID: 2586, Key: "EnteringNihlathakAct5Dru", Value: `Nihlathak... you can't hide from me.`},
	},
	"EnteringNihlathakAct5Nec": {
		{ID: 52, Key: "EnteringNihlathakAct5Nec", Value: `Ahh, the familiar scent of death.`},
	},
	"EnteringNihlathakAct5Pal": {
		{ID: 765, Key: "EnteringNihlathakAct5Pal", Value: `By the Light! What is this place?`},
	},
	"EnteringNihlathakAct5Sor": {
		{ID: 992, Key: "EnteringNihlathakAct5Sor", Value: `Could this be a trap?`},
	},
	"EnteringRadamentAss": {
		{ID: 554, Key: "EnteringRadamentAss", Value: `Why must evil hide in such wretched places?`},
	},
	"EnteringRadamentDru": {
		{ID: 286, Key: "EnteringRadamentDru", Value: `Face the light or lurk in darkness.`},
	},
	"EnteringTopMountAct5Ams": {
		{ID: 380, Key: "EnteringTopMountAct5Ams", Value: `The fabled home of the Ancients.`},
	},
	"EnteringTopMountAct5Ass": {
		{ID: 2531, Key: "EnteringTopMountAct5Ass", Value: `I shall prove worthy.`},
	},
	"EnteringTopMountAct5Bar": {
		{ID: 2530, Key: "EnteringTopMountAct5Bar", Value: `The guardians of Mount Arreat await.`},
	},
	"EnteringTopMountAct5Dru": {
		{ID: 458, Key: "EnteringTopMountAct5Dru", Value: `At last...The summit of Mount Arreat.`},
	},
	"EnteringTopMountAct5Nec": {
		{ID: 680, Key: "EnteringTopMountAct5Nec", Value: `The resting place of the Ancients...`},
	},
	"EnteringTopMountAct5Pal": {
		{ID: 1492, Key: "EnteringTopMountAct5Pal", Value: `The summit - The Barbarian holy ground.`},
	},
	"EnteringTopMountAct5Sor": {
		{ID: 1676, Key: "EnteringTopMountAct5Sor", Value: `The fabled home of the Ancients.`},
	},
	"EnteringWorldstoneAct5Ams": {
		{ID: 2061, Key: "EnteringWorldstoneAct5Ams", Value: `The Worldstone!`},
	},
	"EnteringWorldstoneAct5Ass": {
		{ID: 2064, Key: "EnteringWorldstoneAct5Ass", Value: `The Worldstone. What power.`},
	},
	"EnteringWorldstoneAct5Bar": {
		{ID: 2056, Key: "EnteringWorldstoneAct5Bar", Value: `The halls of the Ancients... Magnificent.`},
	},
	"EnteringWorldstoneAct5Dru": {
		{ID: 2250, Key: "EnteringWorldstoneAct5Dru", Value: `The legendary Worldstone - guardian of the Natural realm.`},
	},
	"EnteringWorldstoneAct5Nec": {
		{ID: 163, Key: "EnteringWorldstoneAct5Nec", Value: `So, this is what the Ancients guard.`},
	},
	"EnteringWorldstoneAct5Pal": {
		{ID: 475, Key: "EnteringWorldstoneAct5Pal", Value: `The Worldstone! Praise the Light.`},
	},
	"EnteringWorldstoneAct5Sor": {
		{ID: 222, Key: "EnteringWorldstoneAct5Sor", Value: `The power of the Worldstone washes over me.`},
	},
	"Entrapping": {
		{ID: 2599, Key: "Entrapping", Value: `Entrapping`},
	},
	"Envenomed": {
		{ID: 2625, Key: "Envenomed", Value: `Foul`},
	},
	"Eocene": {
		{ID: 2680, Key: "Eocene", Value: `Righteous`},
	},
	"Erion's Bonehandle": {
		{ID: 745, Key: "Erion's Bonehandle", Value: `Erion's Bonehandle`},
	},
	"Eschuta's temper": {
		{ID: 539, Key: "Eschuta's temper", Value: `Eschuta's Temper`},
		{ID: 540, Key: "Eschuta's temper", Value: `Eschuta's Temper`},
	},
	"EskillKickPlur": {
		{ID: 754, Key: "EskillKickPlur", Value: ` Kicks`},
	},
	"EskillKickSing": {
		{ID: 1659, Key: "EskillKickSing", Value: ` Kick`},
	},
	"EskillLifeSteal": {
		{ID: 2807, Key: "EskillLifeSteal", Value: `Life Steal: `},
	},
	"EskillManaSteal": {
		{ID: 329, Key: "EskillManaSteal", Value: `Mana Steal: `},
	},
	"EskillPassiveFeral": {
		{ID: 2019, Key: "EskillPassiveFeral", Value: `Passive Bonus to Wolves and Bears`},
	},
	"EskillPerBlade": {
		{ID: 1740, Key: "EskillPerBlade", Value: ` per blade`},
	},
	"EskillPetLife": {
		{ID: 1012, Key: "EskillPetLife", Value: `Pet Life: `},
	},
	"EskillWolfDef": {
		{ID: 2431, Key: "EskillWolfDef", Value: `Wolf Defense: `},
	},
	"Eskillbladesofice1": {
		{ID: 25, Key: "Eskillbladesofice1", Value: `cold damage: `},
	},
	"Eskillbladesofice2": {
		{ID: 27, Key: "Eskillbladesofice2", Value: `cold damage radius: `},
	},
	"Eskillbladesofice3": {
		{ID: 28, Key: "Eskillbladesofice3", Value: `freeze duration: `},
	},
	"Eskillchancetoafflict": {
		{ID: 1296, Key: "Eskillchancetoafflict", Value: `Chance to afflict target: `},
	},
	"Eskillchancetostun": {
		{ID: 2813, Key: "Eskillchancetostun", Value: `Chance to stun: `},
	},
	"Eskillelementaldmg": {
		{ID: 2614, Key: "Eskillelementaldmg", Value: `Elemental Charge-up Damage: `},
	},
	"Eskillferalpets": {
		{ID: 9, Key: "Eskillferalpets", Value: `Feral Pets: `},
	},
	"Eskillfinishmove": {
		{ID: 15, Key: "Eskillfinishmove", Value: `Finishing Move - `},
	},
	"Eskillfistsoffire1": {
		{ID: 19, Key: "Eskillfistsoffire1", Value: `fire damage: `},
	},
	"Eskillfistsoffire2": {
		{ID: 22, Key: "Eskillfistsoffire2", Value: `fire damage radius: `},
	},
	"Eskillfistsoffire3": {
		{ID: 23, Key: "Eskillfistsoffire3", Value: `burning duration: `},
	},
	"Eskillincasehit": {
		{ID: 1251, Key: "Eskillincasehit", Value: ` hit`},
	},
	"Eskillincasemastery": {
		{ID: 2512, Key: "Eskillincasemastery", Value: `%d Percent Chance of Critical Strike`},
	},
	"Eskillincaseraven": {
		{ID: 2286, Key: "Eskillincaseraven", Value: `Mana Cost Per Raven: `},
	},
	"Eskillkickdamage": {
		{ID: 2715, Key: "Eskillkickdamage", Value: `Kick Damage: `},
	},
	"Eskilllowerresis": {
		{ID: 331, Key: "Eskilllowerresis", Value: `Lowers Resistance `},
	},
	"Eskillmanarecov": {
		{ID: 1277, Key: "Eskillmanarecov", Value: `Mana Recovered: `},
	},
	"Eskillpercentatt": {
		{ID: 11, Key: "Eskillpercentatt", Value: ` Percent Attack`},
	},
	"Eskillpercentdmg": {
		{ID: 12, Key: "Eskillpercentdmg", Value: ` Percent Damage`},
	},
	"Eskillpercentlif": {
		{ID: 986, Key: "Eskillpercentlif", Value: ` Percent Life`},
	},
	"Eskillperhit12": {
		{ID: 2021, Key: "Eskillperhit12", Value: ` per hit`},
	},
	"Eskillphoenix1": {
		{ID: 1139, Key: "Eskillphoenix1", Value: `meteor damage: `},
	},
	"Eskillphoenix2": {
		{ID: 1140, Key: "Eskillphoenix2", Value: `chain lightning damage: `},
	},
	"Eskillphoenix3": {
		{ID: 1141, Key: "Eskillphoenix3", Value: `chaos ice bolt damage: `},
	},
	"Eskillpowerup1": {
		{ID: 2815, Key: "Eskillpowerup1", Value: `Charge 1 - `},
	},
	"Eskillpowerup2": {
		{ID: 2816, Key: "Eskillpowerup2", Value: `Charge 2 - `},
	},
	"Eskillpowerup3": {
		{ID: 4, Key: "Eskillpowerup3", Value: `Charge 3 - `},
	},
	"Eskillpowerupadd": {
		{ID: 164, Key: "Eskillpowerupadd", Value: `Adds `},
	},
	"Eskillpudburning": {
		{ID: 18, Key: "Eskillpudburning", Value: `burning damage`},
	},
	"Eskillpuddgmper": {
		{ID: 7, Key: "Eskillpuddgmper", Value: ` percent damage`},
	},
	"Eskillpudlife": {
		{ID: 6, Key: "Eskillpudlife", Value: ` percent life stealing`},
	},
	"Eskillpudmana": {
		{ID: 806, Key: "Eskillpudmana", Value: ` percent life and mana stealing `},
	},
	"Eskillsinishup": {
		{ID: 5, Key: "Eskillsinishup", Value: `Finishing Move Bonuses`},
	},
	"Eskillthunder1": {
		{ID: 418, Key: "Eskillthunder1", Value: `lightning damage: `},
	},
	"Eskillthunder2": {
		{ID: 420, Key: "Eskillthunder2", Value: `nova damage: `},
	},
	"Eskillthunder3": {
		{ID: 421, Key: "Eskillthunder3", Value: `charged bolt damage: `},
	},
	"Eskilltomeleeattacks": {
		{ID: 741, Key: "Eskilltomeleeattacks", Value: ` to melee attacks`},
	},
	"Ethereal Edge": {
		{ID: 2737, Key: "Ethereal Edge", Value: `Ethereal Edge`},
	},
	"Ethereal edge": {
		{ID: 1347, Key: "Ethereal edge", Value: `Ethereal Edge`},
	},
	"Evil": {
		{ID: 1711, Key: "Evil", Value: `Evil`},
	},
	"Evil hut": {
		{ID: 1225, Key: "Evil hut", Value: `Evil Demon Hut`},
	},
	"ExInsertSockets": {
		{ID: 21, Key: "ExInsertSockets", Value: `Can be Inserted into Socketed Items`},
	},
	"Excalibur": {
		{ID: 2775, Key: "Excalibur", Value: `Excalibur`},
	},
	"Executioner's Justice": {
		{ID: 1921, Key: "Executioner's Justice", Value: `Executioner's Justice`},
	},
	"Expert's": {
		{ID: 616, Key: "Expert's", Value: `Expert's`},
	},
	"Eyeback Unleashed": {
		{ID: 515, Key: "Eyeback Unleashed", Value: `Eyeback the Unleashed`},
	},
	"Falconeye": {
		{ID: 448, Key: "Falconeye", Value: `Falconeye`},
	},
	"Fanatic": {
		{ID: 252, Key: "Fanatic", Value: `Fanatic`},
	},
	"FanaticMinion": {
		{ID: 1180, Key: "FanaticMinion", Value: `Fanatic Enslaved`},
	},
	"Fathom": {
		{ID: 1209, Key: "Fathom", Value: `Death's Fathom`},
	},
	"Felicitous": {
		{ID: 2798, Key: "Felicitous", Value: `Felicitous`},
	},
	"Feral": {
		{ID: 146, Key: "Feral", Value: `Feral`},
	},
	"Ferocious": {
		{ID: 2013, Key: "Ferocious", Value: `Ferocious`},
	},
	"Festering": {
		{ID: 776, Key: "Festering", Value: `Festering`},
	},
	"Fierce": {
		{ID: 1282, Key: "Fierce", Value: `Fierce`},
	},
	"FindingBeneathCityAss": {
		{ID: 10, Key: "FindingBeneathCityAss", Value: `And I thought the Forgotten Tower stank.`},
	},
	"FindingBeneathCityDru": {
		{ID: 739, Key: "FindingBeneathCityDru", Value: `This smells worse than the sewers of Lut Gohlein.`},
	},
	"FindingCairnAss": {
		{ID: 2761, Key: "FindingCairnAss", Value: `These stones hold an ancient power.`},
	},
	"FindingCairnDru": {
		{ID: 2513, Key: "FindingCairnDru", Value: `Such stones are common back home.`},
	},
	"FindingDrainLeverAss": {
		{ID: 1101, Key: "FindingDrainLeverAss", Value: `Levers are made to be pulled.`},
	},
	"FindingDrainLeverDru": {
		{ID: 320, Key: "FindingDrainLeverDru", Value: `Finally... The drain lever.`},
	},
	"FindingGuardianTowerAss": {
		{ID: 1490, Key: "FindingGuardianTowerAss", Value: `Mephisto... I'm coming for you.`},
	},
	"FindingGuardianTowerDru": {
		{ID: 2248, Key: "FindingGuardianTowerDru", Value: `Hatred stirs within me.`},
	},
	"FindingInifusAss": {
		{ID: 697, Key: "FindingInifusAss", Value: `How has this tree escaped corruption?`},
	},
	"FindingInifusDru": {
		{ID: 1458, Key: "FindingInifusDru", Value: `This dead tree teems with energy.`},
	},
	"FindingJadeFigAss": {
		{ID: 2645, Key: "FindingJadeFigAss", Value: `Hmm, a jade statue. What should I do with it?`},
	},
	"FindingJadeFigDru": {
		{ID: 2044, Key: "FindingJadeFigDru", Value: `It looks like jade. Perhaps it's worth something.`},
	},
	"FindingLamEsenAss": {
		{ID: 2208, Key: "FindingLamEsenAss", Value: `Black books make for dark thoughts.`},
	},
	"FindingLamEsenDru": {
		{ID: 160, Key: "FindingLamEsenDru", Value: `An ancient manuscript...  This could be useful.`},
	},
	"FindingSummonerAss": {
		{ID: 957, Key: "FindingSummonerAss", Value: `Summoner, the dark magics have corrupted you.`},
	},
	"FindingSummonerDru": {
		{ID: 1170, Key: "FindingSummonerDru", Value: `This place would drive anyone mad.`},
	},
	"FindingTempleAss": {
		{ID: 1099, Key: "FindingTempleAss", Value: `I dread this place of darkness.`},
	},
	"FindingTempleDru": {
		{ID: 321, Key: "FindingTempleDru", Value: `This temple is a nest of evil.`},
	},
	"FindingTristramAss": {
		{ID: 2525, Key: "FindingTristramAss", Value: `Tristram... The first to fall to Diablo's wrath.`},
	},
	"FindingTristramDru": {
		{ID: 1806, Key: "FindingTristramDru", Value: `Ruins... the fate of all cities.`},
	},
	"FindingTrueTombAss": {
		{ID: 2226, Key: "FindingTrueTombAss", Value: `I can sense Tal Rasha's presence now.`},
	},
	"FindingTrueTombDru": {
		{ID: 220, Key: "FindingTrueTombDru", Value: `So, Tal Rasha... This is your resting place.`},
	},
	"FindingdecoyTombAss": {
		{ID: 761, Key: "FindingdecoyTombAss", Value: `The Horadrim have left their mark here.`},
	},
	"FindingdecoyTombDru": {
		{ID: 559, Key: "FindingdecoyTombDru", Value: `These Horadric markings are mysterious.`},
	},
	"FireBoar": {
		{ID: 473, Key: "FireBoar", Value: `Fire Boar`},
	},
	"Firelizard's Talons": {
		{ID: 1930, Key: "Firelizard's Talons", Value: `Firelizard's Talons`},
	},
	"Flamebellow": {
		{ID: 235, Key: "Flamebellow", Value: `Flamebellow`},
	},
	"Flaming": {
		{ID: 2663, Key: "Flaming", Value: `Flaming`},
	},
	"Fleshbone": {
		{ID: 1920, Key: "Fleshbone", Value: `Fleshbone`},
	},
	"Fleshrender": {
		{ID: 2277, Key: "Fleshrender", Value: `Fleshrender`},
	},
	"Fleshripper": {
		{ID: 477, Key: "Fleshripper", Value: `Fleshripper`},
	},
	"Flowkrad's Fur": {
		{ID: 2003, Key: "Flowkrad's Fur", Value: `Flowkrad's Fur`},
	},
	"Flowkrad's Grin": {
		{ID: 285, Key: "Flowkrad's Grin", Value: `Flowkrad's Grin`},
	},
	"Flowkrad's Howl": {
		{ID: 1269, Key: "Flowkrad's Howl", Value: `Flowkrad's Howl`},
	},
	"Flowkrad's Paws": {
		{ID: 2748, Key: "Flowkrad's Paws", Value: `Flowkrad's Paws`},
	},
	"Flowkrad's Sinew": {
		{ID: 289, Key: "Flowkrad's Sinew", Value: `Flowkrad's Sinew`},
	},
	"Foci of Visjerei": {
		{ID: 2524, Key: "Foci of Visjerei", Value: `Foci of the Vizjerei`},
	},
	"Foggy": {
		{ID: 1751, Key: "Foggy", Value: `Foggy`},
	},
	"Frantic": {
		{ID: 483, Key: "Frantic", Value: `Frantic`},
	},
	"Freezing": {
		{ID: 2342, Key: "Freezing", Value: `Freezing`},
	},
	"FreezingIzualAma": {
		{ID: 2514, Key: "FreezingIzualAma", Value: `Goodbye, Izual.`},
	},
	"FreezingIzualAss": {
		{ID: 1603, Key: "FreezingIzualAss", Value: `Corruption... take flight.`},
	},
	"FreezingIzualDru": {
		{ID: 2315, Key: "FreezingIzualDru", Value: `I have no pity for him. Oblivion is his reward.`},
	},
	"FrenziedHellSpawn": {
		{ID: 103, Key: "FrenziedHellSpawn", Value: `Frenzied Hell Spawn`},
	},
	"FrenziedIceSpawn": {
		{ID: 108, Key: "FrenziedIceSpawn", Value: `Frenzied Ice Spawn`},
	},
	"Frigid": {
		{ID: 1503, Key: "Frigid", Value: `Frigid`},
	},
	"Frostwind": {
		{ID: 1975, Key: "Frostwind", Value: `Frostwind`},
	},
	"Frozen Horror1": {
		{ID: 195, Key: "Frozen Horror1", Value: `Frozen Creeper`},
	},
	"Frozen Horror2": {
		{ID: 201, Key: "Frozen Horror2", Value: `Frozen Terror`},
	},
	"Frozen Horror3": {
		{ID: 202, Key: "Frozen Horror3", Value: `Frozen Scourge`},
	},
	"Frozen Horror4": {
		{ID: 203, Key: "Frozen Horror4", Value: `Frozen Horror`},
	},
	"Frozen Horror5": {
		{ID: 204, Key: "Frozen Horror5", Value: `Frozen Scorch`},
	},
	"Frozenstein": {
		{ID: 298, Key: "Frozenstein", Value: `Frozenstein`},
	},
	"Fuego Del Sol": {
		{ID: 1178, Key: "Fuego Del Sol", Value: `Fuego Del Sol`},
	},
	"Fungal": {
		{ID: 1407, Key: "Fungal", Value: `Fungal`},
	},
	"Furious": {
		{ID: 1824, Key: "Furious", Value: `Furious`},
	},
	"Gaea's": {
		{ID: 1585, Key: "Gaea's", Value: `Gaean`},
	},
	"Gaia's Wrath": {
		{ID: 1919, Key: "Gaia's Wrath", Value: `Gaia's Wrath`},
	},
	"Gargoyle's Bite": {
		{ID: 1922, Key: "Gargoyle's Bite", Value: `Gargoyle's Bite`},
	},
	"Gaudy": {
		{ID: 2205, Key: "Gaudy", Value: `Gaudy`},
	},
	"Gaya Wand": {
		{ID: 1283, Key: "Gaya Wand", Value: `Gaya Wand`},
	},
	"GemeffectX11": {
		{ID: 1969, Key: "GemeffectX11", Value: `Adds to Strength`},
	},
	"GemeffectX12": {
		{ID: 1970, Key: "GemeffectX12", Value: `Adds to Defense Rating`},
	},
	"GemeffectX13": {
		{ID: 1971, Key: "GemeffectX13", Value: `Adds to Attack Rating`},
	},
	"GemeffectX21": {
		{ID: 1972, Key: "GemeffectX21", Value: `Adds to Maximum Mana`},
	},
	"GemeffectX22": {
		{ID: 1973, Key: "GemeffectX22", Value: `Adds Resistance to Cold`},
	},
	"GemeffectX23": {
		{ID: 1974, Key: "GemeffectX23", Value: `Adds Cold Damage to Attack`},
	},
	"GemeffectX31": {
		{ID: 1978, Key: "GemeffectX31", Value: `Adds to Dexterity`},
	},
	"GemeffectX32": {
		{ID: 1979, Key: "GemeffectX32", Value: `Adds Resistance to Poison`},
	},
	"GemeffectX33": {
		{ID: 1980, Key: "GemeffectX33", Value: `Adds Poison Damage to Attack`},
	},
	"GemeffectX41": {
		{ID: 1995, Key: "GemeffectX41", Value: `Adds to Maximum Life`},
	},
	"GemeffectX42": {
		{ID: 1997, Key: "GemeffectX42", Value: `Adds Resistance to Fire`},
	},
	"GemeffectX43": {
		{ID: 1998, Key: "GemeffectX43", Value: `Adds Fire Damage to Attack`},
	},
	"GemeffectX51": {
		{ID: 1903, Key: "GemeffectX51", Value: `Adds to Attack Rating`},
	},
	"GemeffectX52": {
		{ID: 1904, Key: "GemeffectX52", Value: `Adds to All Resistances`},
	},
	"GemeffectX53": {
		{ID: 1909, Key: "GemeffectX53", Value: `Adds to Damage vs. Undead`},
	},
	"GemeffectX61": {
		{ID: 1910, Key: "GemeffectX61", Value: `Adds to Chance to Find Magic Items`},
	},
	"GemeffectX62": {
		{ID: 1911, Key: "GemeffectX62", Value: `Adds Resistance to Lightning`},
	},
	"GemeffectX63": {
		{ID: 1912, Key: "GemeffectX63", Value: `Adds Lightning Damage to Attack`},
	},
	"GemeffectX71": {
		{ID: 1914, Key: "GemeffectX71", Value: `Adds Mana and Life Regeneration`},
	},
	"GemeffectX72": {
		{ID: 1915, Key: "GemeffectX72", Value: `Adds Attacker Takes Damage`},
	},
	"GemeffectX73": {
		{ID: 1917, Key: "GemeffectX73", Value: `Adds Mana and Life Steal to Attack`},
	},
	"Gerke's Sanctuary": {
		{ID: 2516, Key: "Gerke's Sanctuary", Value: `Gerke's Sanctuary`},
	},
	"Geronimo's Fury": {
		{ID: 2814, Key: "Geronimo's Fury", Value: `Geronimo's Fury`},
	},
	"Gheed's Fortune": {
		{ID: 47, Key: "Gheed's Fortune", Value: `Gheed's Fortune`},
	},
	"Ghost Liberator": {
		{ID: 1754, Key: "Ghost Liberator", Value: `Ghost Liberator`},
	},
	"Ghostflame": {
		{ID: 922, Key: "Ghostflame", Value: `Ghostflame`},
	},
	"Ghostleach": {
		{ID: 1240, Key: "Ghostleach", Value: `Ghost Leach`},
	},
	"Ghostly": {
		{ID: 251, Key: "Ghostly", Value: `Ghostly`},
	},
	"Ghoulhide": {
		{ID: 2010, Key: "Ghoulhide", Value: `Ghoulhide`},
	},
	"Giantmaimer": {
		{ID: 32, Key: "Giantmaimer", Value: `Giant Maimer`},
	},
	"Giantskull": {
		{ID: 2288, Key: "Giantskull", Value: `Giant Skull`},
	},
	"Gillian's Brazier": {
		{ID: 1208, Key: "Gillian's Brazier", Value: `Gillian's Brazier`},
	},
	"Gimmershred": {
		{ID: 2755, Key: "Gimmershred", Value: `Gimmershred`},
	},
	"Ginther's Rift": {
		{ID: 2289, Key: "Ginther's Rift", Value: `Ginther's Rift`},
	},
	"Glacial": {
		{ID: 2313, Key: "Glacial", Value: `Glacial`},
	},
	"Glacial Caves Level 1": {
		{ID: 528, Key: "Glacial Caves Level 1", Value: `The Ancients' Way`},
	},
	"Glacial Caves Level 2": {
		{ID: 529, Key: "Glacial Caves Level 2", Value: `Icy Cellar`},
	},
	"Glimmershred": {
		{ID: 16, Key: "Glimmershred", Value: `Glimmershred`},
	},
	"Globe of Visjerei": {
		{ID: 2018, Key: "Globe of Visjerei", Value: `Globe of the Vizjerei`},
	},
	"Gloomstrap": {
		{ID: 1597, Key: "Gloomstrap", Value: `Gloom's Trap`},
	},
	"Go North": {
		{ID: 818, Key: "Go North", Value: `go north`},
	},
	"Go South": {
		{ID: 746, Key: "Go South", Value: `return to hell`},
	},
	"Godly": {
		{ID: 1762, Key: "Godly", Value: `Godly`},
	},
	"Godstrike Arch": {
		{ID: 2195, Key: "Godstrike Arch", Value: `Goldstrike Arch`},
	},
	"Golemlord's": {
		{ID: 1084, Key: "Golemlord's", Value: `Golemlord's`},
	},
	"Gore Ripper": {
		{ID: 633, Key: "Gore Ripper", Value: `Gore Ripper`},
	},
	"GoreBearer": {
		{ID: 80, Key: "GoreBearer", Value: `Gore Bearer`},
	},
	"Gorerider": {
		{ID: 2310, Key: "Gorerider", Value: `Gore Rider`},
	},
	"Grandmaster's": {
		{ID: 950, Key: "Grandmaster's", Value: `Grandmaster's`},
	},
	"Gravepalm": {
		{ID: 797, Key: "Gravepalm", Value: `Gravepalm`},
	},
	"Graverobber's": {
		{ID: 3, Key: "Graverobber's", Value: `Graverobber's`},
	},
	"Great Wyrm's": {
		{ID: 1327, Key: "Great Wyrm's", Value: `Great Wyrm's`},
	},
	"GreaterHellSpawn": {
		{ID: 225, Key: "GreaterHellSpawn", Value: `Greater Hell Spawn`},
	},
	"GreaterIceSpawn": {
		{ID: 831, Key: "GreaterIceSpawn", Value: `Greater Ice Spawn`},
	},
	"Green": {
		{ID: 89, Key: "Green", Value: `Green`},
	},
	"Griffon's Eye": {
		{ID: 2162, Key: "Griffon's Eye", Value: `Griffon's Eye`},
	},
	"Grim's Burning Dead": {
		{ID: 588, Key: "Grim's Burning Dead", Value: `Grim's Burning Dead`},
	},
	"Griswold's Heart": {
		{ID: 2002, Key: "Griswold's Heart", Value: `Griswold's Heart`},
	},
	"Griswold's Honor": {
		{ID: 2059, Key: "Griswold's Honor", Value: `Griswold's Honor`},
	},
	"Griswold's Legacy": {
		{ID: 2180, Key: "Griswold's Legacy", Value: `Griswold's Legacy`},
	},
	"Griswold's Valor": {
		{ID: 95, Key: "Griswold's Valor", Value: `Griswold's Valor`},
	},
	"Griswolds's Redemption": {
		{ID: 1216, Key: "Griswolds's Redemption", Value: `Griswold's Redemption`},
	},
	"Gritty": {
		{ID: 2116, Key: "Gritty", Value: `Grinding`},
	},
	"Guardian Angel": {
		{ID: 2055, Key: "Guardian Angel", Value: `Guardian Angel`},
	},
	"Guardian Angle": {
		{ID: 2416, Key: "Guardian Angle", Value: `Guardian Angel`},
	},
	"Guardian Naga": {
		{ID: 829, Key: "Guardian Naga", Value: `Guardian Naga`},
	},
	"Guardian's": {
		{ID: 2221, Key: "Guardian's", Value: `Guardian's`},
	},
	"Guillaume's Face": {
		{ID: 860, Key: "Guillaume's Face", Value: `Guillaume's Face`},
	},
	"Gulletwound": {
		{ID: 1747, Key: "Gulletwound", Value: `Gulletwound`},
	},
	"Gutsiphon": {
		{ID: 2770, Key: "Gutsiphon", Value: `Gut Siphon`},
	},
	"Gymnast's": {
		{ID: 13, Key: "Gymnast's", Value: `Gymnastic`},
	},
	"Hadeshorn": {
		{ID: 1179, Key: "Hadeshorn", Value: `Hadeshorn`},
	},
	"Haemosu's Adament": {
		{ID: 2536, Key: "Haemosu's Adament", Value: `Haemosu's Adamant`},
	},
	"Halaberd's Reign": {
		{ID: 1480, Key: "Halaberd's Reign", Value: `Halaberd's Reign`},
	},
	"Hallowed": {
		{ID: 96, Key: "Hallowed", Value: `Hallowed`},
	},
	"Halls of Anguish": {
		{ID: 542, Key: "Halls of Anguish", Value: `Halls of Anguish`},
	},
	"Halls of Death's Calling": {
		{ID: 1105, Key: "Halls of Death's Calling", Value: `Halls of Pain`},
	},
	"Halls of Tormented Insanity": {
		{ID: 543, Key: "Halls of Tormented Insanity", Value: `Halls of Torment`},
	},
	"Halls of Vaught": {
		{ID: 547, Key: "Halls of Vaught", Value: `Halls of Vaught`},
	},
	"Hand of Blessed Light": {
		{ID: 1418, Key: "Hand of Blessed Light", Value: `Hand of Blessed Light`},
	},
	"Harlequin Crest": {
		{ID: 662, Key: "Harlequin Crest", Value: `Harlequin Crest`},
	},
	"Harmful": {
		{ID: 816, Key: "Harmful", Value: `Harmful`},
	},
	"Harpoonist's": {
		{ID: 217, Key: "Harpoonist's", Value: `Harpoonist's`},
	},
	"Harrogath": {
		{ID: 518, Key: "Harrogath", Value: `Harrogath`},
	},
	"Hawk Branded": {
		{ID: 2206, Key: "Hawk Branded", Value: `Hawk Branded`},
	},
	"Hawkeye": {
		{ID: 1432, Key: "Hawkeye", Value: `Hawkeye`},
	},
	"Hazy": {
		{ID: 452, Key: "Hazy", Value: `Hazy`},
	},
	"Headhunter's Glory": {
		{ID: 1447, Key: "Headhunter's Glory", Value: `Head Hunter's Glory`},
	},
	"Headstriker": {
		{ID: 1769, Key: "Headstriker", Value: `Headstriker`},
	},
	"Heart Carver": {
		{ID: 1419, Key: "Heart Carver", Value: `Heart Carver`},
	},
	"Heart of Wolverine": {
		{ID: 239, Key: "Heart of Wolverine", Value: `Heart of Wolverine`},
	},
	"Heated": {
		{ID: 2193, Key: "Heated", Value: `Heated`},
	},
	"Heaven's Brethren": {
		{ID: 88, Key: "Heaven's Brethren", Value: `Heaven's Brethren`},
	},
	"Heaven's Light": {
		{ID: 2818, Key: "Heaven's Light", Value: `Heaven's Light`},
	},
	"Heaven's Taebaek": {
		{ID: 2031, Key: "Heaven's Taebaek", Value: `Taebaek's Glory`},
	},
	"Hell Temptress": {
		{ID: 113, Key: "Hell Temptress", Value: `Hell Temptress`},
	},
	"Hell Witch": {
		{ID: 118, Key: "Hell Witch", Value: `Hell Witch`},
	},
	"Hell's Whisper": {
		{ID: 39, Key: "Hell's Whisper", Value: `Hell Whisper`},
	},
	"Hell1": {
		{ID: 126, Key: "Hell1", Value: `Abaddon`},
	},
	"Hell2": {
		{ID: 129, Key: "Hell2", Value: `Pit of Acheron`},
	},
	"Hell3": {
		{ID: 130, Key: "Hell3", Value: `Infernal Pit`},
	},
	"HellSpawn": {
		{ID: 978, Key: "HellSpawn", Value: `Hell Spawn`},
	},
	"HellWhip": {
		{ID: 170, Key: "HellWhip", Value: `Hell Whip`},
	},
	"Hellatial": {
		{ID: 2411, Key: "Hellatial", Value: `Hellacious`},
	},
	"Hellmouth": {
		{ID: 1275, Key: "Hellmouth", Value: `Hellmouth`},
	},
	"Hellrack": {
		{ID: 907, Key: "Hellrack", Value: `Hellrack`},
	},
	"Hellslayer": {
		{ID: 1983, Key: "Hellslayer", Value: `Hellslayer`},
	},
	"Hellwarden's Husk": {
		{ID: 644, Key: "Hellwarden's Husk", Value: `Hell Warden's Husk`},
	},
	"Herald of Zakarum": {
		{ID: 135, Key: "Herald of Zakarum", Value: `Herald of Zakarum`},
	},
	"Hermetic": {
		{ID: 2532, Key: "Hermetic", Value: `Hermetic`},
	},
	"Hexfire": {
		{ID: 2053, Key: "Hexfire", Value: `Hexfire`},
	},
	"Hexing": {
		{ID: 708, Key: "Hexing", Value: `Hexing`},
	},
	"Hibernal": {
		{ID: 41, Key: "Hibernal", Value: `Hibernal`},
	},
	"Hierophant's": {
		{ID: 1897, Key: "Hierophant's", Value: `Hierophant's`},
	},
	"Highlord's Wrath": {
		{ID: 128, Key: "Highlord's Wrath", Value: `Highlord's Wrath`},
	},
	"Hightower's Watch": {
		{ID: 668, Key: "Hightower's Watch", Value: `Hightower's Watch`},
	},
	"Homunculus": {
		{ID: 2369, Key: "Homunculus", Value: `Homunculus`},
	},
	"Hone Sundan": {
		{ID: 1000, Key: "Hone Sundan", Value: `Hone Sundan`},
	},
	"HoradricMalusAss": {
		{ID: 857, Key: "HoradricMalusAss", Value: `A Malus! This should go to Charsi.`},
	},
	"HoradricMalusDru": {
		{ID: 617, Key: "HoradricMalusDru", Value: `Charsi will be thankful to get this Malus.`},
	},
	"Horazon's Chalice": {
		{ID: 1129, Key: "Horazon's Chalice", Value: `Horazon's Chalice`},
	},
	"Horizon's Tornado": {
		{ID: 303, Key: "Horizon's Tornado", Value: `Horizon's Tornado`},
	},
	"Howling Visage": {
		{ID: 2421, Key: "Howling Visage", Value: `Howling Visage`},
	},
	"Husoldal Evo": {
		{ID: 734, Key: "Husoldal Evo", Value: `Husoldal Evo`},
	},
	"Hwanin's Cordon": {
		{ID: 1137, Key: "Hwanin's Cordon", Value: `Hwanin's Cordon`},
	},
	"Hwanin's Justice": {
		{ID: 1788, Key: "Hwanin's Justice", Value: `Hwanin's Justice`},
	},
	"Hwanin's Majesty": {
		{ID: 2011, Key: "Hwanin's Majesty", Value: `Hwanin's Majesty`},
	},
	"Hwanin's Refuge": {
		{ID: 1495, Key: "Hwanin's Refuge", Value: `Hwanin's Refuge`},
	},
	"Hwanin's Seal": {
		{ID: 2528, Key: "Hwanin's Seal", Value: `Hwanin's Blessing`},
	},
	"Hwanin's Splendor": {
		{ID: 260, Key: "Hwanin's Splendor", Value: `Hwanin's Splendor`},
	},
	"Iansang's Frenzy": {
		{ID: 2164, Key: "Iansang's Frenzy", Value: `Iansang's Frenzy`},
	},
	"IceBoar": {
		{ID: 92, Key: "IceBoar", Value: `Ice Boar`},
	},
	"IceSpawn": {
		{ID: 784, Key: "IceSpawn", Value: `Ice Spawn`},
	},
	"Immortal King": {
		{ID: 390, Key: "Immortal King", Value: `Immortal King`},
	},
	"Immortal King's Detail": {
		{ID: 2529, Key: "Immortal King's Detail", Value: `Immortal King's Detail`},
	},
	"Immortal King's Forge": {
		{ID: 778, Key: "Immortal King's Forge", Value: `Immortal King's Forge`},
	},
	"Immortal King's Pillar": {
		{ID: 2032, Key: "Immortal King's Pillar", Value: `Immortal King's Pillar`},
	},
	"Immortal King's Soul Cage ": {
		{ID: 399, Key: "Immortal King's Soul Cage ", Value: `Immortal King's Soul Cage`},
	},
	"Immortal King's Stone Crusher": {
		{ID: 849, Key: "Immortal King's Stone Crusher", Value: `Immortal King's Stone Crusher`},
	},
	"Immortal King's Will": {
		{ID: 955, Key: "Immortal King's Will", Value: `Immortal King's Will`},
	},
	"Imp1": {
		{ID: 180, Key: "Imp1", Value: `Demon Imp`},
	},
	"Imp2": {
		{ID: 181, Key: "Imp2", Value: `Demon Rascal`},
	},
	"Imp3": {
		{ID: 182, Key: "Imp3", Value: `Demon Gremlin`},
	},
	"Imp4": {
		{ID: 187, Key: "Imp4", Value: `Demon Trickster`},
	},
	"Imp5": {
		{ID: 194, Key: "Imp5", Value: `Demon Sprite`},
	},
	"Impecable": {
		{ID: 1408, Key: "Impecable", Value: `Impeccable`},
	},
	"Indiego's Fancy": {
		{ID: 2245, Key: "Indiego's Fancy", Value: `Indiego's Fancy`},
	},
	"Infectious": {
		{ID: 226, Key: "Infectious", Value: `Infectious`},
	},
	"Inferno Sentry": {
		{ID: 243, Key: "Inferno Sentry", Value: `Inferno Sentry`},
	},
	"Infernostride": {
		{ID: 1596, Key: "Infernostride", Value: `Infernostride`},
	},
	"InsaneHellSpawn": {
		{ID: 109, Key: "InsaneHellSpawn", Value: `Insane Hell Spawn`},
	},
	"InsaneIceSpawn": {
		{ID: 1181, Key: "InsaneIceSpawn", Value: `Insane Ice Spawn`},
	},
	"Insidious": {
		{ID: 2461, Key: "Insidious", Value: `Insidious`},
	},
	"Invis Pet": {
		{ID: 165, Key: "Invis Pet", Value: `Pet`},
	},
	"Ironpelt": {
		{ID: 1250, Key: "Ironpelt", Value: `Iron Pelt`},
	},
	"Ironward": {
		{ID: 1299, Key: "Ironward", Value: `Astreon's Iron Ward`},
	},
	"Irresistible": {
		{ID: 1403, Key: "Irresistible", Value: `Irresistible`},
	},
	"Islestrike": {
		{ID: 189, Key: "Islestrike", Value: `Islestrike`},
	},
	"ItemExpansiveChanc1": {
		{ID: 557, Key: "ItemExpansiveChanc1", Value: `%d%% Chance to cast level %d %s on striking`},
	},
	"ItemExpansiveChanc2": {
		{ID: 558, Key: "ItemExpansiveChanc2", Value: `%d%% Chance to cast level %d %s when struck`},
	},
	"ItemExpansiveChancX": {
		{ID: 792, Key: "ItemExpansiveChancX", Value: `%d%% Chance to cast level %d %s on attack`},
	},
	"ItemExpcharmdesc": {
		{ID: 651, Key: "ItemExpcharmdesc", Value: `Keep in inventory to gain bonus`},
	},
	"ItemexpED": {
		{ID: 1477, Key: "ItemexpED", Value: ` Elemental Damage`},
	},
	"Ivory": {
		{ID: 918, Key: "Ivory", Value: `Ivory`},
	},
	"Jacinth": {
		{ID: 2238, Key: "Jacinth", Value: `Jacinth`},
	},
	"Jack's": {
		{ID: 1308, Key: "Jack's", Value: `Jack's`},
	},
	"Jadetalon": {
		{ID: 1450, Key: "Jadetalon", Value: `Jade Talon`},
	},
	"Jalal's Mane": {
		{ID: 2591, Key: "Jalal's Mane", Value: `Jalal's Mane`},
	},
	"Jester's": {
		{ID: 1398, Key: "Jester's", Value: `Jester's`},
	},
	"Jeweler's": {
		{ID: 2167, Key: "Jeweler's", Value: `Jeweler's`},
	},
	"Joker's": {
		{ID: 1108, Key: "Joker's", Value: `Joker's`},
	},
	"Kang's Virtue": {
		{ID: 736, Key: "Kang's Virtue", Value: `Kang's Virtue`},
	},
	"Keeper's": {
		{ID: 2721, Key: "Keeper's", Value: `Keeper's`},
	},
	"Kelpie Snare": {
		{ID: 2246, Key: "Kelpie Snare", Value: `Kelpie Snare`},
	},
	"Kenshi's": {
		{ID: 1415, Key: "Kenshi's", Value: `Kenshi's`},
	},
	"Kerke's Sanctuary": {
		{ID: 2355, Key: "Kerke's Sanctuary", Value: `Gerke's Sanctuary`},
	},
	"Khalim's Vengance": {
		{ID: 497, Key: "Khalim's Vengance", Value: `Khalim's Vengeance`},
	},
	"KillingdDiabloAms": {
		{ID: 2522, Key: "KillingdDiabloAms", Value: `The reign of Terror has ended.`},
	},
	"KillingdDiabloAss": {
		{ID: 1760, Key: "KillingdDiabloAss", Value: `A hero's mistake is finally corrected.`},
	},
	"KillingdDiabloBar": {
		{ID: 1675, Key: "KillingdDiabloBar", Value: `Eternal suffering would be too brief for you, Diablo.`},
	},
	"KillingdDiabloDru": {
		{ID: 2049, Key: "KillingdDiabloDru", Value: `Thus ends the plague of Terror.`},
	},
	"KillingdDiabloNec": {
		{ID: 416, Key: "KillingdDiabloNec", Value: `Lord Diablo I have bested you.`},
	},
	"KillingdDiabloPal": {
		{ID: 1002, Key: "KillingdDiabloPal", Value: `Let Diablo's death end the reign of the Three!`},
	},
	"KillingdDiabloSor": {
		{ID: 211, Key: "KillingdDiabloSor", Value: `Terror stalks Hell no more.`},
	},
	"Kira's Guardian": {
		{ID: 2722, Key: "Kira's Guardian", Value: `Kira's Guardian`},
	},
	"Knave's": {
		{ID: 2661, Key: "Knave's", Value: `Knave's`},
	},
	"Kuko Shakaku": {
		{ID: 999, Key: "Kuko Shakaku", Value: `Kuko Shakaku`},
	},
	"Lacerator": {
		{ID: 1712, Key: "Lacerator", Value: `Lacerator`},
	},
	"Laden": {
		{ID: 1402, Key: "Laden", Value: `Laden`},
	},
	"Lance Guard": {
		{ID: 2175, Key: "Lance Guard", Value: `Lance Guard`},
	},
	"Lanceguard": {
		{ID: 2299, Key: "Lanceguard", Value: `Lance Guard`},
	},
	"Lancer's": {
		{ID: 1473, Key: "Lancer's", Value: `Lancer's`},
	},
	"Langer Briser": {
		{ID: 1345, Key: "Langer Briser", Value: `Langer Briser`},
	},
	"Lapis Lazuli": {
		{ID: 154, Key: "Lapis Lazuli", Value: `Lapis Lazuli`},
	},
	"Larry": {
		{ID: 275, Key: "Larry", Value: `Larry`},
	},
	"Larzuk": {
		{ID: 256, Key: "Larzuk", Value: `Larzuk`},
	},
	"Larzuk's Champion": {
		{ID: 266, Key: "Larzuk's Champion", Value: `Larzuk's Champion`},
	},
	"LarzukAct5IntroAmaGossip1": {
		{ID: 2405, Key: "LarzukAct5IntroAmaGossip1", Value: `42
So, you're an Amazon. I have heard 
rumors about your kind. Your unusual 
weapons could prove a challenge to my 
skills, but I'm prepared to meet it.
 
I am Larzuk, the armorer. My ancestors 
were some of the finest craftsmen in 
Harrogath. Regretfully, my supplies run 
lower with every passing day, yet the 
demons beyond the walls have not 
weakened. I fear the time is near when 
I must put down my hammer and take 
up a sword, instead.
`},
	},
	"LarzukAct5IntroGossip1": {
		{ID: 1990, Key: "LarzukAct5IntroGossip1", Value: `55
I am Larzuk, the armorer. My ancestors 
were some of the finest craftsmen in 
Harrogath. 
 
Regretfully, my supplies run lower with 
every passing day, yet the demons 
beyond the walls have not weakened. 
 
I fear the time is near when I must put 
down my hammer and take up a sword, 
instead.
`},
	},
	"LarzukGossip1": {
		{ID: 1874, Key: "LarzukGossip1", Value: `65
I've heard that you Amazons excel at 
killing from a distance. From what I've 
seen, that's the best way to deal with 
Hell's minions. 
 
Are all of your kind so...big?
`},
	},
	"LarzukGossip10": {
		{ID: 1842, Key: "LarzukGossip10", Value: `56
Just so you know...The gold you pay me 
doesn't line my pockets. Much of it 
goes to buy the raw metal I need to 
melt down for weapons and armor. The 
rest -- well, all I can spare -- goes to 
Malah and Qual-Kehk for other 
supplies.
`},
	},
	"LarzukGossip2": {
		{ID: 1875, Key: "LarzukGossip2", Value: `65
Why did you ever leave your beautiful 
islands to come to this frozen 
battleground? Perhaps if we both 
survive this, we'll travel back there 
together.
`},
	},
	"LarzukGossip3": {
		{ID: 1876, Key: "LarzukGossip3", Value: `75
Has that infernal howling been keeping 
you awake at night, too? Some think 
it's the howling of animals, but those 
sounds don't come from any beast I've 
ever known.
`},
	},
	"LarzukGossip4": {
		{ID: 1877, Key: "LarzukGossip4", Value: `69
Demonic forces have been using our 
own towers and barricades against us. 
You know, it would be wise to cut the 
demons down in the open before they 
can mount those fortifications.
`},
	},
	"LarzukGossip5": {
		{ID: 1878, Key: "LarzukGossip5", Value: `70
Nihlathak's despair is infectious -- and 
his behavior does not befit an Elder of 
his stature. I think we'd be better off 
without him.
`},
	},
	"LarzukGossip6": {
		{ID: 1879, Key: "LarzukGossip6", Value: `63
I've offered Qual-Kehk my ideas on how 
to break the siege, but he dismisses 
them. Is it because I lack scars of 
battle, or does he think I'm a couple 
arrows short of a quiver?
`},
	},
	"LarzukGossip7": {
		{ID: 1880, Key: "LarzukGossip7", Value: `64
Legend has it that the top of Mount 
Arreat is guarded by the spirits of our 
ancestors. Though our people are 
forbidden to climb to the mountain's 
summit, some foreign travelers have 
made the attempt.
 
None have ever returned.
`},
	},
	"LarzukGossip8": {
		{ID: 1881, Key: "LarzukGossip8", Value: `66
This is a lively town during our victory 
celebrations. We Barbarians train long 
and hard from childhood to become 
warriors, and we celebrate with equal 
fervor.
 
You didn't happen to bring along any 
ale to trade?
`},
	},
	"LarzukGossip9": {
		{ID: 1882, Key: "LarzukGossip9", Value: `54
Every day, one of my friends dies 
fighting outside the town walls, while I 
tend my anvil here unscathed. If only 
we didn't need weapons so badly, I 
could be out doing my share of the 
fighting. 
 
Good luck to you, warrior.
`},
	},
	"Lasher": {
		{ID: 120, Key: "Lasher", Value: `Lasher`},
	},
	"Lasting": {
		{ID: 467, Key: "Lasting", Value: `Lasting`},
	},
	"Lavagout": {
		{ID: 2430, Key: "Lavagout", Value: `Lava Gout`},
	},
	"Lawful": {
		{ID: 1252, Key: "Lawful", Value: `Lawful`},
	},
	"Laying of Hands": {
		{ID: 2184, Key: "Laying of Hands", Value: `Laying of Hands`},
	},
	"Ldeathpole": {
		{ID: 374, Key: "Ldeathpole", Value: `Death Pole`},
	},
	"LeaveCampAss": {
		{ID: 212, Key: "LeaveCampAss", Value: `They'll never see me coming.`},
	},
	"LeaveCampDru": {
		{ID: 956, Key: "LeaveCampDru", Value: `So, it begins.`},
	},
	"LeavingTownAct5Ams": {
		{ID: 2316, Key: "LeavingTownAct5Ams", Value: `The siege must be stopped.`},
	},
	"LeavingTownAct5Ass": {
		{ID: 33, Key: "LeavingTownAct5Ass", Value: `You'll pay for your atrocities, Baal.`},
	},
	"LeavingTownAct5Bar": {
		{ID: 65, Key: "LeavingTownAct5Bar", Value: `The time has come to cleanse my homeland!`},
	},
	"LeavingTownAct5Dru": {
		{ID: 338, Key: "LeavingTownAct5Dru", Value: `Baal! Nothing will stand in my way.`},
	},
	"LeavingTownAct5Nec": {
		{ID: 301, Key: "LeavingTownAct5Nec", Value: `It takes more than a siege to stop me.`},
	},
	"LeavingTownAct5Pal": {
		{ID: 1290, Key: "LeavingTownAct5Pal", Value: `Baal. I'm coming for you.`},
	},
	"LeavingTownAct5Sor": {
		{ID: 1773, Key: "LeavingTownAct5Sor", Value: `My magic will break the siege.`},
	},
	"Lestron's Mark": {
		{ID: 2372, Key: "Lestron's Mark", Value: `Lestron's Mark`},
	},
	"Leviathan": {
		{ID: 509, Key: "Leviathan", Value: `Leviathan`},
	},
	"Lidless Wall": {
		{ID: 2173, Key: "Lidless Wall", Value: `Lidless Wall`},
	},
	"LifeSeeker": {
		{ID: 81, Key: "LifeSeeker", Value: `Life Seeker`},
	},
	"LifeStealer": {
		{ID: 82, Key: "LifeStealer", Value: `Life Stealer`},
	},
	"Lifechoke": {
		{ID: 2397, Key: "Lifechoke", Value: `Lifechoke`},
	},
	"Lightning Sentry": {
		{ID: 60, Key: "Lightning Sentry", Value: `Lightning Sentry`},
	},
	"Lightsabre": {
		{ID: 1759, Key: "Lightsabre", Value: `Lightsabre`},
	},
	"Lilac": {
		{ID: 2264, Key: "Lilac", Value: `Lilac`},
	},
	"Lion Branded": {
		{ID: 732, Key: "Lion Branded", Value: `Lion Branded`},
	},
	"Lord of Destruction": {
		{ID: 8, Key: "Lord of Destruction", Value: ``},
	},
	"Loud": {
		{ID: 586, Key: "Loud", Value: `Alarming`},
	},
	"Lucky": {
		{ID: 1399, Key: "Lucky", Value: `Lucky`},
	},
	"Lycander's Aim": {
		{ID: 2769, Key: "Lycander's Aim", Value: `Lycander's Aim`},
	},
	"Lycander's Flank": {
		{ID: 830, Key: "Lycander's Flank", Value: `Lycander's Flank`},
	},
	"M'avina's Battle Hymn": {
		{ID: 2060, Key: "M'avina's Battle Hymn", Value: `M'avina's Battle Hymn`},
	},
	"M'avina's Caster": {
		{ID: 1500, Key: "M'avina's Caster", Value: `M'avina's Caster`},
	},
	"M'avina's Embrace": {
		{ID: 2181, Key: "M'avina's Embrace", Value: `M'avina's Embrace`},
	},
	"M'avina's Icy Clutch": {
		{ID: 290, Key: "M'avina's Icy Clutch", Value: `M'avina's Icy Clutch`},
	},
	"M'avina's Tenet": {
		{ID: 2009, Key: "M'avina's Tenet", Value: `M'avina's Tenet`},
	},
	"M'avina's True Sight": {
		{ID: 341, Key: "M'avina's True Sight", Value: `M'avina's True Sight`},
	},
	"Magekiller's": {
		{ID: 1963, Key: "Magekiller's", Value: `Magekiller's`},
	},
	"Magewrath": {
		{ID: 735, Key: "Magewrath", Value: `Magewrath`},
	},
	"Magma Torquer": {
		{ID: 750, Key: "Magma Torquer", Value: `Magma Torquer`},
	},
	"Magnus' Skin": {
		{ID: 507, Key: "Magnus' Skin", Value: `Magnus' Skin`},
	},
	"Majestic": {
		{ID: 2141, Key: "Majestic", Value: `Majestic`},
	},
	"Malah": {
		{ID: 1210, Key: "Malah", Value: `Malah`},
	},
	"MalahAct5IntroBarGossip1": {
		{ID: 2505, Key: "MalahAct5IntroBarGossip1", Value: `34
You have traveled far only to return 
home to us, Barbarian. Ohh...Much has 
happened in Harrogath since you left. 
Our homeland is hardly recognizable 
with so much evil about.
 
Yet, I've managed to survive so far. 
You've seen your share of suffering as 
well, I'm sure. Seeing you again -- alive 
-- does my heart good. I shall pray that 
you can help us out of this hell.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what.
 
Nihlathak is the last survivor of the 
Council of Elders, but I'm afraid he has 
not been himself lately. The other 
Elders sacrificed themselves to place a 
protective ward around Harrogath. If 
Nihlathak is curt with you, it is because 
he feels responsible for our situation. 
He does not relish sending more of our 
people out to die.
 
So much has changed since you left 
that I see little hope for us in this life.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`},
	},
	"MalahAct5IntroGossip1": {
		{ID: 901, Key: "MalahAct5IntroGossip1", Value: `35
I, Malah, welcome you to Harrogath, 
the last stronghold of Order on Mount 
Arreat. You have come to the right 
place if you intend to defeat Baal the 
Lord of Destruction.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what. 
 
All of the Elders, save Nihlathak, 
sacrificed themselves to place a 
protective ward around Harrogath.
 
Some of us here, certainly Nihlathak, do 
not appreciate your presence. We are a 
proud people, and it is not easy for us 
to accept aid. I, however, am glad you 
are here.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`},
	},
	"MalahAct5IntroSorGossip1": {
		{ID: 2594, Key: "MalahAct5IntroSorGossip1", Value: `36
A Sorceress...Here in Harrogath?
 
There was a time, child, when I thought 
I was destined to follow your kind's 
path. However, my powers never 
developed beyond the simplest of 
spells. Although I can heal almost any 
wound with time and energy, there is 
little I can do to help against Baal.
 
But enough of that...I spend too much 
time in reflection and have forgotten 
my manners.
 
I, Malah, welcome you to Harrogath, 
the last stronghold of Order on Mount 
Arreat. You have come to the right 
place, if you intend to defeat Baal the 
Lord of Destruction.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what. 
 
All of the Elders, save Nihlathak, 
sacrificed themselves to place a 
protective ward around Harrogath.
 
Some of us here, certainly Nihlathak, do 
not appreciate your presence. We are a 
proud people, and it is not easy for us 
to accept aid. I, however, am glad you 
are here.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`},
	},
	"MalahGossip1": {
		{ID: 2228, Key: "MalahGossip1", Value: `53
I'm aware of the amazing magical 
powers a Sorceress can channel. If we 
survive Baal's wrath, I would be most 
honored if you could demonstrate and 
perhaps teach me something of what 
you know. 
 
I may be old, but I'm not dead.
`},
	},
	"MalahGossip10": {
		{ID: 2555, Key: "MalahGossip10", Value: `90
Be cautious, warrior.
 
Though I am an experienced healer, 
resurrection is beyond my powers.
`},
	},
	"MalahGossip11": {
		{ID: 2556, Key: "MalahGossip11", Value: `53
I pray for the day when the fields are 
covered in snow unstained by the blood 
of our own. Perhaps that day will come 
soon...Perhaps never...
 
Oh...But I forget myself. How may I aid 
your quest?
`},
	},
	"MalahGossip12": {
		{ID: 2557, Key: "MalahGossip12", Value: `64
Your gold is most helpful. Medical 
supplies for our wounded are scarce 
and very expensive.
 
When we can find a supplier, oh, we 
must pay dearly for what we need.
`},
	},
	"MalahGossip13": {
		{ID: 2558, Key: "MalahGossip13", Value: `44
With the exception of Qual-Kehk, the 
townspeople do not see what I see -- 
the massacre we face.
 
Our bravest, strongest heroes hobble 
back to me begging for help. I do what 
I can -- healing and bandaging some, 
preparing the others for what lies 
beyond.
`},
	},
	"MalahGossip2": {
		{ID: 2229, Key: "MalahGossip2", Value: `42
I know you and my son, Bannuk, did not 
part on the best of terms. He did not 
understand the wanderlust that drove 
you to leave your homeland. However, 
even though Bannuk could never admit 
it to me, as he grew older I could see 
that same desire in his eyes.
 
Oh...It's a pity I didn't encourage him to 
go with you. He might still be alive 
today.
`},
	},
	"MalahGossip3": {
		{ID: 2230, Key: "MalahGossip3", Value: `57
Though once considered shelter by our 
people, the Ice Caves offer no refuge 
from the dark horde. There are 
creatures there that will freeze your 
heart before feasting upon it.
`},
	},
	"MalahGossip4": {
		{ID: 2231, Key: "MalahGossip4", Value: `107
On the battlefield, turning your back on 
even the dead is unwise.
`},
	},
	"MalahGossip5": {
		{ID: 2232, Key: "MalahGossip5", Value: `60
This battle plays mind tricks on our 
warriors. Those fortunate enough to 
have returned from the mountainside 
claim to have seen angels in flight.
 
Fly they might, but that certainly does 
not make them angels.
`},
	},
	"MalahGossip6": {
		{ID: 2233, Key: "MalahGossip6", Value: `53
Perhaps you have heard the accounts 
of my son's horrible death at the hands 
of Baal's minions. He was my last living 
child...
 
The oath of compassion I have taken as 
a healer extends only to humankind.
 
Cut them down, warrior. All of them!
`},
	},
	"MalahGossip7": {
		{ID: 2234, Key: "MalahGossip7", Value: `78
The catapults are infernal machines 
made of demon flesh fused with steel.
 
Be wary of them.
`},
	},
	"MalahGossip8": {
		{ID: 2235, Key: "MalahGossip8", Value: `60
Qual-Kehk is a worthy leader, but the 
losses have borne heavily upon him. He 
is only impatient because he judges the 
worth of a warrior by action, not 
words.
 
You must prove yourself worthy of his 
trust.
`},
	},
	"MalahGossip9": {
		{ID: 2236, Key: "MalahGossip9", Value: `52
Larzuk possesses a good soul, but at 
times his mind seems quite unsound. 
 
He once asked me for twenty of my 
finest sheepskins. He said he would fill 
them with hot air and float like a cloud 
above the battlefield to spy on Baal's 
legions! 
 
I worry the siege has driven him mad.
`},
	},
	"Malicious": {
		{ID: 449, Key: "Malicious", Value: `Malicious`},
	},
	"Malignant": {
		{ID: 2593, Key: "Malignant", Value: `Malignant Skull`},
	},
	"Mang Song's Lesson": {
		{ID: 513, Key: "Mang Song's Lesson", Value: `Mang Song's Lesson`},
	},
	"Mara's Kaleidoscope": {
		{ID: 1932, Key: "Mara's Kaleidoscope", Value: `Mara's Kaleidoscope`},
	},
	"Maroon": {
		{ID: 2622, Key: "Maroon", Value: `Maroon`},
	},
	"Marred": {
		{ID: 402, Key: "Marred", Value: `Marred`},
	},
	"Marrowgrinder": {
		{ID: 2677, Key: "Marrowgrinder", Value: `Marrow Grinder`},
	},
	"Marrowwalk": {
		{ID: 1976, Key: "Marrowwalk", Value: `Marrowwalk`},
	},
	"Marshal's": {
		{ID: 1695, Key: "Marshal's", Value: `Marshal's`},
	},
	"Master's": {
		{ID: 1633, Key: "Master's", Value: `Master's`},
	},
	"McAuley's Folly": {
		{ID: 1756, Key: "McAuley's Folly", Value: `Sander's Folly`},
	},
	"McAuley's Paragon": {
		{ID: 1086, Key: "McAuley's Paragon", Value: `Sander's Paragon`},
	},
	"McAuley's Riprap": {
		{ID: 247, Key: "McAuley's Riprap", Value: `Sander's Riprap`},
	},
	"McAuley's Superstition": {
		{ID: 1122, Key: "McAuley's Superstition", Value: `Sander's Superstition`},
	},
	"McAuley's Taboo": {
		{ID: 1159, Key: "McAuley's Taboo", Value: `Sander's Taboo`},
	},
	"Meatscrape": {
		{ID: 53, Key: "Meatscrape", Value: `Meatscrape`},
	},
	"Mechanist's": {
		{ID: 896, Key: "Mechanist's", Value: `Mechanic's`},
	},
	"Medusa's Gaze": {
		{ID: 411, Key: "Medusa's Gaze", Value: `Medusa's Gaze`},
	},
	"Megaflow Rectifier": {
		{ID: 297, Key: "Megaflow Rectifier", Value: `Eldritch the Rectifier`},
	},
	"Mentalist's": {
		{ID: 902, Key: "Mentalist's", Value: `Mentalist's`},
	},
	"MercX101": {
		{ID: 841, Key: "MercX101", Value: `Vardhaka`},
	},
	"MercX102": {
		{ID: 845, Key: "MercX102", Value: `Khan`},
	},
	"MercX103": {
		{ID: 846, Key: "MercX103", Value: `Klisk`},
	},
	"MercX104": {
		{ID: 847, Key: "MercX104", Value: `Bors`},
	},
	"MercX105": {
		{ID: 859, Key: "MercX105", Value: `Brom`},
	},
	"MercX106": {
		{ID: 861, Key: "MercX106", Value: `Wiglaf`},
	},
	"MercX107": {
		{ID: 864, Key: "MercX107", Value: `Hrothgar`},
	},
	"MercX108": {
		{ID: 865, Key: "MercX108", Value: `Scyld`},
	},
	"MercX109": {
		{ID: 866, Key: "MercX109", Value: `Healfdane`},
	},
	"MercX110": {
		{ID: 867, Key: "MercX110", Value: `Heorogar`},
	},
	"MercX111": {
		{ID: 868, Key: "MercX111", Value: `Halgaunt`},
	},
	"MercX112": {
		{ID: 870, Key: "MercX112", Value: `Hygelac`},
	},
	"MercX113": {
		{ID: 872, Key: "MercX113", Value: `Egtheow`},
	},
	"MercX114": {
		{ID: 873, Key: "MercX114", Value: `Bohdan`},
	},
	"MercX115": {
		{ID: 874, Key: "MercX115", Value: `Wulfgar`},
	},
	"MercX116": {
		{ID: 881, Key: "MercX116", Value: `Hild`},
	},
	"MercX117": {
		{ID: 882, Key: "MercX117", Value: `Heatholaf`},
	},
	"MercX118": {
		{ID: 893, Key: "MercX118", Value: `Weder`},
	},
	"MercX119": {
		{ID: 910, Key: "MercX119", Value: `Vikhyat`},
	},
	"MercX120": {
		{ID: 1032, Key: "MercX120", Value: `Unferth`},
	},
	"MercX121": {
		{ID: 1036, Key: "MercX121", Value: `Sigemund`},
	},
	"MercX122": {
		{ID: 1039, Key: "MercX122", Value: `Heremod`},
	},
	"MercX123": {
		{ID: 1044, Key: "MercX123", Value: `Hengest`},
	},
	"MercX124": {
		{ID: 1045, Key: "MercX124", Value: `Folcwald`},
	},
	"MercX125": {
		{ID: 1047, Key: "MercX125", Value: `Frisian`},
	},
	"MercX126": {
		{ID: 1051, Key: "MercX126", Value: `Hnaef`},
	},
	"MercX127": {
		{ID: 1052, Key: "MercX127", Value: `Guthlaf`},
	},
	"MercX128": {
		{ID: 1053, Key: "MercX128", Value: `Oslaf`},
	},
	"MercX129": {
		{ID: 1063, Key: "MercX129", Value: `Yrmenlaf`},
	},
	"MercX130": {
		{ID: 1064, Key: "MercX130", Value: `Garmund`},
	},
	"MercX131": {
		{ID: 1065, Key: "MercX131", Value: `Freawaru`},
	},
	"MercX132": {
		{ID: 1077, Key: "MercX132", Value: `Eadgils`},
	},
	"MercX133": {
		{ID: 1079, Key: "MercX133", Value: `Onela`},
	},
	"MercX134": {
		{ID: 1080, Key: "MercX134", Value: `Damien`},
	},
	"MercX135": {
		{ID: 1082, Key: "MercX135", Value: `Erfor`},
	},
	"MercX136": {
		{ID: 1116, Key: "MercX136", Value: `Weohstan`},
	},
	"MercX137": {
		{ID: 1118, Key: "MercX137", Value: `Wulf`},
	},
	"MercX138": {
		{ID: 1164, Key: "MercX138", Value: `Bulwye`},
	},
	"MercX139": {
		{ID: 1165, Key: "MercX139", Value: `Lief`},
	},
	"MercX140": {
		{ID: 1166, Key: "MercX140", Value: `Magnus`},
	},
	"MercX141": {
		{ID: 1188, Key: "MercX141", Value: `Klatu`},
	},
	"MercX142": {
		{ID: 1189, Key: "MercX142", Value: `Drus`},
	},
	"MercX143": {
		{ID: 1190, Key: "MercX143", Value: `Hoku`},
	},
	"MercX144": {
		{ID: 1191, Key: "MercX144", Value: `Kord`},
	},
	"MercX145": {
		{ID: 1196, Key: "MercX145", Value: `Uther`},
	},
	"MercX146": {
		{ID: 1197, Key: "MercX146", Value: `Ip`},
	},
	"MercX147": {
		{ID: 1198, Key: "MercX147", Value: `Ulf`},
	},
	"MercX148": {
		{ID: 1200, Key: "MercX148", Value: `Tharr`},
	},
	"MercX149": {
		{ID: 1202, Key: "MercX149", Value: `Kaelim`},
	},
	"MercX150": {
		{ID: 1203, Key: "MercX150", Value: `Ulric`},
	},
	"MercX151": {
		{ID: 1214, Key: "MercX151", Value: `Alaric`},
	},
	"MercX152": {
		{ID: 1228, Key: "MercX152", Value: `Ethelred`},
	},
	"MercX153": {
		{ID: 1229, Key: "MercX153", Value: `Caden`},
	},
	"MercX154": {
		{ID: 1230, Key: "MercX154", Value: `Elgifu`},
	},
	"MercX155": {
		{ID: 1232, Key: "MercX155", Value: `Tostig`},
	},
	"MercX156": {
		{ID: 1235, Key: "MercX156", Value: `Alcuin`},
	},
	"MercX157": {
		{ID: 1237, Key: "MercX157", Value: `Emund`},
	},
	"MercX158": {
		{ID: 1239, Key: "MercX158", Value: `Sigurd`},
	},
	"MercX159": {
		{ID: 1244, Key: "MercX159", Value: `Gorm`},
	},
	"MercX160": {
		{ID: 1246, Key: "MercX160", Value: `Hollis`},
	},
	"MercX161": {
		{ID: 1254, Key: "MercX161", Value: `Ragnar`},
	},
	"MercX162": {
		{ID: 1267, Key: "MercX162", Value: `Torkel`},
	},
	"MercX163": {
		{ID: 1272, Key: "MercX163", Value: `Wulfstan`},
	},
	"MercX164": {
		{ID: 1274, Key: "MercX164", Value: `Alban`},
	},
	"MercX165": {
		{ID: 1279, Key: "MercX165", Value: `Barloc`},
	},
	"MercX166": {
		{ID: 1280, Key: "MercX166", Value: `Bill`},
	},
	"MercX167": {
		{ID: 1281, Key: "MercX167", Value: `Theodoric`},
	},
	"Merman's Speed": {
		{ID: 2225, Key: "Merman's Speed", Value: `Merman's Sprocket`},
	},
	"Messerschmidt's Reaver": {
		{ID: 848, Key: "Messerschmidt's Reaver", Value: `Messerschmidt's Reaver`},
	},
	"Metalgird": {
		{ID: 1091, Key: "Metalgird", Value: `Metalite's Girth`},
	},
	"Metalgrid": {
		{ID: 40, Key: "Metalgrid", Value: `Metalgrid`},
	},
	"MiniPanelHire": {
		{ID: 794, Key: "MiniPanelHire", Value: `Hireling`},
	},
	"MiniPanelHireinv": {
		{ID: 1029, Key: "MiniPanelHireinv", Value: `Hireling Inventory`},
	},
	"MinionGreaterIce/hellSpawnSpawner": {
		{ID: 175, Key: "MinionGreaterIce/hellSpawnSpawner", Value: `Demon Portal`},
	},
	"MinionIce/fireBoarSpawner": {
		{ID: 406, Key: "MinionIce/fireBoarSpawner", Value: `Demon Portal`},
	},
	"MinionSlayerSpawner": {
		{ID: 171, Key: "MinionSlayerSpawner", Value: `Demon Portal`},
	},
	"MinionSpawner": {
		{ID: 842, Key: "MinionSpawner", Value: `Demon Portal`},
	},
	"Minionexp": {
		{ID: 400, Key: "Minionexp", Value: `Enslaved`},
	},
	"Minionice/hellSpawnSpawner": {
		{ID: 172, Key: "Minionice/hellSpawnSpawner", Value: `Demon Portal`},
	},
	"Miocene": {
		{ID: 919, Key: "Miocene", Value: `Enlightened`},
	},
	"Mnemonic": {
		{ID: 1083, Key: "Mnemonic", Value: `Mnemonic`},
	},
	"ModStre10L": {
		{ID: 700, Key: "ModStre10L", Value: `Percent Chance to Cast`},
	},
	"ModStre10a": {
		{ID: 720, Key: "ModStre10a", Value: `Charges`},
	},
	"ModStre10b": {
		{ID: 721, Key: "ModStre10b", Value: `Level`},
	},
	"ModStre10c": {
		{ID: 722, Key: "ModStre10c", Value: `Per Level`},
	},
	"ModStre10d": {
		{ID: 723, Key: "ModStre10d", Value: `(%d/%d Charges)`},
	},
	"ModStre10e": {
		{ID: 724, Key: "ModStre10e", Value: `(`},
	},
	"ModStre10f": {
		{ID: 725, Key: "ModStre10f", Value: `)`},
	},
	"ModStre10g": {
		{ID: 726, Key: "ModStre10g", Value: `Stealth`},
	},
	"ModStre10h": {
		{ID: 728, Key: "ModStre10h", Value: `Immunity to Poison`},
	},
	"ModStre10i": {
		{ID: 729, Key: "ModStre10i", Value: `Cursed`},
	},
	"ModStre10j": {
		{ID: 733, Key: "ModStre10j", Value: `Per Player in Party`},
	},
	"ModStre10k": {
		{ID: 737, Key: "ModStre10k", Value: `Kick Damage:`},
	},
	"ModStre8a": {
		{ID: 2095, Key: "ModStre8a", Value: `to Druid Skills`},
	},
	"ModStre8b": {
		{ID: 2096, Key: "ModStre8b", Value: `to Assassin Skills`},
	},
	"ModStre8c": {
		{ID: 2097, Key: "ModStre8c", Value: ` Sockets`},
	},
	"ModStre8d": {
		{ID: 2098, Key: "ModStre8d", Value: ` to Attack Rating vs. Demons`},
	},
	"ModStre8e": {
		{ID: 2099, Key: "ModStre8e", Value: ` to Attack Rating vs. Undead`},
	},
	"ModStre8f": {
		{ID: 2100, Key: "ModStre8f", Value: ` to Damage vs. Demons`},
	},
	"ModStre8g": {
		{ID: 2111, Key: "ModStre8g", Value: ` percent to Attack Rating`},
	},
	"ModStre8h": {
		{ID: 2112, Key: "ModStre8h", Value: ` to Javelin and Spear Skills`},
	},
	"ModStre8i": {
		{ID: 2113, Key: "ModStre8i", Value: ` to Passive and Magic Skills`},
	},
	"ModStre8j": {
		{ID: 2114, Key: "ModStre8j", Value: ` to Bow and Crossbow Skills`},
	},
	"ModStre8k": {
		{ID: 2118, Key: "ModStre8k", Value: ` to Defensive Aura Skills`},
	},
	"ModStre8l": {
		{ID: 2119, Key: "ModStre8l", Value: ` to Offensive Aura Skills`},
	},
	"ModStre8m": {
		{ID: 2120, Key: "ModStre8m", Value: ` to Combat Skills`},
	},
	"ModStre8n": {
		{ID: 2121, Key: "ModStre8n", Value: ` to Summoning Skills`},
	},
	"ModStre8o": {
		{ID: 2122, Key: "ModStre8o", Value: ` to Poison and Bone Skills`},
	},
	"ModStre8p": {
		{ID: 2078, Key: "ModStre8p", Value: ` to Curses`},
	},
	"ModStre8q": {
		{ID: 2079, Key: "ModStre8q", Value: ` to Warcry Skills`},
	},
	"ModStre8r": {
		{ID: 2080, Key: "ModStre8r", Value: ` to Combat Skills`},
	},
	"ModStre8s": {
		{ID: 2081, Key: "ModStre8s", Value: ` to Masteries Skills`},
	},
	"ModStre8t": {
		{ID: 2082, Key: "ModStre8t", Value: ` to Cold Skills`},
	},
	"ModStre8u": {
		{ID: 2094, Key: "ModStre8u", Value: ` to Lightning Skills`},
	},
	"ModStre8v": {
		{ID: 2123, Key: "ModStre8v", Value: ` to Fire Skills`},
	},
	"ModStre8w": {
		{ID: 2124, Key: "ModStre8w", Value: ` to Summoning Skills`},
	},
	"ModStre8x": {
		{ID: 2126, Key: "ModStre8x", Value: ` to Shape-Shifting Skills`},
	},
	"ModStre8y": {
		{ID: 2128, Key: "ModStre8y", Value: ` to Elemental Skills`},
	},
	"ModStre8z": {
		{ID: 2129, Key: "ModStre8z", Value: ` to Trap Skills`},
	},
	"ModStre9a": {
		{ID: 2130, Key: "ModStre9a", Value: ` to Shadow Discipline Skills`},
	},
	"ModStre9b": {
		{ID: 2137, Key: "ModStre9b", Value: ` to Martial Art Skills`},
	},
	"ModStre9c": {
		{ID: 2138, Key: "ModStre9c", Value: `(Based on Character Level)`},
	},
	"ModStre9d": {
		{ID: 2140, Key: "ModStre9d", Value: `(Increases During Nighttime)`},
	},
	"ModStre9e": {
		{ID: 2143, Key: "ModStre9e", Value: `(Increases During Daytime)`},
	},
	"ModStre9f": {
		{ID: 2146, Key: "ModStre9f", Value: `(Increases Near Dawn)`},
	},
	"ModStre9g": {
		{ID: 2148, Key: "ModStre9g", Value: `(Increases Near Dusk)`},
	},
	"ModStre9h": {
		{ID: 2149, Key: "ModStre9h", Value: ` Charges of`},
	},
	"ModStre9i": {
		{ID: 2157, Key: "ModStre9i", Value: `Increased Stack Size`},
	},
	"ModStre9s": {
		{ID: 2257, Key: "ModStre9s", Value: `Indestructible`},
	},
	"ModStre9t": {
		{ID: 2259, Key: "ModStre9t", Value: `Repairs %d durability per second`},
	},
	"ModStre9u": {
		{ID: 2263, Key: "ModStre9u", Value: `Repairs %d durability in %d seconds`},
	},
	"ModStre9v": {
		{ID: 2265, Key: "ModStre9v", Value: `Replenishes quantity`},
	},
	"ModStre9w": {
		{ID: 2266, Key: "ModStre9w", Value: `Cast a Level %d`},
	},
	"ModStre9x": {
		{ID: 2269, Key: "ModStre9x", Value: `When You Swing`},
	},
	"ModStre9y": {
		{ID: 2273, Key: "ModStre9y", Value: `When You Get Hit`},
	},
	"ModStre9z": {
		{ID: 2275, Key: "ModStre9z", Value: `When You Hit an Enemy`},
	},
	"Moe": {
		{ID: 269, Key: "Moe", Value: `Moe`},
	},
	"Moldy": {
		{ID: 2239, Key: "Moldy", Value: `Moldy`},
	},
	"Monumental": {
		{ID: 2579, Key: "Monumental", Value: `Monumental`},
	},
	"Moonfall": {
		{ID: 1156, Key: "Moonfall", Value: `Moonfall`},
	},
	"Moonrend": {
		{ID: 2533, Key: "Moonrend", Value: `Moonrend`},
	},
	"Moonshadow": {
		{ID: 1030, Key: "Moonshadow", Value: `Moon Shadow`},
	},
	"Mordoc's marauder": {
		{ID: 1571, Key: "Mordoc's marauder", Value: `Mordoc's Marauder`},
	},
	"Mortal Crescent": {
		{ID: 2285, Key: "Mortal Crescent", Value: `Mortal Crescent`},
	},
	"Moser's Blessed Circle": {
		{ID: 249, Key: "Moser's Blessed Circle", Value: `Moser's Blessed Circle`},
	},
	"Mosers Blessed Circle": {
		{ID: 1058, Key: "Mosers Blessed Circle", Value: `Moser's Blessed Circle`},
	},
	"Musty": {
		{ID: 673, Key: "Musty", Value: `Musty`},
	},
	"Nagas": {
		{ID: 640, Key: "Nagas", Value: `Nagas`},
	},
	"Naj's Ancient Set": {
		{ID: 1174, Key: "Naj's Ancient Set", Value: `Naj's Ancient Vestige`},
	},
	"Naj's Circlet": {
		{ID: 2222, Key: "Naj's Circlet", Value: `Naj's Circlet`},
	},
	"Naj's Light Plate": {
		{ID: 2437, Key: "Naj's Light Plate", Value: `Naj's Light Plate`},
	},
	"Naj's Puzzler": {
		{ID: 42, Key: "Naj's Puzzler", Value: `Naj's Puzzler`},
	},
	"Natalya's Mark": {
		{ID: 2478, Key: "Natalya's Mark", Value: `Natalya's Mark`},
	},
	"Natalya's Odium": {
		{ID: 1123, Key: "Natalya's Odium", Value: `Natalya's Odium`},
	},
	"Natalya's Shadow": {
		{ID: 317, Key: "Natalya's Shadow", Value: `Natalya's Shadow`},
	},
	"Natalya's Soul": {
		{ID: 2179, Key: "Natalya's Soul", Value: `Natalya's Soul`},
	},
	"Natalya's Totem": {
		{ID: 2587, Key: "Natalya's Totem", Value: `Natalya's Totem`},
	},
	"Nature's": {
		{ID: 854, Key: "Nature's", Value: `Natural`},
	},
	"Nature's Peace": {
		{ID: 190, Key: "Nature's Peace", Value: `Nature's Peace`},
	},
	"Nebucaneezer's Storm": {
		{ID: 2584, Key: "Nebucaneezer's Storm", Value: `Nebuchadnezzar's Storm`},
	},
	"NecOnly": {
		{ID: 2101, Key: "NecOnly", Value: `(Necromancer Only)`},
	},
	"Nethercrow": {
		{ID: 1095, Key: "Nethercrow", Value: `Nethercrow`},
	},
	"Nickel": {
		{ID: 1887, Key: "Nickel", Value: `Nickel`},
	},
	"Nightsummon": {
		{ID: 134, Key: "Nightsummon", Value: `Nightsummon`},
	},
	"Nightwing's Veil": {
		{ID: 880, Key: "Nightwing's Veil", Value: `Nightwing's Veil`},
	},
	"Nihlathak": {
		{ID: 267, Key: "Nihlathak", Value: `Nihlathak`},
	},
	"Nihlathak Town": {
		{ID: 262, Key: "Nihlathak Town", Value: `Nihlathak`},
	},
	"NihlathakAct5IntroAssGossip1": {
		{ID: 795, Key: "NihlathakAct5IntroAssGossip1", Value: `54
Well, well...An Assassin!
 
Heh, heh...While I am sure we are all 
grateful for your presence in our 
troubled town, you need not have made 
the journey.
 
I can personally say that your skills are 
not required here. You would serve 
your clan better elsewhere.
`},
	},
	"NihlathakAct5IntroGossip1": {
		{ID: 279, Key: "NihlathakAct5IntroGossip1", Value: `43
Well, well. The siege has everything in 
short supply...except fools. 
 
Why would you seek this place, 
stranger? Are you a vulture come to 
loot the bodies of our fallen warriors? 
 
Regardless, this is no place to make a 
name for yourself. The mountain is 
ours to protect. It is only a matter of 
time before Hell's legions are routed.
`},
	},
	"NihlathakAct5IntroNecGossip1": {
		{ID: 137, Key: "NihlathakAct5IntroNecGossip1", Value: `45
Ahhh, a Necromancer.
 
While I admire your courage in seeking 
out the darker side of magic, we really 
have little need of your skills. The 
battle will turn soon enough without 
your meddling. 
 
Yet, I should have expected to see your 
kind here. You are like a moth to the 
flame -- drawn to all this death. It 
feeds you in more ways than one, does 
it not?
`},
	},
	"NihlathakGossip1": {
		{ID: 940, Key: "NihlathakGossip1", Value: `49
If you're looking for cases of 
treacherous magic, Assassin, take a 
hard look at Larzuk. He was the only 
one in town who escaped the Red Fever 
last spring. He claimed his good 
fortune was due to 'hand washing' 
before meals.
 
Hmmm...Very suspicious...
`},
	},
	"NihlathakGossip2": {
		{ID: 941, Key: "NihlathakGossip2", Value: `48
Why would you seek this place, 
stranger? Are you a vulture come to 
loot the bodies of our fallen warriors? 
 
Regardless, this is no place to make a 
name for yourself. The mountain is 
ours to protect. It is only a matter of 
time before Hell's legions are routed.
`},
	},
	"NihlathakGossip3": {
		{ID: 942, Key: "NihlathakGossip3", Value: `67
Qual-Kehk is useless. He has blindly 
sent our warriors to their deaths, 
assuming Baal's legions would fight as 
men do. Of course, everyone knows 
they do not.
`},
	},
	"NihlathakGossip4": {
		{ID: 943, Key: "NihlathakGossip4", Value: `57
The demon hordes have grown powerful 
beyond measure, aided by our foolish 
mistakes. But I may have found a way 
to correct those mistakes...
`},
	},
	"NihlathakGossip5": {
		{ID: 944, Key: "NihlathakGossip5", Value: `37
Oh yes...I remember our warriors as 
children. Malah would set their broken 
bones and give them powders for their 
fevers. Now, they return to her with 
wounds that will never heal.
 
I am tired...Please...leave me.
`},
	},
	"NihlathakGossip6": {
		{ID: 945, Key: "NihlathakGossip6", Value: `133
If you have nothing useful for me, be on 
your way!
`},
	},
	"NihlathakGossip7": {
		{ID: 946, Key: "NihlathakGossip7", Value: `81
Anya's father was my good friend. 
There are so many to mourn...I have 
no time for you!
`},
	},
	"NihlathakGossip8": {
		{ID: 947, Key: "NihlathakGossip8", Value: `53
I have long been criticized, but 
especially of late -- since the deaths of 
my fellow Elders. Through it all, I have 
learned one thing. Each man must do 
what's right, no matter what others 
may think.
`},
	},
	"NihlathakGossip9": {
		{ID: 948, Key: "NihlathakGossip9", Value: `65
The Council of Elders always believed 
itself prepared for the coming of the 
Three. Obviously, we were not prepared 
enough.
`},
	},
	"Nihlathaks Temple": {
		{ID: 532, Key: "Nihlathaks Temple", Value: `Nihlathak's Temple`},
	},
	"Nord's Tenderizer": {
		{ID: 238, Key: "Nord's Tenderizer", Value: `Nord's Tenderizer`},
	},
	"Nosferatu's Coil": {
		{ID: 2523, Key: "Nosferatu's Coil", Value: `Nosferatu's Coil`},
	},
	"Noxious": {
		{ID: 1437, Key: "Noxious", Value: `Noxious`},
	},
	"Null": {
		{ID: 1720, Key: "Null", Value: `Null`},
	},
	"Oak Sage": {
		{ID: 236, Key: "Oak Sage", Value: `Oak Sage`},
	},
	"Obsidian": {
		{ID: 510, Key: "Obsidian", Value: `Obsidian`},
	},
	"Odium": {
		{ID: 358, Key: "Odium", Value: `Odium`},
	},
	"Of Abrasion": {
		{ID: 1900, Key: "Of Abrasion", Value: `of Abrasion`},
	},
	"Of Admiration": {
		{ID: 2041, Key: "Of Admiration", Value: `of Admiration`},
	},
	"Of Attrition": {
		{ID: 1042, Key: "Of Attrition", Value: `of Attrition`},
	},
	"Of Baddass": {
		{ID: 535, Key: "Of Baddass", Value: `of the Abyss`},
	},
	"Of Badness": {
		{ID: 1293, Key: "Of Badness", Value: `of Badness`},
	},
	"Of Beauty": {
		{ID: 295, Key: "Of Beauty", Value: `of Beauty`},
	},
	"Of Bile": {
		{ID: 674, Key: "Of Bile", Value: `of Bile`},
	},
	"Of Blitzen": {
		{ID: 155, Key: "Of Blitzen", Value: `of Blitzen`},
	},
	"Of Combat": {
		{ID: 2742, Key: "Of Combat", Value: `of Combat`},
	},
	"Of Credit": {
		{ID: 1158, Key: "Of Credit", Value: `of Credit`},
	},
	"Of Cremation": {
		{ID: 2642, Key: "Of Cremation", Value: `of Cremation`},
	},
	"Of Cruelty": {
		{ID: 2363, Key: "Of Cruelty", Value: `of Cruelty`},
	},
	"Of Daring": {
		{ID: 1488, Key: "Of Daring", Value: `of Daring`},
	},
	"Of Darkness": {
		{ID: 2168, Key: "Of Darkness", Value: `of Darkness`},
	},
	"Of Dawn": {
		{ID: 1204, Key: "Of Dawn", Value: `of Dawn`},
	},
	"Of Disease": {
		{ID: 1872, Key: "Of Disease", Value: `of Disease`},
	},
	"Of Dispatch": {
		{ID: 1522, Key: "Of Dispatch", Value: `of Dispatch`},
	},
	"Of Doom": {
		{ID: 1901, Key: "Of Doom", Value: `of Doom`},
	},
	"Of Dread": {
		{ID: 929, Key: "Of Dread", Value: `of Dread`},
	},
	"Of Dusk": {
		{ID: 747, Key: "Of Dusk", Value: `of Dusk`},
	},
	"Of Elusion": {
		{ID: 116, Key: "Of Elusion", Value: `of Eluding`},
	},
	"Of Erosion": {
		{ID: 1008, Key: "Of Erosion", Value: `of Erosion`},
	},
	"Of Fortication": {
		{ID: 61, Key: "Of Fortication", Value: `of Fortification`},
	},
	"Of Grounding": {
		{ID: 1474, Key: "Of Grounding", Value: `of Grounding`},
	},
	"Of Honor": {
		{ID: 932, Key: "Of Honor", Value: `of Honor`},
	},
	"Of Horror": {
		{ID: 1965, Key: "Of Horror", Value: `of Horror`},
	},
	"Of Illusion": {
		{ID: 2691, Key: "Of Illusion", Value: `of Illusion`},
	},
	"Of Judgement": {
		{ID: 1962, Key: "Of Judgement", Value: `of Judgement`},
	},
	"Of Karma": {
		{ID: 2360, Key: "Of Karma", Value: `of Karma`},
	},
	"Of Love": {
		{ID: 920, Key: "Of Love", Value: `of Love`},
	},
	"Of Luck": {
		{ID: 2249, Key: "Of Luck", Value: `of Luck`},
	},
	"Of Nobility": {
		{ID: 599, Key: "Of Nobility", Value: `of Nobility`},
	},
	"Of Pilfering": {
		{ID: 2183, Key: "Of Pilfering", Value: `of Pilfering`},
	},
	"Of Quickening": {
		{ID: 1719, Key: "Of Quickening", Value: `of Quickening`},
	},
	"Of Quota": {
		{ID: 1233, Key: "Of Quota", Value: `of the Quota`},
	},
	"Of Redemption": {
		{ID: 1961, Key: "Of Redemption", Value: `of Redemption`},
	},
	"Of Remorse": {
		{ID: 1935, Key: "Of Remorse", Value: `of Remorse`},
	},
	"Of Searing": {
		{ID: 2223, Key: "Of Searing", Value: `of Searing`},
	},
	"Of Stature": {
		{ID: 1288, Key: "Of Stature", Value: `of Stature`},
	},
	"Of Stone": {
		{ID: 622, Key: "Of Stone", Value: `of Stone`},
	},
	"Of Suffering": {
		{ID: 1475, Key: "Of Suffering", Value: `of Suffering`},
	},
	"Of Sweetness": {
		{ID: 504, Key: "Of Sweetness", Value: `of Sweetness`},
	},
	"Of Terror": {
		{ID: 2328, Key: "Of Terror", Value: `of Terror`},
	},
	"Of Tribute": {
		{ID: 1018, Key: "Of Tribute", Value: `of Tribute`},
	},
	"Of Twilight": {
		{ID: 1441, Key: "Of Twilight", Value: `of Twilight`},
	},
	"Of Valhalla": {
		{ID: 1061, Key: "Of Valhalla", Value: `of Valhalla`},
	},
	"Of Vengence": {
		{ID: 1744, Key: "Of Vengence", Value: `of Vengence`},
	},
	"Of Vines": {
		{ID: 74, Key: "Of Vines", Value: `of Vines`},
	},
	"Of Waste": {
		{ID: 2695, Key: "Of Waste", Value: `of Waste`},
	},
	"Of self-repair": {
		{ID: 523, Key: "Of self-repair", Value: `Self-Repair`},
	},
	"Of the Avenger": {
		{ID: 293, Key: "Of the Avenger", Value: `of the Avenger`},
	},
	"Of the Bayou": {
		{ID: 875, Key: "Of the Bayou", Value: `of the Bayou`},
	},
	"Of the Beast": {
		{ID: 1592, Key: "Of the Beast", Value: `of the Beast`},
	},
	"Of the Choir": {
		{ID: 661, Key: "Of the Choir", Value: `Of the Choir`},
	},
	"Of the Earth": {
		{ID: 2142, Key: "Of the Earth", Value: `of the Earth`},
	},
	"Of the Fly": {
		{ID: 2115, Key: "Of the Fly", Value: `of the Fly`},
	},
	"Of the Forest": {
		{ID: 1966, Key: "Of the Forest", Value: `of the Forest`},
	},
	"Of the Ghost": {
		{ID: 2274, Key: "Of the Ghost", Value: `of the Ghost`},
	},
	"Of the Gnat": {
		{ID: 2474, Key: "Of the Gnat", Value: `of the Gnat`},
	},
	"Of the Grassy Gnoll": {
		{ID: 2462, Key: "Of the Grassy Gnoll", Value: `of the Grassy Gnoll`},
	},
	"Of the Hag": {
		{ID: 2305, Key: "Of the Hag", Value: `of the Hag`},
	},
	"Of the Horde": {
		{ID: 900, Key: "Of the Horde", Value: `of the Horde`},
	},
	"Of the Idiot": {
		{ID: 394, Key: "Of the Idiot", Value: `of the Idiot`},
	},
	"Of the Imbecile": {
		{ID: 1416, Key: "Of the Imbecile", Value: `of the Imbecile`},
	},
	"Of the Infantry": {
		{ID: 2272, Key: "Of the Infantry", Value: `of the Infantry`},
	},
	"Of the Jujube": {
		{ID: 1075, Key: "Of the Jujube", Value: `of the Jujube`},
	},
	"Of the Lady": {
		{ID: 1873, Key: "Of the Lady", Value: `of the Lady`},
	},
	"Of the Lake": {
		{ID: 2270, Key: "Of the Lake", Value: `of the Lake`},
	},
	"Of the Lilly": {
		{ID: 963, Key: "Of the Lilly", Value: `of the Lilly`},
	},
	"Of the Maggot": {
		{ID: 438, Key: "Of the Maggot", Value: `of Maggots`},
	},
	"Of the Maiden": {
		{ID: 793, Key: "Of the Maiden", Value: `Of the Maiden`},
	},
	"Of the Moon": {
		{ID: 2590, Key: "Of the Moon", Value: `of the Moon`},
	},
	"Of the Mosquito": {
		{ID: 2020, Key: "Of the Mosquito", Value: `of the Mosquito`},
	},
	"Of the Obscenity": {
		{ID: 192, Key: "Of the Obscenity", Value: `of the Obscene`},
	},
	"Of the Ocean": {
		{ID: 1699, Key: "Of the Ocean", Value: `of the Ocean`},
	},
	"Of the Plague": {
		{ID: 1964, Key: "Of the Plague", Value: `of the Plague`},
	},
	"Of the Retard": {
		{ID: 933, Key: "Of the Retard", Value: `of the Moron`},
	},
	"Of the River": {
		{ID: 2626, Key: "Of the River", Value: `of the River`},
	},
	"Of the Sky": {
		{ID: 87, Key: "Of the Sky", Value: `of the Sky`},
	},
	"Of the Sniper": {
		{ID: 744, Key: "Of the Sniper", Value: `Of the Sniper`},
	},
	"Of the Specter": {
		{ID: 2194, Key: "Of the Specter", Value: `of the Specter`},
	},
	"Of the Stars": {
		{ID: 1590, Key: "Of the Stars", Value: `of the Stars`},
	},
	"Of the Stiletto": {
		{ID: 1081, Key: "Of the Stiletto", Value: `Of the Stiletto`},
	},
	"Of the Stream": {
		{ID: 856, Key: "Of the Stream", Value: `of the Stream`},
	},
	"Of the Unicorn": {
		{ID: 257, Key: "Of the Unicorn", Value: `of the Unicorn`},
	},
	"Of the Virgin": {
		{ID: 2339, Key: "Of the Virgin", Value: `of the Virgin`},
	},
	"Of the Walrus": {
		{ID: 2475, Key: "Of the Walrus", Value: `of the Walrus`},
	},
	"Of the Witch": {
		{ID: 1890, Key: "Of the Witch", Value: `of the Witch`},
	},
	"Ogun's Fierce Visage": {
		{ID: 2618, Key: "Ogun's Fierce Visage", Value: `Ogun's Fierce Visage`},
	},
	"Ogun's Lash": {
		{ID: 110, Key: "Ogun's Lash", Value: `Ogun's Lash`},
	},
	"Ogun's Shadow": {
		{ID: 2619, Key: "Ogun's Shadow", Value: `Ogun's Shadow`},
	},
	"Ogun's Vengeance": {
		{ID: 2022, Key: "Ogun's Vengeance", Value: `Ogun's Vengeance`},
	},
	"Oligocene": {
		{ID: 1957, Key: "Oligocene", Value: `Honorable`},
	},
	"Ondal's Almighty": {
		{ID: 481, Key: "Ondal's Almighty", Value: `Ondal's Almighty`},
	},
	"Ondal's Wisdom": {
		{ID: 970, Key: "Ondal's Wisdom", Value: `Ondal's Wisdom`},
	},
	"Original": {
		{ID: 2553, Key: "Original", Value: `Original`},
	},
	"Ormus' Robes": {
		{ID: 2, Key: "Ormus' Robes", Value: `Ormus' Robes`},
	},
	"Orphan's Call": {
		{ID: 1682, Key: "Orphan's Call", Value: `Orphan's Call`},
	},
	"Orumus' Robes": {
		{ID: 2540, Key: "Orumus' Robes", Value: `Ormus' Robe`},
	},
	"OverLord": {
		{ID: 168, Key: "OverLord", Value: `Overlord`},
	},
	"OverSeer": {
		{ID: 119, Key: "OverSeer", Value: `Overseer`},
	},
	"POW": {
		{ID: 268, Key: "POW", Value: `Barbarian Captive`},
	},
	"Pagan's Athame": {
		{ID: 820, Key: "Pagan's Athame", Value: `Pagan's Athame`},
	},
	"Pain Worm1": {
		{ID: 451, Key: "Pain Worm1", Value: `Pain Worm`},
	},
	"Pain Worm2": {
		{ID: 453, Key: "Pain Worm2", Value: `Torment Worm`},
	},
	"Pain Worm3": {
		{ID: 459, Key: "Pain Worm3", Value: `Agony Worm`},
	},
	"Pain Worm4": {
		{ID: 460, Key: "Pain Worm4", Value: `Menace Worm`},
	},
	"Pain Worm5": {
		{ID: 461, Key: "Pain Worm5", Value: `Anguish Worm`},
	},
	"PalOnly": {
		{ID: 179, Key: "PalOnly", Value: `(Paladin Only)`},
	},
	"Paleocene": {
		{ID: 1647, Key: "Paleocene", Value: `Faithful`},
	},
	"Palo Grande": {
		{ID: 2739, Key: "Palo Grande", Value: `Palo Grande`},
	},
	"Paradox": {
		{ID: 2197, Key: "Paradox", Value: `Paradox`},
	},
	"Parkersor's Calm": {
		{ID: 512, Key: "Parkersor's Calm", Value: `Parkersor's Calm`},
	},
	"Pearl": {
		{ID: 237, Key: "Pearl", Value: `Pearl`},
	},
	"Peasent Crown": {
		{ID: 2535, Key: "Peasent Crown", Value: `Peasant Crown`},
	},
	"Pernicious": {
		{ID: 2438, Key: "Pernicious", Value: `Pernicious`},
	},
	"Perpetual": {
		{ID: 1888, Key: "Perpetual", Value: `Perpetual`},
	},
	"Personalizeui": {
		{ID: 1253, Key: "Personalizeui", Value: `personalize`},
	},
	"Pestilent": {
		{ID: 2578, Key: "Pestilent", Value: `Pestilent`},
	},
	"Phoenix Egg": {
		{ID: 1314, Key: "Phoenix Egg", Value: `Phoenix Egg`},
	},
	"Pierre Tombale Couant": {
		{ID: 1679, Key: "Pierre Tombale Couant", Value: `Pierre Tombale Couant`},
	},
	"Pindleskin": {
		{ID: 1060, Key: "Pindleskin", Value: `Pindleskin`},
	},
	"Pitblood Thirst": {
		{ID: 177, Key: "Pitblood Thirst", Value: `Pitblood Thirst`},
	},
	"Plague Bearer": {
		{ID: 433, Key: "Plague Bearer", Value: `Plague Bearer`},
	},
	"Playersubtitles29": {
		{ID: 413, Key: "Playersubtitles29", Value: `Retreat!`},
	},
	"Playersubtitles30": {
		{ID: 355, Key: "Playersubtitles30", Value: `Run away!`},
	},
	"Pompe's Wrath": {
		{ID: 142, Key: "Pompe's Wrath", Value: `Pompeii's Wrath`},
	},
	"Possessed": {
		{ID: 253, Key: "Possessed", Value: `Possessed`},
	},
	"Powered": {
		{ID: 1736, Key: "Powered", Value: `Powered`},
	},
	"Precocious": {
		{ID: 2762, Key: "Precocious", Value: `Precocious`},
	},
	"Preserver's": {
		{ID: 1840, Key: "Preserver's", Value: `Preserver's`},
	},
	"Prison Door": {
		{ID: 423, Key: "Prison Door", Value: `Prison Door`},
	},
	"ProwlingDead": {
		{ID: 75, Key: "ProwlingDead", Value: `Prowling Dead`},
	},
	"Psychic": {
		{ID: 408, Key: "Psychic", Value: `Psychic`},
	},
	"Punishing": {
		{ID: 1328, Key: "Punishing", Value: `Punishing`},
	},
	"Pure": {
		{ID: 1553, Key: "Pure", Value: `Pure`},
	},
	"Pus Spiter": {
		{ID: 2318, Key: "Pus Spiter", Value: `Pus Spitter`},
	},
	"Putrid Defiler1": {
		{ID: 230, Key: "Putrid Defiler1", Value: `Putrid Defiler`},
	},
	"Putrid Defiler2": {
		{ID: 231, Key: "Putrid Defiler2", Value: `Wretched Defiler`},
	},
	"Putrid Defiler3": {
		{ID: 232, Key: "Putrid Defiler3", Value: `Fetid Defiler`},
	},
	"Putrid Defiler4": {
		{ID: 233, Key: "Putrid Defiler4", Value: `Rancid Defiler`},
	},
	"Putrid Defiler5": {
		{ID: 234, Key: "Putrid Defiler5", Value: `Rank Defiler`},
	},
	"Qual'Kek's Enforcer": {
		{ID: 379, Key: "Qual'Kek's Enforcer", Value: `Qual-Khek's Enforcer`},
	},
	"Qual-Kehk": {
		{ID: 263, Key: "Qual-Kehk", Value: `Qual-Kehk`},
	},
	"QualKehkAct5IntroDruGossip1": {
		{ID: 1369, Key: "QualKehkAct5IntroDruGossip1", Value: `42
A Druid in Harrogath! Have things truly 
come to this?
 
After the Mage Wars, I assumed Druids 
would never be seen in Harrogath 
again. You take a big chance coming 
here! 
 
To be honest, I have never been 
comfortable with your shape-shifting 
kind, but I do respect your search for 
balance and peace. So, if you trust us 
enough to enter our gates, I trust you 
enough to let you stay.
 
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`},
	},
	"QualKehkAct5IntroGossip1": {
		{ID: 936, Key: "QualKehkAct5IntroGossip1", Value: `45
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`},
	},
	"QualKehkAct5IntroPalGossip1": {
		{ID: 161, Key: "QualKehkAct5IntroPalGossip1", Value: `42
A Paladin! I have long heard of your 
people.
 
As a young warrior I even considered 
the pilgrimage to Kurast. But I was 
younger then and foolish. My place has 
always been here -- protecting 
Harrogath, and Mount Arreat with it.
 
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`},
	},
	"QualKehkGossip1": {
		{ID: 1621, Key: "QualKehkGossip1", Value: `67
It would be an honor to have a warrior 
of the Light fighting side-by-side with 
my men.
 
I can see your faith gives you great 
strength, Paladin, but don't expect it to 
keep you out of harm's way.
`},
	},
	"QualKehkGossip2": {
		{ID: 1622, Key: "QualKehkGossip2", Value: `61
Harrogath has great need of your 
powers, noble Druid. However, in the 
face of this supernatural onslaught, 
are your natural powers up to the 
task?
`},
	},
	"QualKehkGossip3": {
		{ID: 1623, Key: "QualKehkGossip3", Value: `44
The death of Malah's son was a great 
tragedy. He was our finest archer.
 
While leading a successful campaign 
against Baal's forces, he was impaled 
on a demon's spear. The wound was 
such that...Well, even Malah herself 
acknowledges that quick death was a 
blessing.
`},
	},
	"QualKehkGossip4": {
		{ID: 1624, Key: "QualKehkGossip4", Value: `57
We have lost many well-trained warriors 
to Baal's siege machines. Their range is 
great. Though, they are vulnerable if 
you close the distance quickly enough.
`},
	},
	"QualKehkGossip5": {
		{ID: 1625, Key: "QualKehkGossip5", Value: `96
Baal's legions seem countless, but 
slaying their commanders takes some 
of the fight out of them.
`},
	},
	"QualKehkGossip6": {
		{ID: 1626, Key: "QualKehkGossip6", Value: `58
Early on, parties of our best scouts 
were ambushed by demons that 
spawned from the very air around 
them. Survivors often mentioned a 
strange creature floating in the 
distance.
 
Perhaps taking it down could prevent 
some nasty surprises.
`},
	},
	"QualKehkGossip7": {
		{ID: 1627, Key: "QualKehkGossip7", Value: `57
This is unlike any battle I have ever 
fought. While we ration food and 
water, the demon hordes feast nightly 
on the flesh and blood of our dead.
`},
	},
	"QualKehkGossip8": {
		{ID: 1628, Key: "QualKehkGossip8", Value: `48
Larzuk is a talented blacksmith, but his 
head is full of some strange ideas. 
 
Just the other day he came to me with a 
plan to break the siege. He wanted to 
fill large pipes with exploding powders 
and steel balls, then...Well, like I said, 
strange.
`},
	},
	"QualKehkGossip9": {
		{ID: 1629, Key: "QualKehkGossip9", Value: `46
Our Elders were wise leaders in more 
peaceful times, but now the survival of 
our people has fallen to me. My men 
and I will fight to the death, but there's 
no way to ensure the outcome.
 
I used to believe that nothing could 
break through our guard and assault 
the holy mountain. I know now that I 
was horribly mistaken.
`},
	},
	"Que-Hegan's Wisdon": {
		{ID: 757, Key: "Que-Hegan's Wisdon", Value: `Que-Hegan's Wisdom`},
	},
	"Que-hegan's Wisdom": {
		{ID: 1268, Key: "Que-hegan's Wisdom", Value: `Que-Hegan's Wisdom`},
	},
	"Quixotic": {
		{ID: 1619, Key: "Quixotic", Value: `Quixotic`},
	},
	"Rabbit Slayer": {
		{ID: 291, Key: "Rabbit Slayer", Value: `Rabbit Slayer`},
	},
	"Radimant's Sphere": {
		{ID: 2787, Key: "Radimant's Sphere", Value: `Radament's Sphere`},
	},
	"Raging": {
		{ID: 652, Key: "Raging", Value: `Raging`},
	},
	"Raiden's Crutch": {
		{ID: 2052, Key: "Raiden's Crutch", Value: `Raiden's Crutch`},
	},
	"Rainbow": {
		{ID: 594, Key: "Rainbow", Value: `Rainbow`},
	},
	"Rainbow Facet": {
		{ID: 1443, Key: "Rainbow Facet", Value: `Rainbow Facet`},
	},
	"Raven Frost": {
		{ID: 1723, Key: "Raven Frost", Value: `Raven Frost`},
	},
	"Ravenlore": {
		{ID: 1446, Key: "Ravenlore", Value: `Ravenlore`},
	},
	"Razoredge": {
		{ID: 2294, Key: "Razoredge", Value: `Razor's Edge`},
	},
	"Razorswitch": {
		{ID: 1310, Key: "Razorswitch", Value: `Razorswitch`},
	},
	"Razortail": {
		{ID: 2001, Key: "Razortail", Value: `Razortail`},
	},
	"Realgar": {
		{ID: 1717, Key: "Realgar", Value: `Realgar`},
	},
	"ReanimatedHorde": {
		{ID: 73, Key: "ReanimatedHorde", Value: `Reanimated Horde`},
	},
	"Reclusive": {
		{ID: 885, Key: "Reclusive", Value: `Reclusive`},
	},
	"Rename Instruct": {
		{ID: 836, Key: "Rename Instruct", Value: `Choose the item that you wish to personalize with your name.`},
	},
	"Repulsive": {
		{ID: 887, Key: "Repulsive", Value: `Repulsive`},
	},
	"RescueCainAss": {
		{ID: 2216, Key: "RescueCainAss", Value: `Cain! Go to the Rogue camp.`},
	},
	"RescueCainDru": {
		{ID: 1459, Key: "RescueCainDru", Value: `Deckard Cain, leave this place!`},
	},
	"RescueQual-KehkAct5Ams": {
		{ID: 1276, Key: "RescueQual-KehkAct5Ams", Value: `Follow me.`},
	},
	"RescueQual-KehkAct5Ass": {
		{ID: 609, Key: "RescueQual-KehkAct5Ass", Value: `Follow me.`},
	},
	"RescueQual-KehkAct5Bar": {
		{ID: 574, Key: "RescueQual-KehkAct5Bar", Value: `Follow me.`},
	},
	"RescueQual-KehkAct5Dru": {
		{ID: 817, Key: "RescueQual-KehkAct5Dru", Value: `Follow me.`},
	},
	"RescueQual-KehkAct5Nec": {
		{ID: 2211, Key: "RescueQual-KehkAct5Nec", Value: `Follow me.`},
	},
	"RescueQual-KehkAct5Pal": {
		{ID: 24, Key: "RescueQual-KehkAct5Pal", Value: `Follow me.`},
	},
	"RescueQual-KehkAct5Sor": {
		{ID: 2209, Key: "RescueQual-KehkAct5Sor", Value: `Follow me.`},
	},
	"Resonant": {
		{ID: 156, Key: "Resonant", Value: `Resonant`},
	},
	"Ribcracker": {
		{ID: 604, Key: "Ribcracker", Value: `Ribcracker`},
	},
	"Riftlash": {
		{ID: 1599, Key: "Riftlash", Value: `Riftlash`},
	},
	"Riftslash": {
		{ID: 1523, Key: "Riftslash", Value: `Riftslash`},
	},
	"Rigid Highlands": {
		{ID: 1006, Key: "Rigid Highlands", Value: `Frigid Highlands`},
	},
	"Rings": {
		{ID: 1452, Key: "Rings", Value: `Rings`},
	},
	"Riphook": {
		{ID: 373, Key: "Riphook", Value: `Riphook`},
	},
	"Rite of Passage": {
		{ID: 387, Key: "Rite of Passage", Value: `Rite of Passage`},
	},
	"Robin's Yolk": {
		{ID: 1680, Key: "Robin's Yolk", Value: `Robin's Yolk`},
	},
	"Robineye": {
		{ID: 2144, Key: "Robineye", Value: `Robineye`},
	},
	"Rockhew": {
		{ID: 1114, Key: "Rockhew", Value: `Rockhew`},
	},
	"Rockstopper": {
		{ID: 149, Key: "Rockstopper", Value: `Rockstopper`},
	},
	"Rocky Summit": {
		{ID: 1163, Key: "Rocky Summit", Value: `Arreat Summit`},
	},
	"Rose": {
		{ID: 2580, Key: "Rose", Value: `Rose`},
	},
	"Rose Branded": {
		{ID: 997, Key: "Rose Branded", Value: `Rose Branded`},
	},
	"RotWalker": {
		{ID: 756, Key: "RotWalker", Value: `Rot Walker`},
	},
	"Rotting": {
		{ID: 2321, Key: "Rotting", Value: `Rotting`},
	},
	"Royal": {
		{ID: 544, Key: "Royal", Value: `Royal`},
	},
	"Rude": {
		{ID: 1056, Key: "Rude", Value: `Rude`},
	},
	"RuneQuote": {
		{ID: 2147, Key: "RuneQuote", Value: `'`},
	},
	"Runemaster": {
		{ID: 2283, Key: "Runemaster", Value: `Rune Master`},
	},
	"Runeslayer": {
		{ID: 2581, Key: "Runeslayer", Value: `Rune Slayer`},
	},
	"Runeword1": {
		{ID: 1421, Key: "Runeword1", Value: `Ancients' Pledge`},
	},
	"Runeword10": {
		{ID: 1333, Key: "Runeword10", Value: `Brand`},
	},
	"Runeword100": {
		{ID: 2559, Key: "Runeword100", Value: `Penitence`},
	},
	"Runeword101": {
		{ID: 2560, Key: "Runeword101", Value: `Peril`},
	},
	"Runeword102": {
		{ID: 2561, Key: "Runeword102", Value: `Pestilence`},
	},
	"Runeword103": {
		{ID: 2562, Key: "Runeword103", Value: `Phoenix`},
	},
	"Runeword104": {
		{ID: 2563, Key: "Runeword104", Value: `Piety`},
	},
	"Runeword105": {
		{ID: 2564, Key: "Runeword105", Value: `Pillar of Faith`},
	},
	"Runeword106": {
		{ID: 2565, Key: "Runeword106", Value: `Plague`},
	},
	"Runeword107": {
		{ID: 2566, Key: "Runeword107", Value: `Praise`},
	},
	"Runeword108": {
		{ID: 2567, Key: "Runeword108", Value: `Prayer`},
	},
	"Runeword109": {
		{ID: 2568, Key: "Runeword109", Value: `Pride`},
	},
	"Runeword11": {
		{ID: 1334, Key: "Runeword11", Value: `Breath of the Dying`},
	},
	"Runeword110": {
		{ID: 2447, Key: "Runeword110", Value: `Principle`},
	},
	"Runeword111": {
		{ID: 2448, Key: "Runeword111", Value: `Prowess in Battle`},
	},
	"Runeword112": {
		{ID: 2449, Key: "Runeword112", Value: `Prudence`},
	},
	"Runeword113": {
		{ID: 2450, Key: "Runeword113", Value: `Punishment`},
	},
	"Runeword114": {
		{ID: 2451, Key: "Runeword114", Value: `Purity`},
	},
	"Runeword115": {
		{ID: 2453, Key: "Runeword115", Value: `Question`},
	},
	"Runeword116": {
		{ID: 2454, Key: "Runeword116", Value: `Radiance`},
	},
	"Runeword117": {
		{ID: 2455, Key: "Runeword117", Value: `Rain`},
	},
	"Runeword118": {
		{ID: 2456, Key: "Runeword118", Value: `Reason`},
	},
	"Runeword119": {
		{ID: 2457, Key: "Runeword119", Value: `Red`},
	},
	"Runeword12": {
		{ID: 1335, Key: "Runeword12", Value: `Broken Promise`},
	},
	"Runeword120": {
		{ID: 2463, Key: "Runeword120", Value: `Rhyme`},
	},
	"Runeword121": {
		{ID: 2464, Key: "Runeword121", Value: `Rift`},
	},
	"Runeword122": {
		{ID: 2465, Key: "Runeword122", Value: `Sanctuary`},
	},
	"Runeword123": {
		{ID: 2466, Key: "Runeword123", Value: `Serendipity`},
	},
	"Runeword124": {
		{ID: 2467, Key: "Runeword124", Value: `Shadow`},
	},
	"Runeword125": {
		{ID: 2468, Key: "Runeword125", Value: `Shadow of Doubt`},
	},
	"Runeword126": {
		{ID: 2469, Key: "Runeword126", Value: `Silence`},
	},
	"Runeword127": {
		{ID: 2470, Key: "Runeword127", Value: `Siren's Song`},
	},
	"Runeword128": {
		{ID: 2471, Key: "Runeword128", Value: `Smoke`},
	},
	"Runeword129": {
		{ID: 2473, Key: "Runeword129", Value: `Sorrow`},
	},
	"Runeword13": {
		{ID: 1336, Key: "Runeword13", Value: `Call to Arms`},
	},
	"Runeword130": {
		{ID: 2479, Key: "Runeword130", Value: `Spirit`},
	},
	"Runeword131": {
		{ID: 2480, Key: "Runeword131", Value: `Splendor`},
	},
	"Runeword132": {
		{ID: 2481, Key: "Runeword132", Value: `Starlight`},
	},
	"Runeword133": {
		{ID: 2482, Key: "Runeword133", Value: `Stealth`},
	},
	"Runeword134": {
		{ID: 2483, Key: "Runeword134", Value: `Steel`},
	},
	"Runeword135": {
		{ID: 2484, Key: "Runeword135", Value: `Still Water`},
	},
	"Runeword136": {
		{ID: 2485, Key: "Runeword136", Value: `Sting`},
	},
	"Runeword137": {
		{ID: 2486, Key: "Runeword137", Value: `Stone`},
	},
	"Runeword138": {
		{ID: 2487, Key: "Runeword138", Value: `Storm`},
	},
	"Runeword139": {
		{ID: 2488, Key: "Runeword139", Value: `Strength`},
	},
	"Runeword14": {
		{ID: 1337, Key: "Runeword14", Value: `Chains of Honor`},
	},
	"Runeword140": {
		{ID: 2495, Key: "Runeword140", Value: `Tempest`},
	},
	"Runeword141": {
		{ID: 2496, Key: "Runeword141", Value: `Temptation`},
	},
	"Runeword142": {
		{ID: 2497, Key: "Runeword142", Value: `Terror`},
	},
	"Runeword143": {
		{ID: 2498, Key: "Runeword143", Value: `Thirst`},
	},
	"Runeword144": {
		{ID: 2500, Key: "Runeword144", Value: `Thought`},
	},
	"Runeword145": {
		{ID: 2501, Key: "Runeword145", Value: `Thunder`},
	},
	"Runeword146": {
		{ID: 2502, Key: "Runeword146", Value: `Time`},
	},
	"Runeword147": {
		{ID: 2503, Key: "Runeword147", Value: `Tradition`},
	},
	"Runeword148": {
		{ID: 2504, Key: "Runeword148", Value: `Treachery`},
	},
	"Runeword149": {
		{ID: 2506, Key: "Runeword149", Value: `Trust`},
	},
	"Runeword15": {
		{ID: 1338, Key: "Runeword15", Value: `Chance`},
	},
	"Runeword150": {
		{ID: 2383, Key: "Runeword150", Value: `Truth`},
	},
	"Runeword151": {
		{ID: 2384, Key: "Runeword151", Value: `Unbending Will`},
	},
	"Runeword152": {
		{ID: 2385, Key: "Runeword152", Value: `Valor`},
	},
	"Runeword153": {
		{ID: 2386, Key: "Runeword153", Value: `Vengeance`},
	},
	"Runeword154": {
		{ID: 2387, Key: "Runeword154", Value: `Venom`},
	},
	"Runeword155": {
		{ID: 2388, Key: "Runeword155", Value: `Victory`},
	},
	"Runeword156": {
		{ID: 2389, Key: "Runeword156", Value: `Voice`},
	},
	"Runeword157": {
		{ID: 2390, Key: "Runeword157", Value: `Void`},
	},
	"Runeword158": {
		{ID: 2391, Key: "Runeword158", Value: `War`},
	},
	"Runeword159": {
		{ID: 2392, Key: "Runeword159", Value: `Water`},
	},
	"Runeword16": {
		{ID: 1339, Key: "Runeword16", Value: `Chaos`},
	},
	"Runeword160": {
		{ID: 2399, Key: "Runeword160", Value: `Wealth`},
	},
	"Runeword161": {
		{ID: 2400, Key: "Runeword161", Value: `Whisper`},
	},
	"Runeword162": {
		{ID: 2401, Key: "Runeword162", Value: `White`},
	},
	"Runeword163": {
		{ID: 2402, Key: "Runeword163", Value: `Wind`},
	},
	"Runeword164": {
		{ID: 2403, Key: "Runeword164", Value: `Wings of Hope`},
	},
	"Runeword165": {
		{ID: 2404, Key: "Runeword165", Value: `Wisdom`},
	},
	"Runeword166": {
		{ID: 2406, Key: "Runeword166", Value: `Woe`},
	},
	"Runeword167": {
		{ID: 2408, Key: "Runeword167", Value: `Wonder`},
	},
	"Runeword168": {
		{ID: 2409, Key: "Runeword168", Value: `Wrath`},
	},
	"Runeword169": {
		{ID: 2410, Key: "Runeword169", Value: `Youth`},
	},
	"Runeword17": {
		{ID: 1340, Key: "Runeword17", Value: `Crescent Moon`},
	},
	"Runeword170": {
		{ID: 2415, Key: "Runeword170", Value: `Zephyr`},
	},
	"Runeword18": {
		{ID: 1341, Key: "Runeword18", Value: `Darkness`},
	},
	"Runeword19": {
		{ID: 1342, Key: "Runeword19", Value: `Daylight`},
	},
	"Runeword2": {
		{ID: 1423, Key: "Runeword2", Value: `Armageddon`},
	},
	"Runeword20": {
		{ID: 1349, Key: "Runeword20", Value: `Death`},
	},
	"Runeword21": {
		{ID: 1359, Key: "Runeword21", Value: `Deception`},
	},
	"Runeword22": {
		{ID: 1360, Key: "Runeword22", Value: `Delerium`},
	},
	"Runeword23": {
		{ID: 1361, Key: "Runeword23", Value: `Desire`},
	},
	"Runeword24": {
		{ID: 1362, Key: "Runeword24", Value: `Despair`},
	},
	"Runeword25": {
		{ID: 1363, Key: "Runeword25", Value: `Destruction`},
	},
	"Runeword26": {
		{ID: 1364, Key: "Runeword26", Value: `Doom`},
	},
	"Runeword27": {
		{ID: 1365, Key: "Runeword27", Value: `Dragon`},
	},
	"Runeword28": {
		{ID: 1373, Key: "Runeword28", Value: `Dread`},
	},
	"Runeword29": {
		{ID: 1375, Key: "Runeword29", Value: `Dream`},
	},
	"Runeword3": {
		{ID: 1424, Key: "Runeword3", Value: `Authority`},
	},
	"Runeword30": {
		{ID: 1376, Key: "Runeword30", Value: `Duress`},
	},
	"Runeword31": {
		{ID: 1378, Key: "Runeword31", Value: `Edge`},
	},
	"Runeword32": {
		{ID: 1379, Key: "Runeword32", Value: `Elation`},
	},
	"Runeword33": {
		{ID: 1380, Key: "Runeword33", Value: `Enigma`},
	},
	"Runeword34": {
		{ID: 1381, Key: "Runeword34", Value: `Enlightenment`},
	},
	"Runeword35": {
		{ID: 1382, Key: "Runeword35", Value: `Envy`},
	},
	"Runeword36": {
		{ID: 1383, Key: "Runeword36", Value: `Eternity`},
	},
	"Runeword37": {
		{ID: 1384, Key: "Runeword37", Value: `Exile`},
	},
	"Runeword38": {
		{ID: 1385, Key: "Runeword38", Value: `Faith`},
	},
	"Runeword39": {
		{ID: 1386, Key: "Runeword39", Value: `Famine`},
	},
	"Runeword4": {
		{ID: 1425, Key: "Runeword4", Value: `Beast`},
	},
	"Runeword40": {
		{ID: 1387, Key: "Runeword40", Value: `Flame`},
	},
	"Runeword41": {
		{ID: 1388, Key: "Runeword41", Value: `Fortitude`},
	},
	"Runeword42": {
		{ID: 1389, Key: "Runeword42", Value: `Fortune`},
	},
	"Runeword43": {
		{ID: 1391, Key: "Runeword43", Value: `Friendship`},
	},
	"Runeword44": {
		{ID: 1392, Key: "Runeword44", Value: `Fury`},
	},
	"Runeword45": {
		{ID: 1393, Key: "Runeword45", Value: `Gloom`},
	},
	"Runeword46": {
		{ID: 1394, Key: "Runeword46", Value: `Glory`},
	},
	"Runeword47": {
		{ID: 1395, Key: "Runeword47", Value: `Grief`},
	},
	"Runeword48": {
		{ID: 1396, Key: "Runeword48", Value: `Hand of Justice`},
	},
	"Runeword49": {
		{ID: 1397, Key: "Runeword49", Value: `Harmony`},
	},
	"Runeword5": {
		{ID: 1426, Key: "Runeword5", Value: `Beauty`},
	},
	"Runeword50": {
		{ID: 1525, Key: "Runeword50", Value: `Hatred`},
	},
	"Runeword51": {
		{ID: 1526, Key: "Runeword51", Value: `Heart of the Oak`},
	},
	"Runeword52": {
		{ID: 1527, Key: "Runeword52", Value: `Heaven's Will`},
	},
	"Runeword53": {
		{ID: 1529, Key: "Runeword53", Value: `Holy Tears`},
	},
	"Runeword54": {
		{ID: 1530, Key: "Runeword54", Value: `Holy Thunder`},
	},
	"Runeword55": {
		{ID: 1531, Key: "Runeword55", Value: `Honor`},
	},
	"Runeword56": {
		{ID: 1532, Key: "Runeword56", Value: `Revenge`},
	},
	"Runeword57": {
		{ID: 1533, Key: "Runeword57", Value: `Humility`},
	},
	"Runeword58": {
		{ID: 1535, Key: "Runeword58", Value: `Hunger`},
	},
	"Runeword59": {
		{ID: 1536, Key: "Runeword59", Value: `Ice`},
	},
	"Runeword6": {
		{ID: 1427, Key: "Runeword6", Value: `Black`},
	},
	"Runeword60": {
		{ID: 1541, Key: "Runeword60", Value: `Infinity`},
	},
	"Runeword61": {
		{ID: 1542, Key: "Runeword61", Value: `Innocence`},
	},
	"Runeword62": {
		{ID: 1543, Key: "Runeword62", Value: `Insight`},
	},
	"Runeword63": {
		{ID: 1544, Key: "Runeword63", Value: `Jealousy`},
	},
	"Runeword64": {
		{ID: 1545, Key: "Runeword64", Value: `Judgement`},
	},
	"Runeword65": {
		{ID: 1546, Key: "Runeword65", Value: `King's Grace`},
	},
	"Runeword66": {
		{ID: 1547, Key: "Runeword66", Value: `Kingslayer`},
	},
	"Runeword67": {
		{ID: 1548, Key: "Runeword67", Value: `Knight's Vigil`},
	},
	"Runeword68": {
		{ID: 1549, Key: "Runeword68", Value: `Knowledge`},
	},
	"Runeword69": {
		{ID: 1550, Key: "Runeword69", Value: `Last Wish`},
	},
	"Runeword7": {
		{ID: 1429, Key: "Runeword7", Value: `Blood`},
	},
	"Runeword70": {
		{ID: 1557, Key: "Runeword70", Value: `Law`},
	},
	"Runeword71": {
		{ID: 1558, Key: "Runeword71", Value: `Lawbringer`},
	},
	"Runeword72": {
		{ID: 1559, Key: "Runeword72", Value: `Leaf`},
	},
	"Runeword73": {
		{ID: 1560, Key: "Runeword73", Value: `Lightning`},
	},
	"Runeword74": {
		{ID: 1561, Key: "Runeword74", Value: `Lionheart`},
	},
	"Runeword75": {
		{ID: 1562, Key: "Runeword75", Value: `Lore`},
	},
	"Runeword76": {
		{ID: 1563, Key: "Runeword76", Value: `Love`},
	},
	"Runeword77": {
		{ID: 1564, Key: "Runeword77", Value: `Loyalty`},
	},
	"Runeword78": {
		{ID: 1565, Key: "Runeword78", Value: `Lust`},
	},
	"Runeword79": {
		{ID: 1566, Key: "Runeword79", Value: `Madness`},
	},
	"Runeword8": {
		{ID: 1430, Key: "Runeword8", Value: `Bone`},
	},
	"Runeword81": {
		{ID: 1574, Key: "Runeword81", Value: `Malice`},
	},
	"Runeword82": {
		{ID: 1575, Key: "Runeword82", Value: `Melody`},
	},
	"Runeword83": {
		{ID: 1578, Key: "Runeword83", Value: `Memory`},
	},
	"Runeword84": {
		{ID: 1579, Key: "Runeword84", Value: `Mist`},
	},
	"Runeword85": {
		{ID: 1580, Key: "Runeword85", Value: `Morning`},
	},
	"Runeword86": {
		{ID: 1581, Key: "Runeword86", Value: `Mystery`},
	},
	"Runeword87": {
		{ID: 1582, Key: "Runeword87", Value: `Myth`},
	},
	"Runeword88": {
		{ID: 1583, Key: "Runeword88", Value: `Nadir`},
	},
	"Runeword89": {
		{ID: 1584, Key: "Runeword89", Value: `Nature's Kingdom`},
	},
	"Runeword9": {
		{ID: 1431, Key: "Runeword9", Value: `Bramble`},
	},
	"Runeword90": {
		{ID: 1461, Key: "Runeword90", Value: `Night`},
	},
	"Runeword91": {
		{ID: 1462, Key: "Runeword91", Value: `Oath`},
	},
	"Runeword92": {
		{ID: 1463, Key: "Runeword92", Value: `Obedience`},
	},
	"Runeword93": {
		{ID: 1464, Key: "Runeword93", Value: `Oblivion`},
	},
	"Runeword94": {
		{ID: 1465, Key: "Runeword94", Value: `Obsession`},
	},
	"Runeword95": {
		{ID: 1466, Key: "Runeword95", Value: `Passion`},
	},
	"Runeword96": {
		{ID: 1467, Key: "Runeword96", Value: `Patience`},
	},
	"Runeword97": {
		{ID: 1468, Key: "Runeword97", Value: `Patter`},
	},
	"Runeword98": {
		{ID: 1469, Key: "Runeword98", Value: `Peace`},
	},
	"Runeword99": {
		{ID: 1471, Key: "Runeword99", Value: `Voice of Reason`},
	},
	"Rusty": {
		{ID: 1400, Key: "Rusty", Value: `Rusty`},
	},
	"Sacred": {
		{ID: 428, Key: "Sacred", Value: `Sacred`},
	},
	"Sacred Charge": {
		{ID: 284, Key: "Sacred Charge", Value: `Sacred Charge`},
	},
	"Samual's Caretaker": {
		{ID: 2042, Key: "Samual's Caretaker", Value: `Samuel's Caretaker`},
	},
	"Sander's Basis": {
		{ID: 1066, Key: "Sander's Basis", Value: `McAuley's Basis`},
	},
	"Sander's Court Jester": {
		{ID: 908, Key: "Sander's Court Jester", Value: `McAuley's Folly`},
	},
	"Sander's Derby": {
		{ID: 318, Key: "Sander's Derby", Value: `McAuley's Pledge`},
	},
	"Sander's Superstition": {
		{ID: 709, Key: "Sander's Superstition", Value: `McAuley's Superstition`},
	},
	"Sander's Taboo": {
		{ID: 246, Key: "Sander's Taboo", Value: `McAuley's Taboo`},
	},
	"Sandstorm Trek": {
		{ID: 2440, Key: "Sandstorm Trek", Value: `Sandstorm Trek`},
	},
	"Sanguinary": {
		{ID: 69, Key: "Sanguinary", Value: `Sanguinary`},
	},
	"Sanguine": {
		{ID: 624, Key: "Sanguine", Value: `Sanguine`},
	},
	"Sankenkur's Resurrection": {
		{ID: 197, Key: "Sankenkur's Resurrection", Value: `Sankekur's Resurrection`},
	},
	"Saracen's Chance": {
		{ID: 1931, Key: "Saracen's Chance", Value: `Saracen's Chance`},
	},
	"Sarmichian Justice": {
		{ID: 1301, Key: "Sarmichian Justice", Value: `Samurai Justice`},
	},
	"Sazabi's Cobalt Redeemer": {
		{ID: 689, Key: "Sazabi's Cobalt Redeemer", Value: `Sazabi's Cobalt Redeemer`},
	},
	"Sazabi's Ghost Liberator": {
		{ID: 2517, Key: "Sazabi's Ghost Liberator", Value: `Sazabi's Ghost Liberator`},
	},
	"Sazabi's Grand Tribute": {
		{ID: 208, Key: "Sazabi's Grand Tribute", Value: `Sazabi's Grand Tribute`},
	},
	"Sazabi's Mental Sheath": {
		{ID: 825, Key: "Sazabi's Mental Sheath", Value: `Sazabi's Mental Sheath`},
	},
	"Scarlet": {
		{ID: 851, Key: "Scarlet", Value: `Scarlet`},
	},
	"Schaefer's Hammer": {
		{ID: 2347, Key: "Schaefer's Hammer", Value: `Schaefer's Hammer`},
	},
	"Scintillating": {
		{ID: 2492, Key: "Scintillating", Value: `Scintillating`},
	},
	"Scorched": {
		{ID: 1868, Key: "Scorched", Value: `Scorched`},
	},
	"Scorching": {
		{ID: 1517, Key: "Scorching", Value: `Condensing`},
	},
	"Screaming": {
		{ID: 2489, Key: "Screaming", Value: `Screaming`},
	},
	"Seige Tower": {
		{ID: 1169, Key: "Seige Tower", Value: `Seige Tower`},
	},
	"Sensei's": {
		{ID: 276, Key: "Sensei's", Value: `Sensei's`},
	},
	"Septic": {
		{ID: 667, Key: "Septic", Value: `Septic`},
	},
	"Seraph's Hymn": {
		{ID: 506, Key: "Seraph's Hymn", Value: `Seraph's Hymn`},
	},
	"Serrated": {
		{ID: 1193, Key: "Serrated", Value: `Serrated`},
	},
	"Shadefalcon": {
		{ID: 1607, Key: "Shadefalcon", Value: `Shade Falcon`},
	},
	"Shadow Master": {
		{ID: 245, Key: "Shadow Master", Value: `Shadow Master`},
	},
	"Shadow Warrior": {
		{ID: 572, Key: "Shadow Warrior", Value: `Shadow Warrior`},
	},
	"Shadow's Touch": {
		{ID: 1157, Key: "Shadow's Touch", Value: `Shadow Touch`},
	},
	"Shadowdancer": {
		{ID: 2278, Key: "Shadowdancer", Value: `Shadow Dancer`},
	},
	"Shadowkiller": {
		{ID: 362, Key: "Shadowkiller", Value: `Shadow Killer`},
	},
	"Shaftstop": {
		{ID: 2413, Key: "Shaftstop", Value: `Shaftstop`},
	},
	"Shakabra's Crux": {
		{ID: 2632, Key: "Shakabra's Crux", Value: `Shakabra's Crux`},
	},
	"Shaman's": {
		{ID: 822, Key: "Shaman's", Value: `Shaman's`},
	},
	"Sharp Tooth Sayer": {
		{ID: 283, Key: "Sharp Tooth Sayer", Value: `Sharptooth Slayer`},
	},
	"Shimmering": {
		{ID: 1731, Key: "Shimmering", Value: `Shimmering`},
	},
	"Shivering": {
		{ID: 490, Key: "Shivering", Value: `Shivering`},
	},
	"Shocking": {
		{ID: 35, Key: "Shocking", Value: `Shocking`},
	},
	"Shogukusha's": {
		{ID: 2393, Key: "Shogukusha's", Value: `Shogukusha's`},
	},
	"Shouting": {
		{ID: 1567, Key: "Shouting", Value: `Howling`},
	},
	"Shrine2wilderness": {
		{ID: 391, Key: "Shrine2wilderness", Value: `Shrine`},
	},
	"Shrine3wilderness": {
		{ID: 1117, Key: "Shrine3wilderness", Value: `Shrine`},
	},
	"Siege Beast": {
		{ID: 78, Key: "Siege Beast", Value: `Siege Beast`},
	},
	"Siege Boss": {
		{ID: 229, Key: "Siege Boss", Value: `Shenk the Overseer`},
	},
	"Siege Control": {
		{ID: 393, Key: "Siege Control", Value: `Siege Control`},
	},
	"Siege Door": {
		{ID: 330, Key: "Siege Door", Value: `Siege Door`},
	},
	"Sightless Veil": {
		{ID: 2620, Key: "Sightless Veil", Value: `Sightless Veil`},
	},
	"Sigurd's Staunch": {
		{ID: 1985, Key: "Sigurd's Staunch", Value: `Siggard's Stealth`},
	},
	"Silkweave": {
		{ID: 524, Key: "Silkweave", Value: `Silkweave`},
	},
	"Sinblade": {
		{ID: 538, Key: "Sinblade", Value: `Sinblade`},
	},
	"Singing": {
		{ID: 340, Key: "Singing", Value: `Singing`},
	},
	"Siren's call": {
		{ID: 2291, Key: "Siren's call", Value: `Siren's Call`},
	},
	"Skillan222": {
		{ID: 2214, Key: "Skillan222", Value: `Raven`},
	},
	"Skillan223": {
		{ID: 2215, Key: "Skillan223", Value: `Psn Creep`},
		{ID: 2218, Key: "Skillan223", Value: `Werewolf`},
	},
	"Skillan225": {
		{ID: 2227, Key: "Skillan225", Value: `Lycanthropy`},
	},
	"Skillan226": {
		{ID: 2252, Key: "Skillan226", Value: `Firestorm`},
	},
	"Skillan227": {
		{ID: 2253, Key: "Skillan227", Value: `Oak Sage`},
	},
	"Skillan228": {
		{ID: 2254, Key: "Skillan228", Value: `Sum Spt Wolf`},
	},
	"Skillan229": {
		{ID: 2256, Key: "Skillan229", Value: `Werebear`},
	},
	"Skillan230": {
		{ID: 2293, Key: "Skillan230", Value: `Molten Boulder`},
	},
	"Skillan231": {
		{ID: 2295, Key: "Skillan231", Value: `Arctic Blast`},
	},
	"Skillan232": {
		{ID: 2296, Key: "Skillan232", Value: `Carrion Vine`},
	},
	"Skillan233": {
		{ID: 2300, Key: "Skillan233", Value: `Feral Rage`},
	},
	"Skillan234": {
		{ID: 2301, Key: "Skillan234", Value: `Maul`},
	},
	"Skillan235": {
		{ID: 2302, Key: "Skillan235", Value: `Fissure`},
	},
	"Skillan236": {
		{ID: 2303, Key: "Skillan236", Value: `Cyclone Armor`},
	},
	"Skillan237": {
		{ID: 2317, Key: "Skillan237", Value: `Wolverine Hrt`},
	},
	"Skillan238": {
		{ID: 2319, Key: "Skillan238", Value: `Summon D Wolf`},
	},
	"Skillan239": {
		{ID: 2323, Key: "Skillan239", Value: `Rabies`},
	},
	"Skillan240": {
		{ID: 2325, Key: "Skillan240", Value: `Fire Claws`},
	},
	"Skillan241": {
		{ID: 2330, Key: "Skillan241", Value: `Twister`},
	},
	"Skillan242": {
		{ID: 2331, Key: "Skillan242", Value: `Sol Creep`},
	},
	"Skillan243": {
		{ID: 2332, Key: "Skillan243", Value: `Hunger`},
	},
	"Skillan244": {
		{ID: 2334, Key: "Skillan244", Value: `Shock Wave`},
	},
	"Skillan245": {
		{ID: 2335, Key: "Skillan245", Value: `Volcano`},
	},
	"Skillan246": {
		{ID: 2336, Key: "Skillan246", Value: `Tornado`},
	},
	"Skillan247": {
		{ID: 2337, Key: "Skillan247", Value: `Spirit Barbs`},
	},
	"Skillan248": {
		{ID: 2349, Key: "Skillan248", Value: `Summon Grizzly`},
	},
	"Skillan249": {
		{ID: 2358, Key: "Skillan249", Value: `Fury`},
	},
	"Skillan250": {
		{ID: 2366, Key: "Skillan250", Value: `Armageddon`},
	},
	"Skillan251": {
		{ID: 2370, Key: "Skillan251", Value: `Hurricane`},
	},
	"Skillan252": {
		{ID: 2375, Key: "Skillan252", Value: `Fire Blast`},
	},
	"Skillan253": {
		{ID: 2378, Key: "Skillan253", Value: `Claw Mastery`},
	},
	"Skillan254": {
		{ID: 2382, Key: "Skillan254", Value: `Psyc Hammer`},
	},
	"Skillan255": {
		{ID: 2424, Key: "Skillan255", Value: `Tiger Strike`},
	},
	"Skillan256": {
		{ID: 2429, Key: "Skillan256", Value: `Dragon Talon`},
	},
	"Skillan257": {
		{ID: 2637, Key: "Skillan257", Value: `Shock Web`},
	},
	"Skillan258": {
		{ID: 2641, Key: "Skillan258", Value: `Blade Sentinel`},
	},
	"Skillan259": {
		{ID: 2648, Key: "Skillan259", Value: `Burst of Speed`},
	},
	"Skillan260": {
		{ID: 2653, Key: "Skillan260", Value: `Fists of Fire`},
	},
	"Skillan261": {
		{ID: 2657, Key: "Skillan261", Value: `Dragon Claw`},
	},
	"Skillan262": {
		{ID: 2670, Key: "Skillan262", Value: `Charged Bolt`},
	},
	"Skillan263": {
		{ID: 2676, Key: "Skillan263", Value: `Wake of Fire`},
	},
	"Skillan264": {
		{ID: 2684, Key: "Skillan264", Value: `Wpn Block`},
	},
	"Skillan265": {
		{ID: 2688, Key: "Skillan265", Value: `Cloak of Shdws`},
	},
	"Skillan266": {
		{ID: 2694, Key: "Skillan266", Value: `Cobra Strike`},
	},
	"Skillan267": {
		{ID: 2703, Key: "Skillan267", Value: `Blade Fury`},
	},
	"Skillan268": {
		{ID: 2708, Key: "Skillan268", Value: `Fade`},
	},
	"Skillan269": {
		{ID: 2712, Key: "Skillan269", Value: `Shdw Warrior`},
	},
	"Skillan270": {
		{ID: 2717, Key: "Skillan270", Value: `Thunder Claws`},
	},
	"Skillan271": {
		{ID: 2724, Key: "Skillan271", Value: `Dragon Tail`},
	},
	"Skillan272": {
		{ID: 2728, Key: "Skillan272", Value: `Lightning Sentry`},
	},
	"Skillan273": {
		{ID: 2732, Key: "Skillan273", Value: `Wake of Inferno`},
	},
	"Skillan274": {
		{ID: 2751, Key: "Skillan274", Value: `Mind Blast`},
	},
	"Skillan275": {
		{ID: 2758, Key: "Skillan275", Value: `Blades of Ice`},
	},
	"Skillan276": {
		{ID: 2766, Key: "Skillan276", Value: `Dragon Flight`},
	},
	"Skillan277": {
		{ID: 2777, Key: "Skillan277", Value: `Death Sentry`},
	},
	"Skillan278": {
		{ID: 2783, Key: "Skillan278", Value: `Blade Shield`},
	},
	"Skillan279": {
		{ID: 2788, Key: "Skillan279", Value: `Venom`},
	},
	"Skillan280": {
		{ID: 2793, Key: "Skillan280", Value: `Shdw Master`},
	},
	"Skillan281": {
		{ID: 2799, Key: "Skillan281", Value: `Phnx Strike`},
	},
	"Skillld222": {
		{ID: 1770, Key: "Skillld222", Value: `the eyes of your enemies
summon ravens to peck out`},
	},
	"Skillld223": {
		{ID: 1771, Key: "Skillld223", Value: `disease to all it contacts
summon a vine that spreads`},
		{ID: 1774, Key: "Skillld223", Value: `transform into a werewolf`},
	},
	"Skillld225": {
		{ID: 1775, Key: "Skillld225", Value: `when in werewolf or werebear form
passive - improves duration and life`},
	},
	"Skillld226": {
		{ID: 1776, Key: "Skillld226", Value: `unleash fiery chaos to burn your enemies`},
	},
	"Skillld227": {
		{ID: 1777, Key: "Skillld227", Value: `life for you and your party
summon a spirit pet that increases`},
	},
	"Skillld228": {
		{ID: 1780, Key: "Skillld228", Value: `to fight by your side
summon a wolf with teleporting ability`},
	},
	"Skillld229": {
		{ID: 1781, Key: "Skillld229", Value: `transform into a werebear`},
	},
	"Skillld230": {
		{ID: 1782, Key: "Skillld230", Value: `that knocks back your enemies
launch a boulder of flaming hot magma`},
	},
	"Skillld231": {
		{ID: 1783, Key: "Skillld231", Value: `to burn your enemies with frost
blast a continuous jet of ice`},
	},
	"Skillld232": {
		{ID: 1786, Key: "Skillld232", Value: `and replenishes your life
summon a vine that eats corpses`},
	},
	"Skillld233": {
		{ID: 1789, Key: "Skillld233", Value: `with successive hits
increasing amounts of life from your enemies
go into a frenzied rage to steal
when in werewolf form,`},
	},
	"Skillld234": {
		{ID: 1791, Key: "Skillld234", Value: `with successive hits
for increasing extra damage
maul your enemies
when in werebear form,`},
	},
	"Skillld235": {
		{ID: 1792, Key: "Skillld235", Value: `burning them to a crisp
open volcanic vents below your enemies,`},
	},
	"Skillld236": {
		{ID: 1793, Key: "Skillld236", Value: `fire, cold, and lightning
shield yourself from damage caused by`},
	},
	"Skillld237": {
		{ID: 1794, Key: "Skillld237", Value: `of you and your party
to the damage and attack rating
summon a spirit pet that adds`},
	},
	"Skillld238": {
		{ID: 1795, Key: "Skillld238", Value: `it does to enemies
eating corpses to increase damage
summon a wolf that becomes enraged,`},
	},
	"Skillld239": {
		{ID: 1796, Key: "Skillld239", Value: `that spreads to other monsters
to inflict them with disease
bite your enemies
when in werewolf form,`},
	},
	"Skillld240": {
		{ID: 1799, Key: "Skillld240", Value: `with a fiery claw attack
form, maul your enemies
when in werewolf or werebear`},
	},
	"Skillld241": {
		{ID: 1802, Key: "Skillld241", Value: `cut a path through your enemies
release several small whirlwinds that`},
	},
	"Skillld242": {
		{ID: 1803, Key: "Skillld242", Value: `and replenishes your mana
summon a vine that eats corpses`},
	},
	"Skillld243": {
		{ID: 1807, Key: "Skillld243", Value: `to gain life and mana
form, bite your enemies
when in werewolf or werebear`},
	},
	"Skillld244": {
		{ID: 1808, Key: "Skillld244", Value: `that stuns nearby enemies
stomp to create a shock wave
when in werebear form,`},
	},
	"Skillld245": {
		{ID: 2070, Key: "Skillld245", Value: `and destruction over your enemies
summon forth a volcano to rain death`},
	},
	"Skillld246": {
		{ID: 2072, Key: "Skillld246", Value: `to blast your enemies
create a funnel of wind and debris`},
	},
	"Skillld247": {
		{ID: 2077, Key: "Skillld247", Value: `back at your enemies
taken by you and your party
summon spirit pet that reflects damage`},
	},
	"Skillld248": {
		{ID: 2340, Key: "Skillld248", Value: `summon a ferocious grizzly bear`},
	},
	"Skillld249": {
		{ID: 2353, Key: "Skillld249", Value: `or one target multiple times
either multiple adjacent targets
when in werewolf form, attack`},
	},
	"Skillld250": {
		{ID: 2362, Key: "Skillld250", Value: `destruction on nearby enemies
create a meteor shower to rain fiery`},
	},
	"Skillld251": {
		{ID: 2368, Key: "Skillld251", Value: `debris to pound your enemies to bits
create a massive storm of wind and`},
	},
	"Skillld252": {
		{ID: 2374, Key: "Skillld252", Value: `to blast your enemies to bits
throw a fire bomb`},
	},
	"Skillld253": {
		{ID: 2377, Key: "Skillld253", Value: `with claw-class weapons
passive - improves your skill`},
	},
	"Skillld254": {
		{ID: 2380, Key: "Skillld254", Value: `to crush and knock back your enemies
to create a psychic blast
use the power of your mind`},
	},
	"Skillld255": {
		{ID: 2423, Key: "Skillld255", Value: `normal attack to release charges
must use a dragon finishing move or
to finishing moves
consecutive hits add damage bonuses

Charge-up Skill
`},
	},
	"Skillld256": {
		{ID: 2427, Key: "Skillld256", Value: `adds charged-up bonuses to the kick
kick your enemies out of your way

Finishing Move
`},
	},
	"Skillld257": {
		{ID: 2435, Key: "Skillld257", Value: `to shock your enemies
throw a web of lightning`},
	},
	"Skillld258": {
		{ID: 2640, Key: "Skillld258", Value: `between you and target point
set a spinning blade to patrol`},
	},
	"Skillld259": {
		{ID: 2647, Key: "Skillld259", Value: `for a period of time
increases attack and movement speed`},
	},
	"Skillld260": {
		{ID: 2652, Key: "Skillld260", Value: `normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add fire damage

Charge-up Skill
`},
	},
	"Skillld261": {
		{ID: 2656, Key: "Skillld261", Value: `adds charged-up bonuses to both claw attacks
with your dual claw-class weapons
slice and dice your enemies

Finishing Move
`},
	},
	"Skillld262": {
		{ID: 2669, Key: "Skillld262", Value: `at enemies that pass near
a trap that emits charged bolts`},
	},
	"Skillld263": {
		{ID: 2675, Key: "Skillld263", Value: `a trap that emits waves of fire`},
	},
	"Skillld264": {
		{ID: 2683, Key: "Skillld264", Value: `you are using dual claw-class weapons
passive - chance to block when`},
	},
	"Skillld265": {
		{ID: 2687, Key: "Skillld265", Value: `lowering their defenses for a period of time
cast a shadow to blind nearby enemies`},
	},
	"Skillld266": {
		{ID: 2693, Key: "Skillld266", Value: `normal attack to release charges
must use a dragon finishing move or
to finishing moves
consecutive hits add life and mana stealing

Charge-up Skill
`},
	},
	"Skillld267": {
		{ID: 2699, Key: "Skillld267", Value: `to slice up your enemies
throw spinning blades`},
	},
	"Skillld268": {
		{ID: 2707, Key: "Skillld268", Value: `for a period of time
raise all resistances and resist curses`},
	},
	"Skillld269": {
		{ID: 2711, Key: "Skillld269", Value: `your skills and fights by your side
summon a shadow of yourself that mimics`},
	},
	"Skillld270": {
		{ID: 2716, Key: "Skillld270", Value: `normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add lightning damage

Charge-up Skill
`},
	},
	"Skillld271": {
		{ID: 2723, Key: "Skillld271", Value: `adds charged-up bonuses to the kick
knock back your enemies with an explosive kick

Finishing Move
`},
	},
	"Skillld272": {
		{ID: 2727, Key: "Skillld272", Value: `to scorch passing enemies
a trap that shoots lightning`},
	},
	"Skillld273": {
		{ID: 2731, Key: "Skillld273", Value: `trap that sprays fire at passing enemies`},
	},
	"Skillld274": {
		{ID: 2746, Key: "Skillld274", Value: `and convert the feeble-minded
stun a group of enemies
using the power of your mind`},
	},
	"Skillld275": {
		{ID: 2757, Key: "Skillld275", Value: `normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add cold damage

Charge-up Skill
`},
	},
	"Skillld276": {
		{ID: 2765, Key: "Skillld276", Value: `adds charged-up bonuses to the kick
and destroy them with a kick
teleport to your enemies

Finishing Move
`},
	},
	"Skillld277": {
		{ID: 2776, Key: "Skillld277", Value: `or explodes nearby corpses laying waste to more enemies
trap that shoots lightning at your enemies`},
	},
	"Skillld278": {
		{ID: 2782, Key: "Skillld278", Value: `who stray too close
spinning blades slice enemies`},
	},
	"Skillld279": {
		{ID: 2786, Key: "Skillld279", Value: `add poison damage to your weapons`},
	},
	"Skillld280": {
		{ID: 2792, Key: "Skillld280", Value: `to fight by your side
summon a powerful shadow of yourself`},
	},
	"Skillld281": {
		{ID: 2796, Key: "Skillld281", Value: `normal attack to release charges
must use a dragon finishing move or
adds elemental novas to finishing moves

Charge-up Skill
`},
	},
	"Skillname222": {
		{ID: 1332, Key: "Skillname222", Value: `Raven`},
	},
	"Skillname223": {
		{ID: 1493, Key: "Skillname223", Value: `Poison Creeper`},
		{ID: 1494, Key: "Skillname223", Value: `Werewolf`},
	},
	"Skillname225": {
		{ID: 1497, Key: "Skillname225", Value: `Lycanthropy`},
	},
	"Skillname226": {
		{ID: 1499, Key: "Skillname226", Value: `Firestorm`},
	},
	"Skillname227": {
		{ID: 1501, Key: "Skillname227", Value: `Oak Sage`},
	},
	"Skillname228": {
		{ID: 1504, Key: "Skillname228", Value: `Summon Spirit Wolf`},
	},
	"Skillname229": {
		{ID: 1505, Key: "Skillname229", Value: `Werebear`},
	},
	"Skillname230": {
		{ID: 1315, Key: "Skillname230", Value: `Molten Boulder`},
	},
	"Skillname231": {
		{ID: 1506, Key: "Skillname231", Value: `Arctic Blast`},
	},
	"Skillname232": {
		{ID: 1507, Key: "Skillname232", Value: `Carrion Vine`},
	},
	"Skillname233": {
		{ID: 1509, Key: "Skillname233", Value: `Feral Rage`},
	},
	"Skillname234": {
		{ID: 1510, Key: "Skillname234", Value: `Maul`},
	},
	"Skillname235": {
		{ID: 1514, Key: "Skillname235", Value: `Fissure`},
	},
	"Skillname236": {
		{ID: 1515, Key: "Skillname236", Value: `Cyclone Armor`},
	},
	"Skillname237": {
		{ID: 1516, Key: "Skillname237", Value: `Heart of Wolverine`},
	},
	"Skillname238": {
		{ID: 1519, Key: "Skillname238", Value: `Summon Dire Wolf`},
	},
	"Skillname239": {
		{ID: 1520, Key: "Skillname239", Value: `Rabies`},
	},
	"Skillname240": {
		{ID: 1298, Key: "Skillname240", Value: `Fire Claws`},
	},
	"Skillname241": {
		{ID: 1521, Key: "Skillname241", Value: `Twister`},
	},
	"Skillname242": {
		{ID: 1524, Key: "Skillname242", Value: `Solar Creeper`},
	},
	"Skillname243": {
		{ID: 1604, Key: "Skillname243", Value: `Hunger`},
	},
	"Skillname244": {
		{ID: 1605, Key: "Skillname244", Value: `Shock Wave`},
	},
	"Skillname245": {
		{ID: 1606, Key: "Skillname245", Value: `Volcano`},
	},
	"Skillname246": {
		{ID: 1611, Key: "Skillname246", Value: `Tornado`},
	},
	"Skillname247": {
		{ID: 1612, Key: "Skillname247", Value: `Spirit of Barbs`},
	},
	"Skillname248": {
		{ID: 1613, Key: "Skillname248", Value: `Summon Grizzly`},
	},
	"Skillname249": {
		{ID: 2351, Key: "Skillname249", Value: `Fury`},
	},
	"Skillname250": {
		{ID: 2359, Key: "Skillname250", Value: `Armageddon`},
	},
	"Skillname251": {
		{ID: 2367, Key: "Skillname251", Value: `Hurricane`},
	},
	"Skillname252": {
		{ID: 2371, Key: "Skillname252", Value: `Fire Blast`},
	},
	"Skillname253": {
		{ID: 2376, Key: "Skillname253", Value: `Claw Mastery`},
	},
	"Skillname254": {
		{ID: 2379, Key: "Skillname254", Value: `Psychic Hammer`},
	},
	"Skillname255": {
		{ID: 2422, Key: "Skillname255", Value: `Tiger Strike`},
	},
	"Skillname256": {
		{ID: 2426, Key: "Skillname256", Value: `Dragon Talon`},
	},
	"Skillname257": {
		{ID: 2433, Key: "Skillname257", Value: `Shock Web`},
	},
	"Skillname258": {
		{ID: 2638, Key: "Skillname258", Value: `Blade Sentinel`},
	},
	"Skillname259": {
		{ID: 2643, Key: "Skillname259", Value: `Burst of Speed`},
	},
	"Skillname260": {
		{ID: 2649, Key: "Skillname260", Value: `Fists of Fire`},
	},
	"Skillname261": {
		{ID: 2654, Key: "Skillname261", Value: `Dragon Claw`},
	},
	"Skillname262": {
		{ID: 2658, Key: "Skillname262", Value: `Charged Bolt Sentry`},
	},
	"Skillname263": {
		{ID: 2671, Key: "Skillname263", Value: `Wake of Fire`},
	},
	"Skillname264": {
		{ID: 2679, Key: "Skillname264", Value: `Weapon Block`},
	},
	"Skillname265": {
		{ID: 2685, Key: "Skillname265", Value: `Cloak of Shadows`},
	},
	"Skillname266": {
		{ID: 2689, Key: "Skillname266", Value: `Cobra Strike`},
	},
	"Skillname267": {
		{ID: 2697, Key: "Skillname267", Value: `Blade Fury`},
	},
	"Skillname268": {
		{ID: 2705, Key: "Skillname268", Value: `Fade`},
	},
	"Skillname269": {
		{ID: 2709, Key: "Skillname269", Value: `Shadow Warrior`},
	},
	"Skillname270": {
		{ID: 2713, Key: "Skillname270", Value: `Claws of Thunder`},
	},
	"Skillname271": {
		{ID: 2718, Key: "Skillname271", Value: `Dragon Tail`},
	},
	"Skillname272": {
		{ID: 2725, Key: "Skillname272", Value: `Lightning Sentry`},
	},
	"Skillname273": {
		{ID: 2729, Key: "Skillname273", Value: `Wake of Inferno`},
	},
	"Skillname274": {
		{ID: 2735, Key: "Skillname274", Value: `Mind Blast`},
	},
	"Skillname275": {
		{ID: 2752, Key: "Skillname275", Value: `Blades of Ice`},
	},
	"Skillname276": {
		{ID: 2760, Key: "Skillname276", Value: `Dragon Flight`},
	},
	"Skillname277": {
		{ID: 2767, Key: "Skillname277", Value: `Death Sentry`},
	},
	"Skillname278": {
		{ID: 2779, Key: "Skillname278", Value: `Blade Shield`},
	},
	"Skillname279": {
		{ID: 2784, Key: "Skillname279", Value: `Venom`},
	},
	"Skillname280": {
		{ID: 2789, Key: "Skillname280", Value: `Shadow Master`},
	},
	"Skillname281": {
		{ID: 2794, Key: "Skillname281", Value: `Phoenix Strike`},
	},
	"Skillsd222": {
		{ID: 2439, Key: "Skillsd222", Value: `summon ravens`},
	},
	"Skillsd223": {
		{ID: 2442, Key: "Skillsd223", Value: `summon disease spreading vine`},
		{ID: 2443, Key: "Skillsd223", Value: `transform into a werewolf`},
	},
	"Skillsd225": {
		{ID: 2445, Key: "Skillsd225", Value: `passive - improves shape-shifting ability`},
	},
	"Skillsd226": {
		{ID: 2446, Key: "Skillsd226", Value: `unleash fiery chaos`},
	},
	"Skillsd227": {
		{ID: 2537, Key: "Skillsd227", Value: `summon the spirit of the oak`},
	},
	"Skillsd228": {
		{ID: 2538, Key: "Skillsd228", Value: `summon a wolf`},
	},
	"Skillsd229": {
		{ID: 2539, Key: "Skillsd229", Value: `transform into a werebear`},
	},
	"Skillsd230": {
		{ID: 2542, Key: "Skillsd230", Value: `launch a molten boulder`},
	},
	"Skillsd231": {
		{ID: 2544, Key: "Skillsd231", Value: `shoot a jet of ice`},
	},
	"Skillsd232": {
		{ID: 2547, Key: "Skillsd232", Value: `summon corpse eating vine`},
	},
	"Skillsd233": {
		{ID: 2550, Key: "Skillsd233", Value: `life-stealing rage attack - werewolf form`},
	},
	"Skillsd234": {
		{ID: 2551, Key: "Skillsd234", Value: `maul your enemies - werebear form`},
	},
	"Skillsd235": {
		{ID: 2571, Key: "Skillsd235", Value: `open the earth to burn enemies`},
	},
	"Skillsd236": {
		{ID: 2588, Key: "Skillsd236", Value: `shield from elemental damage`},
	},
	"Skillsd237": {
		{ID: 2589, Key: "Skillsd237", Value: `summon a wolverine spirit`},
	},
	"Skillsd238": {
		{ID: 2592, Key: "Skillsd238", Value: `summon an enraged wolf`},
	},
	"Skillsd239": {
		{ID: 2597, Key: "Skillsd239", Value: `bite causes disease - werewolf form`},
	},
	"Skillsd240": {
		{ID: 2598, Key: "Skillsd240", Value: `fiery, mauling attack`},
	},
	"Skillsd241": {
		{ID: 2601, Key: "Skillsd241", Value: `release several small whirlwinds`},
	},
	"Skillsd242": {
		{ID: 2602, Key: "Skillsd242", Value: `summon corpse eating vine`},
	},
	"Skillsd243": {
		{ID: 2603, Key: "Skillsd243", Value: `life-and-mana-stealing bite`},
	},
	"Skillsd244": {
		{ID: 2607, Key: "Skillsd244", Value: `create shock waves - werebear form`},
	},
	"Skillsd245": {
		{ID: 2610, Key: "Skillsd245", Value: `create a volcano`},
	},
	"Skillsd246": {
		{ID: 2611, Key: "Skillsd246", Value: `create a tornado`},
	},
	"Skillsd247": {
		{ID: 2612, Key: "Skillsd247", Value: `summon a spirit pet of thorns`},
	},
	"Skillsd248": {
		{ID: 2615, Key: "Skillsd248", Value: `summon a grizzly bear`},
	},
	"Skillsd249": {
		{ID: 2616, Key: "Skillsd249", Value: `multiple attacks - werewolf Form`},
	},
	"Skillsd250": {
		{ID: 2617, Key: "Skillsd250", Value: `rain fire on your enemies`},
	},
	"Skillsd251": {
		{ID: 2621, Key: "Skillsd251", Value: `create a massive wind storm`},
	},
	"Skillsd252": {
		{ID: 2623, Key: "Skillsd252", Value: `throw a fire bomb`},
	},
	"Skillsd253": {
		{ID: 2624, Key: "Skillsd253", Value: `passive - improves claw-class weapons ability`},
	},
	"Skillsd254": {
		{ID: 2630, Key: "Skillsd254", Value: `a mind blast to crush your enemies`},
	},
	"Skillsd255": {
		{ID: 2634, Key: "Skillsd255", Value: `increases damage of finishing moves

Charge-up Skill`},
	},
	"Skillsd256": {
		{ID: 2635, Key: "Skillsd256", Value: `kick your enemies

Finishing Move`},
	},
	"Skillsd257": {
		{ID: 2636, Key: "Skillsd257", Value: `throw a web of lightning`},
	},
	"Skillsd258": {
		{ID: 2639, Key: "Skillsd258", Value: `set a spinning blade`},
	},
	"Skillsd259": {
		{ID: 2646, Key: "Skillsd259", Value: `increases attack and movement speed`},
	},
	"Skillsd260": {
		{ID: 2650, Key: "Skillsd260", Value: `adds fire damage to finishing moves

Charge-up Skill`},
	},
	"Skillsd261": {
		{ID: 2655, Key: "Skillsd261", Value: `double claw attack

Finishing Move`},
	},
	"Skillsd262": {
		{ID: 2659, Key: "Skillsd262", Value: `a trap that emits charged bolts`},
	},
	"Skillsd263": {
		{ID: 2672, Key: "Skillsd263", Value: `a trap that emits waves of fire`},
	},
	"Skillsd264": {
		{ID: 2682, Key: "Skillsd264", Value: `passive - block with two claw-class weapons`},
	},
	"Skillsd265": {
		{ID: 2686, Key: "Skillsd265", Value: `blind your enemies`},
	},
	"Skillsd266": {
		{ID: 2690, Key: "Skillsd266", Value: ` adds life and mana stealing to finishing moves

Charge-up Skill`},
	},
	"Skillsd267": {
		{ID: 2698, Key: "Skillsd267", Value: `throw spinning blades`},
	},
	"Skillsd268": {
		{ID: 2706, Key: "Skillsd268", Value: `temporary resist all`},
	},
	"Skillsd269": {
		{ID: 2710, Key: "Skillsd269", Value: `summon a shadow of yourself`},
	},
	"Skillsd270": {
		{ID: 2714, Key: "Skillsd270", Value: `adds lightning damage to finishing moves
 
Charge-up Skill`},
	},
	"Skillsd271": {
		{ID: 2719, Key: "Skillsd271", Value: `explosive kick

Finishing Move`},
	},
	"Skillsd272": {
		{ID: 2726, Key: "Skillsd272", Value: `a trap that emits lightning`},
	},
	"Skillsd273": {
		{ID: 2730, Key: "Skillsd273", Value: `trap that sprays fire`},
	},
	"Skillsd274": {
		{ID: 2743, Key: "Skillsd274", Value: `compelling psionic blast`},
	},
	"Skillsd275": {
		{ID: 2753, Key: "Skillsd275", Value: `adds cold damage to finishing moves

Charge-up Skill`},
	},
	"Skillsd276": {
		{ID: 2764, Key: "Skillsd276", Value: `teleport and kick enemies

Finishing Move`},
	},
	"Skillsd277": {
		{ID: 2774, Key: "Skillsd277", Value: `trap that explodes nearby corpses`},
	},
	"Skillsd278": {
		{ID: 2781, Key: "Skillsd278", Value: `spinning blades of defense`},
	},
	"Skillsd279": {
		{ID: 2785, Key: "Skillsd279", Value: `poison your weapon`},
	},
	"Skillsd280": {
		{ID: 2790, Key: "Skillsd280", Value: `summon your shadow`},
	},
	"Skillsd281": {
		{ID: 2795, Key: "Skillsd281", Value: `adds elemental novas to finishing moves

Charge-up Skill`},
	},
	"Skin of the Flayed One": {
		{ID: 1013, Key: "Skin of the Flayed One", Value: `Skin of Flayed One`},
	},
	"Skin of the Flayerd One": {
		{ID: 403, Key: "Skin of the Flayerd One", Value: `Skin of the Flayed One`},
	},
	"Skin of the Vipermagi": {
		{ID: 1112, Key: "Skin of the Vipermagi", Value: `Skin of the Vipermagi`},
	},
	"Skullcollector": {
		{ID: 332, Key: "Skullcollector", Value: `Skull Collector`},
	},
	"Skullder's Ire": {
		{ID: 2548, Key: "Skullder's Ire", Value: `Skullder's Ire`},
	},
	"Skulltred": {
		{ID: 2696, Key: "Skulltred", Value: `Skulltred`},
	},
	"Skystrike": {
		{ID: 2394, Key: "Skystrike", Value: `Skystrike`},
	},
	"Slayerexp": {
		{ID: 90, Key: "Slayerexp", Value: `Slayer`},
	},
	"Smiting": {
		{ID: 1801, Key: "Smiting", Value: `Smiting`},
	},
	"Smoking": {
		{ID: 141, Key: "Smoking", Value: `Smoking`},
	},
	"Smoldering": {
		{ID: 2459, Key: "Smoldering", Value: `Smoldering`},
	},
	"Snaketongue": {
		{ID: 2207, Key: "Snaketongue", Value: `Snake Tongue`},
	},
	"Snapchip Shatter": {
		{ID: 296, Key: "Snapchip Shatter", Value: `Snapchip Shatter`},
	},
	"SnowYeti1": {
		{ID: 923, Key: "SnowYeti1", Value: `Snow Drifter`},
	},
	"SnowYeti2": {
		{ID: 979, Key: "SnowYeti2", Value: `Abominable`},
	},
	"SnowYeti3": {
		{ID: 980, Key: "SnowYeti3", Value: `Chilled Froth`},
	},
	"SnowYeti4": {
		{ID: 984, Key: "SnowYeti4", Value: `Frozen Abyss`},
	},
	"Snowclash": {
		{ID: 976, Key: "Snowclash", Value: `Snowclash`},
	},
	"Snowflake": {
		{ID: 377, Key: "Snowflake", Value: `Snowy`},
	},
	"SorOnly": {
		{ID: 1487, Key: "SorOnly", Value: `(Sorceress Only)`},
	},
	"Souldrain": {
		{ID: 1479, Key: "Souldrain", Value: `Soul Drainer`},
	},
	"Soulfeast Tine": {
		{ID: 223, Key: "Soulfeast Tine", Value: `Soulfeast Tine`},
	},
	"Soulless": {
		{ID: 2660, Key: "Soulless", Value: `Soulless`},
	},
	"Soulreaper": {
		{ID: 1595, Key: "Soulreaper", Value: `Soul Reaper`},
	},
	"Sounding": {
		{ID: 221, Key: "Sounding", Value: `Sounding`},
	},
	"Sparking": {
		{ID: 1048, Key: "Sparking", Value: `Sparking`},
	},
	"Sparroweye": {
		{ID: 2460, Key: "Sparroweye", Value: `Sparroweye`},
	},
	"Spearmaiden's": {
		{ID: 2373, Key: "Spearmaiden's", Value: `Spearmaiden's`},
	},
	"Spellsteel": {
		{ID: 790, Key: "Spellsteel", Value: `Spellsteel`},
	},
	"Spike Generator": {
		{ID: 58, Key: "Spike Generator", Value: `Spike Generator`},
	},
	"Spike Thorn": {
		{ID: 1448, Key: "Spike Thorn", Value: `Spike Thorn`},
	},
	"Spineripper": {
		{ID: 983, Key: "Spineripper", Value: `Spineripper`},
	},
	"Spire of Honor": {
		{ID: 774, Key: "Spire of Honor", Value: `Spire of Honor`},
	},
	"Spirit Ward": {
		{ID: 2510, Key: "Spirit Ward", Value: `Spirit Ward`},
	},
	"Spirit of Barbs": {
		{ID: 241, Key: "Spirit of Barbs", Value: `Spirit of Barbs`},
	},
	"Spiritforge": {
		{ID: 1270, Key: "Spiritforge", Value: `Spirit Forge`},
	},
	"Spiritkeeper": {
		{ID: 1435, Key: "Spiritkeeper", Value: `Spirit Keeper`},
	},
	"Spiritual": {
		{ID: 405, Key: "Spiritual", Value: `Spiritual`},
	},
	"Spiritual Custodian": {
		{ID: 799, Key: "Spiritual Custodian", Value: `Dark Adherent`},
	},
	"Spiteful": {
		{ID: 1537, Key: "Spiteful", Value: `Spiteful`},
	},
	"Squire's Cover": {
		{ID: 254, Key: "Squire's Cover", Value: `Squire's Cover`},
	},
	"Stalwart": {
		{ID: 985, Key: "Stalwart", Value: `Stalwart`},
	},
	"Static": {
		{ID: 454, Key: "Static", Value: `Static`},
	},
	"Stealskull": {
		{ID: 2284, Key: "Stealskull", Value: `Stealskull`},
	},
	"Steel Carapice": {
		{ID: 1598, Key: "Steel Carapice", Value: `Steel Carapace`},
	},
	"Steelpillar": {
		{ID: 2174, Key: "Steelpillar", Value: `Steel Pillar`},
	},
	"Steelrend": {
		{ID: 2161, Key: "Steelrend", Value: `Steelrend`},
	},
	"Steelshade": {
		{ID: 1727, Key: "Steelshade", Value: `Steel Shade`},
	},
	"Steller": {
		{ID: 57, Key: "Steller", Value: `Stellar`},
	},
	"Stinging": {
		{ID: 615, Key: "Stinging", Value: `Stinging`},
	},
	"Stinky": {
		{ID: 731, Key: "Stinky", Value: `Odiferous`},
	},
	"Stone Crusher": {
		{ID: 552, Key: "Stone Crusher", Value: `Stone Crusher`},
	},
	"Stonerage": {
		{ID: 2304, Key: "Stonerage", Value: `Stonerage`},
	},
	"Stonerattle": {
		{ID: 64, Key: "Stonerattle", Value: `Stone Rattle`},
	},
	"Stoneraven": {
		{ID: 1026, Key: "Stoneraven", Value: `Stoneraven`},
	},
	"Stormchaser": {
		{ID: 2210, Key: "Stormchaser", Value: `Stormchaser`},
	},
	"Stormlash": {
		{ID: 2629, Key: "Stormlash", Value: `Stormlash`},
	},
	"Stormrider": {
		{ID: 2276, Key: "Stormrider", Value: `Stormrider`},
	},
	"Stormshield": {
		{ID: 1936, Key: "Stormshield", Value: `Stormshield`},
	},
	"Stormspike": {
		{ID: 2314, Key: "Stormspike", Value: `Stormspike`},
	},
	"Stormspire": {
		{ID: 2420, Key: "Stormspire", Value: `Stormspire`},
	},
	"Stormwillow": {
		{ID: 592, Key: "Stormwillow", Value: `Stormwillow`},
	},
	"Stout": {
		{ID: 2576, Key: "Stout", Value: `Stout`},
	},
	"StrGemX1": {
		{ID: 1020, Key: "StrGemX1", Value: `Helms:`},
	},
	"StrGemX2": {
		{ID: 1021, Key: "StrGemX2", Value: `Shields:`},
	},
	"StrGemX3": {
		{ID: 1022, Key: "StrGemX3", Value: `Weapons:`},
	},
	"StrGemX4": {
		{ID: 1023, Key: "StrGemX4", Value: `Armor:`},
	},
	"StrMercEx12": {
		{ID: 2352, Key: "StrMercEx12", Value: `Defensive`},
	},
	"StrMercEx14": {
		{ID: 2356, Key: "StrMercEx14", Value: `Offensive`},
	},
	"StrMercEx15": {
		{ID: 2357, Key: "StrMercEx15", Value: `Combat`},
	},
	"StrSklTree26": {
		{ID: 798, Key: "StrSklTree26", Value: `Summoning`},
	},
	"StrSklTree27": {
		{ID: 800, Key: "StrSklTree27", Value: `Shape`},
	},
	"StrSklTree28": {
		{ID: 801, Key: "StrSklTree28", Value: `Shifting`},
	},
	"StrSklTree29": {
		{ID: 802, Key: "StrSklTree29", Value: `Elemental`},
	},
	"StrSklTree30": {
		{ID: 755, Key: "StrSklTree30", Value: `Traps`},
	},
	"StrSklTree31": {
		{ID: 758, Key: "StrSklTree31", Value: `Shadow`},
	},
	"StrSklTree32": {
		{ID: 759, Key: "StrSklTree32", Value: `Disciplines`},
	},
	"StrSklTree33": {
		{ID: 760, Key: "StrSklTree33", Value: `Martial`},
	},
	"StrSklTree34": {
		{ID: 762, Key: "StrSklTree34", Value: `Arts`},
	},
	"StrUI18": {
		{ID: 974, Key: "StrUI18", Value: `Master`},
	},
	"Strange": {
		{ID: 1234, Key: "Strange", Value: `Strange`},
	},
	"String of Ears": {
		{ID: 2343, Key: "String of Ears", Value: `String of Ears`},
	},
	"Strongoak": {
		{ID: 1714, Key: "Strongoak", Value: `Strong Oak`},
	},
	"StygianFury": {
		{ID: 199, Key: "StygianFury", Value: `Stygian Fury`},
	},
	"StygianHarlot": {
		{ID: 1171, Key: "StygianHarlot", Value: `Stygian Harlot`},
	},
	"Succubusexp": {
		{ID: 111, Key: "Succubusexp", Value: `Succubus`},
	},
	"Suicide Branch": {
		{ID: 906, Key: "Suicide Branch", Value: `Suicide Branch`},
	},
	"Sureshrill Frost": {
		{ID: 1223, Key: "Sureshrill Frost", Value: `Sureshrill Frost`},
	},
	"Swordguard": {
		{ID: 855, Key: "Swordguard", Value: `Swordguard`},
	},
	"Syrian": {
		{ID: 174, Key: "Syrian", Value: `Triad's Foliage`},
	},
	"Tal Rasha's Adjudication": {
		{ID: 1739, Key: "Tal Rasha's Adjudication", Value: `Tal Rasha's Adjudication`},
	},
	"Tal Rasha's Fire-Spun Cloth": {
		{ID: 954, Key: "Tal Rasha's Fire-Spun Cloth", Value: `Tal Rasha's Fine-Spun Cloth`},
	},
	"Tal Rasha's Horadric Crest": {
		{ID: 2311, Key: "Tal Rasha's Horadric Crest", Value: `Tal Rasha's Horadric Crest`},
	},
	"Tal Rasha's Howling Wind": {
		{ID: 1664, Key: "Tal Rasha's Howling Wind", Value: `Tal Rasha's Guardianship`},
	},
	"Tal Rasha's Lidless Eye": {
		{ID: 2200, Key: "Tal Rasha's Lidless Eye", Value: `Tal Rasha's Lidless Eye`},
	},
	"Tal Rasha's Wrappings": {
		{ID: 713, Key: "Tal Rasha's Wrappings", Value: `Tal Rasha's Wrappings`},
	},
	"Talberd's Law": {
		{ID: 1312, Key: "Talberd's Law", Value: `Talberd's Law`},
	},
	"Tang's Battle Standard": {
		{ID: 2348, Key: "Tang's Battle Standard", Value: `Tang's Battle Standard`},
	},
	"Tang's Fore-Fathers": {
		{ID: 2541, Key: "Tang's Fore-Fathers", Value: `Tang's Fore-Fathers`},
	},
	"Tang's Imperial Robes": {
		{ID: 2298, Key: "Tang's Imperial Robes", Value: `Tang's Imperial Robes`},
	},
	"Tang's Rule": {
		{ID: 1070, Key: "Tang's Rule", Value: `Tang's Rule`},
	},
	"Tang's Throne": {
		{ID: 2398, Key: "Tang's Throne", Value: `Tang's Throne`},
	},
	"Telling of Beads": {
		{ID: 412, Key: "Telling of Beads", Value: `Telling of Beads`},
	},
	"Templar's Might": {
		{ID: 2396, Key: "Templar's Might", Value: `Templar's Might`},
	},
	"Terra": {
		{ID: 1735, Key: "Terra", Value: `Terra's Guardian`},
	},
	"Terra's": {
		{ID: 1632, Key: "Terra's", Value: `Terrene`},
	},
	"The Archon Magus": {
		{ID: 619, Key: "The Archon Magus", Value: `The Achron Magus`},
	},
	"The Atlantian": {
		{ID: 2605, Key: "The Atlantian", Value: `The Atlantean`},
	},
	"The Atlantien": {
		{ID: 2668, Key: "The Atlantien", Value: `The Atlantean`},
	},
	"The Black Adder": {
		{ID: 2361, Key: "The Black Adder", Value: `The Black Adder`},
	},
	"The Cat's Eye": {
		{ID: 1977, Key: "The Cat's Eye", Value: `The Cat's Eye`},
	},
	"The Cranium Basher": {
		{ID: 953, Key: "The Cranium Basher", Value: `The Cranium Basher`},
	},
	"The Disciple": {
		{ID: 2182, Key: "The Disciple", Value: `The Disciple`},
	},
	"The Ensanguinator": {
		{ID: 1245, Key: "The Ensanguinator", Value: `The Ensanguinator`},
	},
	"The Fetid Sprinkler": {
		{ID: 1913, Key: "The Fetid Sprinkler", Value: `The Fetid Sprinkler`},
	},
	"The Gavel of Pain": {
		{ID: 302, Key: "The Gavel of Pain", Value: `The Gavel of Pain`},
	},
	"The Gladiator's Bane": {
		{ID: 2644, Key: "The Gladiator's Bane", Value: `The Gladiator's Bane`},
	},
	"The Grandfather": {
		{ID: 1982, Key: "The Grandfather", Value: `The Grandfather`},
	},
	"The Harbinger": {
		{ID: 763, Key: "The Harbinger", Value: `The Harbinger`},
	},
	"The Harvester": {
		{ID: 2279, Key: "The Harvester", Value: `The Harvester`},
	},
	"The Impaler": {
		{ID: 2759, Key: "The Impaler", Value: `The Impaler`},
	},
	"The Long Rod": {
		{ID: 2733, Key: "The Long Rod", Value: `The Longest Rod`},
	},
	"The Meat Scraper": {
		{ID: 2282, Key: "The Meat Scraper", Value: `The Meat Scraper`},
	},
	"The Minataur": {
		{ID: 1248, Key: "The Minataur", Value: `The Minotaur`},
	},
	"The Minotaur": {
		{ID: 2554, Key: "The Minotaur", Value: `The Minotaur`},
	},
	"The Oculus": {
		{ID: 271, Key: "The Oculus", Value: `The Oculus`},
	},
	"The Prowler": {
		{ID: 514, Key: "The Prowler", Value: `The Prowler`},
	},
	"The Reaper's Toll": {
		{ID: 2780, Key: "The Reaper's Toll", Value: `The Reaper's Toll`},
	},
	"The Reedeemer": {
		{ID: 364, Key: "The Reedeemer", Value: `The Redeemer`},
	},
	"The Rising Sun": {
		{ID: 1933, Key: "The Rising Sun", Value: `The Rising Sun`},
	},
	"The Scalper": {
		{ID: 1162, Key: "The Scalper", Value: `The Scalper`},
	},
	"The Spirit Shroud": {
		{ID: 1043, Key: "The Spirit Shroud", Value: `The Spirit Shroud`},
	},
	"The Swashbuckler": {
		{ID: 1436, Key: "The Swashbuckler", Value: `The Swashbuckler`},
	},
	"The TreeEnt": {
		{ID: 1224, Key: "The TreeEnt", Value: `The Treentster`},
	},
	"The Vicar": {
		{ID: 2803, Key: "The Vicar", Value: `The Vicar`},
	},
	"The Vile Husk": {
		{ID: 2309, Key: "The Vile Husk", Value: `The Vile Husk`},
	},
	"The Wanderer's Blade": {
		{ID: 2521, Key: "The Wanderer's Blade", Value: `The Wanderer's Blade`},
	},
	"The Worldstone Chamber": {
		{ID: 548, Key: "The Worldstone Chamber", Value: `The Worldstone Chamber`},
	},
	"The Worldstone Keep Level 1": {
		{ID: 1217, Key: "The Worldstone Keep Level 1", Value: `Worldstone Keep Level 1`},
	},
	"The Worldstone Keep Level 2": {
		{ID: 1218, Key: "The Worldstone Keep Level 2", Value: `Worldstone Keep Level 2`},
	},
	"The Worldstone Keep Level 3": {
		{ID: 1220, Key: "The Worldstone Keep Level 3", Value: `Worldstone Keep Level 3`},
	},
	"Thin": {
		{ID: 371, Key: "Thin", Value: `Thin`},
	},
	"Threash Socket": {
		{ID: 988, Key: "Threash Socket", Value: `Thresh Socket`},
	},
	"Throne of Destruction": {
		{ID: 549, Key: "Throne of Destruction", Value: `Throne of Destruction`},
	},
	"Thudergod's Vigor": {
		{ID: 1133, Key: "Thudergod's Vigor", Value: `Thundergod's Vigor`},
	},
	"Thunderstroke": {
		{ID: 1115, Key: "Thunderstroke", Value: `Thunderstroke`},
	},
	"Tiamat's Rebuke": {
		{ID: 1926, Key: "Tiamat's Rebuke", Value: `Tiamat's Rebuke`},
	},
	"Timeless": {
		{ID: 1713, Key: "Timeless", Value: `Timeless`},
	},
	"Tin": {
		{ID: 743, Key: "Tin", Value: `Tin`},
	},
	"Titan's Revenge": {
		{ID: 591, Key: "Titan's Revenge", Value: `Titan's Revenge`},
	},
	"Titanfist": {
		{ID: 272, Key: "Titanfist", Value: `Titan's Fist`},
	},
	"To Harrogath": {
		{ID: 656, Key: "To Harrogath", Value: `To Harrogath`},
	},
	"To Hell1": {
		{ID: 2575, Key: "To Hell1", Value: `To Abaddon`},
	},
	"To Hell2": {
		{ID: 2582, Key: "To Hell2", Value: `To The Pit of Acheron`},
	},
	"To Hell3": {
		{ID: 2585, Key: "To Hell3", Value: `To The Infernal Pit`},
	},
	"To Nihlathaks Temple": {
		{ID: 577, Key: "To Nihlathaks Temple", Value: `To Nihlathak's Temple`},
	},
	"To The Arreat Plateau": {
		{ID: 564, Key: "To The Arreat Plateau", Value: `To The Arreat Plateau`},
	},
	"To The Bloody Foothills": {
		{ID: 561, Key: "To The Bloody Foothills", Value: `To The Bloody Foothills`},
	},
	"To The Cellar of Pity": {
		{ID: 568, Key: "To The Cellar of Pity", Value: `To The Frozen River`},
	},
	"To The Crystalized Cavern Level 1": {
		{ID: 567, Key: "To The Crystalized Cavern Level 1", Value: `To The Crystalline Passage`},
	},
	"To The Crystalized Cavern Level 2": {
		{ID: 569, Key: "To The Crystalized Cavern Level 2", Value: `To The Glacial Trail`},
	},
	"To The Echo Chamber": {
		{ID: 570, Key: "To The Echo Chamber", Value: `To The Drifter Cavern`},
	},
	"To The Glacier Caves Level 1": {
		{ID: 573, Key: "To The Glacier Caves Level 1", Value: `To The Ancients' Way`},
	},
	"To The Glacier Caves Level 2": {
		{ID: 576, Key: "To The Glacier Caves Level 2", Value: `To The Icy Cellar`},
	},
	"To The Halls of Anguish": {
		{ID: 578, Key: "To The Halls of Anguish", Value: `To The Halls of Anguish`},
	},
	"To The Halls of Death's Calling": {
		{ID: 580, Key: "To The Halls of Death's Calling", Value: `To The Halls of Pain`},
	},
	"To The Halls of Tormented Insanity": {
		{ID: 581, Key: "To The Halls of Tormented Insanity", Value: `To The Halls of Torment`},
	},
	"To The Halls of Vaught": {
		{ID: 584, Key: "To The Halls of Vaught", Value: `To The Halls of Vaught`},
	},
	"To The Rigid Highlands": {
		{ID: 1238, Key: "To The Rigid Highlands", Value: `To The Frigid Highlands`},
	},
	"To The Rocky Summit": {
		{ID: 1024, Key: "To The Rocky Summit", Value: `To The Arreat Summit`},
	},
	"To The Throne of Destruction": {
		{ID: 681, Key: "To The Throne of Destruction", Value: `To The Throne of Destruction`},
	},
	"To The Tundra Wastelands": {
		{ID: 571, Key: "To The Tundra Wastelands", Value: `To The Frozen Tundra`},
	},
	"To The Worldstone Chamber": {
		{ID: 1007, Key: "To The Worldstone Chamber", Value: `To The Worldstone Chamber`},
	},
	"To The Worldstone Keep Level 1": {
		{ID: 589, Key: "To The Worldstone Keep Level 1", Value: `To The Worldstone Keep Level 1`},
	},
	"To The Worldstone Keep Level 2": {
		{ID: 595, Key: "To The Worldstone Keep Level 2", Value: `To The Worldstone Keep Level 2`},
	},
	"To The Worldstone Keep Level 3": {
		{ID: 596, Key: "To The Worldstone Keep Level 3", Value: `To The Worldstone Keep Level 3`},
	},
	"Todesfaelle Flamme": {
		{ID: 1539, Key: "Todesfaelle Flamme", Value: `Todesfaelle Flamme`},
	},
	"Tomb Reaver": {
		{ID: 2016, Key: "Tomb Reaver", Value: `Tomb Reaver`},
	},
	"Toothrow": {
		{ID: 828, Key: "Toothrow", Value: `Toothrow`},
	},
	"Toxic": {
		{ID: 1805, Key: "Toxic", Value: `Toxic`},
	},
	"Trainer's": {
		{ID: 815, Key: "Trainer's", Value: `Trainer's`},
	},
	"Trang-Oul's Avatar": {
		{ID: 1175, Key: "Trang-Oul's Avatar", Value: `Trang-Oul's Avatar`},
	},
	"Trang-Oul's Claws": {
		{ID: 347, Key: "Trang-Oul's Claws", Value: `Trang-Oul's Claws`},
	},
	"Trang-Oul's Girth": {
		{ID: 304, Key: "Trang-Oul's Girth", Value: `Trang-Oul's Girth`},
	},
	"Trang-Oul's Guise": {
		{ID: 1481, Key: "Trang-Oul's Guise", Value: `Trang-Oul's Guise`},
	},
	"Trang-Oul's Mask": {
		{ID: 607, Key: "Trang-Oul's Mask", Value: `Trang-Oul's Mask`},
	},
	"Trang-Oul's Scales": {
		{ID: 1, Key: "Trang-Oul's Scales", Value: `Trang-Oul's Scales`},
	},
	"Trang-Oul's Wing": {
		{ID: 1453, Key: "Trang-Oul's Wing", Value: `Trang-Oul's Wing`},
	},
	"Travel To Harrogath": {
		{ID: 1211, Key: "Travel To Harrogath", Value: `travel to harrogath`},
	},
	"Trickster's": {
		{ID: 1551, Key: "Trickster's", Value: `Trickster's`},
	},
	"Trump": {
		{ID: 2125, Key: "Trump", Value: `Fool's`},
	},
	"Tundra Wastelands": {
		{ID: 1213, Key: "Tundra Wastelands", Value: `Frozen Tundra`},
	},
	"Turquoise": {
		{ID: 1401, Key: "Turquoise", Value: `Turquoise`},
	},
	"Tyrael's Mercy": {
		{ID: 1266, Key: "Tyrael's Mercy", Value: `Tyrael's Mercy`},
	},
	"Tyrael's Might": {
		{ID: 1098, Key: "Tyrael's Might", Value: `Tyrael's Might`},
	},
	"UIFenirsui": {
		{ID: 48, Key: "UIFenirsui", Value: `Dire Wolf`},
	},
	"UiRescuedBarUI": {
		{ID: 49, Key: "UiRescuedBarUI", Value: `Warrior`},
	},
	"UiShadowUI": {
		{ID: 50, Key: "UiShadowUI", Value: `Shadow`},
	},
	"Unearthly": {
		{ID: 951, Key: "Unearthly", Value: `Unearthly`},
	},
	"UnholyCorpse": {
		{ID: 76, Key: "UnholyCorpse", Value: `Unholy Corpse`},
	},
	"Valkiry Wing": {
		{ID: 1037, Key: "Valkiry Wing", Value: `Valkyrie Wing`},
	},
	"Vampiregaze": {
		{ID: 2251, Key: "Vampiregaze", Value: `Vampire Gaze`},
	},
	"Veil of Steel": {
		{ID: 1981, Key: "Veil of Steel", Value: `Veil of Steel`},
	},
	"Venom Grip": {
		{ID: 1511, Key: "Venom Grip", Value: `Venom Grip`},
	},
	"Venomous": {
		{ID: 410, Key: "Venomous", Value: `Venomous`},
	},
	"Verdugo's Hearty Cord": {
		{ID: 2005, Key: "Verdugo's Hearty Cord", Value: `Verdungo's Hearty Cord`},
	},
	"Vermillion": {
		{ID: 1899, Key: "Vermillion", Value: `Vermilion`},
	},
	"Veteran's": {
		{ID: 2345, Key: "Veteran's", Value: `Veteran's`},
	},
	"Victorious": {
		{ID: 2673, Key: "Victorious", Value: `Victorious`},
	},
	"Vigorous": {
		{ID: 1552, Key: "Vigorous", Value: `Vigorous`},
	},
	"VileTemptress": {
		{ID: 112, Key: "VileTemptress", Value: `Vile Temptress`},
	},
	"VileWitch": {
		{ID: 167, Key: "VileWitch", Value: `Vile Witch`},
	},
	"Vine Creature": {
		{ID: 593, Key: "Vine Creature", Value: `Vine Creature`},
	},
	"Vinvear Molech": {
		{ID: 751, Key: "Vinvear Molech", Value: `Vinvear Molech`},
	},
	"Viperfork": {
		{ID: 106, Key: "Viperfork", Value: `Viperfork`},
	},
	"Visceratuant": {
		{ID: 274, Key: "Visceratuant", Value: `Visceratuant`},
	},
	"Visionary": {
		{ID: 642, Key: "Visionary", Value: `Visionary`},
	},
	"Vodoun": {
		{ID: 2255, Key: "Vodoun", Value: `Mojo`},
	},
	"Volcanic": {
		{ID: 2145, Key: "Volcanic", Value: `Volcanic`},
	},
	"Wailing": {
		{ID: 897, Key: "Wailing", Value: `Wailing`},
	},
	"WailingSpirit": {
		{ID: 977, Key: "WailingSpirit", Value: `Wailing Spirit`},
	},
	"Wake of Destruction": {
		{ID: 250, Key: "Wake of Destruction", Value: `Wake of Fire`},
	},
	"Wallace's Tear": {
		{ID: 2662, Key: "Wallace's Tear", Value: `Wallace's Tear`},
	},
	"Warder's": {
		{ID: 1652, Key: "Warder's", Value: `Warden's`},
	},
	"Warhound": {
		{ID: 1168, Key: "Warhound", Value: `Warhound`},
	},
	"Warlord's Trust": {
		{ID: 789, Key: "Warlord's Trust", Value: `Warlord's Trust`},
	},
	"Warpspear": {
		{ID: 1289, Key: "Warpspear", Value: `Warpspear`},
	},
	"Warriv's Warder": {
		{ID: 1927, Key: "Warriv's Warder", Value: `Warriv's Warder`},
	},
	"Warshrike": {
		{ID: 2651, Key: "Warshrike", Value: `Warshrike`},
	},
	"Wartraveler": {
		{ID: 1442, Key: "Wartraveler", Value: `War Traveler`},
	},
	"Waterwalk": {
		{ID: 852, Key: "Waterwalk", Value: `Waterwalk`},
	},
	"Waypoint": {
		{ID: 990, Key: "Waypoint", Value: `Waypoint`},
	},
	"WeaponDescH2H": {
		{ID: 1569, Key: "WeaponDescH2H", Value: `Claw Class`},
	},
	"WeaponDescOrb": {
		{ID: 2546, Key: "WeaponDescOrb", Value: `Orb Class`},
	},
	"Whichwild String": {
		{ID: 2160, Key: "Whichwild String", Value: `Whichwild String`},
	},
	"Wicked": {
		{ID: 1041, Key: "Wicked", Value: `Wicked`},
	},
	"Widow maker": {
		{ID: 1928, Key: "Widow maker", Value: `Widowmaker`},
	},
	"Widowmaker": {
		{ID: 545, Key: "Widowmaker", Value: `Widowmaker`},
	},
	"Wihtstan's Guard": {
		{ID: 478, Key: "Wihtstan's Guard", Value: `Whitstan's Guard`},
	},
	"Wilhelm's Pride": {
		{ID: 2290, Key: "Wilhelm's Pride", Value: `Wilhelm's Pride`},
	},
	"Willhelm's Pride": {
		{ID: 2322, Key: "Willhelm's Pride", Value: `Willhelm's Pride`},
	},
	"Windforce": {
		{ID: 2178, Key: "Windforce", Value: `Windforce`},
	},
	"Windhammer": {
		{ID: 2046, Key: "Windhammer", Value: `Windhammer`},
	},
	"Windstrike": {
		{ID: 1700, Key: "Windstrike", Value: `Windstrike`},
	},
	"Wisp": {
		{ID: 1800, Key: "Wisp", Value: `Wisp Projector`},
	},
	"Witch-hunter's": {
		{ID: 1205, Key: "Witch-hunter's", Value: `Witch-hunter's`},
	},
	"Wizardspike": {
		{ID: 99, Key: "Wizardspike", Value: `Wizardspike`},
	},
	"Wolf": {
		{ID: 396, Key: "Wolf", Value: `Wolf`},
	},
	"WolfRider1": {
		{ID: 672, Key: "WolfRider1", Value: `Fallen Wolfrider`},
	},
	"WolfRider2": {
		{ID: 664, Key: "WolfRider2", Value: `Carver Wolfrider`},
	},
	"WolfRider3": {
		{ID: 671, Key: "WolfRider3", Value: `Devilkin Wolfrider`},
	},
	"WolfRider4": {
		{ID: 666, Key: "WolfRider4", Value: `Darkone Wolfrider`},
	},
	"WolfRider5": {
		{ID: 670, Key: "WolfRider5", Value: `Warped Fallen Wolfrider`},
	},
	"Wolfhowl": {
		{ID: 2441, Key: "Wolfhowl", Value: `Wolfhowl`},
	},
	"Wraithfang": {
		{ID: 2346, Key: "Wraithfang", Value: `Wraithfang`},
	},
	"Wraithflight": {
		{ID: 55, Key: "Wraithflight", Value: `Wraith Flight`},
	},
	"Wrath of Cain": {
		{ID: 2365, Key: "Wrath of Cain", Value: `Wrath of Cain`},
	},
	"Wyvern's Head": {
		{ID: 1138, Key: "Wyvern's Head", Value: `Wyvern's Head`},
	},
	"X": {
		{ID: 809, Key: "X", Value: `Warning:  Once you convert a character to expansion, you cannot play it in original Diablo II games.`},
		{ID: 810, Key: "X", Value: ``},
		{ID: 813, Key: "X", Value: `Are you sure you wish to continue?`},
		{ID: 826, Key: "X", Value: `CONVERT TO EXPANSION`},
		{ID: 832, Key: "X", Value: `CREATE NEW CHARACTER`},
		{ID: 833, Key: "X", Value: `DELETE CHARACTER`},
		{ID: 834, Key: "X", Value: `Ethereal (Cannot be Repaired)`},
		{ID: 835, Key: "X", Value: `This item cannot be repaired.`},
	},
	"Xenos": {
		{ID: 660, Key: "Xenos", Value: `Xenos`},
	},
	"Yelling": {
		{ID: 2507, Key: "Yelling", Value: `Yelling`},
	},
	"Zakarum's Hand": {
		{ID: 2666, Key: "Zakarum's Hand", Value: `Zakarum's Hand`},
	},
	"Zakarum's Salvation": {
		{ID: 259, Key: "Zakarum's Salvation", Value: `Zakarum's Salvation`},
	},
	"Zakrum's Hand": {
		{ID: 508, Key: "Zakrum's Hand", Value: `Zakarum's Hand`},
	},
	"Zircon": {
		{ID: 1128, Key: "Zircon", Value: `Zircon`},
	},
	"act1X": {
		{ID: 2189, Key: "act1X", Value: `THE SISTER'S LAMENT`},
	},
	"act2X": {
		{ID: 2190, Key: "act2X", Value: `DESERT JOURNEY`},
	},
	"act3X": {
		{ID: 2191, Key: "act3X", Value: `MEPHISTO'S JUNGLE`},
	},
	"act4X": {
		{ID: 2192, Key: "act4X", Value: `ENTER HELL`},
	},
	"act5X": {
		{ID: 2199, Key: "act5X", Value: `SEARCH FOR BAAL`},
	},
	"am1": {
		{ID: 1255, Key: "am1", Value: `Stag Bow`},
	},
	"am2": {
		{ID: 1256, Key: "am2", Value: `Reflex Bow`},
	},
	"am3": {
		{ID: 1258, Key: "am3", Value: `Maiden Spear`},
	},
	"am4": {
		{ID: 1259, Key: "am4", Value: `Maiden Pike`},
	},
	"am5": {
		{ID: 1260, Key: "am5", Value: `Maiden Javelin`},
	},
	"am6": {
		{ID: 1261, Key: "am6", Value: `Ashwood Bow`},
	},
	"am7": {
		{ID: 1262, Key: "am7", Value: `Ceremonial Bow`},
	},
	"am8": {
		{ID: 1263, Key: "am8", Value: `Ceremonial Spear`},
	},
	"am9": {
		{ID: 1264, Key: "am9", Value: `Ceremonial Pike`},
	},
	"ama": {
		{ID: 1302, Key: "ama", Value: `Ceremonial Javelin`},
	},
	"amb": {
		{ID: 1303, Key: "amb", Value: `Matriarchal Bow`},
	},
	"amc": {
		{ID: 1304, Key: "amc", Value: `Grand Matron Bow`},
	},
	"amd": {
		{ID: 1305, Key: "amd", Value: `Matriarchal Spear`},
	},
	"ame": {
		{ID: 1306, Key: "ame", Value: `Matriarchal Pike`},
	},
	"amf": {
		{ID: 1307, Key: "amf", Value: `Matriarchal Javelin`},
	},
	"ancientsaltar": {
		{ID: 424, Key: "ancientsaltar", Value: `Altar of the Heavens`},
	},
	"animated skulland rockpile": {
		{ID: 382, Key: "animated skulland rockpile", Value: `Skulls and Rocks`},
	},
	"animatedskullsandrocks": {
		{ID: 343, Key: "animatedskullsandrocks", Value: `Skulls and Rocks`},
	},
	"as1": {
		{ID: 1350, Key: "as1", Value: `Wraps`},
	},
	"as2": {
		{ID: 1351, Key: "as2", Value: `Knuckles`},
	},
	"as3": {
		{ID: 1352, Key: "as3", Value: `Slashers`},
	},
	"as4": {
		{ID: 1353, Key: "as4", Value: `Splay`},
	},
	"as5": {
		{ID: 1354, Key: "as5", Value: `Hook`},
	},
	"as6": {
		{ID: 1356, Key: "as6", Value: `Shank`},
	},
	"as7": {
		{ID: 1357, Key: "as7", Value: `Claws`},
	},
	"axf": {
		{ID: 1483, Key: "axf", Value: `Hatchet Hands`},
	},
	"ba1": {
		{ID: 1318, Key: "ba1", Value: `Jawbone Cap`},
	},
	"ba2": {
		{ID: 1319, Key: "ba2", Value: `Fanged Helm`},
	},
	"ba3": {
		{ID: 1320, Key: "ba3", Value: `Horned Helm`},
	},
	"ba4": {
		{ID: 1321, Key: "ba4", Value: `Assault Helmet`},
	},
	"ba5": {
		{ID: 1322, Key: "ba5", Value: `Avenger Guard`},
	},
	"ba6": {
		{ID: 1323, Key: "ba6", Value: `Jawbone Visor`},
	},
	"ba7": {
		{ID: 1324, Key: "ba7", Value: `Lion Helm`},
	},
	"ba8": {
		{ID: 1325, Key: "ba8", Value: `Rage Mask`},
	},
	"ba9": {
		{ID: 1326, Key: "ba9", Value: `Savage Helmet`},
	},
	"baa": {
		{ID: 1366, Key: "baa", Value: `Slayer Guard`},
	},
	"bab": {
		{ID: 1367, Key: "bab", Value: `Carnage Helm`},
	},
	"bac": {
		{ID: 1368, Key: "bac", Value: `Fury Visor`},
	},
	"bad": {
		{ID: 1370, Key: "bad", Value: `Destroyer Helm`},
	},
	"bae": {
		{ID: 1371, Key: "bae", Value: `Conqueror Crown`},
	},
	"baf": {
		{ID: 1372, Key: "baf", Value: `Guardian Crown`},
	},
	"banner 1": {
		{ID: 998, Key: "banner 1", Value: `Banner`},
	},
	"banner 2": {
		{ID: 996, Key: "banner 2", Value: `Banner`},
	},
	"banner1": {
		{ID: 345, Key: "banner1", Value: `Banner`},
	},
	"banner2": {
		{ID: 346, Key: "banner2", Value: `Banner`},
	},
	"barrel wilderness": {
		{ID: 307, Key: "barrel wilderness", Value: `Barrel`},
	},
	"brz": {
		{ID: 1663, Key: "brz", Value: `Brain`},
	},
	"btl": {
		{ID: 1681, Key: "btl", Value: `Blade Talons`},
	},
	"burialchestL": {
		{ID: 309, Key: "burialchestL", Value: `Burial Chest`},
	},
	"burialchestR": {
		{ID: 310, Key: "burialchestR", Value: `Burial Chest`},
	},
	"burningbodies": {
		{ID: 314, Key: "burningbodies", Value: `Pyre of Flesh`},
	},
	"burningpit": {
		{ID: 322, Key: "burningpit", Value: `Burning Pit`},
	},
	"cagedwussie1": {
		{ID: 367, Key: "cagedwussie1", Value: `Cage`},
	},
	"candles": {
		{ID: 764, Key: "candles", Value: `Candles`},
	},
	"ces": {
		{ID: 1704, Key: "ces", Value: `Cestus`},
	},
	"chan": {
		{ID: 334, Key: "chan", Value: `Chandelier`},
	},
	"chestR": {
		{ID: 1243, Key: "chestR", Value: `Chest`},
	},
	"chestr": {
		{ID: 1278, Key: "chestr", Value: `Chest`},
	},
	"ci0": {
		{ID: 1701, Key: "ci0", Value: `Circlet`},
	},
	"ci1": {
		{ID: 1702, Key: "ci1", Value: `Coronet`},
	},
	"ci2": {
		{ID: 1703, Key: "ci2", Value: `Tiara`},
	},
	"ci3": {
		{ID: 1707, Key: "ci3", Value: `Diadem`},
	},
	"clientsmoke": {
		{ID: 365, Key: "clientsmoke", Value: `Smoke`},
	},
	"clw": {
		{ID: 1820, Key: "clw", Value: `Claws`},
	},
	"cm1": {
		{ID: 1766, Key: "cm1", Value: `Small Charm`},
	},
	"cm2": {
		{ID: 1767, Key: "cm2", Value: `Large Charm`},
	},
	"cm3": {
		{ID: 1768, Key: "cm3", Value: `Grand Charm`},
	},
	"deadbarbarian": {
		{ID: 495, Key: "deadbarbarian", Value: `Dead Barbarian`},
	},
	"deadbarbarian18": {
		{ID: 366, Key: "deadbarbarian18", Value: `Dead Barbarian`},
	},
	"deadperson": {
		{ID: 531, Key: "deadperson", Value: `Corpse`},
	},
	"deadperson2": {
		{ID: 417, Key: "deadperson2", Value: `Corpse`},
	},
	"deathpole": {
		{ID: 493, Key: "deathpole", Value: `Death Pole`},
	},
	"debris": {
		{ID: 463, Key: "debris", Value: `Debris`},
	},
	"dr1": {
		{ID: 2102, Key: "dr1", Value: `Wolf Head`},
	},
	"dr2": {
		{ID: 2103, Key: "dr2", Value: `Hawk Helm`},
	},
	"dr3": {
		{ID: 2104, Key: "dr3", Value: `Antlers`},
	},
	"dr4": {
		{ID: 2105, Key: "dr4", Value: `Falcon Mask`},
	},
	"dr5": {
		{ID: 2106, Key: "dr5", Value: `Spirit Mask`},
	},
	"dr6": {
		{ID: 2107, Key: "dr6", Value: `Alpha Helm`},
	},
	"dr7": {
		{ID: 2108, Key: "dr7", Value: `Griffon Headdress`},
	},
	"dr8": {
		{ID: 2109, Key: "dr8", Value: `Hunter's Guise`},
	},
	"dr9": {
		{ID: 2110, Key: "dr9", Value: `Sacred Feathers`},
	},
	"dra": {
		{ID: 2150, Key: "dra", Value: `Totemic Mask`},
	},
	"drb": {
		{ID: 2151, Key: "drb", Value: `Blood Spirit`},
	},
	"drc": {
		{ID: 2152, Key: "drc", Value: `Sun Spirit`},
	},
	"drd": {
		{ID: 2153, Key: "drd", Value: `Earth Spirit`},
	},
	"dre": {
		{ID: 2154, Key: "dre", Value: `Sky Spirit`},
	},
	"drf": {
		{ID: 2156, Key: "drf", Value: `Dream Spirit`},
	},
	"earmorstandL": {
		{ID: 443, Key: "earmorstandL", Value: `Armor Stand`},
	},
	"earmorstandR": {
		{ID: 444, Key: "earmorstandR", Value: `Armor Stand`},
	},
	"ecfra": {
		{ID: 385, Key: "ecfra", Value: `Fire`},
	},
	"eflg": {
		{ID: 327, Key: "eflg", Value: `Town Flag`},
	},
	"etorch1": {
		{ID: 1076, Key: "etorch1", Value: `Torch`},
	},
	"etorch2": {
		{ID: 1072, Key: "etorch2", Value: `Torch`},
	},
	"ettr": {
		{ID: 384, Key: "ettr", Value: `Torch`},
	},
	"evilurn": {
		{ID: 370, Key: "evilurn", Value: `Evil Urn`},
	},
	"eweaponrackL": {
		{ID: 1226, Key: "eweaponrackL", Value: `Weapon Rack`},
	},
	"eweaponrackR": {
		{ID: 431, Key: "eweaponrackR", Value: `Weapon Rack`},
	},
	"explodingbarrel": {
		{ID: 386, Key: "explodingbarrel", Value: `Barrel`},
	},
	"explodingchest": {
		{ID: 381, Key: "explodingchest", Value: `Chest`},
	},
	"eyz": {
		{ID: 2543, Key: "eyz", Value: `Eye`},
	},
	"fana": {
		{ID: 2024, Key: "fana", Value: `Frozen Anya`},
	},
	"flag widlerness": {
		{ID: 785, Key: "flag widlerness", Value: `Flag`},
	},
	"flag wilderness": {
		{ID: 655, Key: "flag wilderness", Value: `Flag`},
	},
	"flg": {
		{ID: 2572, Key: "flg", Value: `Flag`},
	},
	"fng": {
		{ID: 2604, Key: "fng", Value: `Fang`},
	},
	"gate": {
		{ID: 462, Key: "gate", Value: `Main Gate`},
	},
	"groundtomb": {
		{ID: 853, Key: "groundtomb", Value: `Tomb`},
	},
	"groundtombL": {
		{ID: 401, Key: "groundtombL", Value: `Tomb`},
	},
	"healthshrine": {
		{ID: 1003, Key: "healthshrine", Value: `Health Shrine`},
	},
	"hellgate": {
		{ID: 894, Key: "hellgate", Value: `Hell Gate`},
	},
	"hiddenstash": {
		{ID: 425, Key: "hiddenstash", Value: `Hidden Stash`},
	},
	"hirechat1": {
		{ID: 612, Key: "hirechat1", Value: `Thanks.`},
	},
	"hirechat2": {
		{ID: 614, Key: "hirechat2", Value: `Thank you.`},
	},
	"hirechat3": {
		{ID: 620, Key: "hirechat3", Value: `I needed that.`},
	},
	"hiredanger1": {
		{ID: 621, Key: "hiredanger1", Value: `I sense danger.`},
	},
	"hiredanger2": {
		{ID: 635, Key: "hiredanger2", Value: `I hate these vermin.`},
	},
	"hiredanger3": {
		{ID: 638, Key: "hiredanger3", Value: `I have a bad feeling about this.`},
	},
	"hiredanger4": {
		{ID: 641, Key: "hiredanger4", Value: `Beware!`},
	},
	"hiredanger5": {
		{ID: 646, Key: "hiredanger5", Value: `I detest spiders.`},
	},
	"hiredanger6": {
		{ID: 649, Key: "hiredanger6", Value: `Eek, snakes!`},
	},
	"hiredismiss": {
		{ID: 1025, Key: "hiredismiss", Value: `Dismiss`},
	},
	"hiredismisshire": {
		{ID: 601, Key: "hiredismisshire", Value: `Dismiss Hireling`},
	},
	"hirefeelstronger2": {
		{ID: 650, Key: "hirefeelstronger2", Value: `I am more experienced.`},
	},
	"hiregreets1": {
		{ID: 658, Key: "hiregreets1", Value: `Good morning.`},
	},
	"hiregreets2": {
		{ID: 691, Key: "hiregreets2", Value: `Good afternoon.`},
	},
	"hiregreets3": {
		{ID: 692, Key: "hiregreets3", Value: `Good evening.`},
	},
	"hiregreets4": {
		{ID: 693, Key: "hiregreets4", Value: `Hello.`},
	},
	"hirehelp1": {
		{ID: 1088, Key: "hirehelp1", Value: `I am hurt!`},
	},
	"hirehelp2": {
		{ID: 1089, Key: "hirehelp2", Value: `Help!`},
	},
	"hirehelp3": {
		{ID: 1107, Key: "hirehelp3", Value: `I am dying.`},
	},
	"hirehelp4": {
		{ID: 1113, Key: "hirehelp4", Value: `Help me!`},
	},
	"hireiconinfo1": {
		{ID: 597, Key: "hireiconinfo1", Value: `Drop Potion on Portrait to Heal`},
	},
	"hireiconinfo2": {
		{ID: 600, Key: "hireiconinfo2", Value: `Right-click to Open Inventory (%s)`},
	},
	"hirepraise1": {
		{ID: 686, Key: "hirepraise1", Value: `It is good to work for someone who cares.`},
	},
	"hirepraise2": {
		{ID: 690, Key: "hirepraise2", Value: `Good for you.`},
	},
	"hirerehire": {
		{ID: 602, Key: "hirerehire", Value: `Rehire`},
	},
	"hireresurrect": {
		{ID: 1027, Key: "hireresurrect", Value: `Resurrect`},
	},
	"hireresurrect2": {
		{ID: 605, Key: "hireresurrect2", Value: `Resurrect %s: %d`},
	},
	"hrb": {
		{ID: 357, Key: "hrb", Value: `Herb`},
	},
	"hrn": {
		{ID: 369, Key: "hrn", Value: `Horn`},
	},
	"hrt": {
		{ID: 375, Key: "hrt", Value: `Heart`},
	},
	"ice": {
		{ID: 376, Key: "ice", Value: `Keep it to thaw Anya
Malah's Potion`},
	},
	"icecave_torch1": {
		{ID: 479, Key: "icecave_torch1", Value: `Torch`},
	},
	"icecave_torch2": {
		{ID: 476, Key: "icecave_torch2", Value: `Torch`},
	},
	"icecavejar1": {
		{ID: 442, Key: "icecavejar1", Value: `Jar`},
	},
	"icecavejar2": {
		{ID: 436, Key: "icecavejar2", Value: `Jar`},
	},
	"icecavejar3": {
		{ID: 441, Key: "icecavejar3", Value: `Jar`},
	},
	"icecavejar4": {
		{ID: 439, Key: "icecavejar4", Value: `Jar`},
	},
	"icecaveshrine2": {
		{ID: 368, Key: "icecaveshrine2", Value: `Shrine`},
	},
	"jar": {
		{ID: 613, Key: "jar", Value: `Jar`},
	},
	"jar1": {
		{ID: 383, Key: "jar1", Value: `Jar`},
	},
	"jar2": {
		{ID: 335, Key: "jar2", Value: `Jar`},
	},
	"jar3": {
		{ID: 336, Key: "jar3", Value: `Jar`},
	},
	"jaw": {
		{ID: 618, Key: "jaw", Value: `Jawbone`},
	},
	"jew": {
		{ID: 683, Key: "jew", Value: `Jewel`},
	},
	"ktr": {
		{ID: 1173, Key: "ktr", Value: `Katar`},
	},
	"magic shrine2": {
		{ID: 995, Key: "magic shrine2", Value: `Magic Shrine`},
	},
	"mrbox": {
		{ID: 685, Key: "mrbox", Value: `Box`},
	},
	"mrjar": {
		{ID: 397, Key: "mrjar", Value: `Hidden Stash`},
	},
	"mrpole": {
		{ID: 494, Key: "mrpole", Value: `Pole`},
	},
	"ne1": {
		{ID: 1635, Key: "ne1", Value: `Preserved Head`},
	},
	"ne2": {
		{ID: 1636, Key: "ne2", Value: `Zombie Head`},
	},
	"ne3": {
		{ID: 1637, Key: "ne3", Value: `Unraveller Head`},
	},
	"ne4": {
		{ID: 1638, Key: "ne4", Value: `Gargoyle Head`},
	},
	"ne5": {
		{ID: 1640, Key: "ne5", Value: `Demon Head`},
	},
	"ne6": {
		{ID: 1641, Key: "ne6", Value: `Mummified Trophy`},
	},
	"ne7": {
		{ID: 1643, Key: "ne7", Value: `Fetish Trophy`},
	},
	"ne8": {
		{ID: 1644, Key: "ne8", Value: `Sexton Trophy`},
	},
	"ne9": {
		{ID: 1645, Key: "ne9", Value: `Cantor Trophy`},
	},
	"nea": {
		{ID: 1686, Key: "nea", Value: `Hierophant Trophy`},
	},
	"neb": {
		{ID: 1687, Key: "neb", Value: `Minion Skull`},
	},
	"nec": {
		{ID: 1691, Key: "nec", Value: `Hellspawn Skull`},
	},
	"ned": {
		{ID: 1692, Key: "ned", Value: `Overseer Skull`},
	},
	"nee": {
		{ID: 1693, Key: "nee", Value: `Succubus Skull`},
	},
	"nef": {
		{ID: 1694, Key: "nef", Value: `Bloodlord Skull`},
	},
	"ob1": {
		{ID: 1843, Key: "ob1", Value: `Eagle Orb`},
	},
	"ob2": {
		{ID: 1844, Key: "ob2", Value: `Sacred Globe`},
	},
	"ob3": {
		{ID: 1846, Key: "ob3", Value: `Smoked Sphere`},
	},
	"ob4": {
		{ID: 1847, Key: "ob4", Value: `Clasped Orb`},
	},
	"ob5": {
		{ID: 2023, Key: "ob5", Value: `Jared's Stone`},
	},
	"ob6": {
		{ID: 1848, Key: "ob6", Value: `Glowing Orb`},
	},
	"ob7": {
		{ID: 1849, Key: "ob7", Value: `Crystalline Globe`},
	},
	"ob8": {
		{ID: 1850, Key: "ob8", Value: `Cloudy Sphere`},
	},
	"ob9": {
		{ID: 1851, Key: "ob9", Value: `Sparkling Ball`},
	},
	"oba": {
		{ID: 1891, Key: "oba", Value: `Swirling Crystal`},
	},
	"obb": {
		{ID: 1892, Key: "obb", Value: `Heavenly Stone`},
	},
	"obc": {
		{ID: 1893, Key: "obc", Value: `Eldritch Orb`},
	},
	"obd": {
		{ID: 1894, Key: "obd", Value: `Demon Heart`},
	},
	"obe": {
		{ID: 1895, Key: "obe", Value: `Vortex Orb`},
	},
	"obf": {
		{ID: 1896, Key: "obf", Value: `Dimensional Shard`},
	},
	"object": {
		{ID: 389, Key: "object", Value: `Hidden Stash`},
	},
	"object1": {
		{ID: 351, Key: "object1", Value: `Hidden Stash`},
	},
	"object2": {
		{ID: 398, Key: "object2", Value: `Hidden Stash`},
	},
	"of Acceleration": {
		{ID: 1618, Key: "of Acceleration", Value: `of Acceleration`},
	},
	"of Ages": {
		{ID: 1508, Key: "of Ages", Value: `of Ages`},
	},
	"of Amianthus": {
		{ID: 37, Key: "of Amianthus", Value: `of Amianthus`},
	},
	"of Amplify Damage": {
		{ID: 429, Key: "of Amplify Damage", Value: `of Amplify Damage`},
	},
	"of Anima": {
		{ID: 958, Key: "of Anima", Value: `of Amicae`},
	},
	"of Anthrax": {
		{ID: 1414, Key: "of Anthrax", Value: `of Anthrax`},
	},
	"of Armageddon": {
		{ID: 2241, Key: "of Armageddon", Value: `of Armageddon`},
	},
	"of Atlus": {
		{ID: 210, Key: "of Atlus", Value: `of Atlas`},
	},
	"of Attraction": {
		{ID: 1721, Key: "of Attraction", Value: `of Attract`},
	},
	"of Avarice": {
		{ID: 2117, Key: "of Avarice", Value: `of Avarice`},
	},
	"of Battle Command": {
		{ID: 654, Key: "of Battle Command", Value: `of Battle Command`},
	},
	"of Battle Cry": {
		{ID: 173, Key: "of Battle Cry", Value: `of Battle Cry`},
	},
	"of Battle Orders": {
		{ID: 360, Key: "of Battle Orders", Value: `of Battle Orders`},
	},
	"of Blaze": {
		{ID: 886, Key: "of Blaze", Value: `of Blaze`},
	},
	"of Blazing": {
		{ID: 555, Key: "of Blazing", Value: `of Blazing`},
	},
	"of Blessed Hammers": {
		{ID: 1787, Key: "of Blessed Hammers", Value: `of Blessed Hammer`},
	},
	"of Bliss": {
		{ID: 14, Key: "of Bliss", Value: `of Bliss`},
	},
	"of Blizzard": {
		{ID: 378, Key: "of Blizzard", Value: `of Blizzard`},
	},
	"of Blizzards": {
		{ID: 611, Key: "of Blizzards", Value: `of Blizzard`},
	},
	"of Blood Golem Summoning": {
		{ID: 2754, Key: "of Blood Golem Summoning", Value: `of Blood Golem`},
	},
	"of Bone Armor": {
		{ID: 56, Key: "of Bone Armor", Value: `of Bone Armor`},
	},
	"of Bone Imprisonment": {
		{ID: 1967, Key: "of Bone Imprisonment", Value: `of Bone Prison`},
	},
	"of Bone Spears": {
		{ID: 1201, Key: "of Bone Spears", Value: `of Bone Spear`},
	},
	"of Bone Spirits": {
		{ID: 869, Key: "of Bone Spirits", Value: `of Bone Spirit`},
	},
	"of Bone Walls": {
		{ID: 38, Key: "of Bone Walls", Value: `of Bone Wall`},
	},
	"of Butchery": {
		{ID: 2066, Key: "of Butchery", Value: `of Butchery`},
	},
	"of Chain Lightning": {
		{ID: 1648, Key: "of Chain Lightning", Value: `of Chain Lightning`},
	},
	"of Charged Bolts": {
		{ID: 1330, Key: "of Charged Bolts", Value: `of Charged Bolt`},
	},
	"of Charged Shield": {
		{ID: 2434, Key: "of Charged Shield", Value: `of Charged Bolt`},
	},
	"of Charged Spear": {
		{ID: 969, Key: "of Charged Spear", Value: `of Charged Spear`},
	},
	"of Charged Strike": {
		{ID: 1809, Key: "of Charged Strike", Value: `of Charged Strike`},
	},
	"of Charging": {
		{ID: 2412, Key: "of Charging", Value: `of Charge`},
	},
	"of Chilling Armor": {
		{ID: 2243, Key: "of Chilling Armor", Value: `of Chilling Armor`},
	},
	"of Clay Golem Summoning": {
		{ID: 536, Key: "of Clay Golem Summoning", Value: `of Clay Golem`},
	},
	"of Cold Arrows": {
		{ID: 645, Key: "of Cold Arrows", Value: `of Cold Arrow`},
	},
	"of Concentration": {
		{ID: 1653, Key: "of Concentration", Value: `of Concentration`},
	},
	"of Confusion": {
		{ID: 1750, Key: "of Confusion", Value: `of Confusion`},
	},
	"of Conversion": {
		{ID: 1991, Key: "of Conversion", Value: `of Conversion`},
	},
	"of Coolness": {
		{ID: 2583, Key: "of Coolness", Value: `of Coolness`},
	},
	"of Corpse Explosions": {
		{ID: 1902, Key: "of Corpse Explosions", Value: `of Corpse Explosion`},
	},
	"of Cyclone Armor": {
		{ID: 981, Key: "of Cyclone Armor", Value: `of Cyclone Armor`},
	},
	"of Damage Amplification": {
		{ID: 1096, Key: "of Damage Amplification", Value: `of Amplify Damage`},
	},
	"of Daring": {
		{ID: 1222, Key: "of Daring", Value: `of Daring`},
	},
	"of Dawn": {
		{ID: 1236, Key: "of Dawn", Value: `of Dawn`},
	},
	"of Decrepification": {
		{ID: 1241, Key: "of Decrepification", Value: `of Decrepify`},
	},
	"of Desire": {
		{ID: 2287, Key: "of Desire", Value: `of Desire`},
	},
	"of Dim Vision": {
		{ID: 1009, Key: "of Dim Vision", Value: `of Dim Vision`},
	},
	"of Enchant": {
		{ID: 1460, Key: "of Enchant", Value: `of Enchant`},
	},
	"of Enchantment": {
		{ID: 1329, Key: "of Enchantment", Value: `of Enchantment`},
	},
	"of Energy Shield": {
		{ID: 1040, Key: "of Energy Shield", Value: `of Energy Shield`},
	},
	"of Enlightenment": {
		{ID: 1111, Key: "of Enlightenment", Value: `of Enlightenment`},
	},
	"of Ennui": {
		{ID: 556, Key: "of Ennui", Value: `of Ennui`},
	},
	"of Envy": {
		{ID: 198, Key: "of Envy", Value: `of Envy`},
	},
	"of Eruption": {
		{ID: 1417, Key: "of Eruption", Value: `of Fissure`},
	},
	"of Evisceration": {
		{ID: 598, Key: "of Evisceration", Value: `of Evisceration`},
	},
	"of Exploding Arrows": {
		{ID: 1144, Key: "of Exploding Arrows", Value: `of Exploding Arrow`},
	},
	"of Faith": {
		{ID: 44, Key: "of Faith", Value: `of Faith`},
	},
	"of Fast Repair": {
		{ID: 505, Key: "of Fast Repair", Value: `of Restoration`},
	},
	"of Fending": {
		{ID: 1722, Key: "of Fending", Value: `of Fending`},
	},
	"of Fervor": {
		{ID: 2327, Key: "of Fervor", Value: `of Fervor`},
	},
	"of Fire Arrows": {
		{ID: 840, Key: "of Fire Arrows", Value: `of Fire Arrow`},
	},
	"of Fire Ball": {
		{ID: 1489, Key: "of Fire Ball", Value: `of Fire Ball`},
	},
	"of Fire Balls": {
		{ID: 2436, Key: "of Fire Balls", Value: `of Fire Ball`},
	},
	"of Fire Bolts": {
		{ID: 455, Key: "of Fire Bolts", Value: `of Fire Bolt`},
	},
	"of Fire Golem Summoning": {
		{ID: 876, Key: "of Fire Golem Summoning", Value: `of Fire Golem`},
	},
	"of Fire Quenching": {
		{ID: 2364, Key: "of Fire Quenching", Value: `of Quenching`},
	},
	"of Fire Wall": {
		{ID: 1177, Key: "of Fire Wall", Value: `of Fire Wall`},
	},
	"of Fire Walls": {
		{ID: 248, Key: "of Fire Walls", Value: `of Fire Wall`},
	},
	"of Firebolts": {
		{ID: 563, Key: "of Firebolts", Value: `of Firebolts`},
	},
	"of Firestorms": {
		{ID: 2307, Key: "of Firestorms", Value: `of Firestorms`},
	},
	"of Fist of the Heavens": {
		{ID: 1132, Key: "of Fist of the Heavens", Value: `of Fist of the Heavens`},
	},
	"of Freedom": {
		{ID: 1556, Key: "of Freedom", Value: `of Freedom`},
	},
	"of Freezing Arrows": {
		{ID: 2026, Key: "of Freezing Arrows", Value: `of Freezing Arrow`},
	},
	"of Frigidity": {
		{ID: 827, Key: "of Frigidity", Value: `of Frigidity`},
	},
	"of Frost Novas": {
		{ID: 432, Key: "of Frost Novas", Value: `of Frost Nova`},
	},
	"of Frost Shield": {
		{ID: 2674, Key: "of Frost Shield", Value: `of Frost Shield`},
	},
	"of Frozen Armor": {
		{ID: 1405, Key: "of Frozen Armor", Value: `of Frozen Armor`},
	},
	"of Frozen Orb": {
		{ID: 565, Key: "of Frozen Orb", Value: `of Frozen Orb`},
	},
	"of Frozen Orbs": {
		{ID: 93, Key: "of Frozen Orbs", Value: `of Frozen Orb`},
	},
	"of Glacial Spike": {
		{ID: 1724, Key: "of Glacial Spike", Value: `of Glacial Spike`},
	},
	"of Glacial Spikes": {
		{ID: 2329, Key: "of Glacial Spikes", Value: `of Glacial Spike`},
	},
	"of Good Luck": {
		{ID: 2242, Key: "of Good Luck", Value: `of Good Luck`},
	},
	"of Grace": {
		{ID: 2417, Key: "of Grace", Value: `of Grace`},
	},
	"of Grace and Power": {
		{ID: 1413, Key: "of Grace and Power", Value: `of Grace and Power`},
	},
	"of Grim Ward": {
		{ID: 968, Key: "of Grim Ward", Value: `of Grim Ward`},
	},
	"of Grounding": {
		{ID: 1589, Key: "of Grounding", Value: `of Grounding`},
	},
	"of Guided Arrows": {
		{ID: 214, Key: "of Guided Arrows", Value: `of Guided Arrow`},
	},
	"of Holy Bolts": {
		{ID: 127, Key: "of Holy Bolts", Value: `of Holy Bolt`},
	},
	"of Holy Shield": {
		{ID: 1790, Key: "of Holy Shield", Value: `of Holy Shield`},
	},
	"of Honor": {
		{ID: 1440, Key: "of Honor", Value: `of Honor`},
	},
	"of Hope": {
		{ID: 1317, Key: "of Hope", Value: `of Hope`},
	},
	"of Howling": {
		{ID: 1698, Key: "of Howling", Value: `of Howl`},
	},
	"of Hurricane": {
		{ID: 2127, Key: "of Hurricane", Value: `of Hurricane`},
	},
	"of Hydra Shield": {
		{ID: 1634, Key: "of Hydra Shield", Value: `of Hydra Shield`},
	},
	"of Hydras": {
		{ID: 2139, Key: "of Hydras", Value: `of Hydra`},
	},
	"of Ice Arrows": {
		{ID: 2165, Key: "of Ice Arrows", Value: `of Ice Arrow`},
	},
	"of Ice Blast": {
		{ID: 1656, Key: "of Ice Blast", Value: `of Ice Blast`},
	},
	"of Ice Blasts": {
		{ID: 1242, Key: "of Ice Blasts", Value: `of Ice Blast`},
	},
	"of Ice Bolts": {
		{ID: 631, Key: "of Ice Bolts", Value: `of Ice Bolt`},
	},
	"of Icebolt": {
		{ID: 1433, Key: "of Icebolt", Value: `of Icebolt`},
	},
	"of Immolating Arrows": {
		{ID: 582, Key: "of Immolating Arrows", Value: `of Immolating Arrow`},
	},
	"of Impaling Spear": {
		{ID: 2306, Key: "of Impaling Spear", Value: `of Impaling Spear`},
	},
	"of Impaling Strike": {
		{ID: 1085, Key: "of Impaling Strike", Value: `of Impaling Strike`},
	},
	"of Incineration": {
		{ID: 952, Key: "of Incineration", Value: `of Incineration`},
	},
	"of Incombustibility": {
		{ID: 2065, Key: "of Incombustibility", Value: `of Inflammability`},
	},
	"of Inertia": {
		{ID: 767, Key: "of Inertia", Value: `of Inertia`},
	},
	"of Inner Sight": {
		{ID: 2262, Key: "of Inner Sight", Value: `of Inner Sight`},
	},
	"of Insulation": {
		{ID: 1869, Key: "of Insulation", Value: `of Insulation`},
	},
	"of Ire": {
		{ID: 1346, Key: "of Ire", Value: `of Ire`},
	},
	"of Iron Golem Creation": {
		{ID: 1284, Key: "of Iron Golem Creation", Value: `of Iron Golem`},
	},
	"of Iron Maiden": {
		{ID: 1049, Key: "of Iron Maiden", Value: `of Iron Maiden`},
	},
	"of Item Finding": {
		{ID: 1594, Key: "of Item Finding", Value: `of Find Item`},
	},
	"of Jabbing": {
		{ID: 324, Key: "of Jabbing", Value: `of Jab`},
	},
	"of Joy": {
		{ID: 1588, Key: "of Joy", Value: `of Joy`},
	},
	"of Joyfulness": {
		{ID: 2428, Key: "of Joyfulness", Value: `of Joyfulness`},
	},
	"of Knowledge": {
		{ID: 2198, Key: "of Knowledge", Value: `of Knowledge`},
	},
	"of Life Everlasting": {
		{ID: 1718, Key: "of Life Everlasting", Value: `of Life Everlasting`},
	},
	"of Life Tap": {
		{ID: 1726, Key: "of Life Tap", Value: `of Life Tap`},
	},
	"of Lightning Fury": {
		{ID: 551, Key: "of Lightning Fury", Value: `of Lightning Fury`},
	},
	"of Lightning Javelin": {
		{ID: 521, Key: "of Lightning Javelin", Value: `of Lightning Javelin`},
	},
	"of Lightning Spear": {
		{ID: 2308, Key: "of Lightning Spear", Value: `of Lightning Spear`},
	},
	"of Lightning Strike": {
		{ID: 1518, Key: "of Lightning Strike", Value: `of Lightning Strike`},
	},
	"of Lower Resistance": {
		{ID: 2312, Key: "of Lower Resistance", Value: `of Lower Resistance`},
	},
	"of Luck": {
		{ID: 2281, Key: "of Luck", Value: `of Luck`},
	},
	"of Magic Arrows": {
		{ID: 162, Key: "of Magic Arrows", Value: `of Magic Arrow`},
	},
	"of Malice": {
		{ID: 1287, Key: "of Malice", Value: `of Malice`},
	},
	"of Memory": {
		{ID: 2292, Key: "of Memory", Value: `of Memory`},
	},
	"of Meteor": {
		{ID: 1697, Key: "of Meteor", Value: `of Meteor`},
	},
	"of Meteors": {
		{ID: 1131, Key: "of Meteors", Value: `of Meteor`},
	},
	"of Molten Boulders": {
		{ID: 566, Key: "of Molten Boulders", Value: `of Molten Boulder`},
	},
	"of Multiple Shot": {
		{ID: 659, Key: "of Multiple Shot", Value: `of Multiple Shot`},
	},
	"of Nirvana": {
		{ID: 1120, Key: "of Nirvana", Value: `of Nirvana`},
	},
	"of Nova": {
		{ID: 679, Key: "of Nova", Value: `of Nova`},
	},
	"of Nova Shield": {
		{ID: 2092, Key: "of Nova Shield", Value: `of Nova Shield`},
	},
	"of Novas": {
		{ID: 1968, Key: "of Novas", Value: `of Nova`},
	},
	"of Passion": {
		{ID: 2810, Key: "of Passion", Value: `of Passion`},
	},
	"of Plague Jab": {
		{ID: 2030, Key: "of Plague Jab", Value: `of Plague Jab`},
	},
	"of Plague Javelin": {
		{ID: 1620, Key: "of Plague Javelin", Value: `of Plague Javelin`},
	},
	"of Poison Dagger": {
		{ID: 1033, Key: "of Poison Dagger", Value: `of Poison Dagger`},
	},
	"of Poison Explosion": {
		{ID: 585, Key: "of Poison Explosion", Value: `of Poison Explosion`},
	},
	"of Poison Jab": {
		{ID: 994, Key: "of Poison Jab", Value: `of Poison Jab`},
	},
	"of Poison Javelin": {
		{ID: 132, Key: "of Poison Javelin", Value: `of Poison Javelin`},
	},
	"of Poison Novas": {
		{ID: 1654, Key: "of Poison Novas", Value: `of Poison Nova`},
	},
	"of Potion Finding": {
		{ID: 714, Key: "of Potion Finding", Value: `of Find Potion`},
	},
	"of Power": {
		{ID: 207, Key: "of Power", Value: `of Power`},
	},
	"of Power Spear": {
		{ID: 157, Key: "of Power Spear", Value: `of Power Spear`},
	},
	"of Power Strike": {
		{ID: 1176, Key: "of Power Strike", Value: `of Power Strike`},
	},
	"of Propogation": {
		{ID: 1134, Key: "of Propogation", Value: `of Propogation`},
	},
	"of Prosperity": {
		{ID: 191, Key: "of Prosperity", Value: `of Prosperity`},
	},
	"of Raise Skeletal Mages": {
		{ID: 2093, Key: "of Raise Skeletal Mages", Value: `of Skeletal Mages`},
	},
	"of Raise Skeletons": {
		{ID: 783, Key: "of Raise Skeletons", Value: `of Raise Skeleton`},
	},
	"of Razors": {
		{ID: 742, Key: "of Razors", Value: `of Razors`},
	},
	"of Replenishing": {
		{ID: 537, Key: "of Replenishing", Value: `of Replenishing`},
	},
	"of Resistance": {
		{ID: 153, Key: "of Resistance", Value: `of Resistance`},
	},
	"of Revivification": {
		{ID: 1104, Key: "of Revivification", Value: `of Revivification`},
	},
	"of Sacrifice": {
		{ID: 639, Key: "of Sacrifice", Value: `of Sacrifice`},
	},
	"of Self-Repair": {
		{ID: 1343, Key: "of Self-Repair", Value: `of Self-Repair`},
	},
	"of Shiver Armor": {
		{ID: 140, Key: "of Shiver Armor", Value: `of Shiver Armor`},
	},
	"of Shouting": {
		{ID: 982, Key: "of Shouting", Value: `of Shout`},
	},
	"of Slow Missiles": {
		{ID: 1994, Key: "of Slow Missiles", Value: `of Slow Missile`},
	},
	"of Spirit": {
		{ID: 812, Key: "of Spirit", Value: `of Spirit`},
	},
	"of Static Field": {
		{ID: 1451, Key: "of Static Field", Value: `of Static Field`},
	},
	"of Stoicism": {
		{ID: 148, Key: "of Stoicism", Value: `of Stoicism`},
	},
	"of Storms": {
		{ID: 1649, Key: "of Storms", Value: `of Storms`},
	},
	"of Stunning": {
		{ID: 2381, Key: "of Stunning", Value: `of Stun`},
	},
	"of Substinence": {
		{ID: 1311, Key: "of Substinence", Value: `of Sustenance`},
	},
	"of Sunlight": {
		{ID: 752, Key: "of Sunlight", Value: `of Sunlight`},
	},
	"of Swords": {
		{ID: 525, Key: "of Swords", Value: `of Swords`},
	},
	"of Taunting": {
		{ID: 1087, Key: "of Taunting", Value: `of Taunt`},
	},
	"of Teeth": {
		{ID: 590, Key: "of Teeth", Value: `of Teeth`},
	},
	"of Telekinesis": {
		{ID: 1916, Key: "of Telekinesis", Value: `of Telekinesis`},
	},
	"of Teleport Shield": {
		{ID: 1554, Key: "of Teleport Shield", Value: `of Teleport Shield`},
	},
	"of Teleportation": {
		{ID: 2678, Key: "of Teleportation", Value: `of Teleportation`},
	},
	"of Terror": {
		{ID: 2063, Key: "of Terror", Value: `of Terror`},
	},
	"of Thawing": {
		{ID: 1127, Key: "of Thawing", Value: `of Thawing`},
	},
	"of Thunder Storm": {
		{ID: 2569, Key: "of Thunder Storm", Value: `of Thunder Storm`},
	},
	"of Tornado": {
		{ID: 2631, Key: "of Tornado", Value: `of Tornado`},
	},
	"of Transcendence": {
		{ID: 2069, Key: "of Transcendence", Value: `of Transcendence`},
	},
	"of Traveling": {
		{ID: 1587, Key: "of Traveling", Value: `of Transportation`},
	},
	"of Truth": {
		{ID: 1568, Key: "of Truth", Value: `of Truth`},
	},
	"of Twister": {
		{ID: 1591, Key: "of Twister", Value: `of Twister`},
	},
	"of Vengeance": {
		{ID: 91, Key: "of Vengeance", Value: `of Vengeance`},
	},
	"of Virility": {
		{ID: 2008, Key: "of Virility", Value: `of Virility`},
	},
	"of Vita": {
		{ID: 934, Key: "of Vita", Value: `of Vita`},
	},
	"of Volcano": {
		{ID: 440, Key: "of Volcano", Value: `of Volcano`},
	},
	"of War Cry": {
		{ID: 1130, Key: "of War Cry", Value: `of War Cry`},
	},
	"of Warming": {
		{ID: 1046, Key: "of Warming", Value: `of Warming`},
	},
	"of Weaken": {
		{ID: 1496, Key: "of Weaken", Value: `of Weaken`},
	},
	"of Winter": {
		{ID: 632, Key: "of Winter", Value: `of Winter`},
	},
	"of Wrath": {
		{ID: 579, Key: "of Wrath", Value: `of Wrath`},
	},
	"of Zeal": {
		{ID: 1745, Key: "of Zeal", Value: `of Zeal`},
	},
	"of the Cobra": {
		{ID: 1906, Key: "of the Cobra", Value: `of the Cobra`},
	},
	"of the Colossus": {
		{ID: 2067, Key: "of the Colossus", Value: `of the Gargantuan`},
	},
	"of the Colossus1": {
		{ID: 2261, Key: "of the Colossus1", Value: `of the Colossus`},
	},
	"of the Dunes": {
		{ID: 2508, Key: "of the Dunes", Value: `of the Dunes`},
	},
	"of the Dynamo": {
		{ID: 2494, Key: "of the Dynamo", Value: `of the Dynamo`},
	},
	"of the Efreeti": {
		{ID: 67, Key: "of the Efreeti", Value: `of the Efreeti`},
	},
	"of the Elements": {
		{ID: 1221, Key: "of the Elements", Value: `of the Elements`},
	},
	"of the Elephant": {
		{ID: 457, Key: "of the Elephant", Value: `of the Elephant`},
	},
	"of the Icicle": {
		{ID: 2045, Key: "of the Icicle", Value: `of the Icicle`},
	},
	"of the Kraken": {
		{ID: 342, Key: "of the Kraken", Value: `of the Centaur`},
	},
	"of the Kraken1": {
		{ID: 1870, Key: "of the Kraken1", Value: `of the Kraken`},
	},
	"of the Lamprey": {
		{ID: 292, Key: "of the Lamprey", Value: `of the Lamprey`},
	},
	"of the Phoenix": {
		{ID: 727, Key: "of the Phoenix", Value: `of the Phoenix`},
	},
	"of the Plague": {
		{ID: 2033, Key: "of the Plague", Value: `of the Plague`},
	},
	"of the Sirocco": {
		{ID: 925, Key: "of the Sirocco", Value: `of the Scirocco`},
	},
	"of the Squid": {
		{ID: 1097, Key: "of the Squid", Value: `of the Squid`},
	},
	"of the Whale": {
		{ID: 1871, Key: "of the Whale", Value: `of the Whale`},
	},
	"of the Yeti": {
		{ID: 2432, Key: "of the Yeti", Value: `of the Yeti`},
	},
	"pa1": {
		{ID: 2083, Key: "pa1", Value: `Targe`},
	},
	"pa2": {
		{ID: 2084, Key: "pa2", Value: `Rondache`},
	},
	"pa3": {
		{ID: 2085, Key: "pa3", Value: `Heraldic Shield`},
	},
	"pa4": {
		{ID: 2086, Key: "pa4", Value: `Aerin Shield`},
	},
	"pa5": {
		{ID: 2087, Key: "pa5", Value: `Crown Shield`},
	},
	"pa6": {
		{ID: 2088, Key: "pa6", Value: `Akaran Targe`},
	},
	"pa7": {
		{ID: 2089, Key: "pa7", Value: `Akaran Rondache`},
	},
	"pa8": {
		{ID: 2090, Key: "pa8", Value: `Protector Shield`},
	},
	"pa9": {
		{ID: 2091, Key: "pa9", Value: `Gilded Shield`},
	},
	"paa": {
		{ID: 2131, Key: "paa", Value: `Royal Shield`},
	},
	"pab": {
		{ID: 2132, Key: "pab", Value: `Sacred Targe`},
	},
	"pac": {
		{ID: 2133, Key: "pac", Value: `Sacred Rondache`},
	},
	"pad": {
		{ID: 2185, Key: "pad", Value: `Kurast Shield`},
	},
	"pae": {
		{ID: 2135, Key: "pae", Value: `Zakarum Shield`},
	},
	"paf": {
		{ID: 2136, Key: "paf", Value: `Vortex Shield`},
	},
	"pene": {
		{ID: 350, Key: "pene", Value: `Stash`},
	},
	"pileofskullsandrocks": {
		{ID: 677, Key: "pileofskullsandrocks", Value: `Skulls and Rocks`},
	},
	"pole": {
		{ID: 339, Key: "pole", Value: `Pole`},
	},
	"ptox": {
		{ID: 392, Key: "ptox", Value: `Torch Pit`},
	},
	"pyox": {
		{ID: 1124, Key: "pyox", Value: `Fire Pit`},
	},
	"qll": {
		{ID: 2574, Key: "qll", Value: `Quill`},
	},
	"qsta5q11": {
		{ID: 623, Key: "qsta5q11", Value: `Stop the Siege by killing Shenk the Overseer in the Bloody Foothills.`},
	},
	"qsta5q12": {
		{ID: 629, Key: "qsta5q12", Value: `Kill Shenk the Overseer.`},
	},
	"qsta5q13": {
		{ID: 630, Key: "qsta5q13", Value: `Go talk to Larzuk for your reward.`},
	},
	"qsta5q14": {
		{ID: 626, Key: "qsta5q14", Value: `Larzuk will add sockets to the item of your choice.`},
	},
	"qstsa5q1": {
		{ID: 445, Key: "qstsa5q1", Value: `Siege on Harrogath`},
	},
	"qstsa5q2": {
		{ID: 446, Key: "qstsa5q2", Value: `Rescue on Mount Arreat`},
	},
	"qstsa5q21": {
		{ID: 447, Key: "qstsa5q21", Value: `Find the Soldiers in the Frigid Highlands.`},
	},
	"qstsa5q21a": {
		{ID: 2186, Key: "qstsa5q21a", Value: `Free the soldiers from their prison and lead them back to town.`},
	},
	"qstsa5q22": {
		{ID: 450, Key: "qstsa5q22", Value: `Rescue %d more Soldiers in the Frigid Highlands.`},
	},
	"qstsa5q23": {
		{ID: 464, Key: "qstsa5q23", Value: `Return to Qual-Kehk for your reward.`},
	},
	"qstsa5q24": {
		{ID: 466, Key: "qstsa5q24", Value: `Apply the Runes to a Socketed item in this order:`},
	},
	"qstsa5q3": {
		{ID: 471, Key: "qstsa5q3", Value: `Prison of Ice`},
	},
	"qstsa5q31": {
		{ID: 472, Key: "qstsa5q31", Value: `Look for Anya under the Crystalline Passage by the Frozen River.`},
	},
	"qstsa5q31a": {
		{ID: 1455, Key: "qstsa5q31a", Value: `Rescue Anya.`},
	},
	"qstsa5q32": {
		{ID: 474, Key: "qstsa5q32", Value: `Talk to Malah.`},
	},
	"qstsa5q33": {
		{ID: 480, Key: "qstsa5q33", Value: `Use the potion Malah gave you to thaw Anya.`},
	},
	"qstsa5q34": {
		{ID: 484, Key: "qstsa5q34", Value: `Talk to Malah for your reward.`},
	},
	"qstsa5q35": {
		{ID: 485, Key: "qstsa5q35", Value: `Talk to Anya.`},
	},
	"qstsa5q4": {
		{ID: 487, Key: "qstsa5q4", Value: `Betrayal of Harrogath`},
	},
	"qstsa5q41": {
		{ID: 488, Key: "qstsa5q41", Value: `Take Anya's portal to Nihlathak's Temple.`},
	},
	"qstsa5q42": {
		{ID: 489, Key: "qstsa5q42", Value: `Kill Nihlathak.`},
	},
	"qstsa5q42a": {
		{ID: 1498, Key: "qstsa5q42a", Value: `Find Nihlathak in the Halls of Vaught.`},
	},
	"qstsa5q43": {
		{ID: 492, Key: "qstsa5q43", Value: `Talk to Anya before you continue through the Crystalline Passage, past the Glacial Trail, to proceed up Mount Arreat to the Summit.`},
	},
	"qstsa5q43a": {
		{ID: 1482, Key: "qstsa5q43a", Value: `Anya will personalize an item for you.`},
	},
	"qstsa5q5": {
		{ID: 496, Key: "qstsa5q5", Value: `Rite of Passage`},
	},
	"qstsa5q51": {
		{ID: 498, Key: "qstsa5q51", Value: `Travel through the Ancient's Way to find the Ancients at the Arreat Summit.`},
	},
	"qstsa5q52": {
		{ID: 499, Key: "qstsa5q52", Value: `Consult with the Ancients by activating the Altar of the Heavens.`},
	},
	"qstsa5q53": {
		{ID: 500, Key: "qstsa5q53", Value: `Defeat all three Ancients without leaving Mount Arreat.`},
	},
	"qstsa5q6": {
		{ID: 501, Key: "qstsa5q6", Value: `Eve of Destruction`},
	},
	"qstsa5q61": {
		{ID: 502, Key: "qstsa5q61", Value: `Consult with the Ancients.`},
	},
	"qstsa5q61a": {
		{ID: 2029, Key: "qstsa5q61a", Value: `Find Baal's Throne Room.`},
	},
	"qstsa5q62": {
		{ID: 503, Key: "qstsa5q62", Value: `Kill Baal.`},
	},
	"qstsa5q62a": {
		{ID: 1757, Key: "qstsa5q62a", Value: `Kill Baal in the Worldstone Chamber before he corrupts it.`},
	},
	"qstsa5q62b": {
		{ID: 1758, Key: "qstsa5q62b", Value: `Kill Baal's Minions.`},
	},
	"qstsa5q63": {
		{ID: 516, Key: "qstsa5q63", Value: `Talk to Tyreal.`},
	},
	"qstsa5q64": {
		{ID: 517, Key: "qstsa5q64", Value: `Take Tyreal's Portal to Safety.`},
	},
	"r01": {
		{ID: 1823, Key: "r01", Value: `El Rune`},
	},
	"r01L": {
		{ID: 863, Key: "r01L", Value: `El`},
	},
	"r02": {
		{ID: 1812, Key: "r02", Value: `Eld Rune`},
	},
	"r02L": {
		{ID: 879, Key: "r02L", Value: `Eld`},
	},
	"r03": {
		{ID: 1813, Key: "r03", Value: `Tir Rune`},
	},
	"r03L": {
		{ID: 895, Key: "r03L", Value: `Tir`},
	},
	"r04": {
		{ID: 1814, Key: "r04", Value: `Nef Rune`},
	},
	"r04L": {
		{ID: 911, Key: "r04L", Value: `Nef`},
	},
	"r05": {
		{ID: 1815, Key: "r05", Value: `Eth Rune`},
	},
	"r05L": {
		{ID: 927, Key: "r05L", Value: `Eth`},
	},
	"r06": {
		{ID: 1822, Key: "r06", Value: `Ith Rune`},
	},
	"r06L": {
		{ID: 949, Key: "r06L", Value: `Ith`},
	},
	"r07": {
		{ID: 1817, Key: "r07", Value: `Tal Rune`},
	},
	"r07L": {
		{ID: 959, Key: "r07L", Value: `Tal`},
	},
	"r08": {
		{ID: 1818, Key: "r08", Value: `Ral Rune`},
	},
	"r08L": {
		{ID: 975, Key: "r08L", Value: `Ral`},
	},
	"r09": {
		{ID: 1819, Key: "r09", Value: `Ort Rune`},
	},
	"r09L": {
		{ID: 991, Key: "r09L", Value: `Ort`},
	},
	"r10": {
		{ID: 1839, Key: "r10", Value: `Thul Rune`},
	},
	"r10L": {
		{ID: 1103, Key: "r10L", Value: `Thul`},
	},
	"r11": {
		{ID: 1827, Key: "r11", Value: `Amn Rune`},
	},
	"r11L": {
		{ID: 1119, Key: "r11L", Value: `Amn`},
	},
	"r12": {
		{ID: 1838, Key: "r12", Value: `Sol Rune`},
	},
	"r12L": {
		{ID: 1135, Key: "r12L", Value: `Sol`},
	},
	"r13": {
		{ID: 1829, Key: "r13", Value: `Shae Rune`},
	},
	"r13L": {
		{ID: 1154, Key: "r13L", Value: `Shae`},
	},
	"r14": {
		{ID: 1830, Key: "r14", Value: `Dol Rune`},
	},
	"r14L": {
		{ID: 1167, Key: "r14L", Value: `Dol`},
	},
	"r15": {
		{ID: 1831, Key: "r15", Value: `Hel Rune`},
	},
	"r15L": {
		{ID: 1183, Key: "r15L", Value: `Hel`},
	},
	"r16": {
		{ID: 1837, Key: "r16", Value: `Po Rune`},
	},
	"r16L": {
		{ID: 1199, Key: "r16L", Value: `Po`},
	},
	"r17": {
		{ID: 1836, Key: "r17", Value: `Lum Rune`},
	},
	"r17L": {
		{ID: 1215, Key: "r17L", Value: `Lum`},
	},
	"r18": {
		{ID: 1834, Key: "r18", Value: `Ko Rune`},
	},
	"r18L": {
		{ID: 1231, Key: "r18L", Value: `Ko`},
	},
	"r19": {
		{ID: 1835, Key: "r19", Value: `Fal Rune`},
	},
	"r19L": {
		{ID: 1247, Key: "r19L", Value: `Fal`},
	},
	"r20": {
		{ID: 1866, Key: "r20", Value: `Lem Rune`},
	},
	"r20L": {
		{ID: 1358, Key: "r20L", Value: `Lem`},
	},
	"r21": {
		{ID: 1865, Key: "r21", Value: `Pul Rune`},
	},
	"r21L": {
		{ID: 1374, Key: "r21L", Value: `Pul`},
	},
	"r22": {
		{ID: 1864, Key: "r22", Value: `Um Rune`},
	},
	"r22L": {
		{ID: 1390, Key: "r22L", Value: `Um`},
	},
	"r23": {
		{ID: 1863, Key: "r23", Value: `Mal Rune`},
	},
	"r23L": {
		{ID: 1406, Key: "r23L", Value: `Mal`},
	},
	"r24": {
		{ID: 1862, Key: "r24", Value: `Ist Rune`},
	},
	"r24L": {
		{ID: 1422, Key: "r24L", Value: `Ist`},
	},
	"r25": {
		{ID: 1857, Key: "r25", Value: `Gul Rune`},
	},
	"r25L": {
		{ID: 1438, Key: "r25L", Value: `Gul`},
	},
	"r26": {
		{ID: 1856, Key: "r26", Value: `Vex Rune`},
	},
	"r26L": {
		{ID: 1454, Key: "r26L", Value: `Vex`},
	},
	"r27": {
		{ID: 1855, Key: "r27", Value: `Ohm Rune`},
	},
	"r27L": {
		{ID: 1470, Key: "r27L", Value: `Ohm`},
	},
	"r28": {
		{ID: 1854, Key: "r28", Value: `Lo Rune`},
	},
	"r28L": {
		{ID: 1486, Key: "r28L", Value: `Lo`},
	},
	"r29": {
		{ID: 1853, Key: "r29", Value: `Sur Rune`},
	},
	"r29L": {
		{ID: 1502, Key: "r29L", Value: `Sur`},
	},
	"r30": {
		{ID: 1858, Key: "r30", Value: `Ber Rune`},
	},
	"r30L": {
		{ID: 1614, Key: "r30L", Value: `Ber`},
	},
	"r31": {
		{ID: 1859, Key: "r31", Value: `Jo Rune`},
	},
	"r31L": {
		{ID: 1631, Key: "r31L", Value: `Jo`},
	},
	"r32": {
		{ID: 1860, Key: "r32", Value: `Cham Rune`},
	},
	"r32L": {
		{ID: 1646, Key: "r32L", Value: `Cham`},
	},
	"r33": {
		{ID: 1861, Key: "r33", Value: `Zod Rune`},
	},
	"r33L": {
		{ID: 1662, Key: "r33L", Value: `Zod`},
	},
	"red light": {
		{ID: 678, Key: "red light", Value: `Red Light`},
	},
	"sbub": {
		{ID: 414, Key: "sbub", Value: `Shrub`},
	},
	"scz": {
		{ID: 139, Key: "scz", Value: `Scalp`},
	},
	"secret object": {
		{ID: 372, Key: "secret object", Value: `secret object`},
	},
	"skr": {
		{ID: 258, Key: "skr", Value: `Scissors Katar`},
	},
	"sol": {
		{ID: 316, Key: "sol", Value: `Soul`},
	},
	"spe": {
		{ID: 325, Key: "spe", Value: `Spleen`},
	},
	"strAmazonOnly": {
		{ID: 1297, Key: "strAmazonOnly", Value: `(Amazon Only)`},
	},
	"strAssassinOnly": {
		{ID: 634, Key: "strAssassinOnly", Value: `(Assassin Only)`},
	},
	"strBarbarianOnly": {
		{ID: 300, Key: "strBarbarianOnly", Value: `(Barbarian Only)`},
	},
	"strDruidOnly": {
		{ID: 299, Key: "strDruidOnly", Value: `(Druid Only)`},
	},
	"strUI10": {
		{ID: 965, Key: "strUI10", Value: `Solar Creeper`},
	},
	"strUI11": {
		{ID: 966, Key: "strUI11", Value: `Spirit of Barbs`},
	},
	"strUI12": {
		{ID: 971, Key: "strUI12", Value: `Heart of Wolverine`},
	},
	"strUI13": {
		{ID: 972, Key: "strUI13", Value: `Vine`},
	},
	"strUI14": {
		{ID: 973, Key: "strUI14", Value: `Spirit`},
	},
	"strUI5": {
		{ID: 29, Key: "strUI5", Value: `Raven`},
	},
	"strUI6": {
		{ID: 30, Key: "strUI6", Value: `Wolf`},
	},
	"strUI7": {
		{ID: 31, Key: "strUI7", Value: `Bear`},
	},
	"strUI8": {
		{ID: 34, Key: "strUI8", Value: `Poison Creeper`},
	},
	"strUI9": {
		{ID: 43, Key: "strUI9", Value: `Carrion Vine`},
	},
	"strepilogueX": {
		{ID: 2545, Key: "strepilogueX", Value: `TERROR'S END`},
	},
	"strlastcinematic": {
		{ID: 1734, Key: "strlastcinematic", Value: `DESTRUCTION'S END`},
	},
	"swingingheads": {
		{ID: 337, Key: "swingingheads", Value: `Swinging Heads`},
	},
	"tal": {
		{ID: 348, Key: "tal", Value: `Tail`},
	},
	"tomb1": {
		{ID: 361, Key: "tomb1", Value: `Tomb`},
	},
	"tomb1L": {
		{ID: 909, Key: "tomb1L", Value: `Tomb`},
	},
	"tomb2": {
		{ID: 359, Key: "tomb2", Value: `Tomb`},
	},
	"tomb2L": {
		{ID: 1005, Key: "tomb2L", Value: `Tomb`},
	},
	"tomb3": {
		{ID: 354, Key: "tomb3", Value: `Tomb`},
	},
	"tomb3L": {
		{ID: 1004, Key: "tomb3L", Value: `Tomb`},
	},
	"torch1": {
		{ID: 353, Key: "torch1", Value: `Torch`},
	},
	"torch2": {
		{ID: 352, Key: "torch2", Value: `Torch`},
	},
	"tr2": {
		{ID: 562, Key: "tr2", Value: `Right Click to Cast
Scroll of Resistance`},
	},
	"tribal flag": {
		{ID: 326, Key: "tribal flag", Value: `Tribal Flag`},
	},
	"ttor": {
		{ID: 363, Key: "ttor", Value: `Torch`},
	},
	"uap": {
		{ID: 608, Key: "uap", Value: `Shako`},
	},
	"uar": {
		{ID: 610, Key: "uar", Value: `Sacred Armor`},
	},
	"ubub": {
		{ID: 415, Key: "ubub", Value: `Shrub`},
	},
	"ucl": {
		{ID: 637, Key: "ucl", Value: `Loricated Mail`},
	},
	"uea": {
		{ID: 657, Key: "uea", Value: `Wyrmhide`},
	},
	"uh9": {
		{ID: 665, Key: "uh9", Value: `Bone Visage`},
	},
	"uhb": {
		{ID: 706, Key: "uhb", Value: `Myrmidon Greaves`},
	},
	"uhc": {
		{ID: 707, Key: "uhc", Value: `Colossus Girdle`},
	},
	"uhg": {
		{ID: 711, Key: "uhg", Value: `Ogre Gauntlets`},
	},
	"uhl": {
		{ID: 716, Key: "uhl", Value: `Giant Conch`},
	},
	"uhm": {
		{ID: 717, Key: "uhm", Value: `Spired Helm`},
	},
	"uhn": {
		{ID: 718, Key: "uhn", Value: `Boneweave`},
	},
	"uit": {
		{ID: 740, Key: "uit", Value: `Monarch`},
	},
	"ukp": {
		{ID: 768, Key: "ukp", Value: `Hydraskull`},
	},
	"ula": {
		{ID: 769, Key: "ula", Value: `Scarab Husk`},
	},
	"ulb": {
		{ID: 770, Key: "ulb", Value: `Wyrmhide Boots`},
	},
	"ulc": {
		{ID: 771, Key: "ulc", Value: `Spiderweb Sash`},
	},
	"uld": {
		{ID: 772, Key: "uld", Value: `Kraken Shell`},
	},
	"ulg": {
		{ID: 775, Key: "ulg", Value: `Bramble Mitts`},
	},
	"ulm": {
		{ID: 781, Key: "ulm", Value: `Armet`},
	},
	"ult": {
		{ID: 788, Key: "ult", Value: `Hellforge Plate`},
	},
	"umb": {
		{ID: 786, Key: "umb", Value: `Boneweave Boots`},
	},
	"umc": {
		{ID: 787, Key: "umc", Value: `Mithril Coil`},
	},
	"umg": {
		{ID: 791, Key: "umg", Value: `Vambraces`},
	},
	"uml": {
		{ID: 796, Key: "uml", Value: `Luna`},
	},
	"uncle f#%* comedy central(c)": {
		{ID: 780, Key: "uncle f#%* comedy central(c)", Value: `Moe`},
	},
	"ung": {
		{ID: 807, Key: "ung", Value: `Diamond Mail`},
	},
	"uow": {
		{ID: 839, Key: "uow", Value: `Aegis`},
	},
	"upk": {
		{ID: 843, Key: "upk", Value: `Blade Barrier`},
	},
	"upl": {
		{ID: 844, Key: "upl", Value: `Balrog Skin`},
	},
	"urg": {
		{ID: 871, Key: "urg", Value: `Hyperion`},
	},
	"urn": {
		{ID: 878, Key: "urn", Value: `Corona`},
	},
	"urs": {
		{ID: 884, Key: "urs", Value: `Great Hauberk`},
	},
	"ush": {
		{ID: 888, Key: "ush", Value: `Troll Nest`},
	},
	"usk": {
		{ID: 891, Key: "usk", Value: `Demonhead`},
	},
	"utb": {
		{ID: 898, Key: "utb", Value: `Mirrored Boots`},
	},
	"utc": {
		{ID: 899, Key: "utc", Value: `Troll Belt`},
	},
	"utg": {
		{ID: 903, Key: "utg", Value: `Crusader Gauntlets`},
	},
	"uth": {
		{ID: 904, Key: "uth", Value: `Lacquered Plate`},
	},
	"utp": {
		{ID: 912, Key: "utp", Value: `Archon Plate`},
	},
	"uts": {
		{ID: 916, Key: "uts", Value: `Ward`},
	},
	"utu": {
		{ID: 917, Key: "utu", Value: `Wire Fleece`},
	},
	"uuc": {
		{ID: 915, Key: "uuc", Value: `Heater`},
	},
	"uui": {
		{ID: 921, Key: "uui", Value: `Dusk Shroud`},
	},
	"uul": {
		{ID: 924, Key: "uul", Value: `Shadow Plate`},
	},
	"uvb": {
		{ID: 930, Key: "uvb", Value: `Scarabshell Boots`},
	},
	"uvc": {
		{ID: 931, Key: "uvc", Value: `Vampirefang Belt`},
	},
	"uvg": {
		{ID: 935, Key: "uvg", Value: `Vampirebone Gloves`},
	},
	"well": {
		{ID: 311, Key: "well", Value: `Well`},
	},
	"woodchest2L": {
		{ID: 546, Key: "woodchest2L", Value: `Wooden Chest`},
	},
	"woodchest2R": {
		{ID: 560, Key: "woodchest2R", Value: `Wooden Chest`},
	},
	"woodchestL": {
		{ID: 308, Key: "woodchestL", Value: `Wooden Chest`},
	},
	"woodchestR": {
		{ID: 313, Key: "woodchestR", Value: `Wooden Chest`},
	},
	"wrb": {
		{ID: 1377, Key: "wrb", Value: `Wrist Blade`},
	},
	"x": {
		{ID: 122, Key: "x", Value: `Unable to enter game. Your character must kill Baal in Nightmare difficulty to play in a Hell game.`},
		{ID: 123, Key: "x", Value: `Unable to enter game. Your character must kill Baal to play in a Nightmare game.`},
		{ID: 124, Key: "x", Value: `CREATE NEW`},
		{ID: 125, Key: "x", Value: `CHARACTER`},
		{ID: 144, Key: "x", Value: `Unable to access file. Cannot convert character.`},
		{ID: 150, Key: "x", Value: `Expansion Disc`},
		{ID: 151, Key: "x", Value: `Â© Copyright 2001 Blizzard Entertainment`},
		{ID: 305, Key: "x", Value: `Commanding the forces of nature, he summons wild beasts and raging storms to his side.`},
		{ID: 306, Key: "x", Value: `Schooled in the Martial Arts, her mind and body are deadly weapons.`},
		{ID: 803, Key: "x", Value: `EXPANSION`},
		{ID: 804, Key: "x", Value: `EXPANSION CHARACTER`},
		{ID: 805, Key: "x", Value: `CONVERT TO`},
		{ID: 808, Key: "x", Value: `CLICK OK TO CONVERT YOUR ORIGINAL DIABLO II CHARACTER TO PLAY IN EXPANSION GAMES.`},
		{ID: 819, Key: "x", Value: `NEW EXPANSION CHARACTER`},
		{ID: 821, Key: "x", Value: `NEW EXPANSION OPEN CHARACTER.`},
		{ID: 823, Key: "x", Value: `NEW EXPANSION REALM CHARACTER.`},
	},
}

var expansionstring_TblDataByID = map[int]tbl.TblEntry{
	1:  {ID: 1, Key: "Trang-Oul's Scales", Value: `Trang-Oul's Scales`},
	2:  {ID: 2, Key: "Ormus' Robes", Value: `Ormus' Robes`},
	3:  {ID: 3, Key: "Graverobber's", Value: `Graverobber's`},
	4:  {ID: 4, Key: "Eskillpowerup3", Value: `Charge 3 - `},
	5:  {ID: 5, Key: "Eskillsinishup", Value: `Finishing Move Bonuses`},
	6:  {ID: 6, Key: "Eskillpudlife", Value: ` percent life stealing`},
	7:  {ID: 7, Key: "Eskillpuddgmper", Value: ` percent damage`},
	8:  {ID: 8, Key: "Lord of Destruction", Value: ``},
	9:  {ID: 9, Key: "Eskillferalpets", Value: `Feral Pets: `},
	10: {ID: 10, Key: "FindingBeneathCityAss", Value: `And I thought the Forgotten Tower stank.`},
	11: {ID: 11, Key: "Eskillpercentatt", Value: ` Percent Attack`},
	12: {ID: 12, Key: "Eskillpercentdmg", Value: ` Percent Damage`},
	13: {ID: 13, Key: "Gymnast's", Value: `Gymnastic`},
	14: {ID: 14, Key: "of Bliss", Value: `of Bliss`},
	15: {ID: 15, Key: "Eskillfinishmove", Value: `Finishing Move - `},
	16: {ID: 16, Key: "Glimmershred", Value: `Glimmershred`},
	17: {ID: 17, Key: "Dust Storm", Value: `Dust Storm`},
	18: {ID: 18, Key: "Eskillpudburning", Value: `burning damage`},
	19: {ID: 19, Key: "Eskillfistsoffire1", Value: `fire damage: `},
	20: {ID: 20, Key: "Dusty", Value: `Dusty`},
	21: {ID: 21, Key: "ExInsertSockets", Value: `Can be Inserted into Socketed Items`},
	22: {ID: 22, Key: "Eskillfistsoffire2", Value: `fire damage radius: `},
	23: {ID: 23, Key: "Eskillfistsoffire3", Value: `burning duration: `},
	24: {ID: 24, Key: "RescueQual-KehkAct5Pal", Value: `Follow me.`},
	25: {ID: 25, Key: "Eskillbladesofice1", Value: `cold damage: `},
	26: {ID: 26, Key: "Compact", Value: `Compact`},
	27: {ID: 27, Key: "Eskillbladesofice2", Value: `cold damage radius: `},
	28: {ID: 28, Key: "Eskillbladesofice3", Value: `freeze duration: `},
	29: {ID: 29, Key: "strUI5", Value: `Raven`},
	30: {ID: 30, Key: "strUI6", Value: `Wolf`},
	31: {ID: 31, Key: "strUI7", Value: `Bear`},
	32: {ID: 32, Key: "Giantmaimer", Value: `Giant Maimer`},
	33: {ID: 33, Key: "LeavingTownAct5Ass", Value: `You'll pay for your atrocities, Baal.`},
	34: {ID: 34, Key: "strUI8", Value: `Poison Creeper`},
	35: {ID: 35, Key: "Shocking", Value: `Shocking`},
	36: {ID: 36, Key: "Consecrated", Value: `Consecrated`},
	37: {ID: 37, Key: "of Amianthus", Value: `of Amianthus`},
	38: {ID: 38, Key: "of Bone Walls", Value: `of Bone Wall`},
	39: {ID: 39, Key: "Hell's Whisper", Value: `Hell Whisper`},
	40: {ID: 40, Key: "Metalgrid", Value: `Metalgrid`},
	41: {ID: 41, Key: "Hibernal", Value: `Hibernal`},
	42: {ID: 42, Key: "Naj's Puzzler", Value: `Naj's Puzzler`},
	43: {ID: 43, Key: "strUI9", Value: `Carrion Vine`},
	44: {ID: 44, Key: "of Faith", Value: `of Faith`},
	45: {ID: 45, Key: "A5Q3InitMalah", Value: `43
There is a matter which I hesitate to 
share, but I believe you are the only 
one who can help me now.
 
Anya, the young alchemist and 
daughter to one of our slain Elders, 
has been missing for some time. She is 
a strong, crafty woman with a spirit 
like no other.
 
One night, just before your arrival, I 
overheard her and Nihlathak arguing 
about her father's death. The next 
morning she was gone.
 
Nihlathak has his own tale as to where 
she went and why. Don't believe him! I 
fear he is at the root of her 
disappearance.
 
Please, if you can, search for Anya and 
bring her back to us. She'll know what 
to do about Nihlathak.
`},
	46: {ID: 46, Key: "Bul-Kathos' Might", Value: `Bul-Kathos' Might`},
	47: {ID: 47, Key: "Gheed's Fortune", Value: `Gheed's Fortune`},
	48: {ID: 48, Key: "UIFenirsui", Value: `Dire Wolf`},
	49: {ID: 49, Key: "UiRescuedBarUI", Value: `Warrior`},
	50: {ID: 50, Key: "UiShadowUI", Value: `Shadow`},
	51: {ID: 51, Key: "Credendum", Value: `Credendum`},
	52: {ID: 52, Key: "EnteringNihlathakAct5Nec", Value: `Ahh, the familiar scent of death.`},
	53: {ID: 53, Key: "Meatscrape", Value: `Meatscrape`},
	54: {ID: 54, Key: "CompletingDefeatBaalAct5Dru", Value: `Baal! Join your brothers in oblivion.`},
	55: {ID: 55, Key: "Wraithflight", Value: `Wraith Flight`},
	56: {ID: 56, Key: "of Bone Armor", Value: `of Bone Armor`},
	57: {ID: 57, Key: "Steller", Value: `Stellar`},
	58: {ID: 58, Key: "Spike Generator", Value: `Spike Generator`},
	59: {ID: 59, Key: "Charged Bolt Sentry", Value: `Charged Bolt Sentry`},
	60: {ID: 60, Key: "Lightning Sentry", Value: `Lightning Sentry`},
	61: {ID: 61, Key: "Of Fortication", Value: `of Fortification`},
	62: {ID: 62, Key: "Blade Creeper", Value: `Blade Sentinel`},
	63: {ID: 63, Key: "A5Q4SuccessfulQualKehk", Value: `57
Nihlathak was a vile demon that shall 
find his home among the tortured 
minions of Hell!
 
You battled the Darkness without fear. I 
laud your skill and courage.
`},
	64:  {ID: 64, Key: "Stonerattle", Value: `Stone Rattle`},
	65:  {ID: 65, Key: "LeavingTownAct5Bar", Value: `The time has come to cleanse my homeland!`},
	66:  {ID: 66, Key: "Druid Wolf", Value: `Wolf`},
	67:  {ID: 67, Key: "of the Efreeti", Value: `of the Efreeti`},
	68:  {ID: 68, Key: "Druid Fenris", Value: `Dire Wolf`},
	69:  {ID: 69, Key: "Sanguinary", Value: `Sanguinary`},
	70:  {ID: 70, Key: "Druid Bear", Value: `Grizzly`},
	71:  {ID: 71, Key: "Druid Cycle of Life", Value: `Cycle of Life`},
	72:  {ID: 72, Key: "Bear", Value: `Bear`},
	73:  {ID: 73, Key: "ReanimatedHorde", Value: `Reanimated Horde`},
	74:  {ID: 74, Key: "Of Vines", Value: `of Vines`},
	75:  {ID: 75, Key: "ProwlingDead", Value: `Prowling Dead`},
	76:  {ID: 76, Key: "UnholyCorpse", Value: `Unholy Corpse`},
	77:  {ID: 77, Key: "DefiledWarrior", Value: `Defiled Warrior`},
	78:  {ID: 78, Key: "Siege Beast", Value: `Siege Beast`},
	79:  {ID: 79, Key: "CrushBiest", Value: `Crush Beast`},
	80:  {ID: 80, Key: "GoreBearer", Value: `Gore Bearer`},
	81:  {ID: 81, Key: "LifeSeeker", Value: `Life Seeker`},
	82:  {ID: 82, Key: "LifeStealer", Value: `Life Stealer`},
	83:  {ID: 83, Key: "DeathlyVisage", Value: `Deathly Visage`},
	84:  {ID: 84, Key: "BoundSpirit", Value: `Bound Spirit`},
	85:  {ID: 85, Key: "Deathexp", Value: `Death`},
	86:  {ID: 86, Key: "Blackhorn's Face", Value: `Blackhorn's Face`},
	87:  {ID: 87, Key: "Of the Sky", Value: `of the Sky`},
	88:  {ID: 88, Key: "Heaven's Brethren", Value: `Heaven's Brethren`},
	89:  {ID: 89, Key: "Green", Value: `Green`},
	90:  {ID: 90, Key: "Slayerexp", Value: `Slayer`},
	91:  {ID: 91, Key: "of Vengeance", Value: `of Vengeance`},
	92:  {ID: 92, Key: "IceBoar", Value: `Ice Boar`},
	93:  {ID: 93, Key: "of Frozen Orbs", Value: `of Frozen Orb`},
	94:  {ID: 94, Key: "ConsumedFireBoar", Value: `Consumed Fire Boar`},
	95:  {ID: 95, Key: "Griswold's Valor", Value: `Griswold's Valor`},
	96:  {ID: 96, Key: "Hallowed", Value: `Hallowed`},
	97:  {ID: 97, Key: "Demon's Arch", Value: `Demon's Arch`},
	98:  {ID: 98, Key: "Atma's Wail", Value: `Atma's Wail`},
	99:  {ID: 99, Key: "Wizardspike", Value: `Wizardspike`},
	100: {ID: 100, Key: "ConsumedIceBoar", Value: `Consumed Ice Boar`},
	101: {ID: 101, Key: "Cranebeak", Value: `Cranebeak`},
	102: {ID: 102, Key: "CompletingTombDru", Value: `Diablo... I will find you yet.`},
	103: {ID: 103, Key: "FrenziedHellSpawn", Value: `Frenzied Hell Spawn`},
	104: {ID: 104, Key: "A4Q2ExpansionSuccessCain", Value: `40
I knew there was great potential in you, 
my friend. You've done a fantastic job.
 
Though my ancestors often struggled 
against the Three Evils and their 
minions, I've always lived a shut-in, 
scholarly life. I'm glad that my wisdom 
aided you.
 
Now, I wish to leave this place. Though 
Heaven's Gates are a marvel to behold, 
I hope I won't have to see them again 
for many, many years.
 
Please talk to Tyrael about leaving this 
place now!
`},
	105: {ID: 105, Key: "Eburin", Value: `Eburine`},
	106: {ID: 106, Key: "Viperfork", Value: `Viperfork`},
	107: {ID: 107, Key: "Baal Crab Clone", Value: `Baal`},
	108: {ID: 108, Key: "FrenziedIceSpawn", Value: `Frenzied Ice Spawn`},
	109: {ID: 109, Key: "InsaneHellSpawn", Value: `Insane Hell Spawn`},
	110: {ID: 110, Key: "Ogun's Lash", Value: `Ogun's Lash`},
	111: {ID: 111, Key: "Succubusexp", Value: `Succubus`},
	112: {ID: 112, Key: "VileTemptress", Value: `Vile Temptress`},
	113: {ID: 113, Key: "Hell Temptress", Value: `Hell Temptress`},
	114: {ID: 114, Key: "Blood Temptress", Value: `Blood Temptress`},
	115: {ID: 115, Key: "Dominus", Value: `Siren`},
	116: {ID: 116, Key: "Of Elusion", Value: `of Eluding`},
	117: {ID: 117, Key: "Blood Witch", Value: `Blood Witch`},
	118: {ID: 118, Key: "Hell Witch", Value: `Hell Witch`},
	119: {ID: 119, Key: "OverSeer", Value: `Overseer`},
	120: {ID: 120, Key: "Lasher", Value: `Lasher`},
	121: {ID: 121, Key: "Captain's", Value: `Captain's`},
	122: {ID: 122, Key: "x", Value: `Unable to enter game. Your character must kill Baal in Nightmare difficulty to play in a Hell game.`},
	123: {ID: 123, Key: "x", Value: `Unable to enter game. Your character must kill Baal to play in a Nightmare game.`},
	124: {ID: 124, Key: "x", Value: `CREATE NEW`},
	125: {ID: 125, Key: "x", Value: `CHARACTER`},
	126: {ID: 126, Key: "Hell1", Value: `Abaddon`},
	127: {ID: 127, Key: "of Holy Bolts", Value: `of Holy Bolt`},
	128: {ID: 128, Key: "Highlord's Wrath", Value: `Highlord's Wrath`},
	129: {ID: 129, Key: "Hell2", Value: `Pit of Acheron`},
	130: {ID: 130, Key: "Hell3", Value: `Infernal Pit`},
	131: {ID: 131, Key: "Bush Wacker", Value: `Bush Wacker`},
	132: {ID: 132, Key: "of Poison Javelin", Value: `of Poison Javelin`},
	133: {ID: 133, Key: "A5Q2SuccessfulMalah", Value: `73
You have inspired the people in this 
town -- not only those you rescued, but 
those you've helped in other ways as 
well.
`},
	134: {ID: 134, Key: "Nightsummon", Value: `Nightsummon`},
	135: {ID: 135, Key: "Herald of Zakarum", Value: `Herald of Zakarum`},
	136: {ID: 136, Key: "0sc", Value: `Scroll of Knowledge`},
	137: {ID: 137, Key: "NihlathakAct5IntroNecGossip1", Value: `45
Ahhh, a Necromancer.
 
While I admire your courage in seeking 
out the darker side of magic, we really 
have little need of your skills. The 
battle will turn soon enough without 
your meddling. 
 
Yet, I should have expected to see your 
kind here. You are like a moth to the 
flame -- drawn to all this death. It 
feeds you in more ways than one, does 
it not?
`},
	138: {ID: 138, Key: "A5Q3AfterInitMalah", Value: `94
When you talk to Nihlathak, be careful. 
There is no telling what he will say or 
do.
`},
	139: {ID: 139, Key: "scz", Value: `Scalp`},
	140: {ID: 140, Key: "of Shiver Armor", Value: `of Shiver Armor`},
	141: {ID: 141, Key: "Smoking", Value: `Smoking`},
	142: {ID: 142, Key: "Pompe's Wrath", Value: `Pompeii's Wrath`},
	143: {ID: 143, Key: "Arreat's Face", Value: `Arreat's Face`},
	144: {ID: 144, Key: "x", Value: `Unable to access file. Cannot convert character.`},
	145: {ID: 145, Key: "CainAct5Gossip10", Value: `54
It is my belief that the Soulstones are at 
the center of this conflict. If only that 
fool Marius had not intervened, Baal 
would still be imprisoned within Tal 
Rasha.
`},
	146: {ID: 146, Key: "Feral", Value: `Feral`},
	147: {ID: 147, Key: "Blackhand Key", Value: `Blackhand Key`},
	148: {ID: 148, Key: "of Stoicism", Value: `of Stoicism`},
	149: {ID: 149, Key: "Rockstopper", Value: `Rockstopper`},
	150: {ID: 150, Key: "x", Value: `Expansion Disc`},
	151: {ID: 151, Key: "x", Value: `Â© Copyright 2001 Blizzard Entertainment`},
	152: {ID: 152, Key: "Ambergris", Value: `Ambergris`},
	153: {ID: 153, Key: "of Resistance", Value: `of Resistance`},
	154: {ID: 154, Key: "Lapis Lazuli", Value: `Lapis Lazuli`},
	155: {ID: 155, Key: "Of Blitzen", Value: `of Blitzen`},
	156: {ID: 156, Key: "Resonant", Value: `Resonant`},
	157: {ID: 157, Key: "of Power Spear", Value: `of Power Spear`},
	158: {ID: 158, Key: "Deathbit", Value: `Deathbit`},
	159: {ID: 159, Key: "Crown of Thieves", Value: `Crown of Thieves`},
	160: {ID: 160, Key: "FindingLamEsenDru", Value: `An ancient manuscript...  This could be useful.`},
	161: {ID: 161, Key: "QualKehkAct5IntroPalGossip1", Value: `42
A Paladin! I have long heard of your 
people.
 
As a young warrior I even considered 
the pilgrimage to Kurast. But I was 
younger then and foolish. My place has 
always been here -- protecting 
Harrogath, and Mount Arreat with it.
 
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`},
	162: {ID: 162, Key: "of Magic Arrows", Value: `of Magic Arrow`},
	163: {ID: 163, Key: "EnteringWorldstoneAct5Nec", Value: `So, this is what the Ancients guard.`},
	164: {ID: 164, Key: "Eskillpowerupadd", Value: `Adds `},
	165: {ID: 165, Key: "Invis Pet", Value: `Pet`},
	166: {ID: 166, Key: "BloodBringer", Value: `Blood Bringer`},
	167: {ID: 167, Key: "VileWitch", Value: `Vile Witch`},
	168: {ID: 168, Key: "OverLord", Value: `Overlord`},
	169: {ID: 169, Key: "BloodBoss", Value: `Blood Boss`},
	170: {ID: 170, Key: "HellWhip", Value: `Hell Whip`},
	171: {ID: 171, Key: "MinionSlayerSpawner", Value: `Demon Portal`},
	172: {ID: 172, Key: "Minionice/hellSpawnSpawner", Value: `Demon Portal`},
	173: {ID: 173, Key: "of Battle Cry", Value: `of Battle Cry`},
	174: {ID: 174, Key: "Syrian", Value: `Triad's Foliage`},
	175: {ID: 175, Key: "MinionGreaterIce/hellSpawnSpawner", Value: `Demon Portal`},
	176: {ID: 176, Key: "A5Q3EarlyReturnNihlathak", Value: `46
Look, I've told you! She's dead! If you 
knew what was good for you, you'd 
concentrate your efforts on saving 
Harrogath -- not on lost causes like 
Anya.
`},
	177: {ID: 177, Key: "Pitblood Thirst", Value: `Pitblood Thirst`},
	178: {ID: 178, Key: "EnterForgottenTAss", Value: `Who would want to remember this place?`},
	179: {ID: 179, Key: "PalOnly", Value: `(Paladin Only)`},
	180: {ID: 180, Key: "Imp1", Value: `Demon Imp`},
	181: {ID: 181, Key: "Imp2", Value: `Demon Rascal`},
	182: {ID: 182, Key: "Imp3", Value: `Demon Gremlin`},
	183: {ID: 183, Key: "Elysian", Value: `Elysian`},
	184: {ID: 184, Key: "Chromatic Ire", Value: `Chromatic Ire`},
	185: {ID: 185, Key: "Darkforge Spawn", Value: `Darkforce Spawn`},
	186: {ID: 186, Key: "Bloodletter", Value: `Bloodletter`},
	187: {ID: 187, Key: "Imp4", Value: `Demon Trickster`},
	188: {ID: 188, Key: "A5Q2SuccessfulQualKehk", Value: `51
Thank you for rescuing my men. They 
have spoken well of your bravery in 
battle. Perhaps there is hope for us 
after all. 
 
If you wish, you may hire some of my 
mercenaries that you saved. And 
please...take this set of runes. I had 
been saving them for a socketed shield, 
but I think you'll make better use of 
them.
 
Be sure to set them in the right order 
for their fullest effect.
`},
	189: {ID: 189, Key: "Islestrike", Value: `Islestrike`},
	190: {ID: 190, Key: "Nature's Peace", Value: `Nature's Peace`},
	191: {ID: 191, Key: "of Prosperity", Value: `of Prosperity`},
	192: {ID: 192, Key: "Of the Obscenity", Value: `of the Obscene`},
	193: {ID: 193, Key: "Carnageleaver", Value: `Carnage Leaver`},
	194: {ID: 194, Key: "Imp5", Value: `Demon Sprite`},
	195: {ID: 195, Key: "Frozen Horror1", Value: `Frozen Creeper`},
	196: {ID: 196, Key: "A5Q4AfterInitLarzuk", Value: `56
Now, rescuing Anya -- that's good. 
Talking to me while Nihlathak hands 
over the Relic to Baal -- uh...that's bad!
`},
	197: {ID: 197, Key: "Sankenkur's Resurrection", Value: `Sankekur's Resurrection`},
	198: {ID: 198, Key: "of Envy", Value: `of Envy`},
	199: {ID: 199, Key: "StygianFury", Value: `Stygian Fury`},
	200: {ID: 200, Key: "A5Q5EarlyReturnCain", Value: `52
We have come too far to be defeated 
now, my friend. I have seen you 
complete many difficult quests. Though 
this may be your greatest trial, it is not 
beyond your reach.
`},
	201: {ID: 201, Key: "Frozen Horror2", Value: `Frozen Terror`},
	202: {ID: 202, Key: "Frozen Horror3", Value: `Frozen Scourge`},
	203: {ID: 203, Key: "Frozen Horror4", Value: `Frozen Horror`},
	204: {ID: 204, Key: "Frozen Horror5", Value: `Frozen Scorch`},
	205: {ID: 205, Key: "Catapult Spotter N", Value: `Catapult`},
	206: {ID: 206, Key: "Catapult Spotter S", Value: `Catapult`},
	207: {ID: 207, Key: "of Power", Value: `of Power`},
	208: {ID: 208, Key: "Sazabi's Grand Tribute", Value: `Sazabi's Grand Tribute`},
	209: {ID: 209, Key: "Catapult Spotter E", Value: `Catapult`},
	210: {ID: 210, Key: "of Atlus", Value: `of Atlas`},
	211: {ID: 211, Key: "KillingdDiabloSor", Value: `Terror stalks Hell no more.`},
	212: {ID: 212, Key: "LeaveCampAss", Value: `They'll never see me coming.`},
	213: {ID: 213, Key: "Catapult Spotter W", Value: `Catapult`},
	214: {ID: 214, Key: "of Guided Arrows", Value: `of Guided Arrow`},
	215: {ID: 215, Key: "Catapult Spotter Siege", Value: `Catapult`},
	216: {ID: 216, Key: "BeginTaintedSunDru", Value: `Strange... an unexpected eclipse.`},
	217: {ID: 217, Key: "Harpoonist's", Value: `Harpoonist's`},
	218: {ID: 218, Key: "CatapultSiege", Value: `Catapult`},
	219: {ID: 219, Key: "A5Q1EarlyReturnMalah", Value: `64
This may not be very encouraging, but 
you have fared better than the others 
who passed through those gates.
 
In regards to Shenk the Overseer, 
remember: a general is nothing without 
his men.
`},
	220: {ID: 220, Key: "FindingTrueTombDru", Value: `So, Tal Rasha... This is your resting place.`},
	221: {ID: 221, Key: "Sounding", Value: `Sounding`},
	222: {ID: 222, Key: "EnteringWorldstoneAct5Sor", Value: `The power of the Worldstone washes over me.`},
	223: {ID: 223, Key: "Soulfeast Tine", Value: `Soulfeast Tine`},
	224: {ID: 224, Key: "Barricade Wall Right", Value: `Barricade`},
	225: {ID: 225, Key: "GreaterHellSpawn", Value: `Greater Hell Spawn`},
	226: {ID: 226, Key: "Infectious", Value: `Infectious`},
	227: {ID: 227, Key: "Barricade Door", Value: `Barricaded Door`},
	228: {ID: 228, Key: "Barricade Tower", Value: `Barricaded Tower`},
	229: {ID: 229, Key: "Siege Boss", Value: `Shenk the Overseer`},
	230: {ID: 230, Key: "Putrid Defiler1", Value: `Putrid Defiler`},
	231: {ID: 231, Key: "Putrid Defiler2", Value: `Wretched Defiler`},
	232: {ID: 232, Key: "Putrid Defiler3", Value: `Fetid Defiler`},
	233: {ID: 233, Key: "Putrid Defiler4", Value: `Rancid Defiler`},
	234: {ID: 234, Key: "Putrid Defiler5", Value: `Rank Defiler`},
	235: {ID: 235, Key: "Flamebellow", Value: `Flamebellow`},
	236: {ID: 236, Key: "Oak Sage", Value: `Oak Sage`},
	237: {ID: 237, Key: "Pearl", Value: `Pearl`},
	238: {ID: 238, Key: "Nord's Tenderizer", Value: `Nord's Tenderizer`},
	239: {ID: 239, Key: "Heart of Wolverine", Value: `Heart of Wolverine`},
	240: {ID: 240, Key: "Carmine", Value: `Carmine`},
	241: {ID: 241, Key: "Spirit of Barbs", Value: `Spirit of Barbs`},
	242: {ID: 242, Key: "Death Sentry", Value: `Death Sentry`},
	243: {ID: 243, Key: "Inferno Sentry", Value: `Inferno Sentry`},
	244: {ID: 244, Key: "BerserkSlayer", Value: `Berserker Slayer`},
	245: {ID: 245, Key: "Shadow Master", Value: `Shadow Master`},
	246: {ID: 246, Key: "Sander's Taboo", Value: `McAuley's Taboo`},
	247: {ID: 247, Key: "McAuley's Riprap", Value: `Sander's Riprap`},
	248: {ID: 248, Key: "of Fire Walls", Value: `of Fire Wall`},
	249: {ID: 249, Key: "Moser's Blessed Circle", Value: `Moser's Blessed Circle`},
	250: {ID: 250, Key: "Wake of Destruction", Value: `Wake of Fire`},
	251: {ID: 251, Key: "Ghostly", Value: `Ghostly`},
	252: {ID: 252, Key: "Fanatic", Value: `Fanatic`},
	253: {ID: 253, Key: "Possessed", Value: `Possessed`},
	254: {ID: 254, Key: "Squire's Cover", Value: `Squire's Cover`},
	255: {ID: 255, Key: "Berserk", Value: `Berserker`},
	256: {ID: 256, Key: "Larzuk", Value: `Larzuk`},
	257: {ID: 257, Key: "Of the Unicorn", Value: `of the Unicorn`},
	258: {ID: 258, Key: "skr", Value: `Scissors Katar`},
	259: {ID: 259, Key: "Zakarum's Salvation", Value: `Zakarum's Salvation`},
	260: {ID: 260, Key: "Hwanin's Splendor", Value: `Hwanin's Splendor`},
	261: {ID: 261, Key: "Drehya", Value: `Anya`},
	262: {ID: 262, Key: "Nihlathak Town", Value: `Nihlathak`},
	263: {ID: 263, Key: "Qual-Kehk", Value: `Qual-Kehk`},
	264: {ID: 264, Key: "Act 5 Townguard", Value: `Guard`},
	265: {ID: 265, Key: "Act 5 Combatant", Value: `Combatant`},
	266: {ID: 266, Key: "Larzuk's Champion", Value: `Larzuk's Champion`},
	267: {ID: 267, Key: "Nihlathak", Value: `Nihlathak`},
	268: {ID: 268, Key: "POW", Value: `Barbarian Captive`},
	269: {ID: 269, Key: "Moe", Value: `Moe`},
	270: {ID: 270, Key: "Demonhorn's Edge", Value: `Demonhorn's Edge`},
	271: {ID: 271, Key: "The Oculus", Value: `The Oculus`},
	272: {ID: 272, Key: "Titanfist", Value: `Titan's Fist`},
	273: {ID: 273, Key: "Curly", Value: `Curly`},
	274: {ID: 274, Key: "Visceratuant", Value: `Visceratuant`},
	275: {ID: 275, Key: "Larry", Value: `Larry`},
	276: {ID: 276, Key: "Sensei's", Value: `Sensei's`},
	277: {ID: 277, Key: "A5Q3FoundAnyaAnya", Value: `50
Hero. Nihlathak did this to me!
 
If you've come to help me, my only hope 
lies with Malah.
 
Please...Tell her you've found me...
`},
	278: {ID: 278, Key: "Ancient Barbarian 3", Value: `Korlic`},
	279: {ID: 279, Key: "NihlathakAct5IntroGossip1", Value: `43
Well, well. The siege has everything in 
short supply...except fools. 
 
Why would you seek this place, 
stranger? Are you a vulture come to 
loot the bodies of our fallen warriors? 
 
Regardless, this is no place to make a 
name for yourself. The mountain is 
ours to protect. It is only a matter of 
time before Hell's legions are routed.
`},
	280: {ID: 280, Key: "Ancient Barbarian 2", Value: `Madawc`},
	281: {ID: 281, Key: "Ancient Barbarian 1", Value: `Talic`},
	282: {ID: 282, Key: "Blaze Ripper", Value: `Blaze Ripper`},
	283: {ID: 283, Key: "Sharp Tooth Sayer", Value: `Sharptooth Slayer`},
	284: {ID: 284, Key: "Sacred Charge", Value: `Sacred Charge`},
	285: {ID: 285, Key: "Flowkrad's Grin", Value: `Flowkrad's Grin`},
	286: {ID: 286, Key: "EnteringRadamentDru", Value: `Face the light or lurk in darkness.`},
	287: {ID: 287, Key: "Cinnabar", Value: `Cinnabar`},
	288: {ID: 288, Key: "A5Q5EarlyReturnMalah", Value: `68
Do not doubt yourself. I believe you are 
worthy to pass through the Ancients' 
gates, but you must believe, as well.
`},
	289: {ID: 289, Key: "Flowkrad's Sinew", Value: `Flowkrad's Sinew`},
	290: {ID: 290, Key: "M'avina's Icy Clutch", Value: `M'avina's Icy Clutch`},
	291: {ID: 291, Key: "Rabbit Slayer", Value: `Rabbit Slayer`},
	292: {ID: 292, Key: "of the Lamprey", Value: `of the Lamprey`},
	293: {ID: 293, Key: "Of the Avenger", Value: `of the Avenger`},
	294: {ID: 294, Key: "Anodized Elite", Value: `Hell's Belle`},
	295: {ID: 295, Key: "Of Beauty", Value: `of Beauty`},
	296: {ID: 296, Key: "Snapchip Shatter", Value: `Snapchip Shatter`},
	297: {ID: 297, Key: "Megaflow Rectifier", Value: `Eldritch the Rectifier`},
	298: {ID: 298, Key: "Frozenstein", Value: `Frozenstein`},
	299: {ID: 299, Key: "strDruidOnly", Value: `(Druid Only)`},
	300: {ID: 300, Key: "strBarbarianOnly", Value: `(Barbarian Only)`},
	301: {ID: 301, Key: "LeavingTownAct5Nec", Value: `It takes more than a siege to stop me.`},
	302: {ID: 302, Key: "The Gavel of Pain", Value: `The Gavel of Pain`},
	303: {ID: 303, Key: "Horizon's Tornado", Value: `Horizon's Tornado`},
	304: {ID: 304, Key: "Trang-Oul's Girth", Value: `Trang-Oul's Girth`},
	305: {ID: 305, Key: "x", Value: `Commanding the forces of nature, he summons wild beasts and raging storms to his side.`},
	306: {ID: 306, Key: "x", Value: `Schooled in the Martial Arts, her mind and body are deadly weapons.`},
	307: {ID: 307, Key: "barrel wilderness", Value: `Barrel`},
	308: {ID: 308, Key: "woodchestL", Value: `Wooden Chest`},
	309: {ID: 309, Key: "burialchestL", Value: `Burial Chest`},
	310: {ID: 310, Key: "burialchestR", Value: `Burial Chest`},
	311: {ID: 311, Key: "well", Value: `Well`},
	312: {ID: 312, Key: "ChestL", Value: `Chest`},
	313: {ID: 313, Key: "woodchestR", Value: `Wooden Chest`},
	314: {ID: 314, Key: "burningbodies", Value: `Pyre of Flesh`},
	315: {ID: 315, Key: "Cloudy", Value: `Cloudy`},
	316: {ID: 316, Key: "sol", Value: `Soul`},
	317: {ID: 317, Key: "Natalya's Shadow", Value: `Natalya's Shadow`},
	318: {ID: 318, Key: "Sander's Derby", Value: `McAuley's Pledge`},
	319: {ID: 319, Key: "Blood Chalice", Value: `Blood Chalice`},
	320: {ID: 320, Key: "FindingDrainLeverDru", Value: `Finally... The drain lever.`},
	321: {ID: 321, Key: "FindingTempleDru", Value: `This temple is a nest of evil.`},
	322: {ID: 322, Key: "burningpit", Value: `Burning Pit`},
	323: {ID: 323, Key: "Druid Plague Poppy", Value: `Plague Poppy`},
	324: {ID: 324, Key: "of Jabbing", Value: `of Jab`},
	325: {ID: 325, Key: "spe", Value: `Spleen`},
	326: {ID: 326, Key: "tribal flag", Value: `Tribal Flag`},
	327: {ID: 327, Key: "eflg", Value: `Town Flag`},
	328: {ID: 328, Key: "Eagleexp", Value: `Eagle`},
	329: {ID: 329, Key: "EskillManaSteal", Value: `Mana Steal: `},
	330: {ID: 330, Key: "Siege Door", Value: `Siege Door`},
	331: {ID: 331, Key: "Eskilllowerresis", Value: `Lowers Resistance `},
	332: {ID: 332, Key: "Skullcollector", Value: `Skull Collector`},
	333: {ID: 333, Key: "Eagleeye", Value: `Eagleeye`},
	334: {ID: 334, Key: "chan", Value: `Chandelier`},
	335: {ID: 335, Key: "jar2", Value: `Jar`},
	336: {ID: 336, Key: "jar3", Value: `Jar`},
	337: {ID: 337, Key: "swingingheads", Value: `Swinging Heads`},
	338: {ID: 338, Key: "LeavingTownAct5Dru", Value: `Baal! Nothing will stand in my way.`},
	339: {ID: 339, Key: "pole", Value: `Pole`},
	340: {ID: 340, Key: "Singing", Value: `Singing`},
	341: {ID: 341, Key: "M'avina's True Sight", Value: `M'avina's True Sight`},
	342: {ID: 342, Key: "of the Kraken", Value: `of the Centaur`},
	343: {ID: 343, Key: "animatedskullsandrocks", Value: `Skulls and Rocks`},
	344: {ID: 344, Key: "Coldkill", Value: `Coldkill`},
	345: {ID: 345, Key: "banner1", Value: `Banner`},
	346: {ID: 346, Key: "banner2", Value: `Banner`},
	347: {ID: 347, Key: "Trang-Oul's Claws", Value: `Trang-Oul's Claws`},
	348: {ID: 348, Key: "tal", Value: `Tail`},
	349: {ID: 349, Key: "Camphor", Value: `Camphor`},
	350: {ID: 350, Key: "pene", Value: `Stash`},
	351: {ID: 351, Key: "object1", Value: `Hidden Stash`},
	352: {ID: 352, Key: "torch2", Value: `Torch`},
	353: {ID: 353, Key: "torch1", Value: `Torch`},
	354: {ID: 354, Key: "tomb3", Value: `Tomb`},
	355: {ID: 355, Key: "Playersubtitles30", Value: `Run away!`},
	356: {ID: 356, Key: "EnteringArcaneDru", Value: `This was not designed by nature's Architect.`},
	357: {ID: 357, Key: "hrb", Value: `Herb`},
	358: {ID: 358, Key: "Odium", Value: `Odium`},
	359: {ID: 359, Key: "tomb2", Value: `Tomb`},
	360: {ID: 360, Key: "of Battle Orders", Value: `of Battle Orders`},
	361: {ID: 361, Key: "tomb1", Value: `Tomb`},
	362: {ID: 362, Key: "Shadowkiller", Value: `Shadow Killer`},
	363: {ID: 363, Key: "ttor", Value: `Torch`},
	364: {ID: 364, Key: "The Reedeemer", Value: `The Redeemer`},
	365: {ID: 365, Key: "clientsmoke", Value: `Smoke`},
	366: {ID: 366, Key: "deadbarbarian18", Value: `Dead Barbarian`},
	367: {ID: 367, Key: "cagedwussie1", Value: `Cage`},
	368: {ID: 368, Key: "icecaveshrine2", Value: `Shrine`},
	369: {ID: 369, Key: "hrn", Value: `Horn`},
	370: {ID: 370, Key: "evilurn", Value: `Evil Urn`},
	371: {ID: 371, Key: "Thin", Value: `Thin`},
	372: {ID: 372, Key: "secret object", Value: `secret object`},
	373: {ID: 373, Key: "Riphook", Value: `Riphook`},
	374: {ID: 374, Key: "Ldeathpole", Value: `Death Pole`},
	375: {ID: 375, Key: "hrt", Value: `Heart`},
	376: {ID: 376, Key: "ice", Value: `Keep it to thaw Anya
Malah's Potion`},
	377: {ID: 377, Key: "Snowflake", Value: `Snowy`},
	378: {ID: 378, Key: "of Blizzard", Value: `of Blizzard`},
	379: {ID: 379, Key: "Qual'Kek's Enforcer", Value: `Qual-Khek's Enforcer`},
	380: {ID: 380, Key: "EnteringTopMountAct5Ams", Value: `The fabled home of the Ancients.`},
	381: {ID: 381, Key: "explodingchest", Value: `Chest`},
	382: {ID: 382, Key: "animated skulland rockpile", Value: `Skulls and Rocks`},
	383: {ID: 383, Key: "jar1", Value: `Jar`},
	384: {ID: 384, Key: "ettr", Value: `Torch`},
	385: {ID: 385, Key: "ecfra", Value: `Fire`},
	386: {ID: 386, Key: "explodingbarrel", Value: `Barrel`},
	387: {ID: 387, Key: "Rite of Passage", Value: `Rite of Passage`},
	388: {ID: 388, Key: "Druid Spirit Wolf", Value: `Spirit Wolf`},
	389: {ID: 389, Key: "object", Value: `Hidden Stash`},
	390: {ID: 390, Key: "Immortal King", Value: `Immortal King`},
	391: {ID: 391, Key: "Shrine2wilderness", Value: `Shrine`},
	392: {ID: 392, Key: "ptox", Value: `Torch Pit`},
	393: {ID: 393, Key: "Siege Control", Value: `Siege Control`},
	394: {ID: 394, Key: "Of the Idiot", Value: `of the Idiot`},
	395: {ID: 395, Key: "Doombringer", Value: `Doombringer`},
	396: {ID: 396, Key: "Wolf", Value: `Wolf`},
	397: {ID: 397, Key: "mrjar", Value: `Hidden Stash`},
	398: {ID: 398, Key: "object2", Value: `Hidden Stash`},
	399: {ID: 399, Key: "Immortal King's Soul Cage ", Value: `Immortal King's Soul Cage`},
	400: {ID: 400, Key: "Minionexp", Value: `Enslaved`},
	401: {ID: 401, Key: "groundtombL", Value: `Tomb`},
	402: {ID: 402, Key: "Marred", Value: `Marred`},
	403: {ID: 403, Key: "Skin of the Flayerd One", Value: `Skin of the Flayed One`},
	404: {ID: 404, Key: "Doomseer", Value: `Doomseer`},
	405: {ID: 405, Key: "Spiritual", Value: `Spiritual`},
	406: {ID: 406, Key: "MinionIce/fireBoarSpawner", Value: `Demon Portal`},
	407: {ID: 407, Key: "A5Q6SuccessfulLarzuk", Value: `74
The Ancients themselves will envy our 
songs about you.
 
Please, don't forget about us! Farewell, 
my friend.
`},
	408: {ID: 408, Key: "Psychic", Value: `Psychic`},
	409: {ID: 409, Key: "Calling", Value: `Calling`},
	410: {ID: 410, Key: "Venomous", Value: `Venomous`},
	411: {ID: 411, Key: "Medusa's Gaze", Value: `Medusa's Gaze`},
	412: {ID: 412, Key: "Telling of Beads", Value: `Telling of Beads`},
	413: {ID: 413, Key: "Playersubtitles29", Value: `Retreat!`},
	414: {ID: 414, Key: "sbub", Value: `Shrub`},
	415: {ID: 415, Key: "ubub", Value: `Shrub`},
	416: {ID: 416, Key: "KillingdDiabloNec", Value: `Lord Diablo I have bested you.`},
	417: {ID: 417, Key: "deadperson2", Value: `Corpse`},
	418: {ID: 418, Key: "Eskillthunder1", Value: `lightning damage: `},
	419: {ID: 419, Key: "ESkillTimes", Value: ` Times`},
	420: {ID: 420, Key: "Eskillthunder2", Value: `nova damage: `},
	421: {ID: 421, Key: "Eskillthunder3", Value: `charged bolt damage: `},
	422: {ID: 422, Key: "Eaglewind", Value: `Eaglewind`},
	423: {ID: 423, Key: "Prison Door", Value: `Prison Door`},
	424: {ID: 424, Key: "ancientsaltar", Value: `Altar of the Heavens`},
	425: {ID: 425, Key: "hiddenstash", Value: `Hidden Stash`},
	426: {ID: 426, Key: "A5Q6SuccessfulAnya", Value: `53
You have done the impossible, hero. 
Your defeat of the last of the three 
Prime Evils is a great victory for the 
Light. 
 
Strange that you say that the 
Worldstone must be destroyed. The 
prophecies said nothing about that.
 
Perhaps all we have fought for will be 
lost...or perhaps we'll never need fight 
again!
`},
	427: {ID: 427, Key: "Dawnbringer", Value: `Dawn Bringer`},
	428: {ID: 428, Key: "Sacred", Value: `Sacred`},
	429: {ID: 429, Key: "of Amplify Damage", Value: `of Amplify Damage`},
	430: {ID: 430, Key: "A5Q2SuccessfulLarzuk", Value: `55
Since your arrival, Cain has spoken of 
your deeds in the Southern Kingdoms, 
defeating both Mephisto and Diablo. At 
first, I scoffed at his tales, but I'm 
finding them more believable with every 
passing day.
`},
	431: {ID: 431, Key: "eweaponrackR", Value: `Weapon Rack`},
	432: {ID: 432, Key: "of Frost Novas", Value: `of Frost Nova`},
	433: {ID: 433, Key: "Plague Bearer", Value: `Plague Bearer`},
	434: {ID: 434, Key: "A5Q2EarlyReturnAnya", Value: `115
Well, you found me on the mountain; I'm 
sure you'll find them too.
`},
	435: {ID: 435, Key: "Aldur's Gauntlet", Value: `Aldur's Rhythm`},
	436: {ID: 436, Key: "icecavejar2", Value: `Jar`},
	437: {ID: 437, Key: "CompletingBurialDru", Value: `Your time is past, Blood Raven.`},
	438: {ID: 438, Key: "Of the Maggot", Value: `of Maggots`},
	439: {ID: 439, Key: "icecavejar4", Value: `Jar`},
	440: {ID: 440, Key: "of Volcano", Value: `of Volcano`},
	441: {ID: 441, Key: "icecavejar3", Value: `Jar`},
	442: {ID: 442, Key: "icecavejar1", Value: `Jar`},
	443: {ID: 443, Key: "earmorstandL", Value: `Armor Stand`},
	444: {ID: 444, Key: "earmorstandR", Value: `Armor Stand`},
	445: {ID: 445, Key: "qstsa5q1", Value: `Siege on Harrogath`},
	446: {ID: 446, Key: "qstsa5q2", Value: `Rescue on Mount Arreat`},
	447: {ID: 447, Key: "qstsa5q21", Value: `Find the Soldiers in the Frigid Highlands.`},
	448: {ID: 448, Key: "Falconeye", Value: `Falconeye`},
	449: {ID: 449, Key: "Malicious", Value: `Malicious`},
	450: {ID: 450, Key: "qstsa5q22", Value: `Rescue %d more Soldiers in the Frigid Highlands.`},
	451: {ID: 451, Key: "Pain Worm1", Value: `Pain Worm`},
	452: {ID: 452, Key: "Hazy", Value: `Hazy`},
	453: {ID: 453, Key: "Pain Worm2", Value: `Torment Worm`},
	454: {ID: 454, Key: "Static", Value: `Static`},
	455: {ID: 455, Key: "of Fire Bolts", Value: `of Fire Bolt`},
	456: {ID: 456, Key: "Baezil's Vortex", Value: `Baezil's Vortex`},
	457: {ID: 457, Key: "of the Elephant", Value: `of the Elephant`},
	458: {ID: 458, Key: "EnteringTopMountAct5Dru", Value: `At last...The summit of Mount Arreat.`},
	459: {ID: 459, Key: "Pain Worm3", Value: `Agony Worm`},
	460: {ID: 460, Key: "Pain Worm4", Value: `Menace Worm`},
	461: {ID: 461, Key: "Pain Worm5", Value: `Anguish Worm`},
	462: {ID: 462, Key: "gate", Value: `Main Gate`},
	463: {ID: 463, Key: "debris", Value: `Debris`},
	464: {ID: 464, Key: "qstsa5q23", Value: `Return to Qual-Kehk for your reward.`},
	465: {ID: 465, Key: "A5Q3EarlyReturnCain", Value: `60
Nihlathak's story does sound 
reasonable, considering what I've 
heard about Anya. However, the best 
lies are often hidden within truth.
`},
	466: {ID: 466, Key: "qstsa5q24", Value: `Apply the Runes to a Socketed item in this order:`},
	467: {ID: 467, Key: "Lasting", Value: `Lasting`},
	468: {ID: 468, Key: "EnterMonasteryDru", Value: `Evil flows from here.`},
	469: {ID: 469, Key: "Artificer's", Value: `Artisan's`},
	470: {ID: 470, Key: "Altar", Value: `Altar`},
	471: {ID: 471, Key: "qstsa5q3", Value: `Prison of Ice`},
	472: {ID: 472, Key: "qstsa5q31", Value: `Look for Anya under the Crystalline Passage by the Frozen River.`},
	473: {ID: 473, Key: "FireBoar", Value: `Fire Boar`},
	474: {ID: 474, Key: "qstsa5q32", Value: `Talk to Malah.`},
	475: {ID: 475, Key: "EnteringWorldstoneAct5Pal", Value: `The Worldstone! Praise the Light.`},
	476: {ID: 476, Key: "icecave_torch2", Value: `Torch`},
	477: {ID: 477, Key: "Fleshripper", Value: `Fleshripper`},
	478: {ID: 478, Key: "Wihtstan's Guard", Value: `Whitstan's Guard`},
	479: {ID: 479, Key: "icecave_torch1", Value: `Torch`},
	480: {ID: 480, Key: "qstsa5q33", Value: `Use the potion Malah gave you to thaw Anya.`},
	481: {ID: 481, Key: "Ondal's Almighty", Value: `Ondal's Almighty`},
	482: {ID: 482, Key: "Cliffkiller", Value: `Cliffkiller`},
	483: {ID: 483, Key: "Frantic", Value: `Frantic`},
	484: {ID: 484, Key: "qstsa5q34", Value: `Talk to Malah for your reward.`},
	485: {ID: 485, Key: "qstsa5q35", Value: `Talk to Anya.`},
	486: {ID: 486, Key: "BeginTaintedSunAss", Value: `An eclipse... never a good omen.`},
	487: {ID: 487, Key: "qstsa5q4", Value: `Betrayal of Harrogath`},
	488: {ID: 488, Key: "qstsa5q41", Value: `Take Anya's portal to Nihlathak's Temple.`},
	489: {ID: 489, Key: "qstsa5q42", Value: `Kill Nihlathak.`},
	490: {ID: 490, Key: "Shivering", Value: `Shivering`},
	491: {ID: 491, Key: "Athena's Wrath", Value: `Athena's Wrath`},
	492: {ID: 492, Key: "qstsa5q43", Value: `Talk to Anya before you continue through the Crystalline Passage, past the Glacial Trail, to proceed up Mount Arreat to the Summit.`},
	493: {ID: 493, Key: "deathpole", Value: `Death Pole`},
	494: {ID: 494, Key: "mrpole", Value: `Pole`},
	495: {ID: 495, Key: "deadbarbarian", Value: `Dead Barbarian`},
	496: {ID: 496, Key: "qstsa5q5", Value: `Rite of Passage`},
	497: {ID: 497, Key: "Khalim's Vengance", Value: `Khalim's Vengeance`},
	498: {ID: 498, Key: "qstsa5q51", Value: `Travel through the Ancient's Way to find the Ancients at the Arreat Summit.`},
	499: {ID: 499, Key: "qstsa5q52", Value: `Consult with the Ancients by activating the Altar of the Heavens.`},
	500: {ID: 500, Key: "qstsa5q53", Value: `Defeat all three Ancients without leaving Mount Arreat.`},
	501: {ID: 501, Key: "qstsa5q6", Value: `Eve of Destruction`},
	502: {ID: 502, Key: "qstsa5q61", Value: `Consult with the Ancients.`},
	503: {ID: 503, Key: "qstsa5q62", Value: `Kill Baal.`},
	504: {ID: 504, Key: "Of Sweetness", Value: `of Sweetness`},
	505: {ID: 505, Key: "of Fast Repair", Value: `of Restoration`},
	506: {ID: 506, Key: "Seraph's Hymn", Value: `Seraph's Hymn`},
	507: {ID: 507, Key: "Magnus' Skin", Value: `Magnus' Skin`},
	508: {ID: 508, Key: "Zakrum's Hand", Value: `Zakarum's Hand`},
	509: {ID: 509, Key: "Leviathan", Value: `Leviathan`},
	510: {ID: 510, Key: "Obsidian", Value: `Obsidian`},
	511: {ID: 511, Key: "Demonweb", Value: `Demonweb`},
	512: {ID: 512, Key: "Parkersor's Calm", Value: `Parkersor's Calm`},
	513: {ID: 513, Key: "Mang Song's Lesson", Value: `Mang Song's Lesson`},
	514: {ID: 514, Key: "The Prowler", Value: `The Prowler`},
	515: {ID: 515, Key: "Eyeback Unleashed", Value: `Eyeback the Unleashed`},
	516: {ID: 516, Key: "qstsa5q63", Value: `Talk to Tyreal.`},
	517: {ID: 517, Key: "qstsa5q64", Value: `Take Tyreal's Portal to Safety.`},
	518: {ID: 518, Key: "Harrogath", Value: `Harrogath`},
	519: {ID: 519, Key: "Bloody Foothills", Value: `Bloody Foothills`},
	520: {ID: 520, Key: "Bonesaw Breaker", Value: `Bonesaw Breaker`},
	521: {ID: 521, Key: "of Lightning Javelin", Value: `of Lightning Javelin`},
	522: {ID: 522, Key: "Arreat Plateau", Value: `Arreat Plateau`},
	523: {ID: 523, Key: "Of self-repair", Value: `Self-Repair`},
	524: {ID: 524, Key: "Silkweave", Value: `Silkweave`},
	525: {ID: 525, Key: "of Swords", Value: `of Swords`},
	526: {ID: 526, Key: "Cellar of Pity", Value: `Frozen River`},
	527: {ID: 527, Key: "Echo Chamber", Value: `Drifter Cavern`},
	528: {ID: 528, Key: "Glacial Caves Level 1", Value: `The Ancients' Way`},
	529: {ID: 529, Key: "Glacial Caves Level 2", Value: `Icy Cellar`},
	530: {ID: 530, Key: "ChestSL", Value: `Chest`},
	531: {ID: 531, Key: "deadperson", Value: `Corpse`},
	532: {ID: 532, Key: "Nihlathaks Temple", Value: `Nihlathak's Temple`},
	533: {ID: 533, Key: "Atma's Scarab", Value: `Atma's Scarab`},
	534: {ID: 534, Key: "Arcing", Value: `Arcing`},
	535: {ID: 535, Key: "Of Baddass", Value: `of the Abyss`},
	536: {ID: 536, Key: "of Clay Golem Summoning", Value: `of Clay Golem`},
	537: {ID: 537, Key: "of Replenishing", Value: `of Replenishing`},
	538: {ID: 538, Key: "Sinblade", Value: `Sinblade`},
	539: {ID: 539, Key: "Eschuta's temper", Value: `Eschuta's Temper`},
	540: {ID: 540, Key: "Eschuta's temper", Value: `Eschuta's Temper`},
	541: {ID: 541, Key: "Baal Taunt", Value: `Baal`},
	542: {ID: 542, Key: "Halls of Anguish", Value: `Halls of Anguish`},
	543: {ID: 543, Key: "Halls of Tormented Insanity", Value: `Halls of Torment`},
	544: {ID: 544, Key: "Royal", Value: `Royal`},
	545: {ID: 545, Key: "Widowmaker", Value: `Widowmaker`},
	546: {ID: 546, Key: "woodchest2L", Value: `Wooden Chest`},
	547: {ID: 547, Key: "Halls of Vaught", Value: `Halls of Vaught`},
	548: {ID: 548, Key: "The Worldstone Chamber", Value: `The Worldstone Chamber`},
	549: {ID: 549, Key: "Throne of Destruction", Value: `Throne of Destruction`},
	550: {ID: 550, Key: "Corporal", Value: `Corporal`},
	551: {ID: 551, Key: "of Lightning Fury", Value: `of Lightning Fury`},
	552: {ID: 552, Key: "Stone Crusher", Value: `Stone Crusher`},
	553: {ID: 553, Key: "Barricade Wall Left", Value: `Barricade`},
	554: {ID: 554, Key: "EnteringRadamentAss", Value: `Why must evil hide in such wretched places?`},
	555: {ID: 555, Key: "of Blazing", Value: `of Blazing`},
	556: {ID: 556, Key: "of Ennui", Value: `of Ennui`},
	557: {ID: 557, Key: "ItemExpansiveChanc1", Value: `%d%% Chance to cast level %d %s on striking`},
	558: {ID: 558, Key: "ItemExpansiveChanc2", Value: `%d%% Chance to cast level %d %s when struck`},
	559: {ID: 559, Key: "FindingdecoyTombDru", Value: `These Horadric markings are mysterious.`},
	560: {ID: 560, Key: "woodchest2R", Value: `Wooden Chest`},
	561: {ID: 561, Key: "To The Bloody Foothills", Value: `To The Bloody Foothills`},
	562: {ID: 562, Key: "tr2", Value: `Right Click to Cast
Scroll of Resistance`},
	563: {ID: 563, Key: "of Firebolts", Value: `of Firebolts`},
	564: {ID: 564, Key: "To The Arreat Plateau", Value: `To The Arreat Plateau`},
	565: {ID: 565, Key: "of Frozen Orb", Value: `of Frozen Orb`},
	566: {ID: 566, Key: "of Molten Boulders", Value: `of Molten Boulder`},
	567: {ID: 567, Key: "To The Crystalized Cavern Level 1", Value: `To The Crystalline Passage`},
	568: {ID: 568, Key: "To The Cellar of Pity", Value: `To The Frozen River`},
	569: {ID: 569, Key: "To The Crystalized Cavern Level 2", Value: `To The Glacial Trail`},
	570: {ID: 570, Key: "To The Echo Chamber", Value: `To The Drifter Cavern`},
	571: {ID: 571, Key: "To The Tundra Wastelands", Value: `To The Frozen Tundra`},
	572: {ID: 572, Key: "Shadow Warrior", Value: `Shadow Warrior`},
	573: {ID: 573, Key: "To The Glacier Caves Level 1", Value: `To The Ancients' Way`},
	574: {ID: 574, Key: "RescueQual-KehkAct5Bar", Value: `Follow me.`},
	575: {ID: 575, Key: "Bul-Kathos' Children", Value: `Bul-Kathos' Children`},
	576: {ID: 576, Key: "To The Glacier Caves Level 2", Value: `To The Icy Cellar`},
	577: {ID: 577, Key: "To Nihlathaks Temple", Value: `To Nihlathak's Temple`},
	578: {ID: 578, Key: "To The Halls of Anguish", Value: `To The Halls of Anguish`},
	579: {ID: 579, Key: "of Wrath", Value: `of Wrath`},
	580: {ID: 580, Key: "To The Halls of Death's Calling", Value: `To The Halls of Pain`},
	581: {ID: 581, Key: "To The Halls of Tormented Insanity", Value: `To The Halls of Torment`},
	582: {ID: 582, Key: "of Immolating Arrows", Value: `of Immolating Arrow`},
	583: {ID: 583, Key: "Bonescapel", Value: `Bonescalpel`},
	584: {ID: 584, Key: "To The Halls of Vaught", Value: `To The Halls of Vaught`},
	585: {ID: 585, Key: "of Poison Explosion", Value: `of Poison Explosion`},
	586: {ID: 586, Key: "Loud", Value: `Alarming`},
	587: {ID: 587, Key: "Blackbog's Sharp", Value: `Blackbog's Sharp`},
	588: {ID: 588, Key: "Grim's Burning Dead", Value: `Grim's Burning Dead`},
	589: {ID: 589, Key: "To The Worldstone Keep Level 1", Value: `To The Worldstone Keep Level 1`},
	590: {ID: 590, Key: "of Teeth", Value: `of Teeth`},
	591: {ID: 591, Key: "Titan's Revenge", Value: `Titan's Revenge`},
	592: {ID: 592, Key: "Stormwillow", Value: `Stormwillow`},
	593: {ID: 593, Key: "Vine Creature", Value: `Vine Creature`},
	594: {ID: 594, Key: "Rainbow", Value: `Rainbow`},
	595: {ID: 595, Key: "To The Worldstone Keep Level 2", Value: `To The Worldstone Keep Level 2`},
	596: {ID: 596, Key: "To The Worldstone Keep Level 3", Value: `To The Worldstone Keep Level 3`},
	597: {ID: 597, Key: "hireiconinfo1", Value: `Drop Potion on Portrait to Heal`},
	598: {ID: 598, Key: "of Evisceration", Value: `of Evisceration`},
	599: {ID: 599, Key: "Of Nobility", Value: `of Nobility`},
	600: {ID: 600, Key: "hireiconinfo2", Value: `Right-click to Open Inventory (%s)`},
	601: {ID: 601, Key: "hiredismisshire", Value: `Dismiss Hireling`},
	602: {ID: 602, Key: "hirerehire", Value: `Rehire`},
	603: {ID: 603, Key: "Bing Sz Wang", Value: `Bing Sz Wang`},
	604: {ID: 604, Key: "Ribcracker", Value: `Ribcracker`},
	605: {ID: 605, Key: "hireresurrect2", Value: `Resurrect %s: %d`},
	606: {ID: 606, Key: "Endlesshail", Value: `Endlesshail`},
	607: {ID: 607, Key: "Trang-Oul's Mask", Value: `Trang-Oul's Mask`},
	608: {ID: 608, Key: "uap", Value: `Shako`},
	609: {ID: 609, Key: "RescueQual-KehkAct5Ass", Value: `Follow me.`},
	610: {ID: 610, Key: "uar", Value: `Sacred Armor`},
	611: {ID: 611, Key: "of Blizzards", Value: `of Blizzard`},
	612: {ID: 612, Key: "hirechat1", Value: `Thanks.`},
	613: {ID: 613, Key: "jar", Value: `Jar`},
	614: {ID: 614, Key: "hirechat2", Value: `Thank you.`},
	615: {ID: 615, Key: "Stinging", Value: `Stinging`},
	616: {ID: 616, Key: "Expert's", Value: `Expert's`},
	617: {ID: 617, Key: "HoradricMalusDru", Value: `Charsi will be thankful to get this Malus.`},
	618: {ID: 618, Key: "jaw", Value: `Jawbone`},
	619: {ID: 619, Key: "The Archon Magus", Value: `The Achron Magus`},
	620: {ID: 620, Key: "hirechat3", Value: `I needed that.`},
	621: {ID: 621, Key: "hiredanger1", Value: `I sense danger.`},
	622: {ID: 622, Key: "Of Stone", Value: `of Stone`},
	623: {ID: 623, Key: "qsta5q11", Value: `Stop the Siege by killing Shenk the Overseer in the Bloody Foothills.`},
	624: {ID: 624, Key: "Sanguine", Value: `Sanguine`},
	625: {ID: 625, Key: "EnteringArcaneAss", Value: `The Sanctuary - Horazon's obsession.`},
	626: {ID: 626, Key: "qsta5q14", Value: `Larzuk will add sockets to the item of your choice.`},
	627: {ID: 627, Key: "ESkillWolf", Value: `Wolf: `},
	628: {ID: 628, Key: "Baal Crab", Value: `Baal`},
	629: {ID: 629, Key: "qsta5q12", Value: `Kill Shenk the Overseer.`},
	630: {ID: 630, Key: "qsta5q13", Value: `Go talk to Larzuk for your reward.`},
	631: {ID: 631, Key: "of Ice Bolts", Value: `of Ice Bolt`},
	632: {ID: 632, Key: "of Winter", Value: `of Winter`},
	633: {ID: 633, Key: "Gore Ripper", Value: `Gore Ripper`},
	634: {ID: 634, Key: "strAssassinOnly", Value: `(Assassin Only)`},
	635: {ID: 635, Key: "hiredanger2", Value: `I hate these vermin.`},
	636: {ID: 636, Key: "A5Q5AfterInitQualKehk", Value: `63
The Ancients are not our enemies. 
Remember that! They are our 
ancestors -- our gods.
`},
	637: {ID: 637, Key: "ucl", Value: `Loricated Mail`},
	638: {ID: 638, Key: "hiredanger3", Value: `I have a bad feeling about this.`},
	639: {ID: 639, Key: "of Sacrifice", Value: `of Sacrifice`},
	640: {ID: 640, Key: "Nagas", Value: `Nagas`},
	641: {ID: 641, Key: "hiredanger4", Value: `Beware!`},
	642: {ID: 642, Key: "Visionary", Value: `Visionary`},
	643: {ID: 643, Key: "Buriza-Do Kyanon", Value: `Buriza-Do Kyanon`},
	644: {ID: 644, Key: "Hellwarden's Husk", Value: `Hell Warden's Husk`},
	645: {ID: 645, Key: "of Cold Arrows", Value: `of Cold Arrow`},
	646: {ID: 646, Key: "hiredanger5", Value: `I detest spiders.`},
	647: {ID: 647, Key: "Baal Subject 6a", Value: `The Baker`},
	648: {ID: 648, Key: "Baal Subject 6b", Value: `The Candlestick Maker`},
	649: {ID: 649, Key: "hiredanger6", Value: `Eek, snakes!`},
	650: {ID: 650, Key: "hirefeelstronger2", Value: `I am more experienced.`},
	651: {ID: 651, Key: "ItemExpcharmdesc", Value: `Keep in inventory to gain bonus`},
	652: {ID: 652, Key: "Raging", Value: `Raging`},
	653: {ID: 653, Key: "ESkillStars", Value: `Stars: `},
	654: {ID: 654, Key: "of Battle Command", Value: `of Battle Command`},
	655: {ID: 655, Key: "flag wilderness", Value: `Flag`},
	656: {ID: 656, Key: "To Harrogath", Value: `To Harrogath`},
	657: {ID: 657, Key: "uea", Value: `Wyrmhide`},
	658: {ID: 658, Key: "hiregreets1", Value: `Good morning.`},
	659: {ID: 659, Key: "of Multiple Shot", Value: `of Multiple Shot`},
	660: {ID: 660, Key: "Xenos", Value: `Xenos`},
	661: {ID: 661, Key: "Of the Choir", Value: `Of the Choir`},
	662: {ID: 662, Key: "Harlequin Crest", Value: `Harlequin Crest`},
	663: {ID: 663, Key: "A5Q6SuccessfulMalah", Value: `48
If Tyrael says the Worldstone must be 
destroyed, then it must. We cannot let 
Baal's corruption prevail!
 
The world will change, true -- but who is 
to say it isn't for the better?
`},
	664: {ID: 664, Key: "WolfRider2", Value: `Carver Wolfrider`},
	665: {ID: 665, Key: "uh9", Value: `Bone Visage`},
	666: {ID: 666, Key: "WolfRider4", Value: `Darkone Wolfrider`},
	667: {ID: 667, Key: "Septic", Value: `Septic`},
	668: {ID: 668, Key: "Hightower's Watch", Value: `Hightower's Watch`},
	669: {ID: 669, Key: "CompletingTempleDru", Value: `There is hope once again.`},
	670: {ID: 670, Key: "WolfRider5", Value: `Warped Fallen Wolfrider`},
	671: {ID: 671, Key: "WolfRider3", Value: `Devilkin Wolfrider`},
	672: {ID: 672, Key: "WolfRider1", Value: `Fallen Wolfrider`},
	673: {ID: 673, Key: "Musty", Value: `Musty`},
	674: {ID: 674, Key: "Of Bile", Value: `of Bile`},
	675: {ID: 675, Key: "EnterMonasteryAss", Value: `Such corruption in this place...`},
	676: {ID: 676, Key: "ChestSR", Value: `Chest`},
	677: {ID: 677, Key: "pileofskullsandrocks", Value: `Skulls and Rocks`},
	678: {ID: 678, Key: "red light", Value: `Red Light`},
	679: {ID: 679, Key: "of Nova", Value: `of Nova`},
	680: {ID: 680, Key: "EnteringTopMountAct5Nec", Value: `The resting place of the Ancients...`},
	681: {ID: 681, Key: "To The Throne of Destruction", Value: `To The Throne of Destruction`},
	682: {ID: 682, Key: "A5Q5SuccessfulAnya", Value: `82
You stand before me a worthy hero -- 
and on you rests the last hope of our 
people.
 
Bear it well, warrior.
`},
	683: {ID: 683, Key: "jew", Value: `Jewel`},
	684: {ID: 684, Key: "Communal", Value: `Communal`},
	685: {ID: 685, Key: "mrbox", Value: `Box`},
	686: {ID: 686, Key: "hirepraise1", Value: `It is good to work for someone who cares.`},
	687: {ID: 687, Key: "Druid Totem", Value: `Druid Spirit`},
	688: {ID: 688, Key: "A5Q1EarlyReturnNihlathak", Value: `78
What are you still doing here? I thought 
you were going off to die.
 
Go...Be quick about it.
`},
	689: {ID: 689, Key: "Sazabi's Cobalt Redeemer", Value: `Sazabi's Cobalt Redeemer`},
	690: {ID: 690, Key: "hirepraise2", Value: `Good for you.`},
	691: {ID: 691, Key: "hiregreets2", Value: `Good afternoon.`},
	692: {ID: 692, Key: "hiregreets3", Value: `Good evening.`},
	693: {ID: 693, Key: "hiregreets4", Value: `Hello.`},
	694: {ID: 694, Key: "CfgSkill9", Value: `Skill 9`},
	695: {ID: 695, Key: "CfgSkill10", Value: `Skill 10`},
	696: {ID: 696, Key: "CfgSkill11", Value: `Skill 11`},
	697: {ID: 697, Key: "FindingInifusAss", Value: `How has this tree escaped corruption?`},
	698: {ID: 698, Key: "Baal Throne", Value: `Baal`},
	699: {ID: 699, Key: "A5Q4EarlyReturnAnya", Value: `55
Nihlathak is a traitor! If Baal gets the 
Relic, he shall enter the mountain and 
wreak havoc there!
 
I cannot believe that Nihlathak would 
give the Relic to Baal in any kind of 
trade. He's truly insane if he believes 
that he can deal with the Lord of 
Destruction.
`},
	700: {ID: 700, Key: "ModStre10L", Value: `Percent Chance to Cast`},
	701: {ID: 701, Key: "EnterJailAss", Value: `Try and cage me, demons.`},
	702: {ID: 702, Key: "CfgSkill12", Value: `Skill 12`},
	703: {ID: 703, Key: "CfgSkill13", Value: `Skill 13`},
	704: {ID: 704, Key: "CfgSkill14", Value: `Skill 14`},
	705: {ID: 705, Key: "CfgSkill15", Value: `Skill 15`},
	706: {ID: 706, Key: "uhb", Value: `Myrmidon Greaves`},
	707: {ID: 707, Key: "uhc", Value: `Colossus Girdle`},
	708: {ID: 708, Key: "Hexing", Value: `Hexing`},
	709: {ID: 709, Key: "Sander's Superstition", Value: `McAuley's Superstition`},
	710: {ID: 710, Key: "CfgSkill16", Value: `Skill 16`},
	711: {ID: 711, Key: "uhg", Value: `Ogre Gauntlets`},
	712: {ID: 712, Key: "CfgToggleminimap", Value: `Toggle MiniMap`},
	713: {ID: 713, Key: "Tal Rasha's Wrappings", Value: `Tal Rasha's Wrappings`},
	714: {ID: 714, Key: "of Potion Finding", Value: `of Find Potion`},
	715: {ID: 715, Key: "Cfgswapweapons", Value: `Swap Weapons`},
	716: {ID: 716, Key: "uhl", Value: `Giant Conch`},
	717: {ID: 717, Key: "uhm", Value: `Spired Helm`},
	718: {ID: 718, Key: "uhn", Value: `Boneweave`},
	719: {ID: 719, Key: "EnterBurialAss", Value: `Whose handiwork lies buried here?`},
	720: {ID: 720, Key: "ModStre10a", Value: `Charges`},
	721: {ID: 721, Key: "ModStre10b", Value: `Level`},
	722: {ID: 722, Key: "ModStre10c", Value: `Per Level`},
	723: {ID: 723, Key: "ModStre10d", Value: `(%d/%d Charges)`},
	724: {ID: 724, Key: "ModStre10e", Value: `(`},
	725: {ID: 725, Key: "ModStre10f", Value: `)`},
	726: {ID: 726, Key: "ModStre10g", Value: `Stealth`},
	727: {ID: 727, Key: "of the Phoenix", Value: `of the Phoenix`},
	728: {ID: 728, Key: "ModStre10h", Value: `Immunity to Poison`},
	729: {ID: 729, Key: "ModStre10i", Value: `Cursed`},
	730: {ID: 730, Key: "A5Q1EarlyReturnCain", Value: `105
I understand your reluctance, but now 
is the time to strike.
`},
	731: {ID: 731, Key: "Stinky", Value: `Odiferous`},
	732: {ID: 732, Key: "Lion Branded", Value: `Lion Branded`},
	733: {ID: 733, Key: "ModStre10j", Value: `Per Player in Party`},
	734: {ID: 734, Key: "Husoldal Evo", Value: `Husoldal Evo`},
	735: {ID: 735, Key: "Magewrath", Value: `Magewrath`},
	736: {ID: 736, Key: "Kang's Virtue", Value: `Kang's Virtue`},
	737: {ID: 737, Key: "ModStre10k", Value: `Kick Damage:`},
	738: {ID: 738, Key: "BaalColdMage", Value: `Death Mage`},
	739: {ID: 739, Key: "FindingBeneathCityDru", Value: `This smells worse than the sewers of Lut Gohlein.`},
	740: {ID: 740, Key: "uit", Value: `Monarch`},
	741: {ID: 741, Key: "Eskilltomeleeattacks", Value: ` to melee attacks`},
	742: {ID: 742, Key: "of Razors", Value: `of Razors`},
	743: {ID: 743, Key: "Tin", Value: `Tin`},
	744: {ID: 744, Key: "Of the Sniper", Value: `Of the Sniper`},
	745: {ID: 745, Key: "Erion's Bonehandle", Value: `Erion's Bonehandle`},
	746: {ID: 746, Key: "Go South", Value: `return to hell`},
	747: {ID: 747, Key: "Of Dusk", Value: `of Dusk`},
	748: {ID: 748, Key: "Apothecary's Tote", Value: `Apothecary's Tote`},
	749: {ID: 749, Key: "Alma's Reflection", Value: `Alma's Reflection`},
	750: {ID: 750, Key: "Magma Torquer", Value: `Magma Torquer`},
	751: {ID: 751, Key: "Vinvear Molech", Value: `Vinvear Molech`},
	752: {ID: 752, Key: "of Sunlight", Value: `of Sunlight`},
	753: {ID: 753, Key: "A4Q2ExpansionSuccessTyrael", Value: `40
Praise be to the Light! You have 
accomplished the impossible!
 
Diablo and Mephisto have been 
banished back into the Black Abyss 
that spawned them, and the corrupted 
Soulstones are no more.
 
However, while you were fighting here, 
Baal remained behind in the mortal 
realm, building an army of hellish 
minions. Now, Baal's army is searching 
for the Worldstone, the ancient source 
of all the Soulstones and their power, 
while leaving behind a wake of 
destruction. They have forged deeply 
into the Barbarian homelands, heading 
directly for the summit of Mount 
Arreat!
 
Baal knows, mortal hero! That is the 
very site of the blessed Worldstone!
 
Now, enter the portal I have opened for 
you. It will take you to the Barbarian 
city of Harrogath, the last bastion of 
Order on the slopes of Arreat.
`},
	754: {ID: 754, Key: "EskillKickPlur", Value: ` Kicks`},
	755: {ID: 755, Key: "StrSklTree30", Value: `Traps`},
	756: {ID: 756, Key: "RotWalker", Value: `Rot Walker`},
	757: {ID: 757, Key: "Que-Hegan's Wisdon", Value: `Que-Hegan's Wisdom`},
	758: {ID: 758, Key: "StrSklTree31", Value: `Shadow`},
	759: {ID: 759, Key: "StrSklTree32", Value: `Disciplines`},
	760: {ID: 760, Key: "StrSklTree33", Value: `Martial`},
	761: {ID: 761, Key: "FindingdecoyTombAss", Value: `The Horadrim have left their mark here.`},
	762: {ID: 762, Key: "StrSklTree34", Value: `Arts`},
	763: {ID: 763, Key: "The Harbinger", Value: `The Harbinger`},
	764: {ID: 764, Key: "candles", Value: `Candles`},
	765: {ID: 765, Key: "EnteringNihlathakAct5Pal", Value: `By the Light! What is this place?`},
	766: {ID: 766, Key: "Cfghireling", Value: `Hireling Screen`},
	767: {ID: 767, Key: "of Inertia", Value: `of Inertia`},
	768: {ID: 768, Key: "ukp", Value: `Hydraskull`},
	769: {ID: 769, Key: "ula", Value: `Scarab Husk`},
	770: {ID: 770, Key: "ulb", Value: `Wyrmhide Boots`},
	771: {ID: 771, Key: "ulc", Value: `Spiderweb Sash`},
	772: {ID: 772, Key: "uld", Value: `Kraken Shell`},
	773: {ID: 773, Key: "Demon Machine", Value: `Demon Machine`},
	774: {ID: 774, Key: "Spire of Honor", Value: `Spire of Honor`},
	775: {ID: 775, Key: "ulg", Value: `Bramble Mitts`},
	776: {ID: 776, Key: "Festering", Value: `Festering`},
	777: {ID: 777, Key: "Bloodmoon", Value: `Bloodmoon`},
	778: {ID: 778, Key: "Immortal King's Forge", Value: `Immortal King's Forge`},
	779: {ID: 779, Key: "CompletingBladeDru", Value: `Ormus may know something about this unusual blade.`},
	780: {ID: 780, Key: "uncle f#%* comedy central(c)", Value: `Moe`},
	781: {ID: 781, Key: "ulm", Value: `Armet`},
	782: {ID: 782, Key: "Accursed", Value: `Cursing`},
	783: {ID: 783, Key: "of Raise Skeletons", Value: `of Raise Skeleton`},
	784: {ID: 784, Key: "IceSpawn", Value: `Ice Spawn`},
	785: {ID: 785, Key: "flag widlerness", Value: `Flag`},
	786: {ID: 786, Key: "umb", Value: `Boneweave Boots`},
	787: {ID: 787, Key: "umc", Value: `Mithril Coil`},
	788: {ID: 788, Key: "ult", Value: `Hellforge Plate`},
	789: {ID: 789, Key: "Warlord's Trust", Value: `Warlord's Trust`},
	790: {ID: 790, Key: "Spellsteel", Value: `Spellsteel`},
	791: {ID: 791, Key: "umg", Value: `Vambraces`},
	792: {ID: 792, Key: "ItemExpansiveChancX", Value: `%d%% Chance to cast level %d %s on attack`},
	793: {ID: 793, Key: "Of the Maiden", Value: `Of the Maiden`},
	794: {ID: 794, Key: "MiniPanelHire", Value: `Hireling`},
	795: {ID: 795, Key: "NihlathakAct5IntroAssGossip1", Value: `54
Well, well...An Assassin!
 
Heh, heh...While I am sure we are all 
grateful for your presence in our 
troubled town, you need not have made 
the journey.
 
I can personally say that your skills are 
not required here. You would serve 
your clan better elsewhere.
`},
	796: {ID: 796, Key: "uml", Value: `Luna`},
	797: {ID: 797, Key: "Gravepalm", Value: `Gravepalm`},
	798: {ID: 798, Key: "StrSklTree26", Value: `Summoning`},
	799: {ID: 799, Key: "Spiritual Custodian", Value: `Dark Adherent`},
	800: {ID: 800, Key: "StrSklTree27", Value: `Shape`},
	801: {ID: 801, Key: "StrSklTree28", Value: `Shifting`},
	802: {ID: 802, Key: "StrSklTree29", Value: `Elemental`},
	803: {ID: 803, Key: "x", Value: `EXPANSION`},
	804: {ID: 804, Key: "x", Value: `EXPANSION CHARACTER`},
	805: {ID: 805, Key: "x", Value: `CONVERT TO`},
	806: {ID: 806, Key: "Eskillpudmana", Value: ` percent life and mana stealing `},
	807: {ID: 807, Key: "ung", Value: `Diamond Mail`},
	808: {ID: 808, Key: "x", Value: `CLICK OK TO CONVERT YOUR ORIGINAL DIABLO II CHARACTER TO PLAY IN EXPANSION GAMES.`},
	809: {ID: 809, Key: "X", Value: `Warning:  Once you convert a character to expansion, you cannot play it in original Diablo II games.`},
	810: {ID: 810, Key: "X", Value: ``},
	811: {ID: 811, Key: "Bloodraven's Charge", Value: `Blood Raven's Charge`},
	812: {ID: 812, Key: "of Spirit", Value: `of Spirit`},
	813: {ID: 813, Key: "X", Value: `Are you sure you wish to continue?`},
	814: {ID: 814, Key: "Angel's Song", Value: `Angel's Song`},
	815: {ID: 815, Key: "Trainer's", Value: `Trainer's`},
	816: {ID: 816, Key: "Harmful", Value: `Harmful`},
	817: {ID: 817, Key: "RescueQual-KehkAct5Dru", Value: `Follow me.`},
	818: {ID: 818, Key: "Go North", Value: `go north`},
	819: {ID: 819, Key: "x", Value: `NEW EXPANSION CHARACTER`},
	820: {ID: 820, Key: "Pagan's Athame", Value: `Pagan's Athame`},
	821: {ID: 821, Key: "x", Value: `NEW EXPANSION OPEN CHARACTER.`},
	822: {ID: 822, Key: "Shaman's", Value: `Shaman's`},
	823: {ID: 823, Key: "x", Value: `NEW EXPANSION REALM CHARACTER.`},
	824: {ID: 824, Key: "Dense", Value: `Dense`},
	825: {ID: 825, Key: "Sazabi's Mental Sheath", Value: `Sazabi's Mental Sheath`},
	826: {ID: 826, Key: "X", Value: `CONVERT TO EXPANSION`},
	827: {ID: 827, Key: "of Frigidity", Value: `of Frigidity`},
	828: {ID: 828, Key: "Toothrow", Value: `Toothrow`},
	829: {ID: 829, Key: "Guardian Naga", Value: `Guardian Naga`},
	830: {ID: 830, Key: "Lycander's Flank", Value: `Lycander's Flank`},
	831: {ID: 831, Key: "GreaterIceSpawn", Value: `Greater Ice Spawn`},
	832: {ID: 832, Key: "X", Value: `CREATE NEW CHARACTER`},
	833: {ID: 833, Key: "X", Value: `DELETE CHARACTER`},
	834: {ID: 834, Key: "X", Value: `Ethereal (Cannot be Repaired)`},
	835: {ID: 835, Key: "X", Value: `This item cannot be repaired.`},
	836: {ID: 836, Key: "Rename Instruct", Value: `Choose the item that you wish to personalize with your name.`},
	837: {ID: 837, Key: "Addsocketsui", Value: `add sockets`},
	838: {ID: 838, Key: "Addsocketsui2", Value: `Choose the item to which you wish to add sockets.`},
	839: {ID: 839, Key: "uow", Value: `Aegis`},
	840: {ID: 840, Key: "of Fire Arrows", Value: `of Fire Arrow`},
	841: {ID: 841, Key: "MercX101", Value: `Vardhaka`},
	842: {ID: 842, Key: "MinionSpawner", Value: `Demon Portal`},
	843: {ID: 843, Key: "upk", Value: `Blade Barrier`},
	844: {ID: 844, Key: "upl", Value: `Balrog Skin`},
	845: {ID: 845, Key: "MercX102", Value: `Khan`},
	846: {ID: 846, Key: "MercX103", Value: `Klisk`},
	847: {ID: 847, Key: "MercX104", Value: `Bors`},
	848: {ID: 848, Key: "Messerschmidt's Reaver", Value: `Messerschmidt's Reaver`},
	849: {ID: 849, Key: "Immortal King's Stone Crusher", Value: `Immortal King's Stone Crusher`},
	850: {ID: 850, Key: "Dac Farren", Value: `Dac Farren`},
	851: {ID: 851, Key: "Scarlet", Value: `Scarlet`},
	852: {ID: 852, Key: "Waterwalk", Value: `Waterwalk`},
	853: {ID: 853, Key: "groundtomb", Value: `Tomb`},
	854: {ID: 854, Key: "Nature's", Value: `Natural`},
	855: {ID: 855, Key: "Swordguard", Value: `Swordguard`},
	856: {ID: 856, Key: "Of the Stream", Value: `of the Stream`},
	857: {ID: 857, Key: "HoradricMalusAss", Value: `A Malus! This should go to Charsi.`},
	858: {ID: 858, Key: "Demonlimb", Value: `Demon Limb`},
	859: {ID: 859, Key: "MercX105", Value: `Brom`},
	860: {ID: 860, Key: "Guillaume's Face", Value: `Guillaume's Face`},
	861: {ID: 861, Key: "MercX106", Value: `Wiglaf`},
	862: {ID: 862, Key: "CompletingNihlathakAct5Ams", Value: `Conspiring with Baal... What a tragic mistake.`},
	863: {ID: 863, Key: "r01L", Value: `El`},
	864: {ID: 864, Key: "MercX107", Value: `Hrothgar`},
	865: {ID: 865, Key: "MercX108", Value: `Scyld`},
	866: {ID: 866, Key: "MercX109", Value: `Healfdane`},
	867: {ID: 867, Key: "MercX110", Value: `Heorogar`},
	868: {ID: 868, Key: "MercX111", Value: `Halgaunt`},
	869: {ID: 869, Key: "of Bone Spirits", Value: `of Bone Spirit`},
	870: {ID: 870, Key: "MercX112", Value: `Hygelac`},
	871: {ID: 871, Key: "urg", Value: `Hyperion`},
	872: {ID: 872, Key: "MercX113", Value: `Egtheow`},
	873: {ID: 873, Key: "MercX114", Value: `Bohdan`},
	874: {ID: 874, Key: "MercX115", Value: `Wulfgar`},
	875: {ID: 875, Key: "Of the Bayou", Value: `of the Bayou`},
	876: {ID: 876, Key: "of Fire Golem Summoning", Value: `of Fire Golem`},
	877: {ID: 877, Key: "Aldur's Deception", Value: `Aldur's Deception`},
	878: {ID: 878, Key: "urn", Value: `Corona`},
	879: {ID: 879, Key: "r02L", Value: `Eld`},
	880: {ID: 880, Key: "Nightwing's Veil", Value: `Nightwing's Veil`},
	881: {ID: 881, Key: "MercX116", Value: `Hild`},
	882: {ID: 882, Key: "MercX117", Value: `Heatholaf`},
	883: {ID: 883, Key: "72a", Value: `Ettin Axe`},
	884: {ID: 884, Key: "urs", Value: `Great Hauberk`},
	885: {ID: 885, Key: "Reclusive", Value: `Reclusive`},
	886: {ID: 886, Key: "of Blaze", Value: `of Blaze`},
	887: {ID: 887, Key: "Repulsive", Value: `Repulsive`},
	888: {ID: 888, Key: "ush", Value: `Troll Nest`},
	889: {ID: 889, Key: "A5Q6InitAncients", Value: `31
You are a worthy hero! We augment 
your skill and grant you entry to the 
interior of Mount Arreat, wherein lies 
the Worldstone.
 
Beware. You will not be alone. Baal the 
Lord of Destruction is already inside. 
 
The Archangel Tyrael has always been 
our benefactor, but even he cannot 
help us now. For Baal blocks Tyrael's 
spiritual presence from entering the 
chamber of the Worldstone. Only you, 
mortal, have the power to defeat Baal 
now.
 
Baal threatens the Worldstone -- and 
through it, the mortal realm, itself. You 
must stop him before he gains full 
control of the sacred stone. With it 
under his control, Baal could shatter 
the boundaries between this world and 
the Burning Hells, thus allowing the 
hordes of the Prime Evils to pour forth 
into the mortal realm like an 
unstoppable tide!
 
If you are weak, the world as you know 
it could be lost forever. You must NOT 
fail!
`},
	890: {ID: 890, Key: "72h", Value: `Legend Sword`},
	891: {ID: 891, Key: "usk", Value: `Demonhead`},
	892: {ID: 892, Key: "Blood Rain", Value: `Blood Rain`},
	893: {ID: 893, Key: "MercX118", Value: `Weder`},
	894: {ID: 894, Key: "hellgate", Value: `Hell Gate`},
	895: {ID: 895, Key: "r03L", Value: `Tir`},
	896: {ID: 896, Key: "Mechanist's", Value: `Mechanic's`},
	897: {ID: 897, Key: "Wailing", Value: `Wailing`},
	898: {ID: 898, Key: "utb", Value: `Mirrored Boots`},
	899: {ID: 899, Key: "utc", Value: `Troll Belt`},
	900: {ID: 900, Key: "Of the Horde", Value: `of the Horde`},
	901: {ID: 901, Key: "MalahAct5IntroGossip1", Value: `35
I, Malah, welcome you to Harrogath, 
the last stronghold of Order on Mount 
Arreat. You have come to the right 
place if you intend to defeat Baal the 
Lord of Destruction.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what. 
 
All of the Elders, save Nihlathak, 
sacrificed themselves to place a 
protective ward around Harrogath.
 
Some of us here, certainly Nihlathak, do 
not appreciate your presence. We are a 
proud people, and it is not easy for us 
to accept aid. I, however, am glad you 
are here.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`},
	902: {ID: 902, Key: "Mentalist's", Value: `Mentalist's`},
	903: {ID: 903, Key: "utg", Value: `Crusader Gauntlets`},
	904: {ID: 904, Key: "uth", Value: `Lacquered Plate`},
	905: {ID: 905, Key: "Charmdes", Value: `Keep in Inventory to Gain Bonus`},
	906: {ID: 906, Key: "Suicide Branch", Value: `Suicide Branch`},
	907: {ID: 907, Key: "Hellrack", Value: `Hellrack`},
	908: {ID: 908, Key: "Sander's Court Jester", Value: `McAuley's Folly`},
	909: {ID: 909, Key: "tomb1L", Value: `Tomb`},
	910: {ID: 910, Key: "MercX119", Value: `Vikhyat`},
	911: {ID: 911, Key: "r04L", Value: `Nef`},
	912: {ID: 912, Key: "utp", Value: `Archon Plate`},
	913: {ID: 913, Key: "ChampionFormatX", Value: `%0 %1`},
	914: {ID: 914, Key: "CompletingDOEDru", Value: `Bah! Is that all of them?`},
	915: {ID: 915, Key: "uuc", Value: `Heater`},
	916: {ID: 916, Key: "uts", Value: `Ward`},
	917: {ID: 917, Key: "utu", Value: `Wire Fleece`},
	918: {ID: 918, Key: "Ivory", Value: `Ivory`},
	919: {ID: 919, Key: "Miocene", Value: `Enlightened`},
	920: {ID: 920, Key: "Of Love", Value: `of Love`},
	921: {ID: 921, Key: "uui", Value: `Dusk Shroud`},
	922: {ID: 922, Key: "Ghostflame", Value: `Ghostflame`},
	923: {ID: 923, Key: "SnowYeti1", Value: `Snow Drifter`},
	924: {ID: 924, Key: "uul", Value: `Shadow Plate`},
	925: {ID: 925, Key: "of the Sirocco", Value: `of the Scirocco`},
	926: {ID: 926, Key: "A5Q1InitLarzuk", Value: `45
If you're here to defeat Baal, you must 
prove it!
 
As we speak, Harrogath is under siege 
by Baal's demons. Catapults rain death 
just outside the town walls.
 
Baal himself travels up the sacred 
mountain, having left in charge here 
one of his most vicious generals, Shenk 
the Overseer. A ruthless taskmaster, 
he lashes his own minions into suicidal 
frenzies on the battlefield.
 
If you wish to prove yourself to us, 
destroy the monster, Shenk, that 
commands those infernal catapults 
outside Harrogath.  If you manage to 
do this, return to me.
`},
	927: {ID: 927, Key: "r05L", Value: `Eth`},
	928: {ID: 928, Key: "A5Q4SuccessfulMalah", Value: `65
So, the Relic is lost. Do not dwell on 
failures past. It is your future that 
matters more.
`},
	929: {ID: 929, Key: "Of Dread", Value: `of Dread`},
	930: {ID: 930, Key: "uvb", Value: `Scarabshell Boots`},
	931: {ID: 931, Key: "uvc", Value: `Vampirefang Belt`},
	932: {ID: 932, Key: "Of Honor", Value: `of Honor`},
	933: {ID: 933, Key: "Of the Retard", Value: `of the Moron`},
	934: {ID: 934, Key: "of Vita", Value: `of Vita`},
	935: {ID: 935, Key: "uvg", Value: `Vampirebone Gloves`},
	936: {ID: 936, Key: "QualKehkAct5IntroGossip1", Value: `45
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`},
	937: {ID: 937, Key: "Andariel's Visage", Value: `Andariel's Visage`},
	938: {ID: 938, Key: "A5Q4SuccessfulAnya", Value: `55
You have stopped Nihlathak, but he 
didn't have the Relic! He must have 
already given it to Baal. Now, Baal will 
not be tested when he reaches Arreat's 
summit.
 
...Damn Nihlathak!
 
I do thank you for trying, though. 
Please, allow me to honor your courage 
by magically inscribing your name onto 
an item of your choosing. It's the least 
I can do.
`},
	939: {ID: 939, Key: "Armor", Value: `Armor`},
	940: {ID: 940, Key: "NihlathakGossip1", Value: `49
If you're looking for cases of 
treacherous magic, Assassin, take a 
hard look at Larzuk. He was the only 
one in town who escaped the Red Fever 
last spring. He claimed his good 
fortune was due to 'hand washing' 
before meals.
 
Hmmm...Very suspicious...
`},
	941: {ID: 941, Key: "NihlathakGossip2", Value: `48
Why would you seek this place, 
stranger? Are you a vulture come to 
loot the bodies of our fallen warriors? 
 
Regardless, this is no place to make a 
name for yourself. The mountain is 
ours to protect. It is only a matter of 
time before Hell's legions are routed.
`},
	942: {ID: 942, Key: "NihlathakGossip3", Value: `67
Qual-Kehk is useless. He has blindly 
sent our warriors to their deaths, 
assuming Baal's legions would fight as 
men do. Of course, everyone knows 
they do not.
`},
	943: {ID: 943, Key: "NihlathakGossip4", Value: `57
The demon hordes have grown powerful 
beyond measure, aided by our foolish 
mistakes. But I may have found a way 
to correct those mistakes...
`},
	944: {ID: 944, Key: "NihlathakGossip5", Value: `37
Oh yes...I remember our warriors as 
children. Malah would set their broken 
bones and give them powders for their 
fevers. Now, they return to her with 
wounds that will never heal.
 
I am tired...Please...leave me.
`},
	945: {ID: 945, Key: "NihlathakGossip6", Value: `133
If you have nothing useful for me, be on 
your way!
`},
	946: {ID: 946, Key: "NihlathakGossip7", Value: `81
Anya's father was my good friend. 
There are so many to mourn...I have 
no time for you!
`},
	947: {ID: 947, Key: "NihlathakGossip8", Value: `53
I have long been criticized, but 
especially of late -- since the deaths of 
my fellow Elders. Through it all, I have 
learned one thing. Each man must do 
what's right, no matter what others 
may think.
`},
	948: {ID: 948, Key: "NihlathakGossip9", Value: `65
The Council of Elders always believed 
itself prepared for the coming of the 
Three. Obviously, we were not prepared 
enough.
`},
	949: {ID: 949, Key: "r06L", Value: `Ith`},
	950: {ID: 950, Key: "Grandmaster's", Value: `Grandmaster's`},
	951: {ID: 951, Key: "Unearthly", Value: `Unearthly`},
	952: {ID: 952, Key: "of Incineration", Value: `of Incineration`},
	953: {ID: 953, Key: "The Cranium Basher", Value: `The Cranium Basher`},
	954: {ID: 954, Key: "Tal Rasha's Fire-Spun Cloth", Value: `Tal Rasha's Fine-Spun Cloth`},
	955: {ID: 955, Key: "Immortal King's Will", Value: `Immortal King's Will`},
	956: {ID: 956, Key: "LeaveCampDru", Value: `So, it begins.`},
	957: {ID: 957, Key: "FindingSummonerAss", Value: `Summoner, the dark magics have corrupted you.`},
	958: {ID: 958, Key: "of Anima", Value: `of Amicae`},
	959: {ID: 959, Key: "r07L", Value: `Tal`},
	960: {ID: 960, Key: "EnterBurialDru", Value: `Planting the dead... How odd.`},
	961: {ID: 961, Key: "CompletingNihlathakAct5Bar", Value: `A fitting death for a traitor.`},
	962: {ID: 962, Key: "CompletingNihlathakAct5Ass", Value: `You Dark Mages are all alike - obsessed with power.`},
	963: {ID: 963, Key: "Of the Lilly", Value: `of the Lilly`},
	964: {ID: 964, Key: "A5Q6EarlyReturnAnya", Value: `60
Baal has blocked Tyrael from entering 
the Worldstone Chamber? This truly 
has become a battle against Hell.
 
Whether or not it was the Heavens' 
decree, this is your fight now -- your 
destiny.
`},
	965: {ID: 965, Key: "strUI10", Value: `Solar Creeper`},
	966: {ID: 966, Key: "strUI11", Value: `Spirit of Barbs`},
	967: {ID: 967, Key: "Crow Caw", Value: `Crow Caw`},
	968: {ID: 968, Key: "of Grim Ward", Value: `of Grim Ward`},
	969: {ID: 969, Key: "of Charged Spear", Value: `of Charged Spear`},
	970: {ID: 970, Key: "Ondal's Wisdom", Value: `Ondal's Wisdom`},
	971: {ID: 971, Key: "strUI12", Value: `Heart of Wolverine`},
	972: {ID: 972, Key: "strUI13", Value: `Vine`},
	973: {ID: 973, Key: "strUI14", Value: `Spirit`},
	974: {ID: 974, Key: "StrUI18", Value: `Master`},
	975: {ID: 975, Key: "r08L", Value: `Ral`},
	976: {ID: 976, Key: "Snowclash", Value: `Snowclash`},
	977: {ID: 977, Key: "WailingSpirit", Value: `Wailing Spirit`},
	978: {ID: 978, Key: "HellSpawn", Value: `Hell Spawn`},
	979: {ID: 979, Key: "SnowYeti2", Value: `Abominable`},
	980: {ID: 980, Key: "SnowYeti3", Value: `Chilled Froth`},
	981: {ID: 981, Key: "of Cyclone Armor", Value: `of Cyclone Armor`},
	982: {ID: 982, Key: "of Shouting", Value: `of Shout`},
	983: {ID: 983, Key: "Spineripper", Value: `Spineripper`},
	984: {ID: 984, Key: "SnowYeti4", Value: `Frozen Abyss`},
	985: {ID: 985, Key: "Stalwart", Value: `Stalwart`},
	986: {ID: 986, Key: "Eskillpercentlif", Value: ` Percent Life`},
	987: {ID: 987, Key: "A5Q4EarlyReturnMalah", Value: `63
Nihlathak was never the kindest man. 
But for him to betray the whole world...
 
Ahh...Where shall his soul finally rest?
`},
	988:  {ID: 988, Key: "Threash Socket", Value: `Thresh Socket`},
	989:  {ID: 989, Key: "Axe Dweller", Value: `Axe Dweller`},
	990:  {ID: 990, Key: "Waypoint", Value: `Waypoint`},
	991:  {ID: 991, Key: "r09L", Value: `Ort`},
	992:  {ID: 992, Key: "EnteringNihlathakAct5Sor", Value: `Could this be a trap?`},
	993:  {ID: 993, Key: "Dark Clan Crusher", Value: `Dark Clan Crusher`},
	994:  {ID: 994, Key: "of Poison Jab", Value: `of Poison Jab`},
	995:  {ID: 995, Key: "magic shrine2", Value: `Magic Shrine`},
	996:  {ID: 996, Key: "banner 2", Value: `Banner`},
	997:  {ID: 997, Key: "Rose Branded", Value: `Rose Branded`},
	998:  {ID: 998, Key: "banner 1", Value: `Banner`},
	999:  {ID: 999, Key: "Kuko Shakaku", Value: `Kuko Shakaku`},
	1000: {ID: 1000, Key: "Hone Sundan", Value: `Hone Sundan`},
	1001: {ID: 1001, Key: "Class Specific", Value: `Class-specific`},
	1002: {ID: 1002, Key: "KillingdDiabloPal", Value: `Let Diablo's death end the reign of the Three!`},
	1003: {ID: 1003, Key: "healthshrine", Value: `Health Shrine`},
	1004: {ID: 1004, Key: "tomb3L", Value: `Tomb`},
	1005: {ID: 1005, Key: "tomb2L", Value: `Tomb`},
	1006: {ID: 1006, Key: "Rigid Highlands", Value: `Frigid Highlands`},
	1007: {ID: 1007, Key: "To The Worldstone Chamber", Value: `To The Worldstone Chamber`},
	1008: {ID: 1008, Key: "Of Erosion", Value: `of Erosion`},
	1009: {ID: 1009, Key: "of Dim Vision", Value: `of Dim Vision`},
	1010: {ID: 1010, Key: "Carrion Wind", Value: `Carrion Wind`},
	1011: {ID: 1011, Key: "Blood Lord1", Value: `Moon Lord`},
	1012: {ID: 1012, Key: "EskillPetLife", Value: `Pet Life: `},
	1013: {ID: 1013, Key: "Skin of the Flayed One", Value: `Skin of Flayed One`},
	1014: {ID: 1014, Key: "Blood Lord2", Value: `Night Lord`},
	1015: {ID: 1015, Key: "Blood Lord3", Value: `Blood Lord`},
	1016: {ID: 1016, Key: "Blood Lord4", Value: `Hell Lord`},
	1017: {ID: 1017, Key: "A5Q4SuccessfulCain", Value: `90
Beware! Baal grows stronger with every 
passing moment.
`},
	1018: {ID: 1018, Key: "Of Tribute", Value: `of Tribute`},
	1019: {ID: 1019, Key: "Blood Lord5", Value: `Death Lord`},
	1020: {ID: 1020, Key: "StrGemX1", Value: `Helms:`},
	1021: {ID: 1021, Key: "StrGemX2", Value: `Shields:`},
	1022: {ID: 1022, Key: "StrGemX3", Value: `Weapons:`},
	1023: {ID: 1023, Key: "StrGemX4", Value: `Armor:`},
	1024: {ID: 1024, Key: "To The Rocky Summit", Value: `To The Arreat Summit`},
	1025: {ID: 1025, Key: "hiredismiss", Value: `Dismiss`},
	1026: {ID: 1026, Key: "Stoneraven", Value: `Stoneraven`},
	1027: {ID: 1027, Key: "hireresurrect", Value: `Resurrect`},
	1028: {ID: 1028, Key: "Arkaine's Valor", Value: `Arkaine's Valor`},
	1029: {ID: 1029, Key: "MiniPanelHireinv", Value: `Hireling Inventory`},
	1030: {ID: 1030, Key: "Moonshadow", Value: `Moon Shadow`},
	1031: {ID: 1031, Key: "Druid Hawk", Value: `Raven`},
	1032: {ID: 1032, Key: "MercX120", Value: `Unferth`},
	1033: {ID: 1033, Key: "of Poison Dagger", Value: `of Poison Dagger`},
	1034: {ID: 1034, Key: "CompletingSummonerAss", Value: `Horazon. Your decoy is dead.`},
	1035: {ID: 1035, Key: "Deaths's Web", Value: `Death's Web`},
	1036: {ID: 1036, Key: "MercX121", Value: `Sigemund`},
	1037: {ID: 1037, Key: "Valkiry Wing", Value: `Valkyrie Wing`},
	1038: {ID: 1038, Key: "Bahamut's", Value: `Bahamut's`},
	1039: {ID: 1039, Key: "MercX122", Value: `Heremod`},
	1040: {ID: 1040, Key: "of Energy Shield", Value: `of Energy Shield`},
	1041: {ID: 1041, Key: "Wicked", Value: `Wicked`},
	1042: {ID: 1042, Key: "Of Attrition", Value: `of Attrition`},
	1043: {ID: 1043, Key: "The Spirit Shroud", Value: `The Spirit Shroud`},
	1044: {ID: 1044, Key: "MercX123", Value: `Hengest`},
	1045: {ID: 1045, Key: "MercX124", Value: `Folcwald`},
	1046: {ID: 1046, Key: "of Warming", Value: `of Warming`},
	1047: {ID: 1047, Key: "MercX125", Value: `Frisian`},
	1048: {ID: 1048, Key: "Sparking", Value: `Sparking`},
	1049: {ID: 1049, Key: "of Iron Maiden", Value: `of Iron Maiden`},
	1050: {ID: 1050, Key: "Baal Subject Mummy", Value: `Unraveler`},
	1051: {ID: 1051, Key: "MercX126", Value: `Hnaef`},
	1052: {ID: 1052, Key: "MercX127", Value: `Guthlaf`},
	1053: {ID: 1053, Key: "MercX128", Value: `Oslaf`},
	1054: {ID: 1054, Key: "CompletingDefeatBaalAct5Sor", Value: `The last of the Three has fallen.`},
	1055: {ID: 1055, Key: "Crystalized Cavern Level 1", Value: `Crystalline Passage`},
	1056: {ID: 1056, Key: "Rude", Value: `Rude`},
	1057: {ID: 1057, Key: "Burly", Value: `Burly`},
	1058: {ID: 1058, Key: "Mosers Blessed Circle", Value: `Moser's Blessed Circle`},
	1059: {ID: 1059, Key: "Crystalized Cavern Level 2", Value: `Glacial Trail`},
	1060: {ID: 1060, Key: "Pindleskin", Value: `Pindleskin`},
	1061: {ID: 1061, Key: "Of Valhalla", Value: `of Valhalla`},
	1062: {ID: 1062, Key: "Coldsteel Eye", Value: `Coldsteel Eye`},
	1063: {ID: 1063, Key: "MercX129", Value: `Yrmenlaf`},
	1064: {ID: 1064, Key: "MercX130", Value: `Garmund`},
	1065: {ID: 1065, Key: "MercX131", Value: `Freawaru`},
	1066: {ID: 1066, Key: "Sander's Basis", Value: `McAuley's Basis`},
	1067: {ID: 1067, Key: "CompletingForgottenTPal", Value: `So much treasure almost covers the stench.`},
	1068: {ID: 1068, Key: "CompletingForgottenTPal", Value: `This tower has its charms...`},
	1069: {ID: 1069, Key: "A5Q2EarlyReturnQualKehkMan", Value: `78
More of my men are still alive out there. 
I am certain of it!
 
Find them. Free them from their cages 
and bring them back to me.
`},
	1070: {ID: 1070, Key: "Tang's Rule", Value: `Tang's Rule`},
	1071: {ID: 1071, Key: "BanishedSoul", Value: `Banished Soul`},
	1072: {ID: 1072, Key: "etorch2", Value: `Torch`},
	1073: {ID: 1073, Key: "A5Q1EarlyReturnLarzuk", Value: `90
What's the matter, hero? Questioning 
your fortitude? I know we are.
`},
	1074: {ID: 1074, Key: "Annihilus", Value: `Annihilus`},
	1075: {ID: 1075, Key: "Of the Jujube", Value: `of the Jujube`},
	1076: {ID: 1076, Key: "etorch1", Value: `Torch`},
	1077: {ID: 1077, Key: "MercX132", Value: `Eadgils`},
	1078: {ID: 1078, Key: "A5Q1AfterInitCain", Value: `83
I believe that stopping the siege on 
Harrogath is the only way for you to 
earn the trust of these people.
`},
	1079: {ID: 1079, Key: "MercX133", Value: `Onela`},
	1080: {ID: 1080, Key: "MercX134", Value: `Damien`},
	1081: {ID: 1081, Key: "Of the Stiletto", Value: `Of the Stiletto`},
	1082: {ID: 1082, Key: "MercX135", Value: `Erfor`},
	1083: {ID: 1083, Key: "Mnemonic", Value: `Mnemonic`},
	1084: {ID: 1084, Key: "Golemlord's", Value: `Golemlord's`},
	1085: {ID: 1085, Key: "of Impaling Strike", Value: `of Impaling Strike`},
	1086: {ID: 1086, Key: "McAuley's Paragon", Value: `Sander's Paragon`},
	1087: {ID: 1087, Key: "of Taunting", Value: `of Taunt`},
	1088: {ID: 1088, Key: "hirehelp1", Value: `I am hurt!`},
	1089: {ID: 1089, Key: "hirehelp2", Value: `Help!`},
	1090: {ID: 1090, Key: "Death Mauler1", Value: `Death Mauler`},
	1091: {ID: 1091, Key: "Metalgird", Value: `Metalite's Girth`},
	1092: {ID: 1092, Key: "Death Mauler2", Value: `Death Brawler`},
	1093: {ID: 1093, Key: "Death Mauler3", Value: `Death Slasher`},
	1094: {ID: 1094, Key: "A5Q2AfterInitCain", Value: `74
I know firsthand that captivity is a sad 
fate for a man. Find them quickly.
`},
	1095: {ID: 1095, Key: "Nethercrow", Value: `Nethercrow`},
	1096: {ID: 1096, Key: "of Damage Amplification", Value: `of Amplify Damage`},
	1097: {ID: 1097, Key: "of the Squid", Value: `of the Squid`},
	1098: {ID: 1098, Key: "Tyrael's Might", Value: `Tyrael's Might`},
	1099: {ID: 1099, Key: "FindingTempleAss", Value: `I dread this place of darkness.`},
	1100: {ID: 1100, Key: "Death Mauler4", Value: `Death Berserker`},
	1101: {ID: 1101, Key: "FindingDrainLeverAss", Value: `Levers are made to be pulled.`},
	1102: {ID: 1102, Key: "Death Mauler5", Value: `Death Brigadier`},
	1103: {ID: 1103, Key: "r10L", Value: `Thul`},
	1104: {ID: 1104, Key: "of Revivification", Value: `of Revivification`},
	1105: {ID: 1105, Key: "Halls of Death's Calling", Value: `Halls of Pain`},
	1106: {ID: 1106, Key: "A5Q3SuccessfulLarzuk", Value: `67
I never liked Nihlathak, but I never 
suspected that he'd betray us! I just 
can't understand how an Elder could do 
such a thing.
`},
	1107: {ID: 1107, Key: "hirehelp3", Value: `I am dying.`},
	1108: {ID: 1108, Key: "Joker's", Value: `Joker's`},
	1109: {ID: 1109, Key: "ESkillPerKick", Value: ` per kick`},
	1110: {ID: 1110, Key: "A5Q3AfterInitCain", Value: `59
I would listen to Malah. Nihlathak 
speaks with a venomous tongue and 
acts as if the entire weight of this town 
rests upon his shoulders.
 
Perhaps there is more going on here 
than we know.
`},
	1111: {ID: 1111, Key: "of Enlightenment", Value: `of Enlightenment`},
	1112: {ID: 1112, Key: "Skin of the Vipermagi", Value: `Skin of the Vipermagi`},
	1113: {ID: 1113, Key: "hirehelp4", Value: `Help me!`},
	1114: {ID: 1114, Key: "Rockhew", Value: `Rockhew`},
	1115: {ID: 1115, Key: "Thunderstroke", Value: `Thunderstroke`},
	1116: {ID: 1116, Key: "MercX136", Value: `Weohstan`},
	1117: {ID: 1117, Key: "Shrine3wilderness", Value: `Shrine`},
	1118: {ID: 1118, Key: "MercX137", Value: `Wulf`},
	1119: {ID: 1119, Key: "r11L", Value: `Amn`},
	1120: {ID: 1120, Key: "of Nirvana", Value: `of Nirvana`},
	1121: {ID: 1121, Key: "Baal Tentacle", Value: `Festering Appendages`},
	1122: {ID: 1122, Key: "McAuley's Superstition", Value: `Sander's Superstition`},
	1123: {ID: 1123, Key: "Natalya's Odium", Value: `Natalya's Odium`},
	1124: {ID: 1124, Key: "pyox", Value: `Fire Pit`},
	1125: {ID: 1125, Key: "Bonehew", Value: `Bonehew`},
	1126: {ID: 1126, Key: "A5Q4AfterInitCain", Value: `50
Regretfully, I know very little about this 
Relic. However, if what the others say 
is true, then Baal must not gain 
possession of it.
 
Stop Nihlathak...before all is lost.
`},
	1127: {ID: 1127, Key: "of Thawing", Value: `of Thawing`},
	1128: {ID: 1128, Key: "Zircon", Value: `Zircon`},
	1129: {ID: 1129, Key: "Horazon's Chalice", Value: `Horazon's Chalice`},
	1130: {ID: 1130, Key: "of War Cry", Value: `of War Cry`},
	1131: {ID: 1131, Key: "of Meteors", Value: `of Meteor`},
	1132: {ID: 1132, Key: "of Fist of the Heavens", Value: `of Fist of the Heavens`},
	1133: {ID: 1133, Key: "Thudergod's Vigor", Value: `Thundergod's Vigor`},
	1134: {ID: 1134, Key: "of Propogation", Value: `of Propogation`},
	1135: {ID: 1135, Key: "r12L", Value: `Sol`},
	1136: {ID: 1136, Key: "Echoing", Value: `Echoing`},
	1137: {ID: 1137, Key: "Hwanin's Cordon", Value: `Hwanin's Cordon`},
	1138: {ID: 1138, Key: "Wyvern's Head", Value: `Wyvern's Head`},
	1139: {ID: 1139, Key: "Eskillphoenix1", Value: `meteor damage: `},
	1140: {ID: 1140, Key: "Eskillphoenix2", Value: `chain lightning damage: `},
	1141: {ID: 1141, Key: "Eskillphoenix3", Value: `chaos ice bolt damage: `},
	1142: {ID: 1142, Key: "A5Q5AfterInitCain", Value: `90
A test of mettle is a fitting rite of 
passage for a Barbarian hero.
`},
	1143: {ID: 1143, Key: "A5Q3FoundAnyaLarzuk", Value: `80
Poor Anya! I should've known Nihlathak 
was a traitor...
 
If you find him alive, kill him for me!
`},
	1144: {ID: 1144, Key: "of Exploding Arrows", Value: `of Exploding Arrow`},
	1145: {ID: 1145, Key: "AnyaGossip1", Value: `45
Now that the Elders are all dead, I don't 
know who will guide our people through 
this dark time. I was to be next in line 
after my father, but this burden is too 
great for me to shoulder alone. 
 
We are a people of strong blood. I shall 
do what I can and let fate do the rest.
`},
	1146: {ID: 1146, Key: "AnyaGossip2", Value: `82
Baal's minions are not to be trifled with. 
In their bloodlust they will sacrifice 
themselves to destroy you.
`},
	1147: {ID: 1147, Key: "AnyaGossip3", Value: `60
Many outsiders believe that the 
fantastic stories about our ancestors, 
the Ancients, are but fables. However, I 
believe that the Ancients were more 
than human -- that mankind has fallen 
from what it once was.
`},
	1148: {ID: 1148, Key: "AnyaGossip4", Value: `58
When I was a child, the Elders told us 
stories about the mountain and its 
power...and how the Barbarian people 
are bound to it as protectors.
 
Baal is not just taking our land -- with 
every step he takes up the mountain, 
he takes a part of who we are as a 
people.
`},
	1149: {ID: 1149, Key: "AnyaGossip5", Value: `72
I am truly glad you are here, warrior. 
Perhaps things would be different now 
if we had asked for assistance from 
the neighboring kingdoms.
 
Our foolish, foolish pride...
`},
	1150: {ID: 1150, Key: "AnyaGossip6", Value: `54
My father, Aust, was among those 
Elders wise enough to see that we 
would need outside help to deal with 
Baal's legions. He believed that this 
conflict would affect the entire world, 
not just our homeland. He said that 
Mount Arreat is as necessary to the 
world's survival as food and water is to 
our own.
 
I believe this to be true.
`},
	1151: {ID: 1151, Key: "AnyaGossip7", Value: `50
Qual-Kehk's confidence in his abilities 
once bordered on arrogance, but in the 
early days of the siege, he was 
humbled by Baal. While he saved us 
from immediate destruction, a third of 
our warriors were lost.
 
None felt those losses more than he. 
Though he may not admit it, your 
presence gives him hope, warrior.
`},
	1152: {ID: 1152, Key: "AnyaGossip8", Value: `78
If Larzuk could sing or dance half as 
well as he smiths, he'd be married by 
now.
 
...Just look at those shoulders.
`},
	1153: {ID: 1153, Key: "AnyaGossip9", Value: `57
Nihlathak was the last of our Elders, 
whose charge it was to safeguard the 
mountain. He alone tried to guide us 
through the most desperate time in our 
history -- but he was just as helpless as 
the rest of us.
 
I cannot forgive his betrayal, but I can 
learn from it.
`},
	1154: {ID: 1154, Key: "r13L", Value: `Shae`},
	1155: {ID: 1155, Key: "Burning", Value: `Burning`},
	1156: {ID: 1156, Key: "Moonfall", Value: `Moonfall`},
	1157: {ID: 1157, Key: "Shadow's Touch", Value: `Shadow Touch`},
	1158: {ID: 1158, Key: "Of Credit", Value: `of Credit`},
	1159: {ID: 1159, Key: "McAuley's Taboo", Value: `Sander's Taboo`},
	1160: {ID: 1160, Key: "CatapultE", Value: `Catapult`},
	1161: {ID: 1161, Key: "Crescent Moon", Value: `Crescent Moon`},
	1162: {ID: 1162, Key: "The Scalper", Value: `The Scalper`},
	1163: {ID: 1163, Key: "Rocky Summit", Value: `Arreat Summit`},
	1164: {ID: 1164, Key: "MercX138", Value: `Bulwye`},
	1165: {ID: 1165, Key: "MercX139", Value: `Lief`},
	1166: {ID: 1166, Key: "MercX140", Value: `Magnus`},
	1167: {ID: 1167, Key: "r14L", Value: `Dol`},
	1168: {ID: 1168, Key: "Warhound", Value: `Warhound`},
	1169: {ID: 1169, Key: "Seige Tower", Value: `Seige Tower`},
	1170: {ID: 1170, Key: "FindingSummonerDru", Value: `This place would drive anyone mad.`},
	1171: {ID: 1171, Key: "StygianHarlot", Value: `Stygian Harlot`},
	1172: {ID: 1172, Key: "Blade of Ali Baba", Value: `Blade of Ali Baba`},
	1173: {ID: 1173, Key: "ktr", Value: `Katar`},
	1174: {ID: 1174, Key: "Naj's Ancient Set", Value: `Naj's Ancient Vestige`},
	1175: {ID: 1175, Key: "Trang-Oul's Avatar", Value: `Trang-Oul's Avatar`},
	1176: {ID: 1176, Key: "of Power Strike", Value: `of Power Strike`},
	1177: {ID: 1177, Key: "of Fire Wall", Value: `of Fire Wall`},
	1178: {ID: 1178, Key: "Fuego Del Sol", Value: `Fuego Del Sol`},
	1179: {ID: 1179, Key: "Hadeshorn", Value: `Hadeshorn`},
	1180: {ID: 1180, Key: "FanaticMinion", Value: `Fanatic Enslaved`},
	1181: {ID: 1181, Key: "InsaneIceSpawn", Value: `Insane Ice Spawn`},
	1182: {ID: 1182, Key: "Cairn Shard", Value: `Cairn Shard`},
	1183: {ID: 1183, Key: "r15L", Value: `Hel`},
	1184: {ID: 1184, Key: "CompletingDOEAss", Value: `The Rogues' test is done.`},
	1185: {ID: 1185, Key: "CatapultN", Value: `Catapult`},
	1186: {ID: 1186, Key: "CatapultS", Value: `Catapult`},
	1187: {ID: 1187, Key: "CatapultW", Value: `Catapult`},
	1188: {ID: 1188, Key: "MercX141", Value: `Klatu`},
	1189: {ID: 1189, Key: "MercX142", Value: `Drus`},
	1190: {ID: 1190, Key: "MercX143", Value: `Hoku`},
	1191: {ID: 1191, Key: "MercX144", Value: `Kord`},
	1192: {ID: 1192, Key: "Black Hades", Value: `Black Hades`},
	1193: {ID: 1193, Key: "Serrated", Value: `Serrated`},
	1194: {ID: 1194, Key: "A5Q3SuccessfulAnya", Value: `80
Thank you, hero, for rescuing me.
 
To show my personal gratitude, I give 
you this. I had it custom-made for you 
by Larzuk.
`},
	1195: {ID: 1195, Key: "CompletingRadamentDru", Value: `Return to dust, Radament.`},
	1196: {ID: 1196, Key: "MercX145", Value: `Uther`},
	1197: {ID: 1197, Key: "MercX146", Value: `Ip`},
	1198: {ID: 1198, Key: "MercX147", Value: `Ulf`},
	1199: {ID: 1199, Key: "r16L", Value: `Po`},
	1200: {ID: 1200, Key: "MercX148", Value: `Tharr`},
	1201: {ID: 1201, Key: "of Bone Spears", Value: `of Bone Spear`},
	1202: {ID: 1202, Key: "MercX149", Value: `Kaelim`},
	1203: {ID: 1203, Key: "MercX150", Value: `Ulric`},
	1204: {ID: 1204, Key: "Of Dawn", Value: `of Dawn`},
	1205: {ID: 1205, Key: "Witch-hunter's", Value: `Witch-hunter's`},
	1206: {ID: 1206, Key: "AmaOnly", Value: `(Amazon Only)`},
	1207: {ID: 1207, Key: "Commander's", Value: `Commander's`},
	1208: {ID: 1208, Key: "Gillian's Brazier", Value: `Gillian's Brazier`},
	1209: {ID: 1209, Key: "Fathom", Value: `Death's Fathom`},
	1210: {ID: 1210, Key: "Malah", Value: `Malah`},
	1211: {ID: 1211, Key: "Travel To Harrogath", Value: `travel to harrogath`},
	1212: {ID: 1212, Key: "A5Q6EarlyReturnCain", Value: `51
Remember this. Baal once possessed Tal 
Rasha, one of the most powerful of the 
ancient Horadrim.
 
Your battles with Mephisto and Diablo 
will pale in comparison to your battle 
with Baal.
 
The Lord of Destruction aided by Tal 
Rasha's knowledge...The mountain 
itself will tremble when you clash.
`},
	1213: {ID: 1213, Key: "Tundra Wastelands", Value: `Frozen Tundra`},
	1214: {ID: 1214, Key: "MercX151", Value: `Alaric`},
	1215: {ID: 1215, Key: "r17L", Value: `Lum`},
	1216: {ID: 1216, Key: "Griswolds's Redemption", Value: `Griswold's Redemption`},
	1217: {ID: 1217, Key: "The Worldstone Keep Level 1", Value: `Worldstone Keep Level 1`},
	1218: {ID: 1218, Key: "The Worldstone Keep Level 2", Value: `Worldstone Keep Level 2`},
	1219: {ID: 1219, Key: "CompletingBurialAss", Value: `What I kill stays dead.`},
	1220: {ID: 1220, Key: "The Worldstone Keep Level 3", Value: `Worldstone Keep Level 3`},
	1221: {ID: 1221, Key: "of the Elements", Value: `of the Elements`},
	1222: {ID: 1222, Key: "of Daring", Value: `of Daring`},
	1223: {ID: 1223, Key: "Sureshrill Frost", Value: `Sureshrill Frost`},
	1224: {ID: 1224, Key: "The TreeEnt", Value: `The Treentster`},
	1225: {ID: 1225, Key: "Evil hut", Value: `Evil Demon Hut`},
	1226: {ID: 1226, Key: "eweaponrackL", Value: `Weapon Rack`},
	1227: {ID: 1227, Key: "CompletingNihlathakAct5Nec", Value: `You were a sad little man, Nihlathak.`},
	1228: {ID: 1228, Key: "MercX152", Value: `Ethelred`},
	1229: {ID: 1229, Key: "MercX153", Value: `Caden`},
	1230: {ID: 1230, Key: "MercX154", Value: `Elgifu`},
	1231: {ID: 1231, Key: "r18L", Value: `Ko`},
	1232: {ID: 1232, Key: "MercX155", Value: `Tostig`},
	1233: {ID: 1233, Key: "Of Quota", Value: `of the Quota`},
	1234: {ID: 1234, Key: "Strange", Value: `Strange`},
	1235: {ID: 1235, Key: "MercX156", Value: `Alcuin`},
	1236: {ID: 1236, Key: "of Dawn", Value: `of Dawn`},
	1237: {ID: 1237, Key: "MercX157", Value: `Emund`},
	1238: {ID: 1238, Key: "To The Rigid Highlands", Value: `To The Frigid Highlands`},
	1239: {ID: 1239, Key: "MercX158", Value: `Sigurd`},
	1240: {ID: 1240, Key: "Ghostleach", Value: `Ghost Leach`},
	1241: {ID: 1241, Key: "of Decrepification", Value: `of Decrepify`},
	1242: {ID: 1242, Key: "of Ice Blasts", Value: `of Ice Blast`},
	1243: {ID: 1243, Key: "chestR", Value: `Chest`},
	1244: {ID: 1244, Key: "MercX159", Value: `Gorm`},
	1245: {ID: 1245, Key: "The Ensanguinator", Value: `The Ensanguinator`},
	1246: {ID: 1246, Key: "MercX160", Value: `Hollis`},
	1247: {ID: 1247, Key: "r19L", Value: `Fal`},
	1248: {ID: 1248, Key: "The Minataur", Value: `The Minotaur`},
	1249: {ID: 1249, Key: "CompletingStopSiegeSor", Value: `Harrogath can rest easier now.`},
	1250: {ID: 1250, Key: "Ironpelt", Value: `Iron Pelt`},
	1251: {ID: 1251, Key: "Eskillincasehit", Value: ` hit`},
	1252: {ID: 1252, Key: "Lawful", Value: `Lawful`},
	1253: {ID: 1253, Key: "Personalizeui", Value: `personalize`},
	1254: {ID: 1254, Key: "MercX161", Value: `Ragnar`},
	1255: {ID: 1255, Key: "am1", Value: `Stag Bow`},
	1256: {ID: 1256, Key: "am2", Value: `Reflex Bow`},
	1257: {ID: 1257, Key: "A5Q3EarlyReturnLarzuk", Value: `56
As the daughter of Elder Aust, Anya is 
the only person, besides Nihlathak, 
who has any real knowledge of Mount 
Arreat's secrets. I'd hate to see our 
fate in the hands of Nihlathak alone.
`},
	1258: {ID: 1258, Key: "am3", Value: `Maiden Spear`},
	1259: {ID: 1259, Key: "am4", Value: `Maiden Pike`},
	1260: {ID: 1260, Key: "am5", Value: `Maiden Javelin`},
	1261: {ID: 1261, Key: "am6", Value: `Ashwood Bow`},
	1262: {ID: 1262, Key: "am7", Value: `Ceremonial Bow`},
	1263: {ID: 1263, Key: "am8", Value: `Ceremonial Spear`},
	1264: {ID: 1264, Key: "am9", Value: `Ceremonial Pike`},
	1265: {ID: 1265, Key: "CompletingStopSiegeDru", Value: `The catapults have been silenced.`},
	1266: {ID: 1266, Key: "Tyrael's Mercy", Value: `Tyrael's Mercy`},
	1267: {ID: 1267, Key: "MercX162", Value: `Torkel`},
	1268: {ID: 1268, Key: "Que-hegan's Wisdom", Value: `Que-Hegan's Wisdom`},
	1269: {ID: 1269, Key: "Flowkrad's Howl", Value: `Flowkrad's Howl`},
	1270: {ID: 1270, Key: "Spiritforge", Value: `Spirit Forge`},
	1271: {ID: 1271, Key: "DeamonSteed", Value: `Demon Steed`},
	1272: {ID: 1272, Key: "MercX163", Value: `Wulfstan`},
	1273: {ID: 1273, Key: "A5Q3SuccessfulCain", Value: `62
For one so young, Anya commands 
great respect. Now that she is here, I 
will make it a point to talk to her about 
Mount Arreat.
 
You should do the same.
`},
	1274: {ID: 1274, Key: "MercX164", Value: `Alban`},
	1275: {ID: 1275, Key: "Hellmouth", Value: `Hellmouth`},
	1276: {ID: 1276, Key: "RescueQual-KehkAct5Ams", Value: `Follow me.`},
	1277: {ID: 1277, Key: "Eskillmanarecov", Value: `Mana Recovered: `},
	1278: {ID: 1278, Key: "chestr", Value: `Chest`},
	1279: {ID: 1279, Key: "MercX165", Value: `Barloc`},
	1280: {ID: 1280, Key: "MercX166", Value: `Bill`},
	1281: {ID: 1281, Key: "MercX167", Value: `Theodoric`},
	1282: {ID: 1282, Key: "Fierce", Value: `Fierce`},
	1283: {ID: 1283, Key: "Gaya Wand", Value: `Gaya Wand`},
	1284: {ID: 1284, Key: "of Iron Golem Creation", Value: `of Iron Golem`},
	1285: {ID: 1285, Key: "CainAct5IntroGossip1", Value: `39
I am amazed to find this place so 
untouched. Everything else in the path 
of Baal the Lord of Destruction lies in 
ruin. 
 
These Barbarians must indeed be the 
legendary guardians of Mount Arreat. 
They are a proud, hardy people. Don't 
expect to be greeted warmly -- 
strangers here rarely are.
 
Perhaps I can gain their trust. I'll spend 
some time with the townsfolk and try 
to understand them better. I'll let you 
know what I discover.
`},
	1286: {ID: 1286, Key: "A5Q5SuccessfulQualKehk", Value: `75
Besting the Ancients in battle is a 
mighty feat indeed. I hope this means 
you're ready to battle Baal.
`},
	1287: {ID: 1287, Key: "of Malice", Value: `of Malice`},
	1288: {ID: 1288, Key: "Of Stature", Value: `of Stature`},
	1289: {ID: 1289, Key: "Warpspear", Value: `Warpspear`},
	1290: {ID: 1290, Key: "LeavingTownAct5Pal", Value: `Baal. I'm coming for you.`},
	1291: {ID: 1291, Key: "Endlessshail", Value: `Endlesshail`},
	1292: {ID: 1292, Key: "CompletingLamEsenAss", Value: `One can't judge a book by its cover.`},
	1293: {ID: 1293, Key: "Of Badness", Value: `of Badness`},
	1294: {ID: 1294, Key: "CompletingLamEsenAss", Value: `Ormus... You have strange taste in books.`},
	1295: {ID: 1295, Key: "CompletingDefeatBaalAct5Ams", Value: `My work here is truly done.`},
	1296: {ID: 1296, Key: "Eskillchancetoafflict", Value: `Chance to afflict target: `},
	1297: {ID: 1297, Key: "strAmazonOnly", Value: `(Amazon Only)`},
	1298: {ID: 1298, Key: "Skillname240", Value: `Fire Claws`},
	1299: {ID: 1299, Key: "Ironward", Value: `Astreon's Iron Ward`},
	1300: {ID: 1300, Key: "Crainte Vomir", Value: `Crainte Vomir`},
	1301: {ID: 1301, Key: "Sarmichian Justice", Value: `Samurai Justice`},
	1302: {ID: 1302, Key: "ama", Value: `Ceremonial Javelin`},
	1303: {ID: 1303, Key: "amb", Value: `Matriarchal Bow`},
	1304: {ID: 1304, Key: "amc", Value: `Grand Matron Bow`},
	1305: {ID: 1305, Key: "amd", Value: `Matriarchal Spear`},
	1306: {ID: 1306, Key: "ame", Value: `Matriarchal Pike`},
	1307: {ID: 1307, Key: "amf", Value: `Matriarchal Javelin`},
	1308: {ID: 1308, Key: "Jack's", Value: `Jack's`},
	1309: {ID: 1309, Key: "Continuous", Value: `Everlasting`},
	1310: {ID: 1310, Key: "Razorswitch", Value: `Razorswitch`},
	1311: {ID: 1311, Key: "of Substinence", Value: `of Sustenance`},
	1312: {ID: 1312, Key: "Talberd's Law", Value: `Talberd's Law`},
	1313: {ID: 1313, Key: "Divine", Value: `Divine`},
	1314: {ID: 1314, Key: "Phoenix Egg", Value: `Phoenix Egg`},
	1315: {ID: 1315, Key: "Skillname230", Value: `Molten Boulder`},
	1316: {ID: 1316, Key: "A5Q3FoundAnyaQualKehk", Value: `70
The snake has slipped our grasp! While 
you were gone, Nihlathak disappeared.
 
I'll bet Anya knows how to track him 
down.
`},
	1317: {ID: 1317, Key: "of Hope", Value: `of Hope`},
	1318: {ID: 1318, Key: "ba1", Value: `Jawbone Cap`},
	1319: {ID: 1319, Key: "ba2", Value: `Fanged Helm`},
	1320: {ID: 1320, Key: "ba3", Value: `Horned Helm`},
	1321: {ID: 1321, Key: "ba4", Value: `Assault Helmet`},
	1322: {ID: 1322, Key: "ba5", Value: `Avenger Guard`},
	1323: {ID: 1323, Key: "ba6", Value: `Jawbone Visor`},
	1324: {ID: 1324, Key: "ba7", Value: `Lion Helm`},
	1325: {ID: 1325, Key: "ba8", Value: `Rage Mask`},
	1326: {ID: 1326, Key: "ba9", Value: `Savage Helmet`},
	1327: {ID: 1327, Key: "Great Wyrm's", Value: `Great Wyrm's`},
	1328: {ID: 1328, Key: "Punishing", Value: `Punishing`},
	1329: {ID: 1329, Key: "of Enchantment", Value: `of Enchantment`},
	1330: {ID: 1330, Key: "of Charged Bolts", Value: `of Charged Bolt`},
	1331: {ID: 1331, Key: "Catgut", Value: `Catgut`},
	1332: {ID: 1332, Key: "Skillname222", Value: `Raven`},
	1333: {ID: 1333, Key: "Runeword10", Value: `Brand`},
	1334: {ID: 1334, Key: "Runeword11", Value: `Breath of the Dying`},
	1335: {ID: 1335, Key: "Runeword12", Value: `Broken Promise`},
	1336: {ID: 1336, Key: "Runeword13", Value: `Call to Arms`},
	1337: {ID: 1337, Key: "Runeword14", Value: `Chains of Honor`},
	1338: {ID: 1338, Key: "Runeword15", Value: `Chance`},
	1339: {ID: 1339, Key: "Runeword16", Value: `Chaos`},
	1340: {ID: 1340, Key: "Runeword17", Value: `Crescent Moon`},
	1341: {ID: 1341, Key: "Runeword18", Value: `Darkness`},
	1342: {ID: 1342, Key: "Runeword19", Value: `Daylight`},
	1343: {ID: 1343, Key: "of Self-Repair", Value: `of Self-Repair`},
	1344: {ID: 1344, Key: "Blighting", Value: `Blighting`},
	1345: {ID: 1345, Key: "Langer Briser", Value: `Langer Briser`},
	1346: {ID: 1346, Key: "of Ire", Value: `of Ire`},
	1347: {ID: 1347, Key: "Ethereal edge", Value: `Ethereal Edge`},
	1348: {ID: 1348, Key: "A5Q2EarlyReturnLarzuk", Value: `76
As a kid, I used to play soldier among 
the barricades on the mountain. 
There's no easy way through that maze 
of walls.
`},
	1349: {ID: 1349, Key: "Runeword20", Value: `Death`},
	1350: {ID: 1350, Key: "as1", Value: `Wraps`},
	1351: {ID: 1351, Key: "as2", Value: `Knuckles`},
	1352: {ID: 1352, Key: "as3", Value: `Slashers`},
	1353: {ID: 1353, Key: "as4", Value: `Splay`},
	1354: {ID: 1354, Key: "as5", Value: `Hook`},
	1355: {ID: 1355, Key: "A5Q2AfterInitNihlathak", Value: `94
You have proven you can take life, 
warrior, but can you save it as well?
`},
	1356: {ID: 1356, Key: "as6", Value: `Shank`},
	1357: {ID: 1357, Key: "as7", Value: `Claws`},
	1358: {ID: 1358, Key: "r20L", Value: `Lem`},
	1359: {ID: 1359, Key: "Runeword21", Value: `Deception`},
	1360: {ID: 1360, Key: "Runeword22", Value: `Delerium`},
	1361: {ID: 1361, Key: "Runeword23", Value: `Desire`},
	1362: {ID: 1362, Key: "Runeword24", Value: `Despair`},
	1363: {ID: 1363, Key: "Runeword25", Value: `Destruction`},
	1364: {ID: 1364, Key: "Runeword26", Value: `Doom`},
	1365: {ID: 1365, Key: "Runeword27", Value: `Dragon`},
	1366: {ID: 1366, Key: "baa", Value: `Slayer Guard`},
	1367: {ID: 1367, Key: "bab", Value: `Carnage Helm`},
	1368: {ID: 1368, Key: "bac", Value: `Fury Visor`},
	1369: {ID: 1369, Key: "QualKehkAct5IntroDruGossip1", Value: `42
A Druid in Harrogath! Have things truly 
come to this?
 
After the Mage Wars, I assumed Druids 
would never be seen in Harrogath 
again. You take a big chance coming 
here! 
 
To be honest, I have never been 
comfortable with your shape-shifting 
kind, but I do respect your search for 
balance and peace. So, if you trust us 
enough to enter our gates, I trust you 
enough to let you stay.
 
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`},
	1370: {ID: 1370, Key: "bad", Value: `Destroyer Helm`},
	1371: {ID: 1371, Key: "bae", Value: `Conqueror Crown`},
	1372: {ID: 1372, Key: "baf", Value: `Guardian Crown`},
	1373: {ID: 1373, Key: "Runeword28", Value: `Dread`},
	1374: {ID: 1374, Key: "r21L", Value: `Pul`},
	1375: {ID: 1375, Key: "Runeword29", Value: `Dream`},
	1376: {ID: 1376, Key: "Runeword30", Value: `Duress`},
	1377: {ID: 1377, Key: "wrb", Value: `Wrist Blade`},
	1378: {ID: 1378, Key: "Runeword31", Value: `Edge`},
	1379: {ID: 1379, Key: "Runeword32", Value: `Elation`},
	1380: {ID: 1380, Key: "Runeword33", Value: `Enigma`},
	1381: {ID: 1381, Key: "Runeword34", Value: `Enlightenment`},
	1382: {ID: 1382, Key: "Runeword35", Value: `Envy`},
	1383: {ID: 1383, Key: "Runeword36", Value: `Eternity`},
	1384: {ID: 1384, Key: "Runeword37", Value: `Exile`},
	1385: {ID: 1385, Key: "Runeword38", Value: `Faith`},
	1386: {ID: 1386, Key: "Runeword39", Value: `Famine`},
	1387: {ID: 1387, Key: "Runeword40", Value: `Flame`},
	1388: {ID: 1388, Key: "Runeword41", Value: `Fortitude`},
	1389: {ID: 1389, Key: "Runeword42", Value: `Fortune`},
	1390: {ID: 1390, Key: "r22L", Value: `Um`},
	1391: {ID: 1391, Key: "Runeword43", Value: `Friendship`},
	1392: {ID: 1392, Key: "Runeword44", Value: `Fury`},
	1393: {ID: 1393, Key: "Runeword45", Value: `Gloom`},
	1394: {ID: 1394, Key: "Runeword46", Value: `Glory`},
	1395: {ID: 1395, Key: "Runeword47", Value: `Grief`},
	1396: {ID: 1396, Key: "Runeword48", Value: `Hand of Justice`},
	1397: {ID: 1397, Key: "Runeword49", Value: `Harmony`},
	1398: {ID: 1398, Key: "Jester's", Value: `Jester's`},
	1399: {ID: 1399, Key: "Lucky", Value: `Lucky`},
	1400: {ID: 1400, Key: "Rusty", Value: `Rusty`},
	1401: {ID: 1401, Key: "Turquoise", Value: `Turquoise`},
	1402: {ID: 1402, Key: "Laden", Value: `Laden`},
	1403: {ID: 1403, Key: "Irresistible", Value: `Irresistible`},
	1404: {ID: 1404, Key: "Crude", Value: `Crude`},
	1405: {ID: 1405, Key: "of Frozen Armor", Value: `of Frozen Armor`},
	1406: {ID: 1406, Key: "r23L", Value: `Mal`},
	1407: {ID: 1407, Key: "Fungal", Value: `Fungal`},
	1408: {ID: 1408, Key: "Impecable", Value: `Impeccable`},
	1409: {ID: 1409, Key: "A5Q1SuccessfulMalah", Value: `94
Oh...At last, the siege is ended. You 
truly are an angel.
 
I thank you.
`},
	1410: {ID: 1410, Key: "A5Q3SuccessfulQualKehk", Value: `140
Your rescue of Anya was quite an 
accomplishment.
`},
	1411: {ID: 1411, Key: "6cb", Value: `Great Bow`},
	1412: {ID: 1412, Key: "6bs", Value: `Shillelagh`},
	1413: {ID: 1413, Key: "of Grace and Power", Value: `of Grace and Power`},
	1414: {ID: 1414, Key: "of Anthrax", Value: `of Anthrax`},
	1415: {ID: 1415, Key: "Kenshi's", Value: `Kenshi's`},
	1416: {ID: 1416, Key: "Of the Imbecile", Value: `of the Imbecile`},
	1417: {ID: 1417, Key: "of Eruption", Value: `of Fissure`},
	1418: {ID: 1418, Key: "Hand of Blessed Light", Value: `Hand of Blessed Light`},
	1419: {ID: 1419, Key: "Heart Carver", Value: `Heart Carver`},
	1420: {ID: 1420, Key: "A5Q5AfterInitMalah", Value: `65
By reaching the summit, you cease 
being just a simple warrior. When you 
come back, you will be far more.
`},
	1421: {ID: 1421, Key: "Runeword1", Value: `Ancients' Pledge`},
	1422: {ID: 1422, Key: "r24L", Value: `Ist`},
	1423: {ID: 1423, Key: "Runeword2", Value: `Armageddon`},
	1424: {ID: 1424, Key: "Runeword3", Value: `Authority`},
	1425: {ID: 1425, Key: "Runeword4", Value: `Beast`},
	1426: {ID: 1426, Key: "Runeword5", Value: `Beauty`},
	1427: {ID: 1427, Key: "Runeword6", Value: `Black`},
	1428: {ID: 1428, Key: "6cs", Value: `Elder Staff`},
	1429: {ID: 1429, Key: "Runeword7", Value: `Blood`},
	1430: {ID: 1430, Key: "Runeword8", Value: `Bone`},
	1431: {ID: 1431, Key: "Runeword9", Value: `Bramble`},
	1432: {ID: 1432, Key: "Hawkeye", Value: `Hawkeye`},
	1433: {ID: 1433, Key: "of Icebolt", Value: `of Icebolt`},
	1434: {ID: 1434, Key: "Doppleganger's Shadow", Value: `Doppleganger's Shadow`},
	1435: {ID: 1435, Key: "Spiritkeeper", Value: `Spirit Keeper`},
	1436: {ID: 1436, Key: "The Swashbuckler", Value: `The Swashbuckler`},
	1437: {ID: 1437, Key: "Noxious", Value: `Noxious`},
	1438: {ID: 1438, Key: "r25L", Value: `Gul`},
	1439: {ID: 1439, Key: "AnyaGossip10", Value: `60
Our people believe that the Ancients 
protecting Mount Arreat have the 
power to stop Baal. Unfortunately, the 
Lord of Destruction has shown great 
power to undermine our faith.
`},
	1440: {ID: 1440, Key: "of Honor", Value: `of Honor`},
	1441: {ID: 1441, Key: "Of Twilight", Value: `of Twilight`},
	1442: {ID: 1442, Key: "Wartraveler", Value: `War Traveler`},
	1443: {ID: 1443, Key: "Rainbow Facet", Value: `Rainbow Facet`},
	1444: {ID: 1444, Key: "BarOnly", Value: `(Barbarian Only)`},
	1445: {ID: 1445, Key: "Ashrera's Wired Frame", Value: `Asheara's Wired Frame`},
	1446: {ID: 1446, Key: "Ravenlore", Value: `Ravenlore`},
	1447: {ID: 1447, Key: "Headhunter's Glory", Value: `Head Hunter's Glory`},
	1448: {ID: 1448, Key: "Spike Thorn", Value: `Spike Thorn`},
	1449: {ID: 1449, Key: "A5Q2SuccessfulAnya", Value: `107
I'm sure those men appreciate the 
freedom you gave them...as much as I 
do.
`},
	1450: {ID: 1450, Key: "Jadetalon", Value: `Jade Talon`},
	1451: {ID: 1451, Key: "of Static Field", Value: `of Static Field`},
	1452: {ID: 1452, Key: "Rings", Value: `Rings`},
	1453: {ID: 1453, Key: "Trang-Oul's Wing", Value: `Trang-Oul's Wing`},
	1454: {ID: 1454, Key: "r26L", Value: `Vex`},
	1455: {ID: 1455, Key: "qstsa5q31a", Value: `Rescue Anya.`},
	1456: {ID: 1456, Key: "Baal Crab to Stairs", Value: `Baal`},
	1457: {ID: 1457, Key: "Chaotic", Value: `Chaotic`},
	1458: {ID: 1458, Key: "FindingInifusDru", Value: `This dead tree teems with energy.`},
	1459: {ID: 1459, Key: "RescueCainDru", Value: `Deckard Cain, leave this place!`},
	1460: {ID: 1460, Key: "of Enchant", Value: `of Enchant`},
	1461: {ID: 1461, Key: "Runeword90", Value: `Night`},
	1462: {ID: 1462, Key: "Runeword91", Value: `Oath`},
	1463: {ID: 1463, Key: "Runeword92", Value: `Obedience`},
	1464: {ID: 1464, Key: "Runeword93", Value: `Oblivion`},
	1465: {ID: 1465, Key: "Runeword94", Value: `Obsession`},
	1466: {ID: 1466, Key: "Runeword95", Value: `Passion`},
	1467: {ID: 1467, Key: "Runeword96", Value: `Patience`},
	1468: {ID: 1468, Key: "Runeword97", Value: `Patter`},
	1469: {ID: 1469, Key: "Runeword98", Value: `Peace`},
	1470: {ID: 1470, Key: "r27L", Value: `Ohm`},
	1471: {ID: 1471, Key: "Runeword99", Value: `Voice of Reason`},
	1472: {ID: 1472, Key: "A5Q1AfterInitLarzuk", Value: `150
Uh...Did I mention there were catapults?
`},
	1473: {ID: 1473, Key: "Lancer's", Value: `Lancer's`},
	1474: {ID: 1474, Key: "Of Grounding", Value: `of Grounding`},
	1475: {ID: 1475, Key: "Of Suffering", Value: `of Suffering`},
	1476: {ID: 1476, Key: "A5Q4EarlyReturnCain", Value: `58
Ohh...This is a truly horrible turn of 
events.
 
I know it seems you have always been 
one step behind, my friend. But look at 
it this way...You have evil on the run.
`},
	1477: {ID: 1477, Key: "ItemexpED", Value: ` Elemental Damage`},
	1478: {ID: 1478, Key: "Drulan's Tongue", Value: `Drulan's Tongue`},
	1479: {ID: 1479, Key: "Souldrain", Value: `Soul Drainer`},
	1480: {ID: 1480, Key: "Halaberd's Reign", Value: `Halaberd's Reign`},
	1481: {ID: 1481, Key: "Trang-Oul's Guise", Value: `Trang-Oul's Guise`},
	1482: {ID: 1482, Key: "qstsa5q43a", Value: `Anya will personalize an item for you.`},
	1483: {ID: 1483, Key: "axf", Value: `Hatchet Hands`},
	1484: {ID: 1484, Key: "CompletingRadamentAss", Value: `Vengeance... for Atma.`},
	1485: {ID: 1485, Key: "EnterJailDru", Value: `Bars can't hold a force of nature.`},
	1486: {ID: 1486, Key: "r28L", Value: `Lo`},
	1487: {ID: 1487, Key: "SorOnly", Value: `(Sorceress Only)`},
	1488: {ID: 1488, Key: "Of Daring", Value: `of Daring`},
	1489: {ID: 1489, Key: "of Fire Ball", Value: `of Fire Ball`},
	1490: {ID: 1490, Key: "FindingGuardianTowerAss", Value: `Mephisto... I'm coming for you.`},
	1491: {ID: 1491, Key: "6hb", Value: `Blade Bow`},
	1492: {ID: 1492, Key: "EnteringTopMountAct5Pal", Value: `The summit - The Barbarian holy ground.`},
	1493: {ID: 1493, Key: "Skillname223", Value: `Poison Creeper`},
	1494: {ID: 1494, Key: "Skillname223", Value: `Werewolf`},
	1495: {ID: 1495, Key: "Hwanin's Refuge", Value: `Hwanin's Refuge`},
	1496: {ID: 1496, Key: "of Weaken", Value: `of Weaken`},
	1497: {ID: 1497, Key: "Skillname225", Value: `Lycanthropy`},
	1498: {ID: 1498, Key: "qstsa5q42a", Value: `Find Nihlathak in the Halls of Vaught.`},
	1499: {ID: 1499, Key: "Skillname226", Value: `Firestorm`},
	1500: {ID: 1500, Key: "M'avina's Caster", Value: `M'avina's Caster`},
	1501: {ID: 1501, Key: "Skillname227", Value: `Oak Sage`},
	1502: {ID: 1502, Key: "r29L", Value: `Sur`},
	1503: {ID: 1503, Key: "Frigid", Value: `Frigid`},
	1504: {ID: 1504, Key: "Skillname228", Value: `Summon Spirit Wolf`},
	1505: {ID: 1505, Key: "Skillname229", Value: `Werebear`},
	1506: {ID: 1506, Key: "Skillname231", Value: `Arctic Blast`},
	1507: {ID: 1507, Key: "Skillname232", Value: `Carrion Vine`},
	1508: {ID: 1508, Key: "of Ages", Value: `of Ages`},
	1509: {ID: 1509, Key: "Skillname233", Value: `Feral Rage`},
	1510: {ID: 1510, Key: "Skillname234", Value: `Maul`},
	1511: {ID: 1511, Key: "Venom Grip", Value: `Venom Grip`},
	1512: {ID: 1512, Key: "6l7", Value: `Crusader Bow`},
	1513: {ID: 1513, Key: "6hx", Value: `Colossus Crossbow`},
	1514: {ID: 1514, Key: "Skillname235", Value: `Fissure`},
	1515: {ID: 1515, Key: "Skillname236", Value: `Cyclone Armor`},
	1516: {ID: 1516, Key: "Skillname237", Value: `Heart of Wolverine`},
	1517: {ID: 1517, Key: "Scorching", Value: `Condensing`},
	1518: {ID: 1518, Key: "of Lightning Strike", Value: `of Lightning Strike`},
	1519: {ID: 1519, Key: "Skillname238", Value: `Summon Dire Wolf`},
	1520: {ID: 1520, Key: "Skillname239", Value: `Rabies`},
	1521: {ID: 1521, Key: "Skillname241", Value: `Twister`},
	1522: {ID: 1522, Key: "Of Dispatch", Value: `of Dispatch`},
	1523: {ID: 1523, Key: "Riftslash", Value: `Riftslash`},
	1524: {ID: 1524, Key: "Skillname242", Value: `Solar Creeper`},
	1525: {ID: 1525, Key: "Runeword50", Value: `Hatred`},
	1526: {ID: 1526, Key: "Runeword51", Value: `Heart of the Oak`},
	1527: {ID: 1527, Key: "Runeword52", Value: `Heaven's Will`},
	1528: {ID: 1528, Key: "A5Q6SuccessfulCain", Value: `40
I knew in time you would defeat Baal. 
You have done everything you set out 
to do, my friend.
 
Ever since you rescued me from 
Tristram, I have believed in you. It has 
been a supreme honor to aid you along 
the way. 
 
So...The Worldstone was corrupted by 
Baal. And now Tyrael must destroy it. 
Worry not. Through whatever lies 
ahead I have faith that the Light will 
guide us both.
 
Go, now, back to the Worldstone 
chamber, and enter the portal Tyrael 
has opened for you.
`},
	1529: {ID: 1529, Key: "Runeword53", Value: `Holy Tears`},
	1530: {ID: 1530, Key: "Runeword54", Value: `Holy Thunder`},
	1531: {ID: 1531, Key: "Runeword55", Value: `Honor`},
	1532: {ID: 1532, Key: "Runeword56", Value: `Revenge`},
	1533: {ID: 1533, Key: "Runeword57", Value: `Humility`},
	1534: {ID: 1534, Key: "A5Q1SuccessfulQualKehk", Value: `68
So...You managed to stop the siege.
 
You are more powerful than I gave you 
credit for. You have rightfully earned 
my respect.
`},
	1535: {ID: 1535, Key: "Runeword58", Value: `Hunger`},
	1536: {ID: 1536, Key: "Runeword59", Value: `Ice`},
	1537: {ID: 1537, Key: "Spiteful", Value: `Spiteful`},
	1538: {ID: 1538, Key: "Caretaker's", Value: `Caretaker's`},
	1539: {ID: 1539, Key: "Todesfaelle Flamme", Value: `Todesfaelle Flamme`},
	1540: {ID: 1540, Key: "CompletingStopSiegeNec", Value: `My, my, what a messy little demon.`},
	1541: {ID: 1541, Key: "Runeword60", Value: `Infinity`},
	1542: {ID: 1542, Key: "Runeword61", Value: `Innocence`},
	1543: {ID: 1543, Key: "Runeword62", Value: `Insight`},
	1544: {ID: 1544, Key: "Runeword63", Value: `Jealousy`},
	1545: {ID: 1545, Key: "Runeword64", Value: `Judgement`},
	1546: {ID: 1546, Key: "Runeword65", Value: `King's Grace`},
	1547: {ID: 1547, Key: "Runeword66", Value: `Kingslayer`},
	1548: {ID: 1548, Key: "Runeword67", Value: `Knight's Vigil`},
	1549: {ID: 1549, Key: "Runeword68", Value: `Knowledge`},
	1550: {ID: 1550, Key: "Runeword69", Value: `Last Wish`},
	1551: {ID: 1551, Key: "Trickster's", Value: `Trickster's`},
	1552: {ID: 1552, Key: "Vigorous", Value: `Vigorous`},
	1553: {ID: 1553, Key: "Pure", Value: `Pure`},
	1554: {ID: 1554, Key: "of Teleport Shield", Value: `of Teleport Shield`},
	1555: {ID: 1555, Key: "6lb", Value: `Shadow Bow`},
	1556: {ID: 1556, Key: "of Freedom", Value: `of Freedom`},
	1557: {ID: 1557, Key: "Runeword70", Value: `Law`},
	1558: {ID: 1558, Key: "Runeword71", Value: `Lawbringer`},
	1559: {ID: 1559, Key: "Runeword72", Value: `Leaf`},
	1560: {ID: 1560, Key: "Runeword73", Value: `Lightning`},
	1561: {ID: 1561, Key: "Runeword74", Value: `Lionheart`},
	1562: {ID: 1562, Key: "Runeword75", Value: `Lore`},
	1563: {ID: 1563, Key: "Runeword76", Value: `Love`},
	1564: {ID: 1564, Key: "Runeword77", Value: `Loyalty`},
	1565: {ID: 1565, Key: "Runeword78", Value: `Lust`},
	1566: {ID: 1566, Key: "Runeword79", Value: `Madness`},
	1567: {ID: 1567, Key: "Shouting", Value: `Howling`},
	1568: {ID: 1568, Key: "of Truth", Value: `of Truth`},
	1569: {ID: 1569, Key: "WeaponDescH2H", Value: `Claw Class`},
	1570: {ID: 1570, Key: "Arioc's Needle", Value: `Arioc's Needle`},
	1571: {ID: 1571, Key: "Mordoc's marauder", Value: `Mordoc's Marauder`},
	1572: {ID: 1572, Key: "6ls", Value: `Stalagmite`},
	1573: {ID: 1573, Key: "A5Q3FoundAnyaMalah", Value: `56
So! That snake Nihlathak was behind 
Anya's disappearance...and he trapped 
her with a freezing curse.
 
Here. Take this potion to Anya and give 
it to her. That should release her.
`},
	1574: {ID: 1574, Key: "Runeword81", Value: `Malice`},
	1575: {ID: 1575, Key: "Runeword82", Value: `Melody`},
	1576: {ID: 1576, Key: "6lw", Value: `Hydra Bow`},
	1577: {ID: 1577, Key: "6lx", Value: `Pellet Bow`},
	1578: {ID: 1578, Key: "Runeword83", Value: `Memory`},
	1579: {ID: 1579, Key: "Runeword84", Value: `Mist`},
	1580: {ID: 1580, Key: "Runeword85", Value: `Morning`},
	1581: {ID: 1581, Key: "Runeword86", Value: `Mystery`},
	1582: {ID: 1582, Key: "Runeword87", Value: `Myth`},
	1583: {ID: 1583, Key: "Runeword88", Value: `Nadir`},
	1584: {ID: 1584, Key: "Runeword89", Value: `Nature's Kingdom`},
	1585: {ID: 1585, Key: "Gaea's", Value: `Gaean`},
	1586: {ID: 1586, Key: "Chromatic", Value: `Chromatic`},
	1587: {ID: 1587, Key: "of Traveling", Value: `of Transportation`},
	1588: {ID: 1588, Key: "of Joy", Value: `of Joy`},
	1589: {ID: 1589, Key: "of Grounding", Value: `of Grounding`},
	1590: {ID: 1590, Key: "Of the Stars", Value: `of the Stars`},
	1591: {ID: 1591, Key: "of Twister", Value: `of Twister`},
	1592: {ID: 1592, Key: "Of the Beast", Value: `of the Beast`},
	1593: {ID: 1593, Key: "6mx", Value: `Gorgon Crossbow`},
	1594: {ID: 1594, Key: "of Item Finding", Value: `of Find Item`},
	1595: {ID: 1595, Key: "Soulreaper", Value: `Soul Reaper`},
	1596: {ID: 1596, Key: "Infernostride", Value: `Infernostride`},
	1597: {ID: 1597, Key: "Gloomstrap", Value: `Gloom's Trap`},
	1598: {ID: 1598, Key: "Steel Carapice", Value: `Steel Carapace`},
	1599: {ID: 1599, Key: "Riftlash", Value: `Riftlash`},
	1600: {ID: 1600, Key: "Ancient Statue 3", Value: `Korlic the Protector`},
	1601: {ID: 1601, Key: "Ancient Statue 2", Value: `Madawc the Guardian`},
	1602: {ID: 1602, Key: "Ancient Statue 1", Value: `Talic the Defender`},
	1603: {ID: 1603, Key: "FreezingIzualAss", Value: `Corruption... take flight.`},
	1604: {ID: 1604, Key: "Skillname243", Value: `Hunger`},
	1605: {ID: 1605, Key: "Skillname244", Value: `Shock Wave`},
	1606: {ID: 1606, Key: "Skillname245", Value: `Volcano`},
	1607: {ID: 1607, Key: "Shadefalcon", Value: `Shade Falcon`},
	1608: {ID: 1608, Key: "7b7", Value: `Champion Sword`},
	1609: {ID: 1609, Key: "7b8", Value: `Winged Axe`},
	1610: {ID: 1610, Key: "CompletingBeneathCityDru", Value: `From trash to treasure...`},
	1611: {ID: 1611, Key: "Skillname246", Value: `Tornado`},
	1612: {ID: 1612, Key: "Skillname247", Value: `Spirit of Barbs`},
	1613: {ID: 1613, Key: "Skillname248", Value: `Summon Grizzly`},
	1614: {ID: 1614, Key: "r30L", Value: `Ber`},
	1615: {ID: 1615, Key: "Baal Subject 1", Value: `Colenzo the Annihilator`},
	1616: {ID: 1616, Key: "Baal Subject 2", Value: `Achmel the Cursed`},
	1617: {ID: 1617, Key: "Baal Subject 3", Value: `Bartuc the Bloody`},
	1618: {ID: 1618, Key: "of Acceleration", Value: `of Acceleration`},
	1619: {ID: 1619, Key: "Quixotic", Value: `Quixotic`},
	1620: {ID: 1620, Key: "of Plague Javelin", Value: `of Plague Javelin`},
	1621: {ID: 1621, Key: "QualKehkGossip1", Value: `67
It would be an honor to have a warrior 
of the Light fighting side-by-side with 
my men.
 
I can see your faith gives you great 
strength, Paladin, but don't expect it to 
keep you out of harm's way.
`},
	1622: {ID: 1622, Key: "QualKehkGossip2", Value: `61
Harrogath has great need of your 
powers, noble Druid. However, in the 
face of this supernatural onslaught, 
are your natural powers up to the 
task?
`},
	1623: {ID: 1623, Key: "QualKehkGossip3", Value: `44
The death of Malah's son was a great 
tragedy. He was our finest archer.
 
While leading a successful campaign 
against Baal's forces, he was impaled 
on a demon's spear. The wound was 
such that...Well, even Malah herself 
acknowledges that quick death was a 
blessing.
`},
	1624: {ID: 1624, Key: "QualKehkGossip4", Value: `57
We have lost many well-trained warriors 
to Baal's siege machines. Their range is 
great. Though, they are vulnerable if 
you close the distance quickly enough.
`},
	1625: {ID: 1625, Key: "QualKehkGossip5", Value: `96
Baal's legions seem countless, but 
slaying their commanders takes some 
of the fight out of them.
`},
	1626: {ID: 1626, Key: "QualKehkGossip6", Value: `58
Early on, parties of our best scouts 
were ambushed by demons that 
spawned from the very air around 
them. Survivors often mentioned a 
strange creature floating in the 
distance.
 
Perhaps taking it down could prevent 
some nasty surprises.
`},
	1627: {ID: 1627, Key: "QualKehkGossip7", Value: `57
This is unlike any battle I have ever 
fought. While we ration food and 
water, the demon hordes feast nightly 
on the flesh and blood of our dead.
`},
	1628: {ID: 1628, Key: "QualKehkGossip8", Value: `48
Larzuk is a talented blacksmith, but his 
head is full of some strange ideas. 
 
Just the other day he came to me with a 
plan to break the siege. He wanted to 
fill large pipes with exploding powders 
and steel balls, then...Well, like I said, 
strange.
`},
	1629: {ID: 1629, Key: "QualKehkGossip9", Value: `46
Our Elders were wise leaders in more 
peaceful times, but now the survival of 
our people has fallen to me. My men 
and I will fight to the death, but there's 
no way to ensure the outcome.
 
I used to believe that nothing could 
break through our guard and assault 
the holy mountain. I know now that I 
was horribly mistaken.
`},
	1630: {ID: 1630, Key: "6s7", Value: `Diamond Bow`},
	1631: {ID: 1631, Key: "r31L", Value: `Jo`},
	1632: {ID: 1632, Key: "Terra's", Value: `Terrene`},
	1633: {ID: 1633, Key: "Master's", Value: `Master's`},
	1634: {ID: 1634, Key: "of Hydra Shield", Value: `of Hydra Shield`},
	1635: {ID: 1635, Key: "ne1", Value: `Preserved Head`},
	1636: {ID: 1636, Key: "ne2", Value: `Zombie Head`},
	1637: {ID: 1637, Key: "ne3", Value: `Unraveller Head`},
	1638: {ID: 1638, Key: "ne4", Value: `Gargoyle Head`},
	1639: {ID: 1639, Key: "A5Q3EarlyReturnMalah", Value: `52
If it were anyone else, I would assume 
her dead. However, Anya is not so 
easily beaten. I know she must be alive.
`},
	1640: {ID: 1640, Key: "ne5", Value: `Demon Head`},
	1641: {ID: 1641, Key: "ne6", Value: `Mummified Trophy`},
	1642: {ID: 1642, Key: "A5Q6SuccessfulTyrael", Value: `40
I am impressed, mortal. You have 
overcome the greatest challenge this 
world has ever faced and defeated the 
last of the Prime Evils. However, we are 
too late to save the Worldstone. Baal's 
destructive touch has corrupted it 
completely.
 
Given enough time, the Worldstone's 
energies will drain away and the 
barriers between the worlds will 
shatter -- the powers of Hell will flood 
into this...Sanctuary...and eradicate 
your people and everything you've 
labored to build.
 
Therefore, I must destroy the corrupted 
Worldstone before the powers of Hell 
take root. This act will change your 
world forever -- with consequences 
even I cannot foresee. However, it is 
the only way to ensure mankind's 
survival.
 
Go now, mortal. I have opened a portal 
that will lead you to safety.
 
May the Eternal Light shine upon you 
and your descendants for what you've 
done this day. The continued survival 
of mankind is your legacy! Above all 
else, you have earned a rest from this 
endless battle.
`},
	1643: {ID: 1643, Key: "ne7", Value: `Fetish Trophy`},
	1644: {ID: 1644, Key: "ne8", Value: `Sexton Trophy`},
	1645: {ID: 1645, Key: "ne9", Value: `Cantor Trophy`},
	1646: {ID: 1646, Key: "r32L", Value: `Cham`},
	1647: {ID: 1647, Key: "Paleocene", Value: `Faithful`},
	1648: {ID: 1648, Key: "of Chain Lightning", Value: `of Chain Lightning`},
	1649: {ID: 1649, Key: "of Storms", Value: `of Storms`},
	1650: {ID: 1650, Key: "7ba", Value: `Silver-edged Axe`},
	1651: {ID: 1651, Key: "7ar", Value: `Suwayyah`},
	1652: {ID: 1652, Key: "Warder's", Value: `Warden's`},
	1653: {ID: 1653, Key: "of Concentration", Value: `of Concentration`},
	1654: {ID: 1654, Key: "of Poison Novas", Value: `of Poison Nova`},
	1655: {ID: 1655, Key: "Darkfear", Value: `Darkfear`},
	1656: {ID: 1656, Key: "of Ice Blast", Value: `of Ice Blast`},
	1657: {ID: 1657, Key: "7ax", Value: `Small Crescent`},
	1658: {ID: 1658, Key: "Baranar's Star", Value: `Baranar's Star`},
	1659: {ID: 1659, Key: "EskillKickSing", Value: ` Kick`},
	1660: {ID: 1660, Key: "7bk", Value: `Winged Knife`},
	1661: {ID: 1661, Key: "7bl", Value: `Legend Spike`},
	1662: {ID: 1662, Key: "r33L", Value: `Zod`},
	1663: {ID: 1663, Key: "brz", Value: `Brain`},
	1664: {ID: 1664, Key: "Tal Rasha's Howling Wind", Value: `Tal Rasha's Guardianship`},
	1665: {ID: 1665, Key: "Baal Subject 4", Value: `Ventar the Unholy`},
	1666: {ID: 1666, Key: "Baal Subject 5", Value: `Lister the Tormentor`},
	1667: {ID: 1667, Key: "7br", Value: `Mancatcher`},
	1668: {ID: 1668, Key: "7bs", Value: `Conquest Sword`},
	1669: {ID: 1669, Key: "7bt", Value: `Decapitator`},
	1670: {ID: 1670, Key: "6sb", Value: `Spider Bow`},
	1671: {ID: 1671, Key: "Blademaster", Value: `Blade Master`},
	1672: {ID: 1672, Key: "7bw", Value: `Lich Wand`},
	1673: {ID: 1673, Key: "6rx", Value: `Demon Crossbow`},
	1674: {ID: 1674, Key: "Baal Subject 6", Value: `The Butcher`},
	1675: {ID: 1675, Key: "KillingdDiabloBar", Value: `Eternal suffering would be too brief for you, Diablo.`},
	1676: {ID: 1676, Key: "EnteringTopMountAct5Sor", Value: `The fabled home of the Ancients.`},
	1677: {ID: 1677, Key: "7cl", Value: `Truncheon`},
	1678: {ID: 1678, Key: "7cm", Value: `Highland Blade`},
	1679: {ID: 1679, Key: "Pierre Tombale Couant", Value: `Pierre Tombale Couant`},
	1680: {ID: 1680, Key: "Robin's Yolk", Value: `Robin's Yolk`},
	1681: {ID: 1681, Key: "btl", Value: `Blade Talons`},
	1682: {ID: 1682, Key: "Orphan's Call", Value: `Orphan's Call`},
	1683: {ID: 1683, Key: "7cr", Value: `Phase Blade`},
	1684: {ID: 1684, Key: "7cs", Value: `Battle Cestus`},
	1685: {ID: 1685, Key: "6ss", Value: `Walking Stick`},
	1686: {ID: 1686, Key: "nea", Value: `Hierophant Trophy`},
	1687: {ID: 1687, Key: "neb", Value: `Minion Skull`},
	1688: {ID: 1688, Key: "7dg", Value: `Bone Knife`},
	1689: {ID: 1689, Key: "6sw", Value: `Ward Bow`},
	1690: {ID: 1690, Key: "7di", Value: `Mithril Point`},
	1691: {ID: 1691, Key: "nec", Value: `Hellspawn Skull`},
	1692: {ID: 1692, Key: "ned", Value: `Overseer Skull`},
	1693: {ID: 1693, Key: "nee", Value: `Succubus Skull`},
	1694: {ID: 1694, Key: "nef", Value: `Bloodlord Skull`},
	1695: {ID: 1695, Key: "Marshal's", Value: `Marshal's`},
	1696: {ID: 1696, Key: "Bloody", Value: `Bloody`},
	1697: {ID: 1697, Key: "of Meteor", Value: `of Meteor`},
	1698: {ID: 1698, Key: "of Howling", Value: `of Howl`},
	1699: {ID: 1699, Key: "Of the Ocean", Value: `of the Ocean`},
	1700: {ID: 1700, Key: "Windstrike", Value: `Windstrike`},
	1701: {ID: 1701, Key: "ci0", Value: `Circlet`},
	1702: {ID: 1702, Key: "ci1", Value: `Coronet`},
	1703: {ID: 1703, Key: "ci2", Value: `Tiara`},
	1704: {ID: 1704, Key: "ces", Value: `Cestus`},
	1705: {ID: 1705, Key: "A5Q1SuccessfulAnya", Value: `108
I was starting to wonder how long it 
would take you to stop those beasts.
 
Good job.
`},
	1706: {ID: 1706, Key: "7h7", Value: `Great Poleaxe`},
	1707: {ID: 1707, Key: "ci3", Value: `Diadem`},
	1708: {ID: 1708, Key: "Bloodtree Stump", Value: `Bloodtree Stump`},
	1709: {ID: 1709, Key: "Bowyer's", Value: `Bowyer's`},
	1710: {ID: 1710, Key: "A5Q1EarlyReturnAnya", Value: `66
Those demons have been out there 
since before your arrival. Can you do 
nothing to stop them?
 
Your task is a simple one, warrior. Find 
Shenk and destroy him.
`},
	1711: {ID: 1711, Key: "Evil", Value: `Evil`},
	1712: {ID: 1712, Key: "Lacerator", Value: `Lacerator`},
	1713: {ID: 1713, Key: "Timeless", Value: `Timeless`},
	1714: {ID: 1714, Key: "Strongoak", Value: `Strong Oak`},
	1715: {ID: 1715, Key: "7fb", Value: `Colossus Sword`},
	1716: {ID: 1716, Key: "7fc", Value: `Hydra Edge`},
	1717: {ID: 1717, Key: "Realgar", Value: `Realgar`},
	1718: {ID: 1718, Key: "of Life Everlasting", Value: `of Life Everlasting`},
	1719: {ID: 1719, Key: "Of Quickening", Value: `of Quickening`},
	1720: {ID: 1720, Key: "Null", Value: `Null`},
	1721: {ID: 1721, Key: "of Attraction", Value: `of Attract`},
	1722: {ID: 1722, Key: "of Fending", Value: `of Fending`},
	1723: {ID: 1723, Key: "Raven Frost", Value: `Raven Frost`},
	1724: {ID: 1724, Key: "of Glacial Spike", Value: `of Glacial Spike`},
	1725: {ID: 1725, Key: "7fl", Value: `Scourge`},
	1726: {ID: 1726, Key: "of Life Tap", Value: `of Life Tap`},
	1727: {ID: 1727, Key: "Steelshade", Value: `Steel Shade`},
	1728: {ID: 1728, Key: "Eaglehorn", Value: `Eaglehorn`},
	1729: {ID: 1729, Key: "Cow King's Hide", Value: `Cow King's Hide`},
	1730: {ID: 1730, Key: "7ga", Value: `Champion Axe`},
	1731: {ID: 1731, Key: "Shimmering", Value: `Shimmering`},
	1732: {ID: 1732, Key: "Earthshifter", Value: `Earth Shifter`},
	1733: {ID: 1733, Key: "7gd", Value: `Colossus Blade`},
	1734: {ID: 1734, Key: "strlastcinematic", Value: `DESTRUCTION'S END`},
	1735: {ID: 1735, Key: "Terra", Value: `Terra's Guardian`},
	1736: {ID: 1736, Key: "Powered", Value: `Powered`},
	1737: {ID: 1737, Key: "A5Q3AfterInitLarzuk", Value: `57
Anya is an amazing alchemist, 
especially for her young age. As long 
as I've known her, she's never let 
anything stop her from pursuing what 
she believed in. 
 
I wouldn't doubt that Nihlathak is 
involved. Ever since her father died, 
they haven't gotten along.
`},
	1738: {ID: 1738, Key: "7gi", Value: `Glorious Axe`},
	1739: {ID: 1739, Key: "Tal Rasha's Adjudication", Value: `Tal Rasha's Adjudication`},
	1740: {ID: 1740, Key: "EskillPerBlade", Value: ` per blade`},
	1741: {ID: 1741, Key: "A5Q2EarlyReturnCain", Value: `63
If you are having trouble finding 
Qual-Kehk's soldiers, you should talk to 
Malah. She healed those who made it 
back before. Perhaps she would have 
some advice.
`},
	1742: {ID: 1742, Key: "7gm", Value: `Thunder Maul`},
	1743: {ID: 1743, Key: "7gl", Value: `Ghost Glaive`},
	1744: {ID: 1744, Key: "Of Vengence", Value: `of Vengence`},
	1745: {ID: 1745, Key: "of Zeal", Value: `of Zeal`},
	1746: {ID: 1746, Key: "7ha", Value: `Tomahawk`},
	1747: {ID: 1747, Key: "Gulletwound", Value: `Gulletwound`},
	1748: {ID: 1748, Key: "7gs", Value: `Balrog Blade`},
	1749: {ID: 1749, Key: "6ws", Value: `Archon Staff`},
	1750: {ID: 1750, Key: "of Confusion", Value: `of Confusion`},
	1751: {ID: 1751, Key: "Foggy", Value: `Foggy`},
	1752: {ID: 1752, Key: "7gw", Value: `Unearthed Wand`},
	1753: {ID: 1753, Key: "Aldur's Stony Gaze", Value: `Aldur's Stony Gaze`},
	1754: {ID: 1754, Key: "Ghost Liberator", Value: `Ghost Liberator`},
	1755: {ID: 1755, Key: "Boneshade", Value: `Boneshade`},
	1756: {ID: 1756, Key: "McAuley's Folly", Value: `Sander's Folly`},
	1757: {ID: 1757, Key: "qstsa5q62a", Value: `Kill Baal in the Worldstone Chamber before he corrupts it.`},
	1758: {ID: 1758, Key: "qstsa5q62b", Value: `Kill Baal's Minions.`},
	1759: {ID: 1759, Key: "Lightsabre", Value: `Lightsabre`},
	1760: {ID: 1760, Key: "KillingdDiabloAss", Value: `A hero's mistake is finally corrected.`},
	1761: {ID: 1761, Key: "CompletingStopSiegePal", Value: `Harrogath is free of your kind, demon.`},
	1762: {ID: 1762, Key: "Godly", Value: `Godly`},
	1763: {ID: 1763, Key: "CompletingNihlathakAct5Pal", Value: `Nihlathak. What led you to this end?`},
	1764: {ID: 1764, Key: "CompletingNihlathakAct5Dru", Value: `Betrayer, you've reaped your reward.`},
	1765: {ID: 1765, Key: "CompletingDefeatBaalAct5Nec", Value: `Baal, never doubt my skills.`},
	1766: {ID: 1766, Key: "cm1", Value: `Small Charm`},
	1767: {ID: 1767, Key: "cm2", Value: `Large Charm`},
	1768: {ID: 1768, Key: "cm3", Value: `Grand Charm`},
	1769: {ID: 1769, Key: "Headstriker", Value: `Headstriker`},
	1770: {ID: 1770, Key: "Skillld222", Value: `the eyes of your enemies
summon ravens to peck out`},
	1771: {ID: 1771, Key: "Skillld223", Value: `disease to all it contacts
summon a vine that spreads`},
	1772: {ID: 1772, Key: "Ancient Eye", Value: `Ancient Eye`},
	1773: {ID: 1773, Key: "LeavingTownAct5Sor", Value: `My magic will break the siege.`},
	1774: {ID: 1774, Key: "Skillld223", Value: `transform into a werewolf`},
	1775: {ID: 1775, Key: "Skillld225", Value: `when in werewolf or werebear form
passive - improves duration and life`},
	1776: {ID: 1776, Key: "Skillld226", Value: `unleash fiery chaos to burn your enemies`},
	1777: {ID: 1777, Key: "Skillld227", Value: `life for you and your party
summon a spirit pet that increases`},
	1778: {ID: 1778, Key: "7ja", Value: `Hyperion Javelin`},
	1779: {ID: 1779, Key: "Diamond", Value: `Diamond`},
	1780: {ID: 1780, Key: "Skillld228", Value: `to fight by your side
summon a wolf with teleporting ability`},
	1781: {ID: 1781, Key: "Skillld229", Value: `transform into a werebear`},
	1782: {ID: 1782, Key: "Skillld230", Value: `that knocks back your enemies
launch a boulder of flaming hot magma`},
	1783: {ID: 1783, Key: "Skillld231", Value: `to burn your enemies with frost
blast a continuous jet of ice`},
	1784: {ID: 1784, Key: "A5Q5SuccessfulCain", Value: `67
You have proven yourself to these 
people. They look to you as their 
warrior, their champion.
`},
	1785: {ID: 1785, Key: "7m7", Value: `Ogre Maul`},
	1786: {ID: 1786, Key: "Skillld232", Value: `and replenishes your life
summon a vine that eats corpses`},
	1787: {ID: 1787, Key: "of Blessed Hammers", Value: `of Blessed Hammer`},
	1788: {ID: 1788, Key: "Hwanin's Justice", Value: `Hwanin's Justice`},
	1789: {ID: 1789, Key: "Skillld233", Value: `with successive hits
increasing amounts of life from your enemies
go into a frenzied rage to steal
when in werewolf form,`},
	1790: {ID: 1790, Key: "of Holy Shield", Value: `of Holy Shield`},
	1791: {ID: 1791, Key: "Skillld234", Value: `with successive hits
for increasing extra damage
maul your enemies
when in werebear form,`},
	1792: {ID: 1792, Key: "Skillld235", Value: `burning them to a crisp
open volcanic vents below your enemies,`},
	1793: {ID: 1793, Key: "Skillld236", Value: `fire, cold, and lightning
shield yourself from damage caused by`},
	1794: {ID: 1794, Key: "Skillld237", Value: `of you and your party
to the damage and attack rating
summon a spirit pet that adds`},
	1795: {ID: 1795, Key: "Skillld238", Value: `it does to enemies
eating corpses to increase damage
summon a wolf that becomes enraged,`},
	1796: {ID: 1796, Key: "Skillld239", Value: `that spreads to other monsters
to inflict them with disease
bite your enemies
when in werewolf form,`},
	1797: {ID: 1797, Key: "EnteringNihlathakAct5Bar", Value: `A coward's hiding place.`},
	1798: {ID: 1798, Key: "Darksight Helm", Value: `Darksight Helm`},
	1799: {ID: 1799, Key: "Skillld240", Value: `with a fiery claw attack
form, maul your enemies
when in werewolf or werebear`},
	1800: {ID: 1800, Key: "Wisp", Value: `Wisp Projector`},
	1801: {ID: 1801, Key: "Smiting", Value: `Smiting`},
	1802: {ID: 1802, Key: "Skillld241", Value: `cut a path through your enemies
release several small whirlwinds that`},
	1803: {ID: 1803, Key: "Skillld242", Value: `and replenishes your mana
summon a vine that eats corpses`},
	1804: {ID: 1804, Key: "A5Q4SuccessfulLarzuk", Value: `93
Hmmm...What does Baal plan to do 
inside Mount Arreat?
`},
	1805: {ID: 1805, Key: "Toxic", Value: `Toxic`},
	1806: {ID: 1806, Key: "FindingTristramDru", Value: `Ruins... the fate of all cities.`},
	1807: {ID: 1807, Key: "Skillld243", Value: `to gain life and mana
form, bite your enemies
when in werewolf or werebear`},
	1808: {ID: 1808, Key: "Skillld244", Value: `that stuns nearby enemies
stomp to create a shock wave
when in werebear form,`},
	1809: {ID: 1809, Key: "of Charged Strike", Value: `of Charged Strike`},
	1810: {ID: 1810, Key: "7la", Value: `Feral Axe`},
	1811: {ID: 1811, Key: "7kr", Value: `Fanged Knife`},
	1812: {ID: 1812, Key: "r02", Value: `Eld Rune`},
	1813: {ID: 1813, Key: "r03", Value: `Tir Rune`},
	1814: {ID: 1814, Key: "r04", Value: `Nef Rune`},
	1815: {ID: 1815, Key: "r05", Value: `Eth Rune`},
	1816: {ID: 1816, Key: "7o7", Value: `Ogre Axe`},
	1817: {ID: 1817, Key: "r07", Value: `Tal Rune`},
	1818: {ID: 1818, Key: "r08", Value: `Ral Rune`},
	1819: {ID: 1819, Key: "r09", Value: `Ort Rune`},
	1820: {ID: 1820, Key: "clw", Value: `Claws`},
	1821: {ID: 1821, Key: "AnyaAct5IntroGossip1", Value: `54
You have proven yourself a true hero to 
me and my people.
 
These are dark times, warrior. I hope 
you can bring an end to Baal's reign of 
destruction. 
 
Our Council of Elders is gone -- my 
father, Aust, among them. The one 
thing that keeps us from total despair 
is the promise of vengeance against 
Baal.
`},
	1822: {ID: 1822, Key: "r06", Value: `Ith Rune`},
	1823: {ID: 1823, Key: "r01", Value: `El Rune`},
	1824: {ID: 1824, Key: "Furious", Value: `Furious`},
	1825: {ID: 1825, Key: "Argent", Value: `Argent`},
	1826: {ID: 1826, Key: "7ma", Value: `Reinforced Mace`},
	1827: {ID: 1827, Key: "r11", Value: `Amn Rune`},
	1828: {ID: 1828, Key: "7ls", Value: `Cryptic Sword`},
	1829: {ID: 1829, Key: "r13", Value: `Shae Rune`},
	1830: {ID: 1830, Key: "r14", Value: `Dol Rune`},
	1831: {ID: 1831, Key: "r15", Value: `Hel Rune`},
	1832: {ID: 1832, Key: "7lw", Value: `Feral Claws`},
	1833: {ID: 1833, Key: "7p7", Value: `War Pike`},
	1834: {ID: 1834, Key: "r18", Value: `Ko Rune`},
	1835: {ID: 1835, Key: "r19", Value: `Fal Rune`},
	1836: {ID: 1836, Key: "r17", Value: `Lum Rune`},
	1837: {ID: 1837, Key: "r16", Value: `Po Rune`},
	1838: {ID: 1838, Key: "r12", Value: `Sol Rune`},
	1839: {ID: 1839, Key: "r10", Value: `Thul Rune`},
	1840: {ID: 1840, Key: "Preserver's", Value: `Preserver's`},
	1841: {ID: 1841, Key: "7mp", Value: `War Spike`},
	1842: {ID: 1842, Key: "LarzukGossip10", Value: `56
Just so you know...The gold you pay me 
doesn't line my pockets. Much of it 
goes to buy the raw metal I need to 
melt down for weapons and armor. The 
rest -- well, all I can spare -- goes to 
Malah and Qual-Kehk for other 
supplies.
`},
	1843: {ID: 1843, Key: "ob1", Value: `Eagle Orb`},
	1844: {ID: 1844, Key: "ob2", Value: `Sacred Globe`},
	1845: {ID: 1845, Key: "7mt", Value: `Devil Star`},
	1846: {ID: 1846, Key: "ob3", Value: `Smoked Sphere`},
	1847: {ID: 1847, Key: "ob4", Value: `Clasped Orb`},
	1848: {ID: 1848, Key: "ob6", Value: `Glowing Orb`},
	1849: {ID: 1849, Key: "ob7", Value: `Crystalline Globe`},
	1850: {ID: 1850, Key: "ob8", Value: `Cloudy Sphere`},
	1851: {ID: 1851, Key: "ob9", Value: `Sparkling Ball`},
	1852: {ID: 1852, Key: "DruOnly", Value: `(Druid Only)`},
	1853: {ID: 1853, Key: "r29", Value: `Sur Rune`},
	1854: {ID: 1854, Key: "r28", Value: `Lo Rune`},
	1855: {ID: 1855, Key: "r27", Value: `Ohm Rune`},
	1856: {ID: 1856, Key: "r26", Value: `Vex Rune`},
	1857: {ID: 1857, Key: "r25", Value: `Gul Rune`},
	1858: {ID: 1858, Key: "r30", Value: `Ber Rune`},
	1859: {ID: 1859, Key: "r31", Value: `Jo Rune`},
	1860: {ID: 1860, Key: "r32", Value: `Cham Rune`},
	1861: {ID: 1861, Key: "r33", Value: `Zod Rune`},
	1862: {ID: 1862, Key: "r24", Value: `Ist Rune`},
	1863: {ID: 1863, Key: "r23", Value: `Mal Rune`},
	1864: {ID: 1864, Key: "r22", Value: `Um Rune`},
	1865: {ID: 1865, Key: "r21", Value: `Pul Rune`},
	1866: {ID: 1866, Key: "r20", Value: `Lem Rune`},
	1867: {ID: 1867, Key: "Carbuncle", Value: `Carbuncle`},
	1868: {ID: 1868, Key: "Scorched", Value: `Scorched`},
	1869: {ID: 1869, Key: "of Insulation", Value: `of Insulation`},
	1870: {ID: 1870, Key: "of the Kraken1", Value: `of the Kraken`},
	1871: {ID: 1871, Key: "of the Whale", Value: `of the Whale`},
	1872: {ID: 1872, Key: "Of Disease", Value: `of Disease`},
	1873: {ID: 1873, Key: "Of the Lady", Value: `of the Lady`},
	1874: {ID: 1874, Key: "LarzukGossip1", Value: `65
I've heard that you Amazons excel at 
killing from a distance. From what I've 
seen, that's the best way to deal with 
Hell's minions. 
 
Are all of your kind so...big?
`},
	1875: {ID: 1875, Key: "LarzukGossip2", Value: `65
Why did you ever leave your beautiful 
islands to come to this frozen 
battleground? Perhaps if we both 
survive this, we'll travel back there 
together.
`},
	1876: {ID: 1876, Key: "LarzukGossip3", Value: `75
Has that infernal howling been keeping 
you awake at night, too? Some think 
it's the howling of animals, but those 
sounds don't come from any beast I've 
ever known.
`},
	1877: {ID: 1877, Key: "LarzukGossip4", Value: `69
Demonic forces have been using our 
own towers and barricades against us. 
You know, it would be wise to cut the 
demons down in the open before they 
can mount those fortifications.
`},
	1878: {ID: 1878, Key: "LarzukGossip5", Value: `70
Nihlathak's despair is infectious -- and 
his behavior does not befit an Elder of 
his stature. I think we'd be better off 
without him.
`},
	1879: {ID: 1879, Key: "LarzukGossip6", Value: `63
I've offered Qual-Kehk my ideas on how 
to break the siege, but he dismisses 
them. Is it because I lack scars of 
battle, or does he think I'm a couple 
arrows short of a quiver?
`},
	1880: {ID: 1880, Key: "LarzukGossip7", Value: `64
Legend has it that the top of Mount 
Arreat is guarded by the spirits of our 
ancestors. Though our people are 
forbidden to climb to the mountain's 
summit, some foreign travelers have 
made the attempt.
 
None have ever returned.
`},
	1881: {ID: 1881, Key: "LarzukGossip8", Value: `66
This is a lively town during our victory 
celebrations. We Barbarians train long 
and hard from childhood to become 
warriors, and we celebrate with equal 
fervor.
 
You didn't happen to bring along any 
ale to trade?
`},
	1882: {ID: 1882, Key: "LarzukGossip9", Value: `54
Every day, one of my friends dies 
fighting outside the town walls, while I 
tend my anvil here unscathed. If only 
we didn't need weapons so badly, I 
could be out doing my share of the 
fighting. 
 
Good luck to you, warrior.
`},
	1883: {ID: 1883, Key: "7pi", Value: `Stygian Pilum`},
	1884: {ID: 1884, Key: "7s7", Value: `Balrog Spear`},
	1885: {ID: 1885, Key: "7s8", Value: `Thresher`},
	1886: {ID: 1886, Key: "7pa", Value: `Cryptic Axe`},
	1887: {ID: 1887, Key: "Nickel", Value: `Nickel`},
	1888: {ID: 1888, Key: "Perpetual", Value: `Perpetual`},
	1889: {ID: 1889, Key: "A5Q1AfterInitAnya", Value: `84
You've proven your skill at rescuing fair 
maidens...but how are you at killing 
vicious beasts?
`},
	1890: {ID: 1890, Key: "Of the Witch", Value: `of the Witch`},
	1891: {ID: 1891, Key: "oba", Value: `Swirling Crystal`},
	1892: {ID: 1892, Key: "obb", Value: `Heavenly Stone`},
	1893: {ID: 1893, Key: "obc", Value: `Eldritch Orb`},
	1894: {ID: 1894, Key: "obd", Value: `Demon Heart`},
	1895: {ID: 1895, Key: "obe", Value: `Vortex Orb`},
	1896: {ID: 1896, Key: "obf", Value: `Dimensional Shard`},
	1897: {ID: 1897, Key: "Hierophant's", Value: `Hierophant's`},
	1898: {ID: 1898, Key: "Acrobat's", Value: `Acrobatic`},
	1899: {ID: 1899, Key: "Vermillion", Value: `Vermilion`},
	1900: {ID: 1900, Key: "Of Abrasion", Value: `of Abrasion`},
	1901: {ID: 1901, Key: "Of Doom", Value: `of Doom`},
	1902: {ID: 1902, Key: "of Corpse Explosions", Value: `of Corpse Explosion`},
	1903: {ID: 1903, Key: "GemeffectX51", Value: `Adds to Attack Rating`},
	1904: {ID: 1904, Key: "GemeffectX52", Value: `Adds to All Resistances`},
	1905: {ID: 1905, Key: "A5Q2AfterInitAnya", Value: `85
If those men are being treated like I 
was, you must find them. They won't 
survive much longer.
`},
	1906: {ID: 1906, Key: "of the Cobra", Value: `of the Cobra`},
	1907: {ID: 1907, Key: "7qr", Value: `Scissors Suwayyah`},
	1908: {ID: 1908, Key: "7qs", Value: `Seraph Rod`},
	1909: {ID: 1909, Key: "GemeffectX53", Value: `Adds to Damage vs. Undead`},
	1910: {ID: 1910, Key: "GemeffectX61", Value: `Adds to Chance to Find Magic Items`},
	1911: {ID: 1911, Key: "GemeffectX62", Value: `Adds Resistance to Lightning`},
	1912: {ID: 1912, Key: "GemeffectX63", Value: `Adds Lightning Damage to Attack`},
	1913: {ID: 1913, Key: "The Fetid Sprinkler", Value: `The Fetid Sprinkler`},
	1914: {ID: 1914, Key: "GemeffectX71", Value: `Adds Mana and Life Regeneration`},
	1915: {ID: 1915, Key: "GemeffectX72", Value: `Adds Attacker Takes Damage`},
	1916: {ID: 1916, Key: "of Telekinesis", Value: `of Telekinesis`},
	1917: {ID: 1917, Key: "GemeffectX73", Value: `Adds Mana and Life Steal to Attack`},
	1918: {ID: 1918, Key: "Blackhorn", Value: `Blackhorn`},
	1919: {ID: 1919, Key: "Gaia's Wrath", Value: `Gaia's Wrath`},
	1920: {ID: 1920, Key: "Fleshbone", Value: `Fleshbone`},
	1921: {ID: 1921, Key: "Executioner's Justice", Value: `Executioner's Justice`},
	1922: {ID: 1922, Key: "Gargoyle's Bite", Value: `Gargoyle's Bite`},
	1923: {ID: 1923, Key: "7sb", Value: `Elegant Blade`},
	1924: {ID: 1924, Key: "7sc", Value: `Mighty Scepter`},
	1925: {ID: 1925, Key: "Djinnslayer", Value: `Djinn Slayer`},
	1926: {ID: 1926, Key: "Tiamat's Rebuke", Value: `Tiamat's Rebuke`},
	1927: {ID: 1927, Key: "Warriv's Warder", Value: `Warriv's Warder`},
	1928: {ID: 1928, Key: "Widow maker", Value: `Widowmaker`},
	1929: {ID: 1929, Key: "Blanched", Value: `Blanched`},
	1930: {ID: 1930, Key: "Firelizard's Talons", Value: `Firelizard's Talons`},
	1931: {ID: 1931, Key: "Saracen's Chance", Value: `Saracen's Chance`},
	1932: {ID: 1932, Key: "Mara's Kaleidoscope", Value: `Mara's Kaleidoscope`},
	1933: {ID: 1933, Key: "The Rising Sun", Value: `The Rising Sun`},
	1934: {ID: 1934, Key: "7sm", Value: `Ataghan`},
	1935: {ID: 1935, Key: "Of Remorse", Value: `of Remorse`},
	1936: {ID: 1936, Key: "Stormshield", Value: `Stormshield`},
	1937: {ID: 1937, Key: "A5Q4AfterInitAnya", Value: `90
As noble as Nihlathak's intentions are, 
nothing can excuse his actions.
`},
	1938: {ID: 1938, Key: "7sp", Value: `Tyrant Club`},
	1939: {ID: 1939, Key: "A5Q5SuccessfulMalah", Value: `60
I knew the Ancients would find you 
worthy of Mount Arreat's secrets. Now, 
stop Baal before he destroys all that is 
sacred.
`},
	1940: {ID: 1940, Key: "7ss", Value: `Falcata`},
	1941: {ID: 1941, Key: "7ta", Value: `Flying Axe`},
	1942: {ID: 1942, Key: "7sr", Value: `Hyperion Spear`},
	1943: {ID: 1943, Key: "7st", Value: `Ghost Spear`},
	1944: {ID: 1944, Key: "CainAct5Gossip1", Value: `52
With hellspawn, size is no measure of 
their threat. Demons half the size of 
men can kill with a gesture, while 
hellish pack animals trample any who 
stand in their way.
`},
	1945: {ID: 1945, Key: "CainAct5Gossip2", Value: `55
Though these Barbarians are known 
throughout the kingdoms as ferocious 
fighters, they are also capable of great 
compassion.
 
They have trained throughout history 
for a battle their legends foretell will 
decide the fate of the world.
`},
	1946: {ID: 1946, Key: "CainAct5Gossip3", Value: `55
The angel Tyrael has watched over the 
guardians of Arreat throughout 
history. I do not believe that Baal and 
Tyrael have come to fight over a paltry 
few souls.
 
They are here to settle a conflict as old 
as time itself.
`},
	1947: {ID: 1947, Key: "CainAct5Gossip4", Value: `44
During my time with the Horadrim, we 
often debated the nature of Mount 
Arreat.
 
We knew that the Barbarian clans 
zealously guarded the mountain as 
their sacred duty. However, many 
dismissed their zeal as simple 
superstition...combined with an inborn 
hostility toward outsiders.
 
Those Horadrim who trekked up Arreat 
were never heard from again...Still, I 
do not believe they died at the hands of 
Barbarians.
`},
	1948: {ID: 1948, Key: "CainAct5Gossip5", Value: `41
All users of the magical arts know of 
Mount Arreat, but few understand its 
true nature. It is the nexus of an 
unfathomable magic.
 
It bodes ill that the Lord of Destruction 
races to its summit with such purpose. 
I fear for the whole world should Baal 
gain what he seeks.
`},
	1949: {ID: 1949, Key: "CainAct5Gossip6", Value: `59
I have spent decades trying to 
understand the forces at work in this 
world. But in the face of all that is 
transpiring, I realize how meager my 
knowledge is.
 
I will be of assistance where I can, my 
friend.
`},
	1950: {ID: 1950, Key: "CainAct5Gossip7", Value: `54
Though the Elder Council of Harrogath 
is gone, there are many capable young 
leaders to take their place.
 
Anya certainly has enough courage and 
intelligence to lead them all, if they can 
survive this catastrophe.
`},
	1951: {ID: 1951, Key: "CainAct5Gossip8", Value: `53
Ah, Anya. Such a fine example of 
feminine strength...
 
She reminds me of the Zakarum 
priestesses I knew in my youth. They 
don't take vows of chastity, you know.
`},
	1952: {ID: 1952, Key: "CainAct5Gossip9", Value: `65
It is fortunate that this town has such a 
talented smith.
 
The quality of Larzuk's work surely 
complements your skills. In fact, he 
would have been quite welcome 
amongst the Horadrim.
`},
	1953: {ID: 1953, Key: "A5Q5AfterInitAnya", Value: `72
You wouldn't have to test yourself 
against the Ancients if it weren't for 
Nihlathak's treachery. He was a fool 
and deserves to suffer for eternity.
`},
	1954: {ID: 1954, Key: "7tk", Value: `Flying Knife`},
	1955: {ID: 1955, Key: "7tr", Value: `Stygian Pike`},
	1956: {ID: 1956, Key: "7ts", Value: `Winged Harpoon`},
	1957: {ID: 1957, Key: "Oligocene", Value: `Honorable`},
	1958: {ID: 1958, Key: "Assamic", Value: `Lunar`},
	1959: {ID: 1959, Key: "AssOnly", Value: `(Assassin Only)`},
	1960: {ID: 1960, Key: "7tw", Value: `Runic Talons`},
	1961: {ID: 1961, Key: "Of Redemption", Value: `of Redemption`},
	1962: {ID: 1962, Key: "Of Judgement", Value: `of Judgement`},
	1963: {ID: 1963, Key: "Magekiller's", Value: `Magekiller's`},
	1964: {ID: 1964, Key: "Of the Plague", Value: `of the Plague`},
	1965: {ID: 1965, Key: "Of Horror", Value: `of Horror`},
	1966: {ID: 1966, Key: "Of the Forest", Value: `of the Forest`},
	1967: {ID: 1967, Key: "of Bone Imprisonment", Value: `of Bone Prison`},
	1968: {ID: 1968, Key: "of Novas", Value: `of Nova`},
	1969: {ID: 1969, Key: "GemeffectX11", Value: `Adds to Strength`},
	1970: {ID: 1970, Key: "GemeffectX12", Value: `Adds to Defense Rating`},
	1971: {ID: 1971, Key: "GemeffectX13", Value: `Adds to Attack Rating`},
	1972: {ID: 1972, Key: "GemeffectX21", Value: `Adds to Maximum Mana`},
	1973: {ID: 1973, Key: "GemeffectX22", Value: `Adds Resistance to Cold`},
	1974: {ID: 1974, Key: "GemeffectX23", Value: `Adds Cold Damage to Attack`},
	1975: {ID: 1975, Key: "Frostwind", Value: `Frostwind`},
	1976: {ID: 1976, Key: "Marrowwalk", Value: `Marrowwalk`},
	1977: {ID: 1977, Key: "The Cat's Eye", Value: `The Cat's Eye`},
	1978: {ID: 1978, Key: "GemeffectX31", Value: `Adds to Dexterity`},
	1979: {ID: 1979, Key: "GemeffectX32", Value: `Adds Resistance to Poison`},
	1980: {ID: 1980, Key: "GemeffectX33", Value: `Adds Poison Damage to Attack`},
	1981: {ID: 1981, Key: "Veil of Steel", Value: `Veil of Steel`},
	1982: {ID: 1982, Key: "The Grandfather", Value: `The Grandfather`},
	1983: {ID: 1983, Key: "Hellslayer", Value: `Hellslayer`},
	1984: {ID: 1984, Key: "7vo", Value: `Colossus Voulge`},
	1985: {ID: 1985, Key: "Sigurd's Staunch", Value: `Siggard's Stealth`},
	1986: {ID: 1986, Key: "7wa", Value: `Beserker Axe`},
	1987: {ID: 1987, Key: "7wb", Value: `Wrist Sword`},
	1988: {ID: 1988, Key: "7wc", Value: `Giant Thresher`},
	1989: {ID: 1989, Key: "7wd", Value: `Mythical Sword`},
	1990: {ID: 1990, Key: "LarzukAct5IntroGossip1", Value: `55
I am Larzuk, the armorer. My ancestors 
were some of the finest craftsmen in 
Harrogath. 
 
Regretfully, my supplies run lower with 
every passing day, yet the demons 
beyond the walls have not weakened. 
 
I fear the time is near when I must put 
down my hammer and take up a sword, 
instead.
`},
	1991: {ID: 1991, Key: "of Conversion", Value: `of Conversion`},
	1992: {ID: 1992, Key: "AncientsAct5IntroGossip1", Value: `35
We are the spirits of the Nephalem, the 
Ancient Ones. We have been chosen to 
guard sacred Mount Arreat, wherein 
the Worldstone rests. Few are worthy 
to stand in its presence; fewer still can 
comprehend its true purpose.
 
Before you enter, you must defeat us.
`},
	1993: {ID: 1993, Key: "7wh", Value: `Legendary Mallet`},
	1994: {ID: 1994, Key: "of Slow Missiles", Value: `of Slow Missile`},
	1995: {ID: 1995, Key: "GemeffectX41", Value: `Adds to Maximum Life`},
	1996: {ID: 1996, Key: "Colorful", Value: `Colorful`},
	1997: {ID: 1997, Key: "GemeffectX42", Value: `Adds Resistance to Fire`},
	1998: {ID: 1998, Key: "GemeffectX43", Value: `Adds Fire Damage to Attack`},
	1999: {ID: 1999, Key: "7wn", Value: `Polished Wand`},
	2000: {ID: 2000, Key: "Bling Bling", Value: `Bling Bling`},
	2001: {ID: 2001, Key: "Razortail", Value: `Razortail`},
	2002: {ID: 2002, Key: "Griswold's Heart", Value: `Griswold's Heart`},
	2003: {ID: 2003, Key: "Flowkrad's Fur", Value: `Flowkrad's Fur`},
	2004: {ID: 2004, Key: "7ws", Value: `Caduceus`},
	2005: {ID: 2005, Key: "Verdugo's Hearty Cord", Value: `Verdungo's Hearty Cord`},
	2006: {ID: 2006, Key: "Aldur's Watchtower", Value: `Aldur's Watchtower`},
	2007: {ID: 2007, Key: "7xf", Value: `War Fist`},
	2008: {ID: 2008, Key: "of Virility", Value: `of Virility`},
	2009: {ID: 2009, Key: "M'avina's Tenet", Value: `M'avina's Tenet`},
	2010: {ID: 2010, Key: "Ghoulhide", Value: `Ghoulhide`},
	2011: {ID: 2011, Key: "Hwanin's Majesty", Value: `Hwanin's Majesty`},
	2012: {ID: 2012, Key: "Cow King's Leathers", Value: `Cow King's Leathers`},
	2013: {ID: 2013, Key: "Ferocious", Value: `Ferocious`},
	2014: {ID: 2014, Key: "Aragon's Sunfire", Value: `Aragon's Sunfire`},
	2015: {ID: 2015, Key: "Apocrypha", Value: `Apocrypha`},
	2016: {ID: 2016, Key: "Tomb Reaver", Value: `Tomb Reaver`},
	2017: {ID: 2017, Key: "Aurora's Guard", Value: `Aurora's Guard`},
	2018: {ID: 2018, Key: "Globe of Visjerei", Value: `Globe of the Vizjerei`},
	2019: {ID: 2019, Key: "EskillPassiveFeral", Value: `Passive Bonus to Wolves and Bears`},
	2020: {ID: 2020, Key: "Of the Mosquito", Value: `of the Mosquito`},
	2021: {ID: 2021, Key: "Eskillperhit12", Value: ` per hit`},
	2022: {ID: 2022, Key: "Ogun's Vengeance", Value: `Ogun's Vengeance`},
	2023: {ID: 2023, Key: "ob5", Value: `Jared's Stone`},
	2024: {ID: 2024, Key: "fana", Value: `Frozen Anya`},
	2025: {ID: 2025, Key: "Charged", Value: `Charged`},
	2026: {ID: 2026, Key: "of Freezing Arrows", Value: `of Freezing Arrow`},
	2027: {ID: 2027, Key: "Aragon's Masterpiece", Value: `Aragon's Masterpiece`},
	2028: {ID: 2028, Key: "A5Q5InitQualKehk", Value: `37
Every time I hear of you, warrior, your 
deeds become more legendary.
 
But take heed. You are approaching the 
very summit of Mount Arreat. I have 
never dared venture there. It is sacred 
-- our most holy place. The legends say 
it is guarded by the Ancient Ones, who 
block the path of all who are unworthy.
 
Your reputation here does not 
matter...It will be the Ancients who 
determine your worthiness.
 
Good luck.
`},
	2029: {ID: 2029, Key: "qstsa5q61a", Value: `Find Baal's Throne Room.`},
	2030: {ID: 2030, Key: "of Plague Jab", Value: `of Plague Jab`},
	2031: {ID: 2031, Key: "Heaven's Taebaek", Value: `Taebaek's Glory`},
	2032: {ID: 2032, Key: "Immortal King's Pillar", Value: `Immortal King's Pillar`},
	2033: {ID: 2033, Key: "of the Plague", Value: `of the Plague`},
	2034: {ID: 2034, Key: "EnterDOEDru", Value: `So, this is where evil hides.`},
	2035: {ID: 2035, Key: "EnteringClawViperAss", Value: `Dark magic in a darker tomb...`},
	2036: {ID: 2036, Key: "CompletingTaintedSunDru", Value: `The sun warms the world once more.`},
	2037: {ID: 2037, Key: "CompletingSummonerDru", Value: `Now I can leave this twisted nightmare.`},
	2038: {ID: 2038, Key: "CompletingLamEsenDru", Value: `Ormus... study the book well.`},
	2039: {ID: 2039, Key: "CompletingBeneathCityAss", Value: `This is one drain I don't mind cleaning out.`},
	2040: {ID: 2040, Key: "7yw", Value: `Ghost Wand`},
	2041: {ID: 2041, Key: "Of Admiration", Value: `of Admiration`},
	2042: {ID: 2042, Key: "Samual's Caretaker", Value: `Samuel's Caretaker`},
	2043: {ID: 2043, Key: "Aragon's Icy Stare", Value: `Aragon's Icy Stare`},
	2044: {ID: 2044, Key: "FindingJadeFigDru", Value: `It looks like jade. Perhaps it's worth something.`},
	2045: {ID: 2045, Key: "of the Icicle", Value: `of the Icicle`},
	2046: {ID: 2046, Key: "Windhammer", Value: `Windhammer`},
	2047: {ID: 2047, Key: "CompletingGuardianTowerAss", Value: `Mephisto, you were no match for me.`},
	2048: {ID: 2048, Key: "CompletingGuardianTowerAss", Value: `Mephisto's hatred was a poisonous void.`},
	2049: {ID: 2049, Key: "KillingdDiabloDru", Value: `Thus ends the plague of Terror.`},
	2050: {ID: 2050, Key: "CompletingStopSiegeBar", Value: `The siege is broken.`},
	2051: {ID: 2051, Key: "EnteringNihlathakAct5Ass", Value: `I should have known...`},
	2052: {ID: 2052, Key: "Raiden's Crutch", Value: `Raiden's Crutch`},
	2053: {ID: 2053, Key: "Hexfire", Value: `Hexfire`},
	2054: {ID: 2054, Key: "Boneslayer Blade", Value: `Boneslayer Blade`},
	2055: {ID: 2055, Key: "Guardian Angel", Value: `Guardian Angel`},
	2056: {ID: 2056, Key: "EnteringWorldstoneAct5Bar", Value: `The halls of the Ancients... Magnificent.`},
	2057: {ID: 2057, Key: "Blackleach Blade", Value: `Blackleach Blade`},
	2058: {ID: 2058, Key: "Bul-Kathos' Custodian", Value: `Bul-Kathos' Custodian`},
	2059: {ID: 2059, Key: "Griswold's Honor", Value: `Griswold's Honor`},
	2060: {ID: 2060, Key: "M'avina's Battle Hymn", Value: `M'avina's Battle Hymn`},
	2061: {ID: 2061, Key: "EnteringWorldstoneAct5Ams", Value: `The Worldstone!`},
	2062: {ID: 2062, Key: "Dracul's Grasp", Value: `Dracul's Grasp`},
	2063: {ID: 2063, Key: "of Terror", Value: `of Terror`},
	2064: {ID: 2064, Key: "EnteringWorldstoneAct5Ass", Value: `The Worldstone. What power.`},
	2065: {ID: 2065, Key: "of Incombustibility", Value: `of Inflammability`},
	2066: {ID: 2066, Key: "of Butchery", Value: `of Butchery`},
	2067: {ID: 2067, Key: "of the Colossus", Value: `of the Gargantuan`},
	2068: {ID: 2068, Key: "CompletingStopSiegeAss", Value: `Shenk, your command has ended.`},
	2069: {ID: 2069, Key: "of Transcendence", Value: `of Transcendence`},
	2070: {ID: 2070, Key: "Skillld245", Value: `and destruction over your enemies
summon forth a volcano to rain death`},
	2071: {ID: 2071, Key: "CompletingDefeatBaalAct5Pal", Value: `Baal, you shall no longer taint this mortal realm.`},
	2072: {ID: 2072, Key: "Skillld246", Value: `to blast your enemies
create a funnel of wind and debris`},
	2073: {ID: 2073, Key: "Aldur's Advance", Value: `Aldur's Advance`},
	2074: {ID: 2074, Key: "A5Q4AfterInitQualKehk", Value: `60
I saw Nihlathak leave town just before 
you found Anya. He must be held 
accountable for his criminal deeds.
 
Find him and bring him back, if you can. 
Likely, he won't come willingly, and 
you'll be forced to kill him.
 
So be it.
`},
	2075: {ID: 2075, Key: "Arachnid Mesh", Value: `Arachnid Mesh`},
	2076: {ID: 2076, Key: "CfgSay7", Value: `Say 'Retreat'`},
	2077: {ID: 2077, Key: "Skillld247", Value: `back at your enemies
taken by you and your party
summon spirit pet that reflects damage`},
	2078: {ID: 2078, Key: "ModStre8p", Value: ` to Curses`},
	2079: {ID: 2079, Key: "ModStre8q", Value: ` to Warcry Skills`},
	2080: {ID: 2080, Key: "ModStre8r", Value: ` to Combat Skills`},
	2081: {ID: 2081, Key: "ModStre8s", Value: ` to Masteries Skills`},
	2082: {ID: 2082, Key: "ModStre8t", Value: ` to Cold Skills`},
	2083: {ID: 2083, Key: "pa1", Value: `Targe`},
	2084: {ID: 2084, Key: "pa2", Value: `Rondache`},
	2085: {ID: 2085, Key: "pa3", Value: `Heraldic Shield`},
	2086: {ID: 2086, Key: "pa4", Value: `Aerin Shield`},
	2087: {ID: 2087, Key: "pa5", Value: `Crown Shield`},
	2088: {ID: 2088, Key: "pa6", Value: `Akaran Targe`},
	2089: {ID: 2089, Key: "pa7", Value: `Akaran Rondache`},
	2090: {ID: 2090, Key: "pa8", Value: `Protector Shield`},
	2091: {ID: 2091, Key: "pa9", Value: `Gilded Shield`},
	2092: {ID: 2092, Key: "of Nova Shield", Value: `of Nova Shield`},
	2093: {ID: 2093, Key: "of Raise Skeletal Mages", Value: `of Skeletal Mages`},
	2094: {ID: 2094, Key: "ModStre8u", Value: ` to Lightning Skills`},
	2095: {ID: 2095, Key: "ModStre8a", Value: `to Druid Skills`},
	2096: {ID: 2096, Key: "ModStre8b", Value: `to Assassin Skills`},
	2097: {ID: 2097, Key: "ModStre8c", Value: ` Sockets`},
	2098: {ID: 2098, Key: "ModStre8d", Value: ` to Attack Rating vs. Demons`},
	2099: {ID: 2099, Key: "ModStre8e", Value: ` to Attack Rating vs. Undead`},
	2100: {ID: 2100, Key: "ModStre8f", Value: ` to Damage vs. Demons`},
	2101: {ID: 2101, Key: "NecOnly", Value: `(Necromancer Only)`},
	2102: {ID: 2102, Key: "dr1", Value: `Wolf Head`},
	2103: {ID: 2103, Key: "dr2", Value: `Hawk Helm`},
	2104: {ID: 2104, Key: "dr3", Value: `Antlers`},
	2105: {ID: 2105, Key: "dr4", Value: `Falcon Mask`},
	2106: {ID: 2106, Key: "dr5", Value: `Spirit Mask`},
	2107: {ID: 2107, Key: "dr6", Value: `Alpha Helm`},
	2108: {ID: 2108, Key: "dr7", Value: `Griffon Headdress`},
	2109: {ID: 2109, Key: "dr8", Value: `Hunter's Guise`},
	2110: {ID: 2110, Key: "dr9", Value: `Sacred Feathers`},
	2111: {ID: 2111, Key: "ModStre8g", Value: ` percent to Attack Rating`},
	2112: {ID: 2112, Key: "ModStre8h", Value: ` to Javelin and Spear Skills`},
	2113: {ID: 2113, Key: "ModStre8i", Value: ` to Passive and Magic Skills`},
	2114: {ID: 2114, Key: "ModStre8j", Value: ` to Bow and Crossbow Skills`},
	2115: {ID: 2115, Key: "Of the Fly", Value: `of the Fly`},
	2116: {ID: 2116, Key: "Gritty", Value: `Grinding`},
	2117: {ID: 2117, Key: "of Avarice", Value: `of Avarice`},
	2118: {ID: 2118, Key: "ModStre8k", Value: ` to Defensive Aura Skills`},
	2119: {ID: 2119, Key: "ModStre8l", Value: ` to Offensive Aura Skills`},
	2120: {ID: 2120, Key: "ModStre8m", Value: ` to Combat Skills`},
	2121: {ID: 2121, Key: "ModStre8n", Value: ` to Summoning Skills`},
	2122: {ID: 2122, Key: "ModStre8o", Value: ` to Poison and Bone Skills`},
	2123: {ID: 2123, Key: "ModStre8v", Value: ` to Fire Skills`},
	2124: {ID: 2124, Key: "ModStre8w", Value: ` to Summoning Skills`},
	2125: {ID: 2125, Key: "Trump", Value: `Fool's`},
	2126: {ID: 2126, Key: "ModStre8x", Value: ` to Shape-Shifting Skills`},
	2127: {ID: 2127, Key: "of Hurricane", Value: `of Hurricane`},
	2128: {ID: 2128, Key: "ModStre8y", Value: ` to Elemental Skills`},
	2129: {ID: 2129, Key: "ModStre8z", Value: ` to Trap Skills`},
	2130: {ID: 2130, Key: "ModStre9a", Value: ` to Shadow Discipline Skills`},
	2131: {ID: 2131, Key: "paa", Value: `Royal Shield`},
	2132: {ID: 2132, Key: "pab", Value: `Sacred Targe`},
	2133: {ID: 2133, Key: "pac", Value: `Sacred Rondache`},
	2134: {ID: 2134, Key: "Aureolin", Value: `Aureolic`},
	2135: {ID: 2135, Key: "pae", Value: `Zakarum Shield`},
	2136: {ID: 2136, Key: "paf", Value: `Vortex Shield`},
	2137: {ID: 2137, Key: "ModStre9b", Value: ` to Martial Art Skills`},
	2138: {ID: 2138, Key: "ModStre9c", Value: `(Based on Character Level)`},
	2139: {ID: 2139, Key: "of Hydras", Value: `of Hydra`},
	2140: {ID: 2140, Key: "ModStre9d", Value: `(Increases During Nighttime)`},
	2141: {ID: 2141, Key: "Majestic", Value: `Majestic`},
	2142: {ID: 2142, Key: "Of the Earth", Value: `of the Earth`},
	2143: {ID: 2143, Key: "ModStre9e", Value: `(Increases During Daytime)`},
	2144: {ID: 2144, Key: "Robineye", Value: `Robineye`},
	2145: {ID: 2145, Key: "Volcanic", Value: `Volcanic`},
	2146: {ID: 2146, Key: "ModStre9f", Value: `(Increases Near Dawn)`},
	2147: {ID: 2147, Key: "RuneQuote", Value: `'`},
	2148: {ID: 2148, Key: "ModStre9g", Value: `(Increases Near Dusk)`},
	2149: {ID: 2149, Key: "ModStre9h", Value: ` Charges of`},
	2150: {ID: 2150, Key: "dra", Value: `Totemic Mask`},
	2151: {ID: 2151, Key: "drb", Value: `Blood Spirit`},
	2152: {ID: 2152, Key: "drc", Value: `Sun Spirit`},
	2153: {ID: 2153, Key: "drd", Value: `Earth Spirit`},
	2154: {ID: 2154, Key: "dre", Value: `Sky Spirit`},
	2155: {ID: 2155, Key: "A5Q6EarlyReturnQualKehk", Value: `63
You have ventured to a place beyond 
legend. You rush to face an evil few 
can even imagine.
 
Be careful, my friend, and may the Light 
watch over you.
`},
	2156: {ID: 2156, Key: "drf", Value: `Dream Spirit`},
	2157: {ID: 2157, Key: "ModStre9i", Value: `Increased Stack Size`},
	2158: {ID: 2158, Key: "Arm of King Leoric", Value: `Arm of King Leoric`},
	2159: {ID: 2159, Key: "Cloudcrack", Value: `Cloudcrack`},
	2160: {ID: 2160, Key: "Whichwild String", Value: `Whichwild String`},
	2161: {ID: 2161, Key: "Steelrend", Value: `Steelrend`},
	2162: {ID: 2162, Key: "Griffon's Eye", Value: `Griffon's Eye`},
	2163: {ID: 2163, Key: "9ar", Value: `Quhab`},
	2164: {ID: 2164, Key: "Iansang's Frenzy", Value: `Iansang's Frenzy`},
	2165: {ID: 2165, Key: "of Ice Arrows", Value: `of Ice Arrow`},
	2166: {ID: 2166, Key: "A5Q3AfterInitQualKehk", Value: `43
Anya's father, Aust, was our wisest 
Elder. He was killed along with the 
other Elders who erected the ward that 
protects this city. The ward has kept 
Baal's demons out of Harrogath, but at 
a costly sacrifice. 
 
Nihlathak, on the other hand, was the 
only Elder to escape the demons. 
Somehow, he alone managed to find 
sanctuary, while the others died 
around him.
 
Ever since that day, Nihlathak and Anya 
have been at odds.
`},
	2167: {ID: 2167, Key: "Jeweler's", Value: `Jeweler's`},
	2168: {ID: 2168, Key: "Of Darkness", Value: `of Darkness`},
	2169: {ID: 2169, Key: "Coldsteal Eye", Value: `Coldsteal Eye`},
	2170: {ID: 2170, Key: "Amodeus's Manipulator", Value: `Amodeus' Manipulator`},
	2171: {ID: 2171, Key: "Darksoul", Value: `Darksoul`},
	2172: {ID: 2172, Key: "Cerebus", Value: `Cerebus' Bite`},
	2173: {ID: 2173, Key: "Lidless Wall", Value: `Lidless Wall`},
	2174: {ID: 2174, Key: "Steelpillar", Value: `Steel Pillar`},
	2175: {ID: 2175, Key: "Lance Guard", Value: `Lance Guard`},
	2176: {ID: 2176, Key: "A5Q5EarlyReturnLarzuk", Value: `70
You've walked on the burial grounds of 
our greatest ancestors. I'm not sure 
whether I should bow before you or 
revile you for committing sacrilege.
 
Tread lightly when you walk with gods.
`},
	2177: {ID: 2177, Key: "Boreal", Value: `Boreal`},
	2178: {ID: 2178, Key: "Windforce", Value: `Windforce`},
	2179: {ID: 2179, Key: "Natalya's Soul", Value: `Natalya's Soul`},
	2180: {ID: 2180, Key: "Griswold's Legacy", Value: `Griswold's Legacy`},
	2181: {ID: 2181, Key: "M'avina's Embrace", Value: `M'avina's Embrace`},
	2182: {ID: 2182, Key: "The Disciple", Value: `The Disciple`},
	2183: {ID: 2183, Key: "Of Pilfering", Value: `of Pilfering`},
	2184: {ID: 2184, Key: "Laying of Hands", Value: `Laying of Hands`},
	2185: {ID: 2185, Key: "pad", Value: `Kurast Shield`},
	2186: {ID: 2186, Key: "qstsa5q21a", Value: `Free the soldiers from their prison and lead them back to town.`},
	2187: {ID: 2187, Key: "A5Q4EarlyReturnQualKehk", Value: `65
My advice is to go in quick and hit hard. 
Nihlathak can't be half as tough as the 
beasts you've faced out there.
`},
	2188: {ID: 2188, Key: "A5Q2AfterInitMalah", Value: `84
Qual-Kehk's men have been imprisoned 
for some time. They are certain to be 
tired and weak.
 
You must protect them once you free 
them.
`},
	2189: {ID: 2189, Key: "act1X", Value: `THE SISTER'S LAMENT`},
	2190: {ID: 2190, Key: "act2X", Value: `DESERT JOURNEY`},
	2191: {ID: 2191, Key: "act3X", Value: `MEPHISTO'S JUNGLE`},
	2192: {ID: 2192, Key: "act4X", Value: `ENTER HELL`},
	2193: {ID: 2193, Key: "Heated", Value: `Heated`},
	2194: {ID: 2194, Key: "Of the Specter", Value: `of the Specter`},
	2195: {ID: 2195, Key: "Godstrike Arch", Value: `Goldstrike Arch`},
	2196: {ID: 2196, Key: "9cs", Value: `Hand Scythe`},
	2197: {ID: 2197, Key: "Paradox", Value: `Paradox`},
	2198: {ID: 2198, Key: "of Knowledge", Value: `of Knowledge`},
	2199: {ID: 2199, Key: "act5X", Value: `SEARCH FOR BAAL`},
	2200: {ID: 2200, Key: "Tal Rasha's Lidless Eye", Value: `Tal Rasha's Lidless Eye`},
	2201: {ID: 2201, Key: "CompletingTaintedSunAss", Value: `Serpents! I expected worse.`},
	2202: {ID: 2202, Key: "CompletingTombAss", Value: `I shall track the Prime Evils to the ends of the world.`},
	2203: {ID: 2203, Key: "A5Q5EarlyReturnQualKehk", Value: `114
I warned you!
 
The Ancients are not like the demons 
you're accustomed to fighting.
`},
	2204: {ID: 2204, Key: "A5Q3SuccessfulMalah", Value: `48
Thank you so much for bringing Anya 
back to us. I have devised this spell to 
increase your resistances as a token of 
my thanks. I know it isn't much, but I 
hope you find it helpful.
 
Please go talk to Anya. She has urgent 
news concerning Nihlathak.
`},
	2205: {ID: 2205, Key: "Gaudy", Value: `Gaudy`},
	2206: {ID: 2206, Key: "Hawk Branded", Value: `Hawk Branded`},
	2207: {ID: 2207, Key: "Snaketongue", Value: `Snake Tongue`},
	2208: {ID: 2208, Key: "FindingLamEsenAss", Value: `Black books make for dark thoughts.`},
	2209: {ID: 2209, Key: "RescueQual-KehkAct5Sor", Value: `Follow me.`},
	2210: {ID: 2210, Key: "Stormchaser", Value: `Stormchaser`},
	2211: {ID: 2211, Key: "RescueQual-KehkAct5Nec", Value: `Follow me.`},
	2212: {ID: 2212, Key: "CompletingDefeatBaalAct5Bar", Value: `The Prime Evils are no more.`},
	2213: {ID: 2213, Key: "CompletingDefeatBaalAct5Ass", Value: `The Evil brotherhood is no more.`},
	2214: {ID: 2214, Key: "Skillan222", Value: `Raven`},
	2215: {ID: 2215, Key: "Skillan223", Value: `Psn Creep`},
	2216: {ID: 2216, Key: "RescueCainAss", Value: `Cain! Go to the Rogue camp.`},
	2217: {ID: 2217, Key: "Bul Katho's Wedding Band", Value: `Bul-Kathos' Wedding Band`},
	2218: {ID: 2218, Key: "Skillan223", Value: `Werewolf`},
	2219: {ID: 2219, Key: "A5Q2EarlyReturnQualKehk", Value: `95
They say that discretion, not 
procrastination, is the better part of 
valor.
`},
	2220: {ID: 2220, Key: "CompletingAndarielDru", Value: `Your reign is over, Andariel.`},
	2221: {ID: 2221, Key: "Guardian's", Value: `Guardian's`},
	2222: {ID: 2222, Key: "Naj's Circlet", Value: `Naj's Circlet`},
	2223: {ID: 2223, Key: "Of Searing", Value: `of Searing`},
	2224: {ID: 2224, Key: "Dragontooth", Value: `Dragontooth`},
	2225: {ID: 2225, Key: "Merman's Speed", Value: `Merman's Sprocket`},
	2226: {ID: 2226, Key: "FindingTrueTombAss", Value: `I can sense Tal Rasha's presence now.`},
	2227: {ID: 2227, Key: "Skillan225", Value: `Lycanthropy`},
	2228: {ID: 2228, Key: "MalahGossip1", Value: `53
I'm aware of the amazing magical 
powers a Sorceress can channel. If we 
survive Baal's wrath, I would be most 
honored if you could demonstrate and 
perhaps teach me something of what 
you know. 
 
I may be old, but I'm not dead.
`},
	2229: {ID: 2229, Key: "MalahGossip2", Value: `42
I know you and my son, Bannuk, did not 
part on the best of terms. He did not 
understand the wanderlust that drove 
you to leave your homeland. However, 
even though Bannuk could never admit 
it to me, as he grew older I could see 
that same desire in his eyes.
 
Oh...It's a pity I didn't encourage him to 
go with you. He might still be alive 
today.
`},
	2230: {ID: 2230, Key: "MalahGossip3", Value: `57
Though once considered shelter by our 
people, the Ice Caves offer no refuge 
from the dark horde. There are 
creatures there that will freeze your 
heart before feasting upon it.
`},
	2231: {ID: 2231, Key: "MalahGossip4", Value: `107
On the battlefield, turning your back on 
even the dead is unwise.
`},
	2232: {ID: 2232, Key: "MalahGossip5", Value: `60
This battle plays mind tricks on our 
warriors. Those fortunate enough to 
have returned from the mountainside 
claim to have seen angels in flight.
 
Fly they might, but that certainly does 
not make them angels.
`},
	2233: {ID: 2233, Key: "MalahGossip6", Value: `53
Perhaps you have heard the accounts 
of my son's horrible death at the hands 
of Baal's minions. He was my last living 
child...
 
The oath of compassion I have taken as 
a healer extends only to humankind.
 
Cut them down, warrior. All of them!
`},
	2234: {ID: 2234, Key: "MalahGossip7", Value: `78
The catapults are infernal machines 
made of demon flesh fused with steel.
 
Be wary of them.
`},
	2235: {ID: 2235, Key: "MalahGossip8", Value: `60
Qual-Kehk is a worthy leader, but the 
losses have borne heavily upon him. He 
is only impatient because he judges the 
worth of a warrior by action, not 
words.
 
You must prove yourself worthy of his 
trust.
`},
	2236: {ID: 2236, Key: "MalahGossip9", Value: `52
Larzuk possesses a good soul, but at 
times his mind seems quite unsound. 
 
He once asked me for twenty of my 
finest sheepskins. He said he would fill 
them with hot air and float like a cloud 
above the battlefield to spy on Baal's 
legions! 
 
I worry the siege has driven him mad.
`},
	2237: {ID: 2237, Key: "A5Q3EarlyReturnQualKehk", Value: `81
It seems like everyone feels Nihlathak 
played a part in Anya's disappearance.
 
Why would he do such a thing?
`},
	2238: {ID: 2238, Key: "Jacinth", Value: `Jacinth`},
	2239: {ID: 2239, Key: "Moldy", Value: `Moldy`},
	2240: {ID: 2240, Key: "A5Q5EarlyReturnAnya", Value: `71
Look. I must apologize.
 
I feel responsible for your current 
struggle. If I had only stopped 
Nihlathak from giving Baal the Relic, 
you would not have to fight those 
ghosts.
`},
	2241: {ID: 2241, Key: "of Armageddon", Value: `of Armageddon`},
	2242: {ID: 2242, Key: "of Good Luck", Value: `of Good Luck`},
	2243: {ID: 2243, Key: "of Chilling Armor", Value: `of Chilling Armor`},
	2244: {ID: 2244, Key: "Blood Comet", Value: `Blood Comet`},
	2245: {ID: 2245, Key: "Indiego's Fancy", Value: `Indiego's Fancy`},
	2246: {ID: 2246, Key: "Kelpie Snare", Value: `Kelpie Snare`},
	2247: {ID: 2247, Key: "Blackoak Shield", Value: `Blackoak Shield`},
	2248: {ID: 2248, Key: "FindingGuardianTowerDru", Value: `Hatred stirs within me.`},
	2249: {ID: 2249, Key: "Of Luck", Value: `of Luck`},
	2250: {ID: 2250, Key: "EnteringWorldstoneAct5Dru", Value: `The legendary Worldstone - guardian of the Natural realm.`},
	2251: {ID: 2251, Key: "Vampiregaze", Value: `Vampire Gaze`},
	2252: {ID: 2252, Key: "Skillan226", Value: `Firestorm`},
	2253: {ID: 2253, Key: "Skillan227", Value: `Oak Sage`},
	2254: {ID: 2254, Key: "Skillan228", Value: `Sum Spt Wolf`},
	2255: {ID: 2255, Key: "Vodoun", Value: `Mojo`},
	2256: {ID: 2256, Key: "Skillan229", Value: `Werebear`},
	2257: {ID: 2257, Key: "ModStre9s", Value: `Indestructible`},
	2258: {ID: 2258, Key: "A5Q2AfterInitQualKehk", Value: `85
Those of my men fortunate enough to 
escape on their own told me that they 
were held captive in the highlands and 
plateaus.
`},
	2259: {ID: 2259, Key: "ModStre9t", Value: `Repairs %d durability per second`},
	2260: {ID: 2260, Key: "Brown", Value: `Brown`},
	2261: {ID: 2261, Key: "of the Colossus1", Value: `of the Colossus`},
	2262: {ID: 2262, Key: "of Inner Sight", Value: `of Inner Sight`},
	2263: {ID: 2263, Key: "ModStre9u", Value: `Repairs %d durability in %d seconds`},
	2264: {ID: 2264, Key: "Lilac", Value: `Lilac`},
	2265: {ID: 2265, Key: "ModStre9v", Value: `Replenishes quantity`},
	2266: {ID: 2266, Key: "ModStre9w", Value: `Cast a Level %d`},
	2267: {ID: 2267, Key: "A5Q1EarlyReturnQualKehk", Value: `93
So, you still live. You're either quick or a 
coward.
`},
	2268: {ID: 2268, Key: "A5Q4EarlyReturnLarzuk", Value: `160
...What's there to talk about?
 
Kill Nihlathak!
`},
	2269: {ID: 2269, Key: "ModStre9x", Value: `When You Swing`},
	2270: {ID: 2270, Key: "Of the Lake", Value: `of the Lake`},
	2271: {ID: 2271, Key: "Athlete's", Value: `Athletic`},
	2272: {ID: 2272, Key: "Of the Infantry", Value: `of the Infantry`},
	2273: {ID: 2273, Key: "ModStre9y", Value: `When You Get Hit`},
	2274: {ID: 2274, Key: "Of the Ghost", Value: `of the Ghost`},
	2275: {ID: 2275, Key: "ModStre9z", Value: `When You Hit an Enemy`},
	2276: {ID: 2276, Key: "Stormrider", Value: `Stormrider`},
	2277: {ID: 2277, Key: "Fleshrender", Value: `Fleshrender`},
	2278: {ID: 2278, Key: "Shadowdancer", Value: `Shadow Dancer`},
	2279: {ID: 2279, Key: "The Harvester", Value: `The Harvester`},
	2280: {ID: 2280, Key: "Corpsemourn", Value: `Corpsemourn`},
	2281: {ID: 2281, Key: "of Luck", Value: `of Luck`},
	2282: {ID: 2282, Key: "The Meat Scraper", Value: `The Meat Scraper`},
	2283: {ID: 2283, Key: "Runemaster", Value: `Rune Master`},
	2284: {ID: 2284, Key: "Stealskull", Value: `Stealskull`},
	2285: {ID: 2285, Key: "Mortal Crescent", Value: `Mortal Crescent`},
	2286: {ID: 2286, Key: "Eskillincaseraven", Value: `Mana Cost Per Raven: `},
	2287: {ID: 2287, Key: "of Desire", Value: `of Desire`},
	2288: {ID: 2288, Key: "Giantskull", Value: `Giant Skull`},
	2289: {ID: 2289, Key: "Ginther's Rift", Value: `Ginther's Rift`},
	2290: {ID: 2290, Key: "Wilhelm's Pride", Value: `Wilhelm's Pride`},
	2291: {ID: 2291, Key: "Siren's call", Value: `Siren's Call`},
	2292: {ID: 2292, Key: "of Memory", Value: `of Memory`},
	2293: {ID: 2293, Key: "Skillan230", Value: `Molten Boulder`},
	2294: {ID: 2294, Key: "Razoredge", Value: `Razor's Edge`},
	2295: {ID: 2295, Key: "Skillan231", Value: `Arctic Blast`},
	2296: {ID: 2296, Key: "Skillan232", Value: `Carrion Vine`},
	2297: {ID: 2297, Key: "EnterDOEAss", Value: `So dark... perfect.`},
	2298: {ID: 2298, Key: "Tang's Imperial Robes", Value: `Tang's Imperial Robes`},
	2299: {ID: 2299, Key: "Lanceguard", Value: `Lance Guard`},
	2300: {ID: 2300, Key: "Skillan233", Value: `Feral Rage`},
	2301: {ID: 2301, Key: "Skillan234", Value: `Maul`},
	2302: {ID: 2302, Key: "Skillan235", Value: `Fissure`},
	2303: {ID: 2303, Key: "Skillan236", Value: `Cyclone Armor`},
	2304: {ID: 2304, Key: "Stonerage", Value: `Stonerage`},
	2305: {ID: 2305, Key: "Of the Hag", Value: `of the Hag`},
	2306: {ID: 2306, Key: "of Impaling Spear", Value: `of Impaling Spear`},
	2307: {ID: 2307, Key: "of Firestorms", Value: `of Firestorms`},
	2308: {ID: 2308, Key: "of Lightning Spear", Value: `of Lightning Spear`},
	2309: {ID: 2309, Key: "The Vile Husk", Value: `The Vile Husk`},
	2310: {ID: 2310, Key: "Gorerider", Value: `Gore Rider`},
	2311: {ID: 2311, Key: "Tal Rasha's Horadric Crest", Value: `Tal Rasha's Horadric Crest`},
	2312: {ID: 2312, Key: "of Lower Resistance", Value: `of Lower Resistance`},
	2313: {ID: 2313, Key: "Glacial", Value: `Glacial`},
	2314: {ID: 2314, Key: "Stormspike", Value: `Stormspike`},
	2315: {ID: 2315, Key: "FreezingIzualDru", Value: `I have no pity for him. Oblivion is his reward.`},
	2316: {ID: 2316, Key: "LeavingTownAct5Ams", Value: `The siege must be stopped.`},
	2317: {ID: 2317, Key: "Skillan237", Value: `Wolverine Hrt`},
	2318: {ID: 2318, Key: "Pus Spiter", Value: `Pus Spitter`},
	2319: {ID: 2319, Key: "Skillan238", Value: `Summon D Wolf`},
	2320: {ID: 2320, Key: "Carin Shard", Value: `Carin Shard`},
	2321: {ID: 2321, Key: "Rotting", Value: `Rotting`},
	2322: {ID: 2322, Key: "Willhelm's Pride", Value: `Willhelm's Pride`},
	2323: {ID: 2323, Key: "Skillan239", Value: `Rabies`},
	2324: {ID: 2324, Key: "Drulan's Tounge", Value: `Drulan's Tongue`},
	2325: {ID: 2325, Key: "Skillan240", Value: `Fire Claws`},
	2326: {ID: 2326, Key: "EnterCatacombsAss", Value: `I don't like it down here.`},
	2327: {ID: 2327, Key: "of Fervor", Value: `of Fervor`},
	2328: {ID: 2328, Key: "Of Terror", Value: `of Terror`},
	2329: {ID: 2329, Key: "of Glacial Spikes", Value: `of Glacial Spike`},
	2330: {ID: 2330, Key: "Skillan241", Value: `Twister`},
	2331: {ID: 2331, Key: "Skillan242", Value: `Sol Creep`},
	2332: {ID: 2332, Key: "Skillan243", Value: `Hunger`},
	2333: {ID: 2333, Key: "Buzzing", Value: `Buzzing`},
	2334: {ID: 2334, Key: "Skillan244", Value: `Shock Wave`},
	2335: {ID: 2335, Key: "Skillan245", Value: `Volcano`},
	2336: {ID: 2336, Key: "Skillan246", Value: `Tornado`},
	2337: {ID: 2337, Key: "Skillan247", Value: `Spirit Barbs`},
	2338: {ID: 2338, Key: "A5Q2EarlyReturnMalah", Value: `67
Soldiers who made it back told of a 
system of barricades placed among the 
mountain passes. They said that is 
where the prisoners are held.
`},
	2339: {ID: 2339, Key: "Of the Virgin", Value: `of the Virgin`},
	2340: {ID: 2340, Key: "Skillld248", Value: `summon a ferocious grizzly bear`},
	2341: {ID: 2341, Key: "A5Q4InitAnya", Value: `43
Nihlathak told me he struck a deal with 
Baal to protect Harrogath. In exchange 
for the demon's mercy, the misguided 
fool plans to give Baal the Relic of the 
Ancients, our most holy totem!
 
Doing so will allow Baal to enter Mount 
Arreat unchallenged by the Ancients. I 
tried to stop Nihlathak, but he 
imprisoned me in that icy tomb.
 
Nihlathak must be stopped before he 
dooms the whole world. As much as I 
would love to strangle the life out of 
him, I'm afraid I haven't the strength.
 
You must go to his lair through this 
portal I've opened, kill him, and then 
bring back the Relic of the Ancients.
 
Stop Nihlathak from destroying what we 
have striven for eons to protect.
`},
	2342: {ID: 2342, Key: "Freezing", Value: `Freezing`},
	2343: {ID: 2343, Key: "String of Ears", Value: `String of Ears`},
	2344: {ID: 2344, Key: "9lw", Value: `Greater Claws`},
	2345: {ID: 2345, Key: "Veteran's", Value: `Veteran's`},
	2346: {ID: 2346, Key: "Wraithfang", Value: `Wraithfang`},
	2347: {ID: 2347, Key: "Schaefer's Hammer", Value: `Schaefer's Hammer`},
	2348: {ID: 2348, Key: "Tang's Battle Standard", Value: `Tang's Battle Standard`},
	2349: {ID: 2349, Key: "Skillan248", Value: `Summon Grizzly`},
	2350: {ID: 2350, Key: "A5Q1AfterInitQualKehk", Value: `66
About to face Shenk the Overseer and 
stop the siege, are you? You should 
ask Malah to perform your last rites 
before you go, stranger.
`},
	2351: {ID: 2351, Key: "Skillname249", Value: `Fury`},
	2352: {ID: 2352, Key: "StrMercEx12", Value: `Defensive`},
	2353: {ID: 2353, Key: "Skillld249", Value: `or one target multiple times
either multiple adjacent targets
when in werewolf form, attack`},
	2354: {ID: 2354, Key: "Bul-Kathos' Sacred Charge", Value: `Bul-Kathos' Sacred Charge`},
	2355: {ID: 2355, Key: "Kerke's Sanctuary", Value: `Gerke's Sanctuary`},
	2356: {ID: 2356, Key: "StrMercEx14", Value: `Offensive`},
	2357: {ID: 2357, Key: "StrMercEx15", Value: `Combat`},
	2358: {ID: 2358, Key: "Skillan249", Value: `Fury`},
	2359: {ID: 2359, Key: "Skillname250", Value: `Armageddon`},
	2360: {ID: 2360, Key: "Of Karma", Value: `of Karma`},
	2361: {ID: 2361, Key: "The Black Adder", Value: `The Black Adder`},
	2362: {ID: 2362, Key: "Skillld250", Value: `destruction on nearby enemies
create a meteor shower to rain fiery`},
	2363: {ID: 2363, Key: "Of Cruelty", Value: `of Cruelty`},
	2364: {ID: 2364, Key: "of Fire Quenching", Value: `of Quenching`},
	2365: {ID: 2365, Key: "Wrath of Cain", Value: `Wrath of Cain`},
	2366: {ID: 2366, Key: "Skillan250", Value: `Armageddon`},
	2367: {ID: 2367, Key: "Skillname251", Value: `Hurricane`},
	2368: {ID: 2368, Key: "Skillld251", Value: `debris to pound your enemies to bits
create a massive storm of wind and`},
	2369: {ID: 2369, Key: "Homunculus", Value: `Homunculus`},
	2370: {ID: 2370, Key: "Skillan251", Value: `Hurricane`},
	2371: {ID: 2371, Key: "Skillname252", Value: `Fire Blast`},
	2372: {ID: 2372, Key: "Lestron's Mark", Value: `Lestron's Mark`},
	2373: {ID: 2373, Key: "Spearmaiden's", Value: `Spearmaiden's`},
	2374: {ID: 2374, Key: "Skillld252", Value: `to blast your enemies to bits
throw a fire bomb`},
	2375: {ID: 2375, Key: "Skillan252", Value: `Fire Blast`},
	2376: {ID: 2376, Key: "Skillname253", Value: `Claw Mastery`},
	2377: {ID: 2377, Key: "Skillld253", Value: `with claw-class weapons
passive - improves your skill`},
	2378: {ID: 2378, Key: "Skillan253", Value: `Claw Mastery`},
	2379: {ID: 2379, Key: "Skillname254", Value: `Psychic Hammer`},
	2380: {ID: 2380, Key: "Skillld254", Value: `to crush and knock back your enemies
to create a psychic blast
use the power of your mind`},
	2381: {ID: 2381, Key: "of Stunning", Value: `of Stun`},
	2382: {ID: 2382, Key: "Skillan254", Value: `Psyc Hammer`},
	2383: {ID: 2383, Key: "Runeword150", Value: `Truth`},
	2384: {ID: 2384, Key: "Runeword151", Value: `Unbending Will`},
	2385: {ID: 2385, Key: "Runeword152", Value: `Valor`},
	2386: {ID: 2386, Key: "Runeword153", Value: `Vengeance`},
	2387: {ID: 2387, Key: "Runeword154", Value: `Venom`},
	2388: {ID: 2388, Key: "Runeword155", Value: `Victory`},
	2389: {ID: 2389, Key: "Runeword156", Value: `Voice`},
	2390: {ID: 2390, Key: "Runeword157", Value: `Void`},
	2391: {ID: 2391, Key: "Runeword158", Value: `War`},
	2392: {ID: 2392, Key: "Runeword159", Value: `Water`},
	2393: {ID: 2393, Key: "Shogukusha's", Value: `Shogukusha's`},
	2394: {ID: 2394, Key: "Skystrike", Value: `Skystrike`},
	2395: {ID: 2395, Key: "El Infierno", Value: `El Infierno`},
	2396: {ID: 2396, Key: "Templar's Might", Value: `Templar's Might`},
	2397: {ID: 2397, Key: "Lifechoke", Value: `Lifechoke`},
	2398: {ID: 2398, Key: "Tang's Throne", Value: `Tang's Throne`},
	2399: {ID: 2399, Key: "Runeword160", Value: `Wealth`},
	2400: {ID: 2400, Key: "Runeword161", Value: `Whisper`},
	2401: {ID: 2401, Key: "Runeword162", Value: `White`},
	2402: {ID: 2402, Key: "Runeword163", Value: `Wind`},
	2403: {ID: 2403, Key: "Runeword164", Value: `Wings of Hope`},
	2404: {ID: 2404, Key: "Runeword165", Value: `Wisdom`},
	2405: {ID: 2405, Key: "LarzukAct5IntroAmaGossip1", Value: `42
So, you're an Amazon. I have heard 
rumors about your kind. Your unusual 
weapons could prove a challenge to my 
skills, but I'm prepared to meet it.
 
I am Larzuk, the armorer. My ancestors 
were some of the finest craftsmen in 
Harrogath. Regretfully, my supplies run 
lower with every passing day, yet the 
demons beyond the walls have not 
weakened. I fear the time is near when 
I must put down my hammer and take 
up a sword, instead.
`},
	2406: {ID: 2406, Key: "Runeword166", Value: `Woe`},
	2407: {ID: 2407, Key: "A5Q6EarlyReturnMalah", Value: `70
You knew it would eventually come down 
to this. Kill Baal! Finish the game!
`},
	2408: {ID: 2408, Key: "Runeword167", Value: `Wonder`},
	2409: {ID: 2409, Key: "Runeword168", Value: `Wrath`},
	2410: {ID: 2410, Key: "Runeword169", Value: `Youth`},
	2411: {ID: 2411, Key: "Hellatial", Value: `Hellacious`},
	2412: {ID: 2412, Key: "of Charging", Value: `of Charge`},
	2413: {ID: 2413, Key: "Shaftstop", Value: `Shaftstop`},
	2414: {ID: 2414, Key: "Duriel's Shell", Value: `Duriel's Shell`},
	2415: {ID: 2415, Key: "Runeword170", Value: `Zephyr`},
	2416: {ID: 2416, Key: "Guardian Angle", Value: `Guardian Angel`},
	2417: {ID: 2417, Key: "of Grace", Value: `of Grace`},
	2418: {ID: 2418, Key: "Bul-Kathos", Value: `Bul-Kathos`},
	2419: {ID: 2419, Key: "9qr", Value: `Scissors Quhab`},
	2420: {ID: 2420, Key: "Stormspire", Value: `Stormspire`},
	2421: {ID: 2421, Key: "Howling Visage", Value: `Howling Visage`},
	2422: {ID: 2422, Key: "Skillname255", Value: `Tiger Strike`},
	2423: {ID: 2423, Key: "Skillld255", Value: `normal attack to release charges
must use a dragon finishing move or
to finishing moves
consecutive hits add damage bonuses

Charge-up Skill
`},
	2424: {ID: 2424, Key: "Skillan255", Value: `Tiger Strike`},
	2425: {ID: 2425, Key: "CompletingAndarielAss", Value: `Death becomes you, Andariel.`},
	2426: {ID: 2426, Key: "Skillname256", Value: `Dragon Talon`},
	2427: {ID: 2427, Key: "Skillld256", Value: `adds charged-up bonuses to the kick
kick your enemies out of your way

Finishing Move
`},
	2428: {ID: 2428, Key: "of Joyfulness", Value: `of Joyfulness`},
	2429: {ID: 2429, Key: "Skillan256", Value: `Dragon Talon`},
	2430: {ID: 2430, Key: "Lavagout", Value: `Lava Gout`},
	2431: {ID: 2431, Key: "EskillWolfDef", Value: `Wolf Defense: `},
	2432: {ID: 2432, Key: "of the Yeti", Value: `of the Yeti`},
	2433: {ID: 2433, Key: "Skillname257", Value: `Shock Web`},
	2434: {ID: 2434, Key: "of Charged Shield", Value: `of Charged Bolt`},
	2435: {ID: 2435, Key: "Skillld257", Value: `to shock your enemies
throw a web of lightning`},
	2436: {ID: 2436, Key: "of Fire Balls", Value: `of Fire Ball`},
	2437: {ID: 2437, Key: "Naj's Light Plate", Value: `Naj's Light Plate`},
	2438: {ID: 2438, Key: "Pernicious", Value: `Pernicious`},
	2439: {ID: 2439, Key: "Skillsd222", Value: `summon ravens`},
	2440: {ID: 2440, Key: "Sandstorm Trek", Value: `Sandstorm Trek`},
	2441: {ID: 2441, Key: "Wolfhowl", Value: `Wolfhowl`},
	2442: {ID: 2442, Key: "Skillsd223", Value: `summon disease spreading vine`},
	2443: {ID: 2443, Key: "Skillsd223", Value: `transform into a werewolf`},
	2444: {ID: 2444, Key: "A5Q1AfterInitMalah", Value: `77
Qual-Kehk and his men have been 
fighting to break the siege for some 
time. Where many have failed, you may 
succeed.
`},
	2445: {ID: 2445, Key: "Skillsd225", Value: `passive - improves shape-shifting ability`},
	2446: {ID: 2446, Key: "Skillsd226", Value: `unleash fiery chaos`},
	2447: {ID: 2447, Key: "Runeword110", Value: `Principle`},
	2448: {ID: 2448, Key: "Runeword111", Value: `Prowess in Battle`},
	2449: {ID: 2449, Key: "Runeword112", Value: `Prudence`},
	2450: {ID: 2450, Key: "Runeword113", Value: `Punishment`},
	2451: {ID: 2451, Key: "Runeword114", Value: `Purity`},
	2452: {ID: 2452, Key: "A5Q6EarlyReturnLarzuk", Value: `59
I may be just an armorer, but I know 
this...Baal plans to destroy the world 
with the secrets contained in that 
mountain. It doesn't take a genius to 
know he has to be stopped.
`},
	2453: {ID: 2453, Key: "Runeword115", Value: `Question`},
	2454: {ID: 2454, Key: "Runeword116", Value: `Radiance`},
	2455: {ID: 2455, Key: "Runeword117", Value: `Rain`},
	2456: {ID: 2456, Key: "Runeword118", Value: `Reason`},
	2457: {ID: 2457, Key: "Runeword119", Value: `Red`},
	2458: {ID: 2458, Key: "Antimagic", Value: `Antimagic`},
	2459: {ID: 2459, Key: "Smoldering", Value: `Smoldering`},
	2460: {ID: 2460, Key: "Sparroweye", Value: `Sparroweye`},
	2461: {ID: 2461, Key: "Insidious", Value: `Insidious`},
	2462: {ID: 2462, Key: "Of the Grassy Gnoll", Value: `of the Grassy Gnoll`},
	2463: {ID: 2463, Key: "Runeword120", Value: `Rhyme`},
	2464: {ID: 2464, Key: "Runeword121", Value: `Rift`},
	2465: {ID: 2465, Key: "Runeword122", Value: `Sanctuary`},
	2466: {ID: 2466, Key: "Runeword123", Value: `Serendipity`},
	2467: {ID: 2467, Key: "Runeword124", Value: `Shadow`},
	2468: {ID: 2468, Key: "Runeword125", Value: `Shadow of Doubt`},
	2469: {ID: 2469, Key: "Runeword126", Value: `Silence`},
	2470: {ID: 2470, Key: "Runeword127", Value: `Siren's Song`},
	2471: {ID: 2471, Key: "Runeword128", Value: `Smoke`},
	2472: {ID: 2472, Key: "9tw", Value: `Greater Talons`},
	2473: {ID: 2473, Key: "Runeword129", Value: `Sorrow`},
	2474: {ID: 2474, Key: "Of the Gnat", Value: `of the Gnat`},
	2475: {ID: 2475, Key: "Of the Walrus", Value: `of the Walrus`},
	2476: {ID: 2476, Key: "Dun", Value: `Dun`},
	2477: {ID: 2477, Key: "Alma Negra", Value: `Alma Negra`},
	2478: {ID: 2478, Key: "Natalya's Mark", Value: `Natalya's Mark`},
	2479: {ID: 2479, Key: "Runeword130", Value: `Spirit`},
	2480: {ID: 2480, Key: "Runeword131", Value: `Splendor`},
	2481: {ID: 2481, Key: "Runeword132", Value: `Starlight`},
	2482: {ID: 2482, Key: "Runeword133", Value: `Stealth`},
	2483: {ID: 2483, Key: "Runeword134", Value: `Steel`},
	2484: {ID: 2484, Key: "Runeword135", Value: `Still Water`},
	2485: {ID: 2485, Key: "Runeword136", Value: `Sting`},
	2486: {ID: 2486, Key: "Runeword137", Value: `Stone`},
	2487: {ID: 2487, Key: "Runeword138", Value: `Storm`},
	2488: {ID: 2488, Key: "Runeword139", Value: `Strength`},
	2489: {ID: 2489, Key: "Screaming", Value: `Screaming`},
	2490: {ID: 2490, Key: "Astral", Value: `Astral`},
	2491: {ID: 2491, Key: "Arcadian", Value: `Arcadian`},
	2492: {ID: 2492, Key: "Scintillating", Value: `Scintillating`},
	2493: {ID: 2493, Key: "Chestnut", Value: `Chestnut`},
	2494: {ID: 2494, Key: "of the Dynamo", Value: `of the Dynamo`},
	2495: {ID: 2495, Key: "Runeword140", Value: `Tempest`},
	2496: {ID: 2496, Key: "Runeword141", Value: `Temptation`},
	2497: {ID: 2497, Key: "Runeword142", Value: `Terror`},
	2498: {ID: 2498, Key: "Runeword143", Value: `Thirst`},
	2499: {ID: 2499, Key: "9wb", Value: `Wrist Spike`},
	2500: {ID: 2500, Key: "Runeword144", Value: `Thought`},
	2501: {ID: 2501, Key: "Runeword145", Value: `Thunder`},
	2502: {ID: 2502, Key: "Runeword146", Value: `Time`},
	2503: {ID: 2503, Key: "Runeword147", Value: `Tradition`},
	2504: {ID: 2504, Key: "Runeword148", Value: `Treachery`},
	2505: {ID: 2505, Key: "MalahAct5IntroBarGossip1", Value: `34
You have traveled far only to return 
home to us, Barbarian. Ohh...Much has 
happened in Harrogath since you left. 
Our homeland is hardly recognizable 
with so much evil about.
 
Yet, I've managed to survive so far. 
You've seen your share of suffering as 
well, I'm sure. Seeing you again -- alive 
-- does my heart good. I shall pray that 
you can help us out of this hell.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what.
 
Nihlathak is the last survivor of the 
Council of Elders, but I'm afraid he has 
not been himself lately. The other 
Elders sacrificed themselves to place a 
protective ward around Harrogath. If 
Nihlathak is curt with you, it is because 
he feels responsible for our situation. 
He does not relish sending more of our 
people out to die.
 
So much has changed since you left 
that I see little hope for us in this life.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`},
	2506: {ID: 2506, Key: "Runeword149", Value: `Trust`},
	2507: {ID: 2507, Key: "Yelling", Value: `Yelling`},
	2508: {ID: 2508, Key: "of the Dunes", Value: `of the Dunes`},
	2509: {ID: 2509, Key: "Boneflame", Value: `Boneflame`},
	2510: {ID: 2510, Key: "Spirit Ward", Value: `Spirit Ward`},
	2511: {ID: 2511, Key: "Aragon's Storm Cloud", Value: `Aragon's Storm Cloud`},
	2512: {ID: 2512, Key: "Eskillincasemastery", Value: `%d Percent Chance of Critical Strike`},
	2513: {ID: 2513, Key: "FindingCairnDru", Value: `Such stones are common back home.`},
	2514: {ID: 2514, Key: "FreezingIzualAma", Value: `Goodbye, Izual.`},
	2515: {ID: 2515, Key: "Baezils Vortex", Value: `Baezil's Vortex`},
	2516: {ID: 2516, Key: "Gerke's Sanctuary", Value: `Gerke's Sanctuary`},
	2517: {ID: 2517, Key: "Sazabi's Ghost Liberator", Value: `Sazabi's Ghost Liberator`},
	2518: {ID: 2518, Key: "Dangoon's Teaching", Value: `Dangoon's Teaching`},
	2519: {ID: 2519, Key: "9xf", Value: `Fascia`},
	2520: {ID: 2520, Key: "Cunning", Value: `Cunning`},
	2521: {ID: 2521, Key: "The Wanderer's Blade", Value: `The Wanderer's Blade`},
	2522: {ID: 2522, Key: "KillingdDiabloAms", Value: `The reign of Terror has ended.`},
	2523: {ID: 2523, Key: "Nosferatu's Coil", Value: `Nosferatu's Coil`},
	2524: {ID: 2524, Key: "Foci of Visjerei", Value: `Foci of the Vizjerei`},
	2525: {ID: 2525, Key: "FindingTristramAss", Value: `Tristram... The first to fall to Diablo's wrath.`},
	2526: {ID: 2526, Key: "A5Q5SuccessfulLarzuk", Value: `82
The Ancients have honored you, and in 
turn, so do we. I have no remaining 
doubts about you, now.
`},
	2527: {ID: 2527, Key: "Butcher's Pupil", Value: `Butcher's Pupil`},
	2528: {ID: 2528, Key: "Hwanin's Seal", Value: `Hwanin's Blessing`},
	2529: {ID: 2529, Key: "Immortal King's Detail", Value: `Immortal King's Detail`},
	2530: {ID: 2530, Key: "EnteringTopMountAct5Bar", Value: `The guardians of Mount Arreat await.`},
	2531: {ID: 2531, Key: "EnteringTopMountAct5Ass", Value: `I shall prove worthy.`},
	2532: {ID: 2532, Key: "Hermetic", Value: `Hermetic`},
	2533: {ID: 2533, Key: "Moonrend", Value: `Moonrend`},
	2534: {ID: 2534, Key: "Dwarf Star", Value: `Dwarf Star`},
	2535: {ID: 2535, Key: "Peasent Crown", Value: `Peasant Crown`},
	2536: {ID: 2536, Key: "Haemosu's Adament", Value: `Haemosu's Adamant`},
	2537: {ID: 2537, Key: "Skillsd227", Value: `summon the spirit of the oak`},
	2538: {ID: 2538, Key: "Skillsd228", Value: `summon a wolf`},
	2539: {ID: 2539, Key: "Skillsd229", Value: `transform into a werebear`},
	2540: {ID: 2540, Key: "Orumus' Robes", Value: `Ormus' Robe`},
	2541: {ID: 2541, Key: "Tang's Fore-Fathers", Value: `Tang's Fore-Fathers`},
	2542: {ID: 2542, Key: "Skillsd230", Value: `launch a molten boulder`},
	2543: {ID: 2543, Key: "eyz", Value: `Eye`},
	2544: {ID: 2544, Key: "Skillsd231", Value: `shoot a jet of ice`},
	2545: {ID: 2545, Key: "strepilogueX", Value: `TERROR'S END`},
	2546: {ID: 2546, Key: "WeaponDescOrb", Value: `Orb Class`},
	2547: {ID: 2547, Key: "Skillsd232", Value: `summon corpse eating vine`},
	2548: {ID: 2548, Key: "Skullder's Ire", Value: `Skullder's Ire`},
	2549: {ID: 2549, Key: "A5Q1SuccessfulLarzuk", Value: `63
You're an even greater warrior than I 
expected...Sorry for underestimating 
you.
 
As a token of my appreciation, I will 
craft sockets into an item of your 
choosing, and from now on, you'll get 
the best price for all my wares.
`},
	2550: {ID: 2550, Key: "Skillsd233", Value: `life-stealing rage attack - werewolf form`},
	2551: {ID: 2551, Key: "Skillsd234", Value: `maul your enemies - werebear form`},
	2552: {ID: 2552, Key: "A5Q2SuccessfulCain", Value: `80
You've become a hero to this town, my 
friend. The shadows have lifted ever 
since you brought the Light to 
Harrogath.
`},
	2553: {ID: 2553, Key: "Original", Value: `Original`},
	2554: {ID: 2554, Key: "The Minotaur", Value: `The Minotaur`},
	2555: {ID: 2555, Key: "MalahGossip10", Value: `90
Be cautious, warrior.
 
Though I am an experienced healer, 
resurrection is beyond my powers.
`},
	2556: {ID: 2556, Key: "MalahGossip11", Value: `53
I pray for the day when the fields are 
covered in snow unstained by the blood 
of our own. Perhaps that day will come 
soon...Perhaps never...
 
Oh...But I forget myself. How may I aid 
your quest?
`},
	2557: {ID: 2557, Key: "MalahGossip12", Value: `64
Your gold is most helpful. Medical 
supplies for our wounded are scarce 
and very expensive.
 
When we can find a supplier, oh, we 
must pay dearly for what we need.
`},
	2558: {ID: 2558, Key: "MalahGossip13", Value: `44
With the exception of Qual-Kehk, the 
townspeople do not see what I see -- 
the massacre we face.
 
Our bravest, strongest heroes hobble 
back to me begging for help. I do what 
I can -- healing and bandaging some, 
preparing the others for what lies 
beyond.
`},
	2559: {ID: 2559, Key: "Runeword100", Value: `Penitence`},
	2560: {ID: 2560, Key: "Runeword101", Value: `Peril`},
	2561: {ID: 2561, Key: "Runeword102", Value: `Pestilence`},
	2562: {ID: 2562, Key: "Runeword103", Value: `Phoenix`},
	2563: {ID: 2563, Key: "Runeword104", Value: `Piety`},
	2564: {ID: 2564, Key: "Runeword105", Value: `Pillar of Faith`},
	2565: {ID: 2565, Key: "Runeword106", Value: `Plague`},
	2566: {ID: 2566, Key: "Runeword107", Value: `Praise`},
	2567: {ID: 2567, Key: "Runeword108", Value: `Prayer`},
	2568: {ID: 2568, Key: "Runeword109", Value: `Pride`},
	2569: {ID: 2569, Key: "of Thunder Storm", Value: `of Thunder Storm`},
	2570: {ID: 2570, Key: "EnteringClawViperDru", Value: `Snakes... I hate snakes.`},
	2571: {ID: 2571, Key: "Skillsd235", Value: `open the earth to burn enemies`},
	2572: {ID: 2572, Key: "flg", Value: `Flag`},
	2573: {ID: 2573, Key: "Decaying", Value: `Decaying`},
	2574: {ID: 2574, Key: "qll", Value: `Quill`},
	2575: {ID: 2575, Key: "To Hell1", Value: `To Abaddon`},
	2576: {ID: 2576, Key: "Stout", Value: `Stout`},
	2577: {ID: 2577, Key: "A5Q3AfterInitNihlathak", Value: `41
Anya! Who have you been talking to? 
Likely it was that meddling Malah. 
 
Well, I'll tell you what really happened. 
Anya came to me for guidance, after 
receiving a vision that her mother and 
younger brother were trapped in the 
lands beyond the Ice Caves. She had 
decided to go rescue them. 
 
I told her that her quest was a foolish 
one and that she would be safer 
staying within the city walls. However, 
she is a willful girl and would not listen 
to me. 
 
The next morning, she was gone. No 
one is more distraught than I over 
losing her. 
 
However, if you feel the need to be 
Malah's errand child, I won't try to stop 
you.
`},
	2578: {ID: 2578, Key: "Pestilent", Value: `Pestilent`},
	2579: {ID: 2579, Key: "Monumental", Value: `Monumental`},
	2580: {ID: 2580, Key: "Rose", Value: `Rose`},
	2581: {ID: 2581, Key: "Runeslayer", Value: `Rune Slayer`},
	2582: {ID: 2582, Key: "To Hell2", Value: `To The Pit of Acheron`},
	2583: {ID: 2583, Key: "of Coolness", Value: `of Coolness`},
	2584: {ID: 2584, Key: "Nebucaneezer's Storm", Value: `Nebuchadnezzar's Storm`},
	2585: {ID: 2585, Key: "To Hell3", Value: `To The Infernal Pit`},
	2586: {ID: 2586, Key: "EnteringNihlathakAct5Dru", Value: `Nihlathak... you can't hide from me.`},
	2587: {ID: 2587, Key: "Natalya's Totem", Value: `Natalya's Totem`},
	2588: {ID: 2588, Key: "Skillsd236", Value: `shield from elemental damage`},
	2589: {ID: 2589, Key: "Skillsd237", Value: `summon a wolverine spirit`},
	2590: {ID: 2590, Key: "Of the Moon", Value: `of the Moon`},
	2591: {ID: 2591, Key: "Jalal's Mane", Value: `Jalal's Mane`},
	2592: {ID: 2592, Key: "Skillsd238", Value: `summon an enraged wolf`},
	2593: {ID: 2593, Key: "Malignant", Value: `Malignant Skull`},
	2594: {ID: 2594, Key: "MalahAct5IntroSorGossip1", Value: `36
A Sorceress...Here in Harrogath?
 
There was a time, child, when I thought 
I was destined to follow your kind's 
path. However, my powers never 
developed beyond the simplest of 
spells. Although I can heal almost any 
wound with time and energy, there is 
little I can do to help against Baal.
 
But enough of that...I spend too much 
time in reflection and have forgotten 
my manners.
 
I, Malah, welcome you to Harrogath, 
the last stronghold of Order on Mount 
Arreat. You have come to the right 
place, if you intend to defeat Baal the 
Lord of Destruction.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what. 
 
All of the Elders, save Nihlathak, 
sacrificed themselves to place a 
protective ward around Harrogath.
 
Some of us here, certainly Nihlathak, do 
not appreciate your presence. We are a 
proud people, and it is not easy for us 
to accept aid. I, however, am glad you 
are here.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`},
	2595: {ID: 2595, Key: "Butchers Cleaver", Value: `Butcher's Cleaver`},
	2596: {ID: 2596, Key: "Constricting Ring", Value: `Constricting Ring`},
	2597: {ID: 2597, Key: "Skillsd239", Value: `bite causes disease - werewolf form`},
	2598: {ID: 2598, Key: "Skillsd240", Value: `fiery, mauling attack`},
	2599: {ID: 2599, Key: "Entrapping", Value: `Entrapping`},
	2600: {ID: 2600, Key: "Blazing", Value: `Blazing`},
	2601: {ID: 2601, Key: "Skillsd241", Value: `release several small whirlwinds`},
	2602: {ID: 2602, Key: "Skillsd242", Value: `summon corpse eating vine`},
	2603: {ID: 2603, Key: "Skillsd243", Value: `life-and-mana-stealing bite`},
	2604: {ID: 2604, Key: "fng", Value: `Fang`},
	2605: {ID: 2605, Key: "The Atlantian", Value: `The Atlantean`},
	2606: {ID: 2606, Key: "Cutthroat", Value: `Bartuc's Chop Chop`},
	2607: {ID: 2607, Key: "Skillsd244", Value: `create shock waves - werebear form`},
	2608: {ID: 2608, Key: "A5Q3FoundAnyaCain", Value: `54
Goodness! Anya frozen by that fallen 
Barbarian, Nihlathak...Perhaps Malah 
can help you where I cannot.
`},
	2609: {ID: 2609, Key: "Bul-Kathos' Tribal Guardian", Value: `Bul-Kathos' Tribal Guardian`},
	2610: {ID: 2610, Key: "Skillsd245", Value: `create a volcano`},
	2611: {ID: 2611, Key: "Skillsd246", Value: `create a tornado`},
	2612: {ID: 2612, Key: "Skillsd247", Value: `summon a spirit pet of thorns`},
	2613: {ID: 2613, Key: "Charsi's Favor", Value: `Charsi's Favor`},
	2614: {ID: 2614, Key: "Eskillelementaldmg", Value: `Elemental Charge-up Damage: `},
	2615: {ID: 2615, Key: "Skillsd248", Value: `summon a grizzly bear`},
	2616: {ID: 2616, Key: "Skillsd249", Value: `multiple attacks - werewolf Form`},
	2617: {ID: 2617, Key: "Skillsd250", Value: `rain fire on your enemies`},
	2618: {ID: 2618, Key: "Ogun's Fierce Visage", Value: `Ogun's Fierce Visage`},
	2619: {ID: 2619, Key: "Ogun's Shadow", Value: `Ogun's Shadow`},
	2620: {ID: 2620, Key: "Sightless Veil", Value: `Sightless Veil`},
	2621: {ID: 2621, Key: "Skillsd251", Value: `create a massive wind storm`},
	2622: {ID: 2622, Key: "Maroon", Value: `Maroon`},
	2623: {ID: 2623, Key: "Skillsd252", Value: `throw a fire bomb`},
	2624: {ID: 2624, Key: "Skillsd253", Value: `passive - improves claw-class weapons ability`},
	2625: {ID: 2625, Key: "Envenomed", Value: `Foul`},
	2626: {ID: 2626, Key: "Of the River", Value: `of the River`},
	2627: {ID: 2627, Key: "Aladdin's Eviserator", Value: `Aladdin's Eviscerator`},
	2628: {ID: 2628, Key: "Banshee's Wail", Value: `Banshee's Wail`},
	2629: {ID: 2629, Key: "Stormlash", Value: `Stormlash`},
	2630: {ID: 2630, Key: "Skillsd254", Value: `a mind blast to crush your enemies`},
	2631: {ID: 2631, Key: "of Tornado", Value: `of Tornado`},
	2632: {ID: 2632, Key: "Shakabra's Crux", Value: `Shakabra's Crux`},
	2633: {ID: 2633, Key: "EnterCatacombsDru", Value: `The supernatural is strong here.`},
	2634: {ID: 2634, Key: "Skillsd255", Value: `increases damage of finishing moves

Charge-up Skill`},
	2635: {ID: 2635, Key: "Skillsd256", Value: `kick your enemies

Finishing Move`},
	2636: {ID: 2636, Key: "Skillsd257", Value: `throw a web of lightning`},
	2637: {ID: 2637, Key: "Skillan257", Value: `Shock Web`},
	2638: {ID: 2638, Key: "Skillname258", Value: `Blade Sentinel`},
	2639: {ID: 2639, Key: "Skillsd258", Value: `set a spinning blade`},
	2640: {ID: 2640, Key: "Skillld258", Value: `between you and target point
set a spinning blade to patrol`},
	2641: {ID: 2641, Key: "Skillan258", Value: `Blade Sentinel`},
	2642: {ID: 2642, Key: "Of Cremation", Value: `of Cremation`},
	2643: {ID: 2643, Key: "Skillname259", Value: `Burst of Speed`},
	2644: {ID: 2644, Key: "The Gladiator's Bane", Value: `The Gladiator's Bane`},
	2645: {ID: 2645, Key: "FindingJadeFigAss", Value: `Hmm, a jade statue. What should I do with it?`},
	2646: {ID: 2646, Key: "Skillsd259", Value: `increases attack and movement speed`},
	2647: {ID: 2647, Key: "Skillld259", Value: `for a period of time
increases attack and movement speed`},
	2648: {ID: 2648, Key: "Skillan259", Value: `Burst of Speed`},
	2649: {ID: 2649, Key: "Skillname260", Value: `Fists of Fire`},
	2650: {ID: 2650, Key: "Skillsd260", Value: `adds fire damage to finishing moves

Charge-up Skill`},
	2651: {ID: 2651, Key: "Warshrike", Value: `Warshrike`},
	2652: {ID: 2652, Key: "Skillld260", Value: `normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add fire damage

Charge-up Skill
`},
	2653: {ID: 2653, Key: "Skillan260", Value: `Fists of Fire`},
	2654: {ID: 2654, Key: "Skillname261", Value: `Dragon Claw`},
	2655: {ID: 2655, Key: "Skillsd261", Value: `double claw attack

Finishing Move`},
	2656: {ID: 2656, Key: "Skillld261", Value: `adds charged-up bonuses to both claw attacks
with your dual claw-class weapons
slice and dice your enemies

Finishing Move
`},
	2657: {ID: 2657, Key: "Skillan261", Value: `Dragon Claw`},
	2658: {ID: 2658, Key: "Skillname262", Value: `Charged Bolt Sentry`},
	2659: {ID: 2659, Key: "Skillsd262", Value: `a trap that emits charged bolts`},
	2660: {ID: 2660, Key: "Soulless", Value: `Soulless`},
	2661: {ID: 2661, Key: "Knave's", Value: `Knave's`},
	2662: {ID: 2662, Key: "Wallace's Tear", Value: `Wallace's Tear`},
	2663: {ID: 2663, Key: "Flaming", Value: `Flaming`},
	2664: {ID: 2664, Key: "Dragonscale", Value: `Dragonscale`},
	2665: {ID: 2665, Key: "Drakeflame", Value: `Drakeflame`},
	2666: {ID: 2666, Key: "Zakarum's Hand", Value: `Zakarum's Hand`},
	2667: {ID: 2667, Key: "Cow King's Horns", Value: `Cow King's Horns`},
	2668: {ID: 2668, Key: "The Atlantien", Value: `The Atlantean`},
	2669: {ID: 2669, Key: "Skillld262", Value: `at enemies that pass near
a trap that emits charged bolts`},
	2670: {ID: 2670, Key: "Skillan262", Value: `Charged Bolt`},
	2671: {ID: 2671, Key: "Skillname263", Value: `Wake of Fire`},
	2672: {ID: 2672, Key: "Skillsd263", Value: `a trap that emits waves of fire`},
	2673: {ID: 2673, Key: "Victorious", Value: `Victorious`},
	2674: {ID: 2674, Key: "of Frost Shield", Value: `of Frost Shield`},
	2675: {ID: 2675, Key: "Skillld263", Value: `a trap that emits waves of fire`},
	2676: {ID: 2676, Key: "Skillan263", Value: `Wake of Fire`},
	2677: {ID: 2677, Key: "Marrowgrinder", Value: `Marrow Grinder`},
	2678: {ID: 2678, Key: "of Teleportation", Value: `of Teleportation`},
	2679: {ID: 2679, Key: "Skillname264", Value: `Weapon Block`},
	2680: {ID: 2680, Key: "Eocene", Value: `Righteous`},
	2681: {ID: 2681, Key: "CompletingNihlathakAct5Sor", Value: `Your power was no match for mine.`},
	2682: {ID: 2682, Key: "Skillsd264", Value: `passive - block with two claw-class weapons`},
	2683: {ID: 2683, Key: "Skillld264", Value: `you are using dual claw-class weapons
passive - chance to block when`},
	2684: {ID: 2684, Key: "Skillan264", Value: `Wpn Block`},
	2685: {ID: 2685, Key: "Skillname265", Value: `Cloak of Shadows`},
	2686: {ID: 2686, Key: "Skillsd265", Value: `blind your enemies`},
	2687: {ID: 2687, Key: "Skillld265", Value: `lowering their defenses for a period of time
cast a shadow to blind nearby enemies`},
	2688: {ID: 2688, Key: "Skillan265", Value: `Cloak of Shdws`},
	2689: {ID: 2689, Key: "Skillname266", Value: `Cobra Strike`},
	2690: {ID: 2690, Key: "Skillsd266", Value: ` adds life and mana stealing to finishing moves

Charge-up Skill`},
	2691: {ID: 2691, Key: "Of Illusion", Value: `of Illusion`},
	2692: {ID: 2692, Key: "EnterForgottenTDru", Value: `I can smell why this tower was abandoned.`},
	2693: {ID: 2693, Key: "Skillld266", Value: `normal attack to release charges
must use a dragon finishing move or
to finishing moves
consecutive hits add life and mana stealing

Charge-up Skill
`},
	2694: {ID: 2694, Key: "Skillan266", Value: `Cobra Strike`},
	2695: {ID: 2695, Key: "Of Waste", Value: `of Waste`},
	2696: {ID: 2696, Key: "Skulltred", Value: `Skulltred`},
	2697: {ID: 2697, Key: "Skillname267", Value: `Blade Fury`},
	2698: {ID: 2698, Key: "Skillsd267", Value: `throw spinning blades`},
	2699: {ID: 2699, Key: "Skillld267", Value: `to slice up your enemies
throw spinning blades`},
	2700: {ID: 2700, Key: "A5Q4AfterInitMalah", Value: `65
My worst fear has come true. Nihlathak 
has gone mad.
 
You must stop him, before he gives the 
Relic to the Lord of Destruction!
`},
	2701: {ID: 2701, Key: "A5Q5AfterInitLarzuk", Value: `75
Every night, I've prayed to the Ancients 
to bring us peace...and now you must 
fight them.
 
Huh...Who shall I pray to now, warrior?
`},
	2702: {ID: 2702, Key: "A5Q1AfterInitNihlathak", Value: `70
After so many have died, who are you 
to think you can accomplish what our 
proud warriors could not?
`},
	2703: {ID: 2703, Key: "Skillan267", Value: `Blade Fury`},
	2704: {ID: 2704, Key: "A5Q1SuccessfulNihlathak", Value: `65
Ending the siege does not earn 
immediate respect, outsider.
 
Respect only comes with sacrifice -- 
something I'm sure you know nothing 
of.
`},
	2705: {ID: 2705, Key: "Skillname268", Value: `Fade`},
	2706: {ID: 2706, Key: "Skillsd268", Value: `temporary resist all`},
	2707: {ID: 2707, Key: "Skillld268", Value: `for a period of time
raise all resistances and resist curses`},
	2708: {ID: 2708, Key: "Skillan268", Value: `Fade`},
	2709: {ID: 2709, Key: "Skillname269", Value: `Shadow Warrior`},
	2710: {ID: 2710, Key: "Skillsd269", Value: `summon a shadow of yourself`},
	2711: {ID: 2711, Key: "Skillld269", Value: `your skills and fights by your side
summon a shadow of yourself that mimics`},
	2712: {ID: 2712, Key: "Skillan269", Value: `Shdw Warrior`},
	2713: {ID: 2713, Key: "Skillname270", Value: `Claws of Thunder`},
	2714: {ID: 2714, Key: "Skillsd270", Value: `adds lightning damage to finishing moves
 
Charge-up Skill`},
	2715: {ID: 2715, Key: "Eskillkickdamage", Value: `Kick Damage: `},
	2716: {ID: 2716, Key: "Skillld270", Value: `normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add lightning damage

Charge-up Skill
`},
	2717: {ID: 2717, Key: "Skillan270", Value: `Thunder Claws`},
	2718: {ID: 2718, Key: "Skillname271", Value: `Dragon Tail`},
	2719: {ID: 2719, Key: "Skillsd271", Value: `explosive kick

Finishing Move`},
	2720: {ID: 2720, Key: "A5Q2SuccessfulNihlathak", Value: `72
So. You brought the lost sheep home to 
the shepherd. Well done.
`},
	2721: {ID: 2721, Key: "Keeper's", Value: `Keeper's`},
	2722: {ID: 2722, Key: "Kira's Guardian", Value: `Kira's Guardian`},
	2723: {ID: 2723, Key: "Skillld271", Value: `adds charged-up bonuses to the kick
knock back your enemies with an explosive kick

Finishing Move
`},
	2724: {ID: 2724, Key: "Skillan271", Value: `Dragon Tail`},
	2725: {ID: 2725, Key: "Skillname272", Value: `Lightning Sentry`},
	2726: {ID: 2726, Key: "Skillsd272", Value: `a trap that emits lightning`},
	2727: {ID: 2727, Key: "Skillld272", Value: `to scorch passing enemies
a trap that shoots lightning`},
	2728: {ID: 2728, Key: "Skillan272", Value: `Lightning Sentry`},
	2729: {ID: 2729, Key: "Skillname273", Value: `Wake of Inferno`},
	2730: {ID: 2730, Key: "Skillsd273", Value: `trap that sprays fire`},
	2731: {ID: 2731, Key: "Skillld273", Value: `trap that sprays fire at passing enemies`},
	2732: {ID: 2732, Key: "Skillan273", Value: `Wake of Inferno`},
	2733: {ID: 2733, Key: "The Long Rod", Value: `The Longest Rod`},
	2734: {ID: 2734, Key: "CompletingTempleAss", Value: `The dark magic here is dispelled.`},
	2735: {ID: 2735, Key: "Skillname274", Value: `Mind Blast`},
	2736: {ID: 2736, Key: "Bul-Kathos' Warden", Value: `Bul-Kathos' Warden`},
	2737: {ID: 2737, Key: "Ethereal Edge", Value: `Ethereal Edge`},
	2738: {ID: 2738, Key: "A5Q2EarlyReturnNihlathak", Value: `52
Did you ever stop to think why these 
demons are capturing Qual-Kehk's 
men? Why they are attacking us? Have 
you considered what Baal wants with 
the mountain?
 
No. You've not. You have no idea what 
you are dealing with.
`},
	2739: {ID: 2739, Key: "Palo Grande", Value: `Palo Grande`},
	2740: {ID: 2740, Key: "Baals Minion", Value: `Minion of Destruction`},
	2741: {ID: 2741, Key: "CompletingStopSiegeAms", Value: `Oops...Did I do that?`},
	2742: {ID: 2742, Key: "Of Combat", Value: `of Combat`},
	2743: {ID: 2743, Key: "Skillsd274", Value: `compelling psionic blast`},
	2744: {ID: 2744, Key: "Aldur's Guantlet", Value: `Aldur's Gauntlet`},
	2745: {ID: 2745, Key: "Death's Web", Value: `Death's Web`},
	2746: {ID: 2746, Key: "Skillld274", Value: `and convert the feeble-minded
stun a group of enemies
using the power of your mind`},
	2747: {ID: 2747, Key: "Crown of Ages", Value: `Crown of Ages`},
	2748: {ID: 2748, Key: "Flowkrad's Paws", Value: `Flowkrad's Paws`},
	2749: {ID: 2749, Key: "A5Q2AfterInitLarzuk", Value: `73
I crafted the swords and armor for 
Qual-Kehk's men. They were like 
brothers to me. I can't imagine the pain 
they must be suffering.
 
Save them if you can!
`},
	2750: {ID: 2750, Key: "A5Q2InitQualKehk", Value: `58
My concerns have turned to my men 
taken prisoner on the battlefield by 
Baal's demons. I hate to think what's 
happened to them.
 
As you journey up the mountain, keep 
your eyes open for my soldiers and 
bring them back to me if you can.
`},
	2751: {ID: 2751, Key: "Skillan274", Value: `Mind Blast`},
	2752: {ID: 2752, Key: "Skillname275", Value: `Blades of Ice`},
	2753: {ID: 2753, Key: "Skillsd275", Value: `adds cold damage to finishing moves

Charge-up Skill`},
	2754: {ID: 2754, Key: "of Blood Golem Summoning", Value: `of Blood Golem`},
	2755: {ID: 2755, Key: "Gimmershred", Value: `Gimmershred`},
	2756: {ID: 2756, Key: "A5Q6SuccessfulQualKehk", Value: `60
The destruction of the Worldstone does 
not bode well for our world. But I'll try 
not to worry...
 
After all, we have warriors like you 
fighting for us and for the Light.
 
Farewell!
`},
	2757: {ID: 2757, Key: "Skillld275", Value: `normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add cold damage

Charge-up Skill
`},
	2758: {ID: 2758, Key: "Skillan275", Value: `Blades of Ice`},
	2759: {ID: 2759, Key: "The Impaler", Value: `The Impaler`},
	2760: {ID: 2760, Key: "Skillname276", Value: `Dragon Flight`},
	2761: {ID: 2761, Key: "FindingCairnAss", Value: `These stones hold an ancient power.`},
	2762: {ID: 2762, Key: "Precocious", Value: `Precocious`},
	2763: {ID: 2763, Key: "EnteringNihlathakAct5Ams", Value: `...Nihlathak's home away from home.`},
	2764: {ID: 2764, Key: "Skillsd276", Value: `teleport and kick enemies

Finishing Move`},
	2765: {ID: 2765, Key: "Skillld276", Value: `adds charged-up bonuses to the kick
and destroy them with a kick
teleport to your enemies

Finishing Move
`},
	2766: {ID: 2766, Key: "Skillan276", Value: `Dragon Flight`},
	2767: {ID: 2767, Key: "Skillname277", Value: `Death Sentry`},
	2768: {ID: 2768, Key: "Archangel's Deliverance", Value: `Arch-angel's Deliverance`},
	2769: {ID: 2769, Key: "Lycander's Aim", Value: `Lycander's Aim`},
	2770: {ID: 2770, Key: "Gutsiphon", Value: `Gut Siphon`},
	2771: {ID: 2771, Key: "Cow King's Hoofs", Value: `Cow King's Hooves`},
	2772: {ID: 2772, Key: "Chilling", Value: `Chilling`},
	2773: {ID: 2773, Key: "El Espiritu", Value: `El Espiritu`},
	2774: {ID: 2774, Key: "Skillsd277", Value: `trap that explodes nearby corpses`},
	2775: {ID: 2775, Key: "Excalibur", Value: `Excalibur`},
	2776: {ID: 2776, Key: "Skillld277", Value: `or explodes nearby corpses laying waste to more enemies
trap that shoots lightning at your enemies`},
	2777: {ID: 2777, Key: "Skillan277", Value: `Death Sentry`},
	2778: {ID: 2778, Key: "Deathcleaver", Value: `Death Cleaver`},
	2779: {ID: 2779, Key: "Skillname278", Value: `Blade Shield`},
	2780: {ID: 2780, Key: "The Reaper's Toll", Value: `The Reaper's Toll`},
	2781: {ID: 2781, Key: "Skillsd278", Value: `spinning blades of defense`},
	2782: {ID: 2782, Key: "Skillld278", Value: `who stray too close
spinning blades slice enemies`},
	2783: {ID: 2783, Key: "Skillan278", Value: `Blade Shield`},
	2784: {ID: 2784, Key: "Skillname279", Value: `Venom`},
	2785: {ID: 2785, Key: "Skillsd279", Value: `poison your weapon`},
	2786: {ID: 2786, Key: "Skillld279", Value: `add poison damage to your weapons`},
	2787: {ID: 2787, Key: "Radimant's Sphere", Value: `Radament's Sphere`},
	2788: {ID: 2788, Key: "Skillan279", Value: `Venom`},
	2789: {ID: 2789, Key: "Skillname280", Value: `Shadow Master`},
	2790: {ID: 2790, Key: "Skillsd280", Value: `summon your shadow`},
	2791: {ID: 2791, Key: "Corosive", Value: `Corrosive`},
	2792: {ID: 2792, Key: "Skillld280", Value: `to fight by your side
summon a powerful shadow of yourself`},
	2793: {ID: 2793, Key: "Skillan280", Value: `Shdw Master`},
	2794: {ID: 2794, Key: "Skillname281", Value: `Phoenix Strike`},
	2795: {ID: 2795, Key: "Skillsd281", Value: `adds elemental novas to finishing moves

Charge-up Skill`},
	2796: {ID: 2796, Key: "Skillld281", Value: `normal attack to release charges
must use a dragon finishing move or
adds elemental novas to finishing moves

Charge-up Skill
`},
	2797: {ID: 2797, Key: "Earthshaker", Value: `Earthshaker`},
	2798: {ID: 2798, Key: "Felicitous", Value: `Felicitous`},
	2799: {ID: 2799, Key: "Skillan281", Value: `Phnx Strike`},
	2800: {ID: 2800, Key: "ESkillHawk", Value: `Ravens: `},
	2801: {ID: 2801, Key: "ESkillSpikes", Value: `Spikes: `},
	2802: {ID: 2802, Key: "ESkillWolves", Value: `Wolves: `},
	2803: {ID: 2803, Key: "The Vicar", Value: `The Vicar`},
	2804: {ID: 2804, Key: "ESkillShoots", Value: `Shoots `},
	2805: {ID: 2805, Key: "ESkillSpikes2", Value: ` Spikes`},
	2806: {ID: 2806, Key: "Akarat's Protector", Value: `Akarat's Protector`},
	2807: {ID: 2807, Key: "EskillLifeSteal", Value: `Life Steal: `},
	2808: {ID: 2808, Key: "A5Q1SuccessfulCain", Value: `60
Those catapults were like nothing I have 
ever seen before. You have prevailed 
against Shenk, my friend, but Baal is 
still far ahead of you.
`},
	2809: {ID: 2809, Key: "Crest of Morn", Value: `Crest of Morn`},
	2810: {ID: 2810, Key: "of Passion", Value: `of Passion`},
	2811: {ID: 2811, Key: "Celestial", Value: `Celestial`},
	2812: {ID: 2812, Key: "CompletingBladeAss", Value: `What a delicious blade! I should consult Ormus.`},
	2813: {ID: 2813, Key: "Eskillchancetostun", Value: `Chance to stun: `},
	2814: {ID: 2814, Key: "Geronimo's Fury", Value: `Geronimo's Fury`},
	2815: {ID: 2815, Key: "Eskillpowerup1", Value: `Charge 1 - `},
	2816: {ID: 2816, Key: "Eskillpowerup2", Value: `Charge 2 - `},
	2817: {ID: 2817, Key: "Ember", Value: `Fiery`},
	2818: {ID: 2818, Key: "Heaven's Light", Value: `Heaven's Light`},
}

var expansionstring_TblDataByValue = map[string][]int{
	``:                  []int{0, 8, 810}[1:],
	` Charges of`:       []int{0, 2149}[1:],
	` Elemental Damage`: []int{0, 1477}[1:],
	` Kick`:             []int{0, 1659}[1:],
	` Kicks`:            []int{0, 754}[1:],
	` Percent Attack`:   []int{0, 11}[1:],
	` Percent Damage`:   []int{0, 12}[1:],
	` Percent Life`:     []int{0, 986}[1:],
	` Sockets`:          []int{0, 2097}[1:],
	` Spikes`:           []int{0, 2805}[1:],
	` Times`:            []int{0, 419}[1:],
	` adds life and mana stealing to finishing moves

Charge-up Skill`: []int{0, 2690}[1:],
	` hit`:                                        []int{0, 1251}[1:],
	` per blade`:                                  []int{0, 1740}[1:],
	` per hit`:                                    []int{0, 2021}[1:],
	` per kick`:                                   []int{0, 1109}[1:],
	` percent damage`:                             []int{0, 7}[1:],
	` percent life and mana stealing `:            []int{0, 806}[1:],
	` percent life stealing`:                      []int{0, 6}[1:],
	` percent to Attack Rating`:                   []int{0, 2111}[1:],
	` to Attack Rating vs. Demons`:                []int{0, 2098}[1:],
	` to Attack Rating vs. Undead`:                []int{0, 2099}[1:],
	` to Bow and Crossbow Skills`:                 []int{0, 2114}[1:],
	` to Cold Skills`:                             []int{0, 2082}[1:],
	` to Combat Skills`:                           []int{0, 2080, 2120}[1:],
	` to Curses`:                                  []int{0, 2078}[1:],
	` to Damage vs. Demons`:                       []int{0, 2100}[1:],
	` to Defensive Aura Skills`:                   []int{0, 2118}[1:],
	` to Elemental Skills`:                        []int{0, 2128}[1:],
	` to Fire Skills`:                             []int{0, 2123}[1:],
	` to Javelin and Spear Skills`:                []int{0, 2112}[1:],
	` to Lightning Skills`:                        []int{0, 2094}[1:],
	` to Martial Art Skills`:                      []int{0, 2137}[1:],
	` to Masteries Skills`:                        []int{0, 2081}[1:],
	` to Offensive Aura Skills`:                   []int{0, 2119}[1:],
	` to Passive and Magic Skills`:                []int{0, 2113}[1:],
	` to Poison and Bone Skills`:                  []int{0, 2122}[1:],
	` to Shadow Discipline Skills`:                []int{0, 2130}[1:],
	` to Shape-Shifting Skills`:                   []int{0, 2126}[1:],
	` to Summoning Skills`:                        []int{0, 2121, 2124}[1:],
	` to Trap Skills`:                             []int{0, 2129}[1:],
	` to Warcry Skills`:                           []int{0, 2079}[1:],
	` to melee attacks`:                           []int{0, 741}[1:],
	`%0 %1`:                                       []int{0, 913}[1:],
	`%d Percent Chance of Critical Strike`:        []int{0, 2512}[1:],
	`%d%% Chance to cast level %d %s on attack`:   []int{0, 792}[1:],
	`%d%% Chance to cast level %d %s on striking`: []int{0, 557}[1:],
	`%d%% Chance to cast level %d %s when struck`: []int{0, 558}[1:],
	`'`:                                   []int{0, 2147}[1:],
	`(`:                                   []int{0, 724}[1:],
	`(%d/%d Charges)`:                     []int{0, 723}[1:],
	`(Amazon Only)`:                       []int{0, 1206, 1297}[1:],
	`(Assassin Only)`:                     []int{0, 634, 1959}[1:],
	`(Barbarian Only)`:                    []int{0, 300, 1444}[1:],
	`(Based on Character Level)`:          []int{0, 2138}[1:],
	`(Druid Only)`:                        []int{0, 299, 1852}[1:],
	`(Increases During Daytime)`:          []int{0, 2143}[1:],
	`(Increases During Nighttime)`:        []int{0, 2140}[1:],
	`(Increases Near Dawn)`:               []int{0, 2146}[1:],
	`(Increases Near Dusk)`:               []int{0, 2148}[1:],
	`(Necromancer Only)`:                  []int{0, 2101}[1:],
	`(Paladin Only)`:                      []int{0, 179}[1:],
	`(Sorceress Only)`:                    []int{0, 1487}[1:],
	`)`:                                   []int{0, 725}[1:],
	`...Nihlathak's home away from home.`: []int{0, 2763}[1:],
	`105
I understand your reluctance, but now 
is the time to strike.
`: []int{0, 730}[1:],
	`107
I'm sure those men appreciate the 
freedom you gave them...as much as I 
do.
`: []int{0, 1449}[1:],
	`107
On the battlefield, turning your back on 
even the dead is unwise.
`: []int{0, 2231}[1:],
	`108
I was starting to wonder how long it 
would take you to stop those beasts.
 
Good job.
`: []int{0, 1705}[1:],
	`114
I warned you!
 
The Ancients are not like the demons 
you're accustomed to fighting.
`: []int{0, 2203}[1:],
	`115
Well, you found me on the mountain; I'm 
sure you'll find them too.
`: []int{0, 434}[1:],
	`133
If you have nothing useful for me, be on 
your way!
`: []int{0, 945}[1:],
	`140
Your rescue of Anya was quite an 
accomplishment.
`: []int{0, 1410}[1:],
	`150
Uh...Did I mention there were catapults?
`: []int{0, 1472}[1:],
	`160
...What's there to talk about?
 
Kill Nihlathak!
`: []int{0, 2268}[1:],
	`31
You are a worthy hero! We augment 
your skill and grant you entry to the 
interior of Mount Arreat, wherein lies 
the Worldstone.
 
Beware. You will not be alone. Baal the 
Lord of Destruction is already inside. 
 
The Archangel Tyrael has always been 
our benefactor, but even he cannot 
help us now. For Baal blocks Tyrael's 
spiritual presence from entering the 
chamber of the Worldstone. Only you, 
mortal, have the power to defeat Baal 
now.
 
Baal threatens the Worldstone -- and 
through it, the mortal realm, itself. You 
must stop him before he gains full 
control of the sacred stone. With it 
under his control, Baal could shatter 
the boundaries between this world and 
the Burning Hells, thus allowing the 
hordes of the Prime Evils to pour forth 
into the mortal realm like an 
unstoppable tide!
 
If you are weak, the world as you know 
it could be lost forever. You must NOT 
fail!
`: []int{0, 889}[1:],
	`34
You have traveled far only to return 
home to us, Barbarian. Ohh...Much has 
happened in Harrogath since you left. 
Our homeland is hardly recognizable 
with so much evil about.
 
Yet, I've managed to survive so far. 
You've seen your share of suffering as 
well, I'm sure. Seeing you again -- alive 
-- does my heart good. I shall pray that 
you can help us out of this hell.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what.
 
Nihlathak is the last survivor of the 
Council of Elders, but I'm afraid he has 
not been himself lately. The other 
Elders sacrificed themselves to place a 
protective ward around Harrogath. If 
Nihlathak is curt with you, it is because 
he feels responsible for our situation. 
He does not relish sending more of our 
people out to die.
 
So much has changed since you left 
that I see little hope for us in this life.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`: []int{0, 2505}[1:],
	`35
I, Malah, welcome you to Harrogath, 
the last stronghold of Order on Mount 
Arreat. You have come to the right 
place if you intend to defeat Baal the 
Lord of Destruction.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what. 
 
All of the Elders, save Nihlathak, 
sacrificed themselves to place a 
protective ward around Harrogath.
 
Some of us here, certainly Nihlathak, do 
not appreciate your presence. We are a 
proud people, and it is not easy for us 
to accept aid. I, however, am glad you 
are here.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`: []int{0, 901}[1:],
	`35
We are the spirits of the Nephalem, the 
Ancient Ones. We have been chosen to 
guard sacred Mount Arreat, wherein 
the Worldstone rests. Few are worthy 
to stand in its presence; fewer still can 
comprehend its true purpose.
 
Before you enter, you must defeat us.
`: []int{0, 1992}[1:],
	`36
A Sorceress...Here in Harrogath?
 
There was a time, child, when I thought 
I was destined to follow your kind's 
path. However, my powers never 
developed beyond the simplest of 
spells. Although I can heal almost any 
wound with time and energy, there is 
little I can do to help against Baal.
 
But enough of that...I spend too much 
time in reflection and have forgotten 
my manners.
 
I, Malah, welcome you to Harrogath, 
the last stronghold of Order on Mount 
Arreat. You have come to the right 
place, if you intend to defeat Baal the 
Lord of Destruction.
 
Baal has laid waste to our mountain 
and its denizens. His minions continue 
to attack our town, while Qual-Kehk 
and his men have proven helpless to 
stop them. Baal is still out on the 
mountain looking for something -- but I 
know not what. 
 
All of the Elders, save Nihlathak, 
sacrificed themselves to place a 
protective ward around Harrogath.
 
Some of us here, certainly Nihlathak, do 
not appreciate your presence. We are a 
proud people, and it is not easy for us 
to accept aid. I, however, am glad you 
are here.
 
If you need healing or a potion, please 
come to me. See Larzuk for weapons, 
armor, and repairs. Nihlathak, despite 
his disposition, may be of some 
assistance with other wares. Finally, 
Qual-Kehk, our Man-At-Arms, leads 
Harrogath's remaining forces against 
Baal.
`: []int{0, 2594}[1:],
	`37
Every time I hear of you, warrior, your 
deeds become more legendary.
 
But take heed. You are approaching the 
very summit of Mount Arreat. I have 
never dared venture there. It is sacred 
-- our most holy place. The legends say 
it is guarded by the Ancient Ones, who 
block the path of all who are unworthy.
 
Your reputation here does not 
matter...It will be the Ancients who 
determine your worthiness.
 
Good luck.
`: []int{0, 2028}[1:],
	`37
Oh yes...I remember our warriors as 
children. Malah would set their broken 
bones and give them powders for their 
fevers. Now, they return to her with 
wounds that will never heal.
 
I am tired...Please...leave me.
`: []int{0, 944}[1:],
	`39
I am amazed to find this place so 
untouched. Everything else in the path 
of Baal the Lord of Destruction lies in 
ruin. 
 
These Barbarians must indeed be the 
legendary guardians of Mount Arreat. 
They are a proud, hardy people. Don't 
expect to be greeted warmly -- 
strangers here rarely are.
 
Perhaps I can gain their trust. I'll spend 
some time with the townsfolk and try 
to understand them better. I'll let you 
know what I discover.
`: []int{0, 1285}[1:],
	`40
I am impressed, mortal. You have 
overcome the greatest challenge this 
world has ever faced and defeated the 
last of the Prime Evils. However, we are 
too late to save the Worldstone. Baal's 
destructive touch has corrupted it 
completely.
 
Given enough time, the Worldstone's 
energies will drain away and the 
barriers between the worlds will 
shatter -- the powers of Hell will flood 
into this...Sanctuary...and eradicate 
your people and everything you've 
labored to build.
 
Therefore, I must destroy the corrupted 
Worldstone before the powers of Hell 
take root. This act will change your 
world forever -- with consequences 
even I cannot foresee. However, it is 
the only way to ensure mankind's 
survival.
 
Go now, mortal. I have opened a portal 
that will lead you to safety.
 
May the Eternal Light shine upon you 
and your descendants for what you've 
done this day. The continued survival 
of mankind is your legacy! Above all 
else, you have earned a rest from this 
endless battle.
`: []int{0, 1642}[1:],
	`40
I knew in time you would defeat Baal. 
You have done everything you set out 
to do, my friend.
 
Ever since you rescued me from 
Tristram, I have believed in you. It has 
been a supreme honor to aid you along 
the way. 
 
So...The Worldstone was corrupted by 
Baal. And now Tyrael must destroy it. 
Worry not. Through whatever lies 
ahead I have faith that the Light will 
guide us both.
 
Go, now, back to the Worldstone 
chamber, and enter the portal Tyrael 
has opened for you.
`: []int{0, 1528}[1:],
	`40
I knew there was great potential in you, 
my friend. You've done a fantastic job.
 
Though my ancestors often struggled 
against the Three Evils and their 
minions, I've always lived a shut-in, 
scholarly life. I'm glad that my wisdom 
aided you.
 
Now, I wish to leave this place. Though 
Heaven's Gates are a marvel to behold, 
I hope I won't have to see them again 
for many, many years.
 
Please talk to Tyrael about leaving this 
place now!
`: []int{0, 104}[1:],
	`40
Praise be to the Light! You have 
accomplished the impossible!
 
Diablo and Mephisto have been 
banished back into the Black Abyss 
that spawned them, and the corrupted 
Soulstones are no more.
 
However, while you were fighting here, 
Baal remained behind in the mortal 
realm, building an army of hellish 
minions. Now, Baal's army is searching 
for the Worldstone, the ancient source 
of all the Soulstones and their power, 
while leaving behind a wake of 
destruction. They have forged deeply 
into the Barbarian homelands, heading 
directly for the summit of Mount 
Arreat!
 
Baal knows, mortal hero! That is the 
very site of the blessed Worldstone!
 
Now, enter the portal I have opened for 
you. It will take you to the Barbarian 
city of Harrogath, the last bastion of 
Order on the slopes of Arreat.
`: []int{0, 753}[1:],
	`41
All users of the magical arts know of 
Mount Arreat, but few understand its 
true nature. It is the nexus of an 
unfathomable magic.
 
It bodes ill that the Lord of Destruction 
races to its summit with such purpose. 
I fear for the whole world should Baal 
gain what he seeks.
`: []int{0, 1948}[1:],
	`41
Anya! Who have you been talking to? 
Likely it was that meddling Malah. 
 
Well, I'll tell you what really happened. 
Anya came to me for guidance, after 
receiving a vision that her mother and 
younger brother were trapped in the 
lands beyond the Ice Caves. She had 
decided to go rescue them. 
 
I told her that her quest was a foolish 
one and that she would be safer 
staying within the city walls. However, 
she is a willful girl and would not listen 
to me. 
 
The next morning, she was gone. No 
one is more distraught than I over 
losing her. 
 
However, if you feel the need to be 
Malah's errand child, I won't try to stop 
you.
`: []int{0, 2577}[1:],
	`42
A Druid in Harrogath! Have things truly 
come to this?
 
After the Mage Wars, I assumed Druids 
would never be seen in Harrogath 
again. You take a big chance coming 
here! 
 
To be honest, I have never been 
comfortable with your shape-shifting 
kind, but I do respect your search for 
balance and peace. So, if you trust us 
enough to enter our gates, I trust you 
enough to let you stay.
 
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`: []int{0, 1369}[1:],
	`42
A Paladin! I have long heard of your 
people.
 
As a young warrior I even considered 
the pilgrimage to Kurast. But I was 
younger then and foolish. My place has 
always been here -- protecting 
Harrogath, and Mount Arreat with it.
 
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`: []int{0, 161}[1:],
	`42
I know you and my son, Bannuk, did not 
part on the best of terms. He did not 
understand the wanderlust that drove 
you to leave your homeland. However, 
even though Bannuk could never admit 
it to me, as he grew older I could see 
that same desire in his eyes.
 
Oh...It's a pity I didn't encourage him to 
go with you. He might still be alive 
today.
`: []int{0, 2229}[1:],
	`42
So, you're an Amazon. I have heard 
rumors about your kind. Your unusual 
weapons could prove a challenge to my 
skills, but I'm prepared to meet it.
 
I am Larzuk, the armorer. My ancestors 
were some of the finest craftsmen in 
Harrogath. Regretfully, my supplies run 
lower with every passing day, yet the 
demons beyond the walls have not 
weakened. I fear the time is near when 
I must put down my hammer and take 
up a sword, instead.
`: []int{0, 2405}[1:],
	`43
Anya's father, Aust, was our wisest 
Elder. He was killed along with the 
other Elders who erected the ward that 
protects this city. The ward has kept 
Baal's demons out of Harrogath, but at 
a costly sacrifice. 
 
Nihlathak, on the other hand, was the 
only Elder to escape the demons. 
Somehow, he alone managed to find 
sanctuary, while the others died 
around him.
 
Ever since that day, Nihlathak and Anya 
have been at odds.
`: []int{0, 2166}[1:],
	`43
Nihlathak told me he struck a deal with 
Baal to protect Harrogath. In exchange 
for the demon's mercy, the misguided 
fool plans to give Baal the Relic of the 
Ancients, our most holy totem!
 
Doing so will allow Baal to enter Mount 
Arreat unchallenged by the Ancients. I 
tried to stop Nihlathak, but he 
imprisoned me in that icy tomb.
 
Nihlathak must be stopped before he 
dooms the whole world. As much as I 
would love to strangle the life out of 
him, I'm afraid I haven't the strength.
 
You must go to his lair through this 
portal I've opened, kill him, and then 
bring back the Relic of the Ancients.
 
Stop Nihlathak from destroying what we 
have striven for eons to protect.
`: []int{0, 2341}[1:],
	`43
There is a matter which I hesitate to 
share, but I believe you are the only 
one who can help me now.
 
Anya, the young alchemist and 
daughter to one of our slain Elders, 
has been missing for some time. She is 
a strong, crafty woman with a spirit 
like no other.
 
One night, just before your arrival, I 
overheard her and Nihlathak arguing 
about her father's death. The next 
morning she was gone.
 
Nihlathak has his own tale as to where 
she went and why. Don't believe him! I 
fear he is at the root of her 
disappearance.
 
Please, if you can, search for Anya and 
bring her back to us. She'll know what 
to do about Nihlathak.
`: []int{0, 45}[1:],
	`43
Well, well. The siege has everything in 
short supply...except fools. 
 
Why would you seek this place, 
stranger? Are you a vulture come to 
loot the bodies of our fallen warriors? 
 
Regardless, this is no place to make a 
name for yourself. The mountain is 
ours to protect. It is only a matter of 
time before Hell's legions are routed.
`: []int{0, 279}[1:],
	`44
During my time with the Horadrim, we 
often debated the nature of Mount 
Arreat.
 
We knew that the Barbarian clans 
zealously guarded the mountain as 
their sacred duty. However, many 
dismissed their zeal as simple 
superstition...combined with an inborn 
hostility toward outsiders.
 
Those Horadrim who trekked up Arreat 
were never heard from again...Still, I 
do not believe they died at the hands of 
Barbarians.
`: []int{0, 1947}[1:],
	`44
The death of Malah's son was a great 
tragedy. He was our finest archer.
 
While leading a successful campaign 
against Baal's forces, he was impaled 
on a demon's spear. The wound was 
such that...Well, even Malah herself 
acknowledges that quick death was a 
blessing.
`: []int{0, 1623}[1:],
	`44
With the exception of Qual-Kehk, the 
townspeople do not see what I see -- 
the massacre we face.
 
Our bravest, strongest heroes hobble 
back to me begging for help. I do what 
I can -- healing and bandaging some, 
preparing the others for what lies 
beyond.
`: []int{0, 2558}[1:],
	`45
Ahhh, a Necromancer.
 
While I admire your courage in seeking 
out the darker side of magic, we really 
have little need of your skills. The 
battle will turn soon enough without 
your meddling. 
 
Yet, I should have expected to see your 
kind here. You are like a moth to the 
flame -- drawn to all this death. It 
feeds you in more ways than one, does 
it not?
`: []int{0, 137}[1:],
	`45
I am Qual-Kehk, the Senior Man-At-Arms 
of Harrogath.
 
You have the look of a warrior...An 
extra soldier will be useful. But don't 
expect anyone to mourn if you get 
yourself killed. 
 
Baal is true to his namesake. He has 
ravaged through our lands like a 
merciless plague.
 
The protective ward laid down by our 
lost Elders helps hold the evil at bay, 
but Baal's siege has taken its toll all 
the same.
 
Most of my men are now dead. Others 
are trapped in the mountain passes.
 
But I swear we are not beaten yet! We 
will fight to the end to protect this 
mountain!
`: []int{0, 936}[1:],
	`45
If you're here to defeat Baal, you must 
prove it!
 
As we speak, Harrogath is under siege 
by Baal's demons. Catapults rain death 
just outside the town walls.
 
Baal himself travels up the sacred 
mountain, having left in charge here 
one of his most vicious generals, Shenk 
the Overseer. A ruthless taskmaster, 
he lashes his own minions into suicidal 
frenzies on the battlefield.
 
If you wish to prove yourself to us, 
destroy the monster, Shenk, that 
commands those infernal catapults 
outside Harrogath.  If you manage to 
do this, return to me.
`: []int{0, 926}[1:],
	`45
Now that the Elders are all dead, I don't 
know who will guide our people through 
this dark time. I was to be next in line 
after my father, but this burden is too 
great for me to shoulder alone. 
 
We are a people of strong blood. I shall 
do what I can and let fate do the rest.
`: []int{0, 1145}[1:],
	`46
Look, I've told you! She's dead! If you 
knew what was good for you, you'd 
concentrate your efforts on saving 
Harrogath -- not on lost causes like 
Anya.
`: []int{0, 176}[1:],
	`46
Our Elders were wise leaders in more 
peaceful times, but now the survival of 
our people has fallen to me. My men 
and I will fight to the death, but there's 
no way to ensure the outcome.
 
I used to believe that nothing could 
break through our guard and assault 
the holy mountain. I know now that I 
was horribly mistaken.
`: []int{0, 1629}[1:],
	`48
If Tyrael says the Worldstone must be 
destroyed, then it must. We cannot let 
Baal's corruption prevail!
 
The world will change, true -- but who is 
to say it isn't for the better?
`: []int{0, 663}[1:],
	`48
Larzuk is a talented blacksmith, but his 
head is full of some strange ideas. 
 
Just the other day he came to me with a 
plan to break the siege. He wanted to 
fill large pipes with exploding powders 
and steel balls, then...Well, like I said, 
strange.
`: []int{0, 1628}[1:],
	`48
Thank you so much for bringing Anya 
back to us. I have devised this spell to 
increase your resistances as a token of 
my thanks. I know it isn't much, but I 
hope you find it helpful.
 
Please go talk to Anya. She has urgent 
news concerning Nihlathak.
`: []int{0, 2204}[1:],
	`48
Why would you seek this place, 
stranger? Are you a vulture come to 
loot the bodies of our fallen warriors? 
 
Regardless, this is no place to make a 
name for yourself. The mountain is 
ours to protect. It is only a matter of 
time before Hell's legions are routed.
`: []int{0, 941}[1:],
	`49
If you're looking for cases of 
treacherous magic, Assassin, take a 
hard look at Larzuk. He was the only 
one in town who escaped the Red Fever 
last spring. He claimed his good 
fortune was due to 'hand washing' 
before meals.
 
Hmmm...Very suspicious...
`: []int{0, 940}[1:],
	`50
Hero. Nihlathak did this to me!
 
If you've come to help me, my only hope 
lies with Malah.
 
Please...Tell her you've found me...
`: []int{0, 277}[1:],
	`50
Qual-Kehk's confidence in his abilities 
once bordered on arrogance, but in the 
early days of the siege, he was 
humbled by Baal. While he saved us 
from immediate destruction, a third of 
our warriors were lost.
 
None felt those losses more than he. 
Though he may not admit it, your 
presence gives him hope, warrior.
`: []int{0, 1151}[1:],
	`50
Regretfully, I know very little about this 
Relic. However, if what the others say 
is true, then Baal must not gain 
possession of it.
 
Stop Nihlathak...before all is lost.
`: []int{0, 1126}[1:],
	`51
Remember this. Baal once possessed Tal 
Rasha, one of the most powerful of the 
ancient Horadrim.
 
Your battles with Mephisto and Diablo 
will pale in comparison to your battle 
with Baal.
 
The Lord of Destruction aided by Tal 
Rasha's knowledge...The mountain 
itself will tremble when you clash.
`: []int{0, 1212}[1:],
	`51
Thank you for rescuing my men. They 
have spoken well of your bravery in 
battle. Perhaps there is hope for us 
after all. 
 
If you wish, you may hire some of my 
mercenaries that you saved. And 
please...take this set of runes. I had 
been saving them for a socketed shield, 
but I think you'll make better use of 
them.
 
Be sure to set them in the right order 
for their fullest effect.
`: []int{0, 188}[1:],
	`52
Did you ever stop to think why these 
demons are capturing Qual-Kehk's 
men? Why they are attacking us? Have 
you considered what Baal wants with 
the mountain?
 
No. You've not. You have no idea what 
you are dealing with.
`: []int{0, 2738}[1:],
	`52
If it were anyone else, I would assume 
her dead. However, Anya is not so 
easily beaten. I know she must be alive.
`: []int{0, 1639}[1:],
	`52
Larzuk possesses a good soul, but at 
times his mind seems quite unsound. 
 
He once asked me for twenty of my 
finest sheepskins. He said he would fill 
them with hot air and float like a cloud 
above the battlefield to spy on Baal's 
legions! 
 
I worry the siege has driven him mad.
`: []int{0, 2236}[1:],
	`52
We have come too far to be defeated 
now, my friend. I have seen you 
complete many difficult quests. Though 
this may be your greatest trial, it is not 
beyond your reach.
`: []int{0, 200}[1:],
	`52
With hellspawn, size is no measure of 
their threat. Demons half the size of 
men can kill with a gesture, while 
hellish pack animals trample any who 
stand in their way.
`: []int{0, 1944}[1:],
	`53
Ah, Anya. Such a fine example of 
feminine strength...
 
She reminds me of the Zakarum 
priestesses I knew in my youth. They 
don't take vows of chastity, you know.
`: []int{0, 1951}[1:],
	`53
I have long been criticized, but 
especially of late -- since the deaths of 
my fellow Elders. Through it all, I have 
learned one thing. Each man must do 
what's right, no matter what others 
may think.
`: []int{0, 947}[1:],
	`53
I pray for the day when the fields are 
covered in snow unstained by the blood 
of our own. Perhaps that day will come 
soon...Perhaps never...
 
Oh...But I forget myself. How may I aid 
your quest?
`: []int{0, 2556}[1:],
	`53
I'm aware of the amazing magical 
powers a Sorceress can channel. If we 
survive Baal's wrath, I would be most 
honored if you could demonstrate and 
perhaps teach me something of what 
you know. 
 
I may be old, but I'm not dead.
`: []int{0, 2228}[1:],
	`53
Perhaps you have heard the accounts 
of my son's horrible death at the hands 
of Baal's minions. He was my last living 
child...
 
The oath of compassion I have taken as 
a healer extends only to humankind.
 
Cut them down, warrior. All of them!
`: []int{0, 2233}[1:],
	`53
You have done the impossible, hero. 
Your defeat of the last of the three 
Prime Evils is a great victory for the 
Light. 
 
Strange that you say that the 
Worldstone must be destroyed. The 
prophecies said nothing about that.
 
Perhaps all we have fought for will be 
lost...or perhaps we'll never need fight 
again!
`: []int{0, 426}[1:],
	`54
Every day, one of my friends dies 
fighting outside the town walls, while I 
tend my anvil here unscathed. If only 
we didn't need weapons so badly, I 
could be out doing my share of the 
fighting. 
 
Good luck to you, warrior.
`: []int{0, 1882}[1:],
	`54
Goodness! Anya frozen by that fallen 
Barbarian, Nihlathak...Perhaps Malah 
can help you where I cannot.
`: []int{0, 2608}[1:],
	`54
It is my belief that the Soulstones are at 
the center of this conflict. If only that 
fool Marius had not intervened, Baal 
would still be imprisoned within Tal 
Rasha.
`: []int{0, 145}[1:],
	`54
My father, Aust, was among those 
Elders wise enough to see that we 
would need outside help to deal with 
Baal's legions. He believed that this 
conflict would affect the entire world, 
not just our homeland. He said that 
Mount Arreat is as necessary to the 
world's survival as food and water is to 
our own.
 
I believe this to be true.
`: []int{0, 1150}[1:],
	`54
Though the Elder Council of Harrogath 
is gone, there are many capable young 
leaders to take their place.
 
Anya certainly has enough courage and 
intelligence to lead them all, if they can 
survive this catastrophe.
`: []int{0, 1950}[1:],
	`54
Well, well...An Assassin!
 
Heh, heh...While I am sure we are all 
grateful for your presence in our 
troubled town, you need not have made 
the journey.
 
I can personally say that your skills are 
not required here. You would serve 
your clan better elsewhere.
`: []int{0, 795}[1:],
	`54
You have proven yourself a true hero to 
me and my people.
 
These are dark times, warrior. I hope 
you can bring an end to Baal's reign of 
destruction. 
 
Our Council of Elders is gone -- my 
father, Aust, among them. The one 
thing that keeps us from total despair 
is the promise of vengeance against 
Baal.
`: []int{0, 1821}[1:],
	`55
I am Larzuk, the armorer. My ancestors 
were some of the finest craftsmen in 
Harrogath. 
 
Regretfully, my supplies run lower with 
every passing day, yet the demons 
beyond the walls have not weakened. 
 
I fear the time is near when I must put 
down my hammer and take up a sword, 
instead.
`: []int{0, 1990}[1:],
	`55
Nihlathak is a traitor! If Baal gets the 
Relic, he shall enter the mountain and 
wreak havoc there!
 
I cannot believe that Nihlathak would 
give the Relic to Baal in any kind of 
trade. He's truly insane if he believes 
that he can deal with the Lord of 
Destruction.
`: []int{0, 699}[1:],
	`55
Since your arrival, Cain has spoken of 
your deeds in the Southern Kingdoms, 
defeating both Mephisto and Diablo. At 
first, I scoffed at his tales, but I'm 
finding them more believable with every 
passing day.
`: []int{0, 430}[1:],
	`55
The angel Tyrael has watched over the 
guardians of Arreat throughout 
history. I do not believe that Baal and 
Tyrael have come to fight over a paltry 
few souls.
 
They are here to settle a conflict as old 
as time itself.
`: []int{0, 1946}[1:],
	`55
Though these Barbarians are known 
throughout the kingdoms as ferocious 
fighters, they are also capable of great 
compassion.
 
They have trained throughout history 
for a battle their legends foretell will 
decide the fate of the world.
`: []int{0, 1945}[1:],
	`55
You have stopped Nihlathak, but he 
didn't have the Relic! He must have 
already given it to Baal. Now, Baal will 
not be tested when he reaches Arreat's 
summit.
 
...Damn Nihlathak!
 
I do thank you for trying, though. 
Please, allow me to honor your courage 
by magically inscribing your name onto 
an item of your choosing. It's the least 
I can do.
`: []int{0, 938}[1:],
	`56
As the daughter of Elder Aust, Anya is 
the only person, besides Nihlathak, 
who has any real knowledge of Mount 
Arreat's secrets. I'd hate to see our 
fate in the hands of Nihlathak alone.
`: []int{0, 1257}[1:],
	`56
Just so you know...The gold you pay me 
doesn't line my pockets. Much of it 
goes to buy the raw metal I need to 
melt down for weapons and armor. The 
rest -- well, all I can spare -- goes to 
Malah and Qual-Kehk for other 
supplies.
`: []int{0, 1842}[1:],
	`56
Now, rescuing Anya -- that's good. 
Talking to me while Nihlathak hands 
over the Relic to Baal -- uh...that's bad!
`: []int{0, 196}[1:],
	`56
So! That snake Nihlathak was behind 
Anya's disappearance...and he trapped 
her with a freezing curse.
 
Here. Take this potion to Anya and give 
it to her. That should release her.
`: []int{0, 1573}[1:],
	`57
Anya is an amazing alchemist, 
especially for her young age. As long 
as I've known her, she's never let 
anything stop her from pursuing what 
she believed in. 
 
I wouldn't doubt that Nihlathak is 
involved. Ever since her father died, 
they haven't gotten along.
`: []int{0, 1737}[1:],
	`57
Nihlathak was a vile demon that shall 
find his home among the tortured 
minions of Hell!
 
You battled the Darkness without fear. I 
laud your skill and courage.
`: []int{0, 63}[1:],
	`57
Nihlathak was the last of our Elders, 
whose charge it was to safeguard the 
mountain. He alone tried to guide us 
through the most desperate time in our 
history -- but he was just as helpless as 
the rest of us.
 
I cannot forgive his betrayal, but I can 
learn from it.
`: []int{0, 1153}[1:],
	`57
The demon hordes have grown powerful 
beyond measure, aided by our foolish 
mistakes. But I may have found a way 
to correct those mistakes...
`: []int{0, 943}[1:],
	`57
This is unlike any battle I have ever 
fought. While we ration food and 
water, the demon hordes feast nightly 
on the flesh and blood of our dead.
`: []int{0, 1627}[1:],
	`57
Though once considered shelter by our 
people, the Ice Caves offer no refuge 
from the dark horde. There are 
creatures there that will freeze your 
heart before feasting upon it.
`: []int{0, 2230}[1:],
	`57
We have lost many well-trained warriors 
to Baal's siege machines. Their range is 
great. Though, they are vulnerable if 
you close the distance quickly enough.
`: []int{0, 1624}[1:],
	`58
Early on, parties of our best scouts 
were ambushed by demons that 
spawned from the very air around 
them. Survivors often mentioned a 
strange creature floating in the 
distance.
 
Perhaps taking it down could prevent 
some nasty surprises.
`: []int{0, 1626}[1:],
	`58
My concerns have turned to my men 
taken prisoner on the battlefield by 
Baal's demons. I hate to think what's 
happened to them.
 
As you journey up the mountain, keep 
your eyes open for my soldiers and 
bring them back to me if you can.
`: []int{0, 2750}[1:],
	`58
Ohh...This is a truly horrible turn of 
events.
 
I know it seems you have always been 
one step behind, my friend. But look at 
it this way...You have evil on the run.
`: []int{0, 1476}[1:],
	`58
When I was a child, the Elders told us 
stories about the mountain and its 
power...and how the Barbarian people 
are bound to it as protectors.
 
Baal is not just taking our land -- with 
every step he takes up the mountain, 
he takes a part of who we are as a 
people.
`: []int{0, 1148}[1:],
	`59
I have spent decades trying to 
understand the forces at work in this 
world. But in the face of all that is 
transpiring, I realize how meager my 
knowledge is.
 
I will be of assistance where I can, my 
friend.
`: []int{0, 1949}[1:],
	`59
I may be just an armorer, but I know 
this...Baal plans to destroy the world 
with the secrets contained in that 
mountain. It doesn't take a genius to 
know he has to be stopped.
`: []int{0, 2452}[1:],
	`59
I would listen to Malah. Nihlathak 
speaks with a venomous tongue and 
acts as if the entire weight of this town 
rests upon his shoulders.
 
Perhaps there is more going on here 
than we know.
`: []int{0, 1110}[1:],
	`60
Baal has blocked Tyrael from entering 
the Worldstone Chamber? This truly 
has become a battle against Hell.
 
Whether or not it was the Heavens' 
decree, this is your fight now -- your 
destiny.
`: []int{0, 964}[1:],
	`60
I knew the Ancients would find you 
worthy of Mount Arreat's secrets. Now, 
stop Baal before he destroys all that is 
sacred.
`: []int{0, 1939}[1:],
	`60
I saw Nihlathak leave town just before 
you found Anya. He must be held 
accountable for his criminal deeds.
 
Find him and bring him back, if you can. 
Likely, he won't come willingly, and 
you'll be forced to kill him.
 
So be it.
`: []int{0, 2074}[1:],
	`60
Many outsiders believe that the 
fantastic stories about our ancestors, 
the Ancients, are but fables. However, I 
believe that the Ancients were more 
than human -- that mankind has fallen 
from what it once was.
`: []int{0, 1147}[1:],
	`60
Nihlathak's story does sound 
reasonable, considering what I've 
heard about Anya. However, the best 
lies are often hidden within truth.
`: []int{0, 465}[1:],
	`60
Our people believe that the Ancients 
protecting Mount Arreat have the 
power to stop Baal. Unfortunately, the 
Lord of Destruction has shown great 
power to undermine our faith.
`: []int{0, 1439}[1:],
	`60
Qual-Kehk is a worthy leader, but the 
losses have borne heavily upon him. He 
is only impatient because he judges the 
worth of a warrior by action, not 
words.
 
You must prove yourself worthy of his 
trust.
`: []int{0, 2235}[1:],
	`60
The destruction of the Worldstone does 
not bode well for our world. But I'll try 
not to worry...
 
After all, we have warriors like you 
fighting for us and for the Light.
 
Farewell!
`: []int{0, 2756}[1:],
	`60
This battle plays mind tricks on our 
warriors. Those fortunate enough to 
have returned from the mountainside 
claim to have seen angels in flight.
 
Fly they might, but that certainly does 
not make them angels.
`: []int{0, 2232}[1:],
	`60
Those catapults were like nothing I have 
ever seen before. You have prevailed 
against Shenk, my friend, but Baal is 
still far ahead of you.
`: []int{0, 2808}[1:],
	`61
Harrogath has great need of your 
powers, noble Druid. However, in the 
face of this supernatural onslaught, 
are your natural powers up to the 
task?
`: []int{0, 1622}[1:],
	`62
For one so young, Anya commands 
great respect. Now that she is here, I 
will make it a point to talk to her about 
Mount Arreat.
 
You should do the same.
`: []int{0, 1273}[1:],
	`63
I've offered Qual-Kehk my ideas on how 
to break the siege, but he dismisses 
them. Is it because I lack scars of 
battle, or does he think I'm a couple 
arrows short of a quiver?
`: []int{0, 1879}[1:],
	`63
If you are having trouble finding 
Qual-Kehk's soldiers, you should talk to 
Malah. She healed those who made it 
back before. Perhaps she would have 
some advice.
`: []int{0, 1741}[1:],
	`63
Nihlathak was never the kindest man. 
But for him to betray the whole world...
 
Ahh...Where shall his soul finally rest?
`: []int{0, 987}[1:],
	`63
The Ancients are not our enemies. 
Remember that! They are our 
ancestors -- our gods.
`: []int{0, 636}[1:],
	`63
You have ventured to a place beyond 
legend. You rush to face an evil few 
can even imagine.
 
Be careful, my friend, and may the Light 
watch over you.
`: []int{0, 2155}[1:],
	`63
You're an even greater warrior than I 
expected...Sorry for underestimating 
you.
 
As a token of my appreciation, I will 
craft sockets into an item of your 
choosing, and from now on, you'll get 
the best price for all my wares.
`: []int{0, 2549}[1:],
	`64
Legend has it that the top of Mount 
Arreat is guarded by the spirits of our 
ancestors. Though our people are 
forbidden to climb to the mountain's 
summit, some foreign travelers have 
made the attempt.
 
None have ever returned.
`: []int{0, 1880}[1:],
	`64
This may not be very encouraging, but 
you have fared better than the others 
who passed through those gates.
 
In regards to Shenk the Overseer, 
remember: a general is nothing without 
his men.
`: []int{0, 219}[1:],
	`64
Your gold is most helpful. Medical 
supplies for our wounded are scarce 
and very expensive.
 
When we can find a supplier, oh, we 
must pay dearly for what we need.
`: []int{0, 2557}[1:],
	`65
By reaching the summit, you cease 
being just a simple warrior. When you 
come back, you will be far more.
`: []int{0, 1420}[1:],
	`65
Ending the siege does not earn 
immediate respect, outsider.
 
Respect only comes with sacrifice -- 
something I'm sure you know nothing 
of.
`: []int{0, 2704}[1:],
	`65
I've heard that you Amazons excel at 
killing from a distance. From what I've 
seen, that's the best way to deal with 
Hell's minions. 
 
Are all of your kind so...big?
`: []int{0, 1874}[1:],
	`65
It is fortunate that this town has such a 
talented smith.
 
The quality of Larzuk's work surely 
complements your skills. In fact, he 
would have been quite welcome 
amongst the Horadrim.
`: []int{0, 1952}[1:],
	`65
My advice is to go in quick and hit hard. 
Nihlathak can't be half as tough as the 
beasts you've faced out there.
`: []int{0, 2187}[1:],
	`65
My worst fear has come true. Nihlathak 
has gone mad.
 
You must stop him, before he gives the 
Relic to the Lord of Destruction!
`: []int{0, 2700}[1:],
	`65
So, the Relic is lost. Do not dwell on 
failures past. It is your future that 
matters more.
`: []int{0, 928}[1:],
	`65
The Council of Elders always believed 
itself prepared for the coming of the 
Three. Obviously, we were not prepared 
enough.
`: []int{0, 948}[1:],
	`65
Why did you ever leave your beautiful 
islands to come to this frozen 
battleground? Perhaps if we both 
survive this, we'll travel back there 
together.
`: []int{0, 1875}[1:],
	`66
About to face Shenk the Overseer and 
stop the siege, are you? You should 
ask Malah to perform your last rites 
before you go, stranger.
`: []int{0, 2350}[1:],
	`66
This is a lively town during our victory 
celebrations. We Barbarians train long 
and hard from childhood to become 
warriors, and we celebrate with equal 
fervor.
 
You didn't happen to bring along any 
ale to trade?
`: []int{0, 1881}[1:],
	`66
Those demons have been out there 
since before your arrival. Can you do 
nothing to stop them?
 
Your task is a simple one, warrior. Find 
Shenk and destroy him.
`: []int{0, 1710}[1:],
	`67
I never liked Nihlathak, but I never 
suspected that he'd betray us! I just 
can't understand how an Elder could do 
such a thing.
`: []int{0, 1106}[1:],
	`67
It would be an honor to have a warrior 
of the Light fighting side-by-side with 
my men.
 
I can see your faith gives you great 
strength, Paladin, but don't expect it to 
keep you out of harm's way.
`: []int{0, 1621}[1:],
	`67
Qual-Kehk is useless. He has blindly 
sent our warriors to their deaths, 
assuming Baal's legions would fight as 
men do. Of course, everyone knows 
they do not.
`: []int{0, 942}[1:],
	`67
Soldiers who made it back told of a 
system of barricades placed among the 
mountain passes. They said that is 
where the prisoners are held.
`: []int{0, 2338}[1:],
	`67
You have proven yourself to these 
people. They look to you as their 
warrior, their champion.
`: []int{0, 1784}[1:],
	`68
Do not doubt yourself. I believe you are 
worthy to pass through the Ancients' 
gates, but you must believe, as well.
`: []int{0, 288}[1:],
	`68
So...You managed to stop the siege.
 
You are more powerful than I gave you 
credit for. You have rightfully earned 
my respect.
`: []int{0, 1534}[1:],
	`69
Demonic forces have been using our 
own towers and barricades against us. 
You know, it would be wise to cut the 
demons down in the open before they 
can mount those fortifications.
`: []int{0, 1877}[1:],
	`70
After so many have died, who are you 
to think you can accomplish what our 
proud warriors could not?
`: []int{0, 2702}[1:],
	`70
Nihlathak's despair is infectious -- and 
his behavior does not befit an Elder of 
his stature. I think we'd be better off 
without him.
`: []int{0, 1878}[1:],
	`70
The snake has slipped our grasp! While 
you were gone, Nihlathak disappeared.
 
I'll bet Anya knows how to track him 
down.
`: []int{0, 1316}[1:],
	`70
You knew it would eventually come down 
to this. Kill Baal! Finish the game!
`: []int{0, 2407}[1:],
	`70
You've walked on the burial grounds of 
our greatest ancestors. I'm not sure 
whether I should bow before you or 
revile you for committing sacrilege.
 
Tread lightly when you walk with gods.
`: []int{0, 2176}[1:],
	`71
Look. I must apologize.
 
I feel responsible for your current 
struggle. If I had only stopped 
Nihlathak from giving Baal the Relic, 
you would not have to fight those 
ghosts.
`: []int{0, 2240}[1:],
	`72
I am truly glad you are here, warrior. 
Perhaps things would be different now 
if we had asked for assistance from 
the neighboring kingdoms.
 
Our foolish, foolish pride...
`: []int{0, 1149}[1:],
	`72
So. You brought the lost sheep home to 
the shepherd. Well done.
`: []int{0, 2720}[1:],
	`72
You wouldn't have to test yourself 
against the Ancients if it weren't for 
Nihlathak's treachery. He was a fool 
and deserves to suffer for eternity.
`: []int{0, 1953}[1:],
	`73
I crafted the swords and armor for 
Qual-Kehk's men. They were like 
brothers to me. I can't imagine the pain 
they must be suffering.
 
Save them if you can!
`: []int{0, 2749}[1:],
	`73
You have inspired the people in this 
town -- not only those you rescued, but 
those you've helped in other ways as 
well.
`: []int{0, 133}[1:],
	`74
I know firsthand that captivity is a sad 
fate for a man. Find them quickly.
`: []int{0, 1094}[1:],
	`74
The Ancients themselves will envy our 
songs about you.
 
Please, don't forget about us! Farewell, 
my friend.
`: []int{0, 407}[1:],
	`75
Besting the Ancients in battle is a 
mighty feat indeed. I hope this means 
you're ready to battle Baal.
`: []int{0, 1286}[1:],
	`75
Every night, I've prayed to the Ancients 
to bring us peace...and now you must 
fight them.
 
Huh...Who shall I pray to now, warrior?
`: []int{0, 2701}[1:],
	`75
Has that infernal howling been keeping 
you awake at night, too? Some think 
it's the howling of animals, but those 
sounds don't come from any beast I've 
ever known.
`: []int{0, 1876}[1:],
	`76
As a kid, I used to play soldier among 
the barricades on the mountain. 
There's no easy way through that maze 
of walls.
`: []int{0, 1348}[1:],
	`77
Qual-Kehk and his men have been 
fighting to break the siege for some 
time. Where many have failed, you may 
succeed.
`: []int{0, 2444}[1:],
	`78
If Larzuk could sing or dance half as 
well as he smiths, he'd be married by 
now.
 
...Just look at those shoulders.
`: []int{0, 1152}[1:],
	`78
More of my men are still alive out there. 
I am certain of it!
 
Find them. Free them from their cages 
and bring them back to me.
`: []int{0, 1069}[1:],
	`78
The catapults are infernal machines 
made of demon flesh fused with steel.
 
Be wary of them.
`: []int{0, 2234}[1:],
	`78
What are you still doing here? I thought 
you were going off to die.
 
Go...Be quick about it.
`: []int{0, 688}[1:],
	`80
Poor Anya! I should've known Nihlathak 
was a traitor...
 
If you find him alive, kill him for me!
`: []int{0, 1143}[1:],
	`80
Thank you, hero, for rescuing me.
 
To show my personal gratitude, I give 
you this. I had it custom-made for you 
by Larzuk.
`: []int{0, 1194}[1:],
	`80
You've become a hero to this town, my 
friend. The shadows have lifted ever 
since you brought the Light to 
Harrogath.
`: []int{0, 2552}[1:],
	`81
Anya's father was my good friend. 
There are so many to mourn...I have 
no time for you!
`: []int{0, 946}[1:],
	`81
It seems like everyone feels Nihlathak 
played a part in Anya's disappearance.
 
Why would he do such a thing?
`: []int{0, 2237}[1:],
	`82
Baal's minions are not to be trifled with. 
In their bloodlust they will sacrifice 
themselves to destroy you.
`: []int{0, 1146}[1:],
	`82
The Ancients have honored you, and in 
turn, so do we. I have no remaining 
doubts about you, now.
`: []int{0, 2526}[1:],
	`82
You stand before me a worthy hero -- 
and on you rests the last hope of our 
people.
 
Bear it well, warrior.
`: []int{0, 682}[1:],
	`83
I believe that stopping the siege on 
Harrogath is the only way for you to 
earn the trust of these people.
`: []int{0, 1078}[1:],
	`84
Qual-Kehk's men have been imprisoned 
for some time. They are certain to be 
tired and weak.
 
You must protect them once you free 
them.
`: []int{0, 2188}[1:],
	`84
You've proven your skill at rescuing fair 
maidens...but how are you at killing 
vicious beasts?
`: []int{0, 1889}[1:],
	`85
If those men are being treated like I 
was, you must find them. They won't 
survive much longer.
`: []int{0, 1905}[1:],
	`85
Those of my men fortunate enough to 
escape on their own told me that they 
were held captive in the highlands and 
plateaus.
`: []int{0, 2258}[1:],
	`90
A test of mettle is a fitting rite of 
passage for a Barbarian hero.
`: []int{0, 1142}[1:],
	`90
As noble as Nihlathak's intentions are, 
nothing can excuse his actions.
`: []int{0, 1937}[1:],
	`90
Be cautious, warrior.
 
Though I am an experienced healer, 
resurrection is beyond my powers.
`: []int{0, 2555}[1:],
	`90
Beware! Baal grows stronger with every 
passing moment.
`: []int{0, 1017}[1:],
	`90
What's the matter, hero? Questioning 
your fortitude? I know we are.
`: []int{0, 1073}[1:],
	`93
Hmmm...What does Baal plan to do 
inside Mount Arreat?
`: []int{0, 1804}[1:],
	`93
So, you still live. You're either quick or a 
coward.
`: []int{0, 2267}[1:],
	`94
Oh...At last, the siege is ended. You 
truly are an angel.
 
I thank you.
`: []int{0, 1409}[1:],
	`94
When you talk to Nihlathak, be careful. 
There is no telling what he will say or 
do.
`: []int{0, 138}[1:],
	`94
You have proven you can take life, 
warrior, but can you save it as well?
`: []int{0, 1355}[1:],
	`95
They say that discretion, not 
procrastination, is the better part of 
valor.
`: []int{0, 2219}[1:],
	`96
Baal's legions seem countless, but 
slaying their commanders takes some 
of the fight out of them.
`: []int{0, 1625}[1:],
	`A Malus! This should go to Charsi.`:     []int{0, 857}[1:],
	`A coward's hiding place.`:               []int{0, 1797}[1:],
	`A fitting death for a traitor.`:         []int{0, 961}[1:],
	`A hero's mistake is finally corrected.`: []int{0, 1760}[1:],
	`Abaddon`:                                []int{0, 126}[1:],
	`Abominable`:                             []int{0, 979}[1:],
	`Achmel the Cursed`:                      []int{0, 1616}[1:],
	`Acrobatic`:                              []int{0, 1898}[1:],
	`Adds `:                                  []int{0, 164}[1:],
	`Adds Attacker Takes Damage`:             []int{0, 1915}[1:],
	`Adds Cold Damage to Attack`:             []int{0, 1974}[1:],
	`Adds Fire Damage to Attack`:             []int{0, 1998}[1:],
	`Adds Lightning Damage to Attack`:        []int{0, 1912}[1:],
	`Adds Mana and Life Regeneration`:        []int{0, 1914}[1:],
	`Adds Mana and Life Steal to Attack`:     []int{0, 1917}[1:],
	`Adds Poison Damage to Attack`:           []int{0, 1980}[1:],
	`Adds Resistance to Cold`:                []int{0, 1973}[1:],
	`Adds Resistance to Fire`:                []int{0, 1997}[1:],
	`Adds Resistance to Lightning`:           []int{0, 1911}[1:],
	`Adds Resistance to Poison`:              []int{0, 1979}[1:],
	`Adds to All Resistances`:                []int{0, 1904}[1:],
	`Adds to Attack Rating`:                  []int{0, 1903, 1971}[1:],
	`Adds to Chance to Find Magic Items`:     []int{0, 1910}[1:],
	`Adds to Damage vs. Undead`:              []int{0, 1909}[1:],
	`Adds to Defense Rating`:                 []int{0, 1970}[1:],
	`Adds to Dexterity`:                      []int{0, 1978}[1:],
	`Adds to Maximum Life`:                   []int{0, 1995}[1:],
	`Adds to Maximum Mana`:                   []int{0, 1972}[1:],
	`Adds to Strength`:                       []int{0, 1969}[1:],
	`Aegis`:                                  []int{0, 839}[1:],
	`Aerin Shield`:                           []int{0, 2086}[1:],
	`Agony Worm`:                             []int{0, 459}[1:],
	`Ahh, the familiar scent of death.`:      []int{0, 52}[1:],
	`Akaran Rondache`:                        []int{0, 2089}[1:],
	`Akaran Targe`:                           []int{0, 2088}[1:],
	`Akarat's Protector`:                     []int{0, 2806}[1:],
	`Aladdin's Eviscerator`:                  []int{0, 2627}[1:],
	`Alaric`:                                 []int{0, 1214}[1:],
	`Alarming`:                               []int{0, 586}[1:],
	`Alban`:                                  []int{0, 1274}[1:],
	`Alcuin`:                                 []int{0, 1235}[1:],
	`Aldur's Advance`:                        []int{0, 2073}[1:],
	`Aldur's Deception`:                      []int{0, 877}[1:],
	`Aldur's Gauntlet`:                       []int{0, 2744}[1:],
	`Aldur's Rhythm`:                         []int{0, 435}[1:],
	`Aldur's Stony Gaze`:                     []int{0, 1753}[1:],
	`Aldur's Watchtower`:                     []int{0, 2006}[1:],
	`Alma Negra`:                             []int{0, 2477}[1:],
	`Alma's Reflection`:                      []int{0, 749}[1:],
	`Alpha Helm`:                             []int{0, 2107}[1:],
	`Altar`:                                  []int{0, 470}[1:],
	`Altar of the Heavens`:                   []int{0, 424}[1:],
	`Ambergris`:                              []int{0, 152}[1:],
	`Amn`:                                    []int{0, 1119}[1:],
	`Amn Rune`:                               []int{0, 1827}[1:],
	`Amodeus' Manipulator`:                   []int{0, 2170}[1:],
	`An ancient manuscript...  This could be useful.`: []int{0, 160}[1:],
	`An eclipse... never a good omen.`:                []int{0, 486}[1:],
	`Ancient Eye`:                                     []int{0, 1772}[1:],
	`Ancients' Pledge`:                                []int{0, 1421}[1:],
	`And I thought the Forgotten Tower stank.`:        []int{0, 10}[1:],
	`Andariel's Visage`:                               []int{0, 937}[1:],
	`Angel's Song`:                                    []int{0, 814}[1:],
	`Anguish Worm`:                                    []int{0, 461}[1:],
	`Annihilus`:                                       []int{0, 1074}[1:],
	`Antimagic`:                                       []int{0, 2458}[1:],
	`Antlers`:                                         []int{0, 2104}[1:],
	`Anya`:                                            []int{0, 261}[1:],
	`Anya will personalize an item for you.`:          []int{0, 1482}[1:],
	`Apocrypha`:                                       []int{0, 2015}[1:],
	`Apothecary's Tote`:                               []int{0, 748}[1:],
	`Apply the Runes to a Socketed item in this order:`: []int{0, 466}[1:],
	`Arachnid Mesh`:                         []int{0, 2075}[1:],
	`Aragon's Icy Stare`:                    []int{0, 2043}[1:],
	`Aragon's Masterpiece`:                  []int{0, 2027}[1:],
	`Aragon's Storm Cloud`:                  []int{0, 2511}[1:],
	`Aragon's Sunfire`:                      []int{0, 2014}[1:],
	`Arcadian`:                              []int{0, 2491}[1:],
	`Arch-angel's Deliverance`:              []int{0, 2768}[1:],
	`Archon Plate`:                          []int{0, 912}[1:],
	`Archon Staff`:                          []int{0, 1749}[1:],
	`Arcing`:                                []int{0, 534}[1:],
	`Arctic Blast`:                          []int{0, 1506, 2295}[1:],
	`Are you sure you wish to continue?`:    []int{0, 813}[1:],
	`Argent`:                                []int{0, 1825}[1:],
	`Arioc's Needle`:                        []int{0, 1570}[1:],
	`Arkaine's Valor`:                       []int{0, 1028}[1:],
	`Arm of King Leoric`:                    []int{0, 2158}[1:],
	`Armageddon`:                            []int{0, 1423, 2359, 2366}[1:],
	`Armet`:                                 []int{0, 781}[1:],
	`Armor`:                                 []int{0, 939}[1:],
	`Armor Stand`:                           []int{0, 443, 444}[1:],
	`Armor:`:                                []int{0, 1023}[1:],
	`Arreat Plateau`:                        []int{0, 522}[1:],
	`Arreat Summit`:                         []int{0, 1163}[1:],
	`Arreat's Face`:                         []int{0, 143}[1:],
	`Artisan's`:                             []int{0, 469}[1:],
	`Arts`:                                  []int{0, 762}[1:],
	`Asheara's Wired Frame`:                 []int{0, 1445}[1:],
	`Ashwood Bow`:                           []int{0, 1261}[1:],
	`Assault Helmet`:                        []int{0, 1321}[1:],
	`Astral`:                                []int{0, 2490}[1:],
	`Astreon's Iron Ward`:                   []int{0, 1299}[1:],
	`At last...The summit of Mount Arreat.`: []int{0, 458}[1:],
	`Ataghan`:                               []int{0, 1934}[1:],
	`Athena's Wrath`:                        []int{0, 491}[1:],
	`Athletic`:                              []int{0, 2271}[1:],
	`Atma's Scarab`:                         []int{0, 533}[1:],
	`Atma's Wail`:                           []int{0, 98}[1:],
	`Aureolic`:                              []int{0, 2134}[1:],
	`Aurora's Guard`:                        []int{0, 2017}[1:],
	`Authority`:                             []int{0, 1424}[1:],
	`Avenger Guard`:                         []int{0, 1322}[1:],
	`Axe Dweller`:                           []int{0, 989}[1:],
	`Baal`:                                  []int{0, 107, 541, 628, 698, 1456}[1:],
	`Baal! Join your brothers in oblivion.`: []int{0, 54}[1:],
	`Baal! Nothing will stand in my way.`:   []int{0, 338}[1:],
	`Baal, never doubt my skills.`:          []int{0, 1765}[1:],
	`Baal, you shall no longer taint this mortal realm.`: []int{0, 2071}[1:],
	`Baal. I'm coming for you.`:                          []int{0, 1290}[1:],
	`Baezil's Vortex`:                                    []int{0, 456, 2515}[1:],
	`Bah! Is that all of them?`:                          []int{0, 914}[1:],
	`Bahamut's`:                                          []int{0, 1038}[1:],
	`Balrog Blade`:                                       []int{0, 1748}[1:],
	`Balrog Skin`:                                        []int{0, 844}[1:],
	`Balrog Spear`:                                       []int{0, 1884}[1:],
	`Banished Soul`:                                      []int{0, 1071}[1:],
	`Banner`:                                             []int{0, 345, 346, 996, 998}[1:],
	`Banshee's Wail`:                                     []int{0, 2628}[1:],
	`Baranar's Star`:                                     []int{0, 1658}[1:],
	`Barbarian Captive`:                                  []int{0, 268}[1:],
	`Barloc`:                                             []int{0, 1279}[1:],
	`Barrel`:                                             []int{0, 307, 386}[1:],
	`Barricade`:                                          []int{0, 224, 553}[1:],
	`Barricaded Door`:                                    []int{0, 227}[1:],
	`Barricaded Tower`:                                   []int{0, 228}[1:],
	`Bars can't hold a force of nature.`:                 []int{0, 1485}[1:],
	`Bartuc the Bloody`:                                  []int{0, 1617}[1:],
	`Bartuc's Chop Chop`:                                 []int{0, 2606}[1:],
	`Battle Cestus`:                                      []int{0, 1684}[1:],
	`Bear`:                                               []int{0, 31, 72}[1:],
	`Beast`:                                              []int{0, 1425}[1:],
	`Beauty`:                                             []int{0, 1426}[1:],
	`Ber`:                                                []int{0, 1614}[1:],
	`Ber Rune`:                                           []int{0, 1858}[1:],
	`Berserker`:                                          []int{0, 255}[1:],
	`Berserker Slayer`:                                   []int{0, 244}[1:],
	`Beserker Axe`:                                       []int{0, 1986}[1:],
	`Betrayal of Harrogath`:                              []int{0, 487}[1:],
	`Betrayer, you've reaped your reward.`:               []int{0, 1764}[1:],
	`Beware!`:                                            []int{0, 641}[1:],
	`Bill`:                                               []int{0, 1280}[1:],
	`Bing Sz Wang`:                                       []int{0, 603}[1:],
	`Black`:                                              []int{0, 1427}[1:],
	`Black Hades`:                                        []int{0, 1192}[1:],
	`Black books make for dark thoughts.`:                []int{0, 2208}[1:],
	`Blackbog's Sharp`:                                   []int{0, 587}[1:],
	`Blackhand Key`:                                      []int{0, 147}[1:],
	`Blackhorn`:                                          []int{0, 1918}[1:],
	`Blackhorn's Face`:                                   []int{0, 86}[1:],
	`Blackleach Blade`:                                   []int{0, 2057}[1:],
	`Blackoak Shield`:                                    []int{0, 2247}[1:],
	`Blade Barrier`:                                      []int{0, 843}[1:],
	`Blade Bow`:                                          []int{0, 1491}[1:],
	`Blade Fury`:                                         []int{0, 2697, 2703}[1:],
	`Blade Master`:                                       []int{0, 1671}[1:],
	`Blade Sentinel`:                                     []int{0, 62, 2638, 2641}[1:],
	`Blade Shield`:                                       []int{0, 2779, 2783}[1:],
	`Blade Talons`:                                       []int{0, 1681}[1:],
	`Blade of Ali Baba`:                                  []int{0, 1172}[1:],
	`Blades of Ice`:                                      []int{0, 2752, 2758}[1:],
	`Blanched`:                                           []int{0, 1929}[1:],
	`Blaze Ripper`:                                       []int{0, 282}[1:],
	`Blazing`:                                            []int{0, 2600}[1:],
	`Blighting`:                                          []int{0, 1344}[1:],
	`Bling Bling`:                                        []int{0, 2000}[1:],
	`Blood`:                                              []int{0, 1429}[1:],
	`Blood Boss`:                                         []int{0, 169}[1:],
	`Blood Bringer`:                                      []int{0, 166}[1:],
	`Blood Chalice`:                                      []int{0, 319}[1:],
	`Blood Comet`:                                        []int{0, 2244}[1:],
	`Blood Lord`:                                         []int{0, 1015}[1:],
	`Blood Rain`:                                         []int{0, 892}[1:],
	`Blood Raven's Charge`:                               []int{0, 811}[1:],
	`Blood Spirit`:                                       []int{0, 2151}[1:],
	`Blood Temptress`:                                    []int{0, 114}[1:],
	`Blood Witch`:                                        []int{0, 117}[1:],
	`Bloodletter`:                                        []int{0, 186}[1:],
	`Bloodlord Skull`:                                    []int{0, 1694}[1:],
	`Bloodmoon`:                                          []int{0, 777}[1:],
	`Bloodtree Stump`:                                    []int{0, 1708}[1:],
	`Bloody`:                                             []int{0, 1696}[1:],
	`Bloody Foothills`:                                   []int{0, 519}[1:],
	`Bohdan`:                                             []int{0, 873}[1:],
	`Bone`:                                               []int{0, 1430}[1:],
	`Bone Knife`:                                         []int{0, 1688}[1:],
	`Bone Visage`:                                        []int{0, 665}[1:],
	`Boneflame`:                                          []int{0, 2509}[1:],
	`Bonehew`:                                            []int{0, 1125}[1:],
	`Bonesaw Breaker`:                                    []int{0, 520}[1:],
	`Bonescalpel`:                                        []int{0, 583}[1:],
	`Boneshade`:                                          []int{0, 1755}[1:],
	`Boneslayer Blade`:                                   []int{0, 2054}[1:],
	`Boneweave`:                                          []int{0, 718}[1:],
	`Boneweave Boots`:                                    []int{0, 786}[1:],
	`Boreal`:                                             []int{0, 2177}[1:],
	`Bors`:                                               []int{0, 847}[1:],
	`Bound Spirit`:                                       []int{0, 84}[1:],
	`Bowyer's`:                                           []int{0, 1709}[1:],
	`Box`:                                                []int{0, 685}[1:],
	`Brain`:                                              []int{0, 1663}[1:],
	`Bramble`:                                            []int{0, 1431}[1:],
	`Bramble Mitts`:                                      []int{0, 775}[1:],
	`Brand`:                                              []int{0, 1333}[1:],
	`Breath of the Dying`:                                []int{0, 1334}[1:],
	`Broken Promise`:                                     []int{0, 1335}[1:],
	`Brom`:                                               []int{0, 859}[1:],
	`Brown`:                                              []int{0, 2260}[1:],
	`Bul-Kathos`:                                         []int{0, 2418}[1:],
	`Bul-Kathos' Children`:                               []int{0, 575}[1:],
	`Bul-Kathos' Custodian`:                              []int{0, 2058}[1:],
	`Bul-Kathos' Might`:                                  []int{0, 46}[1:],
	`Bul-Kathos' Sacred Charge`:                          []int{0, 2354}[1:],
	`Bul-Kathos' Tribal Guardian`:                        []int{0, 2609}[1:],
	`Bul-Kathos' Warden`:                                 []int{0, 2736}[1:],
	`Bul-Kathos' Wedding Band`:                           []int{0, 2217}[1:],
	`Bulwye`:                                             []int{0, 1164}[1:],
	`Burial Chest`:                                       []int{0, 309, 310}[1:],
	`Buriza-Do Kyanon`:                                   []int{0, 643}[1:],
	`Burly`:                                              []int{0, 1057}[1:],
	`Burning`:                                            []int{0, 1155}[1:],
	`Burning Pit`:                                        []int{0, 322}[1:],
	`Burst of Speed`:                                     []int{0, 2643, 2648}[1:],
	`Bush Wacker`:                                        []int{0, 131}[1:],
	`Butcher's Cleaver`:                                  []int{0, 2595}[1:],
	`Butcher's Pupil`:                                    []int{0, 2527}[1:],
	`Buzzing`:                                            []int{0, 2333}[1:],
	`By the Light! What is this place?`:                  []int{0, 765}[1:],
	`CHARACTER`:                                          []int{0, 125}[1:],
	`CLICK OK TO CONVERT YOUR ORIGINAL DIABLO II CHARACTER TO PLAY IN EXPANSION GAMES.`: []int{0, 808}[1:],
	`CONVERT TO`:                          []int{0, 805}[1:],
	`CONVERT TO EXPANSION`:                []int{0, 826}[1:],
	`CREATE NEW`:                          []int{0, 124}[1:],
	`CREATE NEW CHARACTER`:                []int{0, 832}[1:],
	`Caden`:                               []int{0, 1229}[1:],
	`Caduceus`:                            []int{0, 2004}[1:],
	`Cage`:                                []int{0, 367}[1:],
	`Cain! Go to the Rogue camp.`:         []int{0, 2216}[1:],
	`Cairn Shard`:                         []int{0, 1182}[1:],
	`Call to Arms`:                        []int{0, 1336}[1:],
	`Calling`:                             []int{0, 409}[1:],
	`Camphor`:                             []int{0, 349}[1:],
	`Can be Inserted into Socketed Items`: []int{0, 21}[1:],
	`Candles`:                             []int{0, 764}[1:],
	`Cantor Trophy`:                       []int{0, 1645}[1:],
	`Captain's`:                           []int{0, 121}[1:],
	`Carbuncle`:                           []int{0, 1867}[1:],
	`Caretaker's`:                         []int{0, 1538}[1:],
	`Carin Shard`:                         []int{0, 2320}[1:],
	`Carmine`:                             []int{0, 240}[1:],
	`Carnage Helm`:                        []int{0, 1367}[1:],
	`Carnage Leaver`:                      []int{0, 193}[1:],
	`Carrion Vine`:                        []int{0, 43, 1507, 2296}[1:],
	`Carrion Wind`:                        []int{0, 1010}[1:],
	`Carver Wolfrider`:                    []int{0, 664}[1:],
	`Cast a Level %d`:                     []int{0, 2266}[1:],
	`Catapult`:                            []int{0, 205, 206, 209, 213, 215, 218, 1160, 1185, 1186, 1187}[1:],
	`Catgut`:                              []int{0, 1331}[1:],
	`Celestial`:                           []int{0, 2811}[1:],
	`Cerebus' Bite`:                       []int{0, 2172}[1:],
	`Ceremonial Bow`:                      []int{0, 1262}[1:],
	`Ceremonial Javelin`:                  []int{0, 1302}[1:],
	`Ceremonial Pike`:                     []int{0, 1264}[1:],
	`Ceremonial Spear`:                    []int{0, 1263}[1:],
	`Cestus`:                              []int{0, 1704}[1:],
	`Chains of Honor`:                     []int{0, 1337}[1:],
	`Cham`:                                []int{0, 1646}[1:],
	`Cham Rune`:                           []int{0, 1860}[1:],
	`Champion Axe`:                        []int{0, 1730}[1:],
	`Champion Sword`:                      []int{0, 1608}[1:],
	`Chance`:                              []int{0, 1338}[1:],
	`Chance to afflict target: `:          []int{0, 1296}[1:],
	`Chance to stun: `:                    []int{0, 2813}[1:],
	`Chandelier`:                          []int{0, 334}[1:],
	`Chaos`:                               []int{0, 1339}[1:],
	`Chaotic`:                             []int{0, 1457}[1:],
	`Charge 1 - `:                         []int{0, 2815}[1:],
	`Charge 2 - `:                         []int{0, 2816}[1:],
	`Charge 3 - `:                         []int{0, 4}[1:],
	`Charged`:                             []int{0, 2025}[1:],
	`Charged Bolt`:                        []int{0, 2670}[1:],
	`Charged Bolt Sentry`:                 []int{0, 59, 2658}[1:],
	`Charges`:                             []int{0, 720}[1:],
	`Charsi will be thankful to get this Malus.`: []int{0, 617}[1:],
	`Charsi's Favor`: []int{0, 2613}[1:],
	`Chest`:          []int{0, 312, 381, 530, 676, 1243, 1278}[1:],
	`Chestnut`:       []int{0, 2493}[1:],
	`Chilled Froth`:  []int{0, 980}[1:],
	`Chilling`:       []int{0, 2772}[1:],
	`Choose the item that you wish to personalize with your name.`: []int{0, 836}[1:],
	`Choose the item to which you wish to add sockets.`:            []int{0, 838}[1:],
	`Chromatic`:               []int{0, 1586}[1:],
	`Chromatic Ire`:           []int{0, 184}[1:],
	`Cinnabar`:                []int{0, 287}[1:],
	`Circlet`:                 []int{0, 1701}[1:],
	`Clasped Orb`:             []int{0, 1847}[1:],
	`Class-specific`:          []int{0, 1001}[1:],
	`Claw Class`:              []int{0, 1569}[1:],
	`Claw Mastery`:            []int{0, 2376, 2378}[1:],
	`Claws`:                   []int{0, 1357, 1820}[1:],
	`Claws of Thunder`:        []int{0, 2713}[1:],
	`Cliffkiller`:             []int{0, 482}[1:],
	`Cloak of Shadows`:        []int{0, 2685}[1:],
	`Cloak of Shdws`:          []int{0, 2688}[1:],
	`Cloudcrack`:              []int{0, 2159}[1:],
	`Cloudy`:                  []int{0, 315}[1:],
	`Cloudy Sphere`:           []int{0, 1850}[1:],
	`Cobra Strike`:            []int{0, 2689, 2694}[1:],
	`Coldkill`:                []int{0, 344}[1:],
	`Coldsteal Eye`:           []int{0, 2169}[1:],
	`Coldsteel Eye`:           []int{0, 1062}[1:],
	`Colenzo the Annihilator`: []int{0, 1615}[1:],
	`Colorful`:                []int{0, 1996}[1:],
	`Colossus Blade`:          []int{0, 1733}[1:],
	`Colossus Crossbow`:       []int{0, 1513}[1:],
	`Colossus Girdle`:         []int{0, 707}[1:],
	`Colossus Sword`:          []int{0, 1715}[1:],
	`Colossus Voulge`:         []int{0, 1984}[1:],
	`Combat`:                  []int{0, 2357}[1:],
	`Combatant`:               []int{0, 265}[1:],
	`Commander's`:             []int{0, 1207}[1:],
	`Commanding the forces of nature, he summons wild beasts and raging storms to his side.`: []int{0, 305}[1:],
	`Communal`:        []int{0, 684}[1:],
	`Compact`:         []int{0, 26}[1:],
	`Condensing`:      []int{0, 1517}[1:],
	`Conqueror Crown`: []int{0, 1371}[1:],
	`Conquest Sword`:  []int{0, 1668}[1:],
	`Consecrated`:     []int{0, 36}[1:],
	`Conspiring with Baal... What a tragic mistake.`: []int{0, 862}[1:],
	`Constricting Ring`: []int{0, 2596}[1:],
	`Consult with the Ancients by activating the Altar of the Heavens.`: []int{0, 499}[1:],
	`Consult with the Ancients.`:                                        []int{0, 502}[1:],
	`Consumed Fire Boar`:                                                []int{0, 94}[1:],
	`Consumed Ice Boar`:                                                 []int{0, 100}[1:],
	`Corona`:                                                            []int{0, 878}[1:],
	`Coronet`:                                                           []int{0, 1702}[1:],
	`Corporal`:                                                          []int{0, 550}[1:],
	`Corpse`:                                                            []int{0, 417, 531}[1:],
	`Corpsemourn`:                                                       []int{0, 2280}[1:],
	`Corrosive`:                                                         []int{0, 2791}[1:],
	`Corruption... take flight.`:                                        []int{0, 1603}[1:],
	`Could this be a trap?`:                                             []int{0, 992}[1:],
	`Cow King's Hide`:                                                   []int{0, 1729}[1:],
	`Cow King's Hooves`:                                                 []int{0, 2771}[1:],
	`Cow King's Horns`:                                                  []int{0, 2667}[1:],
	`Cow King's Leathers`:                                               []int{0, 2012}[1:],
	`Crainte Vomir`:                                                     []int{0, 1300}[1:],
	`Cranebeak`:                                                         []int{0, 101}[1:],
	`Credendum`:                                                         []int{0, 51}[1:],
	`Crescent Moon`:                                                     []int{0, 1161, 1340}[1:],
	`Crest of Morn`:                                                     []int{0, 2809}[1:],
	`Crow Caw`:                                                          []int{0, 967}[1:],
	`Crown Shield`:                                                      []int{0, 2087}[1:],
	`Crown of Ages`:                                                     []int{0, 2747}[1:],
	`Crown of Thieves`:                                                  []int{0, 159}[1:],
	`Crude`:                                                             []int{0, 1404}[1:],
	`Crusader Bow`:                                                      []int{0, 1512}[1:],
	`Crusader Gauntlets`:                                                []int{0, 903}[1:],
	`Crush Beast`:                                                       []int{0, 79}[1:],
	`Cryptic Axe`:                                                       []int{0, 1886}[1:],
	`Cryptic Sword`:                                                     []int{0, 1828}[1:],
	`Crystalline Globe`:                                                 []int{0, 1849}[1:],
	`Crystalline Passage`:                                               []int{0, 1055}[1:],
	`Cunning`:                                                           []int{0, 2520}[1:],
	`Curly`:                                                             []int{0, 273}[1:],
	`Cursed`:                                                            []int{0, 729}[1:],
	`Cursing`:                                                           []int{0, 782}[1:],
	`Cycle of Life`:                                                     []int{0, 71}[1:],
	`Cyclone Armor`:                                                     []int{0, 1515, 2303}[1:],
	`DELETE CHARACTER`:                                                  []int{0, 833}[1:],
	`DESERT JOURNEY`:                                                    []int{0, 2190}[1:],
	`DESTRUCTION'S END`:                                                 []int{0, 1734}[1:],
	`Dac Farren`:                                                        []int{0, 850}[1:],
	`Damien`:                                                            []int{0, 1080}[1:],
	`Dangoon's Teaching`:                                                []int{0, 2518}[1:],
	`Dark Adherent`:                                                     []int{0, 799}[1:],
	`Dark Clan Crusher`:                                                 []int{0, 993}[1:],
	`Dark magic in a darker tomb...`:                                    []int{0, 2035}[1:],
	`Darkfear`:                                                          []int{0, 1655}[1:],
	`Darkforce Spawn`:                                                   []int{0, 185}[1:],
	`Darkness`:                                                          []int{0, 1341}[1:],
	`Darkone Wolfrider`:                                                 []int{0, 666}[1:],
	`Darksight Helm`:                                                    []int{0, 1798}[1:],
	`Darksoul`:                                                          []int{0, 2171}[1:],
	`Dawn Bringer`:                                                      []int{0, 427}[1:],
	`Daylight`:                                                          []int{0, 1342}[1:],
	`Dead Barbarian`:                                                    []int{0, 366, 495}[1:],
	`Death`:                                                             []int{0, 85, 1349}[1:],
	`Death Berserker`:                                                   []int{0, 1100}[1:],
	`Death Brawler`:                                                     []int{0, 1092}[1:],
	`Death Brigadier`:                                                   []int{0, 1102}[1:],
	`Death Cleaver`:                                                     []int{0, 2778}[1:],
	`Death Lord`:                                                        []int{0, 1019}[1:],
	`Death Mage`:                                                        []int{0, 738}[1:],
	`Death Mauler`:                                                      []int{0, 1090}[1:],
	`Death Pole`:                                                        []int{0, 374, 493}[1:],
	`Death Sentry`:                                                      []int{0, 242, 2767, 2777}[1:],
	`Death Slasher`:                                                     []int{0, 1093}[1:],
	`Death becomes you, Andariel.`:                                      []int{0, 2425}[1:],
	`Death's Fathom`:                                                    []int{0, 1209}[1:],
	`Death's Web`:                                                       []int{0, 1035, 2745}[1:],
	`Deathbit`:                                                          []int{0, 158}[1:],
	`Deathly Visage`:                                                    []int{0, 83}[1:],
	`Debris`:                                                            []int{0, 463}[1:],
	`Decapitator`:                                                       []int{0, 1669}[1:],
	`Decaying`:                                                          []int{0, 2573}[1:],
	`Deception`:                                                         []int{0, 1359}[1:],
	`Deckard Cain, leave this place!`:                                   []int{0, 1459}[1:],
	`Defeat all three Ancients without leaving Mount Arreat.`: []int{0, 500}[1:],
	`Defensive`:                       []int{0, 2352}[1:],
	`Defiled Warrior`:                 []int{0, 77}[1:],
	`Delerium`:                        []int{0, 1360}[1:],
	`Demon Crossbow`:                  []int{0, 1673}[1:],
	`Demon Gremlin`:                   []int{0, 182}[1:],
	`Demon Head`:                      []int{0, 1640}[1:],
	`Demon Heart`:                     []int{0, 1894}[1:],
	`Demon Imp`:                       []int{0, 180}[1:],
	`Demon Limb`:                      []int{0, 858}[1:],
	`Demon Machine`:                   []int{0, 773}[1:],
	`Demon Portal`:                    []int{0, 171, 172, 175, 406, 842}[1:],
	`Demon Rascal`:                    []int{0, 181}[1:],
	`Demon Sprite`:                    []int{0, 194}[1:],
	`Demon Steed`:                     []int{0, 1271}[1:],
	`Demon Trickster`:                 []int{0, 187}[1:],
	`Demon's Arch`:                    []int{0, 97}[1:],
	`Demonhead`:                       []int{0, 891}[1:],
	`Demonhorn's Edge`:                []int{0, 270}[1:],
	`Demonweb`:                        []int{0, 511}[1:],
	`Dense`:                           []int{0, 824}[1:],
	`Desire`:                          []int{0, 1361}[1:],
	`Despair`:                         []int{0, 1362}[1:],
	`Destroyer Helm`:                  []int{0, 1370}[1:],
	`Destruction`:                     []int{0, 1363}[1:],
	`Devil Star`:                      []int{0, 1845}[1:],
	`Devilkin Wolfrider`:              []int{0, 671}[1:],
	`Diablo... I will find you yet.`:  []int{0, 102}[1:],
	`Diadem`:                          []int{0, 1707}[1:],
	`Diamond`:                         []int{0, 1779}[1:],
	`Diamond Bow`:                     []int{0, 1630}[1:],
	`Diamond Mail`:                    []int{0, 807}[1:],
	`Dimensional Shard`:               []int{0, 1896}[1:],
	`Dire Wolf`:                       []int{0, 48, 68}[1:],
	`Disciplines`:                     []int{0, 759}[1:],
	`Dismiss`:                         []int{0, 1025}[1:],
	`Dismiss Hireling`:                []int{0, 601}[1:],
	`Divine`:                          []int{0, 1313}[1:],
	`Djinn Slayer`:                    []int{0, 1925}[1:],
	`Dol`:                             []int{0, 1167}[1:],
	`Dol Rune`:                        []int{0, 1830}[1:],
	`Doom`:                            []int{0, 1364}[1:],
	`Doombringer`:                     []int{0, 395}[1:],
	`Doomseer`:                        []int{0, 404}[1:],
	`Doppleganger's Shadow`:           []int{0, 1434}[1:],
	`Dracul's Grasp`:                  []int{0, 2062}[1:],
	`Dragon`:                          []int{0, 1365}[1:],
	`Dragon Claw`:                     []int{0, 2654, 2657}[1:],
	`Dragon Flight`:                   []int{0, 2760, 2766}[1:],
	`Dragon Tail`:                     []int{0, 2718, 2724}[1:],
	`Dragon Talon`:                    []int{0, 2426, 2429}[1:],
	`Dragonscale`:                     []int{0, 2664}[1:],
	`Dragontooth`:                     []int{0, 2224}[1:],
	`Drakeflame`:                      []int{0, 2665}[1:],
	`Dread`:                           []int{0, 1373}[1:],
	`Dream`:                           []int{0, 1375}[1:],
	`Dream Spirit`:                    []int{0, 2156}[1:],
	`Drifter Cavern`:                  []int{0, 527}[1:],
	`Drop Potion on Portrait to Heal`: []int{0, 597}[1:],
	`Druid Spirit`:                    []int{0, 687}[1:],
	`Drulan's Tongue`:                 []int{0, 1478, 2324}[1:],
	`Drus`:                            []int{0, 1189}[1:],
	`Dun`:                             []int{0, 2476}[1:],
	`Duress`:                          []int{0, 1376}[1:],
	`Duriel's Shell`:                  []int{0, 2414}[1:],
	`Dusk Shroud`:                     []int{0, 921}[1:],
	`Dust Storm`:                      []int{0, 17}[1:],
	`Dusty`:                           []int{0, 20}[1:],
	`Dwarf Star`:                      []int{0, 2534}[1:],
	`ENTER HELL`:                      []int{0, 2192}[1:],
	`EXPANSION`:                       []int{0, 803}[1:],
	`EXPANSION CHARACTER`:             []int{0, 804}[1:],
	`Eadgils`:                         []int{0, 1077}[1:],
	`Eagle`:                           []int{0, 328}[1:],
	`Eagle Orb`:                       []int{0, 1843}[1:],
	`Eagleeye`:                        []int{0, 333}[1:],
	`Eaglehorn`:                       []int{0, 1728}[1:],
	`Eaglewind`:                       []int{0, 422}[1:],
	`Earth Shifter`:                   []int{0, 1732}[1:],
	`Earth Spirit`:                    []int{0, 2153}[1:],
	`Earthshaker`:                     []int{0, 2797}[1:],
	`Eburine`:                         []int{0, 105}[1:],
	`Echoing`:                         []int{0, 1136}[1:],
	`Edge`:                            []int{0, 1378}[1:],
	`Eek, snakes!`:                    []int{0, 649}[1:],
	`Egtheow`:                         []int{0, 872}[1:],
	`El`:                              []int{0, 863}[1:],
	`El Espiritu`:                     []int{0, 2773}[1:],
	`El Infierno`:                     []int{0, 2395}[1:],
	`El Rune`:                         []int{0, 1823}[1:],
	`Elation`:                         []int{0, 1379}[1:],
	`Eld`:                             []int{0, 879}[1:],
	`Eld Rune`:                        []int{0, 1812}[1:],
	`Elder Staff`:                     []int{0, 1428}[1:],
	`Eldritch Orb`:                    []int{0, 1893}[1:],
	`Eldritch the Rectifier`:          []int{0, 297}[1:],
	`Elegant Blade`:                   []int{0, 1923}[1:],
	`Elemental`:                       []int{0, 802}[1:],
	`Elemental Charge-up Damage: `:    []int{0, 2614}[1:],
	`Elgifu`:                          []int{0, 1230}[1:],
	`Elysian`:                         []int{0, 183}[1:],
	`Emund`:                           []int{0, 1237}[1:],
	`Endlesshail`:                     []int{0, 606, 1291}[1:],
	`Enigma`:                          []int{0, 1380}[1:],
	`Enlightened`:                     []int{0, 919}[1:],
	`Enlightenment`:                   []int{0, 1381}[1:],
	`Enslaved`:                        []int{0, 400}[1:],
	`Entrapping`:                      []int{0, 2599}[1:],
	`Envy`:                            []int{0, 1382}[1:],
	`Erfor`:                           []int{0, 1082}[1:],
	`Erion's Bonehandle`:              []int{0, 745}[1:],
	`Eschuta's Temper`:                []int{0, 539, 540}[1:],
	`Eternal suffering would be too brief for you, Diablo.`: []int{0, 1675}[1:],
	`Eternity`:                               []int{0, 1383}[1:],
	`Eth`:                                    []int{0, 927}[1:],
	`Eth Rune`:                               []int{0, 1815}[1:],
	`Ethelred`:                               []int{0, 1228}[1:],
	`Ethereal (Cannot be Repaired)`:          []int{0, 834}[1:],
	`Ethereal Edge`:                          []int{0, 1347, 2737}[1:],
	`Ettin Axe`:                              []int{0, 883}[1:],
	`Eve of Destruction`:                     []int{0, 501}[1:],
	`Everlasting`:                            []int{0, 1309}[1:],
	`Evil`:                                   []int{0, 1711}[1:],
	`Evil Demon Hut`:                         []int{0, 1225}[1:],
	`Evil Urn`:                               []int{0, 370}[1:],
	`Evil flows from here.`:                  []int{0, 468}[1:],
	`Excalibur`:                              []int{0, 2775}[1:],
	`Executioner's Justice`:                  []int{0, 1921}[1:],
	`Exile`:                                  []int{0, 1384}[1:],
	`Expansion Disc`:                         []int{0, 150}[1:],
	`Expert's`:                               []int{0, 616}[1:],
	`Eye`:                                    []int{0, 2543}[1:],
	`Eyeback the Unleashed`:                  []int{0, 515}[1:],
	`Face the light or lurk in darkness.`:    []int{0, 286}[1:],
	`Fade`:                                   []int{0, 2705, 2708}[1:],
	`Faith`:                                  []int{0, 1385}[1:],
	`Faithful`:                               []int{0, 1647}[1:],
	`Fal`:                                    []int{0, 1247}[1:],
	`Fal Rune`:                               []int{0, 1835}[1:],
	`Falcata`:                                []int{0, 1940}[1:],
	`Falcon Mask`:                            []int{0, 2105}[1:],
	`Falconeye`:                              []int{0, 448}[1:],
	`Fallen Wolfrider`:                       []int{0, 672}[1:],
	`Famine`:                                 []int{0, 1386}[1:],
	`Fanatic`:                                []int{0, 252}[1:],
	`Fanatic Enslaved`:                       []int{0, 1180}[1:],
	`Fang`:                                   []int{0, 2604}[1:],
	`Fanged Helm`:                            []int{0, 1319}[1:],
	`Fanged Knife`:                           []int{0, 1811}[1:],
	`Fascia`:                                 []int{0, 2519}[1:],
	`Felicitous`:                             []int{0, 2798}[1:],
	`Feral`:                                  []int{0, 146}[1:],
	`Feral Axe`:                              []int{0, 1810}[1:],
	`Feral Claws`:                            []int{0, 1832}[1:],
	`Feral Pets: `:                           []int{0, 9}[1:],
	`Feral Rage`:                             []int{0, 1509, 2300}[1:],
	`Ferocious`:                              []int{0, 2013}[1:],
	`Festering`:                              []int{0, 776}[1:],
	`Festering Appendages`:                   []int{0, 1121}[1:],
	`Fetid Defiler`:                          []int{0, 232}[1:],
	`Fetish Trophy`:                          []int{0, 1643}[1:],
	`Fierce`:                                 []int{0, 1282}[1:],
	`Fiery`:                                  []int{0, 2817}[1:],
	`Finally... The drain lever.`:            []int{0, 320}[1:],
	`Find Baal's Throne Room.`:               []int{0, 2029}[1:],
	`Find Nihlathak in the Halls of Vaught.`: []int{0, 1498}[1:],
	`Find the Soldiers in the Frigid Highlands.`: []int{0, 447}[1:],
	`Finishing Move - `:                          []int{0, 15}[1:],
	`Finishing Move Bonuses`:                     []int{0, 5}[1:],
	`Fire`:                                       []int{0, 385}[1:],
	`Fire Blast`:                                 []int{0, 2371, 2375}[1:],
	`Fire Boar`:                                  []int{0, 473}[1:],
	`Fire Claws`:                                 []int{0, 1298, 2325}[1:],
	`Fire Pit`:                                   []int{0, 1124}[1:],
	`Firelizard's Talons`:                        []int{0, 1930}[1:],
	`Firestorm`:                                  []int{0, 1499, 2252}[1:],
	`Fissure`:                                    []int{0, 1514, 2302}[1:],
	`Fists of Fire`:                              []int{0, 2649, 2653}[1:],
	`Flag`:                                       []int{0, 655, 785, 2572}[1:],
	`Flame`:                                      []int{0, 1387}[1:],
	`Flamebellow`:                                []int{0, 235}[1:],
	`Flaming`:                                    []int{0, 2663}[1:],
	`Fleshbone`:                                  []int{0, 1920}[1:],
	`Fleshrender`:                                []int{0, 2277}[1:],
	`Fleshripper`:                                []int{0, 477}[1:],
	`Flowkrad's Fur`:                             []int{0, 2003}[1:],
	`Flowkrad's Grin`:                            []int{0, 285}[1:],
	`Flowkrad's Howl`:                            []int{0, 1269}[1:],
	`Flowkrad's Paws`:                            []int{0, 2748}[1:],
	`Flowkrad's Sinew`:                           []int{0, 289}[1:],
	`Flying Axe`:                                 []int{0, 1941}[1:],
	`Flying Knife`:                               []int{0, 1954}[1:],
	`Foci of the Vizjerei`:                       []int{0, 2524}[1:],
	`Foggy`:                                      []int{0, 1751}[1:],
	`Folcwald`:                                   []int{0, 1045}[1:],
	`Follow me.`:                                 []int{0, 24, 574, 609, 817, 1276, 2209, 2211}[1:],
	`Fool's`:                                     []int{0, 2125}[1:],
	`Fortitude`:                                  []int{0, 1388}[1:],
	`Fortune`:                                    []int{0, 1389}[1:],
	`Foul`:                                       []int{0, 2625}[1:],
	`Frantic`:                                    []int{0, 483}[1:],
	`Freawaru`:                                   []int{0, 1065}[1:],
	`Free the soldiers from their prison and lead them back to town.`: []int{0, 2186}[1:],
	`Freezing`:                               []int{0, 2342}[1:],
	`Frenzied Hell Spawn`:                    []int{0, 103}[1:],
	`Frenzied Ice Spawn`:                     []int{0, 108}[1:],
	`Friendship`:                             []int{0, 1391}[1:],
	`Frigid`:                                 []int{0, 1503}[1:],
	`Frigid Highlands`:                       []int{0, 1006}[1:],
	`Frisian`:                                []int{0, 1047}[1:],
	`From trash to treasure...`:              []int{0, 1610}[1:],
	`Frostwind`:                              []int{0, 1975}[1:],
	`Frozen Abyss`:                           []int{0, 984}[1:],
	`Frozen Anya`:                            []int{0, 2024}[1:],
	`Frozen Creeper`:                         []int{0, 195}[1:],
	`Frozen Horror`:                          []int{0, 203}[1:],
	`Frozen River`:                           []int{0, 526}[1:],
	`Frozen Scorch`:                          []int{0, 204}[1:],
	`Frozen Scourge`:                         []int{0, 202}[1:],
	`Frozen Terror`:                          []int{0, 201}[1:],
	`Frozen Tundra`:                          []int{0, 1213}[1:],
	`Frozenstein`:                            []int{0, 298}[1:],
	`Fuego Del Sol`:                          []int{0, 1178}[1:],
	`Fungal`:                                 []int{0, 1407}[1:],
	`Furious`:                                []int{0, 1824}[1:],
	`Fury`:                                   []int{0, 1392, 2351, 2358}[1:],
	`Fury Visor`:                             []int{0, 1368}[1:],
	`Gaean`:                                  []int{0, 1585}[1:],
	`Gaia's Wrath`:                           []int{0, 1919}[1:],
	`Gargoyle Head`:                          []int{0, 1638}[1:],
	`Gargoyle's Bite`:                        []int{0, 1922}[1:],
	`Garmund`:                                []int{0, 1064}[1:],
	`Gaudy`:                                  []int{0, 2205}[1:],
	`Gaya Wand`:                              []int{0, 1283}[1:],
	`Gerke's Sanctuary`:                      []int{0, 2355, 2516}[1:],
	`Geronimo's Fury`:                        []int{0, 2814}[1:],
	`Gheed's Fortune`:                        []int{0, 47}[1:],
	`Ghost Glaive`:                           []int{0, 1743}[1:],
	`Ghost Leach`:                            []int{0, 1240}[1:],
	`Ghost Liberator`:                        []int{0, 1754}[1:],
	`Ghost Spear`:                            []int{0, 1943}[1:],
	`Ghost Wand`:                             []int{0, 2040}[1:],
	`Ghostflame`:                             []int{0, 922}[1:],
	`Ghostly`:                                []int{0, 251}[1:],
	`Ghoulhide`:                              []int{0, 2010}[1:],
	`Giant Conch`:                            []int{0, 716}[1:],
	`Giant Maimer`:                           []int{0, 32}[1:],
	`Giant Skull`:                            []int{0, 2288}[1:],
	`Giant Thresher`:                         []int{0, 1988}[1:],
	`Gilded Shield`:                          []int{0, 2091}[1:],
	`Gillian's Brazier`:                      []int{0, 1208}[1:],
	`Gimmershred`:                            []int{0, 2755}[1:],
	`Ginther's Rift`:                         []int{0, 2289}[1:],
	`Glacial`:                                []int{0, 2313}[1:],
	`Glacial Trail`:                          []int{0, 1059}[1:],
	`Glimmershred`:                           []int{0, 16}[1:],
	`Globe of the Vizjerei`:                  []int{0, 2018}[1:],
	`Gloom`:                                  []int{0, 1393}[1:],
	`Gloom's Trap`:                           []int{0, 1597}[1:],
	`Glorious Axe`:                           []int{0, 1738}[1:],
	`Glory`:                                  []int{0, 1394}[1:],
	`Glowing Orb`:                            []int{0, 1848}[1:],
	`Go talk to Larzuk for your reward.`:     []int{0, 630}[1:],
	`Godly`:                                  []int{0, 1762}[1:],
	`Goldstrike Arch`:                        []int{0, 2195}[1:],
	`Golemlord's`:                            []int{0, 1084}[1:],
	`Good afternoon.`:                        []int{0, 691}[1:],
	`Good evening.`:                          []int{0, 692}[1:],
	`Good for you.`:                          []int{0, 690}[1:],
	`Good morning.`:                          []int{0, 658}[1:],
	`Goodbye, Izual.`:                        []int{0, 2514}[1:],
	`Gore Bearer`:                            []int{0, 80}[1:],
	`Gore Rider`:                             []int{0, 2310}[1:],
	`Gore Ripper`:                            []int{0, 633}[1:],
	`Gorgon Crossbow`:                        []int{0, 1593}[1:],
	`Gorm`:                                   []int{0, 1244}[1:],
	`Grand Charm`:                            []int{0, 1768}[1:],
	`Grand Matron Bow`:                       []int{0, 1304}[1:],
	`Grandmaster's`:                          []int{0, 950}[1:],
	`Gravepalm`:                              []int{0, 797}[1:],
	`Graverobber's`:                          []int{0, 3}[1:],
	`Great Bow`:                              []int{0, 1411}[1:],
	`Great Hauberk`:                          []int{0, 884}[1:],
	`Great Poleaxe`:                          []int{0, 1706}[1:],
	`Great Wyrm's`:                           []int{0, 1327}[1:],
	`Greater Claws`:                          []int{0, 2344}[1:],
	`Greater Hell Spawn`:                     []int{0, 225}[1:],
	`Greater Ice Spawn`:                      []int{0, 831}[1:],
	`Greater Talons`:                         []int{0, 2472}[1:],
	`Green`:                                  []int{0, 89}[1:],
	`Grief`:                                  []int{0, 1395}[1:],
	`Griffon Headdress`:                      []int{0, 2108}[1:],
	`Griffon's Eye`:                          []int{0, 2162}[1:],
	`Grim's Burning Dead`:                    []int{0, 588}[1:],
	`Grinding`:                               []int{0, 2116}[1:],
	`Griswold's Heart`:                       []int{0, 2002}[1:],
	`Griswold's Honor`:                       []int{0, 2059}[1:],
	`Griswold's Legacy`:                      []int{0, 2180}[1:],
	`Griswold's Redemption`:                  []int{0, 1216}[1:],
	`Griswold's Valor`:                       []int{0, 95}[1:],
	`Grizzly`:                                []int{0, 70}[1:],
	`Guard`:                                  []int{0, 264}[1:],
	`Guardian Angel`:                         []int{0, 2055, 2416}[1:],
	`Guardian Crown`:                         []int{0, 1372}[1:],
	`Guardian Naga`:                          []int{0, 829}[1:],
	`Guardian's`:                             []int{0, 2221}[1:],
	`Guillaume's Face`:                       []int{0, 860}[1:],
	`Gul`:                                    []int{0, 1438}[1:],
	`Gul Rune`:                               []int{0, 1857}[1:],
	`Gulletwound`:                            []int{0, 1747}[1:],
	`Gut Siphon`:                             []int{0, 2770}[1:],
	`Guthlaf`:                                []int{0, 1052}[1:],
	`Gymnastic`:                              []int{0, 13}[1:],
	`Hadeshorn`:                              []int{0, 1179}[1:],
	`Haemosu's Adamant`:                      []int{0, 2536}[1:],
	`Halaberd's Reign`:                       []int{0, 1480}[1:],
	`Halgaunt`:                               []int{0, 868}[1:],
	`Hallowed`:                               []int{0, 96}[1:],
	`Halls of Anguish`:                       []int{0, 542}[1:],
	`Halls of Pain`:                          []int{0, 1105}[1:],
	`Halls of Torment`:                       []int{0, 543}[1:],
	`Halls of Vaught`:                        []int{0, 547}[1:],
	`Hand Scythe`:                            []int{0, 2196}[1:],
	`Hand of Blessed Light`:                  []int{0, 1418}[1:],
	`Hand of Justice`:                        []int{0, 1396}[1:],
	`Harlequin Crest`:                        []int{0, 662}[1:],
	`Harmful`:                                []int{0, 816}[1:],
	`Harmony`:                                []int{0, 1397}[1:],
	`Harpoonist's`:                           []int{0, 217}[1:],
	`Harrogath`:                              []int{0, 518}[1:],
	`Harrogath can rest easier now.`:         []int{0, 1249}[1:],
	`Harrogath is free of your kind, demon.`: []int{0, 1761}[1:],
	`Hatchet Hands`:                          []int{0, 1483}[1:],
	`Hatred`:                                 []int{0, 1525}[1:],
	`Hatred stirs within me.`:                []int{0, 2248}[1:],
	`Hawk Branded`:                           []int{0, 2206}[1:],
	`Hawk Helm`:                              []int{0, 2103}[1:],
	`Hawkeye`:                                []int{0, 1432}[1:],
	`Hazy`:                                   []int{0, 452}[1:],
	`Head Hunter's Glory`:                    []int{0, 1447}[1:],
	`Headstriker`:                            []int{0, 1769}[1:],
	`Healfdane`:                              []int{0, 866}[1:],
	`Health Shrine`:                          []int{0, 1003}[1:],
	`Heart`:                                  []int{0, 375}[1:],
	`Heart Carver`:                           []int{0, 1419}[1:],
	`Heart of Wolverine`:                     []int{0, 239, 971, 1516}[1:],
	`Heart of the Oak`:                       []int{0, 1526}[1:],
	`Heated`:                                 []int{0, 2193}[1:],
	`Heater`:                                 []int{0, 915}[1:],
	`Heatholaf`:                              []int{0, 882}[1:],
	`Heaven's Brethren`:                      []int{0, 88}[1:],
	`Heaven's Light`:                         []int{0, 2818}[1:],
	`Heaven's Will`:                          []int{0, 1527}[1:],
	`Heavenly Stone`:                         []int{0, 1892}[1:],
	`Hel`:                                    []int{0, 1183}[1:],
	`Hel Rune`:                               []int{0, 1831}[1:],
	`Hell Gate`:                              []int{0, 894}[1:],
	`Hell Lord`:                              []int{0, 1016}[1:],
	`Hell Spawn`:                             []int{0, 978}[1:],
	`Hell Temptress`:                         []int{0, 113}[1:],
	`Hell Warden's Husk`:                     []int{0, 644}[1:],
	`Hell Whip`:                              []int{0, 170}[1:],
	`Hell Whisper`:                           []int{0, 39}[1:],
	`Hell Witch`:                             []int{0, 118}[1:],
	`Hell's Belle`:                           []int{0, 294}[1:],
	`Hellacious`:                             []int{0, 2411}[1:],
	`Hellforge Plate`:                        []int{0, 788}[1:],
	`Hellmouth`:                              []int{0, 1275}[1:],
	`Hello.`:                                 []int{0, 693}[1:],
	`Hellrack`:                               []int{0, 907}[1:],
	`Hellslayer`:                             []int{0, 1983}[1:],
	`Hellspawn Skull`:                        []int{0, 1691}[1:],
	`Helms:`:                                 []int{0, 1020}[1:],
	`Help me!`:                               []int{0, 1113}[1:],
	`Help!`:                                  []int{0, 1089}[1:],
	`Hengest`:                                []int{0, 1044}[1:],
	`Heorogar`:                               []int{0, 867}[1:],
	`Herald of Zakarum`:                      []int{0, 135}[1:],
	`Heraldic Shield`:                        []int{0, 2085}[1:],
	`Herb`:                                   []int{0, 357}[1:],
	`Heremod`:                                []int{0, 1039}[1:],
	`Hermetic`:                               []int{0, 2532}[1:],
	`Hexfire`:                                []int{0, 2053}[1:],
	`Hexing`:                                 []int{0, 708}[1:],
	`Hibernal`:                               []int{0, 41}[1:],
	`Hidden Stash`:                           []int{0, 351, 389, 397, 398, 425}[1:],
	`Hierophant Trophy`:                      []int{0, 1686}[1:],
	`Hierophant's`:                           []int{0, 1897}[1:],
	`Highland Blade`:                         []int{0, 1678}[1:],
	`Highlord's Wrath`:                       []int{0, 128}[1:],
	`Hightower's Watch`:                      []int{0, 668}[1:],
	`Hild`:                                   []int{0, 881}[1:],
	`Hireling`:                               []int{0, 794}[1:],
	`Hireling Inventory`:                     []int{0, 1029}[1:],
	`Hireling Screen`:                        []int{0, 766}[1:],
	`Hmm, a jade statue. What should I do with it?`: []int{0, 2645}[1:],
	`Hnaef`:                                 []int{0, 1051}[1:],
	`Hoku`:                                  []int{0, 1190}[1:],
	`Hollis`:                                []int{0, 1246}[1:],
	`Holy Tears`:                            []int{0, 1529}[1:],
	`Holy Thunder`:                          []int{0, 1530}[1:],
	`Homunculus`:                            []int{0, 2369}[1:],
	`Hone Sundan`:                           []int{0, 1000}[1:],
	`Honor`:                                 []int{0, 1531}[1:],
	`Honorable`:                             []int{0, 1957}[1:],
	`Hook`:                                  []int{0, 1354}[1:],
	`Horazon's Chalice`:                     []int{0, 1129}[1:],
	`Horazon. Your decoy is dead.`:          []int{0, 1034}[1:],
	`Horizon's Tornado`:                     []int{0, 303}[1:],
	`Horn`:                                  []int{0, 369}[1:],
	`Horned Helm`:                           []int{0, 1320}[1:],
	`How has this tree escaped corruption?`: []int{0, 697}[1:],
	`Howling`:                               []int{0, 1567}[1:],
	`Howling Visage`:                        []int{0, 2421}[1:],
	`Hrothgar`:                              []int{0, 864}[1:],
	`Humility`:                              []int{0, 1533}[1:],
	`Hunger`:                                []int{0, 1535, 1604, 2332}[1:],
	`Hunter's Guise`:                        []int{0, 2109}[1:],
	`Hurricane`:                             []int{0, 2367, 2370}[1:],
	`Husoldal Evo`:                          []int{0, 734}[1:],
	`Hwanin's Blessing`:                     []int{0, 2528}[1:],
	`Hwanin's Cordon`:                       []int{0, 1137}[1:],
	`Hwanin's Justice`:                      []int{0, 1788}[1:],
	`Hwanin's Majesty`:                      []int{0, 2011}[1:],
	`Hwanin's Refuge`:                       []int{0, 1495}[1:],
	`Hwanin's Splendor`:                     []int{0, 260}[1:],
	`Hydra Bow`:                             []int{0, 1576}[1:],
	`Hydra Edge`:                            []int{0, 1716}[1:],
	`Hydraskull`:                            []int{0, 768}[1:],
	`Hygelac`:                               []int{0, 870}[1:],
	`Hyperion`:                              []int{0, 871}[1:],
	`Hyperion Javelin`:                      []int{0, 1778}[1:],
	`Hyperion Spear`:                        []int{0, 1942}[1:],
	`I am dying.`:                           []int{0, 1107}[1:],
	`I am hurt!`:                            []int{0, 1088}[1:],
	`I am more experienced.`:                []int{0, 650}[1:],
	`I can sense Tal Rasha's presence now.`: []int{0, 2226}[1:],
	`I can smell why this tower was abandoned.`:               []int{0, 2692}[1:],
	`I detest spiders.`:                                       []int{0, 646}[1:],
	`I don't like it down here.`:                              []int{0, 2326}[1:],
	`I dread this place of darkness.`:                         []int{0, 1099}[1:],
	`I hate these vermin.`:                                    []int{0, 635}[1:],
	`I have a bad feeling about this.`:                        []int{0, 638}[1:],
	`I have no pity for him. Oblivion is his reward.`:         []int{0, 2315}[1:],
	`I needed that.`:                                          []int{0, 620}[1:],
	`I sense danger.`:                                         []int{0, 621}[1:],
	`I shall prove worthy.`:                                   []int{0, 2531}[1:],
	`I shall track the Prime Evils to the ends of the world.`: []int{0, 2202}[1:],
	`I should have known...`:                                  []int{0, 2051}[1:],
	`Iansang's Frenzy`:                                        []int{0, 2164}[1:],
	`Ice`:                                                     []int{0, 1536}[1:],
	`Ice Boar`:                                                []int{0, 92}[1:],
	`Ice Spawn`:                                               []int{0, 784}[1:],
	`Icy Cellar`:                                              []int{0, 529}[1:],
	`Immortal King`:                                           []int{0, 390}[1:],
	`Immortal King's Detail`:                                  []int{0, 2529}[1:],
	`Immortal King's Forge`:                                   []int{0, 778}[1:],
	`Immortal King's Pillar`:                                  []int{0, 2032}[1:],
	`Immortal King's Soul Cage`:                               []int{0, 399}[1:],
	`Immortal King's Stone Crusher`:                           []int{0, 849}[1:],
	`Immortal King's Will`:                                    []int{0, 955}[1:],
	`Immunity to Poison`:                                      []int{0, 728}[1:],
	`Impeccable`:                                              []int{0, 1408}[1:],
	`Increased Stack Size`:                                    []int{0, 2157}[1:],
	`Indestructible`:                                          []int{0, 2257}[1:],
	`Indiego's Fancy`:                                         []int{0, 2245}[1:],
	`Infectious`:                                              []int{0, 226}[1:],
	`Infernal Pit`:                                            []int{0, 130}[1:],
	`Inferno Sentry`:                                          []int{0, 243}[1:],
	`Infernostride`:                                           []int{0, 1596}[1:],
	`Infinity`:                                                []int{0, 1541}[1:],
	`Innocence`:                                               []int{0, 1542}[1:],
	`Insane Hell Spawn`:                                       []int{0, 109}[1:],
	`Insane Ice Spawn`:                                        []int{0, 1181}[1:],
	`Insidious`:                                               []int{0, 2461}[1:],
	`Insight`:                                                 []int{0, 1543}[1:],
	`Ip`:                                                      []int{0, 1197}[1:],
	`Iron Pelt`:                                               []int{0, 1250}[1:],
	`Irresistible`:                                            []int{0, 1403}[1:],
	`Islestrike`:                                              []int{0, 189}[1:],
	`Ist`:                                                     []int{0, 1422}[1:],
	`Ist Rune`:                                                []int{0, 1862}[1:],
	`It is good to work for someone who cares.`:         []int{0, 686}[1:],
	`It looks like jade. Perhaps it's worth something.`: []int{0, 2044}[1:],
	`It takes more than a siege to stop me.`:            []int{0, 301}[1:],
	`Ith`:                                               []int{0, 949}[1:],
	`Ith Rune`:                                          []int{0, 1822}[1:],
	`Ivory`:                                             []int{0, 918}[1:],
	`Jacinth`:                                           []int{0, 2238}[1:],
	`Jack's`:                                            []int{0, 1308}[1:],
	`Jade Talon`:                                        []int{0, 1450}[1:],
	`Jalal's Mane`:                                      []int{0, 2591}[1:],
	`Jar`:                                               []int{0, 335, 336, 383, 436, 439, 441, 442, 613}[1:],
	`Jared's Stone`:                                     []int{0, 2023}[1:],
	`Jawbone`:                                           []int{0, 618}[1:],
	`Jawbone Cap`:                                       []int{0, 1318}[1:],
	`Jawbone Visor`:                                     []int{0, 1323}[1:],
	`Jealousy`:                                          []int{0, 1544}[1:],
	`Jester's`:                                          []int{0, 1398}[1:],
	`Jewel`:                                             []int{0, 683}[1:],
	`Jeweler's`:                                         []int{0, 2167}[1:],
	`Jo`:                                                []int{0, 1631}[1:],
	`Jo Rune`:                                           []int{0, 1859}[1:],
	`Joker's`:                                           []int{0, 1108}[1:],
	`Judgement`:                                         []int{0, 1545}[1:],
	`Kaelim`:                                            []int{0, 1202}[1:],
	`Kang's Virtue`:                                     []int{0, 736}[1:],
	`Katar`:                                             []int{0, 1173}[1:],
	`Keep in Inventory to Gain Bonus`:                   []int{0, 905}[1:],
	`Keep in inventory to gain bonus`:                   []int{0, 651}[1:],
	`Keep it to thaw Anya
Malah's Potion`: []int{0, 376}[1:],
	`Keeper's`:           []int{0, 2721}[1:],
	`Kelpie Snare`:       []int{0, 2246}[1:],
	`Kenshi's`:           []int{0, 1415}[1:],
	`Khalim's Vengeance`: []int{0, 497}[1:],
	`Khan`:               []int{0, 845}[1:],
	`Kick Damage:`:       []int{0, 737}[1:],
	`Kick Damage: `:      []int{0, 2715}[1:],
	`Kill Baal in the Worldstone Chamber before he corrupts it.`: []int{0, 1757}[1:],
	`Kill Baal's Minions.`:     []int{0, 1758}[1:],
	`Kill Baal.`:               []int{0, 503}[1:],
	`Kill Nihlathak.`:          []int{0, 489}[1:],
	`Kill Shenk the Overseer.`: []int{0, 629}[1:],
	`King's Grace`:             []int{0, 1546}[1:],
	`Kingslayer`:               []int{0, 1547}[1:],
	`Kira's Guardian`:          []int{0, 2722}[1:],
	`Klatu`:                    []int{0, 1188}[1:],
	`Klisk`:                    []int{0, 846}[1:],
	`Knave's`:                  []int{0, 2661}[1:],
	`Knight's Vigil`:           []int{0, 1548}[1:],
	`Knowledge`:                []int{0, 1549}[1:],
	`Knuckles`:                 []int{0, 1351}[1:],
	`Ko`:                       []int{0, 1231}[1:],
	`Ko Rune`:                  []int{0, 1834}[1:],
	`Kord`:                     []int{0, 1191}[1:],
	`Korlic`:                   []int{0, 278}[1:],
	`Korlic the Protector`:     []int{0, 1600}[1:],
	`Kraken Shell`:             []int{0, 772}[1:],
	`Kuko Shakaku`:             []int{0, 999}[1:],
	`Kurast Shield`:            []int{0, 2185}[1:],
	`Lacerator`:                []int{0, 1712}[1:],
	`Lacquered Plate`:          []int{0, 904}[1:],
	`Laden`:                    []int{0, 1402}[1:],
	`Lance Guard`:              []int{0, 2175, 2299}[1:],
	`Lancer's`:                 []int{0, 1473}[1:],
	`Langer Briser`:            []int{0, 1345}[1:],
	`Lapis Lazuli`:             []int{0, 154}[1:],
	`Large Charm`:              []int{0, 1767}[1:],
	`Larry`:                    []int{0, 275}[1:],
	`Larzuk`:                   []int{0, 256}[1:],
	`Larzuk will add sockets to the item of your choice.`: []int{0, 626}[1:],
	`Larzuk's Champion`: []int{0, 266}[1:],
	`Lasher`:            []int{0, 120}[1:],
	`Last Wish`:         []int{0, 1550}[1:],
	`Lasting`:           []int{0, 467}[1:],
	`Lava Gout`:         []int{0, 2430}[1:],
	`Law`:               []int{0, 1557}[1:],
	`Lawbringer`:        []int{0, 1558}[1:],
	`Lawful`:            []int{0, 1252}[1:],
	`Laying of Hands`:   []int{0, 2184}[1:],
	`Leaf`:              []int{0, 1559}[1:],
	`Legend Spike`:      []int{0, 1661}[1:],
	`Legend Sword`:      []int{0, 890}[1:],
	`Legendary Mallet`:  []int{0, 1993}[1:],
	`Lem`:               []int{0, 1358}[1:],
	`Lem Rune`:          []int{0, 1866}[1:],
	`Lestron's Mark`:    []int{0, 2372}[1:],
	`Let Diablo's death end the reign of the Three!`: []int{0, 1002}[1:],
	`Level`:                         []int{0, 721}[1:],
	`Levers are made to be pulled.`: []int{0, 1101}[1:],
	`Leviathan`:                     []int{0, 509}[1:],
	`Lich Wand`:                     []int{0, 1672}[1:],
	`Lidless Wall`:                  []int{0, 2173}[1:],
	`Lief`:                          []int{0, 1165}[1:],
	`Life Seeker`:                   []int{0, 81}[1:],
	`Life Steal: `:                  []int{0, 2807}[1:],
	`Life Stealer`:                  []int{0, 82}[1:],
	`Lifechoke`:                     []int{0, 2397}[1:],
	`Lightning`:                     []int{0, 1560}[1:],
	`Lightning Sentry`:              []int{0, 60, 2725, 2728}[1:],
	`Lightsabre`:                    []int{0, 1759}[1:],
	`Lilac`:                         []int{0, 2264}[1:],
	`Lion Branded`:                  []int{0, 732}[1:],
	`Lion Helm`:                     []int{0, 1324}[1:],
	`Lionheart`:                     []int{0, 1561}[1:],
	`Lister the Tormentor`:          []int{0, 1666}[1:],
	`Lo`:                            []int{0, 1486}[1:],
	`Lo Rune`:                       []int{0, 1854}[1:],
	`Look for Anya under the Crystalline Passage by the Frozen River.`: []int{0, 472}[1:],
	`Lord Diablo I have bested you.`:                                   []int{0, 416}[1:],
	`Lore`:                                                             []int{0, 1562}[1:],
	`Loricated Mail`:                                                   []int{0, 637}[1:],
	`Love`:                                                             []int{0, 1563}[1:],
	`Lowers Resistance `:                                               []int{0, 331}[1:],
	`Loyalty`:                                                          []int{0, 1564}[1:],
	`Lucky`:                                                            []int{0, 1399}[1:],
	`Lum`:                                                              []int{0, 1215}[1:],
	`Lum Rune`:                                                         []int{0, 1836}[1:],
	`Luna`:                                                             []int{0, 796}[1:],
	`Lunar`:                                                            []int{0, 1958}[1:],
	`Lust`:                                                             []int{0, 1565}[1:],
	`Lycander's Aim`:                                                   []int{0, 2769}[1:],
	`Lycander's Flank`:                                                 []int{0, 830}[1:],
	`Lycanthropy`:                                                      []int{0, 1497, 2227}[1:],
	`M'avina's Battle Hymn`:                                            []int{0, 2060}[1:],
	`M'avina's Caster`:                                                 []int{0, 1500}[1:],
	`M'avina's Embrace`:                                                []int{0, 2181}[1:],
	`M'avina's Icy Clutch`:                                             []int{0, 290}[1:],
	`M'avina's Tenet`:                                                  []int{0, 2009}[1:],
	`M'avina's True Sight`:                                             []int{0, 341}[1:],
	`MEPHISTO'S JUNGLE`:                                                []int{0, 2191}[1:],
	`Madawc`:                                                           []int{0, 280}[1:],
	`Madawc the Guardian`:                                              []int{0, 1601}[1:],
	`Madness`:                                                          []int{0, 1566}[1:],
	`Magekiller's`:                                                     []int{0, 1963}[1:],
	`Magewrath`:                                                        []int{0, 735}[1:],
	`Magic Shrine`:                                                     []int{0, 995}[1:],
	`Magma Torquer`:                                                    []int{0, 750}[1:],
	`Magnus`:                                                           []int{0, 1166}[1:],
	`Magnus' Skin`:                                                     []int{0, 507}[1:],
	`Maiden Javelin`:                                                   []int{0, 1260}[1:],
	`Maiden Pike`:                                                      []int{0, 1259}[1:],
	`Maiden Spear`:                                                     []int{0, 1258}[1:],
	`Main Gate`:                                                        []int{0, 462}[1:],
	`Majestic`:                                                         []int{0, 2141}[1:],
	`Mal`:                                                              []int{0, 1406}[1:],
	`Mal Rune`:                                                         []int{0, 1863}[1:],
	`Malah`:                                                            []int{0, 1210}[1:],
	`Malice`:                                                           []int{0, 1574}[1:],
	`Malicious`:                                                        []int{0, 449}[1:],
	`Malignant Skull`:                                                  []int{0, 2593}[1:],
	`Mana Cost Per Raven: `:                                            []int{0, 2286}[1:],
	`Mana Recovered: `:                                                 []int{0, 1277}[1:],
	`Mana Steal: `:                                                     []int{0, 329}[1:],
	`Mancatcher`:                                                       []int{0, 1667}[1:],
	`Mang Song's Lesson`:                                               []int{0, 513}[1:],
	`Mara's Kaleidoscope`:                                              []int{0, 1932}[1:],
	`Maroon`:                                                           []int{0, 2622}[1:],
	`Marred`:                                                           []int{0, 402}[1:],
	`Marrow Grinder`:                                                   []int{0, 2677}[1:],
	`Marrowwalk`:                                                       []int{0, 1976}[1:],
	`Marshal's`:                                                        []int{0, 1695}[1:],
	`Martial`:                                                          []int{0, 760}[1:],
	`Master`:                                                           []int{0, 974}[1:],
	`Master's`:                                                         []int{0, 1633}[1:],
	`Matriarchal Bow`:                                                  []int{0, 1303}[1:],
	`Matriarchal Javelin`:                                              []int{0, 1307}[1:],
	`Matriarchal Pike`:                                                 []int{0, 1306}[1:],
	`Matriarchal Spear`:                                                []int{0, 1305}[1:],
	`Maul`:                                                             []int{0, 1510, 2301}[1:],
	`McAuley's Basis`:                                                  []int{0, 1066}[1:],
	`McAuley's Folly`:                                                  []int{0, 908}[1:],
	`McAuley's Pledge`:                                                 []int{0, 318}[1:],
	`McAuley's Superstition`:                                           []int{0, 709}[1:],
	`McAuley's Taboo`:                                                  []int{0, 246}[1:],
	`Meatscrape`:                                                       []int{0, 53}[1:],
	`Mechanic's`:                                                       []int{0, 896}[1:],
	`Medusa's Gaze`:                                                    []int{0, 411}[1:],
	`Melody`:                                                           []int{0, 1575}[1:],
	`Memory`:                                                           []int{0, 1578}[1:],
	`Menace Worm`:                                                      []int{0, 460}[1:],
	`Mentalist's`:                                                      []int{0, 902}[1:],
	`Mephisto's hatred was a poisonous void.`: []int{0, 2048}[1:],
	`Mephisto, you were no match for me.`:     []int{0, 2047}[1:],
	`Mephisto... I'm coming for you.`:         []int{0, 1490}[1:],
	`Merman's Sprocket`:                       []int{0, 2225}[1:],
	`Messerschmidt's Reaver`:                  []int{0, 848}[1:],
	`Metalgrid`:                               []int{0, 40}[1:],
	`Metalite's Girth`:                        []int{0, 1091}[1:],
	`Mighty Scepter`:                          []int{0, 1924}[1:],
	`Mind Blast`:                              []int{0, 2735, 2751}[1:],
	`Minion Skull`:                            []int{0, 1687}[1:],
	`Minion of Destruction`:                   []int{0, 2740}[1:],
	`Mirrored Boots`:                          []int{0, 898}[1:],
	`Mist`:                                    []int{0, 1579}[1:],
	`Mithril Coil`:                            []int{0, 787}[1:],
	`Mithril Point`:                           []int{0, 1690}[1:],
	`Mnemonic`:                                []int{0, 1083}[1:],
	`Moe`:                                     []int{0, 269, 780}[1:],
	`Mojo`:                                    []int{0, 2255}[1:],
	`Moldy`:                                   []int{0, 2239}[1:],
	`Molten Boulder`:                          []int{0, 1315, 2293}[1:],
	`Monarch`:                                 []int{0, 740}[1:],
	`Monumental`:                              []int{0, 2579}[1:],
	`Moon Lord`:                               []int{0, 1011}[1:],
	`Moon Shadow`:                             []int{0, 1030}[1:],
	`Moonfall`:                                []int{0, 1156}[1:],
	`Moonrend`:                                []int{0, 2533}[1:],
	`Mordoc's Marauder`:                       []int{0, 1571}[1:],
	`Morning`:                                 []int{0, 1580}[1:],
	`Mortal Crescent`:                         []int{0, 2285}[1:],
	`Moser's Blessed Circle`:                  []int{0, 249, 1058}[1:],
	`Mummified Trophy`:                        []int{0, 1641}[1:],
	`Musty`:                                   []int{0, 673}[1:],
	`My magic will break the siege.`:          []int{0, 1773}[1:],
	`My work here is truly done.`:             []int{0, 1295}[1:],
	`My, my, what a messy little demon.`:      []int{0, 1540}[1:],
	`Myrmidon Greaves`:                        []int{0, 706}[1:],
	`Mystery`:                                 []int{0, 1581}[1:],
	`Myth`:                                    []int{0, 1582}[1:],
	`Mythical Sword`:                          []int{0, 1989}[1:],
	`NEW EXPANSION CHARACTER`:                 []int{0, 819}[1:],
	`NEW EXPANSION OPEN CHARACTER.`:           []int{0, 821}[1:],
	`NEW EXPANSION REALM CHARACTER.`:          []int{0, 823}[1:],
	`Nadir`:                                   []int{0, 1583}[1:],
	`Nagas`:                                   []int{0, 640}[1:],
	`Naj's Ancient Vestige`:                   []int{0, 1174}[1:],
	`Naj's Circlet`:                           []int{0, 2222}[1:],
	`Naj's Light Plate`:                       []int{0, 2437}[1:],
	`Naj's Puzzler`:                           []int{0, 42}[1:],
	`Natalya's Mark`:                          []int{0, 2478}[1:],
	`Natalya's Odium`:                         []int{0, 1123}[1:],
	`Natalya's Shadow`:                        []int{0, 317}[1:],
	`Natalya's Soul`:                          []int{0, 2179}[1:],
	`Natalya's Totem`:                         []int{0, 2587}[1:],
	`Natural`:                                 []int{0, 854}[1:],
	`Nature's Kingdom`:                        []int{0, 1584}[1:],
	`Nature's Peace`:                          []int{0, 190}[1:],
	`Nebuchadnezzar's Storm`:                  []int{0, 2584}[1:],
	`Nef`:                                     []int{0, 911}[1:],
	`Nef Rune`:                                []int{0, 1814}[1:],
	`Nethercrow`:                              []int{0, 1095}[1:],
	`Nickel`:                                  []int{0, 1887}[1:],
	`Night`:                                   []int{0, 1461}[1:],
	`Night Lord`:                              []int{0, 1014}[1:],
	`Nightsummon`:                             []int{0, 134}[1:],
	`Nightwing's Veil`:                        []int{0, 880}[1:],
	`Nihlathak`:                               []int{0, 262, 267}[1:],
	`Nihlathak's Temple`:                      []int{0, 532}[1:],
	`Nihlathak. What led you to this end?`:    []int{0, 1763}[1:],
	`Nihlathak... you can't hide from me.`:    []int{0, 2586}[1:],
	`Nord's Tenderizer`:                       []int{0, 238}[1:],
	`Nosferatu's Coil`:                        []int{0, 2523}[1:],
	`Now I can leave this twisted nightmare.`: []int{0, 2037}[1:],
	`Noxious`:                              []int{0, 1437}[1:],
	`Null`:                                 []int{0, 1720}[1:],
	`Oak Sage`:                             []int{0, 236, 1501, 2253}[1:],
	`Oath`:                                 []int{0, 1462}[1:],
	`Obedience`:                            []int{0, 1463}[1:],
	`Oblivion`:                             []int{0, 1464}[1:],
	`Obsession`:                            []int{0, 1465}[1:],
	`Obsidian`:                             []int{0, 510}[1:],
	`Odiferous`:                            []int{0, 731}[1:],
	`Odium`:                                []int{0, 358}[1:],
	`Of the Choir`:                         []int{0, 661}[1:],
	`Of the Maiden`:                        []int{0, 793}[1:],
	`Of the Sniper`:                        []int{0, 744}[1:],
	`Of the Stiletto`:                      []int{0, 1081}[1:],
	`Offensive`:                            []int{0, 2356}[1:],
	`Ogre Axe`:                             []int{0, 1816}[1:],
	`Ogre Gauntlets`:                       []int{0, 711}[1:],
	`Ogre Maul`:                            []int{0, 1785}[1:],
	`Ogun's Fierce Visage`:                 []int{0, 2618}[1:],
	`Ogun's Lash`:                          []int{0, 110}[1:],
	`Ogun's Shadow`:                        []int{0, 2619}[1:],
	`Ogun's Vengeance`:                     []int{0, 2022}[1:],
	`Ohm`:                                  []int{0, 1470}[1:],
	`Ohm Rune`:                             []int{0, 1855}[1:],
	`Ondal's Almighty`:                     []int{0, 481}[1:],
	`Ondal's Wisdom`:                       []int{0, 970}[1:],
	`One can't judge a book by its cover.`: []int{0, 1292}[1:],
	`Onela`:                                []int{0, 1079}[1:],
	`Oops...Did I do that?`:                []int{0, 2741}[1:],
	`Orb Class`:                            []int{0, 2546}[1:],
	`Original`:                             []int{0, 2553}[1:],
	`Ormus may know something about this unusual blade.`: []int{0, 779}[1:],
	`Ormus' Robe`:  []int{0, 2540}[1:],
	`Ormus' Robes`: []int{0, 2}[1:],
	`Ormus... You have strange taste in books.`: []int{0, 1294}[1:],
	`Ormus... study the book well.`:             []int{0, 2038}[1:],
	`Orphan's Call`:                             []int{0, 1682}[1:],
	`Ort`:                                       []int{0, 991}[1:],
	`Ort Rune`:                                  []int{0, 1819}[1:],
	`Oslaf`:                                     []int{0, 1053}[1:],
	`Overlord`:                                  []int{0, 168}[1:],
	`Overseer`:                                  []int{0, 119}[1:],
	`Overseer Skull`:                            []int{0, 1692}[1:],
	`Pagan's Athame`:                            []int{0, 820}[1:],
	`Pain Worm`:                                 []int{0, 451}[1:],
	`Palo Grande`:                               []int{0, 2739}[1:],
	`Paradox`:                                   []int{0, 2197}[1:],
	`Parkersor's Calm`:                          []int{0, 512}[1:],
	`Passion`:                                   []int{0, 1466}[1:],
	`Passive Bonus to Wolves and Bears`:         []int{0, 2019}[1:],
	`Patience`:                                  []int{0, 1467}[1:],
	`Patter`:                                    []int{0, 1468}[1:],
	`Peace`:                                     []int{0, 1469}[1:],
	`Pearl`:                                     []int{0, 237}[1:],
	`Peasant Crown`:                             []int{0, 2535}[1:],
	`Pellet Bow`:                                []int{0, 1577}[1:],
	`Penitence`:                                 []int{0, 2559}[1:],
	`Per Level`:                                 []int{0, 722}[1:],
	`Per Player in Party`:                       []int{0, 733}[1:],
	`Percent Chance to Cast`:                    []int{0, 700}[1:],
	`Peril`:                                     []int{0, 2560}[1:],
	`Pernicious`:                                []int{0, 2438}[1:],
	`Perpetual`:                                 []int{0, 1888}[1:],
	`Pestilence`:                                []int{0, 2561}[1:],
	`Pestilent`:                                 []int{0, 2578}[1:],
	`Pet`:                                       []int{0, 165}[1:],
	`Pet Life: `:                                []int{0, 1012}[1:],
	`Phase Blade`:                               []int{0, 1683}[1:],
	`Phnx Strike`:                               []int{0, 2799}[1:],
	`Phoenix`:                                   []int{0, 2562}[1:],
	`Phoenix Egg`:                               []int{0, 1314}[1:],
	`Phoenix Strike`:                            []int{0, 2794}[1:],
	`Pierre Tombale Couant`:                     []int{0, 1679}[1:],
	`Piety`:                                     []int{0, 2563}[1:],
	`Pillar of Faith`:                           []int{0, 2564}[1:],
	`Pindleskin`:                                []int{0, 1060}[1:],
	`Pit of Acheron`:                            []int{0, 129}[1:],
	`Pitblood Thirst`:                           []int{0, 177}[1:],
	`Plague`:                                    []int{0, 2565}[1:],
	`Plague Bearer`:                             []int{0, 433}[1:],
	`Plague Poppy`:                              []int{0, 323}[1:],
	`Planting the dead... How odd.`:             []int{0, 960}[1:],
	`Po`:                                        []int{0, 1199}[1:],
	`Po Rune`:                                   []int{0, 1837}[1:],
	`Poison Creeper`:                            []int{0, 34, 1493}[1:],
	`Pole`:                                      []int{0, 339, 494}[1:],
	`Polished Wand`:                             []int{0, 1999}[1:],
	`Pompeii's Wrath`:                           []int{0, 142}[1:],
	`Possessed`:                                 []int{0, 253}[1:],
	`Powered`:                                   []int{0, 1736}[1:],
	`Praise`:                                    []int{0, 2566}[1:],
	`Prayer`:                                    []int{0, 2567}[1:],
	`Precocious`:                                []int{0, 2762}[1:],
	`Preserved Head`:                            []int{0, 1635}[1:],
	`Preserver's`:                               []int{0, 1840}[1:],
	`Pride`:                                     []int{0, 2568}[1:],
	`Principle`:                                 []int{0, 2447}[1:],
	`Prison Door`:                               []int{0, 423}[1:],
	`Prison of Ice`:                             []int{0, 471}[1:],
	`Protector Shield`:                          []int{0, 2090}[1:],
	`Prowess in Battle`:                         []int{0, 2448}[1:],
	`Prowling Dead`:                             []int{0, 75}[1:],
	`Prudence`:                                  []int{0, 2449}[1:],
	`Psn Creep`:                                 []int{0, 2215}[1:],
	`Psyc Hammer`:                               []int{0, 2382}[1:],
	`Psychic`:                                   []int{0, 408}[1:],
	`Psychic Hammer`:                            []int{0, 2379}[1:],
	`Pul`:                                       []int{0, 1374}[1:],
	`Pul Rune`:                                  []int{0, 1865}[1:],
	`Punishing`:                                 []int{0, 1328}[1:],
	`Punishment`:                                []int{0, 2450}[1:],
	`Pure`:                                      []int{0, 1553}[1:],
	`Purity`:                                    []int{0, 2451}[1:],
	`Pus Spitter`:                               []int{0, 2318}[1:],
	`Putrid Defiler`:                            []int{0, 230}[1:],
	`Pyre of Flesh`:                             []int{0, 314}[1:],
	`Qual-Kehk`:                                 []int{0, 263}[1:],
	`Qual-Khek's Enforcer`:                      []int{0, 379}[1:],
	`Que-Hegan's Wisdom`:                        []int{0, 757, 1268}[1:],
	`Question`:                                  []int{0, 2453}[1:],
	`Quhab`:                                     []int{0, 2163}[1:],
	`Quill`:                                     []int{0, 2574}[1:],
	`Quixotic`:                                  []int{0, 1619}[1:],
	`Rabbit Slayer`:                             []int{0, 291}[1:],
	`Rabies`:                                    []int{0, 1520, 2323}[1:],
	`Radament's Sphere`:                         []int{0, 2787}[1:],
	`Radiance`:                                  []int{0, 2454}[1:],
	`Rage Mask`:                                 []int{0, 1325}[1:],
	`Raging`:                                    []int{0, 652}[1:],
	`Ragnar`:                                    []int{0, 1254}[1:],
	`Raiden's Crutch`:                           []int{0, 2052}[1:],
	`Rain`:                                      []int{0, 2455}[1:],
	`Rainbow`:                                   []int{0, 594}[1:],
	`Rainbow Facet`:                             []int{0, 1443}[1:],
	`Ral`:                                       []int{0, 975}[1:],
	`Ral Rune`:                                  []int{0, 1818}[1:],
	`Rancid Defiler`:                            []int{0, 233}[1:],
	`Rank Defiler`:                              []int{0, 234}[1:],
	`Raven`:                                     []int{0, 29, 1031, 1332, 2214}[1:],
	`Raven Frost`:                               []int{0, 1723}[1:],
	`Ravenlore`:                                 []int{0, 1446}[1:],
	`Ravens: `:                                  []int{0, 2800}[1:],
	`Razor's Edge`:                              []int{0, 2294}[1:],
	`Razorswitch`:                               []int{0, 1310}[1:],
	`Razortail`:                                 []int{0, 2001}[1:],
	`Realgar`:                                   []int{0, 1717}[1:],
	`Reanimated Horde`:                          []int{0, 73}[1:],
	`Reason`:                                    []int{0, 2456}[1:],
	`Reclusive`:                                 []int{0, 885}[1:],
	`Red`:                                       []int{0, 2457}[1:],
	`Red Light`:                                 []int{0, 678}[1:],
	`Reflex Bow`:                                []int{0, 1256}[1:],
	`Rehire`:                                    []int{0, 602}[1:],
	`Reinforced Mace`:                           []int{0, 1826}[1:],
	`Repairs %d durability in %d seconds`:       []int{0, 2263}[1:],
	`Repairs %d durability per second`:          []int{0, 2259}[1:],
	`Replenishes quantity`:                      []int{0, 2265}[1:],
	`Repulsive`:                                 []int{0, 887}[1:],
	`Rescue %d more Soldiers in the Frigid Highlands.`: []int{0, 450}[1:],
	`Rescue Anya.`:                         []int{0, 1455}[1:],
	`Rescue on Mount Arreat`:               []int{0, 446}[1:],
	`Resonant`:                             []int{0, 156}[1:],
	`Resurrect`:                            []int{0, 1027}[1:],
	`Resurrect %s: %d`:                     []int{0, 605}[1:],
	`Retreat!`:                             []int{0, 413}[1:],
	`Return to Qual-Kehk for your reward.`: []int{0, 464}[1:],
	`Return to dust, Radament.`:            []int{0, 1195}[1:],
	`Revenge`:                              []int{0, 1532}[1:],
	`Rhyme`:                                []int{0, 2463}[1:],
	`Ribcracker`:                           []int{0, 604}[1:],
	`Rift`:                                 []int{0, 2464}[1:],
	`Riftlash`:                             []int{0, 1599}[1:],
	`Riftslash`:                            []int{0, 1523}[1:],
	`Right Click to Cast
Scroll of Resistance`: []int{0, 562}[1:],
	`Right-click to Open Inventory (%s)`: []int{0, 600}[1:],
	`Righteous`:                          []int{0, 2680}[1:],
	`Rings`:                              []int{0, 1452}[1:],
	`Riphook`:                            []int{0, 373}[1:],
	`Rite of Passage`:                    []int{0, 387, 496}[1:],
	`Robin's Yolk`:                       []int{0, 1680}[1:],
	`Robineye`:                           []int{0, 2144}[1:],
	`Rockhew`:                            []int{0, 1114}[1:],
	`Rockstopper`:                        []int{0, 149}[1:],
	`Rondache`:                           []int{0, 2084}[1:],
	`Rose`:                               []int{0, 2580}[1:],
	`Rose Branded`:                       []int{0, 997}[1:],
	`Rot Walker`:                         []int{0, 756}[1:],
	`Rotting`:                            []int{0, 2321}[1:],
	`Royal`:                              []int{0, 544}[1:],
	`Royal Shield`:                       []int{0, 2131}[1:],
	`Rude`:                               []int{0, 1056}[1:],
	`Ruins... the fate of all cities.`:   []int{0, 1806}[1:],
	`Run away!`:                          []int{0, 355}[1:],
	`Rune Master`:                        []int{0, 2283}[1:],
	`Rune Slayer`:                        []int{0, 2581}[1:],
	`Runic Talons`:                       []int{0, 1960}[1:],
	`Rusty`:                              []int{0, 1400}[1:],
	`SEARCH FOR BAAL`:                    []int{0, 2199}[1:],
	`Sacred`:                             []int{0, 428}[1:],
	`Sacred Armor`:                       []int{0, 610}[1:],
	`Sacred Charge`:                      []int{0, 284}[1:],
	`Sacred Feathers`:                    []int{0, 2110}[1:],
	`Sacred Globe`:                       []int{0, 1844}[1:],
	`Sacred Rondache`:                    []int{0, 2133}[1:],
	`Sacred Targe`:                       []int{0, 2132}[1:],
	`Samuel's Caretaker`:                 []int{0, 2042}[1:],
	`Samurai Justice`:                    []int{0, 1301}[1:],
	`Sanctuary`:                          []int{0, 2465}[1:],
	`Sander's Folly`:                     []int{0, 1756}[1:],
	`Sander's Paragon`:                   []int{0, 1086}[1:],
	`Sander's Riprap`:                    []int{0, 247}[1:],
	`Sander's Superstition`:              []int{0, 1122}[1:],
	`Sander's Taboo`:                     []int{0, 1159}[1:],
	`Sandstorm Trek`:                     []int{0, 2440}[1:],
	`Sanguinary`:                         []int{0, 69}[1:],
	`Sanguine`:                           []int{0, 624}[1:],
	`Sankekur's Resurrection`:            []int{0, 197}[1:],
	`Saracen's Chance`:                   []int{0, 1931}[1:],
	`Savage Helmet`:                      []int{0, 1326}[1:],
	`Say 'Retreat'`:                      []int{0, 2076}[1:],
	`Sazabi's Cobalt Redeemer`:           []int{0, 689}[1:],
	`Sazabi's Ghost Liberator`:           []int{0, 2517}[1:],
	`Sazabi's Grand Tribute`:             []int{0, 208}[1:],
	`Sazabi's Mental Sheath`:             []int{0, 825}[1:],
	`Scalp`:                              []int{0, 139}[1:],
	`Scarab Husk`:                        []int{0, 769}[1:],
	`Scarabshell Boots`:                  []int{0, 930}[1:],
	`Scarlet`:                            []int{0, 851}[1:],
	`Schaefer's Hammer`:                  []int{0, 2347}[1:],
	`Schooled in the Martial Arts, her mind and body are deadly weapons.`: []int{0, 306}[1:],
	`Scintillating`:                  []int{0, 2492}[1:],
	`Scissors Katar`:                 []int{0, 258}[1:],
	`Scissors Quhab`:                 []int{0, 2419}[1:],
	`Scissors Suwayyah`:              []int{0, 1907}[1:],
	`Scorched`:                       []int{0, 1868}[1:],
	`Scourge`:                        []int{0, 1725}[1:],
	`Screaming`:                      []int{0, 2489}[1:],
	`Scroll of Knowledge`:            []int{0, 136}[1:],
	`Scyld`:                          []int{0, 865}[1:],
	`Seige Tower`:                    []int{0, 1169}[1:],
	`Self-Repair`:                    []int{0, 523}[1:],
	`Sensei's`:                       []int{0, 276}[1:],
	`Septic`:                         []int{0, 667}[1:],
	`Seraph Rod`:                     []int{0, 1908}[1:],
	`Seraph's Hymn`:                  []int{0, 506}[1:],
	`Serendipity`:                    []int{0, 2466}[1:],
	`Serpents! I expected worse.`:    []int{0, 2201}[1:],
	`Serrated`:                       []int{0, 1193}[1:],
	`Sexton Trophy`:                  []int{0, 1644}[1:],
	`Shade Falcon`:                   []int{0, 1607}[1:],
	`Shadow`:                         []int{0, 50, 758, 2467}[1:],
	`Shadow Bow`:                     []int{0, 1555}[1:],
	`Shadow Dancer`:                  []int{0, 2278}[1:],
	`Shadow Killer`:                  []int{0, 362}[1:],
	`Shadow Master`:                  []int{0, 245, 2789}[1:],
	`Shadow Plate`:                   []int{0, 924}[1:],
	`Shadow Touch`:                   []int{0, 1157}[1:],
	`Shadow Warrior`:                 []int{0, 572, 2709}[1:],
	`Shadow of Doubt`:                []int{0, 2468}[1:],
	`Shae`:                           []int{0, 1154}[1:],
	`Shae Rune`:                      []int{0, 1829}[1:],
	`Shaftstop`:                      []int{0, 2413}[1:],
	`Shakabra's Crux`:                []int{0, 2632}[1:],
	`Shako`:                          []int{0, 608}[1:],
	`Shaman's`:                       []int{0, 822}[1:],
	`Shank`:                          []int{0, 1356}[1:],
	`Shape`:                          []int{0, 800}[1:],
	`Sharptooth Slayer`:              []int{0, 283}[1:],
	`Shdw Master`:                    []int{0, 2793}[1:],
	`Shdw Warrior`:                   []int{0, 2712}[1:],
	`Shenk the Overseer`:             []int{0, 229}[1:],
	`Shenk, your command has ended.`: []int{0, 2068}[1:],
	`Shields:`:                       []int{0, 1021}[1:],
	`Shifting`:                       []int{0, 801}[1:],
	`Shillelagh`:                     []int{0, 1412}[1:],
	`Shimmering`:                     []int{0, 1731}[1:],
	`Shivering`:                      []int{0, 490}[1:],
	`Shock Wave`:                     []int{0, 1605, 2334}[1:],
	`Shock Web`:                      []int{0, 2433, 2637}[1:],
	`Shocking`:                       []int{0, 35}[1:],
	`Shogukusha's`:                   []int{0, 2393}[1:],
	`Shoots `:                        []int{0, 2804}[1:],
	`Shrine`:                         []int{0, 368, 391, 1117}[1:],
	`Shrub`:                          []int{0, 414, 415}[1:],
	`Siege Beast`:                    []int{0, 78}[1:],
	`Siege Control`:                  []int{0, 393}[1:],
	`Siege Door`:                     []int{0, 330}[1:],
	`Siege on Harrogath`:             []int{0, 445}[1:],
	`Sigemund`:                       []int{0, 1036}[1:],
	`Siggard's Stealth`:              []int{0, 1985}[1:],
	`Sightless Veil`:                 []int{0, 2620}[1:],
	`Sigurd`:                         []int{0, 1239}[1:],
	`Silence`:                        []int{0, 2469}[1:],
	`Silkweave`:                      []int{0, 524}[1:],
	`Silver-edged Axe`:               []int{0, 1650}[1:],
	`Sinblade`:                       []int{0, 538}[1:],
	`Singing`:                        []int{0, 340}[1:],
	`Siren`:                          []int{0, 115}[1:],
	`Siren's Call`:                   []int{0, 2291}[1:],
	`Siren's Song`:                   []int{0, 2470}[1:],
	`Skill 10`:                       []int{0, 695}[1:],
	`Skill 11`:                       []int{0, 696}[1:],
	`Skill 12`:                       []int{0, 702}[1:],
	`Skill 13`:                       []int{0, 703}[1:],
	`Skill 14`:                       []int{0, 704}[1:],
	`Skill 15`:                       []int{0, 705}[1:],
	`Skill 16`:                       []int{0, 710}[1:],
	`Skill 9`:                        []int{0, 694}[1:],
	`Skin of Flayed One`:             []int{0, 1013}[1:],
	`Skin of the Flayed One`:         []int{0, 403}[1:],
	`Skin of the Vipermagi`:          []int{0, 1112}[1:],
	`Skull Collector`:                []int{0, 332}[1:],
	`Skullder's Ire`:                 []int{0, 2548}[1:],
	`Skulls and Rocks`:               []int{0, 343, 382, 677}[1:],
	`Skulltred`:                      []int{0, 2696}[1:],
	`Sky Spirit`:                     []int{0, 2154}[1:],
	`Skystrike`:                      []int{0, 2394}[1:],
	`Slashers`:                       []int{0, 1352}[1:],
	`Slayer`:                         []int{0, 90}[1:],
	`Slayer Guard`:                   []int{0, 1366}[1:],
	`Small Charm`:                    []int{0, 1766}[1:],
	`Small Crescent`:                 []int{0, 1657}[1:],
	`Smiting`:                        []int{0, 1801}[1:],
	`Smoke`:                          []int{0, 365, 2471}[1:],
	`Smoked Sphere`:                  []int{0, 1846}[1:],
	`Smoking`:                        []int{0, 141}[1:],
	`Smoldering`:                     []int{0, 2459}[1:],
	`Snake Tongue`:                   []int{0, 2207}[1:],
	`Snakes... I hate snakes.`:       []int{0, 2570}[1:],
	`Snapchip Shatter`:               []int{0, 296}[1:],
	`Snow Drifter`:                   []int{0, 923}[1:],
	`Snowclash`:                      []int{0, 976}[1:],
	`Snowy`:                          []int{0, 377}[1:],
	`So dark... perfect.`:            []int{0, 2297}[1:],
	`So much treasure almost covers the stench.`:   []int{0, 1067}[1:],
	`So, Tal Rasha... This is your resting place.`: []int{0, 220}[1:],
	`So, it begins.`:                       []int{0, 956}[1:],
	`So, this is what the Ancients guard.`: []int{0, 163}[1:],
	`So, this is where evil hides.`:        []int{0, 2034}[1:],
	`Sol`:                                  []int{0, 1135}[1:],
	`Sol Creep`:                            []int{0, 2331}[1:],
	`Sol Rune`:                             []int{0, 1838}[1:],
	`Solar Creeper`:                        []int{0, 965, 1524}[1:],
	`Sorrow`:                               []int{0, 2473}[1:],
	`Soul`:                                 []int{0, 316}[1:],
	`Soul Drainer`:                         []int{0, 1479}[1:],
	`Soul Reaper`:                          []int{0, 1595}[1:],
	`Soulfeast Tine`:                       []int{0, 223}[1:],
	`Soulless`:                             []int{0, 2660}[1:],
	`Sounding`:                             []int{0, 221}[1:],
	`Sparking`:                             []int{0, 1048}[1:],
	`Sparkling Ball`:                       []int{0, 1851}[1:],
	`Sparroweye`:                           []int{0, 2460}[1:],
	`Spearmaiden's`:                        []int{0, 2373}[1:],
	`Spellsteel`:                           []int{0, 790}[1:],
	`Spider Bow`:                           []int{0, 1670}[1:],
	`Spiderweb Sash`:                       []int{0, 771}[1:],
	`Spike Generator`:                      []int{0, 58}[1:],
	`Spike Thorn`:                          []int{0, 1448}[1:],
	`Spikes: `:                             []int{0, 2801}[1:],
	`Spineripper`:                          []int{0, 983}[1:],
	`Spire of Honor`:                       []int{0, 774}[1:],
	`Spired Helm`:                          []int{0, 717}[1:],
	`Spirit`:                               []int{0, 973, 2479}[1:],
	`Spirit Barbs`:                         []int{0, 2337}[1:],
	`Spirit Forge`:                         []int{0, 1270}[1:],
	`Spirit Keeper`:                        []int{0, 1435}[1:],
	`Spirit Mask`:                          []int{0, 2106}[1:],
	`Spirit Ward`:                          []int{0, 2510}[1:],
	`Spirit Wolf`:                          []int{0, 388}[1:],
	`Spirit of Barbs`:                      []int{0, 241, 966, 1612}[1:],
	`Spiritual`:                            []int{0, 405}[1:],
	`Spiteful`:                             []int{0, 1537}[1:],
	`Splay`:                                []int{0, 1353}[1:],
	`Spleen`:                               []int{0, 325}[1:],
	`Splendor`:                             []int{0, 2480}[1:],
	`Squire's Cover`:                       []int{0, 254}[1:],
	`Stag Bow`:                             []int{0, 1255}[1:],
	`Stalagmite`:                           []int{0, 1572}[1:],
	`Stalwart`:                             []int{0, 985}[1:],
	`Starlight`:                            []int{0, 2481}[1:],
	`Stars: `:                              []int{0, 653}[1:],
	`Stash`:                                []int{0, 350}[1:],
	`Static`:                               []int{0, 454}[1:],
	`Stealskull`:                           []int{0, 2284}[1:],
	`Stealth`:                              []int{0, 726, 2482}[1:],
	`Steel`:                                []int{0, 2483}[1:],
	`Steel Carapace`:                       []int{0, 1598}[1:],
	`Steel Pillar`:                         []int{0, 2174}[1:],
	`Steel Shade`:                          []int{0, 1727}[1:],
	`Steelrend`:                            []int{0, 2161}[1:],
	`Stellar`:                              []int{0, 57}[1:],
	`Still Water`:                          []int{0, 2484}[1:],
	`Sting`:                                []int{0, 2485}[1:],
	`Stinging`:                             []int{0, 615}[1:],
	`Stone`:                                []int{0, 2486}[1:],
	`Stone Crusher`:                        []int{0, 552}[1:],
	`Stone Rattle`:                         []int{0, 64}[1:],
	`Stonerage`:                            []int{0, 2304}[1:],
	`Stoneraven`:                           []int{0, 1026}[1:],
	`Stop the Siege by killing Shenk the Overseer in the Bloody Foothills.`: []int{0, 623}[1:],
	`Storm`:                             []int{0, 2487}[1:],
	`Stormchaser`:                       []int{0, 2210}[1:],
	`Stormlash`:                         []int{0, 2629}[1:],
	`Stormrider`:                        []int{0, 2276}[1:],
	`Stormshield`:                       []int{0, 1936}[1:],
	`Stormspike`:                        []int{0, 2314}[1:],
	`Stormspire`:                        []int{0, 2420}[1:],
	`Stormwillow`:                       []int{0, 592}[1:],
	`Stout`:                             []int{0, 2576}[1:],
	`Strange`:                           []int{0, 1234}[1:],
	`Strange... an unexpected eclipse.`: []int{0, 216}[1:],
	`Strength`:                          []int{0, 2488}[1:],
	`String of Ears`:                    []int{0, 2343}[1:],
	`Strong Oak`:                        []int{0, 1714}[1:],
	`Stygian Fury`:                      []int{0, 199}[1:],
	`Stygian Harlot`:                    []int{0, 1171}[1:],
	`Stygian Pike`:                      []int{0, 1955}[1:],
	`Stygian Pilum`:                     []int{0, 1883}[1:],
	`Succubus`:                          []int{0, 111}[1:],
	`Succubus Skull`:                    []int{0, 1693}[1:],
	`Such corruption in this place...`:  []int{0, 675}[1:],
	`Such stones are common back home.`: []int{0, 2513}[1:],
	`Suicide Branch`:                    []int{0, 906}[1:],
	`Sum Spt Wolf`:                      []int{0, 2254}[1:],
	`Summon D Wolf`:                     []int{0, 2319}[1:],
	`Summon Dire Wolf`:                  []int{0, 1519}[1:],
	`Summon Grizzly`:                    []int{0, 1613, 2349}[1:],
	`Summon Spirit Wolf`:                []int{0, 1504}[1:],
	`Summoner, the dark magics have corrupted you.`: []int{0, 957}[1:],
	`Summoning`:           []int{0, 798}[1:],
	`Sun Spirit`:          []int{0, 2152}[1:],
	`Sur`:                 []int{0, 1502}[1:],
	`Sur Rune`:            []int{0, 1853}[1:],
	`Sureshrill Frost`:    []int{0, 1223}[1:],
	`Suwayyah`:            []int{0, 1651}[1:],
	`Swap Weapons`:        []int{0, 715}[1:],
	`Swinging Heads`:      []int{0, 337}[1:],
	`Swirling Crystal`:    []int{0, 1891}[1:],
	`Swordguard`:          []int{0, 855}[1:],
	`TERROR'S END`:        []int{0, 2545}[1:],
	`THE SISTER'S LAMENT`: []int{0, 2189}[1:],
	`Taebaek's Glory`:     []int{0, 2031}[1:],
	`Tail`:                []int{0, 348}[1:],
	`Take Anya's portal to Nihlathak's Temple.`: []int{0, 488}[1:],
	`Take Tyreal's Portal to Safety.`:           []int{0, 517}[1:],
	`Tal`:                                       []int{0, 959}[1:],
	`Tal Rasha's Adjudication`:                  []int{0, 1739}[1:],
	`Tal Rasha's Fine-Spun Cloth`:               []int{0, 954}[1:],
	`Tal Rasha's Guardianship`:                  []int{0, 1664}[1:],
	`Tal Rasha's Horadric Crest`:                []int{0, 2311}[1:],
	`Tal Rasha's Lidless Eye`:                   []int{0, 2200}[1:],
	`Tal Rasha's Wrappings`:                     []int{0, 713}[1:],
	`Tal Rune`:                                  []int{0, 1817}[1:],
	`Talberd's Law`:                             []int{0, 1312}[1:],
	`Talic`:                                     []int{0, 281}[1:],
	`Talic the Defender`:                        []int{0, 1602}[1:],
	`Talk to Anya before you continue through the Crystalline Passage, past the Glacial Trail, to proceed up Mount Arreat to the Summit.`: []int{0, 492}[1:],
	`Talk to Anya.`:                    []int{0, 485}[1:],
	`Talk to Malah for your reward.`:   []int{0, 484}[1:],
	`Talk to Malah.`:                   []int{0, 474}[1:],
	`Talk to Tyreal.`:                  []int{0, 516}[1:],
	`Tang's Battle Standard`:           []int{0, 2348}[1:],
	`Tang's Fore-Fathers`:              []int{0, 2541}[1:],
	`Tang's Imperial Robes`:            []int{0, 2298}[1:],
	`Tang's Rule`:                      []int{0, 1070}[1:],
	`Tang's Throne`:                    []int{0, 2398}[1:],
	`Targe`:                            []int{0, 2083}[1:],
	`Telling of Beads`:                 []int{0, 412}[1:],
	`Tempest`:                          []int{0, 2495}[1:],
	`Templar's Might`:                  []int{0, 2396}[1:],
	`Temptation`:                       []int{0, 2496}[1:],
	`Terra's Guardian`:                 []int{0, 1735}[1:],
	`Terrene`:                          []int{0, 1632}[1:],
	`Terror`:                           []int{0, 2497}[1:],
	`Terror stalks Hell no more.`:      []int{0, 211}[1:],
	`Thank you.`:                       []int{0, 614}[1:],
	`Thanks.`:                          []int{0, 612}[1:],
	`Tharr`:                            []int{0, 1200}[1:],
	`The Achron Magus`:                 []int{0, 619}[1:],
	`The Ancients' Way`:                []int{0, 528}[1:],
	`The Atlantean`:                    []int{0, 2605, 2668}[1:],
	`The Baker`:                        []int{0, 647}[1:],
	`The Black Adder`:                  []int{0, 2361}[1:],
	`The Butcher`:                      []int{0, 1674}[1:],
	`The Candlestick Maker`:            []int{0, 648}[1:],
	`The Cat's Eye`:                    []int{0, 1977}[1:],
	`The Cranium Basher`:               []int{0, 953}[1:],
	`The Disciple`:                     []int{0, 2182}[1:],
	`The Ensanguinator`:                []int{0, 1245}[1:],
	`The Evil brotherhood is no more.`: []int{0, 2213}[1:],
	`The Fetid Sprinkler`:              []int{0, 1913}[1:],
	`The Gavel of Pain`:                []int{0, 302}[1:],
	`The Gladiator's Bane`:             []int{0, 2644}[1:],
	`The Grandfather`:                  []int{0, 1982}[1:],
	`The Harbinger`:                    []int{0, 763}[1:],
	`The Harvester`:                    []int{0, 2279}[1:],
	`The Horadrim have left their mark here.`: []int{0, 761}[1:],
	`The Impaler`:                               []int{0, 2759}[1:],
	`The Longest Rod`:                           []int{0, 2733}[1:],
	`The Meat Scraper`:                          []int{0, 2282}[1:],
	`The Minotaur`:                              []int{0, 1248, 2554}[1:],
	`The Oculus`:                                []int{0, 271}[1:],
	`The Prime Evils are no more.`:              []int{0, 2212}[1:],
	`The Prowler`:                               []int{0, 514}[1:],
	`The Reaper's Toll`:                         []int{0, 2780}[1:],
	`The Redeemer`:                              []int{0, 364}[1:],
	`The Rising Sun`:                            []int{0, 1933}[1:],
	`The Rogues' test is done.`:                 []int{0, 1184}[1:],
	`The Sanctuary - Horazon's obsession.`:      []int{0, 625}[1:],
	`The Scalper`:                               []int{0, 1162}[1:],
	`The Spirit Shroud`:                         []int{0, 1043}[1:],
	`The Swashbuckler`:                          []int{0, 1436}[1:],
	`The Treentster`:                            []int{0, 1224}[1:],
	`The Vicar`:                                 []int{0, 2803}[1:],
	`The Vile Husk`:                             []int{0, 2309}[1:],
	`The Wanderer's Blade`:                      []int{0, 2521}[1:],
	`The Worldstone Chamber`:                    []int{0, 548}[1:],
	`The Worldstone!`:                           []int{0, 2061}[1:],
	`The Worldstone! Praise the Light.`:         []int{0, 475}[1:],
	`The Worldstone. What power.`:               []int{0, 2064}[1:],
	`The catapults have been silenced.`:         []int{0, 1265}[1:],
	`The dark magic here is dispelled.`:         []int{0, 2734}[1:],
	`The fabled home of the Ancients.`:          []int{0, 380, 1676}[1:],
	`The guardians of Mount Arreat await.`:      []int{0, 2530}[1:],
	`The halls of the Ancients... Magnificent.`: []int{0, 2056}[1:],
	`The last of the Three has fallen.`:         []int{0, 1054}[1:],
	`The legendary Worldstone - guardian of the Natural realm.`: []int{0, 2250}[1:],
	`The power of the Worldstone washes over me.`:               []int{0, 222}[1:],
	`The reign of Terror has ended.`:                            []int{0, 2522}[1:],
	`The resting place of the Ancients...`:                      []int{0, 680}[1:],
	`The siege is broken.`:                                      []int{0, 2050}[1:],
	`The siege must be stopped.`:                                []int{0, 2316}[1:],
	`The summit - The Barbarian holy ground.`:                   []int{0, 1492}[1:],
	`The sun warms the world once more.`:                        []int{0, 2036}[1:],
	`The supernatural is strong here.`:                          []int{0, 2633}[1:],
	`The time has come to cleanse my homeland!`:                 []int{0, 65}[1:],
	`Theodoric`:                                         []int{0, 1281}[1:],
	`There is hope once again.`:                         []int{0, 669}[1:],
	`These Horadric markings are mysterious.`:           []int{0, 559}[1:],
	`These stones hold an ancient power.`:               []int{0, 2761}[1:],
	`They'll never see me coming.`:                      []int{0, 212}[1:],
	`Thin`:                                              []int{0, 371}[1:],
	`Thirst`:                                            []int{0, 2498}[1:],
	`This dead tree teems with energy.`:                 []int{0, 1458}[1:],
	`This is one drain I don't mind cleaning out.`:      []int{0, 2039}[1:],
	`This item cannot be repaired.`:                     []int{0, 835}[1:],
	`This place would drive anyone mad.`:                []int{0, 1170}[1:],
	`This smells worse than the sewers of Lut Gohlein.`: []int{0, 739}[1:],
	`This temple is a nest of evil.`:                    []int{0, 321}[1:],
	`This tower has its charms...`:                      []int{0, 1068}[1:],
	`This was not designed by nature's Architect.`:      []int{0, 356}[1:],
	`Thought`:                         []int{0, 2500}[1:],
	`Thresh Socket`:                   []int{0, 988}[1:],
	`Thresher`:                        []int{0, 1885}[1:],
	`Throne of Destruction`:           []int{0, 549}[1:],
	`Thul`:                            []int{0, 1103}[1:],
	`Thul Rune`:                       []int{0, 1839}[1:],
	`Thunder`:                         []int{0, 2501}[1:],
	`Thunder Claws`:                   []int{0, 2717}[1:],
	`Thunder Maul`:                    []int{0, 1742}[1:],
	`Thundergod's Vigor`:              []int{0, 1133}[1:],
	`Thunderstroke`:                   []int{0, 1115}[1:],
	`Thus ends the plague of Terror.`: []int{0, 2049}[1:],
	`Tiamat's Rebuke`:                 []int{0, 1926}[1:],
	`Tiara`:                           []int{0, 1703}[1:],
	`Tiger Strike`:                    []int{0, 2422, 2424}[1:],
	`Time`:                            []int{0, 2502}[1:],
	`Timeless`:                        []int{0, 1713}[1:],
	`Tin`:                             []int{0, 743}[1:],
	`Tir`:                             []int{0, 895}[1:],
	`Tir Rune`:                        []int{0, 1813}[1:],
	`Titan's Fist`:                    []int{0, 272}[1:],
	`Titan's Revenge`:                 []int{0, 591}[1:],
	`To Abaddon`:                      []int{0, 2575}[1:],
	`To Harrogath`:                    []int{0, 656}[1:],
	`To Nihlathak's Temple`:           []int{0, 577}[1:],
	`To The Ancients' Way`:            []int{0, 573}[1:],
	`To The Arreat Plateau`:           []int{0, 564}[1:],
	`To The Arreat Summit`:            []int{0, 1024}[1:],
	`To The Bloody Foothills`:         []int{0, 561}[1:],
	`To The Crystalline Passage`:      []int{0, 567}[1:],
	`To The Drifter Cavern`:           []int{0, 570}[1:],
	`To The Frigid Highlands`:         []int{0, 1238}[1:],
	`To The Frozen River`:             []int{0, 568}[1:],
	`To The Frozen Tundra`:            []int{0, 571}[1:],
	`To The Glacial Trail`:            []int{0, 569}[1:],
	`To The Halls of Anguish`:         []int{0, 578}[1:],
	`To The Halls of Pain`:            []int{0, 580}[1:],
	`To The Halls of Torment`:         []int{0, 581}[1:],
	`To The Halls of Vaught`:          []int{0, 584}[1:],
	`To The Icy Cellar`:               []int{0, 576}[1:],
	`To The Infernal Pit`:             []int{0, 2585}[1:],
	`To The Pit of Acheron`:           []int{0, 2582}[1:],
	`To The Throne of Destruction`:    []int{0, 681}[1:],
	`To The Worldstone Chamber`:       []int{0, 1007}[1:],
	`To The Worldstone Keep Level 1`:  []int{0, 589}[1:],
	`To The Worldstone Keep Level 2`:  []int{0, 595}[1:],
	`To The Worldstone Keep Level 3`:  []int{0, 596}[1:],
	`Todesfaelle Flamme`:              []int{0, 1539}[1:],
	`Toggle MiniMap`:                  []int{0, 712}[1:],
	`Tomahawk`:                        []int{0, 1746}[1:],
	`Tomb`:                            []int{0, 354, 359, 361, 401, 853, 909, 1004, 1005}[1:],
	`Tomb Reaver`:                     []int{0, 2016}[1:],
	`Toothrow`:                        []int{0, 828}[1:],
	`Torch`:                           []int{0, 352, 353, 363, 384, 476, 479, 1072, 1076}[1:],
	`Torch Pit`:                       []int{0, 392}[1:],
	`Torkel`:                          []int{0, 1267}[1:],
	`Torment Worm`:                    []int{0, 453}[1:],
	`Tornado`:                         []int{0, 1611, 2336}[1:],
	`Tostig`:                          []int{0, 1232}[1:],
	`Totemic Mask`:                    []int{0, 2150}[1:],
	`Town Flag`:                       []int{0, 327}[1:],
	`Toxic`:                           []int{0, 1805}[1:],
	`Tradition`:                       []int{0, 2503}[1:],
	`Trainer's`:                       []int{0, 815}[1:],
	`Trang-Oul's Avatar`:              []int{0, 1175}[1:],
	`Trang-Oul's Claws`:               []int{0, 347}[1:],
	`Trang-Oul's Girth`:               []int{0, 304}[1:],
	`Trang-Oul's Guise`:               []int{0, 1481}[1:],
	`Trang-Oul's Mask`:                []int{0, 607}[1:],
	`Trang-Oul's Scales`:              []int{0, 1}[1:],
	`Trang-Oul's Wing`:                []int{0, 1453}[1:],
	`Traps`:                           []int{0, 755}[1:],
	`Travel through the Ancient's Way to find the Ancients at the Arreat Summit.`: []int{0, 498}[1:],
	`Treachery`:       []int{0, 2504}[1:],
	`Triad's Foliage`: []int{0, 174}[1:],
	`Tribal Flag`:     []int{0, 326}[1:],
	`Trickster's`:     []int{0, 1551}[1:],
	`Tristram... The first to fall to Diablo's wrath.`: []int{0, 2525}[1:],
	`Troll Belt`:               []int{0, 899}[1:],
	`Troll Nest`:               []int{0, 888}[1:],
	`Truncheon`:                []int{0, 1677}[1:],
	`Trust`:                    []int{0, 2506}[1:],
	`Truth`:                    []int{0, 2383}[1:],
	`Try and cage me, demons.`: []int{0, 701}[1:],
	`Turquoise`:                []int{0, 1401}[1:],
	`Twister`:                  []int{0, 1521, 2330}[1:],
	`Tyrael's Mercy`:           []int{0, 1266}[1:],
	`Tyrael's Might`:           []int{0, 1098}[1:],
	`Tyrant Club`:              []int{0, 1938}[1:],
	`Ulf`:                      []int{0, 1198}[1:],
	`Ulric`:                    []int{0, 1203}[1:],
	`Um`:                       []int{0, 1390}[1:],
	`Um Rune`:                  []int{0, 1864}[1:],
	`Unable to access file. Cannot convert character.`:                                                    []int{0, 144}[1:],
	`Unable to enter game. Your character must kill Baal in Nightmare difficulty to play in a Hell game.`: []int{0, 122}[1:],
	`Unable to enter game. Your character must kill Baal to play in a Nightmare game.`:                    []int{0, 123}[1:],
	`Unbending Will`:  []int{0, 2384}[1:],
	`Unearthed Wand`:  []int{0, 1752}[1:],
	`Unearthly`:       []int{0, 951}[1:],
	`Unferth`:         []int{0, 1032}[1:],
	`Unholy Corpse`:   []int{0, 76}[1:],
	`Unraveler`:       []int{0, 1050}[1:],
	`Unraveller Head`: []int{0, 1637}[1:],
	`Use the potion Malah gave you to thaw Anya.`: []int{0, 480}[1:],
	`Uther`:                  []int{0, 1196}[1:],
	`Valkyrie Wing`:          []int{0, 1037}[1:],
	`Valor`:                  []int{0, 2385}[1:],
	`Vambraces`:              []int{0, 791}[1:],
	`Vampire Gaze`:           []int{0, 2251}[1:],
	`Vampirebone Gloves`:     []int{0, 935}[1:],
	`Vampirefang Belt`:       []int{0, 931}[1:],
	`Vardhaka`:               []int{0, 841}[1:],
	`Veil of Steel`:          []int{0, 1981}[1:],
	`Vengeance`:              []int{0, 2386}[1:],
	`Vengeance... for Atma.`: []int{0, 1484}[1:],
	`Venom`:                  []int{0, 2387, 2784, 2788}[1:],
	`Venom Grip`:             []int{0, 1511}[1:],
	`Venomous`:               []int{0, 410}[1:],
	`Ventar the Unholy`:      []int{0, 1665}[1:],
	`Verdungo's Hearty Cord`: []int{0, 2005}[1:],
	`Vermilion`:              []int{0, 1899}[1:],
	`Veteran's`:              []int{0, 2345}[1:],
	`Vex`:                    []int{0, 1454}[1:],
	`Vex Rune`:               []int{0, 1856}[1:],
	`Victorious`:             []int{0, 2673}[1:],
	`Victory`:                []int{0, 2388}[1:],
	`Vigorous`:               []int{0, 1552}[1:],
	`Vikhyat`:                []int{0, 910}[1:],
	`Vile Temptress`:         []int{0, 112}[1:],
	`Vile Witch`:             []int{0, 167}[1:],
	`Vine`:                   []int{0, 972}[1:],
	`Vine Creature`:          []int{0, 593}[1:],
	`Vinvear Molech`:         []int{0, 751}[1:],
	`Viperfork`:              []int{0, 106}[1:],
	`Visceratuant`:           []int{0, 274}[1:],
	`Visionary`:              []int{0, 642}[1:],
	`Voice`:                  []int{0, 2389}[1:],
	`Voice of Reason`:        []int{0, 1471}[1:],
	`Void`:                   []int{0, 2390}[1:],
	`Volcanic`:               []int{0, 2145}[1:],
	`Volcano`:                []int{0, 1606, 2335}[1:],
	`Vortex Orb`:             []int{0, 1895}[1:],
	`Vortex Shield`:          []int{0, 2136}[1:],
	`Wailing`:                []int{0, 897}[1:],
	`Wailing Spirit`:         []int{0, 977}[1:],
	`Wake of Fire`:           []int{0, 250, 2671, 2676}[1:],
	`Wake of Inferno`:        []int{0, 2729, 2732}[1:],
	`Walking Stick`:          []int{0, 1685}[1:],
	`Wallace's Tear`:         []int{0, 2662}[1:],
	`War`:                    []int{0, 2391}[1:],
	`War Fist`:               []int{0, 2007}[1:],
	`War Pike`:               []int{0, 1833}[1:],
	`War Spike`:              []int{0, 1841}[1:],
	`War Traveler`:           []int{0, 1442}[1:],
	`Ward`:                   []int{0, 916}[1:],
	`Ward Bow`:               []int{0, 1689}[1:],
	`Warden's`:               []int{0, 1652}[1:],
	`Warhound`:               []int{0, 1168}[1:],
	`Warlord's Trust`:        []int{0, 789}[1:],
	`Warning:  Once you convert a character to expansion, you cannot play it in original Diablo II games.`: []int{0, 809}[1:],
	`Warped Fallen Wolfrider`: []int{0, 670}[1:],
	`Warpspear`:               []int{0, 1289}[1:],
	`Warrior`:                 []int{0, 49}[1:],
	`Warriv's Warder`:         []int{0, 1927}[1:],
	`Warshrike`:               []int{0, 2651}[1:],
	`Water`:                   []int{0, 2392}[1:],
	`Waterwalk`:               []int{0, 852}[1:],
	`Waypoint`:                []int{0, 990}[1:],
	`Wealth`:                  []int{0, 2399}[1:],
	`Weapon Block`:            []int{0, 2679}[1:],
	`Weapon Rack`:             []int{0, 431, 1226}[1:],
	`Weapons:`:                []int{0, 1022}[1:],
	`Weder`:                   []int{0, 893}[1:],
	`Well`:                    []int{0, 311}[1:],
	`Weohstan`:                []int{0, 1116}[1:],
	`Werebear`:                []int{0, 1505, 2256}[1:],
	`Werewolf`:                []int{0, 1494, 2218}[1:],
	`What I kill stays dead.`: []int{0, 1219}[1:],
	`What a delicious blade! I should consult Ormus.`: []int{0, 2812}[1:],
	`When You Get Hit`:                       []int{0, 2273}[1:],
	`When You Hit an Enemy`:                  []int{0, 2275}[1:],
	`When You Swing`:                         []int{0, 2269}[1:],
	`Whichwild String`:                       []int{0, 2160}[1:],
	`Whisper`:                                []int{0, 2400}[1:],
	`White`:                                  []int{0, 2401}[1:],
	`Whitstan's Guard`:                       []int{0, 478}[1:],
	`Who would want to remember this place?`: []int{0, 178}[1:],
	`Whose handiwork lies buried here?`:      []int{0, 719}[1:],
	`Why must evil hide in such wretched places?`: []int{0, 554}[1:],
	`Wicked`:                  []int{0, 1041}[1:],
	`Widowmaker`:              []int{0, 545, 1928}[1:],
	`Wiglaf`:                  []int{0, 861}[1:],
	`Wilhelm's Pride`:         []int{0, 2290}[1:],
	`Willhelm's Pride`:        []int{0, 2322}[1:],
	`Wind`:                    []int{0, 2402}[1:],
	`Windforce`:               []int{0, 2178}[1:],
	`Windhammer`:              []int{0, 2046}[1:],
	`Windstrike`:              []int{0, 1700}[1:],
	`Winged Axe`:              []int{0, 1609}[1:],
	`Winged Harpoon`:          []int{0, 1956}[1:],
	`Winged Knife`:            []int{0, 1660}[1:],
	`Wings of Hope`:           []int{0, 2403}[1:],
	`Wire Fleece`:             []int{0, 917}[1:],
	`Wisdom`:                  []int{0, 2404}[1:],
	`Wisp Projector`:          []int{0, 1800}[1:],
	`Witch-hunter's`:          []int{0, 1205}[1:],
	`Wizardspike`:             []int{0, 99}[1:],
	`Woe`:                     []int{0, 2406}[1:],
	`Wolf`:                    []int{0, 30, 66, 396}[1:],
	`Wolf Defense: `:          []int{0, 2431}[1:],
	`Wolf Head`:               []int{0, 2102}[1:],
	`Wolf: `:                  []int{0, 627}[1:],
	`Wolfhowl`:                []int{0, 2441}[1:],
	`Wolverine Hrt`:           []int{0, 2317}[1:],
	`Wolves: `:                []int{0, 2802}[1:],
	`Wonder`:                  []int{0, 2408}[1:],
	`Wooden Chest`:            []int{0, 308, 313, 546, 560}[1:],
	`Worldstone Keep Level 1`: []int{0, 1217}[1:],
	`Worldstone Keep Level 2`: []int{0, 1218}[1:],
	`Worldstone Keep Level 3`: []int{0, 1220}[1:],
	`Wpn Block`:               []int{0, 2684}[1:],
	`Wraith Flight`:           []int{0, 55}[1:],
	`Wraithfang`:              []int{0, 2346}[1:],
	`Wraps`:                   []int{0, 1350}[1:],
	`Wrath`:                   []int{0, 2409}[1:],
	`Wrath of Cain`:           []int{0, 2365}[1:],
	`Wretched Defiler`:        []int{0, 231}[1:],
	`Wrist Blade`:             []int{0, 1377}[1:],
	`Wrist Spike`:             []int{0, 2499}[1:],
	`Wrist Sword`:             []int{0, 1987}[1:],
	`Wulf`:                    []int{0, 1118}[1:],
	`Wulfgar`:                 []int{0, 874}[1:],
	`Wulfstan`:                []int{0, 1272}[1:],
	`Wyrmhide`:                []int{0, 657}[1:],
	`Wyrmhide Boots`:          []int{0, 770}[1:],
	`Wyvern's Head`:           []int{0, 1138}[1:],
	`Xenos`:                   []int{0, 660}[1:],
	`Yelling`:                 []int{0, 2507}[1:],
	`You Dark Mages are all alike - obsessed with power.`: []int{0, 962}[1:],
	`You were a sad little man, Nihlathak.`:               []int{0, 1227}[1:],
	`You'll pay for your atrocities, Baal.`:               []int{0, 33}[1:],
	`Your power was no match for mine.`:                   []int{0, 2681}[1:],
	`Your reign is over, Andariel.`:                       []int{0, 2220}[1:],
	`Your time is past, Blood Raven.`:                     []int{0, 437}[1:],
	`Youth`:                                               []int{0, 2410}[1:],
	`Yrmenlaf`:                                            []int{0, 1063}[1:],
	`Zakarum Shield`:                                      []int{0, 2135}[1:],
	`Zakarum's Hand`:                                      []int{0, 508, 2666}[1:],
	`Zakarum's Salvation`:                                 []int{0, 259}[1:],
	`Zephyr`:                                              []int{0, 2415}[1:],
	`Zircon`:                                              []int{0, 1128}[1:],
	`Zod`:                                                 []int{0, 1662}[1:],
	`Zod Rune`:                                            []int{0, 1861}[1:],
	`Zombie Head`:                                         []int{0, 1636}[1:],
	`a mind blast to crush your enemies`:                  []int{0, 2630}[1:],
	`a trap that emits charged bolts`:                     []int{0, 2659}[1:],
	`a trap that emits lightning`:                         []int{0, 2726}[1:],
	`a trap that emits waves of fire`:                     []int{0, 2672, 2675}[1:],
	`add poison damage to your weapons`:                   []int{0, 2786}[1:],
	`add sockets`:                                         []int{0, 837}[1:],
	`adds charged-up bonuses to both claw attacks
with your dual claw-class weapons
slice and dice your enemies

Finishing Move
`: []int{0, 2656}[1:],
	`adds charged-up bonuses to the kick
and destroy them with a kick
teleport to your enemies

Finishing Move
`: []int{0, 2765}[1:],
	`adds charged-up bonuses to the kick
kick your enemies out of your way

Finishing Move
`: []int{0, 2427}[1:],
	`adds charged-up bonuses to the kick
knock back your enemies with an explosive kick

Finishing Move
`: []int{0, 2723}[1:],
	`adds cold damage to finishing moves

Charge-up Skill`: []int{0, 2753}[1:],
	`adds elemental novas to finishing moves

Charge-up Skill`: []int{0, 2795}[1:],
	`adds fire damage to finishing moves

Charge-up Skill`: []int{0, 2650}[1:],
	`adds lightning damage to finishing moves
 
Charge-up Skill`: []int{0, 2714}[1:],
	`and convert the feeble-minded
stun a group of enemies
using the power of your mind`: []int{0, 2746}[1:],
	`and destruction over your enemies
summon forth a volcano to rain death`: []int{0, 2070}[1:],
	`and replenishes your life
summon a vine that eats corpses`: []int{0, 1786}[1:],
	`and replenishes your mana
summon a vine that eats corpses`: []int{0, 1803}[1:],
	`at enemies that pass near
a trap that emits charged bolts`: []int{0, 2669}[1:],
	`back at your enemies
taken by you and your party
summon spirit pet that reflects damage`: []int{0, 2077}[1:],
	`between you and target point
set a spinning blade to patrol`: []int{0, 2640}[1:],
	`bite causes disease - werewolf form`: []int{0, 2597}[1:],
	`blind your enemies`:                  []int{0, 2686}[1:],
	`burning damage`:                      []int{0, 18}[1:],
	`burning duration: `:                  []int{0, 23}[1:],
	`burning them to a crisp
open volcanic vents below your enemies,`: []int{0, 1792}[1:],
	`chain lightning damage: `:           []int{0, 1140}[1:],
	`chaos ice bolt damage: `:            []int{0, 1141}[1:],
	`charged bolt damage: `:              []int{0, 421}[1:],
	`cold damage radius: `:               []int{0, 27}[1:],
	`cold damage: `:                      []int{0, 25}[1:],
	`compelling psionic blast`:           []int{0, 2743}[1:],
	`create a massive wind storm`:        []int{0, 2621}[1:],
	`create a tornado`:                   []int{0, 2611}[1:],
	`create a volcano`:                   []int{0, 2610}[1:],
	`create shock waves - werebear form`: []int{0, 2607}[1:],
	`cut a path through your enemies
release several small whirlwinds that`: []int{0, 1802}[1:],
	`debris to pound your enemies to bits
create a massive storm of wind and`: []int{0, 2368}[1:],
	`destruction on nearby enemies
create a meteor shower to rain fiery`: []int{0, 2362}[1:],
	`disease to all it contacts
summon a vine that spreads`: []int{0, 1771}[1:],
	`double claw attack

Finishing Move`: []int{0, 2655}[1:],
	`explosive kick

Finishing Move`: []int{0, 2719}[1:],
	`fiery, mauling attack`: []int{0, 2598}[1:],
	`fire damage radius: `:  []int{0, 22}[1:],
	`fire damage: `:         []int{0, 19}[1:],
	`fire, cold, and lightning
shield yourself from damage caused by`: []int{0, 1793}[1:],
	`for a period of time
increases attack and movement speed`: []int{0, 2647}[1:],
	`for a period of time
raise all resistances and resist curses`: []int{0, 2707}[1:],
	`freeze duration: `:                   []int{0, 28}[1:],
	`go north`:                            []int{0, 818}[1:],
	`increases attack and movement speed`: []int{0, 2646}[1:],
	`increases damage of finishing moves

Charge-up Skill`: []int{0, 2634}[1:],
	`it does to enemies
eating corpses to increase damage
summon a wolf that becomes enraged,`: []int{0, 1795}[1:],
	`kick your enemies

Finishing Move`: []int{0, 2635}[1:],
	`launch a molten boulder`: []int{0, 2542}[1:],
	`life for you and your party
summon a spirit pet that increases`: []int{0, 1777}[1:],
	`life-and-mana-stealing bite`:               []int{0, 2603}[1:],
	`life-stealing rage attack - werewolf form`: []int{0, 2550}[1:],
	`lightning damage: `:                        []int{0, 418}[1:],
	`lowering their defenses for a period of time
cast a shadow to blind nearby enemies`: []int{0, 2687}[1:],
	`maul your enemies - werebear form`: []int{0, 2551}[1:],
	`meteor damage: `:                   []int{0, 1139}[1:],
	`multiple attacks - werewolf Form`:  []int{0, 2616}[1:],
	`normal attack to release charges
must use a dragon finishing move or
adds elemental novas to finishing moves

Charge-up Skill
`: []int{0, 2796}[1:],
	`normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add cold damage

Charge-up Skill
`: []int{0, 2757}[1:],
	`normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add fire damage

Charge-up Skill
`: []int{0, 2652}[1:],
	`normal attack to release charges
must use a dragon finishing move or
can only be used with claw-class weapons
to finishing moves
consecutive hits add lightning damage

Charge-up Skill
`: []int{0, 2716}[1:],
	`normal attack to release charges
must use a dragon finishing move or
to finishing moves
consecutive hits add damage bonuses

Charge-up Skill
`: []int{0, 2423}[1:],
	`normal attack to release charges
must use a dragon finishing move or
to finishing moves
consecutive hits add life and mana stealing

Charge-up Skill
`: []int{0, 2693}[1:],
	`nova damage: `:          []int{0, 420}[1:],
	`of Abrasion`:            []int{0, 1900}[1:],
	`of Acceleration`:        []int{0, 1618}[1:],
	`of Admiration`:          []int{0, 2041}[1:],
	`of Ages`:                []int{0, 1508}[1:],
	`of Amianthus`:           []int{0, 37}[1:],
	`of Amicae`:              []int{0, 958}[1:],
	`of Amplify Damage`:      []int{0, 429, 1096}[1:],
	`of Anthrax`:             []int{0, 1414}[1:],
	`of Armageddon`:          []int{0, 2241}[1:],
	`of Atlas`:               []int{0, 210}[1:],
	`of Attract`:             []int{0, 1721}[1:],
	`of Attrition`:           []int{0, 1042}[1:],
	`of Avarice`:             []int{0, 2117}[1:],
	`of Badness`:             []int{0, 1293}[1:],
	`of Battle Command`:      []int{0, 654}[1:],
	`of Battle Cry`:          []int{0, 173}[1:],
	`of Battle Orders`:       []int{0, 360}[1:],
	`of Beauty`:              []int{0, 295}[1:],
	`of Bile`:                []int{0, 674}[1:],
	`of Blaze`:               []int{0, 886}[1:],
	`of Blazing`:             []int{0, 555}[1:],
	`of Blessed Hammer`:      []int{0, 1787}[1:],
	`of Bliss`:               []int{0, 14}[1:],
	`of Blitzen`:             []int{0, 155}[1:],
	`of Blizzard`:            []int{0, 378, 611}[1:],
	`of Blood Golem`:         []int{0, 2754}[1:],
	`of Bone Armor`:          []int{0, 56}[1:],
	`of Bone Prison`:         []int{0, 1967}[1:],
	`of Bone Spear`:          []int{0, 1201}[1:],
	`of Bone Spirit`:         []int{0, 869}[1:],
	`of Bone Wall`:           []int{0, 38}[1:],
	`of Butchery`:            []int{0, 2066}[1:],
	`of Chain Lightning`:     []int{0, 1648}[1:],
	`of Charge`:              []int{0, 2412}[1:],
	`of Charged Bolt`:        []int{0, 1330, 2434}[1:],
	`of Charged Spear`:       []int{0, 969}[1:],
	`of Charged Strike`:      []int{0, 1809}[1:],
	`of Chilling Armor`:      []int{0, 2243}[1:],
	`of Clay Golem`:          []int{0, 536}[1:],
	`of Cold Arrow`:          []int{0, 645}[1:],
	`of Combat`:              []int{0, 2742}[1:],
	`of Concentration`:       []int{0, 1653}[1:],
	`of Confusion`:           []int{0, 1750}[1:],
	`of Conversion`:          []int{0, 1991}[1:],
	`of Coolness`:            []int{0, 2583}[1:],
	`of Corpse Explosion`:    []int{0, 1902}[1:],
	`of Credit`:              []int{0, 1158}[1:],
	`of Cremation`:           []int{0, 2642}[1:],
	`of Cruelty`:             []int{0, 2363}[1:],
	`of Cyclone Armor`:       []int{0, 981}[1:],
	`of Daring`:              []int{0, 1222, 1488}[1:],
	`of Darkness`:            []int{0, 2168}[1:],
	`of Dawn`:                []int{0, 1204, 1236}[1:],
	`of Decrepify`:           []int{0, 1241}[1:],
	`of Desire`:              []int{0, 2287}[1:],
	`of Dim Vision`:          []int{0, 1009}[1:],
	`of Disease`:             []int{0, 1872}[1:],
	`of Dispatch`:            []int{0, 1522}[1:],
	`of Doom`:                []int{0, 1901}[1:],
	`of Dread`:               []int{0, 929}[1:],
	`of Dusk`:                []int{0, 747}[1:],
	`of Eluding`:             []int{0, 116}[1:],
	`of Enchant`:             []int{0, 1460}[1:],
	`of Enchantment`:         []int{0, 1329}[1:],
	`of Energy Shield`:       []int{0, 1040}[1:],
	`of Enlightenment`:       []int{0, 1111}[1:],
	`of Ennui`:               []int{0, 556}[1:],
	`of Envy`:                []int{0, 198}[1:],
	`of Erosion`:             []int{0, 1008}[1:],
	`of Evisceration`:        []int{0, 598}[1:],
	`of Exploding Arrow`:     []int{0, 1144}[1:],
	`of Faith`:               []int{0, 44}[1:],
	`of Fending`:             []int{0, 1722}[1:],
	`of Fervor`:              []int{0, 2327}[1:],
	`of Find Item`:           []int{0, 1594}[1:],
	`of Find Potion`:         []int{0, 714}[1:],
	`of Fire Arrow`:          []int{0, 840}[1:],
	`of Fire Ball`:           []int{0, 1489, 2436}[1:],
	`of Fire Bolt`:           []int{0, 455}[1:],
	`of Fire Golem`:          []int{0, 876}[1:],
	`of Fire Wall`:           []int{0, 248, 1177}[1:],
	`of Firebolts`:           []int{0, 563}[1:],
	`of Firestorms`:          []int{0, 2307}[1:],
	`of Fissure`:             []int{0, 1417}[1:],
	`of Fist of the Heavens`: []int{0, 1132}[1:],
	`of Fortification`:       []int{0, 61}[1:],
	`of Freedom`:             []int{0, 1556}[1:],
	`of Freezing Arrow`:      []int{0, 2026}[1:],
	`of Frigidity`:           []int{0, 827}[1:],
	`of Frost Nova`:          []int{0, 432}[1:],
	`of Frost Shield`:        []int{0, 2674}[1:],
	`of Frozen Armor`:        []int{0, 1405}[1:],
	`of Frozen Orb`:          []int{0, 93, 565}[1:],
	`of Glacial Spike`:       []int{0, 1724, 2329}[1:],
	`of Good Luck`:           []int{0, 2242}[1:],
	`of Grace`:               []int{0, 2417}[1:],
	`of Grace and Power`:     []int{0, 1413}[1:],
	`of Grim Ward`:           []int{0, 968}[1:],
	`of Grounding`:           []int{0, 1474, 1589}[1:],
	`of Guided Arrow`:        []int{0, 214}[1:],
	`of Holy Bolt`:           []int{0, 127}[1:],
	`of Holy Shield`:         []int{0, 1790}[1:],
	`of Honor`:               []int{0, 932, 1440}[1:],
	`of Hope`:                []int{0, 1317}[1:],
	`of Horror`:              []int{0, 1965}[1:],
	`of Howl`:                []int{0, 1698}[1:],
	`of Hurricane`:           []int{0, 2127}[1:],
	`of Hydra`:               []int{0, 2139}[1:],
	`of Hydra Shield`:        []int{0, 1634}[1:],
	`of Ice Arrow`:           []int{0, 2165}[1:],
	`of Ice Blast`:           []int{0, 1242, 1656}[1:],
	`of Ice Bolt`:            []int{0, 631}[1:],
	`of Icebolt`:             []int{0, 1433}[1:],
	`of Illusion`:            []int{0, 2691}[1:],
	`of Immolating Arrow`:    []int{0, 582}[1:],
	`of Impaling Spear`:      []int{0, 2306}[1:],
	`of Impaling Strike`:     []int{0, 1085}[1:],
	`of Incineration`:        []int{0, 952}[1:],
	`of Inertia`:             []int{0, 767}[1:],
	`of Inflammability`:      []int{0, 2065}[1:],
	`of Inner Sight`:         []int{0, 2262}[1:],
	`of Insulation`:          []int{0, 1869}[1:],
	`of Ire`:                 []int{0, 1346}[1:],
	`of Iron Golem`:          []int{0, 1284}[1:],
	`of Iron Maiden`:         []int{0, 1049}[1:],
	`of Jab`:                 []int{0, 324}[1:],
	`of Joy`:                 []int{0, 1588}[1:],
	`of Joyfulness`:          []int{0, 2428}[1:],
	`of Judgement`:           []int{0, 1962}[1:],
	`of Karma`:               []int{0, 2360}[1:],
	`of Knowledge`:           []int{0, 2198}[1:],
	`of Life Everlasting`:    []int{0, 1718}[1:],
	`of Life Tap`:            []int{0, 1726}[1:],
	`of Lightning Fury`:      []int{0, 551}[1:],
	`of Lightning Javelin`:   []int{0, 521}[1:],
	`of Lightning Spear`:     []int{0, 2308}[1:],
	`of Lightning Strike`:    []int{0, 1518}[1:],
	`of Love`:                []int{0, 920}[1:],
	`of Lower Resistance`:    []int{0, 2312}[1:],
	`of Luck`:                []int{0, 2249, 2281}[1:],
	`of Maggots`:             []int{0, 438}[1:],
	`of Magic Arrow`:         []int{0, 162}[1:],
	`of Malice`:              []int{0, 1287}[1:],
	`of Memory`:              []int{0, 2292}[1:],
	`of Meteor`:              []int{0, 1131, 1697}[1:],
	`of Molten Boulder`:      []int{0, 566}[1:],
	`of Multiple Shot`:       []int{0, 659}[1:],
	`of Nirvana`:             []int{0, 1120}[1:],
	`of Nobility`:            []int{0, 599}[1:],
	`of Nova`:                []int{0, 679, 1968}[1:],
	`of Nova Shield`:         []int{0, 2092}[1:],
	`of Passion`:             []int{0, 2810}[1:],
	`of Pilfering`:           []int{0, 2183}[1:],
	`of Plague Jab`:          []int{0, 2030}[1:],
	`of Plague Javelin`:      []int{0, 1620}[1:],
	`of Poison Dagger`:       []int{0, 1033}[1:],
	`of Poison Explosion`:    []int{0, 585}[1:],
	`of Poison Jab`:          []int{0, 994}[1:],
	`of Poison Javelin`:      []int{0, 132}[1:],
	`of Poison Nova`:         []int{0, 1654}[1:],
	`of Power`:               []int{0, 207}[1:],
	`of Power Spear`:         []int{0, 157}[1:],
	`of Power Strike`:        []int{0, 1176}[1:],
	`of Propogation`:         []int{0, 1134}[1:],
	`of Prosperity`:          []int{0, 191}[1:],
	`of Quenching`:           []int{0, 2364}[1:],
	`of Quickening`:          []int{0, 1719}[1:],
	`of Raise Skeleton`:      []int{0, 783}[1:],
	`of Razors`:              []int{0, 742}[1:],
	`of Redemption`:          []int{0, 1961}[1:],
	`of Remorse`:             []int{0, 1935}[1:],
	`of Replenishing`:        []int{0, 537}[1:],
	`of Resistance`:          []int{0, 153}[1:],
	`of Restoration`:         []int{0, 505}[1:],
	`of Revivification`:      []int{0, 1104}[1:],
	`of Sacrifice`:           []int{0, 639}[1:],
	`of Searing`:             []int{0, 2223}[1:],
	`of Self-Repair`:         []int{0, 1343}[1:],
	`of Shiver Armor`:        []int{0, 140}[1:],
	`of Shout`:               []int{0, 982}[1:],
	`of Skeletal Mages`:      []int{0, 2093}[1:],
	`of Slow Missile`:        []int{0, 1994}[1:],
	`of Spirit`:              []int{0, 812}[1:],
	`of Static Field`:        []int{0, 1451}[1:],
	`of Stature`:             []int{0, 1288}[1:],
	`of Stoicism`:            []int{0, 148}[1:],
	`of Stone`:               []int{0, 622}[1:],
	`of Storms`:              []int{0, 1649}[1:],
	`of Stun`:                []int{0, 2381}[1:],
	`of Suffering`:           []int{0, 1475}[1:],
	`of Sunlight`:            []int{0, 752}[1:],
	`of Sustenance`:          []int{0, 1311}[1:],
	`of Sweetness`:           []int{0, 504}[1:],
	`of Swords`:              []int{0, 525}[1:],
	`of Taunt`:               []int{0, 1087}[1:],
	`of Teeth`:               []int{0, 590}[1:],
	`of Telekinesis`:         []int{0, 1916}[1:],
	`of Teleport Shield`:     []int{0, 1554}[1:],
	`of Teleportation`:       []int{0, 2678}[1:],
	`of Terror`:              []int{0, 2063, 2328}[1:],
	`of Thawing`:             []int{0, 1127}[1:],
	`of Thunder Storm`:       []int{0, 2569}[1:],
	`of Tornado`:             []int{0, 2631}[1:],
	`of Transcendence`:       []int{0, 2069}[1:],
	`of Transportation`:      []int{0, 1587}[1:],
	`of Tribute`:             []int{0, 1018}[1:],
	`of Truth`:               []int{0, 1568}[1:],
	`of Twilight`:            []int{0, 1441}[1:],
	`of Twister`:             []int{0, 1591}[1:],
	`of Valhalla`:            []int{0, 1061}[1:],
	`of Vengeance`:           []int{0, 91}[1:],
	`of Vengence`:            []int{0, 1744}[1:],
	`of Vines`:               []int{0, 74}[1:],
	`of Virility`:            []int{0, 2008}[1:],
	`of Vita`:                []int{0, 934}[1:],
	`of Volcano`:             []int{0, 440}[1:],
	`of War Cry`:             []int{0, 1130}[1:],
	`of Warming`:             []int{0, 1046}[1:],
	`of Waste`:               []int{0, 2695}[1:],
	`of Weaken`:              []int{0, 1496}[1:],
	`of Winter`:              []int{0, 632}[1:],
	`of Wrath`:               []int{0, 579}[1:],
	`of Zeal`:                []int{0, 1745}[1:],
	`of the Abyss`:           []int{0, 535}[1:],
	`of the Avenger`:         []int{0, 293}[1:],
	`of the Bayou`:           []int{0, 875}[1:],
	`of the Beast`:           []int{0, 1592}[1:],
	`of the Centaur`:         []int{0, 342}[1:],
	`of the Cobra`:           []int{0, 1906}[1:],
	`of the Colossus`:        []int{0, 2261}[1:],
	`of the Dunes`:           []int{0, 2508}[1:],
	`of the Dynamo`:          []int{0, 2494}[1:],
	`of the Earth`:           []int{0, 2142}[1:],
	`of the Efreeti`:         []int{0, 67}[1:],
	`of the Elements`:        []int{0, 1221}[1:],
	`of the Elephant`:        []int{0, 457}[1:],
	`of the Fly`:             []int{0, 2115}[1:],
	`of the Forest`:          []int{0, 1966}[1:],
	`of the Gargantuan`:      []int{0, 2067}[1:],
	`of the Ghost`:           []int{0, 2274}[1:],
	`of the Gnat`:            []int{0, 2474}[1:],
	`of the Grassy Gnoll`:    []int{0, 2462}[1:],
	`of the Hag`:             []int{0, 2305}[1:],
	`of the Horde`:           []int{0, 900}[1:],
	`of the Icicle`:          []int{0, 2045}[1:],
	`of the Idiot`:           []int{0, 394}[1:],
	`of the Imbecile`:        []int{0, 1416}[1:],
	`of the Infantry`:        []int{0, 2272}[1:],
	`of the Jujube`:          []int{0, 1075}[1:],
	`of the Kraken`:          []int{0, 1870}[1:],
	`of the Lady`:            []int{0, 1873}[1:],
	`of the Lake`:            []int{0, 2270}[1:],
	`of the Lamprey`:         []int{0, 292}[1:],
	`of the Lilly`:           []int{0, 963}[1:],
	`of the Moon`:            []int{0, 2590}[1:],
	`of the Moron`:           []int{0, 933}[1:],
	`of the Mosquito`:        []int{0, 2020}[1:],
	`of the Obscene`:         []int{0, 192}[1:],
	`of the Ocean`:           []int{0, 1699}[1:],
	`of the Phoenix`:         []int{0, 727}[1:],
	`of the Plague`:          []int{0, 1964, 2033}[1:],
	`of the Quota`:           []int{0, 1233}[1:],
	`of the River`:           []int{0, 2626}[1:],
	`of the Scirocco`:        []int{0, 925}[1:],
	`of the Sky`:             []int{0, 87}[1:],
	`of the Specter`:         []int{0, 2194}[1:],
	`of the Squid`:           []int{0, 1097}[1:],
	`of the Stars`:           []int{0, 1590}[1:],
	`of the Stream`:          []int{0, 856}[1:],
	`of the Unicorn`:         []int{0, 257}[1:],
	`of the Virgin`:          []int{0, 2339}[1:],
	`of the Walrus`:          []int{0, 2475}[1:],
	`of the Whale`:           []int{0, 1871}[1:],
	`of the Witch`:           []int{0, 1890}[1:],
	`of the Yeti`:            []int{0, 2432}[1:],
	`of you and your party
to the damage and attack rating
summon a spirit pet that adds`: []int{0, 1794}[1:],
	`open the earth to burn enemies`: []int{0, 2571}[1:],
	`or explodes nearby corpses laying waste to more enemies
trap that shoots lightning at your enemies`: []int{0, 2776}[1:],
	`or one target multiple times
either multiple adjacent targets
when in werewolf form, attack`: []int{0, 2353}[1:],
	`passive - block with two claw-class weapons`:   []int{0, 2682}[1:],
	`passive - improves claw-class weapons ability`: []int{0, 2624}[1:],
	`passive - improves shape-shifting ability`:     []int{0, 2445}[1:],
	`personalize`:                      []int{0, 1253}[1:],
	`poison your weapon`:               []int{0, 2785}[1:],
	`rain fire on your enemies`:        []int{0, 2617}[1:],
	`release several small whirlwinds`: []int{0, 2601}[1:],
	`return to hell`:                   []int{0, 746}[1:],
	`secret object`:                    []int{0, 372}[1:],
	`set a spinning blade`:             []int{0, 2639}[1:],
	`shield from elemental damage`:     []int{0, 2588}[1:],
	`shoot a jet of ice`:               []int{0, 2544}[1:],
	`spinning blades of defense`:       []int{0, 2781}[1:],
	`summon a ferocious grizzly bear`:  []int{0, 2340}[1:],
	`summon a grizzly bear`:            []int{0, 2615}[1:],
	`summon a shadow of yourself`:      []int{0, 2710}[1:],
	`summon a spirit pet of thorns`:    []int{0, 2612}[1:],
	`summon a wolf`:                    []int{0, 2538}[1:],
	`summon a wolverine spirit`:        []int{0, 2589}[1:],
	`summon an enraged wolf`:           []int{0, 2592}[1:],
	`summon corpse eating vine`:        []int{0, 2547, 2602}[1:],
	`summon disease spreading vine`:    []int{0, 2442}[1:],
	`summon ravens`:                    []int{0, 2439}[1:],
	`summon the spirit of the oak`:     []int{0, 2537}[1:],
	`summon your shadow`:               []int{0, 2790}[1:],
	`teleport and kick enemies

Finishing Move`: []int{0, 2764}[1:],
	`temporary resist all`: []int{0, 2706}[1:],
	`that knocks back your enemies
launch a boulder of flaming hot magma`: []int{0, 1782}[1:],
	`that spreads to other monsters
to inflict them with disease
bite your enemies
when in werewolf form,`: []int{0, 1796}[1:],
	`that stuns nearby enemies
stomp to create a shock wave
when in werebear form,`: []int{0, 1808}[1:],
	`the eyes of your enemies
summon ravens to peck out`: []int{0, 1770}[1:],
	`throw a fire bomb`:        []int{0, 2623}[1:],
	`throw a web of lightning`: []int{0, 2636}[1:],
	`throw spinning blades`:    []int{0, 2698}[1:],
	`to Assassin Skills`:       []int{0, 2096}[1:],
	`to Druid Skills`:          []int{0, 2095}[1:],
	`to blast your enemies
create a funnel of wind and debris`: []int{0, 2072}[1:],
	`to blast your enemies to bits
throw a fire bomb`: []int{0, 2374}[1:],
	`to burn your enemies with frost
blast a continuous jet of ice`: []int{0, 1783}[1:],
	`to crush and knock back your enemies
to create a psychic blast
use the power of your mind`: []int{0, 2380}[1:],
	`to fight by your side
summon a powerful shadow of yourself`: []int{0, 2792}[1:],
	`to fight by your side
summon a wolf with teleporting ability`: []int{0, 1780}[1:],
	`to gain life and mana
form, bite your enemies
when in werewolf or werebear`: []int{0, 1807}[1:],
	`to scorch passing enemies
a trap that shoots lightning`: []int{0, 2727}[1:],
	`to shock your enemies
throw a web of lightning`: []int{0, 2435}[1:],
	`to slice up your enemies
throw spinning blades`: []int{0, 2699}[1:],
	`transform into a werebear`:                []int{0, 1781, 2539}[1:],
	`transform into a werewolf`:                []int{0, 1774, 2443}[1:],
	`trap that explodes nearby corpses`:        []int{0, 2774}[1:],
	`trap that sprays fire`:                    []int{0, 2730}[1:],
	`trap that sprays fire at passing enemies`: []int{0, 2731}[1:],
	`travel to harrogath`:                      []int{0, 1211}[1:],
	`unleash fiery chaos`:                      []int{0, 2446}[1:],
	`unleash fiery chaos to burn your enemies`: []int{0, 1776}[1:],
	`when in werewolf or werebear form
passive - improves duration and life`: []int{0, 1775}[1:],
	`who stray too close
spinning blades slice enemies`: []int{0, 2782}[1:],
	`with a fiery claw attack
form, maul your enemies
when in werewolf or werebear`: []int{0, 1799}[1:],
	`with claw-class weapons
passive - improves your skill`: []int{0, 2377}[1:],
	`with successive hits
for increasing extra damage
maul your enemies
when in werebear form,`: []int{0, 1791}[1:],
	`with successive hits
increasing amounts of life from your enemies
go into a frenzied rage to steal
when in werewolf form,`: []int{0, 1789}[1:],
	`you are using dual claw-class weapons
passive - chance to block when`: []int{0, 2683}[1:],
	`your skills and fights by your side
summon a shadow of yourself that mimics`: []int{0, 2711}[1:],
	`Â© Copyright 2001 Blizzard Entertainment`: []int{0, 151}[1:],
}
