package config

import (
	"testing"
	"time"
)

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("HOST", "")
	t.Setenv("PORT", "")
	t.Setenv("READ_TIMEOUT", "")
	t.Setenv("WRITE_TIMEOUT", "")
	t.Setenv("IDLE_TIMEOUT", "")
	t.Setenv("HEADER_TIMEOUT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.Host != "localhost" {
		t.Fatalf("expected default host localhost, got %q", cfg.Host)
	}
	if cfg.Port != 8080 {
		t.Fatalf("expected default port 8080, got %d", cfg.Port)
	}
	if cfg.ReadTimeout != 5*time.Second {
		t.Fatalf("expected read timeout 5s, got %s", cfg.ReadTimeout)
	}
	if cfg.WriteTimeout != 20*time.Second {
		t.Fatalf("expected write timeout 20s, got %s", cfg.WriteTimeout)
	}
	if cfg.IdleTimeout != 30*time.Second {
		t.Fatalf("expected idle timeout 30s, got %s", cfg.IdleTimeout)
	}
	if cfg.HeaderTimeout != 5*time.Second {
		t.Fatalf("expected header timeout 5s, got %s", cfg.HeaderTimeout)
	}
}

func TestLoadReadsEnvironment(t *testing.T) {
	t.Setenv("HOST", "0.0.0.0")
	t.Setenv("PORT", "9090")
	t.Setenv("READ_TIMEOUT", "2s")
	t.Setenv("WRITE_TIMEOUT", "3s")
	t.Setenv("IDLE_TIMEOUT", "4s")
	t.Setenv("HEADER_TIMEOUT", "5s")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.Host != "0.0.0.0" {
		t.Fatalf("expected host from env, got %q", cfg.Host)
	}
	if cfg.Port != 9090 {
		t.Fatalf("expected port from env, got %d", cfg.Port)
	}
	if cfg.ReadTimeout != 2*time.Second ||
		cfg.WriteTimeout != 3*time.Second ||
		cfg.IdleTimeout != 4*time.Second ||
		cfg.HeaderTimeout != 5*time.Second {
		t.Fatalf("unexpected timeout config: %+v", cfg)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "not an integer", value: "abc"},
		{name: "zero", value: "0"},
		{name: "negative", value: "-1"},
		{name: "too high", value: "70000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PORT", tt.value)

			if _, err := Load(); err == nil {
				t.Fatalf("expected invalid port %q to fail", tt.value)
			}
		})
	}
}

func TestLoadRejectsInvalidDurations(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value string
	}{
		{name: "bad read timeout", key: "READ_TIMEOUT", value: "slow"},
		{name: "zero write timeout", key: "WRITE_TIMEOUT", value: "0s"},
		{name: "negative idle timeout", key: "IDLE_TIMEOUT", value: "-1s"},
		{name: "zero header timeout", key: "HEADER_TIMEOUT", value: "0s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.value)

			if _, err := Load(); err == nil {
				t.Fatalf("expected invalid duration %s=%q to fail", tt.key, tt.value)
			}
		})
	}
}
