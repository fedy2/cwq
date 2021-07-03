package cimport

import (
	"fmt"

	"github.com/fedy2/cwq/internal/context"
)

type ImportCmd struct {
}

func (r *ImportCmd) Run(ctx *context.Context) error {
	fmt.Println("Import")
	return nil
}
