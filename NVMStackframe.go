package main

type NVMStackframe struct {
	function       *NVMFunction
	programCounter int
	executor       *NVMExecutor
	variables      []NVMVariable
}

func generateStackframe(function *NVMFunction, executor *NVMExecutor) NVMStackframe {
	var result NVMStackframe
	result.function = function
	result.programCounter = 0
	result.executor = executor
	return result
}

func (this *NVMStackframe) print() {
	this.function.print()
}
