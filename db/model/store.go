package db

import "database/sql"

type Store interface {
	Querier
}

type SqlStore struct {
	*Queries
}

func SetupDB(db *sql.DB) Store {
	return &SqlStore{
		Queries: New(db),
	}
}
