package main

import (
	"bufio"
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

func readFile(scanner *bufio.Scanner) GamePool {
	pool := GamePool{}
	for scanner.Scan() {
		line := scanner.Text()
		newGame := Game{}
		fmt.Sscanf(line, "Game %v:", &newGame.number)
		newGame.content = line
		newGame.possible = true

		pool = append(pool, newGame)

	}
	return pool
}

func day2part1() int {
	scanner, file := getScanner(D2FP)
	defer file.Close()
	pool := readFile(scanner)
	// for _, game := range pool {
	// 	colorCheck("red", &game, 12)
	// 	if game.possible {
	// 		colorCheck("green", &game, 13)
	// 	}
	// 	if game.possible {
	// 		colorCheck("blue", &game, 14)
	// 	}
	// }
	for i := range pool {
		colorCheck("red", &pool[i], 12)
		if pool[i].possible {
			colorCheck("green", &pool[i], 13)
		}
		if pool[i].possible {
			colorCheck("blue", &pool[i], 14)
		}
	}

	score := 0

	for _, game := range pool {
		if game.possible {
			score += game.number
		}
	}

	return score

}

func day2part2() int {
	scanner, file := getScanner(D2FP)
	defer file.Close()
	pool := readFile(scanner)
	red, green, blue := 0, 0, 0
	score := 0
	for i := range pool {
		red = minimumCheck("red", &pool[i])
		green = minimumCheck("green", &pool[i])
		blue = minimumCheck("blue", &pool[i])
		score += red * green * blue
	}
	return score
}
func patternFind(color string, game *Game) [][]string {
	pattern := fmt.Sprintf(`\d+ %s`, color)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(game.content, -1)

	return matches
}

func minimumCheck(color string, game *Game) int {
	matches := patternFind(color, game)
	atleast := 0

	for _, match := range matches {
		parts := strings.Split(match[0], " ")
		number, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		if number > atleast {
			atleast = number
		}
	}

	return atleast
}

func colorCheck(color string, game *Game, limit int) {
	matches := patternFind(color, game)
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
