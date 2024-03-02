package Compiler

import (
	"strconv"
)

type ExpressionSimplifier struct {
	newAllocations      []Token
	newAllocationsTypes []int
	counter             int
}

func generateExpressionSimplifier() *ExpressionSimplifier {
	var result ExpressionSimplifier
	result.counter = 0
	return &result
}

func (this *ExpressionSimplifier) processTokens(input []Token) []Token {
	//printTokens(input)
	var inputLines = splitToLines(input)
	var expressionLineIndex = getExpressionLineIndex(inputLines)
	var resultBuffer [][]Token
	resultBuffer = inputLines
	for expressionLineIndex != -1 {
		//printTokens(flatten(inputLines))
		resultBuffer = nil
		var convertedExpression = this.convertExpression(inputLines[expressionLineIndex], inputLines)
		//until the last old refaLine
		for index, line := range inputLines {
			if index > 0 && line[0].content != "REFA" {
				break
			} else {
				resultBuffer = append(resultBuffer, line)
			}
		}
		//add the new refaLines
		for index, varNameToken := range this.newAllocations {
			resultBuffer = append(resultBuffer, generateRefaLine(varNameToken.content, this.newAllocationsTypes[index]))
		}
		//add the content after the last refa until the new expression
		var write = false
		for index, line := range inputLines {
			if index > 0 && line[0].content != "REFA" {
				write = true
			}
			if index == expressionLineIndex {
				break
			}
			if write {
				resultBuffer = append(resultBuffer, line)
			}

		}
		//add the converted expression lines
		for _, line := range convertedExpression {
			resultBuffer = append(resultBuffer, line)
		}
		//add after the expression old expression
		for i := expressionLineIndex + 1; i < len(inputLines); i++ {
			resultBuffer = append(resultBuffer, inputLines[i])
		}

		//copy reset and set new expressionLineIndex
		inputLines = nil
		for i := 0; i < len(resultBuffer); i++ {
			inputLines = append(inputLines, resultBuffer[i])
		}
		expressionLineIndex = getExpressionLineIndex(resultBuffer)
		this.newAllocations = nil
		this.newAllocationsTypes = nil
	}
	var result []Token
	for _, line := range inputLines {
		for _, token := range line {
			result = append(result, token)
		}
	}
	return result
}

// converts the lines
// sets the new refas
// sets the new refa type
func (this *ExpressionSimplifier) convertExpression(line []Token, wholeFunction [][]Token) [][]Token {
	return this.simplify(line[0], extractExpression(line), wholeFunction)
}

func extractExpression(line []Token) []Token {
	var result []Token
	for i := 2; i < len(line)-1; i++ {
		result = append(result, line[i])
	}
	return result
}

func (this *ExpressionSimplifier) simplify(varName Token, expression []Token, wholeFunction [][]Token) [][]Token {
	var result []Token
	var reversePolish = this.shuntingYard(expression)
	var stack []Token
	var expressionConversionResult []Token

	for _, token := range reversePolish {
		if token.isOperator() {
			var right = stack[len(stack)-1]
			stack = stack[:(len(stack) - 1)]
			var left = stack[len(stack)-1]
			stack = stack[:(len(stack) - 1)]
			expressionConversionResult = this.constructExpression(left, right, token, wholeFunction)
			for _, expressionToken := range expressionConversionResult {
				result = append(result, expressionToken)
			}
			stack = append(stack, this.newAllocations[len(this.newAllocations)-1])
		} else {
			stack = append(stack, token)
		}
	}
	result = this.substituteLastBufferWithVarName(result, varName)
	return splitToLines(result)
}

func (this *ExpressionSimplifier) constructExpression(left, right, operator Token, wholeFunction [][]Token) []Token {
	var result []Token
	var buffer = generateToken("____expressionBuffer_"+strconv.Itoa(this.counter), NAME)
	result = append(result, buffer)
	result = append(result, generateToken("=", OPERATOR_SINGLE_EQUALS))
	result = append(result, left)
	result = append(result, operator)
	result = append(result, right)
	result = append(result, generateToken(";", SEMICOLON))
	var allocationType int
	if isOperatorThatAlwaysOutPutsBool(operator) {
		allocationType = 0
	} else {
		allocationType = this.getType(left, wholeFunction)
	}
	this.newAllocations = append(this.newAllocations, buffer)
	this.newAllocationsTypes = append(this.newAllocationsTypes, allocationType)
	this.counter++
	return result
}

