#!/bin/bash
set -eux

echo "host replication replicator all scram-sha-256" >> "$PGDATA/pg_hba.conf"
