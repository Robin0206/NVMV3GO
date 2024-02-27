package Compiler

import (
	"strconv"
)

type WhileConverter struct {
	counter int
}

func generateWhileConverter() *WhileConverter {
	var result WhileConverter
	result.counter = 0
	return &result
}

func (this *WhileConverter) processTokens(input []Token) []Token {
	var result []Token
	for _, token := range input {
		result = append(result, token)
	}
	var whileCoordinates = findWhileToken(result)
	if whileCoordinates != -1 {
		var loop = true
		for loop {
			result = this.substituteWhileAt(result, whileCoordinates)
			loop = false
		}
	}

	return result
}

func (this *WhileConverter) substituteWhileAt(wholeFunction []Token, coordinate int) []Token {
	var result []Token
	var startOfWhile = coordinate
	startOfBlock, endOfBlock := getBlockAfter(coordinate, wholeFunction)
	var beforeWhile []Token
	var expression []Token
	var body []Token
	var afterWhile []Token
	//fillBeforeWhile
	for i := 0; i < startOfWhile; i++ {
		beforeWhile = append(beforeWhile, wholeFunction[i])
	}
	//fill expression
	for i := startOfWhile + 1; i < len(wholeFunction); i++ {
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
		afterWhile = append(afterWhile, wholeFunction[i])
	}
	//convert
	var convertedWhileLoop = this.convertWhileLoop(expression, body)
	//add before while to result
	for _, token := range beforeWhile {
		result = append(result, token)
	}
	//add converted while to result
	for _, token := range convertedWhileLoop {
		result = append(result, token)
	}
	//add after while to result
	for _, token := range afterWhile {
		result = append(result, token)
	}
	return result
}

func findWhileToken(input []Token) int {
	for i := 0; i < len(input); i++ {
		if input[i].content == "while" {
			return i
		}
	}
	return -1
}

func (this *WhileConverter) convertWhileLoop(expression []Token, body []Token) []Token {
	var result []Token
	var bufferName = "____while_buffer_" + strconv.Itoa(this.counter)
	var bodyLabelName = "____while_body_label_" + strconv.Itoa(this.counter)
	var afterBodyLabel = "____after_while_body_label_" + strconv.Itoa(this.counter)

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

	// add the setline for the buffer
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
