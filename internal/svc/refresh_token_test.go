package services

import (
	"context"
	"testing"
	"time"

	"RD-Clone-NAPI/internal/models"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type refreshTokenSuite struct {
	svcSuite
}

func TestRefreshToken(t *testing.T) {
	suite.Run(t, &refreshTokenSuite{})
}

func (u *refreshTokenSuite) TestRefreshToken() {
	ctx := context.TODO()
	u.refreshTokenRepository.EXPECT().Save(ctx, gomock.Any()).Return(nil)

	token, err := u.refreshTokenService.Create(context.TODO())
	u.Nilf(err, "failed to create refresh token")
	u.NotZero(token)
}

func (u *refreshTokenSuite) TestRefreshToken_Create_Error() {
	ctx := context.TODO()
	u.refreshTokenRepository.EXPECT().Save(ctx, gomock.Any()).Return(errCommon)

	token, err := u.refreshTokenService.Create(context.TODO())
	u.Error(err, "expected error when saving token fails")
	u.Contains(err.Error(), "error while saving refresh token")
	u.Empty(token, "token should be empty when error occurs")
}

func (u *refreshTokenSuite) TestRefreshToken_Validate_Success() {
	ctx := context.TODO()
	token := "valid-token"

	// Create a valid token that hasn't expired yet
	refreshToken := &models.RefreshToken{
		Token:     token,
		ExpiresAt: time.Now().Add(1 * time.Hour), // expires in the future
	}

	u.refreshTokenRepository.EXPECT().FindByToken(ctx, token).Return(refreshToken, nil)

	err := u.refreshTokenService.Validate(ctx, token)
	u.Nil(err, "validation should succeed for non-expired token")
}

func (u *refreshTokenSuite) TestRefreshToken_Validate_TokenExpired() {
	ctx := context.TODO()
	token := "expired-token"

	// Create an expired token
	refreshToken := &models.RefreshToken{
		Token:     token,
		ExpiresAt: time.Now().Add(-1 * time.Hour), // expired 1 hour ago
	}

	u.refreshTokenRepository.EXPECT().FindByToken(ctx, token).Return(refreshToken, nil)

	err := u.refreshTokenService.Validate(ctx, token)
	u.Error(err, "validation should fail for expired token")
	u.Equal("token expired", err.Error(), "should return token expired error")
}

func (u *refreshTokenSuite) TestRefreshToken_Validate_TokenNotFound() {
	ctx := context.TODO()
	token := "non-existent-token"

	u.refreshTokenRepository.EXPECT().FindByToken(ctx, token).Return(nil, errCommon)

	err := u.refreshTokenService.Validate(ctx, token)
	u.Error(err, "validation should fail when token not found")
	u.Contains(err.Error(), "error while finding the current token")
}

func (u *refreshTokenSuite) TestRefreshToken_Validate_EmptyToken() {
	ctx := context.TODO()
	token := ""

	u.refreshTokenRepository.EXPECT().FindByToken(ctx, token).Return(nil, errCommon)

	err := u.refreshTokenService.Validate(ctx, token)
	u.Error(err, "validation should fail with empty token")
	u.Contains(err.Error(), "error while finding the current token")
}
