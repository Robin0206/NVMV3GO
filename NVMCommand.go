package main

import (
	"fmt"
	"strings"
)

type NVMCommand struct {
	commandName  string
	functionName string // only used if it is a FUNC or an CALL command
	arguments    []NVMArgument
	commandindex uint32
}

func generateNVMCommand(inputLine string) NVMCommand {
	var splitted = strings.Split(inputLine, " ")
	var commandName = splitted[0]
	var args []string
	var result NVMCommand
	if commandName != "CALL" && commandName != "FUNC" && commandName != "MAIN" {
		for i := 1; i < len(splitted); i++ {
			if splitted[i] != "" {
				args = append(args, splitted[i])
			}
		}
		result.commandName = commandName
		for i := 0; i < len(args); i++ {
			result.arguments = append(result.arguments, generateNVMArgument(args[i]))
		}
		result.commandindex = getCommandIndex(commandName)
		return result
	} else {
		if commandName == "MAIN" {
			result.commandName = commandName
			result.commandindex = getCommandIndex(commandName)
			return result
		} else {
			result.commandName = commandName
			result.functionName = splitted[1]
			result.commandindex = getCommandIndex(commandName)
			return result
		}
	}
}
func (this *NVMCommand) print() {
	fmt.Print("| " + this.commandName + " | ")
	for i := 0; i < len(this.arguments); i++ {
		if this.arguments[i].valueType == 0 {
			fmt.Printf("%d | ", this.arguments[i].integer)
		} else {
			fmt.Printf("%d | ", this.arguments[i].integer)
		}
	}
	fmt.Println()

}

func getCommandIndex(commandName string) uint32 {
	switch commandName {
	case "NOOP":
		return 0
	case "REFA":
		return 1
	case "MOV":
		return 2
	case "SET":
		return 3
	case "SETV":
		return 4
	case "ASET":
		return 5
	case "AGET":
		return 6
	case "CPY":
		return 7
	case "ADD":
		return 8
	case "SUB":
		return 9
	case "MUL":
		return 10
	case "DIV":
		return 11
	case "BINOR":
		return 12
	case "BINAND":
		return 13
	case "BINXOR":
		return 14
	case "BINNOT":
		return 15
	case "LESSTHAN":
		return 16
	case "GREATERTHAN":
		return 17
	case "LOGOR":
		return 18
	case "LOGAND":
		return 19
	case "LOGNOT":
		return 20
	case "LOGEQ":
		return 21
	case "PRINT":
		return 22
	case "RETURN":
		return 23
	case "CALL":
		return 24
	case "PARG":
		return 25
	case "BEQ":
		return 26
	case "LABEL":
		return 27
	case "FUNC":
		return 28
	case "FEND":
		return 29
	case "GARG":
		return 30
	case "MAIN":
		return 31
	case "MEND":
		return 32
	case "RETG":
		return 33

	}
	return 42
}
