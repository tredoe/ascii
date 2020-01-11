// Copyright 2014 Jonas mg
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ascii

// GetTable returns all characters from a list of ranges.
func GetTable(table ...[]Range) []byte {
	chars := make([]int, 0)

	for _, eachTable := range table {
		for _, hex := range eachTable {
			for char := hex.Lo; char <= hex.Hi; char++ {
				chars = append(chars, int(char))
			}
		}
	}

	// They are copied to a new variables with the exact length to save memory.
	chars2 := make([]byte, len(chars))
	for i, char := range chars {
		chars2[i] = byte(char)
	}

	return chars2
}
