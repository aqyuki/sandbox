#!/bin/bash
set -eux

if [ ! -f "$PGDATA/PG_VERSION" ]; then
  PGPASSWORD="replicator_password" pg_basebackup -h postgres-primary -D "$PGDATA" -U replicator -Fp -Xs -P -R
fi

exec docker-entrypoint.sh "$@"
