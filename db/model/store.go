package db

import (
	"database/sql"
	"fmt"
)

type Store interface {
	CreateProfileTX(t CreateProfileTXParam) (CreateProfileTXResultType, error)
	Querier
}

type SqlStore struct {
	db *sql.DB
	*Queries
}

func SetupDB(db *sql.DB) Store {
	return &SqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SqlStore) execTX(fn func(*Queries) error) error {
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	qTX := store.WithTx(tx)
	err = fn(qTX)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
