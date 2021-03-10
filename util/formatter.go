package util

import (
	"fmt"
	"regexp"
	"strings"
)

const idfRegex = `idf, computed as log\(1 \+ \(N - n \+ 0\.5\) \/ \(n \+ 0\.5\)\) from:`
const tfRegex = `tf, computed as freq \/ \(freq \+ k1 \* \(1 \- b \+ b \* dl \/ avgdl\)\) from:`

const DefaultExplanationIndentation = "  "
const EmptySpace = "   "
const ISpape = "│  "
const TShape = "├─ "
const LShape = "└─ "

const EmptyStartString = ""

const DocumentInfoHeaderFormat = "index: %s\ndocumentId: %s\nmatched: %t\nexplanation:\n%s"
const ExplainNodeFormat = "%s%s%g (%s)\n"
const idfFormularFormat = "idf, computed as log(1 + (%g - %g + 0.5) / (%g + 0.5))"                                                  // N, n, n
const idfFormularDetailedFormat = "idf, computed as log(1 + (%g [N] - %g [n] + 0.5) / (%g [n] + 0.5))"                              // N, n, n
const tfFormularFormat = "tf, computed as %g / (%g + %g * (1 - %g + %g * %g / %g))"                                                 // freq, freq, k1, b, b, dl, avgdl
const tfFormularDetailedFormat = "tf, computed as %g [freq] / (%g [freq] + %g [k1] * (1 - %g [b] + %g [b] * %g [dl] / %g [avgdl]))" // freq, freq, k1, b, b, dl, avgdl

func FormatExplainApiDocument(doc ExplainAPIDocument, formatOptions *FormatOptions) string {
	var formattedExplanation string

	if formatOptions.UseTreeFormat {
		formattedExplanation = formatExplainNodesToTreeFormat(EmptyStartString, true, formatOptions, doc.Explanation)
	} else {
		formattedExplanation = formatExplainNodesToSimpleFormat(0, formatOptions, doc.Explanation)
	}

	return fmt.Sprintf(DocumentInfoHeaderFormat, doc.Index, doc.DocumentId, doc.Matched, formattedExplanation)
}

func formatExplainNodesToSimpleFormat(treeLevel int, formatOptions *FormatOptions, nodes ...ExplanationNode) string {
	indentation := strings.Repeat(DefaultExplanationIndentation, treeLevel)
	treeLevel++
	var result string

	for _, node := range nodes {
		showDeeperNodes := true
		descriptionAsByteArray := []byte(node.Description)

		tf := regexp.MustCompile(tfRegex)
		idf := regexp.MustCompile(idfRegex)

		matchesTf := tf.Match(descriptionAsByteArray)
		matchesIdf := idf.Match(descriptionAsByteArray)

		if formatOptions.ShowCompactFunction && matchesIdf {
			// TODO: Alles in der Methode in eine Methode refactoren, die die fertige Zeile ausspuckt
			showDeeperNodes = false

			N := node.Details[0].Value
			n := node.Details[1].Value

			result += fmt.Sprintf("%s%g (%s)\n", indentation, node.Value,
				formatIdfFunction(N, n, formatOptions.ShowVariableNamesInFunction))
		} else if formatOptions.ShowCompactFunction && matchesTf {
			// TODO: Alles in der Methode in eine Methode refactoren, die die fertige Zeile ausspuckt
			// so sieht es dann hier aus: result += formatLineWithTfFunction
			showDeeperNodes = false

			freq := node.Details[0].Value
			k1 := node.Details[1].Value
			b := node.Details[2].Value
			dl := node.Details[3].Value
			avgdl := node.Details[4].Value

			result += fmt.Sprintf("%s%g (%s)\n", indentation, node.Value,
				formatTfFunction(freq, k1, b, dl, avgdl, formatOptions.ShowVariableNamesInFunction))
		} else {
			result += fmt.Sprintf("%s%g (%s)\n", indentation, node.Value, node.Description)
		}

		if len(node.Details) > 0 && showDeeperNodes {
			result += formatExplainNodesToSimpleFormat(treeLevel, formatOptions, node.Details...)
		}
	}

	return result
}

