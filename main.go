package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// args := os.Args
	// if len(args) != 2 {
	// 	fmt.Println("Usage: gobot <script-file>")
	// 	os.Exit(0)
	// }
	// scriptpath := os.Args[1]
	scriptpath := "test.gb"

	if _, err := os.Stat(scriptpath); os.IsNotExist(err) {
		fmt.Printf("%s not found\n", scriptpath)
		os.Exit(0)
	}

	file, err := os.Open(scriptpath)
	if err != nil {
		fmt.Println("Error opening script file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sp := NewScriptProcessor()
	sp.executeScript(lines)
}
