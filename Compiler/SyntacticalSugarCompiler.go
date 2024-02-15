package Compiler

import "NVMV3/Executor"

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
	var functions = splitToFunctions(result)
	for i := 0; i < len(functions); i++ {
		for _, stage := range this.syntacticalSugarProcessingChain {
			functions[i] = stage.processTokens(functions[i])
		}
	}
	result = nil
	for _, function := range functions {
		for _, token := range function {
			result = append(result, token)
		}
	}
	var resultCommands = tokenDoubleArrToCommandArr(splitToLines(result))
	return resultCommands
}

func splitToFunctions(tokens []Token) [][]Token {
	var result [][]Token
	result = append(result, []Token{})
	var lines = splitToLines(tokens)
	for _, line := range lines {
		for _, token := range line {
			result[len(result)-1] = append(result[len(result)-1], token)
		}
		if line[0].content == "FEND" || line[0].content == "MEND" {
			result = append(result, []Token{})
		}
	}
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
