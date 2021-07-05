package cimport

import (
	"github.com/fedy2/cwq/internal/context"
)

type ImportCmd struct {
	File     string `arg:"" required:"" help:"File containing the query descriptions" type:"existingfile"`
	Format   string `help:"File format: csv or json" enum:"csv,json" default:"json"`
	ClearIds bool   `help:"Strips out the ids from the query definitions" default:"false"`
}

func (arguments *ImportCmd) Run(ctx *context.Context) error {
	return Import(arguments)
}
