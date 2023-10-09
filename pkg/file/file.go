package file

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadFromYAML(path string, target any) error {
	yf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yf, target)
}
