package services

import (
	"RD-Clone-NAPI/internal/db"
	"RD-Clone-NAPI/internal/dtos"
	"RD-Clone-NAPI/internal/models"
	"RD-Clone-NAPI/internal/security"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/mail"
	"time"
)

const (
	verificationTokenExpiration = 24
)

var (
	errFailedPasswordVerification = errors.New("invalid password")
)

type userSvc struct {
	userDB  db.UserRepository
	tokenDB db.TokenRepository
	rtSvc   RefreshTokenService
}

// NewUserService returns a new instance of user service.
func NewUserService(uR db.UserRepository, tR db.TokenRepository, rTSvc RefreshTokenService) UserService {
	return &userSvc{userDB: uR, tokenDB: tR, rtSvc: rTSvc}
}

// SignUp executes core logic in order to save the user and generate its verification token for the first time.
func (u *userSvc) SignUp(ctx context.Context, req *dtos.RegisterRequest,
) (*dtos.RegisterResponse, error) {
	pass, err := security.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	currentTime := time.Now().Local()
	user := &models.User{
		Username:  req.Username,
		Password:  pass,
		Email:     req.Email,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	user, saveErr := u.userDB.Save(ctx, user)
	if saveErr != nil {
		return nil, saveErr
	}

	_, err = u.generateVerificationToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate verification token: %w", err)
	}

	return dtos.BuildRegisterResponse(user), nil
}

func (u *userSvc) Get(ctx context.Context, username string) (*dtos.UserResponse, error) {
	user, commonError := u.userDB.FindByUsername(ctx, username)
	if commonError != nil {
		return nil, commonError
	}

	return &dtos.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Enabled:   user.Enabled,
	}, nil
}

// VerifyAccount verifies the account.
func (u *userSvc) VerifyAccount(ctx context.Context, tStr string) error {
	token, verErr := u.tokenDB.FindByToken(ctx, tStr)
	if verErr != nil {
		return verErr
	}

	token.User.Enabled = true
	token.User.UpdatedAt = time.Now()

	updateErr := u.userDB.Update(ctx, token.User)

	if updateErr != nil {
		return updateErr
	}

	return nil
}

// Login validates username/email and password returning a JWT token and a refresh token with expiration.
func (u *userSvc) Login(ctx context.Context, loginReq *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	var user *models.User

	_, err := mail.ParseAddress(loginReq.UserOrEmail)
	if err != nil {
		user, err = u.userDB.FindByUsername(ctx, loginReq.UserOrEmail)
	} else {
		user, err = u.userDB.FindByEmail(ctx, loginReq.UserOrEmail)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	validPass := security.CheckHash(loginReq.Password, user.Password)
	if !validPass {
		return nil, errFailedPasswordVerification
	}

	JWT, expDate, err := security.GenerateTokenWithExp(user.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token while login: %w", err)
	}

	refreshToken, err := u.rtSvc.Create(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	return dtos.BuildLoginResponse(user.Username, user.Email, JWT, refreshToken, expDate), nil
}

func (u *userSvc) RefreshToken(ctx context.Context, rtReq *dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse,
	error) {
	err := u.rtSvc.Validate(ctx, rtReq.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to validate refresh token: %w", err)
	}

	JWT, expDate, err := security.GenerateTokenWithExp(rtReq.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	refreshToken, err := u.rtSvc.Create(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new refresh token: %w", err)
	}

	return &dtos.RefreshTokenResponse{
		Username:     rtReq.Username,
		Token:        JWT,
		RefreshToken: refreshToken,
		ExpiresAt:    expDate,
	}, nil
}

func (u *userSvc) generateVerificationToken(ctx context.Context, user *models.User) (string, error) {
	token := uuid.New().String()
	vToken := models.VerificationToken{
		Token:      uuid.New().String(),
		User:       user,
		ExpiryDate: time.Now().Add(time.Hour * verificationTokenExpiration),
	}

	err := u.tokenDB.Save(ctx, &vToken)
	if err != nil {
		return "", fmt.Errorf("failed to save verification token: %w", err)
	}

	return token, nil
}
