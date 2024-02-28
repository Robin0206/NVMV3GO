package Compiler

type LabelSubstitutor struct {
	counter int // needs to be reset for every precessed function!!!!
}

func (this *LabelSubstitutor) processTokens(input []Token) []Token {
	this.counter = 0
	var inputLines = splitToLines(input)
	//read all labelNames
	for i := 0; i < len(inputLines); i++ {
		if inputLines[i][0].content == "LABEL" {
			inputLines = substituteNameWithNumber(inputLines, inputLines[i][1].content, this.counter)
			this.counter++
		}
	}
	return flatten(inputLines)
}
