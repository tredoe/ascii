// Copyright 2014 Jonas mg
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ascii

import (
	"testing"
)

func TestSpace(t *testing.T) {
	sp := []byte{'\t', '\n', '\v', '\f', '\r', ' '}

	for _, b := range sp {
		if !IsSpace(b) {
			t.Errorf("character should be of kind space: %#x", b)
		}
	}
}

func TestDigit(t *testing.T) {
	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for _, b := range digits {
		if !IsDigit(b) {
			t.Errorf("character '%#x' should be a digit", b)
		}
	}
}

func TestSymbol(t *testing.T) {
	chars := []byte{0x21, 0x2F, 0x3A, 0x40, 0x5B, 0x60, 0x7B, 0x7E, '.'}

	for _, b := range chars {
		if !IsSymbol(b) {
			t.Errorf("character '%c' should be a symbol", b)
		}
	}
}
