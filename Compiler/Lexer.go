package Compiler

import (
	"strconv"
	"strings"
)

type Lexer struct {
}

func generateLexer() Lexer {
	var result Lexer
	return result
}

func (this *Lexer) tokenize(line string) []Token {
	var splittedLine = strings.Split(line, " ")
	var result []Token
	for _, stringToken := range splittedLine {
		var tokenType = this.determineType(stringToken)
		if tokenType != -1 {
			result = append(result, generateToken(stringToken, tokenType))
		}
	}
	return result
}

func (this *Lexer) determineType(token string) int {
	if token == " " || token == "" || token == "  " {
		return -1
	}
	if _, err := strconv.Atoi(token); err == nil {
		return 23
	}
	if _, err := strconv.ParseFloat(token, 64); err == nil {
		return 23
	}
	if strings.Contains(token, "b_") {
		return 23
	}
	var systemFunctionsUpperCase = []string{
		"NOOP",
		"REFA",
		"MOV",
		"SET",
		"SETV",
		"ASET",
		"AGET",
		"CPY",
		"ADD",
		"SUB",
		"MUL",
		"DIV",
		"BINOR",
		"BINAND",
		"BINNOT",
		"BINXOR",
		"LESSTHAN",
		"GREATERTHAN",
		"LOGOR",
		"LOGAND",
		"LOGNOT",
		"LOGEQ",
		"PRINT",
		"RETURN",
		"CALL",
		"PARG",
		"BEQ",
		"LABEL",
		"FUNC",
		"FEND",
		"GARG",
		"MAIN",
		"MEND",
		"RETG",
		"MOD",
	}

	for _, sysFunc := range systemFunctionsUpperCase {
		if sysFunc == strings.ToUpper(token) {
			return SYSTEM_FUNCTION
		}
	}
	var keyWords = []string{
		"for",
		"if",
		"while",
		"func",
		"else",
		"with",
		"return",
		"from",
		"to",
		"for",
		"do",
		"bool",
		"int",
		"byte",
		"real",
		"times",
	}
	for _, keyWord := range keyWords {
		if keyWord == strings.ToLower(token) {
			return KEYWORD
		}
		if strings.Contains(strings.ToLower(token), keyWord) && strings.Contains(token, ".") {
			return KEYWORD
		}
	}

	switch token {
	case CURLY_BRACE_LEFT_STRING:
		return 0
	case CURLY_BRACE_RIGHT_STRING:
		return 1
	case BRACE_RIGHT_STRING:
		return 2
	case BRACE_LEFT_STRING:
		return 3

	case OPERATOR_PLUS_STRING:
		return 4
	case OPERATOR_MINUS_STRING:
		return 5
	case OPERATOR_DIVIDE_STRING:
		return 6
	case OPERATOR_MULTIPLY_STRING:
		return 7
	case OPERATOR_XOR_STRING:
		return 9
	case OPERATOR_MODULO_STRING:
		return 10

	case OPERATOR_SINGLE_EQUALS_STRING:
		return 11
	case OPERATOR_SINGLE_OR_STRING:
		return 12
	case OPERATOR_SINGLE_AND_STRING:
		return 13

	case OPERATOR_DOUBLE_EQUALS_STRING:
		return 14
	case OPERATOR_DOUBLE_OR_STRING:
		return 15
	case OPERATOR_DOUBLE_AND_STRING:
		return 16

	case OPERATOR_LESS_STRING:
		return 17
	case OPERATOR_MORE_STRING:
		return 18

	case SEMICOLON_STRING:
		return 19
	case COMMA_STRING:
		return 20

	case ARR_BRACE_RIGHT_STRING:
		return 24
	case ARR_BRACE_LEFT_STRING:
		return 25
	default:
		return 21 // NAME
	}
}
