package Compiler

import (
	"NVMV3/Executor"
	"fmt"
	"strconv"
)

type SyntacticalSugarCompiler struct {
	preprocessor                    Preprocessor
	lexer                           Lexer
	syntacticalSugarProcessingChain []SyntacticalSugarStage
}

func GenerateSyntacticalSugarCompiler() SyntacticalSugarCompiler {
	var result SyntacticalSugarCompiler
	result.preprocessor = GeneratePreprocessor()
	result.lexer = generateLexer()
	result.syntacticalSugarProcessingChain = []SyntacticalSugarStage{
		generateExpressionSimplifier(),
		generateInlineNumberConverter(),
		&SetConverter{},
		&OperationConverter{},
		&RefaUpPuller{},
		&FunctionSignatureConverter{},
		&BracketAndCommaRemover{},
		generateVariableNameConverter(),
	}
	return result
}

func (this *SyntacticalSugarCompiler) Compile(input []string) []Executor.NVMCommand {
	var tempResult []string
	tempResult = this.preprocessor.ProcessLines(input)
	var result []Token
	for _, line := range tempResult {
		var tokens = this.lexer.tokenize(line)
		for _, token := range tokens {
			result = append(result, token)
		}
	}
	var functions = splitTokensToFunctions(result)
	for i := 0; i < len(functions); i++ {
		for _, stage := range this.syntacticalSugarProcessingChain {
			functions[i] = stage.processTokens(functions[i])
			printTokens(functions[i])
		}
	}
	var resultBuffer []Token
	for _, function := range functions {
		for _, token := range function {
			resultBuffer = append(resultBuffer, token)
		}
	}
	var resultLines = splitToLines(resultBuffer)
	var resultCommands = tokenDoubleArrToCommandArr(resultLines)
	return resultCommands
}

func splitTokensToFunctions(tokens []Token) [][]Token {
	var result [][]Token
	var current []Token
	var lines = splitToLines(tokens)
	for i, line := range lines {
		if i > 0 && len(line) > 0 && (line[0].content == "func") {
			result = append(result, current)
			current = nil
		}
		for _, token := range line {
			current = append(current, token)
		}
	}
	result = append(result, current)
	current = nil
	return result
}

func tokenDoubleArrToCommandArr(input [][]Token) []Executor.NVMCommand {
	var result []Executor.NVMCommand
	for _, line := range input {
		result = append(result, tokenLineToCommand(line))
	}
	return result
}

func tokenLineToCommand(tokenLine []Token) Executor.NVMCommand {
	var resultLine = ""
	for _, token := range tokenLine {
		if token.tokenType != SEMICOLON {
			resultLine = resultLine + token.content + " "
		}
	}
	return Executor.GenerateNVMCommand(resultLine)
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
