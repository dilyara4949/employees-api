//go:build !integration

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
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"GRPC_PORT":          "qw",
				"DB_HOST":            "qw",
				"DB_PORT":            "qw",
				"DB_USER":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			want: Config{
				Address:        "address",
				RestPort:       "restport",
				JWTTokenSecret: "secret",
				GrpcPort:       "qw",
				DB: DB{
					Host:     "qw",
					Port:     "qw",
					User:     "qw",
					Password: "qw",
					Name:     "qw",
					Timeout:  1,
					MaxConn:  1,
				},
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
				t.Errorf("got: %v, wantErr: %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
