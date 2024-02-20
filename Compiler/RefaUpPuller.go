package Compiler

type RefaUpPuller struct {
}

func (this *RefaUpPuller) processTokens(tokens []Token) []Token {
	var result []Token
	var resultLines [][]Token
	var refaLines [][]Token
	var bodyWithoutRefaLines [][]Token
	var lines = splitToLines(tokens)

	for i := 0; i < len(lines); i++ {
		if lines[i][0].content == "REFA" {
			refaLines = append(refaLines, lines[i])
		} else {
			bodyWithoutRefaLines = append(bodyWithoutRefaLines, lines[i])
		}
	}
	resultLines = append(resultLines, bodyWithoutRefaLines[0])
	for i := 0; i < len(refaLines); i++ {
		resultLines = append(resultLines, refaLines[i])
	}
	for i := 1; i < len(bodyWithoutRefaLines); i++ {
		resultLines = append(resultLines, bodyWithoutRefaLines[i])
	}
	for i := 0; i < len(resultLines); i++ {
		for j := 0; j < len(resultLines[i]); j++ {
			result = append(result, resultLines[i][j])
		}
	}
	return result
}
