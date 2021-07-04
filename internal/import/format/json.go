package format

import (
	"encoding/json"

	"github.com/fedy2/cwq/internal/cloudwatch"
)

func FromJson(data string) ([]cloudwatch.QueryDefinition, error) {
	var queryDefinitions []cloudwatch.QueryDefinition
	err := json.Unmarshal([]byte(data), &queryDefinitions)
	if err != nil {
		return nil, err
	}
	return queryDefinitions, nil
}
