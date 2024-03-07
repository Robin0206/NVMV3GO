package Compiler

import "strconv"

type ForConverter struct {
	counter int
}

func (this *ForConverter) processTokens(input []Token) []Token {
	var result = input
	var forIndex = findForIndex(result)
	for forIndex != -1 {
		result = this.substituteForLoop(result, forIndex)
		forIndex = findForIndex(result)
	}
	return result
}

func (this *ForConverter) substituteForLoop(input []Token, index int) []Token {
	var result []Token
	var forStart = index
	var blockStart, blockEnd = getBlockAfter(index, input)
	for i := 0; i < forStart; i++ {
		result = append(result, input[i])
	}
	var convertedForLoop = this.convertForLoop(input, blockStart, blockEnd, forStart)
	for i := 0; i < len(convertedForLoop); i++ {
		result = append(result, convertedForLoop[i])
	}
	for i := blockEnd + 1; i < len(input); i++ {
		result = append(result, input[i])
	}
	return result
}

func (this *ForConverter) convertForLoop(input []Token, start int, end int, forStart int) []Token {
	var result [][]Token
	var exprBufferName = "____for_expr_buffer_" + strconv.Itoa(this.counter)
	var labelStartName = "____label_for_start_ " + strconv.Itoa(this.counter)
	var labelEndName = "____label_for_end_ " + strconv.Itoa(this.counter)
	this.counter++
	var runVar = input[forStart+1].content

	//append before body
	result = append(result, generateRefaLine(runVar, INT))
	result = append(result, generateRefaLine(exprBufferName, BOOL))
	result = append(result, generateSetLineToken(runVar, input[forStart+3]))

	//=================================

	result = append(result, generateSetLineWithExpression(exprBufferName, []Token{
		generateToken(runVar, NAME),
		generateToken("<", OPERATOR_LESS),
		input[forStart+5],
	}))

	result = append(result,
		[]Token{
			generateToken("BEQ", SYSTEM_FUNCTION),
			generateToken("(", BRACE_LEFT),
			generateToken(exprBufferName, NAME),
			generateToken(",", COMMA),
			generateToken("true", NAME),
			generateToken(",", COMMA),
			generateToken(labelStartName, NAME),
			generateToken(")", BRACE_RIGHT),
			generateToken(";", SEMICOLON),
		})
	result = append(result,
		[]Token{
			generateToken("BEQ", SYSTEM_FUNCTION),
			generateToken("(", BRACE_LEFT),
			generateToken("true", NAME),
			generateToken(",", COMMA),
			generateToken("true", NAME),
			generateToken(",", COMMA),
			generateToken(labelEndName, NAME),
			generateToken(")", BRACE_RIGHT),
			generateToken(";", SEMICOLON),
		})
	//=================================
	result = append(result, generateLabelLine(labelStartName))
	//append body
	var body []Token
	for i := start; i < end; i++ {
		body = append(body, input[i])
	}
	var lineBody = splitToLines(body)
	for _, line := range lineBody {
		result = append(result, line)
	}
	//append after body
	result = append(result, generateSetLineWithExpression(runVar, []Token{
		generateToken(runVar, NAME),
		generateToken("+", OPERATOR_PLUS),
		generateToken("1", NUMBER),
	}))
	result = append(result, generateSetLineWithExpression(exprBufferName, []Token{
		generateToken(runVar, NAME),
		generateToken("<", OPERATOR_LESS),
		input[forStart+5],
	}))
	result = append(result,
		[]Token{
			generateToken("BEQ", SYSTEM_FUNCTION),
			generateToken("(", BRACE_LEFT),
			generateToken(exprBufferName, NAME),
			generateToken(",", COMMA),
			generateToken("true", NAME),
			generateToken(",", COMMA),
			generateToken(labelStartName, NAME),
			generateToken(")", BRACE_RIGHT),
			generateToken(";", SEMICOLON),
		})
	result = append(result, generateLabelLine(labelEndName))
	return flatten(result)
}

func findForIndex(input []Token) int {
	for index, token := range input {
		if token.content == "for" {
			return index
		}
	}
	return -1
}
