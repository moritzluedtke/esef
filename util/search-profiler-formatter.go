package util

import (
	"strings"
)

const (
	tabSize    = 2
	whitespace = " "
)

func FormatSearchProfilerOutput(searchProfilerOutput string) string {
	output := strings.Replace(searchProfilerOutput, "(", "(\n", -1)
	output = strings.Replace(output, ")", "\n)", -1)
	output = strings.Replace(output, ") |", ")\n|", -1)
	output = strings.Replace(output, ") (", ")\n(", -1)
	output = strings.Replace(output, "| ", "|\n", -1)
	output = strings.Replace(output, " +", "\n+", -1)

	outputAsArray := strings.Split(output, "\n")

	currentIndentationLevel := 0
	var indentedOutputArray []string

	for _, line := range outputAsArray {
		line = strings.Trim(line, whitespace)

		if strings.Contains(line, ")") {
			currentIndentationLevel--
		}

		indentedOutputArray = append(indentedOutputArray, addIndentationToLine(line, currentIndentationLevel))

		if strings.Contains(line, "(") {
			currentIndentationLevel++
		} else if strings.Contains(line, "+(") {
			currentIndentationLevel++
		}
	}

	return strings.Join(indentedOutputArray, "\n")
}

func addIndentationToLine(line string, indentationLevel int) string {
	return strings.Repeat(whitespace, tabSize*indentationLevel) + line
}
