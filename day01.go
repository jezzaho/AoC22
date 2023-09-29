package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
)

func day1p1() []int {
	// Open a file
	f, err := os.Open("./data/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()

	// Reading one by one using scanner
	scanner := bufio.NewScanner(f)
	count := 0
	// Declare a storage for elves calories count
	elfCalories := make([]int, 0)
	for scanner.Scan() {
		// fmt.Printf("Line: %v\n", scanner.Text())
		if scanner.Text() == "" {
			elfCalories = append(elfCalories, count)
			count = 0
		} else {
			val, err := strconv.ParseInt(scanner.Text(), 0, 0)
			if err != nil {
				fmt.Println("Error during parsing: String -> Integer")
			}
			count += int(val)
		}
	}
	return elfCalories
}

func day1p2() int {
	elfCalories := day1p1()
	// Finding top three elves by amount of calories
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	return (elfCalories[0] + elfCalories[1] + elfCalories[2])

}

func day1Results() {
	// Day 1 Results Part 1
	fmt.Printf("The maximum calories carried by elf is: %v\n", slices.Max(day1p1()))

	// Day 1 Results Part 2
	fmt.Printf("The sum of top three elves is: %v\n", day1p2())

}
