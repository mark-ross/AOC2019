package main

import (
	"fmt"
	"testing"
)

func TestInterpretInstruction(t *testing.T) {
	var tests = []struct {
		in  int
		out string
	}{
		{1002, "2,true,false,true"},
		{11101, "1,false,false,false"},
		{99, "99,true,true,true"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintln(tt.in), func(t *testing.T) {
			opCode, ai, li, ri := interpretInstruction(tt.in)
			out := fmt.Sprint(opCode, ",", ai, ",", li, ",", ri)
			if out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}

func TestGetValueFromSlice(t *testing.T) {
	type input struct {
		slice    []int
		ind      int
		indirect bool
	}

	var tests = []struct {
		in  input
		out int
	}{
		{input{
			[]int{02,}
		}}
	}
}
