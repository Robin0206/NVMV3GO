package Compiler

type SetConverter struct {
}

func (this *SetConverter) processTokens(tokens []Token) []Token {
	var result []Token
	var lines = splitToLines(tokens)
	var resultLines [][]Token
	for _, line := range lines {
		if lineContainsToken("space", line) {
			for _, token := range line {
				token.Print()
			}
		}
		if isSetLine(line) {
			resultLines = append(resultLines, convertSetLine(line))
			continue
		}
		if isSetVLine(line) {
			resultLines = append(resultLines, convertSetVLine(line))
			continue
		}
		resultLines = append(resultLines, line)
	}
	for i := 0; i < len(resultLines); i++ {
		for j := 0; j < len(resultLines[i]); j++ {
			result = append(result, resultLines[i][j])
		}
	}
	return result
}

func convertSetVLine(line []Token) []Token {
	var result []Token
	var dst = line[0]
	var src = line[2]
	result = append(result, generateToken("SETV", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(dst.content, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(src.content, NAME))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func containsOperatorOtherThanSingleEquals(line []Token) bool {
	for _, token := range line {
		if token.tokenType == OPERATOR_PLUS ||
			token.tokenType == OPERATOR_MINUS ||
			token.tokenType == OPERATOR_DIVIDE ||
			token.tokenType == OPERATOR_MULTIPLY ||
			token.tokenType == OPERATOR_XOR ||
			token.tokenType == OPERATOR_MODULO ||
			token.tokenType == OPERATOR_SINGLE_OR ||
			token.tokenType == OPERATOR_SINGLE_AND ||
			token.tokenType == OPERATOR_DOUBLE_EQUALS ||
			token.tokenType == OPERATOR_DOUBLE_OR ||
			token.tokenType == OPERATOR_DOUBLE_AND ||
			token.tokenType == OPERATOR_LESS ||
			token.tokenType == OPERATOR_MORE {
			return true
		}

	}
	return false
}

func convertSetLine(line []Token) []Token {
	var result []Token
	var dst = line[0]
	var src = line[2]
	result = append(result, generateToken("SET", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(dst.content, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(src.content, NUMBER))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func isSetVLine(line []Token) bool {
	return !containsOperatorOtherThanSingleEquals(line) && !containsNumber(line) && containsSingleEquals(line)
}

func containsSingleEquals(line []Token) bool {
	for i := 0; i < len(line); i++ {
		if line[i].tokenType == OPERATOR_SINGLE_EQUALS {
			return true
		}
	}
	return false
}

func containsNumber(line []Token) bool {
	for i := 0; i < len(line); i++ {
		if line[i].tokenType == NUMBER {
			return true
		}
	}
	return false
}

func isSetLine(line []Token) bool {
	return !containsOperatorOtherThanSingleEquals(line) && containsNumber(line) && containsSingleEquals(line)
}
