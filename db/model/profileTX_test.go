package db

import (
	"cloudplane/internal"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateProfileTX(t *testing.T) {
	userRes, err := createFakeDBUser()
	require.NoError(t, err)

	profParam := CreateProfileTXParam{
		AccessID:    internal.CreateRandomString(10),
		SecretKey:   internal.CreateRandomString(10),
		Description: internal.CreateRandomString(10),
		Region:      internal.CreateRandomString(5),
		Username:    userRes.Username,
	}
	profRes, err := testStore.CreateProfileTX(profParam)
	require.NoError(t, err)
	require.NotZero(t, profRes.ProfID)

	// test to check if we can retrieve the same profile using GetProfiles
	getProfParam := GetProfilesParams{
		Username: userRes.Username,
		Limit:    1,
	}
	gotProfRes, err := testStore.GetProfiles(context.Background(), getProfParam)
	require.NoError(t, err)
	require.Equal(t, gotProfRes[0].AccessID, profParam.AccessID)
	require.Equal(t, gotProfRes[0].SecretKey, profParam.SecretKey)
	require.Equal(t, gotProfRes[0].Region, profParam.Region)
}
