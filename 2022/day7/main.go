package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here /////////////

type t_fileNode struct {
	name     string
	isFile   bool
	size     int
	depth    int
	children []*t_fileNode
	parent   *t_fileNode
}

func addNode(node *t_fileNode, parent *t_fileNode) {
	node.depth = parent.depth + 1
	node.parent = parent
	parent.children = append(parent.children, node)
	addSizeToParents(node, node.size)
}

func makeNode(name string, isFile bool, size int) *t_fileNode {
	return &t_fileNode{
		name:     name,
		isFile:   isFile,
		size:     size,
		depth:    0,
		children: make([]*t_fileNode, 0),
		parent:   nil,
	}
}

func addSizeToParents(node *t_fileNode, size int) {
	node.parent.size += size

	if node.parent.parent != nil {
		addSizeToParents(node.parent, size)
	}
}

func printFs(node *t_fileNode) {

	ident := ""

	for i := 1; i < node.depth; i++ {
		ident += "    "
	}

	typeInfo := ""

	if !node.isFile {
		typeInfo = fmt.Sprintf("(dir, size=%d)", node.size)
	} else {
		typeInfo = fmt.Sprintf("(file, size=%d)", node.size)
	}

	fmt.Printf("%s- %s %s\n", ident, node.name, typeInfo)

	for _, child := range node.children {
		printFs(child)
	}
}

func getResultPart1(node *t_fileNode, result *int) {

	if !node.isFile && node.size < 100000 {
		*result += node.size
	}

	for _, child := range node.children {
		getResultPart1(child, result)
	}
}

func getResultPart2(node *t_fileNode, result *int, toDelete int) {

	// init
	if *result == 0 {
		*result = node.size
	}

	if !node.isFile && node.size > toDelete && node.size < *result {
		*result = node.size
	}

	for _, child := range node.children {
		getResultPart2(child, result, toDelete)
	}
}

////////// Main Function ///////////

func main() {

	// Abfrage Inputparameter
	fileName, part, err := utils.Initializer(os.Args)
	if err != nil {
		fmt.Printf("Fehler: %s\n", err)
		return
	}

	// Datei einlesen
	input, err := utils.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Fehler: %s\n", err)
		return
	}

	// gemeinsamer Code

	var fs *t_fileNode
	var currentDir *t_fileNode

	// build tree
	for _, line := range input {
		fmt.Println(line)

		if line == "$ cd /" {
			fs = makeNode("/", false, 0)
			currentDir = fs
			continue
		}

		if line == "$ ls" {
			continue
		}

		if line == "$ cd .." {
			currentDir = currentDir.parent
			continue
		}

		if strings.HasPrefix(line, "dir") {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			dirname := strings.Split(line, " ")[2]
			dir := makeNode(dirname, false, 0)
			addNode(dir, currentDir)
			currentDir = dir
			continue
		} else {
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			filename := strings.Split(line, " ")[1]

			file := makeNode(filename, true, size)
			addNode(file, currentDir)
		}
	}

	println()
	println("print filesystem")
	// print fs
	printFs(fs)

	result := 0

	// Ausführung des Rätselcodes
	if part == "part1" {

		getResultPart1(fs, &result)
		fmt.Printf("Result: %d\n", result)
	}

	if part == "part2" {

		totalDiskspace := 70000000
		neededSpace := 30000000
		unusedSpace := totalDiskspace - fs.size
		toDelete := neededSpace - unusedSpace

		getResultPart2(fs, &result, toDelete)

		fmt.Printf("Result: %d\n", result)
	}
}
