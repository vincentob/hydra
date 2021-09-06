package yaml

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func Marshal(in interface{}) (out []byte, err error) {
	return yaml.Marshal(in)
}

func Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}

func UnmarshalFromFile(filename string, out interface{}) (err error) {
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "read yaml file failed")
	}

	return yaml.Unmarshal(bts, out)
}

func MarshalToFile(filename string, data interface{}) error {
	var err error
	var out []byte

	if out, err = yaml.Marshal(data); err != nil {
		return errors.Wrap(err, "marshal data failed")
	}

	if err = ioutil.WriteFile(filename, out, 0600); err != nil {
		return errors.Wrap(err, "write file failed")
	}

	return nil
}
