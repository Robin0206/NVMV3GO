package Compiler

import "strconv"

type VariableNameConverter struct {
	counter int
}

func (this *VariableNameConverter) processTokens(tokens []Token) []Token {
	this.counter = 0
	var nameIndex = thereExistsARefaWithFollowingName(tokens)
	for nameIndex != -1 {
		tokens = this.substituteNameWithNumber(tokens, nameIndex)
		nameIndex = thereExistsARefaWithFollowingName(tokens)
	}
	return tokens
}

func generateVariableNameConverter() *VariableNameConverter {
	var result VariableNameConverter
	result.counter = 0
	return &result
}

func (this *VariableNameConverter) substituteNameWithNumber(tokens []Token, index int) []Token {
	var counterStr = strconv.Itoa(this.counter)
	var result []Token
	tokens[index].content = counterStr
	tokens[index].tokenType = NUMBER
	for _, token := range tokens {
		if token.tokenType != NAME && token.content == tokens[index].content {
			result = append(result, generateToken(counterStr, NUMBER))
		} else {
			result = append(result, token)
		}
	}
	this.counter++
	return result
}

func thereExistsARefaWithFollowingName(tokens []Token) int {
	for i := 1; i < len(tokens); i++ {
		if tokens[i-1].content == "REFA" && tokens[i].tokenType == NAME {
			return i
		}
	}
	return -1
}
