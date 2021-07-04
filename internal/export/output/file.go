package output

import (
	"os"
)

func ToFile(name string, data string) error {
	// TODO: handle override
	return os.WriteFile(name, []byte(data), 0644)
}
