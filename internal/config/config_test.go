package config

import "testing"

func TestNewConfig(t *testing.T) {
	cfg, err := NewConfig()

	if cfg.JWTTokenSecret == "" {
		t.Fatal("jwt token is empty")
	}
	if cfg.Port == "" {
		t.Fatal("jwt token is empty")
	}
	if err != nil {
		t.Fatal(err)
	}
}
