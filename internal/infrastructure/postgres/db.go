package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB = pgxpool.Pool