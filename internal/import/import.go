package cimport

import (
	"github.com/fedy2/cwq/internal/cloudwatch"
	"github.com/fedy2/cwq/internal/import/format"
	"github.com/fedy2/cwq/internal/import/input"
)

func Import(arguments *ImportCmd) error {
	data, err := input.FromFile(arguments.File)
	if err != nil {
		return err
	}

	queryDefinitions := []cloudwatch.QueryDefinition{}
	err = nil
	switch arguments.Format {
	case "json":
		queryDefinitions, err = format.FromJson(data)
	case "csv":
		queryDefinitions, err = format.FromCsv(data)
	}
	if err != nil {
		return err
	}

	if arguments.ClearIds {
		for i := range queryDefinitions {
			queryDefinitions[i].QueryDefinitionId = nil
		}
	}

	client, err := cloudwatch.NewClient()
	if err != nil {
		return err
	}

	err = client.PutQueryDefinitions(queryDefinitions)
	if err != nil {
		return err
	}

	return nil
}
