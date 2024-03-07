package Compiler

// needs to be placed after the return converter and the return type deducer
type NestedFunctionCallConverter struct {
	counter  int
	compiler *SyntacticalSugarCompiler
}

func (this *NestedFunctionCallConverter) processTokens(input []Token) []Token {
	var resultLines = splitToLines(input)
	var nestedFunctionCoords = findNestedFunctionCall(resultLines)
	for nestedFunctionCoords != -1 {
		resultLines = substituteNestedFunctionCall(resultLines, nestedFunctionCoords)
	}
	return flatten(resultLines)
}

func substituteNestedFunctionCall(wholeFunction [][]Token, coords int) [][]Token {
	var resultLines [][]Token
	for i := 0; i < coords; i++ {
		resultLines = append(resultLines, wholeFunction[i])
	}
	var convertedNestedFunctionCall [][]Token
	convertedNestedFunctionCall = convertNestedFunctionCall(wholeFunction, coords)
	for _, line := range convertedNestedFunctionCall {
		resultLines = append(resultLines, line)
	}
	for i := coords + 1; i < len(wholeFunction); i++ {
		resultLines = append(resultLines, wholeFunction[i])
	}
	return resultLines
}

func convertNestedFunctionCall(function [][]Token, coords int) [][]Token {

	return [][]Token{}
}

func findNestedFunctionCall(lines [][]Token) int {
	for index, line := range lines {
		if isNestedFunctionCall(line) {
			return index
		}
	}
	return -1
}

func isNestedFunctionCall(line []Token) bool {
	return lineIsFunctionCall(line) && functionArgsContainFunctionCall(line)
}

func functionArgsContainFunctionCall(line []Token) bool {
	var args = extractArgumentsFromFunctionCall(line)
	for _, arg := range args {
		if lineContainsTokenType(BRACE_LEFT, arg) {
			return true
		}
	}
	return false
}
