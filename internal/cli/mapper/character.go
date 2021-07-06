package mapper

import (
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
)

type CharacterMapper struct{}

func (CharacterMapper) Decode(ctx *kong.DecodeContext, target reflect.Value) error {
	if target.Kind() != reflect.Int32 {
		return errors.Errorf("\"character\" must be applied to a rune not %s", target.Type())
	}
	var stringValue string
	err := ctx.Scan.PopValueInto("character", &stringValue)
	if err != nil {
		return err
	}
	if len(stringValue) != 1 {
		return errors.Errorf("the value should be one character")
	}
	chars := []rune(stringValue)
	target.Set(reflect.ValueOf(chars[0]))
	return nil
}
