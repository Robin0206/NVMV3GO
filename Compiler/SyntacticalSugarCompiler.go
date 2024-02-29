package Compiler

import (
	"NVMV3/Executor"
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
		&ArgumentRemover{},
		&TrueAndFalseConverter{},
		generateElseConverter(),
		generateIfConverter(),
		generateWhileConverter(),
		generateExpressionSimplifier(),
		generateInlineNumberConverter(),
		&SetConverter{},
		&OperationConverter{},
		&RefaUpPuller{},
		&MultipleRefaRemover{},
		&FunctionSignatureConverter{},
		&BracketAndCommaRemover{},
		generateVariableNameConverter(),
		&LabelSubstitutor{},
	}
	return result
}

func (this *SyntacticalSugarCompiler) Compile(input []string, debugPrint bool) []Executor.NVMCommand {
	var tempResult []string

	//remove blank lines
	for _, line := range input {
		if len(line) > 0 {
			tempResult = append(tempResult, line)
		}
	}

	//run preprocessor
	tempResult = this.preprocessor.ProcessLines(tempResult)

	//tokenize
	var result []Token
	for _, line := range tempResult {
		var tokens = this.lexer.tokenize(line)
		for _, token := range tokens {
			result = append(result, token)
		}
	}
	//run SyntacticalSugarChain
	var functions = splitTokensToFunctions(result)
	if debugPrint {
		printTokens(functions[0])
	}
	for i := 0; i < len(functions); i++ {
		for _, stage := range this.syntacticalSugarProcessingChain {
			functions[i] = stage.processTokens(functions[i])
			if debugPrint {
				printTokens(functions[i])
			}
		}
	}
	//flatten functions to 1d Token Array
	var resultBuffer []Token
	for _, function := range functions {
		for _, token := range function {
			resultBuffer = append(resultBuffer, token)
		}
	}
	//split them again to lines
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
