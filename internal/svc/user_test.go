package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"RD-Clone-NAPI/internal/dtos"
	"RD-Clone-NAPI/internal/models"
	"RD-Clone-NAPI/internal/security"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var errCommon = errors.New("common error")

type userSuite struct {
	svcSuite
}

func TestUser(t *testing.T) {
	suite.Run(t, &userSuite{})
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

	res, err := u.userService.SignUp(ctx, &req)
	u.NoErrorf(err, "failed to sign up")

	u.Equalf(res.Name, req.Name, "name should be the same")
	u.Equalf(res.Email, req.Email, "email should be the same")
	u.Equalf(res.LastName, req.LastName, "last name should be the same")
	u.Equalf(int8(0), res.Enabled, "user should be disabled")
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
		{
			name: "failure on user save",
			req: dtos.RegisterRequest{
				Name:     "Test",
				LastName: "User",
				Password: "SecurePass123!",
				Email:    "test@example.com",
			},
			expectErr: true,
			mockFunc: func(ctx context.Context) {
				u.userRepository.EXPECT().Save(ctx, gomock.Any()).Return(nil, errCommon)
			},
		},
		{
			name: "failure on token save",
			req: dtos.RegisterRequest{
				Name:     "Test",
				LastName: "User",
				Password: "SecurePass123!",
				Email:    "test@example.com",
			},
			expectErr: true,
			mockFunc: func(ctx context.Context) {
				expectedUser := models.User{
					ID:        2,
					Name:      "Test",
					LastName:  "User",
					Password:  "SecurePass123!",
					Email:     "test@example.com",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Enabled:   0,
				}
				u.userRepository.EXPECT().Save(ctx, gomock.Any()).Return(&expectedUser, nil)
				u.tokenRepository.EXPECT().Save(ctx, gomock.Any()).Return(errCommon)
			},
		},
	}

	for _, tc := range cases {
		u.Run(tc.name, func() {
			ctx := context.TODO()
			tc.mockFunc(ctx)

			res, err := u.userService.SignUp(ctx, &tc.req)
			if tc.expectErr {
				u.Errorf(err, "error expected")

				return
			}

			u.NoErrorf(err, "error not expected")
			u.Equalf(res.Name, tc.req.Name, "name should be the same")
			u.Equalf(res.Email, tc.req.Email, "email should be the same")
			u.Equalf(res.LastName, tc.req.LastName, "last name should be the same")
			u.Equalf(int8(0), res.Enabled, "user should be disabled")
		})
	}
}

func (u *userSuite) TestGet() {
	ctx := context.TODO()
	username := "daniel.gomez"

	expectedUser := &models.User{
		ID:        1,
		Name:      "Daniel",
		LastName:  "Gomez",
		Email:     "dga_355@hotmail.com",
		Password:  "hashed_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Enabled:   1,
	}

	u.userRepository.EXPECT().FindByUsername(ctx, username).Return(expectedUser, nil)

	response, err := u.userService.Get(ctx, username)
	u.NoErrorf(err, "failed to get user")
	u.NotNil(response, "response should not be nil")
	u.Equal(expectedUser.ID, response.ID, "ID should match")
	u.Equal(expectedUser.Name, response.Name, "name should match")
	u.Equal(expectedUser.LastName, response.LastName, "last name should match")
	u.Equal(expectedUser.Email, response.Email, "email should match")
	u.Equal(expectedUser.Enabled, response.Enabled, "enabled status should match")
}

func (u *userSuite) TestGetUserNotFound() {
	ctx := context.TODO()
	username := "nonexistent"

	u.userRepository.EXPECT().FindByUsername(ctx, username).Return(nil, errCommon)

	response, err := u.userService.Get(ctx, username)
	u.Error(err, "error expected when user not found")
	u.Nil(response, "response should be nil")
	u.Contains(err.Error(), "failed to find user", "error message should indicate user not found")
}

