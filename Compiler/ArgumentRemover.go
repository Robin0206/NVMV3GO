package Compiler

import (
	"strings"
)

type ArgumentRemover struct {
}

func (this *ArgumentRemover) processTokens(input []Token) []Token {
	var inputLines = splitToLines(input)
	var firstLine = inputLines[0]
	if firstLine[2].tokenType != BRACE_LEFT {
		return input
	}
	var refaLines [][]Token
	var newRefaLines [][]Token
	var gargLines [][]Token
	var rest [][]Token
	var inRefas = true
	var resultLines [][]Token
	//fill the arrays
	for i := 1; i < len(inputLines); i++ {
		if inputLines[i][0].content != "REFA" {
			inRefas = false
		}
		if inRefas {
			refaLines = append(refaLines, inputLines[i])
		}
		if !inRefas {
			rest = append(rest, inputLines[i])
		}
	}
	var arguments = extractArguments(firstLine)
	for index, arg := range arguments {
		newRefaLines = append(newRefaLines, arg.generateRefaLine())
		gargLines = append(gargLines, arg.generateGargLine(index))
	}
	firstLine = nil
	firstLine = append(firstLine, input[0])
	firstLine = append(firstLine, input[1])
	firstLine = append(firstLine, generateToken("{", CURLY_BRACE_LEFT))
	//construct the result
	resultLines = append(resultLines, firstLine)
	for _, token := range refaLines {
		resultLines = append(resultLines, token)
	}
	for _, token := range newRefaLines {
		resultLines = append(resultLines, token)
	}
	for _, token := range gargLines {
		resultLines = append(resultLines, token)
	}
	for _, token := range rest {
		resultLines = append(resultLines, token)
	}
	return flatten(resultLines)
}

func extractArguments(line []Token) []ArgumentPair {
	var argumentsString string
	var write = false
	for _, token := range line {
		if token.tokenType == BRACE_RIGHT {
			break
		}
		if write {
			argumentsString += token.content + " "
		}
		if token.tokenType == BRACE_LEFT {
			write = true
		}
	}
	var splittedAtCommaArguments = strings.Split(argumentsString, ",")

	var result []ArgumentPair
	for _, rawPair := range splittedAtCommaArguments {

		result = append(result, generateArgumentPair(rawPair))
	}
	return result
}
