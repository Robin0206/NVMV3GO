package Compiler

import (
	"strings"
)

type WhitespaceAdder struct{}

func (this *WhitespaceAdder) process(line string) string {
	var result string
	result = ""
	var symbols = "{}();%+-/*=<>|&!^,[]"
	for i := 0; i < len(line); i++ {
		if i != len(line)-1 {
			if strings.Contains(symbols, string(line[i])) &&
				line[i] == line[i+1] &&
				!(line[i] == '(' || line[i] == ')' || line[i] == '{' || line[i] == '}') {
				var toAppend = " " + string(line[i]) + string(line[i+1]) + " "
				result = result + toAppend
			} else if strings.Contains(symbols, string(line[i])) {
				var toAppend = " " + string(line[i]) + " "
				result = result + toAppend
			} else {
				result = result + string(line[i])
			}
		} else {
			if strings.Contains(symbols, string(line[i])) {
				var toAppend = " " + string(line[i]) + " "
				result = result + toAppend
			} else {
				result = result + string(line[i])
			}
		}
	}
	return result
}
