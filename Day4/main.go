package main

import (
	"fmt"
	"math"
)

func main() {
	n := getNumPossibleCombos(172851, 675869)
	fmt.Println("N: ", n)
}

func getNumPossibleCombos(min int, max int) int {
	// loop over the values (inclusively) to find
	// which of the integer values would work as
	// a password, given the rules.
	possible := 0
	for curr := min; curr <= max; curr++ {
		// get the []uint8 slice from the int
		passHolder := getPassFromInt(curr)
		// if the password is valid, increment
		if isValidPass(passHolder) {
			possible++
		}
	}

	// return back the possible number of passwords
	return possible
}

func getPassFromInt(val int) []uint8 {
	// prebake the allocation since we know the size
	passHolder := make([]uint8, 6)

	// loop over the values so that the largest
	// digit is at index 0, but means we have to
	// loop over m so 10^m will get us a single digit
	for i, m := 0, 5; i < 6; i, m = i+1, m-1 {
		// find the digit modifier we'll use
		digitMod := math.Pow10(m)
		// round down after getting some value >= 0 && < 10
		digitVal := math.Floor(float64(val) / digitMod)
		// cast it down and store in the array
		passHolder[i] = uint8(digitVal)
		// mutate the val we've been passing around
		val %= int(digitMod)
	}

	// finally return the results
	return passHolder
}

func isValidPass(pass []uint8) bool {
	// only allow increasing values across digits
	if !monotonicallyIncreasing(pass) {
		return false
	}

	// // for each digit in the pass
	// // check if it's part of a pair
	// for i := range pass {
	// 	if isPaired(i, pass) {
	// 		return true
	// 	}
	// }

	// part 2!
	if atLeastOneExclusivePair(pass) {
		return true
	}

	// after all the checks, return false if we get here...
	return false
}

func monotonicallyIncreasing(arr []uint8) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func isPaired(i int, arr []uint8) bool {
	// see if the digits next to the current
	// digit are pairs
	if i-1 > 0 {
		if arr[i-1] == arr[i] {
			return true
		}
	}
	if i+1 < len(arr)-1 {
		if arr[i+1] == arr[i] {
			return true
		}
	}
	return false
}

func atLeastOneExclusivePair(arr []uint8) bool {
	pairs := 0
	currentDigit := arr[0]
	count := 1
	for i := 0; i < len(arr); i++ {
		currentDigit = arr[i]
		// loop over all the digits after the value
		// to see if we have pairs (vs. larger sets)
		for j := i + 1; j < len(arr); j++ {
			// if the digit matches, incr counter
			if arr[j] == currentDigit {
				count++
				continue
			}

			// if we only had 2 next to each other
			// vs. > 2 together, increment the pairs
			// and reset the count
			if count == 2 {
				pairs++
				count = 1
			} else if count > 2 {
				// otherwise, just reset the count var
				count = 1
			}
			// if we get here, move to the next digit to check
			i = j - 1
			break
		}
	}

	if pairs > 0 || count == 2 {
		return true
	}
	return false
}
