package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day2p1() {
	// Open a file
	f, err := os.Open("./data/02.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()

	// Reading one by one using scanner
	scanner := bufio.NewScanner(f)
	counterPartOne := 0
	counterPartTwo := 0
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Fields(line)
		counterPartOne += calculatePointsPartOne(moves[0], moves[1])
		counterPartTwo += calculatePointsPartTwo(moves[0], moves[1])

	}
	fmt.Printf("Total points score for part One is: %v\n", counterPartOne)
	fmt.Printf("Total points score for part Two is: %v\n", counterPartTwo)
}

func calculatePointsPartOne(opo, you string) int {
	counter := 0
	if you == "X" {
		counter += 1
		if opo == "A" {
			counter += 3
		}
		if opo == "C" {
			counter += 6
		}

	}
	if you == "Y" {
		counter += 2
		if opo == "A" {
			counter += 6
		}
		if opo == "B" {
			counter += 3
		}
	}
	if you == "Z" {
		counter += 3
		if opo == "B" {
			counter += 6
		}
		if opo == "C" {
			counter += 3
		}
	}

	return counter
}
func calculatePointsPartTwo(opo, you string) int {
	m := make(map[string]int)
	m["A"], m["B"], m["C"] = 1, 2, 3
	// Case lose
	if you == "X" {
		if opo == "A" {
			return 3
		}
		if opo == "B" {
			return 1
		}
		if opo == "C" {
			return 2
		}

	}
	// case draw
	if you == "Y" {
		return m[opo] + 3
	}
	// case win
	if you == "Z" {
		if opo == "A" {
			return 2 + 6 // or just 2+6= 8 because P+6
		}
		if opo == "B" {
			return 3 + 6 // or just 3+6 = 9 because S+6
		}
		if opo == "C" {
			return 1 + 6 // or just 1+6 = 7 because R+6
		}
	}
	return 0
}

// A - X, A - Y, A - Z
// B - X, B - Y, B - Z
// C - X, C - Y, C - Z
// A,X - 1, B,Y - 2, C,Z - 3
// LOSE 0 DRAW 3 WIN 6
//
//
