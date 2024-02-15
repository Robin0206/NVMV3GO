package Executor

type NVMDelegate interface {
	setMachine(executor *NVMExecutor)
	runNoArg(stackframe *NVMStackframe)
	runOneArg(stackframe *NVMStackframe, a *NVMArgument)
	runTwoArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument)
	runThreeArgs(stackframe *NVMStackframe, a *NVMArgument, b *NVMArgument, c *NVMArgument)
}
