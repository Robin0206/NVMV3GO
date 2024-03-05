package Compiler

type ReturnConverter struct {
}

func (this *ReturnConverter) processTokens(input []Token) []Token {
	//write inputLines to resultLines and substitute returns
	var inputLines = splitToLines(input)
	var resultLines [][]Token
	for _, line := range inputLines {
		if lineContainsToken("return", line) {
			var converted = convertReturnLine(line)
			for _, convertedLine := range converted {
				resultLines = append(resultLines, convertedLine)
			}
		} else {
			resultLines = append(resultLines, line)
		}
	}
	//cut the } at the end
	resultLines = resultLines[:(len(resultLines) - 1)]
	//generate and append label and one }
	var labelLine []Token
	labelLine = append(labelLine, generateToken("LABEL", SYSTEM_FUNCTION))
	labelLine = append(labelLine, generateToken("(", BRACE_LEFT))
	labelLine = append(labelLine, generateToken("____label_function_end", NAME))
	labelLine = append(labelLine, generateToken(")", BRACE_RIGHT))
	labelLine = append(labelLine, generateToken(";", SEMICOLON))
	resultLines = append(resultLines, labelLine)
	var result = flatten(resultLines)
	result = append(result, generateToken("}", CURLY_BRACE_RIGHT))
	return result
}

func convertReturnLine(line []Token) [][]Token {

	//generate returnLine
	var returnLine []Token
	returnLine = append(returnLine, generateToken("RETURN", SYSTEM_FUNCTION))
	returnLine = append(returnLine, generateToken("(", BRACE_LEFT))
	for i := 1; i < len(line)-1; i++ {
		returnLine = append(returnLine, line[i])
	}
	returnLine = append(returnLine, generateToken(")", BRACE_RIGHT))
	returnLine = append(returnLine, generateToken(";", SEMICOLON))

	//generate BEQ
	var beqLine []Token
	beqLine = append(beqLine, generateToken("BEQ", SYSTEM_FUNCTION))
	beqLine = append(beqLine, generateToken("(", BRACE_LEFT))
	beqLine = append(beqLine, generateToken("true", NAME))
	beqLine = append(beqLine, generateToken(",", COMMA))
	beqLine = append(beqLine, generateToken("true", NAME))
	beqLine = append(beqLine, generateToken(",", COMMA))
	beqLine = append(beqLine, generateToken("____label_function_end", NAME))
	beqLine = append(beqLine, generateToken(")", BRACE_RIGHT))
	beqLine = append(beqLine, generateToken(";", SEMICOLON))

	return [][]Token{returnLine, beqLine}
}
