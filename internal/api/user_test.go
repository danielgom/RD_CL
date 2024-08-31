package api

import (
	"net/http"
	"testing"

	"RD-Clone-NAPI/internal/dtos"
	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/suite"
)

type userSuite struct {
	apiSuite
}

func TestUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, &userSuite{apiSuite{
		dbName: "user_api_test",
	}})
}

//nolint:funlen // Testing function length
func (u *userSuite) TestUser() {
	testCases := []struct {
		name       string
		register   dtos.RegisterRequest
		expected   string
		statusCode int
	}{
		{
			name: "successful sign up",
			register: dtos.RegisterRequest{
				Name:     "Daniel",
				LastName: "Gomez",
				Password: "Goodpassword1234@@",
				Email:    "dga_355@hotmail.com",
			},
			expected: `{
				"name": "Daniel",
				"last_name": "Gomez",
				"email": "dga_355@hotmail.com",
				"created_at": "$created_at",
				"enabled": 0
			}`,
			statusCode: http.StatusCreated,
		},
		{
			name: "missing name",
			register: dtos.RegisterRequest{
				LastName: "Gomez",
				Password: "Goodpassword1234@@",
				Email:    "dga_355@hotmail.com",
			},
			expected: `{
				"request_id": "$request_id",
				"error": "invalid register request",
    			"reason": "failed to validate request body: Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
			}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			register: dtos.RegisterRequest{
				Name:     "Daniel",
				LastName: "Gomez",
				Password: "Goodpassword1234@@",
				Email:    "invalid_email",
			},
			expected: `{
				"request_id": "$request_id",
				"error": "invalid register request",
    			"reason": "failed to validate request body: Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
			}`,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		u.Run(tc.name, func() {
			w := u.Post("/v1/user/signup", u.toReader(tc.register))
			writerResult := w.Result()
			defer func() {
				err := writerResult.Body.Close()
				u.Nilf(err, "failed to close body")
			}()

			u.Equalf(tc.statusCode, writerResult.StatusCode, "Status code should be %v", tc.statusCode)
			u.jsonEq(writerResult.Body, tc.expected,
				td.Tag("request_id", td.NotZero()),
				td.Tag("created_at", td.NotZero()))
		})
	}
}

func (u *userSuite) TestUserV2() {
	u.PostEq("/v1/user/signup",
		`{
				"name":"Daniel",
				"last_name":"Gierre",
				"password":"Goodpassword1234@@",
				"email":"new_email@hotmail.com"}`, `{
			"name": "Daniel",
			"last_name": "Gierre",
			"email": "new_email@hotmail.com",
			"created_at": "$created_at",
			"enabled": 0
	}`, td.Tag("created_at", td.NotZero()))
}
