package Compiler

type InlineNumberConverter struct {
	counter int
}

func (this *InlineNumberConverter) processTokens(tokens []Token) []Token {
	var result []Token
	var inputLines = splitToLines(tokens)
	var numberLineIndex = containsInlineNumber(inputLines)
	for numberLineIndex != -1 {
		inputLines = substituteNumberLine(&inputLines)
		numberLineIndex = containsInlineNumber(inputLines)
	}
	return result
}

// TODO
func substituteNumberLine(lines *[][]Token) [][]Token {
	return *(lines)
}

func containsInlineNumber(lines [][]Token) int {
	for i := 0; i < len(lines); i++ {
		if lineContainsInlineNumber(lines[i]) && containsOperatorOtherThanSingleEquals(lines[i]) {
			return i
		}
	}
	return -1
}

func lineContainsInlineNumber(line []Token) bool {
	for i := 0; i < len(line); i++ {
		if line[i].tokenType == NUMBER {
			return true
		}
	}
	return false
}
