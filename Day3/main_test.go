package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	in  []string
	out int
}{
	{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 6},
	{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
	{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135},
}

func TestDoWork(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			s := getBestDistance(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}

var fullTests = []struct {
	in  []string
	out int
}{
	{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 40},
	{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
	{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410},
}

func TestDoWorkPart2(t *testing.T) {
	for _, tt := range fullTests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			s := getBestDistance(tt.in)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
