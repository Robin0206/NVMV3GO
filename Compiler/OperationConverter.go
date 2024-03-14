package Compiler

import "strings"

type OperationConverter struct {
}

func (this *OperationConverter) processTokens(tokens []Token) []Token {
	var inputLines = splitToLines(tokens)
	var resultLines [][]Token
	for _, line := range inputLines {
		if !this.isOperationLine(line) {
			resultLines = append(resultLines, line)
		} else {
			var arguments = getArguments(line)
			var functionName = generateToken(strings.ToUpper(getFunctionName(line).content), SYSTEM_FUNCTION)
			resultLines = append(resultLines, generateFunctionCall(arguments, functionName))
		}
	}
	var result []Token
	for i := 0; i < len(resultLines); i++ {
		for j := 0; j < len(resultLines[i]); j++ {
			result = append(result, resultLines[i][j])
		}
	}
	return result
}

func generateFunctionCall(arguments []Token, name Token) []Token {

	return []Token{
		name,
		generateToken("(", BRACE_LEFT),
		arguments[0],
		generateToken(",", COMMA),
		arguments[1],
		generateToken(",", COMMA),
		arguments[2],
		generateToken(")", BRACE_LEFT),
		generateToken(";", SEMICOLON),
	}
}

func getFunctionName(line []Token) Token {
	if line[3].tokenType == OPERATOR_PLUS {
		return generateToken("add", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_MINUS {
		return generateToken("sub", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_DIVIDE {
		return generateToken("div", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_MULTIPLY {
		return generateToken("mul", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_XOR {
		return generateToken("xor", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_MODULO {
		return generateToken("mod", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_SINGLE_OR {
		return generateToken("binOr", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_SINGLE_AND {
		return generateToken("binAnd", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_DOUBLE_EQUALS {
		return generateToken("logEq", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_DOUBLE_OR {
		return generateToken("logOr", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_DOUBLE_AND {
		return generateToken("logAnd", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_LESS {
		return generateToken("lessThan", SYSTEM_FUNCTION)
	}
	if line[3].tokenType == OPERATOR_MORE {
		return generateToken("greaterThan", SYSTEM_FUNCTION)
	}
	return generateToken("noop", SYSTEM_FUNCTION)
}

func getArguments(line []Token) []Token {
	return []Token{line[0], line[2], line[4]}
}

func (this *OperationConverter) isOperationLine(line []Token) bool {
	var numOperators = 0
	for _, token := range line {
		if token.tokenType == OPERATOR_PLUS ||
			token.tokenType == OPERATOR_MINUS ||
			token.tokenType == OPERATOR_DIVIDE ||
			token.tokenType == OPERATOR_MULTIPLY ||
			token.tokenType == OPERATOR_XOR ||
			token.tokenType == OPERATOR_MODULO ||
			token.tokenType == OPERATOR_SINGLE_EQUALS ||
			token.tokenType == OPERATOR_SINGLE_OR ||
			token.tokenType == OPERATOR_SINGLE_AND ||
			token.tokenType == OPERATOR_DOUBLE_EQUALS ||
			token.tokenType == OPERATOR_DOUBLE_OR ||
			token.tokenType == OPERATOR_DOUBLE_AND ||
			token.tokenType == OPERATOR_LESS ||
			token.tokenType == OPERATOR_MORE {
			numOperators++
		}

	}
	return numOperators >= 2
}
