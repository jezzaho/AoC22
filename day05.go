package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type movement struct {
	amount, from, to int
}

func day5() {
	f, err := os.Open("./data/05.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()

	// Rather ugly, but i am stupid
	container, movements := loadData(f)
	f.Seek(0, io.SeekStart)
	containerDup, movementsDup := loadData(f)

	container = manipulatePart1(movements, container)
	containerDup = manipulatePart2(movementsDup, containerDup)

	resultPartOne := showTop(container)
	fmt.Println("The result for Part 1 is: " + resultPartOne)
	resultPartTwo := showTop(containerDup)
	fmt.Println("The result for Part 2 is: " + resultPartTwo)
}
func loadData(f *os.File) (map[int][]string, []movement) {
	stack := make(map[int][]string)
	var movements []movement
	scanner := bufio.NewScanner(f)
	lineCounter := 1
	for scanner.Scan() {
		line := scanner.Text()
		if lineCounter <= 8 {
			for i := 0; i < len(line); i++ {
				if line[i] >= 'A' && line[i] <= 'Z' {
					stack[i/4] = append(stack[i/4], string(line[i]))
				}
			}
		}
		if lineCounter >= 11 {
			var a, f, t int
			fmt.Sscanf(line, "move %v from %v to %v", &a, &f, &t)
			movements = append(movements, movement{a, f, t})

		}
		lineCounter++
	}

	return stack, movements
}
func manipulatePart1(mList []movement, c map[int][]string) map[int][]string {
	for i := 0; i < len(mList); i++ {
		m := mList[i]
		for i := 0; i < m.amount; i++ {
			c[m.to-1] = append([]string{c[m.from-1][0]}, c[m.to-1]...)
			c[m.from-1] = c[m.from-1][1:]
		}
	}
	return c

}
func manipulatePart2(mList []movement, c map[int][]string) map[int][]string {
	for i := 0; i < len(mList); i++ {
		m := mList[i]
		c[m.to-1] = append(c[m.from-1][:m.amount:m.amount], c[m.to-1]...)
		c[m.from-1] = c[m.from-1][m.amount:]
	}

	return c
}
func showTop(c map[int][]string) string {
	s := ""
	keys := make([]int, 0)
	for k := range c {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if len(c[k]) > 0 {
			s += c[k][0]
		}
	}
	return s
}
