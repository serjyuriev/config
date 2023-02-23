package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

type parameters struct {
	config interface{}
	once   sync.Once
	path   string
}

var (
	errPathIsEmpty             = fmt.Errorf("path can't be empty")
	errPathIsNotSet            = fmt.Errorf("path wasn't set")
	errReadingConfigFile       = fmt.Errorf("unable to read config file")
	errUnmarshallingConfigFile = fmt.Errorf("unable to unmarshal yaml in config file")

	params *parameters
)

func GetConfig() (interface{}, error) {
	if params == nil {
		return nil, errPathIsNotSet
	}

	var externalErr error = nil
	params.once.Do(func() {
		file, err := os.ReadFile(params.path)
		if err != nil {
			externalErr = errors.Join(errReadingConfigFile, err)
			return
		}

		if err = yaml.Unmarshal(file, &params.config); err != nil {
			externalErr = errors.Join(errUnmarshallingConfigFile, err)
			return
		}
	})
	return params.config, externalErr
}

func SetPath(path string) error {
	if path == "" {
		return errPathIsEmpty
	}

	params = &parameters{
		config: nil,
		path:   path,
	}

	return nil
}
