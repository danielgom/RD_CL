package services

import (
	"context"
	"testing"

	"RD-Clone-NAPI/internal/config"
	"RD-Clone-NAPI/internal/dtos"
	"RD-Clone-NAPI/internal/testutils"
	"github.com/stretchr/testify/suite"
)

type userSuite struct {
	suite.Suite

	pgCont testutils.Container

	svc UserService
}

func TestUser(t *testing.T) {
	suite.Run(t, &userSuite{})
}

func (u *userSuite) SetupSuite() {
	u.pgCont = testutils.CreatePGContainer()
	config.InitialiseTest(u.pgCont.ConnectionString(), "user_service_test")

	serviceFactory := NewFactory()
	u.svc = serviceFactory.UserService
}

func (u *userSuite) TestUserSignup() {
	ctx := context.TODO()
	req := dtos.RegisterRequest{
		Name:     "Daniel",
		LastName: "Gomez",
		Password: "Password1234@@",
		Email:    "dga_355@hotmail.com",
	}
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
	}{
		{
			name: "successful sign up",
			req: dtos.RegisterRequest{
				Name:     "Daniel",
				LastName: "Gomez",
				Password: "Password1234@@",
				Email:    "dga_355_2@hotmail.com",
			},
			expectErr: false,
		},
	}

	for _, tc := range cases {
		u.Run(tc.name, func() {
			ctx := context.TODO()
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