func (u *userSuite) TestVerifyAccount() {
	ctx := context.TODO()
	token := "verification-token"

	user := &models.User{
		ID:        1,
		Name:      "Daniel",
		LastName:  "Gomez",
		Email:     "dga_355@hotmail.com",
		Enabled:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	verificationToken := &models.VerificationToken{
		Token:      token,
		User:       user,
		ExpiryDate: time.Now().Add(24 * time.Hour),
	}

	u.tokenRepository.EXPECT().FindByToken(ctx, token).Return(verificationToken, nil)
	u.userRepository.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, updatedUser *models.User) error {
		u.Equal(int8(1), updatedUser.Enabled, "user should be enabled")

		return nil
	})

	err := u.userService.VerifyAccount(ctx, token)
	u.NoErrorf(err, "account verification should succeed")
}

func (u *userSuite) TestVerifyAccountFailures() {
	cases := []struct {
		name     string
		token    string
		mockFunc func(ctx context.Context, token string)
	}{
		{
			name:  "token not found",
			token: "invalid-token",
			mockFunc: func(ctx context.Context, token string) {
				u.tokenRepository.EXPECT().FindByToken(ctx, token).Return(nil, errCommon)
			},
		},
		{
			name:  "update user fails",
			token: "valid-token",
			mockFunc: func(ctx context.Context, token string) {
				user := &models.User{
					ID:       1,
					Name:     "Daniel",
					LastName: "Gomez",
					Email:    "dga_355@hotmail.com",
					Enabled:  0,
				}

				verificationToken := &models.VerificationToken{
					Token:      token,
					User:       user,
					ExpiryDate: time.Now().Add(24 * time.Hour),
				}

				u.tokenRepository.EXPECT().FindByToken(ctx, token).Return(verificationToken, nil)
				u.userRepository.EXPECT().Update(ctx, gomock.Any()).Return(errCommon)
			},
		},
	}

	for _, tc := range cases {
		u.Run(tc.name, func() {
			ctx := context.TODO()
			tc.mockFunc(ctx, tc.token)
			err := u.userService.VerifyAccount(ctx, tc.token)
			u.Error(err, "error expected for unsuccessful verification")
		})
	}
}

func (u *userSuite) TestLogin() {
	ctx := context.TODO()
	loginReq := &dtos.LoginRequest{
		UserOrEmail: "dga_355@hotmail.com",
		Password:    "Password1234@@",
	}

	hash, err := security.Hash("Password1234@@")
	u.NoErrorf(err, "failed to hash password")

	user := &models.User{
		ID:       1,
		Name:     "Daniel",
		LastName: "Gomez",
		Email:    "dga_355@hotmail.com",
		Password: hash,
		Enabled:  1,
	}

	refreshToken := "new-refresh-token"

	u.userRepository.EXPECT().FindByEmail(ctx, loginReq.UserOrEmail).Return(user, nil)
	u.mockRefreshTokenService.EXPECT().Create(ctx).Return(refreshToken, nil)

	response, err := u.userService.Login(ctx, loginReq)
	u.NoErrorf(err, "login should succeed")
	u.NotNil(response, "response should not be nil")
	u.Equal(user.Email, response.Email, "email should match")
	u.Equal(user.Email, response.Username, "username should match")
	u.NotEmpty(response.Token, "JWT token should not be empty")
	u.Equal(refreshToken, response.RefreshToken, "refresh token should match")
	u.NotZero(response.ExpiresAt, "expiration time should not be zero")
}

