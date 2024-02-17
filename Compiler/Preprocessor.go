package Compiler

type Preprocessor struct {
	stages []PreprocessorStage
}

func (this *Preprocessor) ProcessLines(input []string) []string {
	var result []string
	for _, line := range input {
		for _, stage := range this.stages {
			line = stage.process(line)
		}
		result = append(result, line)
	}
	return result
}

func GeneratePreprocessor() Preprocessor {
	var result Preprocessor

	result.stages = append(result.stages, &WhitespaceAdder{})
	result.stages = append(result.stages, &SemicolonAdder{})
	return result
}
