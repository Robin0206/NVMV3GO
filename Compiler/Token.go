package Compiler

import "strings"

const CURLY_BRACE_LEFT_STRING = "{"
const CURLY_BRACE_RIGHT_STRING = "}"
const BRACE_RIGHT_STRING = ")"
const BRACE_LEFT_STRING = "("

const OPERATOR_PLUS_STRING = "+"
const OPERATOR_MINUS_STRING = "-"
const OPERATOR_DIVIDE_STRING = "/"
const OPERATOR_MULTIPLY_STRING = "*"
const OPERATOR_XOR_STRING = "^"
const OPERATOR_MODULO_STRING = "%"

const OPERATOR_SINGLE_EQUALS_STRING = "="
const OPERATOR_SINGLE_OR_STRING = "|"
const OPERATOR_SINGLE_AND_STRING = "&"

const OPERATOR_DOUBLE_EQUALS_STRING = "=="
const OPERATOR_DOUBLE_OR_STRING = "||"
const OPERATOR_DOUBLE_AND_STRING = "&&"

const OPERATOR_LESS_STRING = "<"
const OPERATOR_MORE_STRING = ">"

const SEMICOLON_STRING = ";"
const COMMA_STRING = ","

const ARR_BRACE_RIGHT_STRING = "]"
const ARR_BRACE_LEFT_STRING = "["

const CURLY_BRACE_LEFT = 0
const CURLY_BRACE_RIGHT = 1
const BRACE_RIGHT = 2
const BRACE_LEFT = 3

const OPERATOR_PLUS = 4
const OPERATOR_MINUS = 5
const OPERATOR_DIVIDE = 6
const OPERATOR_MULTIPLY = 7
const OPERATOR_NOT = 8
const OPERATOR_XOR = 9
const OPERATOR_MODULO = 10

const OPERATOR_SINGLE_EQUALS = 11
const OPERATOR_SINGLE_OR = 12
const OPERATOR_SINGLE_AND = 13

const OPERATOR_DOUBLE_EQUALS = 14
const OPERATOR_DOUBLE_OR = 15
const OPERATOR_DOUBLE_AND = 16

const OPERATOR_LESS = 17
const OPERATOR_MORE = 18

const SEMICOLON = 19
const COMMA = 20

const NAME = 21
const KEYWORD = 22
const NUMBER = 23
const ARR_BRACE_RIGHT = 24
const ARR_BRACE_LEFT = 25
const SYSTEM_FUNCTION = 26

type Token struct {
	content   string
	tokenType int
}

func generateToken(content string, tokenType int) Token {
	var result Token
	result.tokenType = tokenType
	result.content = strings.Replace(content, " ", "", -1)
	return result
}
