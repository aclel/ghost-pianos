package main

var Sequences = map[int][]uint8{
	60: []uint8{
		KeyInt("AB", 1), KeyInt("C", 2), KeyInt("D", 3), KeyInt("EB", 4), KeyInt("G", 5), KeyInt("C", 6), KeyInt("D", 7),
		KeyInt("AB", 7), KeyInt("C", 6), KeyInt("D", 5), KeyInt("EB", 4), KeyInt("G", 3), KeyInt("C", 2), KeyInt("D", 1),
		KeyInt("EB", 1), KeyInt("G", 2), KeyInt("BB", 3), KeyInt("D", 4), KeyInt("F", 5), KeyInt("A", 6), KeyInt("D", 7),
		KeyInt("EB", 7), KeyInt("G", 6), KeyInt("BB", 5), KeyInt("D", 4), KeyInt("F", 3), KeyInt("A", 2), KeyInt("D", 1),
	},
	62: []uint8{KeyInt("EB", 3), KeyInt("G", 3), KeyInt("BB", 3), KeyInt("D", 3), KeyInt("F", 3), KeyInt("A", 3), KeyInt("D", 3)},
}
