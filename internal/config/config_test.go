package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]string
		want    Config
		wantErr bool
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
				"PORT":             "",
			},
			wantErr: true,
		},
		{
			name: "empty address",
			input: map[string]string{
				"PORT":    "port",
				"":        "secret",
				"ADDRESS": "",
			},
			wantErr: true,
		},
		{
			name: "empty jwt secret",
			input: map[string]string{
				"ADDRESS":          "address",
				"PORT":             "port",
				"JWT_TOKEN_SECRET": "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for confName, confValue := range tt.input {
				os.Setenv(confName, confValue)
			}
			got, err := NewConfig()
			fmt.Println(err, tt.wantErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error: %v, wantErr: %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
