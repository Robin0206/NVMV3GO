package Compiler

type TrueAndFalseConverter struct {
}

func (this *TrueAndFalseConverter) processTokens(input []Token) []Token {
	var result []Token
	var firstLineAfterRefaIndex = 0
	var inputLines = splitToLines(input)
	var resultLines [][]Token
	for i := 0; i < len(inputLines); i++ {
		if i != 0 && inputLines[i][0].content != "REFA" {
			firstLineAfterRefaIndex = i
			break
		}
		resultLines = append(resultLines, inputLines[i])
	}
	//add refaLines
	resultLines = append(resultLines, generateRefaLine("true", BOOL))
	resultLines = append(resultLines, generateRefaLine("false", BOOL))
	//add setLines
	var trueLine []Token
	var falseLine []Token
	trueLine = append(trueLine, generateToken("true", NAME))
	trueLine = append(trueLine, generateToken("=", OPERATOR_SINGLE_EQUALS))
	trueLine = append(trueLine, generateToken("1", NUMBER))
	trueLine = append(trueLine, generateToken(";", SEMICOLON))
	falseLine = append(falseLine, generateToken("false", NAME))
	falseLine = append(falseLine, generateToken("=", OPERATOR_SINGLE_EQUALS))
	falseLine = append(falseLine, generateToken("0", NUMBER))
	falseLine = append(falseLine, generateToken(";", SEMICOLON))
	resultLines = append(resultLines, trueLine)
	resultLines = append(resultLines, falseLine)
	//add after the refaLines
	for i := firstLineAfterRefaIndex; i < len(inputLines); i++ {
		resultLines = append(resultLines, inputLines[i])
	}
	//flatten into result
	for _, line := range resultLines {
		for _, token := range line {
			result = append(result, token)
		}
	}
	return result
}
