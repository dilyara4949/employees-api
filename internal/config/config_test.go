package config

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]string
		want    Config
		wantErr error
	}{
		{
			name: "OK postgres",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_DB":              "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"MONGO_HOST":               "mongohost",
				"MONGO_PORT":               "mongoport",
				"MONGO_USER":               "mongouser",
				"MONGO_PASSWORD":           "mongopass",
				"MONGO_NAME":               "mongodbname",
				"MONGO_TIMEOUT":            "1",
				"REDIS_HOST":               "redishost",
				"REDIS_PORT":               "redisport",
				"REDIS_PASSWORD":           "redispassword",
				"DATABASE_TYPE":            "postgres",
			},
			want: Config{
				Address:        "address",
				RestPort:       "restport",
				JWTTokenSecret: "secret",
				GrpcPort:       "grpcport",
				PostgresConfig: PostgresConfig{
					DB: DB{
						Host:     "postgreshost",
						Port:     "postgresport",
						User:     "postgresuser",
						Password: "postgrespass",
						Name:     "postgresdbname",
						Timeout:  1,
					},
					MaxConn: 2,
				},
				MongoConfig: MongoConfig{
					DB: DB{
						Host:     "mongohost",
						Port:     "mongoport",
						User:     "mongouser",
						Password: "mongopass",
						Name:     "mongodbname",
						Timeout:  1,
					},
					Collections: Collections{
						Positions: positionsCollection,
						Employees: employeesCollection,
					},
				},
				RedisConfig: RedisConfig{
					"redishost",
					"redisport",
					"redispassword",
					time.Second * 10,
					10,
					0,
					time.Hour * 5,
				},
				DatabaseType: "postgres",
			},
		},
		{
			name: "empty rest port",
			input: map[string]string{
				"ADDRESS":          "address",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingRestPort,
		},
		{
			name: "empty address",
			input: map[string]string{
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingAddress,
		},
		{
			name: "empty jwt secret",
			input: map[string]string{
				"ADDRESS":        "address",
				"REST_PORT":      "restport",
				"GRPC_PORT":      "grpcport",
				"MONGO_HOST":     "mongohost",
				"MONGO_PORT":     "mongoport",
				"MONGO_USER":     "mongouser",
				"MONGO_PASSWORD": "mongopass",
				"MONGO_NAME":     "mongodbname",
				"MONGO_TIMEOUT":  "1",
				"DATABASE_TYPE":  "mongo",
			},
			wantErr: errMissingJWTTokenSecret,
		},
		{
			name: "empty grpc port",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingGrpcPort,
		},
		{
			name: "empty mongo host",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingMongoHost,
		},
		{
			name: "empty mongo port",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingMongoPort,
		},
		{
			name: "empty mongo user",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingMongoUser,
		},
		{
			name: "empty mongo password",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_NAME":       "mongodbname",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingMongoPassword,
		},
		{
			name: "empty mongo name",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_TIMEOUT":    "1",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingMongoName,
		},
		{
			name: "empty mongo timeout",
			input: map[string]string{
				"ADDRESS":          "address",
				"REST_PORT":        "restport",
				"JWT_TOKEN_SECRET": "secret",
				"GRPC_PORT":        "grpcport",
				"MONGO_HOST":       "mongohost",
				"MONGO_PORT":       "mongoport",
				"MONGO_USER":       "mongouser",
				"MONGO_PASSWORD":   "mongopass",
				"MONGO_NAME":       "mongodbname",
				"DATABASE_TYPE":    "mongo",
			},
			wantErr: errMissingMongoTimeout,
		},
		{
			name: "empty postgres host",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_NAME":            "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"DATABASE_TYPE":            "postgres",
			},
			wantErr: errMissingPostgresHost,
		},
		{
			name: "empty postgres port",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_NAME":            "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"DATABASE_TYPE":            "postgres",
			},
			wantErr: errMissingPostgresPort,
		},
		{
			name: "empty postgres user",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_NAME":            "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"DATABASE_TYPE":            "postgres",
			},
			wantErr: errMissingPostgresUser,
		},
		{
			name: "empty postgres password",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_NAME":            "postgresdbname",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"DATABASE_TYPE":            "postgres",
			},
			wantErr: errMissingPostgresPassword,
		},
		{
			name: "empty postgres name",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_TIMEOUT":         "1",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"DATABASE_TYPE":            "postgres",
			},
			wantErr: errMissingPostgresName,
		},
		{
			name: "empty postgres timeout",
			input: map[string]string{
				"ADDRESS":                  "address",
				"REST_PORT":                "restport",
				"JWT_TOKEN_SECRET":         "secret",
				"GRPC_PORT":                "grpcport",
				"POSTGRES_HOST":            "postgreshost",
				"POSTGRES_PORT":            "postgresport",
				"POSTGRES_USER":            "postgresuser",
				"POSTGRES_PASSWORD":        "postgrespass",
				"POSTGRES_NAME":            "postgresdbname",
				"POSTGRES_MAX_CONNECTIONS": "2",
				"DATABASE_TYPE":            "postgres",
			},
			wantErr: errMissingPostgresTimeout,
		},
		{
			name: "empty postgres max connections",
			input: map[string]string{
				"ADDRESS":           "address",
				"REST_PORT":         "restport",
				"JWT_TOKEN_SECRET":  "secret",
				"GRPC_PORT":         "grpcport",
				"POSTGRES_HOST":     "postgreshost",
				"POSTGRES_PORT":     "postgresport",
				"POSTGRES_USER":     "postgresuser",
				"POSTGRES_PASSWORD": "postgrespass",
				"POSTGRES_NAME":     "postgresdbname",
				"POSTGRES_TIMEOUT":  "1",
				"DATABASE_TYPE":     "postgres",
			},
			wantErr: errMissingPostgresMaxConn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.input {
				t.Setenv(k, v)
			}

			got, err := NewConfig()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
