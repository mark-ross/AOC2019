package main

import "testing"

var tests = []struct {
	in  string
	out string
}{
	{"1,0,0,0,99", "2,0,0,0,99"},
	{"2,3,0,3,99", "2,3,0,6,99"},
	{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
	{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
}

func TestDoWork(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := doWork(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
