package Executor

import (
	"fmt"
	"math/rand"
	"os"
)

type NOOP struct{ executor *NVMExecutor }

func (this *NOOP) setMachine(executor *NVMExecutor)   { this.executor = executor }
func (this *NOOP) runNoArg(stackframe *NVMStackframe) {}
func (this *NOOP) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	var _ = 0
}
func (this *NOOP) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
}
func (this *NOOP) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
}

type REFA struct{ executor *NVMExecutor }

func (this *REFA) setMachine(executor *NVMExecutor) {
	this.executor = executor
}
func (this *REFA) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate REFA, Method runNoArg called!")
}
func (this *REFA) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate REFA, Method runOneArg called!")
}
func (this *REFA) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate REFA, Method runTwoArgs called!")
}
func (this *REFA) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	stackframe.variables = append(stackframe.variables, generateNVMVariable(b.integer, c.integer))
}

type MOV struct{ executor *NVMExecutor }

func (this *MOV) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *MOV) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate MOV, Method runNoArg called!")
}
func (this *MOV) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate MOV, Method runOneArg called!")
}
func (this *MOV) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate MOV, Method runTwoArgs called!")
}
func (this *MOV) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate MOV, Method runThreeArgs called!")
}

type SET struct{ executor *NVMExecutor }

func (this *SET) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *SET) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate SET, Method runNoArg called!")
}
func (this *SET) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate SET, Method runOneArg called!")
}
func (this *SET) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	switch stackframe.variables[a.integer].valueType {
	case 0:
		if b.integer == 1 {
			stackframe.variables[a.integer].boolValue[0] = true
		} else {
			stackframe.variables[a.integer].boolValue[0] = false
		}
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = uint8(b.integer)
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = int32(b.integer)
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = b.realValue
		break
	}
}
func (this *SET) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate SET, Method runThreeArgs called!")
}

type SETV struct{ executor *NVMExecutor }

func (this *SETV) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *SETV) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate SETV, Method runNoArg called!")
}
func (this *SETV) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate SETV, Method runOneArg called!")
}
func (this *SETV) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	if stackframe.variables[a.integer].valueType != stackframe.variables[b.integer].valueType {
		fmt.Println("ERROR: Types of args dont match in SETV")
	}
	switch stackframe.variables[a.integer].valueType {
	case 0:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].boolValue[0]
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = stackframe.variables[b.integer].realValue[0]
		break
	}
}
func (this *SETV) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate SETV, Method runThreeArgs called!")
}

type ASET struct{ executor *NVMExecutor }

func (this *ASET) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *ASET) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate ASET, Method runNoArg called!")
}
func (this *ASET) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate ASET, Method runOneArg called!")
}
func (this *ASET) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate ASET, Method runTwoArgs called!")
}
func (this *ASET) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {

	switch stackframe.variables[a.integer].valueType {
	case 0:
		stackframe.variables[a.integer].boolValue[stackframe.variables[b.integer].integerValue[0]] = stackframe.variables[c.integer].boolValue[0]
		break
	case 1:
		stackframe.variables[a.integer].byteValue[stackframe.variables[b.integer].integerValue[0]] = stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[stackframe.variables[b.integer].integerValue[0]] = stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].realValue[stackframe.variables[b.integer].integerValue[0]] = stackframe.variables[c.integer].realValue[0]
		break
	}
}

type AGET struct{ executor *NVMExecutor }

func (this *AGET) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *AGET) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate AGET, Method runNoArg called!")
}
func (this *AGET) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate AGET, Method runOneArg called!")
}
func (this *AGET) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate AGET, Method runTwoArgs called!")
}
func (this *AGET) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) { //dstindex, srcValArr, srcIndex
	switch stackframe.variables[a.integer].valueType {
	case 0:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].boolValue[stackframe.variables[c.integer].integerValue[0]]
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[stackframe.variables[c.integer].integerValue[0]]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[stackframe.variables[c.integer].integerValue[0]]
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = stackframe.variables[b.integer].realValue[stackframe.variables[c.integer].integerValue[0]]
		break
	}
}

