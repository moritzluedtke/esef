package main

import (
	"fmt"
	"strings"
)

const TAB = "  "

func FormatExplainApiDocument(doc ExplainAPIDocument) string {
	explanation := formatExplainNodes(0, doc.Explanation)

	return fmt.Sprintf("index: %s\ndocumentId: %s\nmatched: %t\nexplanation:\n%s",
		doc.Index, doc.DocumentId, doc.Matched, explanation)
}

func formatExplainNodes(treeLevel int, nodes ...ExplainNode) string {
	indentation := strings.Repeat(TAB, treeLevel)
	treeLevel++
	var result string

	for _, node := range nodes {
		result += fmt.Sprintf(indentation+"%f (%s)\n", node.Value, node.Description)

		if len(node.Details) > 0 {
			result += formatExplainNodes(treeLevel, node.Details...)
		}
	}

	return result
}
