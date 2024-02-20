package Compiler

type BracketAndCommaRemover struct {
}

func (this *BracketAndCommaRemover) processTokens(tokens []Token) []Token {
	var result []Token
	for i := 0; i < len(tokens); i++ {
		if !(tokens[i].tokenType == BRACE_LEFT || tokens[i].tokenType == BRACE_RIGHT || tokens[i].tokenType == COMMA) {
			result = append(result, tokens[i])
		}
	}
	return result
}
