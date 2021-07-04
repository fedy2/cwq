package input

import "io/ioutil"

func FromFile(name string) (string, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
