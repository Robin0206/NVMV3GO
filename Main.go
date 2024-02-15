package main

import (
	"NVMV3/Compiler"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string
	file, err := os.Open("./resources/PrimeSieve.nvm")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	var preprocessor = Compiler.GeneratePreprocessor()
	var preprocessedLines = preprocessor.ProcessLines(lines)
	for i := 0; i < len(preprocessedLines); i++ {
		fmt.Println(preprocessedLines[i])
	}
}
