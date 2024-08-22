package services

import (
	"RD-Clone-NAPI/internal/config"
	"RD-Clone-NAPI/internal/db"
	"RD-Clone-NAPI/internal/dtos"
	"context"
	"log"
)

type ServiceFactory struct {
	UserService         UserService
	RefreshTokenService RefreshTokenService
}

func NewFactory() *ServiceFactory {
	c := config.Load()
	dbPool, err := config.NewDB(c)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := db.NewUserRepository(dbPool)
	tokenRepository := db.NewTokenRepository(dbPool)
	refreshTokenRepository := db.NewRTRepository(dbPool)

	refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
	userService := NewUserService(userRepository, tokenRepository, refreshTokenService)

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
