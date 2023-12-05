package main

import (
	"strings"
	"unicode"
)

const FILEPATH = "./files/day1.txt"

func day1part1() int {

	scanner, file := getScanner(FILEPATH)

	defer file.Close()

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, second := 0, 0
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if first == 0 {
					first = int(line[i]) - 48
				}
				second = int(line[i]) - 48
			}
		}
		score += 10*first + second
	}

	return score

}

func day1part2() int {
	digits := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	scanner, file := getScanner(FILEPATH)
	defer file.Close()
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		for k, v := range digits {
			replacement := []rune(k)
			replacement[len(k)/2] = []rune(v)[0]
			line = strings.Replace(line, k, string(replacement), -1)
		}
		first, second := 0, 0
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if first == 0 {
					first = int(line[i]) - 48
				}
				second = int(line[i]) - 48
			}
		}
		score += 10*first + second
	}

	return score
}
