package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func mainFindValue() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to open the file...", err)
	}
	defer func() {
		if err = inputFile.Close(); err != nil {
			log.Fatal("Failed to close file...", err)
		}
	}()

	ctr := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		result := doWork(scanner.Text())
		log.Printf("Results of line %d: %s", ctr, result)
		ctr++
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
		baseInput := scanner.Text()

		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				intSlice := getIntSlice(baseInput)
				(*intSlice)[1] = i
				(*intSlice)[2] = j

				performComputation(intSlice)
				if (*intSlice)[0] == 19690720 {
					log.Println("Solution [100 * noun + verb]: ", (100*i + j))
					return
				}
			}
		}
	}
}

func doWork(inputString string) string {
	intSlice := getIntSlice(inputString)
	performComputation(intSlice)

	var sb strings.Builder
	for i, v := range *intSlice {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

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

func performComputation(dataPtr *[]int) {
	// dereference our integer slice pointer
	// and make a quick length variable
	data := *dataPtr
	dataLen := len(data)

	// loop over the code
	for iter := 0; iter < dataLen; iter += 4 {
		// validate the opcode -- fatal error if bad
		opCode := data[iter]
		if opCode != 1 && opCode != 2 && opCode != 99 {
			log.Fatal("Illegal opCode given! ", opCode)
			return
		}

		// if we've reached this code, terminate execution
		if opCode == 99 {
			// log.Println("Stop code received... Ending program")
			return
		}

		// validate the addresses
		addr1 := data[iter+1]
		validateAddress(addr1, dataLen)

		addr2 := data[iter+2]
		validateAddress(addr2, dataLen)

		resultsAddr := data[iter+3]
		validateAddress(resultsAddr, dataLen)

		// get the operating function based on the opCode
		operFunc := getFunc(opCode)

		// reassign the value in the data slice based on the contents
		// of the slice at the given addresses
		data[resultsAddr] = operFunc(data[addr1], data[addr2])
	}
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

func getFunc(opCode int) func(int, int) int {
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
