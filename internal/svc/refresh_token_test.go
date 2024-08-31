package services

import (
	"context"
	"testing"

	"RD-Clone-NAPI/internal/config"
	"RD-Clone-NAPI/internal/testutils"
	"github.com/stretchr/testify/suite"
)

type refreshTokenSuite struct {
	suite.Suite

	pgCont testutils.Container

	svc RefreshTokenService
}

func TestRefreshToken(t *testing.T) {
	suite.Run(t, &refreshTokenSuite{})
}

func (u *refreshTokenSuite) SetupSuite() {
	u.pgCont = testutils.CreatePGContainer()
	config.InitialiseTest(u.pgCont.ConnectionString(), "refresh_token_service_test")

	serviceFactory := NewFactory()
	u.svc = serviceFactory.RefreshTokenService
}

func (u *refreshTokenSuite) TestRefreshToken() {
	token, err := u.svc.Create(context.TODO())
	u.Nilf(err, "failed to create refresh token")
	u.NotZero(token)
}
