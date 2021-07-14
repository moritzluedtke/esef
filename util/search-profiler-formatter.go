package util

import (
	"strings"
)

const (
	tabSize         = 2
	whitespace      = " "
	OpenBracket     = "("
	OpenBracketPlus = "+("
	ClosingBracket  = ")"
	LineBreak       = "\n"
)

func FormatSearchProfilerOutput(searchProfilerOutput string) string {
	output := strings.Replace(searchProfilerOutput, "(", "(\n", -1)
	output = strings.Replace(output, ")", "\n)", -1)
	output = strings.Replace(output, ") |", ")\n|", -1)
	output = strings.Replace(output, ") (", ")\n(", -1)
	output = strings.Replace(output, "| ", "|\n", -1)
	output = strings.Replace(output, " +", "\n+", -1)

	outputAsArray := strings.Split(output, LineBreak)

	currentIndentationLevel := 0
	var indentedOutputArray []string

	for _, line := range outputAsArray {
		line = strings.Trim(line, whitespace)

		if strings.Contains(line, ClosingBracket) {
			currentIndentationLevel--
		}

		indentedOutputArray = append(indentedOutputArray, addIndentationToLine(line, currentIndentationLevel))

		if strings.Contains(line, OpenBracket) {
			currentIndentationLevel++
		} else if strings.Contains(line, OpenBracketPlus) {
			currentIndentationLevel++
		}
	}

	return strings.Join(indentedOutputArray, LineBreak)
}

func addIndentationToLine(line string, indentationLevel int) string {
	return strings.Repeat(whitespace, tabSize*indentationLevel) + line
}
