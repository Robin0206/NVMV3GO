package Compiler

import "strconv"

type IfConverter struct {
	counter int
}

func generateIfConverter() *IfConverter {
	var result IfConverter
	result.counter = 0
	return &result
}

func (this *IfConverter) processTokens(input []Token) []Token {
	var result []Token
	for _, token := range input {
		result = append(result, token)
	}
	var ifCoordinates = findIfToken(result)
	for ifCoordinates != -1 {
		result = this.substituteIfAt(result, ifCoordinates)
		ifCoordinates = findIfToken(result)
	}

	return result
}

func (this *IfConverter) substituteIfAt(wholeFunction []Token, coordinate int) []Token {
	var result []Token
	var startOfIf = coordinate
	startOfBlock, endOfBlock := getBlockAfter(coordinate, wholeFunction)
	var beforeIf []Token
	var expression []Token
	var body []Token
	var afterIf []Token
	//fillBeforeIf
	for i := 0; i < startOfIf; i++ {
		beforeIf = append(beforeIf, wholeFunction[i])
	}
	//fill expression
	for i := startOfIf + 1; i < len(wholeFunction); i++ {
		if wholeFunction[i].tokenType == CURLY_BRACE_LEFT {
			break
		}
		expression = append(expression, wholeFunction[i])
	}
	//fill Body
	for i := startOfBlock; i < endOfBlock; i++ {
		body = append(body, wholeFunction[i])
	}
	//fill after Body
	for i := endOfBlock + 1; i < len(wholeFunction); i++ {
		afterIf = append(afterIf, wholeFunction[i])
	}
	//convert
	var convertedIfLoop = this.convertIfStatement(expression, body)
	//add before if to result
	for _, token := range beforeIf {
		result = append(result, token)
	}
	//add converted if to result
	for _, token := range convertedIfLoop {
		result = append(result, token)
	}
	//add after if to result
	for _, token := range afterIf {
		result = append(result, token)
	}
	return result
}

func findIfToken(input []Token) int {
	for i := 0; i < len(input); i++ {
		if input[i].content == "if" {
			return i
		}
	}
	return -1
}

func (this *IfConverter) convertIfStatement(expression []Token, body []Token) []Token {
	var result []Token
	var bufferName = "____if_buffer_" + strconv.Itoa(this.counter)
	var bodyLabelName = "____if_body_label_" + strconv.Itoa(this.counter)
	var afterBodyLabel = "____after_if_body_label_" + strconv.Itoa(this.counter)

	//add the refaline for the buffer
	var refaCall = generateRefaLine(bufferName, BOOL)
	for _, token := range refaCall {
		result = append(result, token)
	}

	//add the setline for the buffer
	result = append(result, generateToken(bufferName, NAME))
	result = append(result, generateToken("=", OPERATOR_SINGLE_EQUALS))
	for _, token := range expression {
		result = append(result, token)
	}
	result = append(result, generateToken(";", SEMICOLON))

	//add the beq to the body
	result = append(result, generateToken("BEQ", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(bufferName, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken("true", NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(bodyLabelName, NAME))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))

	//add the beq after the body
	result = append(result, generateToken("BEQ", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken("true", NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken("true", NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(afterBodyLabel, NAME))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))

	//add the body label
	result = append(result, generateToken("LABEL", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(bodyLabelName, NAME))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))

	//add the body
	for _, token := range body {
		result = append(result, token)
	}
	//add the label after the body
	result = append(result, generateToken("LABEL", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(afterBodyLabel, NAME))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))
	//increment counter
	this.counter++

	return result
}
