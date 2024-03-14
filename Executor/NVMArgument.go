package Executor

import (
	"strconv"
	"strings"
)

type NVMArgument struct {
	integer   int
	realValue float64
	valueType uint8 // 0 = int 1 = float
}

func generateNVMArgument(value string) NVMArgument {
	if strings.Contains(value, "b_") {
		value = strings.ReplaceAll(value, "b_", "")
		parsedValue, _ := strconv.ParseInt(value, 10, 8)
		return NVMArgument{realValue: 0.0, integer: int(parsedValue), valueType: 0}
	}
	if strings.Contains(value, ".") {
		parsedValue, _ := strconv.ParseFloat(value, 64)
		return NVMArgument{realValue: parsedValue, integer: 0, valueType: 1}
	} else {
		parsedValue, _ := strconv.ParseInt(value, 10, 32)
		return NVMArgument{realValue: 0.0, integer: int(parsedValue), valueType: 0}
	}
}
