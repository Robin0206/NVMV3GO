package Compiler

import "strconv"

// needs to be placed after:
// with remover
// arg remover
// true/false converter
// returnconverter
type ReturnTypeDeducer struct {
	compiler *SyntacticalSugarCompiler
}

func (this *ReturnTypeDeducer) processTokens(input []Token) []Token {
	var functionName = input[1].content
	var returnType, returnSize = deduceReturnTypeAndSize(input)
	if returnType != -2 { // if returns void dont add
		this.compiler.functionNames = append(this.compiler.functionNames, functionName)
		this.compiler.functionReturnTypes = append(this.compiler.functionReturnTypes, returnType)
		this.compiler.functionReturnSizes = append(this.compiler.functionReturnSizes, returnSize)
	}
	return input
}

// type/size
func deduceReturnTypeAndSize(input []Token) (int, int) {
	var inputLines = splitToLines(input)
	var returnLineIndex = getReturnLineIndex(inputLines)
	if returnLineIndex == -1 { // no returnLine found --> void
		return -2, -2
	}
	return extractTypeAndSize(inputLines, returnLineIndex)
}

func extractTypeAndSize(lines [][]Token, index int) (int, int) {
	var refaLine = lines[getRefaLineIndexFromReturnStatement(lines, index)]
	return extractTypeAndSizeFromRefaLine(refaLine)
}

func extractTypeAndSizeFromRefaLine(line []Token) (int, int) {
	var returnType, _ = strconv.ParseInt(line[4].content, 10, 64)
	var returnSize, _ = strconv.ParseInt(line[6].content, 10, 64)
	return int(returnType), int(returnSize)
}

func getRefaLineIndexFromReturnStatement(lines [][]Token, returnLineIndex int) int {
	var returnVarName = lines[returnLineIndex][2].content
	for index, line := range lines {
		if line[0].content == "REFA" && line[2].content == returnVarName {
			return index
		}
	}
	return -1
}

func getReturnLineIndex(lines [][]Token) int {
	for index, line := range lines {
		if lineContainsToken("RETURN", line) {
			return index
		}
	}
	return -1
}
