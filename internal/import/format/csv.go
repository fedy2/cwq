package format

import (
	"encoding/csv"
	"strings"

	"github.com/fedy2/cwq/internal/cloudwatch"
)

func FromCsv(data string) ([]cloudwatch.QueryDefinition, error) {

	// TODO: enable passing CSV options
	reader := csv.NewReader(strings.NewReader(data))

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return []cloudwatch.QueryDefinition{}, nil
	}

	// TODO: make header optional
	recordsWithoutHeader := records[1:]

	queryDefinitions := make([]cloudwatch.QueryDefinition, len(recordsWithoutHeader))

	for i, record := range recordsWithoutHeader {
		queryDefinitions[i] = fromRecord(record)
	}
	return queryDefinitions, nil
}

func fromRecord(record []string) cloudwatch.QueryDefinition {
	// TODO: handle different columns order or get column index from header
	var queryDefinition = cloudwatch.QueryDefinition{
		Name:              &record[0],
		QueryDefinitionId: &record[1],
		QueryString:       &record[2],
		// TODO: pass custom separator
		LogGroupNames: strings.Split(record[3], " "),
	}
	return queryDefinition
}
