package util

const ExplainApiOptionAsText = "Explain API"
const SearchProfilerOptionAsText = "Search Profiler"

type InputOptions struct {
	_UseSearchProfilerInput bool
}

func (io *InputOptions) SetCurrentInputToExplainApi() {
	io._UseSearchProfilerInput = false
}

func (io *InputOptions) SetCurrentInputToSearchProfiler() {
	io._UseSearchProfilerInput = true
}

func (io *InputOptions) GetCurrentInputOption() string {
	if io._UseSearchProfilerInput {
		return SearchProfilerOptionAsText
	} else {
		return ExplainApiOptionAsText
	}
}
