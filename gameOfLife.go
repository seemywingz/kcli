package main

import (
	"fmt"
	"math/rand"
	"time"
)

const height = 20
const width = 50

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

func draw() {
	print("\033[H\033[2J") // clear screen
	Loop2D(height, width, func(row, col int) {
		switch cells[row][col] {
		case 0:
			fmt.Print(".")
		case 1:
			fmt.Print("•")
		default:
			fmt.Print("X")
		}
		// fmt.Print(cells[row][col])
		if col == width-1 {
			fmt.Println("")
		}
	})
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

	buf := make([][]int, height)
	for i := range buf {
		buf[i] = make([]int, width)
	}

	n := getNeighbor(-1, 0)
	ne := getNeighbor(-1, +1)

	e := getNeighbor(0, +1)
	se := getNeighbor(+1, +1)

	s := getNeighbor(+1, 0)
	sw := getNeighbor(+1, -1)

	w := getNeighbor(0, -1)
	nw := getNeighbor(-1, -1)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			neighbors := 0
			if !inBounds(row, col) {
				buf[row][col] = 3
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
				if cells[row][col] > 0 { // alive
					if neighbors >= 4 || neighbors < 2 { // dies
						buf[row][col] = 0
					}
				} else { // dead
					if neighbors == 2 { // born
						buf[row][col] = 1
					}
				}
			}
		}
	}

	// for i := range buf {
	// 	fmt.Println(buf[i])
	// }
	for i := range buf {
		cells[i] = buf[i]
	}
	// fmt.Printf("LENGTH: %v\nCAPACITY: %v\nSEED: %v\n", len(buf), cap(buf), buf)
	// fmt.Printf("LENGTH: %v\nCAPACITY: %v\nCELLS: %v\n", len(cells), cap(cells), cells)
}

// GameOfLife : conways game of life
func GameOfLife() {
	println("Conways Game of Life")
	randSeed()

	for i := 0; i < 1; i++ {
		draw()
		applyRules()
		time.Sleep(time.Second)
	}

}
