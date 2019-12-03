package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	in  int
	out int
}{
	{5, 0},
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestDoWork(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			s := calculateFuelRequirement(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

var fullTests = []struct {
	in  int
	out int
}{
	{1969, 966},
	{100756, 50346},
}

func TestFullCalculate(t *testing.T) {
	for _, tt := range fullTests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			s := fullyCalculateFuel(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
