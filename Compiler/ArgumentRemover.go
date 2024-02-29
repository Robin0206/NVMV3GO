package Compiler

import (
	"strconv"
	"strings"
)

type ArgumentRemover struct {
}

func (this *ArgumentRemover) processTokens(input []Token) []Token {
	var inputLines = splitToLines(input)
	var firstLine = inputLines[0]
	if firstLine[2].tokenType != BRACE_LEFT {
		return input
	}
	var refaLines [][]Token
	var newRefaLines [][]Token
	var gargLines [][]Token
	var rest [][]Token
	var inRefas = true
	var resultLines [][]Token
	//fill the arrays
	for i := 1; i < len(inputLines); i++ {
		if inputLines[i][0].content != "REFA" {
			inRefas = false
		}
		if inRefas {
			refaLines = append(refaLines, inputLines[i])
		}
		if !inRefas {
			rest = append(rest, inputLines[i])
		}
	}
	var arguments = extractArguments(firstLine)
	for index, arg := range arguments {
		argType, _ := strconv.ParseInt(arg[0].content, 10, 64)
		newRefaLines = append(newRefaLines, generateRefaLine(arg[1].content, int(argType)))
		var gargLine []Token
		gargLine = append(gargLine, generateToken("GARG", SYSTEM_FUNCTION))
		gargLine = append(gargLine, generateToken("(", BRACE_LEFT))
		gargLine = append(gargLine, arg[1])
		gargLine = append(gargLine, generateToken(",", COMMA))
		gargLine = append(gargLine, generateToken(strconv.Itoa(index), NUMBER))
		gargLine = append(gargLine, generateToken(")", BRACE_RIGHT))
		gargLine = append(gargLine, generateToken(";", SEMICOLON))
		gargLines = append(gargLines, gargLine)
	}
	firstLine = nil
	firstLine = append(firstLine, input[0])
	firstLine = append(firstLine, input[1])
	firstLine = append(firstLine, generateToken("{", CURLY_BRACE_LEFT))
	//construct the result
	resultLines = append(resultLines, firstLine)
	for _, token := range refaLines {
		resultLines = append(resultLines, token)
	}
	for _, token := range newRefaLines {
		resultLines = append(resultLines, token)
	}
	for _, token := range gargLines {
		resultLines = append(resultLines, token)
	}
	for _, token := range rest {
		resultLines = append(resultLines, token)
	}
	return flatten(resultLines)
}

func extractArguments(line []Token) [][]Token {
	var argumentsString string
	var write = false
	for _, token := range line {
		if token.tokenType == BRACE_RIGHT {
			break
		}
		if write {
			argumentsString += token.content + " "
		}
		if token.tokenType == BRACE_LEFT {
			write = true
		}
	}
	var splittedAtCommaArguments = strings.Split(argumentsString, ",")
	var resultStrings [][]string
	for _, str := range splittedAtCommaArguments {
		resultStrings = append(resultStrings, strings.Split(str, " "))
	}
	var result [][]Token
	for _, rawPair := range resultStrings {
		var pair []string
		for _, str := range rawPair {
			if str != "" {
				pair = append(pair, str)
			}
		}
		var argType Token
		switch pair[0] {
		case "bool":
			argType = generateToken("0", NUMBER)
			break
		case "byte":
			argType = generateToken("1", NUMBER)
			break
		case "int":
			argType = generateToken("2", NUMBER)
			break
		case "real":
			argType = generateToken("3", NUMBER)
			break
		default:
			argType = generateToken("3", NUMBER)
			continue
		}
		var argName = generateToken(pair[1], NAME)
		result = append(result, []Token{argType, argName})
	}
	return result
}
