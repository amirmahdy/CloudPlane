// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: credentials.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createCredential = `-- name: CreateCredential :one
INSERT INTO credentials (id, access_id, secret_key)
VALUES ($1, $2, $3)
returning id, access_id, secret_key
`

type CreateCredentialParams struct {
	ID        uuid.UUID `json:"id"`
	AccessID  string    `json:"access_id"`
	SecretKey string    `json:"secret_key"`
}

func (q *Queries) CreateCredential(ctx context.Context, arg CreateCredentialParams) (Credential, error) {
	row := q.queryRow(ctx, q.createCredentialStmt, createCredential, arg.ID, arg.AccessID, arg.SecretKey)
	var i Credential
	err := row.Scan(&i.ID, &i.AccessID, &i.SecretKey)
	return i, err
}