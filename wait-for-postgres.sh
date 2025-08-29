#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

echo "Waiting for PostgreSQL at $host..."

# Loop until Postgres is ready
until pg_isready -h "$host" -p 5432 -U "$DB_USER"; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

echo "PostgreSQL is up - executing command"
exec $cmd
