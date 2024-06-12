package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]string
		want    Config
		wantErr error
	}{
		{
			name: "OK",
			input: map[string]string{
				"ADDRESS":          "address",
				"PORT":             "port",
				"JWT_TOKEN_SECRET": "secret",
			},
			want: Config{
				Address:        "address",
				Port:           "port",
				JWTTokenSecret: "secret",
			},
		},
		{
			name: "empty port",
			input: map[string]string{
				"ADDRESS":          "address",
				"JWT_TOKEN_SECRET": "secret",
			},
			wantErr: errMissingPort,
		},
		{
			name: "empty address",
			input: map[string]string{
				"PORT":             "port",
				"JWT_TOKEN_SECRET": "secret",
			},
			wantErr: errMissingAddress,
		},
		{
			name: "empty jwt secret",
			input: map[string]string{
				"ADDRESS": "address",
				"PORT":    "port",
			},
			wantErr: errMissingJWTTokenSecret,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for confName, confValue := range tt.input {
				t.Setenv(confName, confValue)
			}
			got, err := NewConfig()
			fmt.Println(err, tt.wantErr)
			if (tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("NewConfig() error: %v, wantErr: %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