func isOperatorThatAlwaysOutPutsBool(operator Token) bool {
	return operator.tokenType == OPERATOR_DOUBLE_EQUALS || operator.tokenType == OPERATOR_LESS || operator.tokenType == OPERATOR_MORE
}

func (this *ExpressionSimplifier) getType(varName Token, wholeFunction [][]Token) int {
	for _, line := range wholeFunction {
		if line[0].content == "REFA" && line[2].content == varName.content {
			result, _ := strconv.ParseInt(line[4].content, 10, 64)
			return int(result)
		}
	}
	for index, refaName := range this.newAllocations {
		if varName.content == refaName.content {
			return this.newAllocationsTypes[index]
		}
	}
	return -1
}

func getExpressionLineIndex(lines [][]Token) int {

	for index, line := range lines {
		if lineIsComplexExpression(line) {
			return index
		}
	}
	return -1
}

func lineIsComplexExpression(line []Token) bool {
	var numOperators = 0
	for _, token := range line {
		if token.isOperatorOtherThanSingleEquals() {
			numOperators++
		}
	}
	return numOperators >= 2
}

func operatorPrescedence(operator int) int {
	switch operator {
	case OPERATOR_PLUS:
		return 11
	case OPERATOR_MINUS:
		return 11
	case OPERATOR_DIVIDE:
		return 12
	case OPERATOR_MULTIPLY:
		return 12
	case OPERATOR_NOT:
		return 13
	case OPERATOR_XOR:
		return 6
	case OPERATOR_MODULO:
		return 12
	case OPERATOR_SINGLE_OR:
		return 5
	case OPERATOR_SINGLE_AND:
		return 7
	case OPERATOR_DOUBLE_EQUALS:
		return 8
	case OPERATOR_DOUBLE_OR:
		return 3
	case OPERATOR_DOUBLE_AND:
		return 4
	case OPERATOR_LESS:
		return 9
	case OPERATOR_MORE:
		return 9
	default:
		return -1
	}
}

func (this *ExpressionSimplifier) shuntingYard(expression []Token) []Token {
	var result []Token
	var stack []Token
	for _, token := range expression {
		if token.tokenType == NUMBER || token.tokenType == NAME {
			result = append(result, token)
		}
		if token.isOperator() {
			for len(stack) != 0 &&
				stack[len(stack)-1].isOperator() &&
				operatorPrescedence(token.tokenType) <= operatorPrescedence(stack[len(stack)-1].tokenType) {

				result = append(result, stack[len(stack)-1])
				stack = stack[:(len(stack) - 1)]
			}
			stack = append(stack, token)
		}
		if token.tokenType == BRACE_LEFT {
			stack = append(stack, token)
		}
		if token.tokenType == BRACE_RIGHT {
			for len(stack) != 0 && stack[len(stack)-1].tokenType != BRACE_LEFT {
				result = append(result, stack[len(stack)-1])
				stack = stack[:(len(stack) - 1)]
			}
			if len(stack) != 0 && stack[len(stack)-1].tokenType != BRACE_LEFT {
				stack = stack[:(len(stack) - 1)]
			}
		}
	}
	for len(stack) != 0 {
		if stack[len(stack)-1].tokenType != BRACE_RIGHT && stack[len(stack)-1].tokenType != BRACE_LEFT {
			result = append(result, stack[len(stack)-1])
			stack = stack[:(len(stack) - 1)]
		} else {
			stack = stack[:(len(stack) - 1)]
		}
	}
	return result
}

func (this *ExpressionSimplifier) substituteLastBufferWithVarName(input []Token, name Token) []Token {
	var result []Token
	var sawSingleEquals = false
	for i := len(input) - 1; i >= 0; i-- {
		if input[i].tokenType == OPERATOR_SINGLE_EQUALS {
			sawSingleEquals = true
		}
		if sawSingleEquals && input[i].content == this.newAllocations[len(this.newAllocations)-1].content {
			input[i].content = name.content
		}
	}
	result = input
	return result
}
