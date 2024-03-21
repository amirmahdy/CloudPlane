package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secret := "changethissecretkeyforproduction"
	jwtMaker, err := NewJWTMaker(secret)
	require.NoError(t, err)
	token, payload, err := jwtMaker.CreateToken("userTest", time.Duration(1*time.Minute))
	require.NoError(t, err)

	payloadVerfied, err := jwtMaker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, payloadVerfied.UserID, payload.UserID)
}