func (u *userSuite) TestLoginFailures() {
	cases := []struct {
		name     string
		request  *dtos.LoginRequest
		mockFunc func(ctx context.Context, req *dtos.LoginRequest)
	}{
		{
			name: "user not found by username",
			request: &dtos.LoginRequest{
				UserOrEmail: "nonexistent",
				Password:    "password",
			},
			mockFunc: func(ctx context.Context, req *dtos.LoginRequest) {
				u.userRepository.EXPECT().FindByUsername(ctx, req.UserOrEmail).Return(nil, errCommon)
			},
		},
		{
			name: "user not found by email",
			request: &dtos.LoginRequest{
				UserOrEmail: "nonexistent@example.com",
				Password:    "password",
			},
			mockFunc: func(ctx context.Context, req *dtos.LoginRequest) {
				u.userRepository.EXPECT().FindByEmail(ctx, req.UserOrEmail).Return(nil, errCommon)
			},
		},
		{
			name: "incorrect password",
			request: &dtos.LoginRequest{
				UserOrEmail: "dga_355@hotmail.com",
				Password:    "WrongPassword123",
			},
			mockFunc: func(ctx context.Context, req *dtos.LoginRequest) {
				user := &models.User{
					Email:    "dga_355@hotmail.com",
					Password: "$2a$10$some.hashed.password", // Doesn't match the password provided
				}
				u.userRepository.EXPECT().FindByEmail(ctx, req.UserOrEmail).Return(user, nil)
			},
		},
		{
			name: "refresh token creation fails",
			request: &dtos.LoginRequest{
				UserOrEmail: "dga_355@hotmail.com",
				Password:    "Password1234@@",
			},
			mockFunc: func(ctx context.Context, req *dtos.LoginRequest) {
				hash, err := security.Hash("Password1234@@")
				u.NoErrorf(err, "failed to hash password")

				user := &models.User{
					Email:    "dga_355@hotmail.com",
					Password: hash,
				}
				u.userRepository.EXPECT().FindByEmail(ctx, req.UserOrEmail).Return(user, nil)
				u.mockRefreshTokenService.EXPECT().Create(ctx).Return("", errCommon)
			},
		},
	}

	for _, tc := range cases {
		u.Run(tc.name, func() {
			ctx := context.TODO()
			tc.mockFunc(ctx, tc.request)
			response, err := u.userService.Login(ctx, tc.request)
			u.Error(err, "error expected for login failure")
			u.Nil(response, "response should be nil")
		})
	}
}

func (u *userSuite) TestRefreshToken() {
	ctx := context.TODO()
	refreshTokenReq := &dtos.RefreshTokenRequest{
		Username:     "dga_355@hotmail.com",
		RefreshToken: "current-refresh-token",
	}

	newRefreshToken := "new-refresh-token"

	u.mockRefreshTokenService.EXPECT().Validate(ctx, refreshTokenReq.RefreshToken).Return(nil)
	u.mockRefreshTokenService.EXPECT().Create(ctx).Return(newRefreshToken, nil)

	response, err := u.userService.RefreshToken(ctx, refreshTokenReq)
	u.NoErrorf(err, "token refresh should succeed")
	u.NotNil(response, "response should not be nil")
	u.Equal(refreshTokenReq.Username, response.Username, "username should match")
	u.NotEmpty(response.Token, "JWT token should not be empty")
	u.Equal(newRefreshToken, response.RefreshToken, "refresh token should be updated")
	u.NotZero(response.ExpiresAt, "expiration time should not be zero")
}

func (u *userSuite) TestRefreshTokenFailures() {
	cases := []struct {
		name     string
		request  *dtos.RefreshTokenRequest
		mockFunc func(ctx context.Context, req *dtos.RefreshTokenRequest)
	}{
		{
			name: "invalid refresh token",
			request: &dtos.RefreshTokenRequest{
				Username:     "dga_355@hotmail.com",
				RefreshToken: "invalid-token",
			},
			mockFunc: func(ctx context.Context, req *dtos.RefreshTokenRequest) {
				u.mockRefreshTokenService.EXPECT().Validate(ctx, req.RefreshToken).Return(errCommon)
			},
		},
		{
			name: "refresh token creation fails",
			request: &dtos.RefreshTokenRequest{
				Username:     "dga_355@hotmail.com",
				RefreshToken: "valid-token",
			},
			mockFunc: func(ctx context.Context, req *dtos.RefreshTokenRequest) {
				u.mockRefreshTokenService.EXPECT().Validate(ctx, req.RefreshToken).Return(nil)
				u.mockRefreshTokenService.EXPECT().Create(ctx).Return("", errCommon)
			},
		},
	}

	for _, tc := range cases {
		u.Run(tc.name, func() {
			ctx := context.TODO()
			tc.mockFunc(ctx, tc.request)
			response, err := u.userService.RefreshToken(ctx, tc.request)
			u.Error(err, "error expected for refresh token failure")
			u.Nil(response, "response should be nil")
		})
	}
}
