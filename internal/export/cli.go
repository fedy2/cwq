package export

import (
	"github.com/fedy2/cwq/internal/context"
)

type ExportCmd struct {
	Output           string `help:"Write to file instead of stdout" short:"o" placeholder:"file"`
	Format           string `help:"Output format: csv or json" enum:"csv,json" default:"json"`
	Prefix           string `help:"Exports the query definitions with a name that starts with the specified prefix."`
	CsvIncludeHeader bool   `help:"Adds an header to the CSV file. Default: true" default:"true" negatable:""`
}

func (arguments *ExportCmd) Run(ctx *context.Context) error {
	return Export(arguments)
}
