package Compiler

import (
	"fmt"
	"strconv"
)

type NestedFunctionCallConverter struct {
	counter  int
	compiler *SyntacticalSugarCompiler
}

func (this *NestedFunctionCallConverter) processTokens(input []Token) []Token {
	var resultLines = splitToLines(input)
	var nestedUserFunctionCallIndex = findNestedFunctionCall(resultLines)
	for nestedUserFunctionCallIndex != -1 {
		resultLines = this.substituteNestedUserFunctionCall(resultLines, nestedUserFunctionCallIndex)
		nestedUserFunctionCallIndex = findNestedFunctionCall(resultLines)
	}
	return flatten(resultLines)
}

func (this *NestedFunctionCallConverter) substituteNestedUserFunctionCall(lines [][]Token, index int) [][]Token {
	var resultLines [][]Token
	for i := 0; i < index; i++ {
		resultLines = append(resultLines, lines[i])
	}
	var convertedFunctionCall = this.convertNestedFunctionCall(lines, index)
	for _, line := range convertedFunctionCall {
		resultLines = append(resultLines, line)
	}
	convertedFunctionCall = this.convertNestedFunctionCall(lines, index)
	for i := index + 1; i < len(lines); i++ {
		resultLines = append(resultLines, lines[i])
	}
	return resultLines
}

func (this *NestedFunctionCallConverter) convertNestedFunctionCall(lines [][]Token, index int) [][]Token {
	var result [][]Token
	var argLine = extractArgLine(lines[index])
	var args = splitAt(argLine)
	var bufferName = "____nested_function_call_buffer_" + strconv.Itoa(this.counter)
	this.counter++
	//determine nestedFunction, its returnType, returnSize and index in argList
	var nestedFunction []Token
	var nestedFunctionReturnType, nestedFunctionReturnSize int
	var argIndex int
	for i, arg := range args {
		if lineIsFunctionCall(arg) {
			nestedFunction = arg
			nestedFunctionReturnType, nestedFunctionReturnSize = this.compiler.getTypeFromFunctionName(nestedFunction[0].content)
			argIndex = i
			if nestedFunctionReturnType == -1 {
				fmt.Println("ERROR: Nested Function Call doesnt return anything")
				return result
			}
			break
		}
	}

	//append refa and set lines
	result = append(result, generateRefaArrLine(bufferName, nestedFunctionReturnType, nestedFunctionReturnSize))
	result = append(result, generateSetLineWithExpression(bufferName, nestedFunction))

	//rebuild and append original line
	var convertedOriginalLine []Token
	for _, token := range lines[index] {
		convertedOriginalLine = append(convertedOriginalLine, token)
		if token.tokenType == BRACE_LEFT {
			break
		}
	}
	for i, arg := range args {
		if i == argIndex {
			convertedOriginalLine = append(convertedOriginalLine, generateToken(bufferName, NAME))
		} else {
			for _, token := range arg {
				convertedOriginalLine = append(convertedOriginalLine, token)
			}
		}
		convertedOriginalLine = append(convertedOriginalLine, generateToken(",", COMMA))
	}
	convertedOriginalLine = convertedOriginalLine[:(len(convertedOriginalLine) - 1)]
	convertedOriginalLine = append(convertedOriginalLine, generateToken(")", BRACE_RIGHT))
	convertedOriginalLine = append(convertedOriginalLine, generateToken(";", SEMICOLON))
	result = append(result, convertedOriginalLine)
	return result
}

func findNestedFunctionCall(lines [][]Token) int {
	for i, line := range lines {
		if lineIsNestedUserFunctionCall(line) {
			return i
		}
	}
	return -1
}

func lineIsNestedUserFunctionCall(line []Token) bool {
	if !lineContainsTokenType(OPERATOR_SINGLE_EQUALS, line) && !lineIsFunctionCall(line) {
		return false
	}

	//extract argLine
	var argLine = extractArgLine(line)

	//return true if argLineContainsFunctionCall
	for i := 0; i < len(argLine)-1; i++ {
		if argLine[i].tokenType == NAME && argLine[i+1].tokenType == BRACE_LEFT {
			return true
		}
	}

	return false
}

func extractArgLine(line []Token) []Token {
	var result []Token
	var alreadyHitBraceLeft = false
	for i := 0; i < len(line)-2; i++ {
		if alreadyHitBraceLeft {
			result = append(result, line[i])
		}
		if line[i].tokenType == BRACE_LEFT {
			alreadyHitBraceLeft = true
		}
	}
	return result
}