type CPY struct{ executor *NVMExecutor }

func (this *CPY) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *CPY) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate CPY, Method runNoArg called!")
}
func (this *CPY) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate CPY, Method runOneArg called!")
}
func (this *CPY) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate CPY, Method runTwoArgs called!")
}
func (this *CPY) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate CPY, Method runThreeArgs called!")
}

type ADD struct{ executor *NVMExecutor }

func (this *ADD) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *ADD) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate ADD, Method runNoArg called!")
}
func (this *ADD) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate ADD, Method runOneArg called!")
}
func (this *ADD) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate ADD, Method runTwoArgs called!")
}
func (this *ADD) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called ADD on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] + stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] + stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = stackframe.variables[b.integer].realValue[0] + stackframe.variables[c.integer].realValue[0]
		break
	}
}

type SUB struct{ executor *NVMExecutor }

func (this *SUB) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *SUB) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate SUB, Method runNoArg called!")
}
func (this *SUB) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate SUB, Method runOneArg called!")
}
func (this *SUB) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate SUB, Method runTwoArgs called!")
}
func (this *SUB) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called SUB on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] - stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] - stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = stackframe.variables[b.integer].realValue[0] - stackframe.variables[c.integer].realValue[0]
		break
	}
}

type MUL struct{ executor *NVMExecutor }

func (this *MUL) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *MUL) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate MUL, Method runNoArg called!")
}
func (this *MUL) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate MUL, Method runOneArg called!")
}
func (this *MUL) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate MUL, Method runTwoArgs called!")
}
func (this *MUL) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called MUL on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] * stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] * stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = stackframe.variables[b.integer].realValue[0] * stackframe.variables[c.integer].realValue[0]
		break
	}
}

type DIV struct{ executor *NVMExecutor }

func (this *DIV) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *DIV) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate DIV, Method runNoArg called!")
}
func (this *DIV) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate DIV, Method runOneArg called!")
}
func (this *DIV) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate DIV, Method runTwoArgs called!")
}
func (this *DIV) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called DIV on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] / stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] / stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].realValue[0] = stackframe.variables[b.integer].realValue[0] / stackframe.variables[c.integer].realValue[0]
		break
	}
}

type MOD struct {
	executor *NVMExecutor
}

func (this *MOD) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *MOD) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate MOD, Method runNoArg called!")
}
func (this *MOD) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate MOD, Method runOneArg called!")
}
func (this *MOD) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate MOD, Method runTwoArgs called!")
}
func (this *MOD) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called MOD on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] % stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] % stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		fmt.Println("ERROR: Called MOD on real values")
		break
	}
}

type BINOR struct{ executor *NVMExecutor }

func (this *BINOR) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *BINOR) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate BINOR, Method runNoArg called!")
}
func (this *BINOR) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate BINOR, Method runOneArg called!")
}
func (this *BINOR) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate BINOR, Method runTwoArgs called!")
}
func (this *BINOR) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called BINOR on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] | stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] | stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		fmt.Println("ERROR: Called BINOR on real values")
		break
	}
}

type BINAND struct{ executor *NVMExecutor }

func (this *BINAND) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *BINAND) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate BINAND, Method runNoArg called!")
}
func (this *BINAND) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate BINAND, Method runOneArg called!")
}
func (this *BINAND) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate BINAND, Method runTwoArgs called!")
}
func (this *BINAND) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called BINAND on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] & stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] & stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		fmt.Println("ERROR: Called BINAND on real values")
		break
	}
}

type BINXOR struct{ executor *NVMExecutor }

func (this *BINXOR) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *BINXOR) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate BINXOR, Method runNoArg called!")
}
func (this *BINXOR) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate BINXOR, Method runOneArg called!")
}
func (this *BINXOR) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate BINXOR, Method runTwoArgs called!")
}
func (this *BINXOR) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called BINXOR on bool values")
		break
	case 1:
		stackframe.variables[a.integer].byteValue[0] = stackframe.variables[b.integer].byteValue[0] ^ stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].integerValue[0] = stackframe.variables[b.integer].integerValue[0] ^ stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		fmt.Println("ERROR: Called BINXOR on real values")
		break
	}
}

