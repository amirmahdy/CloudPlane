package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateProfileTXParam struct {
	Region      string `json:"region"`
	AccessID    string `json:"access_id"`
	SecretKey   string `json:"secret_key"`
	Username    string `json:"username"`
	Description string `json:"description"`
}
type CreateProfileTXResultType struct {
	ProfID uuid.UUID `json:"profile_id"`
}

func (store *SqlStore) CreateProfileTX(t CreateProfileTXParam) (CreateProfileTXResultType, error) {
	res := CreateProfileTXResultType{}
	err := store.execTX(func(q *Queries) error {
		credParam := CreateCredentialParams{
			ID:        uuid.New(),
			AccessID:  t.AccessID,
			SecretKey: t.SecretKey,
		}
		credRes, err := q.CreateCredential(context.Background(), credParam)
		if err != nil {
			return err
		}
		profParam := CreateProfileParams{
			ID:          uuid.New(),
			Description: t.Description,
			Region:      t.Region,
			CredID:      credRes.ID,
			Username:    t.Username,
		}
		profRes, err := q.CreateProfile(context.Background(), profParam)
		if err != nil {
			return err
		}

		res.ProfID = profRes.ID
		return nil
	})
	return res, err
}
