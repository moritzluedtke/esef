package util

import (
	"encoding/json"
	"errors"
)

const ErrorMessage = "error during JSON Data extraction\n\nplease check your input string for any json errors"

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

func ExtractDataFromExplainAPI(explainAPIOutput string) (ExplainAPIDocument, error) {
	return extractDocumentFromJson(explainAPIOutput)
}

func extractDocumentFromJson(inputJson string) (ExplainAPIDocument, error) {
	var explainAPIDocument ExplainAPIDocument
	byteData := []byte(inputJson)
	err := json.Unmarshal(byteData, &explainAPIDocument)

	if err != nil {
		return ExplainAPIDocument{}, errors.New(ErrorMessage)
	}

	return explainAPIDocument, nil
}
