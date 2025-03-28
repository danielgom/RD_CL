package dtos

import (
	"testing"
	"time"

	"RD-Clone-NAPI/internal/models"
	"github.com/stretchr/testify/require"
)

func TestBuildRegisterResponse(t *testing.T) {
	t.Parallel()

	currentTime := time.Now()

	want := &RegisterResponse{
		Name:      "Daniel",
		LastName:  "Gomez",
		Email:     "dga_355@hotmail.com",
		CreatedAt: currentTime,
		Enabled:   0,
	}

	response := BuildRegisterResponse(&models.User{
		ID:        1,
		Name:      "Daniel",
		LastName:  "Gomez",
		Email:     "dga_355@hotmail.com",
		CreatedAt: currentTime,
		Enabled:   0,
	})

	require.Equal(t, want, response)
}
