package main

import (
	"fmt"
	"math/rand"
	"time"
)

const height = 50
const width = 100

var cells = make([][]int, height)

func randSeed() {
	seed := make([][]int, height)
	for i := range seed {
		seed[i] = make([]int, width)
	}
	Loop2D(height, width, func(row, col int) {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100)
		n := 0
		switch {
		case !inBounds(row, col):
			n = 2
		case randNum%2 == 0:
			n = 1
		case randNum%2 > 0:
			n = 0
		}
		seed[row][col] = n
	})
	for i := range seed {
		cells[i] = seed[i]
	}
	// fmt.Printf("LENGTH: %v\nCAPACITY: %v\nSEED: %v\n", len(seed), cap(seed), seed)
	// fmt.Printf("LENGTH: %v\nCAPACITY: %v\nCELLS: %v\n", len(cells), cap(cells), cells)
}

func inBounds(row, col int) bool {
	i := true
	if row == 0 || row == height-1 {
		i = false
	}
	if col == 0 || col == width-1 {
		i = false
	}
	return i
}

func getNeighbor(deltaRow, deltaCol int) func(row, col int) bool {
	return func(row, col int) bool {
		if !inBounds(row, col) {
			return false
		}
		return cells[row+deltaRow][col+deltaCol] == 1
	}
}

func applyRules() {

	buf := cells[:][:]

	n := getNeighbor(-1, 0)
	ne := getNeighbor(-1, +1)

	e := getNeighbor(0, +1)
	se := getNeighbor(+1, +1)

	s := getNeighbor(+1, 0)
	sw := getNeighbor(+1, -1)

	w := getNeighbor(0, -1)
	nw := getNeighbor(-1, -1)

	Loop2D(height, width, func(row, col int) {

		neighbors := 0
		if !inBounds(row, col) {
			buf[row][col] = 2
		} else {
			if n(row, col) {
				neighbors++
			}
			if ne(row, col) {
				neighbors++
			}
			if e(row, col) {
				neighbors++
			}
			if se(row, col) {
				neighbors++
			}
			if s(row, col) {
				neighbors++
			}
			if sw(row, col) {
				neighbors++
			}
			if w(row, col) {
				neighbors++
			}
			if nw(row, col) {
				neighbors++
			}
			// fmt.Printf("CEll (%v,%v): alive: %v neighbors: %v\n", row, col, cells[row][col] > 0, neighbors)
			cell := cells[row][col]
			if cell == 1 { // alive
				if neighbors < 2 || neighbors > 3 { // dies
					buf[row][col] = 0
				}
			}
			if cell == 0 { // alive
				if neighbors == 3 { // born
					buf[row][col] = 1
				}
			}
		}
	})
	cells = buf[:][:]
	// fmt.Printf("LENGTH: %v\nCAPACITY: %v\nSEED: %v\n", len(buf), cap(buf), buf)
	// fmt.Printf("LENGTH: %v\nCAPACITY: %v\nCELLS: %v\n", len(cells), cap(cells), cells)
}

func draw() {
	print("\033[H\033[2J") // clear screen
	Loop2D(height, width, func(row, col int) {
		switch cells[row][col] {
		case 0:
			fmt.Print(" ")
		case 1:
			fmt.Print("•")
		default:
			fmt.Print("X")
		}
		if col == width-1 {
			fmt.Println("")
		}
	})
}

// GameOfLife : conways game of life
func GameOfLife() {
	println("Conways Game of Life")
	randSeed()

	// for i := 0; i < 1; i++ {
	for {
		draw()
		applyRules()
		time.Sleep(100 * time.Millisecond)
	}

}
