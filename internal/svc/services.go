package services

import (
	"context"

	"RD-Clone-NAPI/internal/dtos"
)

type ServiceFactory struct {
	UserService         UserService
	RefreshTokenService RefreshTokenService
}

func NewFactory(userService UserService, refreshTokenService RefreshTokenService) *ServiceFactory {
	if userService == nil {
		panic("userService is required")
	}

	if refreshTokenService == nil {
		panic("refreshTokenService is required")
	}

	return &ServiceFactory{
		UserService:         userService,
		RefreshTokenService: refreshTokenService,
	}
}

// UserService contains all the business logic for the user.
type UserService interface {
	SignUp(context.Context, *dtos.RegisterRequest) (*dtos.RegisterResponse, error)
	Get(context.Context, string) (*dtos.UserResponse, error)
	VerifyAccount(context.Context, string) error
	Login(context.Context, *dtos.LoginRequest) (*dtos.LoginResponse, error)
	RefreshToken(context.Context, *dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse, error)
}

// RefreshTokenService contains all the business logic for the RefreshToken service.
type RefreshTokenService interface {
	Create(context.Context) (string, error)
	Validate(context.Context, string) error
}
