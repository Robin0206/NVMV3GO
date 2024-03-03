package Executor

type NVMVariable struct {
	boolValue    []bool
	byteValue    []uint8
	integerValue []int32
	realValue    []float64
	valueType    uint8
	isArray      bool
	isResizable  bool
}

func generateNVMVariable(valueType, length int) NVMVariable {
	switch valueType {
	case 0:
		return generateBoolVariable(length + 1)
	case 1:
		return generateByteVariable(length + 1)
	case 2:
		return generateIntegerVariable(length + 1)
	case 3:
		return generateRealVariable(length + 1)
	default:
		return generateIntegerVariable(length + 1)
	}
}

func generateRealVariable(length int) NVMVariable {
	var result NVMVariable
	result.valueType = 3
	if length == 0 {
		result.isArray = true
		result.isResizable = true
		return result
	}
	if length == 1 {
		result.realValue = append(result.realValue, 0)
		result.isArray = false
	} else {
		result.realValue = make([]float64, length)
		result.isArray = true
		result.isResizable = false
	}
	return result
}

func generateIntegerVariable(length int) NVMVariable {
	var result NVMVariable
	result.valueType = 2
	if length == 0 {
		result.isArray = true
		result.isResizable = true
		return result
	}
	if length == 1 {
		result.integerValue = append(result.integerValue, 0)
		result.isArray = false
	} else {
		result.integerValue = make([]int32, length)
		result.isArray = true
	}
	return result
}

func generateByteVariable(length int) NVMVariable {
	var result NVMVariable
	result.valueType = 1
	if length == 0 {
		result.isArray = true
		result.isResizable = true
		return result
	}
	if length == 1 {
		result.byteValue = append(result.byteValue, 0)
		result.isArray = false
	} else {
		result.byteValue = make([]uint8, length)
		result.isArray = true
	}
	return result
}

func generateBoolVariable(length int) NVMVariable {
	var result NVMVariable
	result.valueType = 0
	if length == 0 {
		result.isArray = true
		result.isResizable = true
		return result
	}
	if length == 1 {
		result.boolValue = append(result.boolValue, false)
		result.isArray = false
	} else {
		result.boolValue = make([]bool, length)
		result.isArray = true
	}
	return result
}
