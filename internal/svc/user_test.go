package services

import (
	"RD-Clone-NAPI/internal/models"
	"RD-Clone-NAPI/internal/svc/mock_repository"
	"RD-Clone-NAPI/internal/svc/mock_service"
	"context"
	"go.uber.org/mock/gomock"
	"testing"
	"time"

	"RD-Clone-NAPI/internal/dtos"
	"github.com/stretchr/testify/suite"
)

type userSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller

	svc UserService

	userRepository  *mock_repository.MockUserRepository
	tokenRepository *mock_repository.MockTokenRepository
	rTSvc           *mock_service.MockRefreshTokenService
}

func TestUser(t *testing.T) {
	suite.Run(t, &userSuite{})
}

func (u *userSuite) SetupSuite() {
	u.mockCtrl = gomock.NewController(u.T())

	u.userRepository = mock_repository.NewMockUserRepository(u.mockCtrl)
	u.tokenRepository = mock_repository.NewMockTokenRepository(u.mockCtrl)
	u.rTSvc = mock_service.NewMockRefreshTokenService(u.mockCtrl)

	u.svc = NewUserService(u.userRepository, u.tokenRepository, u.rTSvc)

	//u.pgCont = testutils.CreatePGContainer()
	//config.InitialiseTest(u.pgCont.ConnectionString(), "user_service_test")

	//serviceFactory := NewFactory()
	//u.svc = serviceFactory.UserService
}

func (u *userSuite) TearDownSuite() {
	u.mockCtrl.Finish()
}

func (u *userSuite) TestUserSignup() {
	ctx := context.TODO()
	req := dtos.RegisterRequest{
		Name:     "Daniel",
		LastName: "Gomez",
		Password: "Password1234@@",
		Email:    "dga_355@hotmail.com",
	}

	expectedUser := models.User{
		ID:        1,
		Name:      "Daniel",
		LastName:  "Gomez",
		Password:  "Password1234@@",
		Email:     "dga_355@hotmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Enabled:   0,
	}

	u.userRepository.EXPECT().Save(ctx, gomock.Any()).Return(&expectedUser, nil)
	u.tokenRepository.EXPECT().Save(ctx, gomock.Any()).Return(nil)

	res, err := u.svc.SignUp(ctx, &req)
	u.Nilf(err, "failed to sign up")

	u.Equalf(res.Name, req.Name, "name should be the same")
	u.Equalf(res.Email, req.Email, "email should be the same")
	u.Equalf(res.LastName, req.LastName, "last name should be the same")
	u.Equalf(res.Enabled, int8(0), "user should be disabled")
}

func (u *userSuite) TestUserSignupTableDriven() {
	cases := []struct {
		name      string
		req       dtos.RegisterRequest
		expectErr bool
		mockFunc  func(ctx context.Context)
	}{
		{
			name: "successful sign up",
			req: dtos.RegisterRequest{
				Name:     "Daniel",
				LastName: "Gomez",
				Password: "Password1234@@",
				Email:    "dga_355@hotmail.com",
			},
			expectErr: false,
			mockFunc: func(ctx context.Context) {
				expectedUser := models.User{
					ID:        1,
					Name:      "Daniel",
					LastName:  "Gomez",
					Password:  "Password1234@@",
					Email:     "dga_355@hotmail.com",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Enabled:   0,
				}
				u.userRepository.EXPECT().Save(ctx, gomock.Any()).Return(&expectedUser, nil)
				u.tokenRepository.EXPECT().Save(ctx, gomock.Any()).Return(nil)
			},
		},
	}

	for _, tc := range cases {
		u.Run(tc.name, func() {
			ctx := context.TODO()
			tc.mockFunc(ctx)
			res, err := u.svc.SignUp(ctx, &tc.req)
			if tc.expectErr {
				u.NotNilf(err, "error expected")
				return
			}
			u.Nilf(err, "error not expected")
			u.Equalf(res.Name, tc.req.Name, "name should be the same")
			u.Equalf(res.Email, tc.req.Email, "email should be the same")
			u.Equalf(res.LastName, tc.req.LastName, "last name should be the same")
			u.Equalf(res.Enabled, int8(0), "user should be disabled")
		})
	}
}
