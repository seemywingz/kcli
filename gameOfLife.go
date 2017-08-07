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

	for i := 0; i < 500; i++ {
		rand.Seed(time.Now().UnixNano())
		seed[rand.Intn(height)][rand.Intn(width)] = 1
	}

	cells = seed[:][:]
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
	f := []func(int, int) bool{}
	// f[0] = getNeighbor(-1, -1)

	for row := -1; row < 2; row++ {
		for col := -1; col < 2; col++ {
			f[0] = getNeighbor(-1, -1)
		}
	}

	// f[0].(getNeighbor(-1, -1)
	// f[0] = getNeighbor(-1, 0)
	// f[0] = getNeighbor(-1, +1)
	// f[0] = getNeighbor(0, +1)
	// f[0] = getNeighbor(0, -1)
	// f[0] = getNeighbor(+1, +1)
	// f[0] = getNeighbor(+1, 0)
	// f[0] = getNeighbor(+1, -1)

	// nw := getNeighbor(-1, -1)
	// n := getNeighbor(-1, 0)
	// ne := getNeighbor(-1, +1)
	// e := getNeighbor(0, +1)
	// w := getNeighbor(0, -1)
	// se := getNeighbor(+1, +1)
	// s := getNeighbor(+1, 0)
	// sw := getNeighbor(+1, -1)

	Loop2D(height, width, func(row, col int) {

		neighbors := 0
		if !inBounds(row, col) {
			buf[row][col] = 2
		} else {
			for i := range f {
				if f[i](row, col) {
					neighbors++
				}
			}
			// if n(row, col) {
			// 	neighbors++
			// }
			// if ne(row, col) {
			// 	neighbors++
			// }
			// if e(row, col) {
			// 	neighbors++
			// }
			// if se(row, col) {
			// 	neighbors++
			// }
			// if s(row, col) {
			// 	neighbors++
			// }
			// if sw(row, col) {
			// 	neighbors++
			// }
			// if w(row, col) {
			// 	neighbors++
			// }
			// if nw(row, col) {
			// 	neighbors++
			// }
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
}

func draw() {
	print("\033[H\033[2J") // clear screen
	Loop2D(height, width, func(row, col int) {
		switch cells[row][col] {
		case 0:
			fmt.Print(" ")
		case 1:
			// fmt.Print("▒")
			fmt.Print("•")
		default:
			fmt.Print("x")
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
		time.Sleep(300 * time.Millisecond)
	}

}
