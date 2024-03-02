package Compiler

import (
	"strconv"
	"strings"
)

type ArgumentPair struct {
	name        string
	argType     int
	isArray     bool
	size        int
	isResizable bool
}

func (this *ArgumentPair) generateRefaLine() []Token {
	var result []Token
	result = append(result, generateToken("REFA", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(this.name, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(strconv.Itoa(this.argType), NUMBER))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(strconv.Itoa(this.size), NUMBER))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func (this *ArgumentPair) generateGargLine(index int) []Token {
	var result []Token
	result = append(result, generateToken("GARG", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(this.name, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(strconv.Itoa(index), NUMBER))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func generateArgumentPair(input string) ArgumentPair {
	input = strings.TrimSpace(input)
	var result ArgumentPair
	result.name = getArgName(input)
	result.argType = getArgType(input)
	result.size = getArgArrSize(input)
	result.isArray = result.size != 0
	if result.size == -1 {
		result.isResizable = true
	} else {
		result.isResizable = false
	}
	return result
}

// returns 0 if it isnt an array
// returns -1 if it is an resizable array
// else returns the size
func getArgArrSize(input string) int {
	var t = strings.Split(strings.Split(input, ",")[0], " ")[0]
	if !strings.Contains(t, ".") {
		return 0
	}
	if t[len(t)-1] == '.' {
		return -1
	} else {
		var splitted = strings.Split(t, ".")
		var intValue, _ = strconv.ParseInt(splitted[1], 10, 64)
		return int(intValue)
	}

}

func getArgType(input string) int {
	var argTypeString string = ""
	if strings.Contains(input, ".") {
		argTypeString = strings.Split(input, ".")[0]
	} else {
		argTypeString = strings.Split(input, " ")[0]
	}
	switch argTypeString {
	case "bool":
		return 0
	case "byte":
		return 1
	case "int":
		return 2
	case "real":
		return 3
	default:
		return 42
	}
}

func getArgName(input string) string {
	return strings.Split(input, " ")[1]
}
