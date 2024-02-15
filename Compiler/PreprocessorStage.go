package Compiler

type PreprocessorStage interface {
	process(line string) string
}
