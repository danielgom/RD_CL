package dtos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBuildLoginResponse(t *testing.T) {
	t.Parallel()

	expiration := time.Now()
	want := &LoginResponse{
		Username:     "daniel",
		Email:        "dga_355@hotmail.com",
		Token:        "asdsdasd",
		ExpiresAt:    expiration,
		RefreshToken: "rt",
	}

	response := BuildLoginResponse("daniel", "dga_355@hotmail.com", "asdsdasd", "rt", expiration)

	require.Equal(t, want, response)
}
