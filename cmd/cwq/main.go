package main

import (
	"github.com/alecthomas/kong"

	"github.com/fedy2/cwq/internal/context"
	"github.com/fedy2/cwq/internal/export"
	cimport "github.com/fedy2/cwq/internal/import"
)

const (
	version = "0.0.1"
)

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Export  export.ExportCmd  `cmd:"" help:"Exports query definitions."`
	Import  cimport.ImportCmd `cmd:"" help:"Imports query definitions."`
	Version kong.VersionFlag  `name:"version" help:"Print version information and quit"`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Vars{"version": version},
		kong.UsageOnError(),
		kong.Name("cwq"),
		kong.Description("A small tool to import and export CloudWatch query definitions."))

	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&context.Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
