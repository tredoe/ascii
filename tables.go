// Copyright 2014 Jonas mg
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ascii

// Range representates a range of ASCII characters.
// The range runs from Lo to Hi inclusive.
type Range struct {
	Lo byte
	Hi byte
}

var (
	// Digit is the set of ASCII digits.
	Digit = []Range{
		Range{0x30, 0x39},
	}

	// Upper is the set of ASCII upper case letters.
	Upper = []Range{
		Range{0x41, 0x5a},
	}

	// Lower is the set of ASCII lower case letters.
	Lower = []Range{
		Range{0x61, 0x7a},
	}

	// Letter is the set of ASCII letters.
	Letter = []Range{
		Range{0x41, 0x5a},
		Range{0x61, 0x7a},
	}

	// Space is the set of ASCII space characters.
	Space = []Range{
		Range{0x09, 0x0d},
		Range{0x20, 0x20},
	}

	// Symbol is the set of ASCII symbols.
	Symbol = []Range{
		Range{0x21, 0x2F},
		Range{0x3A, 0x40},
		Range{0x5B, 0x60},
		Range{0x7B, 0x7E},
	}

	// Control is the set of ASCII control characters.
	Control = []Range{
		Range{0x01, 0x1f},
		Range{0x7f, 0x7f},
	}
)
