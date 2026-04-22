package db

import (
	"context"
	"fmt"

	"github.com/aqyuki/sandbox/postgres-replication/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bunexp"
)

func CreateBunConn(ctx context.Context, conf *config.Config) (*bun.DB, error) {
	primaryPool, err := createPgxPool(ctx, conf.PrimaryConnURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection to primary: %w", err)
	}
	replicaPool, err := createPgxPool(ctx, conf.ReplicaConnURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection to replica: %w", err)
	}

	primaryConn := stdlib.OpenDBFromPool(primaryPool)
	replicaConn := stdlib.OpenDBFromPool(replicaPool)

	bunDB := bun.NewDB(
		primaryConn,
		pgdialect.New(),
		bun.WithConnResolver(bunexp.NewReadWriteConnResolver(bunexp.WithDBReplica(replicaConn, bunexp.DBReplicaReadOnly))),
	)

	return bunDB, nil
}

func createPgxPool(ctx context.Context, connURL string) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(connURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	conf.PrepareConn = func(ctx context.Context, c *pgx.Conn) (bool, error) {
		return c.Ping(ctx) == nil, nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return pool, nil
}
