package Compiler

type SyntacticalSugarStage interface {
	processTokens(tokens []Token) []Token
}

func splitToLines(tokens []Token) [][]Token {
	var result [][]Token
	var current []Token
	for _, token := range tokens {
		current = append(current, token)
		if token.tokenType == SEMICOLON || token.tokenType == CURLY_BRACE_LEFT || token.tokenType == CURLY_BRACE_RIGHT {
			result = append(result, current)
			current = nil
		}
	}
	return result
}
