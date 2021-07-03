package format

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/fedy2/cwq/internal/cloudwatch"
)

var headerRecord = []string{
	"Name",
	"QueryDefinitionId",
	"QueryString",
	"LastModified",
	"LogGroupNames",
}

func ToCsv(queryDefinitions []cloudwatch.QueryDefinition) string {

	buffer := new(bytes.Buffer)

	// TODO: enable passing CSV options
	writer := csv.NewWriter(buffer)

	// TODO: make header optional
	writer.Write(headerRecord)

	for _, queryDefinition := range queryDefinitions {
		var record = toRecord(queryDefinition)
		writer.Write(record)
	}

	writer.Flush()

	return buffer.String()
}

func toRecord(queryDefinition cloudwatch.QueryDefinition) []string {
	var record = make([]string, 5)
	record[0] = *queryDefinition.Name
	record[1] = *queryDefinition.QueryDefinitionId
	record[2] = *queryDefinition.QueryString
	record[3] = fmt.Sprint(*queryDefinition.LastModified)
	// TODO: pass separator as option
	record[4] = strings.Join(queryDefinition.LogGroupNames, " ")
	return record
}
