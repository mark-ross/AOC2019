package main

import (
	"log"
	"strconv"
	"strings"
)

func getIntSlice(inputString string) *[]int {
	csv := strings.Split(inputString, ",")
	if len(csv) < 2 {
		log.Fatal("Error parsing the string. Is this correct?: ", inputString)
		return nil
	}

	var err error
	results := make([]int, len(csv))
	for i, s := range csv {
		results[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatal("Got a non-integer...:", s)
		}
	}

	return &results
}

func validateAddress(addr, lengthOfSlice int) bool {
	if addr < 0 {
		log.Fatal("Memory address cannot be negative! :", addr)
		return false
	}

	if addr >= lengthOfSlice {
		log.Fatal("Memory address cannot exceed the length of the slice! :", addr)
		return false
	}

	return true
}
