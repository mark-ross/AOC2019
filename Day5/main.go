package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFuncTwoInput(opCode int) func(int, int) int {
	if opCode == 1 {
		return func(a, b int) int {
			return a + b
		}
	} else if opCode == 2 {
		return func(a, b int) int {
			return a * b
		}
	}
	return nil
}

func getValueFromSlice(dat *[]int, ind int, isRef bool) int {
	arr := *dat

	if isRef {
		addr1 := arr[ind]
		validateAddress(ind, len(arr))
		return arr[addr1]
	}

	return arr[ind]
}

func interpretInstruction(instruction int) (int, bool, bool, bool) {
	opCode := instruction % 100
	instruction /= 100
	isRightIndirect := instruction%10 == 0
	instruction /= 10
	isLeftIndirect := instruction%10 == 0
	instruction /= 10
	isAddrIndirect := instruction%10 == 0
	return opCode, isAddrIndirect, isLeftIndirect, isRightIndirect
}

func performComputation(dataPtr *[]int) {
	// dereference our integer slice pointer
	// and make a quick length variable
	data := *dataPtr
	dataLen := len(data)

	// loop over the code
	for iter := 0; iter < dataLen; {
		// validate the opcode -- fatal error if bad
		instruction := data[iter]
		opCode, _, isLeftIndirect, isRightIndirect := interpretInstruction(instruction)

		// if we've reached this code, terminate execution
		if opCode == 99 {
			return
		}

		// math operations
		if opCode < 3 && opCode > 0 {
			val1 := getValueFromSlice(&data, iter+1, isLeftIndirect)
			val2 := getValueFromSlice(&data, iter+2, isRightIndirect)
			outAddr := getValueFromSlice(&data, iter+3, true)

			// get the operating function based on the opCode
			operFunc := getFuncTwoInput(opCode)

			// reassign the value in the data slice based on the contents
			// of the slice at the given addresses
			data[outAddr] = operFunc(val1, val2)

			iter += 4

			// these are for read/write operations
		} else if opCode > 2 && opCode < 5 {
			if opCode == 3 {
				addr := data[iter+1]

				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal("Problem getting input")
				}

				val, err := strconv.Atoi(input)
				if err != nil {
					log.Fatal("Incorrect format for integer given!")
				}

				data[addr] = val
			} else if opCode == 4 {
				addr := data[iter+1]
				val := data[addr]
				fmt.Println(val)
			}

			iter += 2
		}

	}
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to open the file...", err)
	}
	defer func() {
		if err = inputFile.Close(); err != nil {
			log.Fatal("Failed to close file...", err)
		}
	}()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		txt := scanner.Text()
		codes := getIntSlice(txt)
		performComputation(codes)
	}
}
