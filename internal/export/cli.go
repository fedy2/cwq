package export

import (
	"github.com/fedy2/cwq/internal/context"
)

type ExportCmd struct {
	Output string `help:"Write to file instead of stdout" short:"o" placeholder:"file"`
	Format string `help:"Output format: csv or json" enum:"csv,json" default:"json"`
}

func (arguments *ExportCmd) Run(ctx *context.Context) error {
	Export(arguments)
	return nil
}
