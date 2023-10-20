package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Node struct {
	name      string
	size      int
	directory string
	children  []*Node
	parent    *Node
	level     int
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}
func day7p1() {
	f, err := os.Open("data/07.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	defer timeTrack(time.Now(), "input loader")

	scanner := bufio.NewScanner(f)
	level := 0
	root := &Node{name: "/", size: 0, directory: "/", children: []*Node{}}
	parentNode := root
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("%v\n", line)
		if line == "$ cd /" {
			fmt.Println("Starting on parent directory")
		}
		re := regexp.MustCompile(`^\$ cd .+`)
		if re.MatchString(line) && line != "$ cd /" && line != "$ cd .." {
			var name string
			fmt.Sscanf(line, "$ cd %s", &name)
			level++
			// fmt.Println("Going into directory", line)
			parentNode = findChild(*parentNode, name)
		}
		if line == "$ cd .." {
			level--
			if level > 0 {
				parentNode = parentNode.parent
			}
			if level == 0 {
				parentNode = root
			}
			// fmt.Printf("Going back to %v directory\n", parentNode.name)
		}
		if line[0:3] == "dir" {
			newNode := Node{name: line[4:], size: 0, directory: line[4:] + "/", children: []*Node{}, level: (level + 1), parent: parentNode}
			parentNode.AddChild(&newNode)
			// fmt.Printf("Added a directory: %v to parent %v\n", newNode.name, parentNode.name)
		}
		reg := regexp.MustCompile(`^\d+ .+`)
		if reg.MatchString(line) {
			var size int
			var name string
			fmt.Sscanf(line, "%d %s", &size, &name)
			newNode := Node{name: name, size: size, directory: parentNode.directory + name + "/", children: nil, level: (level + 1), parent: parentNode}
			parentNode.AddChild(&newNode)
			// fmt.Printf("Added a file: %v with size %d to parent %v in directory %v\n", name, size, parentNode.name, newNode.directory)

		}
	}
	recursiveWriteSize(root)
	recursiveWrite(root)
	// totalSize := recursiveWrite100k(root)
	// fmt.Printf("Total size is: %d\n", totalSize)
	toDelete := 30_000_000 - (70_000_000 - 44_274_331)
	smallestNodeToDelete := recursiveDeleteSmallestPossible(root, toDelete, nil)
	fmt.Printf("Smallest node to delete is: %s and it's size is: %d\n", smallestNodeToDelete.name, smallestNodeToDelete.size)

}
func findChild(pn Node, name string) *Node {
	for _, child := range pn.children {
		if child.name == name {
			return child
		}
	}
	return nil
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
func recursiveWriteSize(node *Node) int {
	cumulative := node.size
	for _, child := range node.children {
		cumulative += recursiveWriteSize(child)
	}
	node.size = cumulative
	return cumulative
}
func recursiveDeleteSmallestPossible(node *Node, toDelete int, currentSmallestNode *Node) *Node {

	if node.size >= toDelete {
		if currentSmallestNode == nil || node.size < currentSmallestNode.size {
			currentSmallestNode = node
		}
	}
	for _, child := range node.children {
		currentSmallestNode = recursiveDeleteSmallestPossible(child, toDelete, currentSmallestNode)
	}

	return currentSmallestNode
}
func recursiveWrite100k(node *Node) int {
	totalSize := 0
	if node.size < 100_000 && len(node.children) != 0 {
		totalSize += node.size
		fmt.Printf("Node %s is less thank 100k: %d\n", node.name, node.size)
	}
	for _, child := range node.children {
		totalSize += recursiveWrite100k(child)
	}
	return totalSize
}

func recursiveWrite(node *Node) {
	if node.level == 1 {
		fmt.Printf("====\n")
	}
	if node.level == 0 {
		if node.size > 0 {
			fmt.Printf("| %s %d\n", node.name, node.size)
		} else {
			fmt.Printf("| %s \n", node.name)
		}

	} else {
		line := strings.Repeat("--", node.level)
		if node.size > 0 {
			fmt.Printf("|"+line+"%s %d\n", node.name, node.size)
		} else {
			fmt.Printf("|"+line+"%s \n", node.name)
		}

	}
	if node.level == 1 {
		fmt.Printf("====\n")
	}
	for _, child := range node.children {
		recursiveWrite(child)
	}
}

// |-- sdsadas 123131kb
// |-- sdasdas 1231kb
// |	|-- sdsadasd 12 kb
// |	|-- sdasdasd 32 kb
// |-- ababa 231lb
