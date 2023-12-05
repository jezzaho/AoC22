package main

import (
	"fmt"
	"regexp"
)

const D2FP = "./files/day2.txt"

type Game struct {
	number   int
	content  string
	possible bool
}

type GamePool []Game

func day2part1() int {
	scanner, file := getScanner(D2FP)

	defer file.Close()
	pool := GamePool{}

	for scanner.Scan() {
		line := scanner.Text()
		newGame := Game{}
		fmt.Sscanf(line, "Game %v:", &newGame.number)
		newGame.content = line
		newGame.possible = true

		pool = append(pool, newGame)

	}

	return -1

}

func colorCheck(color string, game *Game) {
	pattern := fmt.Sprintf(`(d+) %s)`, color)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(game.content, -1)

	for _, match := range matches {
		for _, m := range match {
			print(m)
		}
	}

}
