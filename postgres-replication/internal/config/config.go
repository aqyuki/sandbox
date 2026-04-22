package config

import (
	"errors"
	"os"
)

type Config struct {
	PrimaryConnURL string
	ReplicaConnURL string
}

const (
	PrimaryConnURLKey = "PRIMARY_DB_URL"
	ReplicaConnURLKey = "REPLICA_DB_URL"
)

var (
	ErrEmptyPrimaryConnURL = errors.New("PRIMARY_DB_URL is empty")
	ErrEmptyReplicaConnURL = errors.New("REPLICA_DB_URL is empty")
)

func LoadEnv() (*Config, error) {
	primaryConnURL, ok := os.LookupEnv(PrimaryConnURLKey)
	if !ok {
		return nil, ErrEmptyPrimaryConnURL
	}

	replicaConnURL, ok := os.LookupEnv(ReplicaConnURLKey)
	if !ok {
		return nil, ErrEmptyReplicaConnURL
	}

	return &Config{
		PrimaryConnURL: primaryConnURL,
		ReplicaConnURL: replicaConnURL,
	}, nil
}
