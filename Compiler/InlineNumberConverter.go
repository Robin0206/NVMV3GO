package Compiler

import (
	"strings"
)

type InlineNumberConverter struct {
	alreadyConverted []string
}

func generateInlineNumberConverter() *InlineNumberConverter {
	var result InlineNumberConverter
	return &result
}

func (this *InlineNumberConverter) processTokens(tokens []Token) []Token {
	var result []Token
	var inputLines = splitToLines(tokens)
	var numberLineIndex = containsInlineNumber(inputLines)
	for numberLineIndex != -1 {
		inputLines = this.substituteNumberLine(inputLines, numberLineIndex)
		numberLineIndex = containsInlineNumber(inputLines)
	}
	for _, line := range inputLines {
		for _, token := range line {
			result = append(result, token)
		}
	}
	return result
}

// TODO
func (this *InlineNumberConverter) substituteNumberLine(lines [][]Token, numberLineIndex int) [][]Token {
	var resultLines [][]Token
	for i := 0; i < numberLineIndex; i++ {
		resultLines = append(resultLines, lines[i])
	}
	var converted = this.convertNumberLine(lines[numberLineIndex])
	for _, line := range converted {
		resultLines = append(resultLines, line)
	}
	for i := numberLineIndex + 1; i < len(lines); i++ {
		resultLines = append(resultLines, lines[i])
	}
	return resultLines
}

func (this *InlineNumberConverter) convertNumberLine(line []Token) [][]Token {
	var result [][]Token
	var resultExpression []Token
	var allocate []string
	for _, token := range line {
		if token.tokenType != NUMBER {
			resultExpression = append(resultExpression, token)
		} else {
			resultExpression = append(resultExpression, generateToken("____numBuffer_"+token.content, NAME))
			if !this.alreadyAllocated(token.content) {
				allocate = append(allocate, token.content)
			}
		}
	}
	//generate refas
	var refaLines [][]Token
	var setLines [][]Token
	for _, numberToAllocate := range allocate {
		if strings.Contains(numberToAllocate, ".") {
			refaLines = append(refaLines, generateRefaLine("____numBuffer_"+numberToAllocate, REAL))
		} else {
			refaLines = append(refaLines, generateRefaLine("____numBuffer_"+numberToAllocate, INT))
		}
	}
	//generate sets
	for _, numberToAllocate := range allocate {
		if strings.Contains(numberToAllocate, ".") {
			refaLines = append(refaLines, generateSetLine("____numBuffer_"+numberToAllocate, numberToAllocate))
		} else {
			refaLines = append(refaLines, generateSetLine("____numBuffer_"+numberToAllocate, numberToAllocate))
		}
	}
	//add them all
	for _, refaLine := range refaLines {
		result = append(result, refaLine)
	}
	for _, setLine := range setLines {
		result = append(result, setLine)
	}
	result = append(result, resultExpression)
	return result
}

func (this *InlineNumberConverter) alreadyAllocated(content string) bool {
	for _, tokenContent := range this.alreadyConverted {
		if tokenContent == content {
			return true
		}
	}
	return false
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
