package system

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

//ReadYamlFile reads data from yaml file
func ReadYamlFile(filepath string, o interface{}) (exists bool, err error) {
	var data []byte
	_, err = os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			WriteYamlFile(filepath, o)
			return false, nil
		}
		return false, errors.Wrap(err, fmt.Sprintf("unable to open file for read %s", filepath))
	}
	data, err = ioutil.ReadFile(filepath)
	if err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("failed to read file for read %s", filepath))
	}
	return true, yaml.Unmarshal(data, o)
}

//WriteYamlFile writes to a yaml file
func WriteYamlFile(filepath string, o interface{}) error {
	data, err := yaml.Marshal(o)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to marshal data %#v", o))
	}
	_, err = os.Stat(filepath)
	if err != nil && !os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("unable to open file for write %s", filepath))
	}
	err = ioutil.WriteFile(filepath, data, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "unable to write data to file")
	}
	return nil
}
