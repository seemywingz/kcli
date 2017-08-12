package main

import (
	"fmt"
	"math/rand"
	"time"

	gt "github.com/seemywingz/gtils"
)

const height = 40
const width = 100

var cells = make([][]int, height)
var iterations int

func randSeed() {
	seed := make([][]int, height)
	for i := range seed {
		seed[i] = make([]int, width)
	}

	for i := 0; i < width*2; i++ {
		rand.Seed(time.Now().UnixNano())
		seed[rand.Intn(height)][rand.Intn(width)] = 1
	}

	cells = seed
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
	checkNeighbors := make([]func(int, int) bool, 8)
	checkNeighbors[0] = getNeighbor(-1, -1)
	checkNeighbors[1] = getNeighbor(-1, 0)
	checkNeighbors[2] = getNeighbor(-1, +1)
	checkNeighbors[3] = getNeighbor(0, +1)
	checkNeighbors[4] = getNeighbor(0, -1)
	checkNeighbors[5] = getNeighbor(+1, +1)
	checkNeighbors[6] = getNeighbor(+1, 0)
	checkNeighbors[7] = getNeighbor(+1, -1)

	gt.Loop2D(height, width, func(row, col int) {

		neighbors := 0
		if !inBounds(row, col) {
			buf[row][col] = 2
		} else {
			for i := range checkNeighbors {
				if checkNeighbors[i](row, col) {
					neighbors++
				}
			}
			cell := cells[row][col]
			if cell == 1 { // alive
				if neighbors < 2 || neighbors > 3 { // dies
					buf[row][col] = 0
				}
			}
			if cell == 0 { // dead
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
	gt.Loop2D(height, width, func(row, col int) {
		switch cells[row][col] {
		case 0:
			fmt.Print(" ")
		case 1:
			// fmt.Print("◉")
			// fmt.Print("▒")
			fmt.Print("•")
		default:
			fmt.Print("x")
		}
		if col == width-1 {
			fmt.Println("")
		}
	})
	fmt.Println("Iterations:", iterations)
}

// GameOfLife : conways game of life
func GameOfLife() {
	println("Conways Game of Life")
	randSeed()

	// for i := 0; i < 1; i++ {
	for {
		draw()
		applyRules()
		time.Sleep(40 * time.Millisecond)
		iterations++
	}

}
