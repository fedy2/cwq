package format

import (
	"encoding/json"

	"github.com/fedy2/cwq/internal/cloudwatch"
)

func ToJson(queryDefinitions []cloudwatch.QueryDefinition) (string, error) {
	// TODO: pass indentention as argument
	queryDefinitionsJson, err := json.MarshalIndent(queryDefinitions, "", " ")
	if err != nil {
		return "", err
	}
	return string(queryDefinitionsJson), nil
}
