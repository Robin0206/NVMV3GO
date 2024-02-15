package Reading

import (
	"NVMV3/Executor"
	"bufio"
	"fmt"
	"os"
)

func Read(path string) []Executor.NVMCommand {
	var result []Executor.NVMCommand
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		currentCommand := Executor.GenerateNVMCommand(fileScanner.Text())
		result = append(result, currentCommand)
	}
	err = file.Close()
	if err != nil {
		return nil
	}
	return result
}
