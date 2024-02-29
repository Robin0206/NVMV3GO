package Compiler

import "slices"

type MultipleRefaRemover struct {
}

func (this *MultipleRefaRemover) processTokens(input []Token) []Token {
	var inputLines = splitToLines(input)
	var resultLines [][]Token
	var refaNames []string
	for _, line := range inputLines {
		if line[0].content == "REFA" && !slices.Contains(refaNames, line[2].content) {
			resultLines = append(resultLines, line)
			refaNames = append(refaNames, line[2].content)
		}
		if line[0].content != "REFA" {
			resultLines = append(resultLines, line)
		}
	}
	return flatten(resultLines)
}
