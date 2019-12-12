package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetPlanetsFromStringAsMap(t *testing.T) {
	var tests = []struct {
		in  string
		out map[string][]string
	}{
		{`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`, map[string][]string{
			"COM": {"B"},
			"B":   {"C", "G"},
			"C":   {"D"},
			"D":   {"E", "I"},
			"E":   {"F", "J"},
			"G":   {"H"},
			"J":   {"K"},
			"K":   {"L"},
		}},
	}

	i := -1
	for _, tt := range tests {
		i++
		t.Run(fmt.Sprintln(i), func(t *testing.T){
			res := getPlanetsFromStringAsMap(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("Got: %+v\nExpected: %+v", res, tt.out)
			}
		})
	}
}

func TestGetPlanetsFromString(t *testing.T) {
	var tests = []struct {
		in  string
		out map[string][]string
	}{
		{`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`, map[string][]string{
			"COM": {"B"},
			"B":   {"C", "G"},
			"C":   {"D"},
			"D":   {"E", "I"},
			"E":   {"F", "J"},
			"G":   {"H"},
			"J":   {"K"},
			"K":   {"L"},
		}},
	}

	i := -1
	for _, tt := range tests {
		i++
		t.Run(fmt.Sprintln(i), func(t *testing.T){
			getPlanetsFromString(tt.in)
			//if !reflect.DeepEqual(res, tt.out) {
			//	t.Errorf("Got: %+v\nExpected: %+v", res, tt.out)
			//}
		})
	}
}

func TestGetOrbitCount(t *testing.T) {
	inStr := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	_, mp := getPlanetsFromString(inStr)
	cnt := getOrbitCount(mp)
	if cnt != 42 {
		log.Fatal("Incorrect number of connections: ", cnt, " vs. Expected: 42")
	}
}