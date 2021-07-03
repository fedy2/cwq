package export

import (
	"fmt"

	"github.com/fedy2/cwq/internal/cloudwatch"
	"github.com/fedy2/cwq/internal/export/format"
	"github.com/fedy2/cwq/internal/export/output"
)

func Export(arguments *ExportCmd) {
	client := cloudwatch.NewClient()

	queryDefinitions, err := client.DescribeQueryDefinitions()
	if err != nil {
		fmt.Println(err)
	}

	queryDefinitionsAsString := ""
	switch arguments.Format {
	case "json":
		queryDefinitionsAsString = format.ToJson(queryDefinitions)
	case "csv":
		queryDefinitionsAsString = format.ToCsv(queryDefinitions)
	}

	if arguments.Output != "" {
		output.ToFile(arguments.Output, queryDefinitionsAsString)
	} else {
		output.ToConsole(queryDefinitionsAsString)
	}
}
