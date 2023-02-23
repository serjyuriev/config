package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	type want struct {
		config interface{}
		err    error
	}

	tests := []struct {
		name   string
		params *parameters
		want   want
	}{
		{
			name:   "path was not set should return error",
			params: nil,
			want: want{
				config: nil,
				err:    errPathIsNotSet,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params = tt.params

			cfg, err := GetConfig()

			assert.Equal(t, tt.want.config, cfg)
			if tt.want.err != nil {
				assert.ErrorIs(t, err, tt.want.err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSetPath(t *testing.T) {
	type want struct {
		err    error
		params *parameters
	}

	tests := []struct {
		name string
		path string
		want want
	}{
		{
			name: "empty path should return error",
			path: "",
			want: want{
				err:    errPathIsEmpty,
				params: nil,
			},
		},
		{
			name: "not empty path should set path",
			path: "/some/path/to/config/file.yaml",
			want: want{
				err:    nil,
				params: &parameters{path: "/some/path/to/config/file.yaml"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetPath(tt.path)

			assert.Equal(t, tt.want.params, params)
			if tt.want.err != nil {
				assert.ErrorIs(t, err, tt.want.err)
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
