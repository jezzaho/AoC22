package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func day6p1() {
	f, err := os.Open("./data/06.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		storage := []string{string(line[0]), string(line[1]), string(line[2]), string(line[3])}
		for i := 4; i < len(line); i++ {
			if !isNotPacketStarter(storage) {
				fmt.Printf("Value: %v is the packet marker.\nLocation: %v\n", string(line[i-1]), i)
				fmt.Printf("Packet: %v%v%v%v\n", string(line[i-4]), string(line[i-3]), string(line[i-2]), string(line[i-1]))
				break
			}
			storage = append(storage[1:len(storage):4], string(line[i]))
		}
	}
}
func day6p2() {
	// UGLY DRY
	f, err := os.Open("./data/06.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close a file
	defer f.Close()

	f.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		storage := make([]string, 14)
		for i := 0; i < 14; i++ {
			storage[i] = string(line[i])
		}
		for i := 14; i < len(line); i++ {
			if !isNotPacketStarter(storage) {
				fmt.Printf("Value: %v is the packet message.\nLocation: %v\n", string(line[i-1]), i)
				fmt.Print("Packet: ")
				for j := 14; j > 0; j-- {
					fmt.Printf("%v", string(line[i-j]))
				}
				break
			}
			storage = append(storage[1:], string(line[i]))
		}
	}

}
func isNotPacketStarter(storage []string) bool {
	dups := make(map[string]bool)
	for _, v := range storage {
		if dups[v] {
			return true
		}
		dups[v] = true
	}
	return false
}
