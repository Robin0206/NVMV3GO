package main

func main() {
	functions := substituteFunctionIndices(splitFunctions(read("./resources/Test.nvm")))
	var executor = generateExecutor(functions)
	executor.run()
}
