package export

import (
	"github.com/fedy2/cwq/internal/cloudwatch"
	"github.com/fedy2/cwq/internal/export/format"
	"github.com/fedy2/cwq/internal/export/output"
)

func Export(arguments *ExportCmd) error {
	client, err := cloudwatch.NewClient()
	if err != nil {
		return err
	}

	queryDefinitions, err := client.DescribeQueryDefinitions(arguments.Prefix)
	if err != nil {
		return err
	}

	queryDefinitionsAsString := ""
	err = nil
	switch arguments.Format {
	case "json":
		queryDefinitionsAsString, err = format.ToJson(queryDefinitions)
	case "csv":
		options := format.CsvOptions{
			IncludeHeader: arguments.CsvIncludeHeader,
		}
		queryDefinitionsAsString, err = format.ToCsv(queryDefinitions, options)
	}
	if err != nil {
		return err
	}

	if arguments.Output != "" {
		err = output.ToFile(arguments.Output, queryDefinitionsAsString)
		if err != nil {
			return err
		}
	} else {
		output.ToConsole(queryDefinitionsAsString)
	}

	return nil
}
