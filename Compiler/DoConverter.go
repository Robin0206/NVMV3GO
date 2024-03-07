package Compiler

import (
	"strconv"
)

type DoConverter struct {
	counter int
}

func (this *DoConverter) processTokens(input []Token) []Token {
	var result = input
	var doIndex = findDoIndex(result)
	for doIndex != -1 {
		result = this.substituteDoBlock(result, doIndex)
		doIndex = findDoIndex(result)
	}
	return result
}

func (this *DoConverter) substituteDoBlock(input []Token, index int) []Token {
	var doIndex = index
	var doBlockStart, doBlockEnd = getBlockAfter(doIndex, input)
	var result []Token

	for i := 0; i < doIndex; i++ {
		result = append(result, input[i])
	}

	var converted = this.convertDoLoop(input, doBlockStart, doBlockEnd, doIndex)
	for _, token := range converted {
		result = append(result, token)
	}

	for i := doBlockEnd; i < len(input); i++ {
		result = append(result, input[i])
	}

	return result
}

func (this *DoConverter) convertDoLoop(input []Token, start int, end int, doIndex int) []Token {
	var result []Token
	var howManyTimes = input[doIndex+1]
	var bufferName = "____do_buffer_" + strconv.Itoa(this.counter)
	this.counter++

	//append the refa line
	var refaLine = generateRefaLine(bufferName, INT)
	for _, token := range refaLine {
		result = append(result, token)
	}

	//build and append the for line
	var forLine = []Token{
		generateToken("for", KEYWORD),
		generateToken(bufferName, NAME),
		generateToken("from", KEYWORD),
		generateToken("0", NUMBER),
		generateToken("to", KEYWORD),
		howManyTimes,
	}
	for _, token := range forLine {
		result = append(result, token)
	}

	//append {
	result = append(result, generateToken("{", CURLY_BRACE_LEFT))

	//append body
	for i := start; i < end; i++ {
		result = append(result, input[i])
	}

	return result
}

func findDoIndex(result []Token) int {
	for index, token := range result {
		if token.content == "do" {
			return index
		}
	}
	return -1
}
