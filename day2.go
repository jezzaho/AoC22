package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
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
	for _, game := range pool {
		colorCheck("red", &game, 12)
		if &game.possible {
			colorCheck("green", &game, 13)
		}
		if &game.possible {
			colorCheck("blue", &game, 14)
		}
	}
	
	count := 100

	for _, game := range pool {
		if !game.possible { 
			count--
		}
	}

	return count

}

func colorCheck(color string, game *Game, limit int) {
	pattern := fmt.Sprintf(`\d+ %s`, color)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(game.content, -1)
	biggest := 0
	// var biggest int
	for _, match := range matches {
		parts := strings.Split(match[0], " ")
		number, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatal(err)
		}
		if number > biggest {
			biggest = number

		}
	}
	if biggest > limit {
		game.possible = false
	}
}
