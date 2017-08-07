package main

import (
	"fmt"
	"math/rand"
	"time"
)

const height = 30
const width = 100

var cells [height][width]int

func randSeed() {
	Loop2D(height, width, func(row, col int) {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		switch {
		case !inBounds(row, col):
			cells[row][col] = 2
		case n%2 == 0:
			cells[row][col] = 1
		case n%2 > 0:
			cells[row][col] = 0
		}
	})
}

func draw() {
	Loop2D(height, width, func(row, col int) {
		switch cells[row][col] {
		case 0:
			fmt.Print(" ")
		case 1:
			fmt.Print("•")
		default:
			fmt.Print(" ")
		}
		if col == width-1 {
			fmt.Println("")
		}
	})
}

func inBounds(row, col int) bool {
	if row == height-1 || row == 0 {
		return false
	}
	if col == width-1 || col == 0 {
		return false
	}
	return true
}

func n(row, col int) bool {
	// if row
	return cells[row-1][col] > 0
}

func applyRules() {
	// buf := cells[:][:]
	Loop2D(height, width, func(row, col int) {
		switch {
		case n(row, col):
			// fmt.Println("n alive")
		}
	})

}

// GameOfLife : conways game of life
func GameOfLife() {
	println("Conways Game of Life")
	randSeed()
	draw()
	// applyRules()

}
