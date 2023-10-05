package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type section struct {
	left, right int
}

func day4p1() {
	// Open a file
	f, err := os.Open("./data/04.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()
	// Init a counter
	counterF := 0
	counterO := 0
	// Reading one by one using scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		left, right := parseLine(line)
		if doesCover(left, right) {
			counterF++
		}
		if doesOverlap(left, right) {
			counterO++
		}

	}

	fmt.Printf("Number of assigment pairs when one fully contains other is: %v\n", counterF)
	fmt.Printf("Number of assigment pairs when they is: %v\n", counterO)
}
func parseLine(s string) (section, section) {
	var first, second section
	fmt.Sscanf(s, "%d-%d,%d-%d", &first.left, &first.right, &second.left, &second.right)
	return first, second
}
func doesCover(a, b section) bool {
	if (a.left <= b.left && a.right >= b.right) || (b.left <= a.left && b.right >= a.right) {
		return true
	}
	return false
}
func doesOverlap(a, b section) bool {
	if (a.left <= b.left && a.right >= b.right) || (a.left <= b.right && a.right >= b.left) {
		return true
	}
	return false
}
