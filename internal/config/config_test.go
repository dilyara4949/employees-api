//go:build all
// +build all

package config

import "testing"

func TestNewConfig(t *testing.T) {
	cfg, err := NewConfig()

	if cfg.JWTTokenSecret == "" {
		t.Fatal("config jwt token secret is empty")
	}
	if cfg.Port == "" {
		t.Fatal("config port token is empty")
	}
	if cfg.Address == "" {
		t.Fatal("config address is empty")
	}
	if err != nil {
		t.Fatal(err)
	}
}
