package output

import (
	"fmt"
	"os"
)

func ToFile(name string, data string) {
	// TODO: handle override
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
