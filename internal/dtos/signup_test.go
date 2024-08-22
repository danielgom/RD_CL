package dtos

import (
	"RD-Clone-NAPI/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBuildRegisterResponse(t *testing.T) {
	t.Parallel()

	currentTime := time.Now()

	want := &RegisterResponse{
		Username:  "daniel",
		Email:     "dga_355@hotmail.com",
		CreatedAt: currentTime,
		Enabled:   false,
	}

	response := BuildRegisterResponse(&models.User{
		ID:        1,
		Username:  "daniel",
		Email:     "dga_355@hotmail.com",
		CreatedAt: currentTime,
		Enabled:   false,
	})

	require.Equal(t, want, response)
}
