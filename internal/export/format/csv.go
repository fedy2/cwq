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

type CsvOptions struct {
	IncludeHeader bool
	Delimiter     rune
}

func ToCsv(queryDefinitions []cloudwatch.QueryDefinition, options CsvOptions) (string, error) {

	buffer := new(bytes.Buffer)

	writer := csv.NewWriter(buffer)
	writer.Comma = options.Delimiter

	if options.IncludeHeader {
		err := writer.Write(headerRecord)
		if err != nil {
			return "", err
		}
	}

	for _, queryDefinition := range queryDefinitions {
		var record = toRecord(queryDefinition)
		err := writer.Write(record)
		if err != nil {
			return "", err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return buffer.String(), nil
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
