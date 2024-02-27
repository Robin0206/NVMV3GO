package main

import (
	"NVMV3/Compiler"
	"NVMV3/Executor"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string
	file, err := os.Open("./resources/Test.num")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	var compiler = Compiler.GenerateSyntacticalSugarCompiler()
	var commands = compiler.Compile(lines)
	var executor = Executor.GenerateExecutor(Executor.SplitFunctions(commands))
	executor.Run()

}
