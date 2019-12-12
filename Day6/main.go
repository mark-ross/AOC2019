package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

	allText, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal("Unable to read file")
	}
	doWork(string(allText))
}

func doWork(txt string) {
	_, res := getPlanetsFromString(txt)
	fmt.Println(getOrbitCount(res))
}

func getPlanetsFromStringAsMap(in string) map[string][]string {
	result := make(map[string][]string)
	allEntries := strings.Split(in, "\n")
	for _, entry := range allEntries {
		split := strings.Split(entry, ")")
		parent := split[0]
		child := split[1]

		val, ok := result[parent]
		if !ok {
			lis :=make([]string, 0)
			result[parent] = lis
			val = result[parent]
		}

		val = append(val, child)
		result[parent] = val
	}
	return result
}

type Planet struct {
	name string
	orbitedBy []*Planet
	orbits *Planet
}

func getPlanetsFromString(in string) (*Planet, map[string]*Planet) {
	results := make(map[string]*Planet)

	isFirst := true
	root := &Planet{
		name: "Root",
		orbitedBy: []*Planet{},
		orbits: nil,
	}

	allEntries := strings.Split(in, "\n")
	for _, entry := range allEntries {
		split := strings.Split(entry, ")")
		base := split[0]
		newOrbit := split[1]

		// update the name if the first one...
		if isFirst {
			root.name = base
			results[base] = root
			isFirst = false
		}

		nPtr, ok := results[newOrbit]
		if !ok {
			nPtr = &Planet{
				name:      newOrbit,
				orbitedBy: []*Planet{},
				orbits:    nil,
			}
			results[newOrbit] = nPtr
		}

		ptr, ok := results[base]
		if !ok {
			ptr = &Planet{
				name:      base,
				orbitedBy: []*Planet{nPtr},
				orbits:    nil,
			}
			results[base] = ptr
		} else {
			ptr.orbitedBy = append(ptr.orbitedBy, nPtr)
		}
		nPtr.orbits = ptr
	}

	return root, results
}

func getOrbitCount(res map[string]*Planet) int {

	directOrbitCount := len(res) -1

	// find all the leaf nodes (last nodes)
	leaves := make([]*Planet, 0, len(res)/2)
	for _, v := range res {
		if len(v.orbitedBy) == 0 {
			leaves = append(leaves, v)
		}
	}


	tracked := map[string]bool{}

	// loop over each leaf's linked list
	// and increment the steps to the source
	indirectOrbitCount := 0
	for _, leaf := range leaves {
		curr := leaf
		// loop until we hit the source
		for i:=0;; {
			curr = curr.orbits
			if curr == nil {
				break
			}

			if _, ok := tracked[curr.name]; !ok {
				i++
				tracked[curr.name] = true
			}
			indirectOrbitCount += i
		}
	}
	return directOrbitCount + indirectOrbitCount
}