type BINNOT struct{ executor *NVMExecutor }

func (this *BINNOT) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *BINNOT) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate BINNOT, Method runNoArg called!")
}
func (this *BINNOT) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate BINNOT, Method runOneArg called!")
}
func (this *BINNOT) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate BINNOT, Method runTwoArgs called!")
}
func (this *BINNOT) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate BINNOT, Method runThreeArgs called!")
}

type LESSTHAN struct{ executor *NVMExecutor }

func (this *LESSTHAN) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *LESSTHAN) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate LESSTHAN, Method runNoArg called!")
}
func (this *LESSTHAN) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate LESSTHAN, Method runOneArg called!")
}
func (this *LESSTHAN) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate LESSTHAN, Method runTwoArgs called!")
}
func (this *LESSTHAN) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called LESSTHAN on bool values")
		break
	case 1:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].byteValue[0] < stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].integerValue[0] < stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].realValue[0] < stackframe.variables[c.integer].realValue[0]
		break
	}
}

type GREATERTHAN struct{ executor *NVMExecutor }

func (this *GREATERTHAN) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *GREATERTHAN) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate GREATERTHAN, Method runNoArg called!")
}
func (this *GREATERTHAN) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate GREATERTHAN, Method runOneArg called!")
}
func (this *GREATERTHAN) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate GREATERTHAN, Method runTwoArgs called!")
}
func (this *GREATERTHAN) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		fmt.Println("ERROR: Called GREATERTHAN on bool values")
		break
	case 1:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].byteValue[0] > stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].integerValue[0] > stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].realValue[0] > stackframe.variables[c.integer].realValue[0]
		break
	}
}

type LOGOR struct{ executor *NVMExecutor }

func (this *LOGOR) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *LOGOR) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate LOGOR, Method runNoArg called!")
}
func (this *LOGOR) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGOR, Method runOneArg called!")
}
func (this *LOGOR) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGOR, Method runTwoArgs called!")
}
func (this *LOGOR) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].boolValue[0] || stackframe.variables[c.integer].boolValue[0]
}

type LOGAND struct{ executor *NVMExecutor }

func (this *LOGAND) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *LOGAND) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate LOGAND, Method runNoArg called!")
}
func (this *LOGAND) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGAND, Method runOneArg called!")
}
func (this *LOGAND) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGAND, Method runTwoArgs called!")
}
func (this *LOGAND) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].boolValue[0] && stackframe.variables[c.integer].boolValue[0]
}

type LOGNOT struct{ executor *NVMExecutor }

func (this *LOGNOT) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *LOGNOT) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate LOGNOT, Method runNoArg called!")
}
func (this *LOGNOT) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGNOT, Method runOneArg called!")
}
func (this *LOGNOT) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	stackframe.variables[a.integer].boolValue[0] = !(stackframe.variables[b.integer].boolValue[0])
}
func (this *LOGNOT) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGNOT, Method runThreeArgs called!")
}

type LOGEQ struct{ executor *NVMExecutor }

func (this *LOGEQ) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *LOGEQ) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate LOGEQ, Method runNoArg called!")
}
func (this *LOGEQ) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGEQ, Method runOneArg called!")
}
func (this *LOGEQ) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate LOGEQ, Method runTwoArgs called!")
}
func (this *LOGEQ) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	switch stackframe.variables[b.integer].valueType {
	case 0:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].boolValue[0] == stackframe.variables[c.integer].boolValue[0]
		break
	case 1:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].byteValue[0] == stackframe.variables[c.integer].byteValue[0]
		break
	case 2:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].integerValue[0] == stackframe.variables[c.integer].integerValue[0]
		break
	case 3:
		stackframe.variables[a.integer].boolValue[0] = stackframe.variables[b.integer].realValue[0] == stackframe.variables[c.integer].realValue[0]
		break
	}
}

