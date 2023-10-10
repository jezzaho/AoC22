package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type movement struct {
	amount, from, to int
}

func day5p1() {
	f, err := os.Open("./data/05.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()

	container, movements := loadData(f)
	for i := 0; i < len(movements); i++ {
		manipulate(movements[i], container)
	}
	result := showTop(container)
	fmt.Println("The result is: " + result)
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
			fmt.Sscanf(line, "move %d from %d to %d", &a, &f, &t)
			movements = append(movements, movement{a, f, t})
			// fmt.Printf("%v\n", movements)
		}
		lineCounter++
	}

	return stack, movements
}
func manipulate(m movement, c map[int][]string) {
	moveFrom := c[m.from-1]
	moveTo := c[m.to-1]
	for i := 0; i < m.amount; i++ {
		el := moveFrom[0]
		moveTo = append([]string{el}, moveTo...)
		moveFrom = moveFrom[1:]
	}
	c[m.from-1] = moveFrom
	c[m.to-1] = moveTo
}
func showTop(c map[int][]string) string {
	s := ""
	keys := make([]int, 0)
	for k, _ := range c {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		s += c[k][0]
	}
	return s
}
