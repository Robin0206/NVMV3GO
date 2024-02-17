package Compiler

import "strconv"

type LoopLabelConverter struct {
	counter int
}

func generateLoopLabelConverter() *LoopLabelConverter {
	var result LoopLabelConverter
	result.counter = 0
	return &result
}

func (this *LoopLabelConverter) processTokens(tokens []Token) []Token {
	this.counter = 0
	var nameIndex = this.thereExistsALabelWithFollowingName(tokens)
	for nameIndex != -1 {
		tokens = this.substituteLabelWithNumber(tokens, nameIndex)
		nameIndex = this.thereExistsALabelWithFollowingName(tokens)
	}
	return tokens
}

func (this *LoopLabelConverter) thereExistsALabelWithFollowingName(tokens []Token) int {
	for i := 1; i < len(tokens); i++ {
		if tokens[i-1].content == "LABEL" && tokens[i].tokenType != NUMBER {
			return i
		}
	}
	return -1
}

func (this *LoopLabelConverter) substituteLabelWithNumber(tokens []Token, index int) []Token {
	var labelName = tokens[index].content
	var counterStr = strconv.Itoa(this.counter)
	tokens[index].content = counterStr
	tokens[index].tokenType = NUMBER
	for i := 0; i < len(tokens); i++ {
		if tokens[i].content == labelName {
			tokens[i].content = counterStr
		}
	}
	this.counter++
	return tokens
}
