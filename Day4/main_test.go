package main

import (
	"fmt"
	"testing"
)

func TestGetPassFromInt(t *testing.T) {
	var tests = []struct {
		in  int
		out []uint8
	}{
		{123456, []uint8{1, 2, 3, 4, 5, 6}},
		{456789, []uint8{4, 5, 6, 7, 8, 9}},
		{147258, []uint8{1, 4, 7, 2, 5, 8}},
		{258369, []uint8{2, 5, 8, 3, 6, 9}},
		{1123456, []uint8{11, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			v := getPassFromInt(tt.in)
			if len(v) != len(tt.out) {
				t.Error("Arrays not of the same length")
			}

			for i := 0; i < len(v); i++ {
				if v[i] != tt.out[i] {
					t.Errorf("Incorrect. Got digit %d as %d, expected: %d", i, v[i], tt.out[i])
				}
			}
		})
	}
}

func TestMonotonicallyIncreasing(t *testing.T) {
	var tests = []struct {
		in  int
		out bool
	}{
		// problem examples
		{111111, true},
		{223450, false},
		{123789, true},
		// additional examples
		{223456, true},
		{123799, true},
		{000000, true},
		{177777, true},
		{176777, false},
		{987654, false},
		// From the interweb
		{123456, true},
		{111111, true},
		{112233, true},
		{111999, true},
		{123454, false},
		{212345, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			pass := getPassFromInt(tt.in)
			v := monotonicallyIncreasing(pass)
			if v != tt.out {
				t.Error("Incorrect. Expected: ", tt.out, " Got: ", v)
			}
		})
	}
}

func TestIsPaired(t *testing.T) {
	var tests = []struct {
		in  int
		out bool
	}{
		// problem examples
		{111111, true},
		{223450, true},
		{123789, false},
		// additional examples
		{223456, true},
		{123799, true},
		{000000, true},
		{177777, true},
		{176777, true},
		{987654, false},
		// From the interweb
		{873224, true},
		{857211, true},
		{883263, true},
		{837243, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			pass := getPassFromInt(tt.in)
			hasPaired := false
			for i := range pass {
				if isPaired(i, pass) {
					hasPaired = true
				}
			}
			if hasPaired != tt.out {
				t.Error("Incorrect. Expected: ", tt.out, " Got: ", hasPaired)
			}
		})
	}
}

func TestAtLeastOneExclusivePair(t *testing.T) {
	var tests = []struct {
		in  int
		out bool
	}{
		// problem examples
		{111111, false},
		{223450, true},
		{123789, false},
		// additional examples
		{223456, true},
		{123799, true},
		{000000, false},
		{177777, false},
		{176777, false},
		{987654, false},
		// From the interweb
		{873224, true},
		{857211, true},
		{888263, false},
		{837243, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			pass := getPassFromInt(tt.in)
			v := atLeastOneExclusivePair(pass)
			if v != tt.out {
				t.Error("Incorrect. Expected: ", tt.out, " Got: ", v)
			}
		})
	}
}

func TestValidation(t *testing.T) {
	var tests = []struct {
		in  int
		out bool
	}{
		// problem examples
		{111111, true},
		{223450, false},
		{123789, false},
		// additional examples
		{223456, true},
		{123799, true},
		{000000, true},
		{177777, true},
		{176777, false},
		{987654, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			pass := getPassFromInt(tt.in)
			v := isValidPass(pass)
			if v != tt.out {
				t.Error("Incorrect. Expected: ", tt.out, " Got: ", v)
			}
		})
	}
}