type PRINT struct{ executor *NVMExecutor }

func (this *PRINT) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *PRINT) runNoArg(stackframe *NVMStackframe) {
	fmt.Println()
}
func (this *PRINT) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	switch stackframe.variables[a.integer].valueType {
	case 0:
		fmt.Print(stackframe.variables[a.integer].boolValue[0])
		break
	case 1:
		var str = string(stackframe.variables[a.integer].byteValue[0])
		fmt.Print(str)
		break
	case 2:
		fmt.Print(stackframe.variables[a.integer].integerValue[0])
		break
	case 3:
		fmt.Printf("%f", stackframe.variables[a.integer].realValue[0])
		break
	}
}
func (this *PRINT) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate PRINT, Method runTwoArgs called!")
}
func (this *PRINT) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate PRINT, Method runThreeArgs called!")
}

type RETURN struct{ executor *NVMExecutor }

func (this *RETURN) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *RETURN) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate RETURN, Method runNoArg called!")
}
func (this *RETURN) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	switch stackframe.variables[a.integer].valueType {
	case 0:
		stackframe.executor.returnRegister.boolValue = stackframe.variables[a.integer].boolValue
		break
	case 1:
		stackframe.executor.returnRegister.byteValue = stackframe.variables[a.integer].byteValue
		break
	case 2:
		stackframe.executor.returnRegister.integerValue = stackframe.variables[a.integer].integerValue
		break
	case 3:
		stackframe.executor.returnRegister.realValue = stackframe.variables[a.integer].realValue
		break
	}
}
func (this *RETURN) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate RETURN, Method runTwoArgs called!")
}
func (this *RETURN) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate RETURN, Method runThreeArgs called!")
}

type CALL struct{ executor *NVMExecutor }

func (this *CALL) setMachine(executor *NVMExecutor) {
	this.executor = executor
}
func (this *CALL) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate CALL, Method runNoArg called!")
}
func (this *CALL) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate CALL, Method runOneArg called!")
}
func (this *CALL) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	var functionAddress = &stackframe.executor.functions[a.integer]
	var newStackframe = generateStackframe(functionAddress, stackframe.executor)
	stackframe.executor.stack = append(stackframe.executor.stack, newStackframe)
}
func (this *CALL) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate CALL, Method runThreeArgs called!")
}

type PARG struct{ executor *NVMExecutor }

func (this *PARG) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *PARG) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate PARG, Method runNoArg called!")
}
func (this *PARG) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	stackframe.executor.argRegister = append(stackframe.executor.argRegister, stackframe.variables[a.integer])
}
func (this *PARG) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate PARG, Method runTwoArgs called!")
}
func (this *PARG) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate PARG, Method runThreeArgs called!")
}

type BEQ struct{ executor *NVMExecutor }

func (this *BEQ) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *BEQ) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate BEQ, Method runNoArg called!")
}
func (this *BEQ) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate BEQ, Method runOneArg called!")
}
func (this *BEQ) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate BEQ, Method runTwoArgs called!")
}
func (this *BEQ) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	if stackframe.variables[a.integer].boolValue[0] && stackframe.variables[b.integer].boolValue[0] {
		stackframe.programCounter = stackframe.function.labels[c.integer]
	}
}

type LABEL struct{ executor *NVMExecutor }

func (this *LABEL) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *LABEL) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate LABEL, Method runNoArg called!")
}
func (this *LABEL) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate LABEL, Method runOneArg called!")
}
func (this *LABEL) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate LABEL, Method runTwoArgs called!")
}
func (this *LABEL) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate LABEL, Method runThreeArgs called!")
}

type FUNC struct{ executor *NVMExecutor }

func (this *FUNC) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *FUNC) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate FUNC, Method runNoArg called!")
}
func (this *FUNC) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate FUNC, Method runOneArg called!")
}
func (this *FUNC) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate FUNC, Method runTwoArgs called!")
}
func (this *FUNC) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate FUNC, Method runThreeArgs called!")
}

type FEND struct{ executor *NVMExecutor }

