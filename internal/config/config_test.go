//go:build !integration
// +build !integration

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
				Mongo: Mongo{
					Collections{
						Positions: positionsCollection,
						Employees: employeesCollection,
					},
				},
			},
		},
		{
			name: "empty  rest port",
			input: map[string]string{
				"ADDRESS":            "address",
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
			wantErr: errMissingRestPort,
		},
		{
			name: "empty address",
			input: map[string]string{
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
			}, wantErr: errMissingAddress,
		},
		{
			name: "empty jwt secret",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"GRPC_PORT":          "qw",
				"DB_HOST":            "qw",
				"DB_PORT":            "qw",
				"DB_USER":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingJWTTokenSecret,
		},
		{
			name: "empty grpc port",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"DB_HOST":            "qw",
				"DB_PORT":            "qw",
				"DB_USER":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingGrpcPort,
		},
		{
			name: "empty db host",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"GRPC_PORT":          "qw",
				"DB_PORT":            "qw",
				"DB_USER":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingDbHost,
		},
		{
			name: "empty db port",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"GRPC_PORT":          "qw",
				"DB_HOST":            "qw",
				"DB_USER":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingDbPort,
		}, {
			name: "empty db user",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"GRPC_PORT":          "qw",
				"DB_HOST":            "qw",
				"DB_PORT":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingDbUser,
		}, {
			name: "empty db password",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"GRPC_PORT":          "qw",
				"DB_HOST":            "qw",
				"DB_PORT":            "qw",
				"DB_USER":            "qw",
				"DB_NAME":            "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingDbPassword,
		}, {
			name: "empty db name",
			input: map[string]string{
				"ADDRESS":            "address",
				"REST_PORT":          "restport",
				"JWT_TOKEN_SECRET":   "secret",
				"GRPC_PORT":          "qw",
				"DB_HOST":            "qw",
				"DB_PORT":            "qw",
				"DB_USER":            "qw",
				"DB_PASSWORD":        "qw",
				"DB_TIMEOUT":         "1",
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingDbName,
		},
		{
			name: "empty db timeout",
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
				"DB_MAX_CONNECTIONS": "1",
			},
			wantErr: errMissingDbTimeout,
		},
		{
			name: "empty db maxconnections",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "qw",
				"DB_HOST":          "qw",
				"DB_PORT":          "qw",
				"DB_USER":          "qw",
				"DB_PASSWORD":      "qw",
				"DB_NAME":          "qw",
				"DB_TIMEOUT":       "1",
			},
			wantErr: errMissingDbMaxConn,
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
