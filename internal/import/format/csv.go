package format

import (
	"encoding/csv"
	"strings"

	"github.com/fedy2/cwq/internal/cloudwatch"
)

type CsvOptions struct {
	HasHeader bool
	Delimiter rune
	Comment   rune
}

func FromCsv(data string, options CsvOptions) ([]cloudwatch.QueryDefinition, error) {

	reader := csv.NewReader(strings.NewReader(data))

	reader.Comma = options.Delimiter
	reader.Comment = options.Comment

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return []cloudwatch.QueryDefinition{}, nil
	}

	if options.HasHeader {
		records = records[1:]
	}

	queryDefinitions := make([]cloudwatch.QueryDefinition, len(records))

	for i, record := range records {
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
