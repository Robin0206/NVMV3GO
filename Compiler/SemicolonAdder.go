package Compiler

import (
	"fmt"
	"strings"
)

type SemicolonAdder struct{}

func (this *SemicolonAdder) process(line string) string {
	if this.shouldHaveSemicolon(line) {
		return line + ";"
	} else {
		return line
	}
}

func (this *SemicolonAdder) shouldHaveSemicolon(line string) bool {
	keyWords := []string{
		"{",
		"}",
		"func ",
		"int ",
		"bool ",
		"byte ",
		"real ",
	}
	if strings.Contains(line, "print") {
		fmt.Println("=======")
	}
	for _, symbol := range keyWords {
		if strings.Contains(line, symbol) {
			return false
		}
	}
	return true
}
