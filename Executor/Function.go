package Executor

import (
	"fmt"
	"strconv"
)

type NVMFunction struct {
	name     string
	commands []NVMCommand
	labels   []int
}

func SplitFunctions(inputCommands []NVMCommand) []NVMFunction {
	var current []NVMCommand
	var result []NVMFunction
	for i := 0; i < len(inputCommands); i++ {
		current = append(current, inputCommands[i])
		if inputCommands[i].commandName == "FEND" || inputCommands[i].commandName == "MEND" {
			result = append(result, generateFunction(current))
			current = nil
		}
	}
	return result
}

func SubstituteFunctionIndices(function []NVMFunction) []NVMFunction {
	var result = function
	var functionNames []string
	var counter = 0
	//get functionNames and append the counter to the FUNC Call
	for i := 0; i < len(result); i++ {
		if result[i].commands[0].commandName == "FUNC" {
			result[i].commands[0].arguments = append(result[i].commands[0].arguments, generateNVMArgument(strconv.Itoa(counter)))
			counter++
			functionNames = append(functionNames, result[i].name)
		}
	}
	//substitute in CALL statements
	counter = 0
	for i := 0; i < len(functionNames); i++ { // for every functionName
		var name = functionNames[i]
		for j := 0; j < len(result); j++ { //for every Function
			for k := 0; k < len(result[j].commands); k++ { //for every command in function
				if result[j].commands[k].commandName == "CALL" && result[j].commands[k].functionName == name && len(result[j].commands[k].arguments) != 1 {
					result[j].commands[k].arguments = append(result[i].commands[0].arguments, generateNVMArgument(strconv.Itoa(counter)))
				}
			}
		}
		counter++
	}
	return result
}

func generateFunction(inputCommands []NVMCommand) NVMFunction {
	var result NVMFunction
	result.name = inputCommands[0].functionName
	for i := 0; i < len(inputCommands); i++ {
		result.commands = append(result.commands, inputCommands[i])
	}
	generateLabels(&result)
	return result
}

func generateLabels(function *NVMFunction) {
	for i := 0; i < len(function.commands); i++ {
		if function.commands[i].commandName == "LABEL" {
			function.labels = append(function.labels, i)
			function.commands[i] = GenerateNVMCommand("NOOP")
		}
	}
}

func (this *NVMFunction) print() {
	fmt.Println("================================")
	fmt.Println("Name: " + this.name)
	for i := 0; i < len(this.commands); i++ {
		this.commands[i].print()
	}
}
