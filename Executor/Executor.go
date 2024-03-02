package Executor

import (
	"fmt"
	"time"
)

type NVMExecutor struct {
	stop             bool
	pop              bool
	functions        []NVMFunction
	functionPointers []*NVMFunction
	stack            []NVMStackframe
	delegateTable    []NVMDelegate
	argRegister      []NVMVariable
	returnRegister   NVMVariable
}

func GenerateExecutor(functions []NVMFunction) *NVMExecutor {
	var result NVMExecutor
	functions = SubstituteFunctionIndices(functions)
	for i := 0; i < len(functions); i++ {
		result.functions = append(result.functions, functions[i])
	}
	for i := 0; i < len(result.functions); i++ {
		result.functionPointers = append(result.functionPointers, &(result.functions[i]))
	}
	result.stop = false
	result.pop = false
	var stackframe = &(result.functions[result.getMainFunctionIndex()])
	result.stack = append(result.stack, generateStackframe(stackframe, &result))
	result.fillDelegateTable()
	return &result
}

func (execPtr *NVMExecutor) getMainFunctionIndex() int {
	for i := 0; i < len(execPtr.functions); i++ {
		if execPtr.functions[i].commands[0].commandName == "MAIN" {
			return i
		}
	}
	fmt.Println("ERROR: Machine cant find Main-Function!")
	return -1
}

func (execPtr *NVMExecutor) printStackframes() {
	for i := 0; i < len(execPtr.stack); i++ {
		fmt.Printf("Stackframe: %d\n", i)
		execPtr.stack[i].print()
	}
}

func (execPtr *NVMExecutor) fillDelegateTable() {
	var NOOPDelegate NOOP
	NOOPDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &NOOPDelegate)
	var REFADelegate REFA
	REFADelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &REFADelegate)
	var MOVDelegate MOV
	MOVDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &MOVDelegate)
	var SETDelegate SET
	SETDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &SETDelegate)
	var SETVDelegate SETV
	SETVDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &SETVDelegate)
	var ASETDelegate ASET
	ASETDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &ASETDelegate)
	var AGETDelegate AGET
	AGETDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &AGETDelegate)
	var CPYDelegate CPY
	CPYDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &CPYDelegate)
	var ADDDelegate ADD
	ADDDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &ADDDelegate)
	var SUBDelegate SUB
	SUBDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &SUBDelegate)
	var MULDelegate MUL
	MULDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &MULDelegate)
	var DIVDelegate DIV
	DIVDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &DIVDelegate)
	var BINORDelegate BINOR
	BINORDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &BINORDelegate)
	var BINANDDelegate BINAND
	BINANDDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &BINANDDelegate)
	var BINXORDelegate BINXOR
	BINXORDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &BINXORDelegate)
	var BINNOTDelegate BINNOT
	BINNOTDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &BINNOTDelegate)
	var LESSTHANDelegate LESSTHAN
	LESSTHANDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &LESSTHANDelegate)
	var GREATERTHANDelegate GREATERTHAN
	GREATERTHANDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &GREATERTHANDelegate)
	var LOGORDelegate LOGOR
	LOGORDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &LOGORDelegate)
	var LOGANDDelegate LOGAND
	LOGANDDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &LOGANDDelegate)
	var LOGNOTDelegate LOGNOT
	LOGNOTDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &LOGNOTDelegate)
	var LOGEQDelegate LOGEQ
	LOGEQDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &LOGEQDelegate)
	var PRINTDelegate PRINT
	PRINTDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &PRINTDelegate)
	var RETURNDelegate RETURN
	RETURNDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &RETURNDelegate)
	var CALLDelegate CALL
	CALLDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &CALLDelegate)
	var PARGDelegate PARG
	PARGDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &PARGDelegate)
	var BEQDelegate BEQ
	BEQDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &BEQDelegate)
	var LABELDelegate LABEL
	LABELDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &LABELDelegate)
	var FUNCDelegate FUNC
	FUNCDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &FUNCDelegate)
	var FENDDelegate FEND
	FENDDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &FENDDelegate)
	var GARGDelegate GARG
	GARGDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &GARGDelegate)
	var MAINDelegate MAIN
	MAINDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &MAINDelegate)
	var MENDDelegate MEND
	MENDDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &MENDDelegate)
	var RETGDelegate RETG
	RETGDelegate.setMachine(execPtr)
	execPtr.delegateTable = append(execPtr.delegateTable, &RETGDelegate)

}

func (execPtr *NVMExecutor) Run(debugPrint bool) {
	if debugPrint {
		for _, function := range execPtr.functions {
			function.print()
		}
		fmt.Println("=====================================================")
	}
	var currentCommand NVMCommand
	for !execPtr.stop && len(execPtr.stack) > 0 {
		currentCommand = execPtr.stack[len(execPtr.stack)-1].function.commands[execPtr.stack[len(execPtr.stack)-1].programCounter]
		if debugPrint {
			currentCommand.Print()
		}
		switch len(currentCommand.arguments) {
		case 0:
			execPtr.delegateTable[currentCommand.commandindex].runNoArg(&execPtr.stack[len(execPtr.stack)-1])
			break
		case 1:
			execPtr.delegateTable[currentCommand.commandindex].runOneArg(
				&execPtr.stack[len(execPtr.stack)-1],
				&currentCommand.arguments[0],
			)
			break
		case 2:
			execPtr.delegateTable[currentCommand.commandindex].runTwoArgs(
				&execPtr.stack[len(execPtr.stack)-1],
				&currentCommand.arguments[0],
				&currentCommand.arguments[1],
			)
			break
		case 3:
			execPtr.delegateTable[currentCommand.commandindex].runThreeArgs(
				&execPtr.stack[len(execPtr.stack)-1],
				&currentCommand.arguments[0],
				&currentCommand.arguments[1],
				&currentCommand.arguments[2],
			)
			break
		}
		if debugPrint {
			time.Sleep(10 * time.Millisecond)
		}
		execPtr.stack[len(execPtr.stack)-1].programCounter++
	}
}
