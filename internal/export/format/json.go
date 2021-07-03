package format

import (
	"encoding/json"
	"fmt"

	"github.com/fedy2/cwq/internal/cloudwatch"
)

func ToJson(queryDefinitions []cloudwatch.QueryDefinition) string {
	// TODO: pass indentention as argument
	queryDefinitionsJson, err := json.MarshalIndent(queryDefinitions, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	return string(queryDefinitionsJson)
}
