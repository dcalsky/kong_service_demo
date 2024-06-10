package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := "123456"
	hash, err := hashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	require.True(t, checkPasswordHash(password, hash))
}
