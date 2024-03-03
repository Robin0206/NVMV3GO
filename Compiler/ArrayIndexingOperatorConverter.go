package Compiler

import "strings"

type ArrayIndexingOperatorConverter struct {
}

func (this *ArrayIndexingOperatorConverter) processTokens(input []Token) []Token {
	var resultLines [][]Token
	input = changeArrIndexingOperatorTokenTypes(input)
	resultLines = convertArrayIndexingOperationsContainingSetCalls(input, resultLines)
	return flatten(resultLines)
}

func convertArrayIndexingOperationsContainingSetCalls(input []Token, resultLines [][]Token) [][]Token {
	var inputLines = splitToLines(input)
	for _, line := range inputLines {
		if !lineContainsToken("REFA", line) && lineContainsToken("=", line) && lineContainsTokenType(ARR_INDEXING_OP, line) {
			resultLines = append(resultLines, convertArrayIndexingLine(line))

		} else {
			resultLines = append(resultLines, line)
		}
	}
	return resultLines
}

func convertArrayIndexingLine(line []Token) []Token {
	if ArrayIndexingOpIsOnLeftSide(line) {
		return convertArrayIndexingOpOnLeftSide(line)
	} else {
		return convertArrayIndexingOpOnRightSide(line)
	}
}

func convertArrayIndexingOpOnLeftSide(line []Token) []Token {
	var dst = strings.Split(line[0].content, ".")[0]
	var dstIndexVarName = strings.Split(line[0].content, ".")[1]
	var src = line[2].content
	var result []Token
	result = append(result, generateToken("ASET", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(dst, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(dstIndexVarName, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(src, NAME))
	result = append(result, generateToken(")", BRACE_LEFT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func convertArrayIndexingOpOnRightSide(line []Token) []Token {
	var dst = line[0].content
	var src = strings.Split(line[2].content, ".")[0]
	var srcIndexVarName = strings.Split(line[2].content, ".")[1]
	var result []Token
	result = append(result, generateToken("AGET", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(dst, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(src, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(srcIndexVarName, NAME))
	result = append(result, generateToken(")", BRACE_LEFT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func ArrayIndexingOpIsOnLeftSide(line []Token) bool {
	var lineString = ""
	for _, token := range line {
		lineString += token.content + " "
	}
	var sides = strings.Split(lineString, "=")
	return strings.Contains(sides[0], ".")
}

func changeArrIndexingOperatorTokenTypes(input []Token) []Token {
	var result []Token
	for _, token := range input {
		if token.tokenType == NAME && strings.Contains(token.content, ".") {
			token.tokenType = ARR_INDEXING_OP
		}
		result = append(result, token)
	}
	return result
}