func formatExplainNodesToTreeFormat(previousIndentation string,
	isRootNode bool,
	formatOptions *FormatOptions,
	nodes ...ExplanationNode) string {
	var result string
	numberOfNodes := len(nodes)

	for i, node := range nodes {
		isLastInTreeLevel := isLastInTreeLevel(i, numberOfNodes)
		lineSymbol := getLineSymbol(isLastInTreeLevel, isRootNode)
		showDeeperNodes := true

		descriptionAsByteArray := []byte(node.Description)

		tf := regexp.MustCompile(tfRegex)
		idf := regexp.MustCompile(idfRegex)

		matchesTf := tf.Match(descriptionAsByteArray)
		matchesIdf := idf.Match(descriptionAsByteArray)

		if formatOptions.ShowCompactFunction && matchesIdf {
			// TODO: Alles in der Methode in eine Methode refactoren, die die fertige Zeile ausspuckt
			showDeeperNodes = false

			N := node.Details[0].Value
			n := node.Details[1].Value

			result += fmt.Sprintf(ExplainNodeFormat, previousIndentation, lineSymbol, node.Value,
				formatIdfFunction(N, n, formatOptions.ShowVariableNamesInFunction))
		} else if formatOptions.ShowCompactFunction && matchesTf {
			// TODO: Alles in der Methode in eine Methode refactoren, die die fertige Zeile ausspuckt
			showDeeperNodes = false

			freq := node.Details[0].Value
			k1 := node.Details[1].Value
			b := node.Details[2].Value
			dl := node.Details[3].Value
			avgdl := node.Details[4].Value

			result += fmt.Sprintf(ExplainNodeFormat, previousIndentation, lineSymbol, node.Value,
				formatTfFunction(freq, k1, b, dl, avgdl, formatOptions.ShowVariableNamesInFunction))
		} else {
			result += fmt.Sprintf(ExplainNodeFormat, previousIndentation, lineSymbol, node.Value, node.Description)
		}

		if len(node.Details) > 0 && showDeeperNodes {
			newIndentation := createNewIndentation(previousIndentation, isRootNode, isLastInTreeLevel)
			result += formatExplainNodesToTreeFormat(newIndentation, false, formatOptions, node.Details...)
		}
	}

	return result
}

func formatIdfFunction(N float64, n float64, showVariableNames bool) string {
	if showVariableNames {
		return fmt.Sprintf(idfFormularDetailedFormat, N, n, n)
	} else {
		return fmt.Sprintf(idfFormularFormat, N, n, n)
	}
}

func formatTfFunction(freq float64, k1 float64, b float64, dl float64, avgdl float64, showVariableNames bool) string {
	if showVariableNames {
		return fmt.Sprintf(tfFormularDetailedFormat, freq, freq, k1, b, b, dl, avgdl)
	} else {
		return fmt.Sprintf(tfFormularFormat, freq, freq, k1, b, b, dl, avgdl)
	}
}

func isLastInTreeLevel(i int, numberOfChildren int) bool {
	if i == (numberOfChildren - 1) {
		return true
	} else {
		return false
	}
}

func createNewIndentation(previousIndentation string, isFirst bool, isLastInTreeLevel bool) string {
	if isFirst {
		return DefaultExplanationIndentation
	} else if isLastInTreeLevel {
		return previousIndentation + EmptySpace
	} else {
		return previousIndentation + ISpape
	}
}

func getLineSymbol(isLastInTreeLevel bool, isFirst bool) string {
	if isFirst {
		return DefaultExplanationIndentation
	} else if isLastInTreeLevel {
		return LShape
	} else {
		return TShape
	}
}
