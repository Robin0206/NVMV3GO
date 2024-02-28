package Compiler

type ElseConverter struct {
	counter int
}

func generateElseConverter() *ElseConverter {
	var result ElseConverter
	result.counter = 0
	return &result
}

func (this *ElseConverter) processTokens(input []Token) []Token {
	var result []Token
	result = input
	var elseCoordinates, associatedIfCoordinates = findElseCoordinates(result)
	for elseCoordinates != -1 {
		result = this.convertElse(result, elseCoordinates, associatedIfCoordinates)
		elseCoordinates, associatedIfCoordinates = findElseCoordinates(result)
	}
	return result
}

func (this *ElseConverter) convertElse(input []Token, elseCoordinates int, ifCoordinates int) []Token {
	var result []Token
	var expression = this.extractExpression(input, ifCoordinates)
	//add until the else to result
	for i := 0; i < elseCoordinates; i++ {
		result = append(result, input[i])
	}
	//add the converted else
	var convertedElse = this.constructElse(expression)
	for _, token := range convertedElse {
		result = append(result, token)
	}
	//add after the else
	for i := elseCoordinates + 1; i < len(input); i++ {
		result = append(result, input[i])
	}
	return result
}

func (this *ElseConverter) constructElse(expression []Token) []Token {
	var result []Token
	//add if negatedExpression
	result = append(result, generateToken("if", NAME))
	for _, token := range expression {
		result = append(result, token)
	}
	result = append(result, generateToken("==", OPERATOR_DOUBLE_EQUALS))
	result = append(result, generateToken("false", NAME))
	return result
}

func (this *ElseConverter) extractExpression(input []Token, ifIndex int) []Token {
	var result []Token
	for i := ifIndex + 1; i < len(input); i++ {
		if input[i].tokenType == CURLY_BRACE_LEFT {
			break
		}
		result = append(result, input[i])
	}
	return result
}

func findElseCoordinates(result []Token) (int, int) {
	var ifCoordinate = -1
	var elseCoordinate = -1
	for index, token := range result {
		if token.content == "if" {
			_, elseCoordinate = getBlockAfter(index, result)
			ifCoordinate = index
			elseCoordinate = elseCoordinate + 1
			if len(result) > elseCoordinate && result[elseCoordinate].content == "else" {
				return elseCoordinate, ifCoordinate
			} else {
				ifCoordinate = -1
				elseCoordinate = -1
				continue
			}
		}
	}
	return elseCoordinate, ifCoordinate
}
