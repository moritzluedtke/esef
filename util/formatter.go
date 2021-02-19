package util

import (
	"fmt"
	"strings"
)

const DefaultExplanationIndentation = "  "
const EmptySpace = "   "
const ISpape = "│  "
const TShape = "├─ "
const LShape = "└─ "

const EmptyStartString = ""

const DocumentInfoHeaderFormat = "index: %s\ndocumentId: %s\nmatched: %t\nexplanation:\n%s"
const ExplainNoteFormat = "%s%s%f (%s)\n"

func FormatExplainApiDocument(doc ExplainAPIDocument, useTreeFormat bool) string {
	var formattedExplanation string

	if useTreeFormat {
		formattedExplanation = formatExplainNodesToTreeFormat(EmptyStartString, true, doc.Explanation)
	} else {
		formattedExplanation = formatExplainNodesToSimpleFormat(0, doc.Explanation)
	}

	return fmt.Sprintf(DocumentInfoHeaderFormat, doc.Index, doc.DocumentId, doc.Matched, formattedExplanation)
}

func formatExplainNodesToSimpleFormat(treeLevel int, nodes ...ExplainNode) string {
	indentation := strings.Repeat(DefaultExplanationIndentation, treeLevel)
	treeLevel++
	var result string

	for _, node := range nodes {
		result += fmt.Sprintf(indentation+"%f (%s)\n", node.Value, node.Description)

		if len(node.Details) > 0 {
			result += formatExplainNodesToSimpleFormat(treeLevel, node.Details...)
		}
	}

	return result
}

func formatExplainNodesToTreeFormat(previousIndentation string, isRootNode bool, nodes ...ExplainNode) string {
	var result string
	numberOfNodes := len(nodes)

	for i, node := range nodes {
		isLastInTreeLevel := isLastInTreeLevel(i, numberOfNodes)
		lineSymbol := getLineSymbol(isLastInTreeLevel, isRootNode)

		result += fmt.Sprintf(ExplainNoteFormat, previousIndentation, lineSymbol, node.Value, node.Description)

		if len(node.Details) > 0 {
			newIndentation := createNewIndentation(previousIndentation, isRootNode, isLastInTreeLevel)

			result += formatExplainNodesToTreeFormat(newIndentation, false, node.Details...)
		}
	}

	return result
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
