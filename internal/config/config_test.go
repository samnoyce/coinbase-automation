package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name       string
		appEnv     string
		secretName string
		want       *Config
		wantErr    bool
	}{
		{
			name:       "APP_ENV unset or empty",
			appEnv:     "",
			secretName: "supersecret",
			want: &Config{
				AppEnv:     "dev",
				SecretName: "supersecret",
			},
		},
		{
			name:       "COINBASE_SECRET_NAME unset or empty",
			appEnv:     "",
			secretName: "",
			wantErr:    true,
		},
		{
			name:       "APP_ENV and COINBASE_SECRET_NAME set and non-empty",
			appEnv:     "prod",
			secretName: "supersecret",
			want: &Config{
				AppEnv:     "prod",
				SecretName: "supersecret",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("APP_ENV", tt.appEnv)
			t.Setenv("COINBASE_SECRET_NAME", tt.secretName)

			cfg, err := Load()

			if tt.wantErr {
				if err == nil {
					t.Fatalf("got err = nil, want non-nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("got err = %v, want nil", err)
			}

			if cfg == nil {
				t.Fatalf("got cfg = nil, want non-nil")
			}

			if cfg.AppEnv != tt.want.AppEnv {
				t.Errorf("got AppEnv = %v, want %v", cfg.AppEnv, tt.want.AppEnv)
			}

			if cfg.SecretName != tt.want.SecretName {
				t.Errorf("got SecretName = %v, want %v", cfg.SecretName, tt.want.SecretName)
			}
		})
	}
}
