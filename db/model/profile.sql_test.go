// go:build integration

package db

import (
	"cloudplane/internal"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	userRes, err := createFakeDBUser()
	require.NoError(t, err)

	credParam := CreateCredentialParams{
		ID:        uuid.New(),
		AccessID:  internal.CreateRandomString(10),
		SecretKey: internal.CreateRandomString(10),
	}
	credRes, err := testStore.CreateCredential(context.Background(), credParam)
	require.NoError(t, err)

	profParam := CreateProfileParams{
		ID:          uuid.New(),
		Description: internal.CreateRandomString(10),
		Region:      internal.CreateRandomString(5),
		CredID:      credRes.ID,
		Username:    userRes.Username,
	}
	profRes, err := testStore.CreateProfile(context.Background(), profParam)
	require.NoError(t, err)
	require.Equal(t, profRes.CredID, credRes.ID)

	// test to check if we can retrieve the same profile using GetProfiles
	getProfParam := GetProfilesParams{
		Username: userRes.Username,
		Limit:    1,
	}
	gotProfRes, err := testStore.GetProfiles(context.Background(), getProfParam)
	require.NoError(t, err)
	require.Equal(t, gotProfRes[0].AccessID, credRes.AccessID)
	require.Equal(t, gotProfRes[0].SecretKey, credRes.SecretKey)
	require.Equal(t, gotProfRes[0].Region, profRes.Region)
}
