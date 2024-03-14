package Compiler

import (
	"fmt"
	"strconv"
	"strings"
)

type InlineOperationRemover struct {
	counter int
}

func (this *InlineOperationRemover) processTokens(input []Token) []Token {
	var resultLines = splitToLines(input)
	var lineIndex = findInlineOperation(resultLines)
	for lineIndex != -1 {
		resultLines = this.substituteInlineOp(resultLines, lineIndex)
		lineIndex = findInlineOperation(resultLines)
	}

	return flatten(resultLines)
}

func (this *InlineOperationRemover) substituteInlineOp(lines [][]Token, index int) [][]Token {
	var result [][]Token
	for i := 0; i < index; i++ {
		result = append(result, lines[i])
	}
	var converted = this.convertInlineOp(lines[index], lines)
	for _, line := range converted {
		result = append(result, line)
	}
	for i := index + 1; i < len(lines); i++ {
		result = append(result, lines[i])
	}
	return result
}

func (this *InlineOperationRemover) convertInlineOp(lineToConvert []Token, wholeFunction [][]Token) [][]Token {
	var functionName = lineToConvert[0]
	var rawArgs = extractArgumentsFromFunctionCall(lineToConvert)
	var bufferName = "____inline_op_buffer_" + strconv.Itoa(this.counter)
	var result [][]Token
	this.counter++
	var expression []Token
	var expressionIndex = -1
	for index, arg := range rawArgs {
		if containsOperatorOtherThanSingleEquals(arg) {
			expression = arg
			expressionIndex = index
			break
		}
	}
	//append refa and set
	var expressionType = determineBufferTypeFromExpression(expression, wholeFunction)
	result = append(result, generateRefaLine(bufferName, expressionType))
	result = append(result, generateSetLineWithExpression(bufferName, expression))
	//append last line
	var convertedFunctionCall []Token
	convertedFunctionCall = append(convertedFunctionCall, functionName)
	convertedFunctionCall = append(convertedFunctionCall, generateToken("(", BRACE_LEFT))
	for index, arg := range rawArgs {
		if index == expressionIndex {
			convertedFunctionCall = append(convertedFunctionCall, generateToken(bufferName, NAME))
			convertedFunctionCall = append(convertedFunctionCall, generateToken(",", COMMA))
		} else {
			for _, token := range arg {
				convertedFunctionCall = append(convertedFunctionCall, token)
			}
			convertedFunctionCall = append(convertedFunctionCall, generateToken(",", COMMA))
		}
	}
	convertedFunctionCall = convertedFunctionCall[:(len(convertedFunctionCall) - 1)]
	convertedFunctionCall = append(convertedFunctionCall, generateToken(")", BRACE_RIGHT))
	convertedFunctionCall = append(convertedFunctionCall, generateToken(";", SEMICOLON))
	result = append(result, convertedFunctionCall)
	return result
}

func determineBufferTypeFromExpression(expression []Token, wholeFunction [][]Token) int {
	if lineContainsOperatorThatAlwaysOutputsBool(expression) {
		return BOOL
	} else {
		for _, token := range expression {
			if token.tokenType == NUMBER {
				if strings.Contains(token.content, ".") {
					return REAL
				} else {
					return INT
				}
			}
			if token.tokenType == NAME {
				return getType(token, wholeFunction)
			}
		}
	}
	fmt.Println("ERROR: InlineOperationRemover cant determine type!")
	return -1
}

func extractArgumentsFromFunctionCall(line []Token) [][]Token {
	var result [][]Token
	var inArgs = false

	var currentLine []Token
	for _, token := range line {
		if token.tokenType == BRACE_RIGHT {
			result = append(result, currentLine)
			break
		}
		if inArgs {
			if token.tokenType != COMMA {
				currentLine = append(currentLine, token)
			} else {
				result = append(result, currentLine)
				currentLine = nil
			}
		}
		if token.tokenType == BRACE_LEFT {
			inArgs = true
		}
	}

	return result
}

func findInlineOperation(lines [][]Token) int {
	for index, line := range lines {
		if lineIsFunctionCall(line) && lineContainsOperator(line) {
			return index
		}
	}
	return -1
}
