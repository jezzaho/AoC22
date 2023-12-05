package main

import (
	"bufio"
	"log"
	"os"
)

func getScanner(filename string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	return scanner, f
}
