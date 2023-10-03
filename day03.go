package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day3p1() {
	// Open a file
	f, err := os.Open("./data/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()
	// Reading one by one using scanner
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		comp1 := scanner.Text()[:(len(scanner.Text()) / 2)]
		comp2 := scanner.Text()[(len(scanner.Text()) / 2):]
		r := findRune(comp1, comp2)
		counter += calcRune(r)
		// fmt.Println(comp1)
		// fmt.Println(comp2)
	}
	fmt.Printf(`The total value of item is: %v`, counter)

}
func findRune(str1, str2 string) rune {
	for _, char := range str1 {
		if strings.ContainsRune(str2, char) {
			return char
		}
	}
	return 0
}
func findThreeRune(str1, str2, str3 string) rune {
	for _, char := range str1 {
		if strings.ContainsRune(str2, char) && strings.ContainsRune(str3, char) {
			return char
		}
	}
	return 0

}
func calcRune(r rune) int {
	if r > 96 {
		return (int(r) - 96)
	}
	return (int(r) - 38)
}

func day3p2() {
	// Open a file
	f, err := os.Open("./data/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()
	// Reading one by one using scanner
	scanner := bufio.NewScanner(f)
	counter := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if len(lines) == 3 {
			counter += processGroup(lines)
			lines = nil
		}
	}
	fmt.Printf(`The value is: %v`, counter)
}

func processGroup(lines []string) int {
	r := findThreeRune(lines[0], lines[1], lines[2])
	return calcRune(r)
}
