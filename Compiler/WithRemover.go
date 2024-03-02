package Compiler

type WithRemover struct {
}

func (this *WithRemover) processTokens(input []Token) []Token {
	if notContainsWith(input) {
		return input
	}
	var resultLines [][]Token
	var withBlock = extractWithBlock(input)
	var refaLines = withBlockToRefaLines(withBlock)
	var inputLines = splitToLines(input)
	inputLines[0] = removeWithBlock(inputLines[0])
	resultLines = append(resultLines, inputLines[0])
	for i := 0; i < len(refaLines); i++ {
		resultLines = append(resultLines, refaLines[i])
	}
	for i := 1; i < len(inputLines); i++ {
		resultLines = append(resultLines, inputLines[i])
	}
	return flatten(resultLines)
}

func notContainsWith(input []Token) bool {
	for _, token := range input {
		if token.content == "with" {
			return false
		}
	}
	return true
}

func removeWithBlock(tokens []Token) []Token {
	var result []Token
	for i := 0; i < len(tokens); i++ {
		if tokens[i].content == "with" {
			break
		}
		result = append(result, tokens[i])
	}
	result = append(result, generateToken("{", CURLY_BRACE_LEFT))
	return result
}

func withBlockToRefaLines(withBlockContents []Token) [][]Token {
	var result [][]Token
	var withLines [][]Token
	var currentLine []Token
	for i := 0; i < len(withBlockContents); i++ {
		if i != len(withBlockContents)-1 && withBlockContents[i].tokenType == BRACE_RIGHT && withBlockContents[i+1].tokenType == COMMA {
			withLines = append(withLines, currentLine)
			currentLine = nil
		} else {
			if withBlockContents[i].tokenType == NAME || withBlockContents[i].tokenType == KEYWORD {
				currentLine = append(currentLine, withBlockContents[i])
			}
		}
	}
	withLines = append(withLines, currentLine)
	for _, line := range withLines {
		if len(line) > 0 {
			var t = line[0].content
			for i := 1; i < len(line); i++ {
				var argPair = generateArgumentPair(t + " " + line[i].content)
				result = append(result, argPair.generateRefaLine())
			}
		}
	}
	return result
}

func extractWithBlock(input []Token) []Token {
	var withBlock []Token
	var write = false
	for _, token := range input {
		if token.content == "with" {
			write = true
		}
		if write {
			if token.tokenType == CURLY_BRACE_LEFT {
				break
			}
			withBlock = append(withBlock, token)
		}
	}
	var result []Token
	for i := 2; i < len(withBlock)-1; i++ {
		result = append(result, withBlock[i])
	}
	return result
}
