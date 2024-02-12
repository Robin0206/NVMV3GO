package main

import (
	"bufio"
	"fmt"
	"os"
)

func read(path string) []NVMCommand {
	var result []NVMCommand
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		currentCommand := generateNVMCommand(fileScanner.Text())
		result = append(result, currentCommand)
	}
	err = file.Close()
	if err != nil {
		return nil
	}
	return result
}
