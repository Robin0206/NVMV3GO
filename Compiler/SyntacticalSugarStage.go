package Compiler

import (
	"fmt"
	"strconv"
)

type SyntacticalSugarStage interface {
	processTokens(tokens []Token) []Token
}

func splitToLines(tokens []Token) [][]Token {
	var result [][]Token
	var current []Token
	for _, token := range tokens {
		current = append(current, token)
		if token.tokenType == SEMICOLON || token.tokenType == CURLY_BRACE_LEFT || token.tokenType == CURLY_BRACE_RIGHT {
			result = append(result, current)
			current = nil
		}
	}
	return result
}

func generateRefaLine(allocate string, varType int) []Token {
	var result []Token
	result = append(result, generateToken("REFA", SYSTEM_FUNCTION))
	result = append(result, generateToken("(", BRACE_LEFT))
	result = append(result, generateToken(allocate, NAME))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken(strconv.Itoa(varType), NUMBER))
	result = append(result, generateToken(",", COMMA))
	result = append(result, generateToken("0", NUMBER))
	result = append(result, generateToken(")", BRACE_RIGHT))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func generateSetLine(name, value string) []Token {
	var result []Token
	result = append(result, generateToken(name, NAME))
	result = append(result, generateToken("=", OPERATOR_SINGLE_EQUALS))
	result = append(result, generateToken(value, NUMBER))
	result = append(result, generateToken(";", SEMICOLON))
	return result
}

func printTokens(tokens []Token) {
	fmt.Println("==================================================")
	var tabLevel = 0
	for _, token := range tokens {

		if token.tokenType == CURLY_BRACE_RIGHT {
			fmt.Println()
			tabLevel--
			for i := 0; i < tabLevel; i++ {
				fmt.Print("\t")
			}
		}
		fmt.Print(token.content + " ")
		if token.tokenType == CURLY_BRACE_LEFT {
			fmt.Println()
			tabLevel++
			for i := 0; i < tabLevel; i++ {
				fmt.Print("\t")
			}
		}
		if token.tokenType == SEMICOLON {
			fmt.Println()
			for i := 0; i < tabLevel; i++ {
				fmt.Print("\t")
			}
		}
	}
	fmt.Println()
}

func getBlockAfter(coordinate int, function []Token) (int, int) {
	var start = -1 // the index after the first {
	var end = -1   //the index of the ending }
	var tabLevel = 0
	for i := coordinate; i < len(function); i++ {
		if coordinate == -1 {
			break
		}
		if start == -1 && function[i].tokenType == CURLY_BRACE_LEFT {
			start = i + 1
			tabLevel++
			continue
		}
		if function[i].tokenType == CURLY_BRACE_LEFT {
			tabLevel++
			continue
		}
		if function[i].tokenType == CURLY_BRACE_RIGHT {
			tabLevel--
			if tabLevel == 0 {
				end = i
				break
			}
			continue
		}
	}
	return start, end
}

func flatten(lines [][]Token) []Token {
	var result []Token
	for _, line := range lines {
		for _, token := range line {
			result = append(result, token)
		}
	}
	return result
}

func substituteNameWithNumber(lines [][]Token, content string, counter int) [][]Token {
	var inputFlattened = flatten(lines)
	var resultFlattened []Token
	for _, token := range inputFlattened {
		if token.content == content {
			token.content = strconv.Itoa(counter)
			token.tokenType = NUMBER
		}
		resultFlattened = append(resultFlattened, token)
	}
	return splitToLines(resultFlattened)
}

func lineContainsToken(content string, line []Token) bool {
	for _, token := range line {
		if token.content == content {
			return true
		}
	}
	return false
}

func lineContainsTokenType(t int, line []Token) bool {
	for _, token := range line {
		if token.tokenType == t {
			return true
		}
	}
	return false
}
