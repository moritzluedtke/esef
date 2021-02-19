package main

import (
	"encoding/json"
)

type ExplainAPIDocument struct {
	Index       string      `json:"_index"`
	Type        string      `json:"_type"`
	DocumentId  string      `json:"_id"`
	Matched     bool        `json:"matched"`
	Explanation ExplainNode `json:"explanation"`
}

type ExplainNode struct {
	Value       float64       `json:"value"`
	Description string        `json:"description"`
	Details     []ExplainNode `json:"details"`
}

func ExtractDataFromExplainAPI(explainAPIOutput string) ExplainAPIDocument {
	return extractDocumentFromJson(explainAPIOutput)
}

func extractDocumentFromJson(inputJson string) ExplainAPIDocument {
	var explainAPIDocument ExplainAPIDocument
	byteData := []byte(inputJson)
	err := json.Unmarshal(byteData, &explainAPIDocument)

	if err != nil {
		return ExplainAPIDocument{}
	}

	return explainAPIDocument
}
