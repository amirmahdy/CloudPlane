// go:build integration

package db

import (
	"cloudplane/internal"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createFakeDBUser() (params CreateUserParams, err error) {
	username := internal.CreateRandomName()
	email := internal.CreateRandomEmail()
	fullName := internal.CreateRandomName()
	password := internal.CreateRandomString(8)
	hash, _ := internal.CreateHashPassword(password)

	params = CreateUserParams{
		Username:       username,
		FullName:       fullName,
		Email:          email,
		HashedPassword: hash,
	}

	// Call the create user function
	_, err = testStore.CreateUser(context.Background(), params)
	return
}

func TestCreateAndGetUser(t *testing.T) {
	params, err := createFakeDBUser()
	require.NoError(t, err)

	hashPass, err := testStore.GetUserPassword(context.Background(), params.Username)
	require.NoError(t, err)
	require.Equal(t, hashPass, params.HashedPassword)
}
