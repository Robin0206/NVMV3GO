package Compiler

type FunctionSignatureConverter struct {
}

func (this *FunctionSignatureConverter) processTokens(tokens []Token) []Token {
	var lines = splitToLines(tokens)
	var result []Token
	if lines[0][1].content == "main" {
		result = append(result, generateToken("MAIN", SYSTEM_FUNCTION))
		result = append(result, generateToken("(", BRACE_LEFT))
		result = append(result, generateToken(")", BRACE_RIGHT))
		result = append(result, generateToken(";", SEMICOLON))
		for lineIndex := 1; lineIndex < len(lines)-1; lineIndex++ {
			for tokenIndex := 0; tokenIndex < len(lines[lineIndex]); tokenIndex++ {
				result = append(result, lines[lineIndex][tokenIndex])
			}
		}
		result = append(result, generateToken("MEND", SYSTEM_FUNCTION))
		result = append(result, generateToken("(", BRACE_LEFT))
		result = append(result, generateToken(")", BRACE_RIGHT))
		result = append(result, generateToken(";", SEMICOLON))
	} else {
		result = append(result, generateToken("FUNC", SYSTEM_FUNCTION))
		result = append(result, generateToken("(", BRACE_LEFT))
		result = append(result, generateToken(lines[0][1].content, NAME))
		result = append(result, generateToken(")", BRACE_RIGHT))
		result = append(result, generateToken(";", SEMICOLON))
		for lineIndex := 1; lineIndex < len(lines)-1; lineIndex++ {
			for tokenIndex := 0; tokenIndex < len(lines[lineIndex]); tokenIndex++ {
				result = append(result, lines[lineIndex][tokenIndex])
			}
		}
		result = append(result, generateToken("FEND", SYSTEM_FUNCTION))
		result = append(result, generateToken("(", BRACE_LEFT))
		result = append(result, generateToken(")", BRACE_RIGHT))
		result = append(result, generateToken(";", SEMICOLON))
	}
	return result
}
