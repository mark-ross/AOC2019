package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p *point) manhattanDistance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (p *point) distance(q *point) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}

func (p *point) up(v int) *point {
	return &point{p.x, p.y + v}
}

func (p *point) down(v int) *point {
	return &point{p.x, p.y - v}
}

func (p *point) left(v int) *point {
	return &point{p.x - v, p.y}
}

func (p *point) right(v int) *point {
	return &point{p.x + v, p.y}
}

// custom sorting method
type byManhattan []*point

func (s byManhattan) Len() int {
	return len(s)
}
func (s byManhattan) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byManhattan) Less(i, j int) bool {
	return s[i].manhattanDistance() < s[j].manhattanDistance()
}

type line struct {
	start point
	end   point
	min   point
	max   point
}

func newLine(start, end point) *line {
	minX := start.x
	maxX := end.x
	if maxX < minX {
		minX, maxX = maxX, minX
	}

	minY := start.y
	maxY := end.y
	if maxY < minY {
		minY, maxY = maxY, minY
	}

	return &line{
		start: start,
		end:   end,
		min:   point{minX, minY},
		max:   point{maxX, maxY},
	}
}

func (l *line) isHorizontal() bool {
	if l.start.y == l.end.y {
		return true
	}
	return false
}

func (l *line) isVertical() bool {
	if l.start.x == l.end.x {
		return true
	}
	return false
}

func (l *line) length() int {
	// hacky solution since we know that it's gridded distances
	return l.max.x - l.min.x + l.max.y - l.min.y
}

func (l *line) intersects(m *line) (bool, *point) {

	var horizLine *line
	if l.isHorizontal() {
		horizLine = l
	} else if m.isHorizontal() {
		horizLine = m
	}
	// if they're parallel, they don't intersect
	if l.isHorizontal() && m.isHorizontal() {
		return false, nil
	}

	var vertLine *line
	if l.isVertical() {
		vertLine = l
	} else if m.isVertical() {
		vertLine = m
	}
	// if they're parallel, they don't intersect
	if l.isVertical() && m.isVertical() {
		return false, nil
	}

	if horizLine.min.x <= vertLine.start.x && horizLine.max.x >= vertLine.start.x {
		if vertLine.min.y <= horizLine.start.y && vertLine.max.y >= horizLine.start.y {

			// disallow the origin from being included
			if vertLine.start.x == 0 && horizLine.start.y == 0 {
				return false, nil
			}

			return true, &point{vertLine.start.x, horizLine.start.y}
		}
	}

	return false, nil
}

type cable struct {
	lines []*line
}

func newCable(input string) *cable {
	csv := strings.Split(input, ",")
	if len(csv) < 2 {
		log.Fatal("Should have more than 2!")
	}

	c := cable{}
	c.lines = make([]*line, len(csv))
	prevPnt := point{0, 0} // always start at origin
	for i, ele := range csv {
		// grab the direction piece
		dir := string(ele[0])

		// pull the intensity and turn it into an int
		valStr := ele[1:]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			log.Fatal("Incorrect int format!: ", valStr)
		}

		// switch on the direction variable
		var nextPnt point
		if dir == "U" {
			nextPnt = *prevPnt.up(val)
		} else if dir == "D" {
			nextPnt = *prevPnt.down(val)
		} else if dir == "L" {
			nextPnt = *prevPnt.left(val)
		} else if dir == "R" {
			nextPnt = *prevPnt.right(val)
		} else {
			log.Fatal("Unrecognized direction variable!: ", dir)
		}

		// generate a line from the two points, and update the
		// references for the next pass
		c.lines[i] = newLine(prevPnt, nextPnt)
		prevPnt = nextPnt
	}

	// return the generated cable
	return &c
}

func (c *cable) getIntersections(d *cable) []int {
	results := make([]int, 0)
	for i, s1 := range c.lines {
		for j, s2 := range d.lines {
			if ok, pnt := s1.intersects(s2); ok {
				// Part 1 solver
				// results = append(results, pnt.manhattanDistance())

				steps := c.getStepsAtLineIndex(i) + d.getStepsAtLineIndex(j)
				steps += s1.start.distance(pnt) + s2.start.distance(pnt)
				results = append(results, steps)
			}
		}
	}
	return results
}

func (c *cable) getStepsAtLineIndex(ind int) int {
	if ind < 0 {
		log.Fatal("Index value must be >0")
	}
	if ind >= len(c.lines) {
		log.Fatal("Index request is > length of slice")
	}

	sum := 0
	for i := 0; i < ind; i++ {
		sum += c.lines[i].length()
	}

	return sum
}

func getBestDistance(input []string) int {
	cable1 := newCable(input[0])
	cable2 := newCable(input[1])

	intersectionPoints := cable1.getIntersections(cable2)
	if len(intersectionPoints) < 1 {
		log.Fatal("No intersection points.")
	}
	var sb strings.Builder
	for _, pt := range intersectionPoints {
		sb.WriteString(fmt.Sprintf("%v, ", pt))
	}
	log.Printf("Result points: %+v", sb.String())
	sort.Ints(intersectionPoints)
	dist := intersectionPoints[0]
	log.Println("Best Score: ", dist)
	return dist
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

	ctr := 0
	inputHolder := make([]string, 2)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lineText := strings.Trim(scanner.Text(), "\r\n")

		inputHolder[ctr%2] = lineText
		if ctr%2 == 1 {
			getBestDistance(inputHolder)
		}
		ctr++
	}
}
