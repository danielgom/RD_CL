package services

import (
	"RD-Clone-NAPI/internal/svc/mock_repository"
	"RD-Clone-NAPI/internal/svc/mock_service"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type svcSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller

	userService         UserService
	refreshTokenService RefreshTokenService

	userRepository          *mock_repository.MockUserRepository
	tokenRepository         *mock_repository.MockTokenRepository
	refreshTokenRepository  *mock_repository.MockRefreshTokenRepository
	mockRefreshTokenService *mock_service.MockRefreshTokenService
}

func (u *svcSuite) SetupTest() {
	u.mockCtrl = gomock.NewController(u.T())

	u.userRepository = mock_repository.NewMockUserRepository(u.mockCtrl)
	u.tokenRepository = mock_repository.NewMockTokenRepository(u.mockCtrl)
	u.refreshTokenRepository = mock_repository.NewMockRefreshTokenRepository(u.mockCtrl)
	u.mockRefreshTokenService = mock_service.NewMockRefreshTokenService(u.mockCtrl)

	u.userService = NewUserService(u.userRepository, u.tokenRepository, u.mockRefreshTokenService)
	u.refreshTokenService = NewRefreshTokenService(u.refreshTokenRepository)
}

func (u *svcSuite) TearDownTest() {
	u.mockCtrl.Finish()
}