func (this *FEND) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *FEND) runNoArg(stackframe *NVMStackframe) {
	stackframe.executor.stack = stackframe.executor.stack[:len(stackframe.executor.stack)-1]
	stackframe.executor.argRegister = nil
}
func (this *FEND) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate FEND, Method runOneArg called!")
}
func (this *FEND) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate FEND, Method runTwoArgs called!")
}
func (this *FEND) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate FEND, Method runThreeArgs called!")
}

type GARG struct{ executor *NVMExecutor }

func (this *GARG) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *GARG) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate GARG, Method runNoArg called!")
}
func (this *GARG) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate GARG, Method runOneArg called!")
}
func (this *GARG) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	stackframe.variables[a.integer] = stackframe.executor.argRegister[b.integer]
	if b.integer == len(stackframe.executor.argRegister)-1 {
		stackframe.executor.argRegister = nil
	}
}
func (this *GARG) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate GARG, Method runThreeArgs called!")
}

type MAIN struct{ executor *NVMExecutor }

func (this *MAIN) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *MAIN) runNoArg(stackframe *NVMStackframe) {
}
func (this *MAIN) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate MAIN, Method runOneArg called!")
}
func (this *MAIN) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate MAIN, Method runTwoArgs called!")
}
func (this *MAIN) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate MAIN, Method runThreeArgs called!")
}

type MEND struct{ executor *NVMExecutor }

func (this *MEND) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *MEND) runNoArg(stackframe *NVMStackframe) {
	os.Exit(0)
}
func (this *MEND) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate MEND, Method runOneArg called!")
}
func (this *MEND) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate MEND, Method runTwoArgs called!")
}
func (this *MEND) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate MEND, Method runThreeArgs called!")
}

type RETG struct{ executor *NVMExecutor }

func (this *RETG) setMachine(executor *NVMExecutor) { this.executor = executor }
func (this *RETG) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate RETG, Method runNoArg called!")
}
func (this *RETG) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	switch stackframe.variables[a.integer].valueType {
	case 0:
		stackframe.variables[a.integer].boolValue = stackframe.executor.returnRegister.boolValue
		break
	case 1:
		stackframe.variables[a.integer].byteValue = stackframe.executor.returnRegister.byteValue
		break
	case 2:
		stackframe.variables[a.integer].integerValue = stackframe.executor.returnRegister.integerValue
		break
	case 3:
		stackframe.variables[a.integer].realValue = stackframe.executor.returnRegister.realValue
		break
	}
	stackframe.executor.argRegister = nil
}
func (this *RETG) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	fmt.Println("ERROR: Delegate RETG, Method runTwoArgs called!")
}
func (this *RETG) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate RETG, Method runThreeArgs called!")
}

type RAND struct{ executor *NVMExecutor }

func (this *RAND) setMachine(executor *NVMExecutor) {
	this.executor = executor
}
func (this *RAND) runNoArg(stackframe *NVMStackframe) {
	fmt.Println("ERROR: Delegate RAND, Method runNoArg called!")
}
func (this *RAND) runOneArg(stackframe *NVMStackframe, a *NVMArgument) {
	fmt.Println("ERROR: Delegate RAND, Method runOneArg called!")
}
func (this *RAND) runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument) {
	switch stackframe.variables[a.integer].valueType {
	case 0:
		stackframe.variables[a.integer].boolValue[stackframe.variables[b.integer].integerValue[0]] = rand.Int()%2 == 0
	case 1:
		stackframe.variables[a.integer].byteValue[stackframe.variables[b.integer].integerValue[0]] = uint8(rand.Uint32() % 256)
	case 2:
		stackframe.variables[a.integer].integerValue[stackframe.variables[b.integer].integerValue[0]] = rand.Int31()
	case 3:
		stackframe.variables[a.integer].realValue[stackframe.variables[b.integer].integerValue[0]] = rand.Float64()
	}
}
func (this *RAND) runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument) {
	fmt.Println("ERROR: Delegate RAND, Method runThreeArgs called!")
}
