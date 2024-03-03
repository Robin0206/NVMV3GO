package Compiler

type UserFunctionCallConverter struct {
	counter int
}

func (this *UserFunctionCallConverter) processTokens(input []Token) []Token {
	var resultLines = splitToLines(input)
	var functionCallIndex = findUserFunctionLineIndex(resultLines)
	for functionCallIndex != -1 {
		resultLines = substituteUserFunctionCall(resultLines, functionCallIndex)
		functionCallIndex = findUserFunctionLineIndex(resultLines)
	}
	return flatten(resultLines)
}

func substituteUserFunctionCall(lines [][]Token, index int) [][]Token {
	var result [][]Token
	for i := 0; i < index; i++ {
		result = append(result, lines[i])
	}
	var converted = convertUserFunctionCallLine(lines[index])
	for _, line := range converted {
		result = append(result, line)
	}
	for i := index + 1; i < len(lines); i++ {
		result = append(result, lines[i])
	}
	return result
}

func convertUserFunctionCallLine(line []Token) [][]Token {
	if lineContainsTokenType(OPERATOR_SINGLE_EQUALS, line) {
		return convertUserFunctionCallWithSingleEquals(line)
	} else {
		return convertUserFunctionCallWithoutSingleEquals(line)
	}
}

func convertUserFunctionCallWithoutSingleEquals(line []Token) [][]Token {
	var funcName = line[0]
	var args = extractArgumentsFromFunctionCall(line)
	var result [][]Token
	for _, arg := range args {
		result = append(result, generatePargLine(arg))
	}
	result = append(result, generateCallLine(funcName))
	return result
}

func convertUserFunctionCallWithSingleEquals(line []Token) [][]Token {
	var funcName = line[2]
	var dstName = line[0]
	var functionCall = line[2:]
	var args = extractArgumentsFromFunctionCall(functionCall)
	var result [][]Token
	for _, arg := range args {
		result = append(result, generatePargLine(arg))
	}
	result = append(result, generateCallLine(funcName))
	result = append(result, generateRetgLine(dstName))
	return result
}

func findUserFunctionLineIndex(lines [][]Token) int {
	for index, line := range lines {
		for i := 0; i < len(line)-1; i++ {
			if line[i].tokenType == NAME && line[i+1].tokenType == BRACE_LEFT {
				return index
			}
		}
	}
	return -1
}
