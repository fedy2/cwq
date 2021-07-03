package export

import (
	"github.com/fedy2/cwq/internal/context"
)

type ExportCmd struct {
}

func (r *ExportCmd) Run(ctx *context.Context) error {
	Export()
	return nil
}
