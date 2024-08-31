package api

import (
	"RD-Clone-NAPI/internal/config"
	"RD-Clone-NAPI/internal/testutils"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	services "RD-Clone-NAPI/internal/svc"

	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/suite"
)

const jwtEx = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJHTy1SZWRkaXQtQ0wiLCJzdWIiOiJkZ2" +
	"FfMzU1QG91dGxvb2suY29tIiwiZXhwIjoxNzI0NzIxNTkyLCJpYXQiOjE3MjQ3MTc5OTJ9" +
	".U9NNd1jg-yGeKUd4PYHss9USMdn-C1TtFbS5uwMGMM7qrtKf_ij7vtpAJXSu87LNQA2d4jkKgl4lDPX0mZ850A"

type apiSuite struct {
	suite.Suite
	deepTester *td.T

	server  *httptest.Server
	handler http.Handler
	dbName  string

	pgCont testutils.Container

	UserHandler *UserHandler
}

func (s *apiSuite) SetupSuite() {
	s.pgCont = testutils.CreatePGContainer()
	config.InitialiseTest(s.pgCont.ConnectionString(), s.dbName)

	api := New()
	s.server = httptest.NewServer(api.Router())
	s.handler = api.Router()

	factory := services.NewFactory()
	s.UserHandler = NewUserHandler(factory.UserService, api)
}

func (s *apiSuite) TearDownSuite() {
	config.CloseDB()
}

func (s *apiSuite) dt() *td.T {
	if s.deepTester == nil {
		s.deepTester = td.NewT(s.T())
	}
	return s.deepTester
}

func (s *apiSuite) request(
	method, path string, body io.Reader,
) *http.Response {
	ctx := context.TODO()
	r, err := http.NewRequestWithContext(ctx, method, s.server.URL+path, body)
	if err != nil {
		s.T().Fatal(err)
		return nil
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		s.T().Fatal(err)
		return nil
	}
	return resp
}

func req(method, path string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, path, body)
	return req
}

func (s *apiSuite) Get(path string) *httptest.ResponseRecorder {
	rs := req("GET", path, nil)
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, rs)
	return w
}

func (s *apiSuite) GetWithJWT(path string) *httptest.ResponseRecorder {
	rs := req("GET", path, nil)
	rs.Header.Set(AuthorizationTokenHeader, "Bearer "+jwtEx)
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, rs)
	return w
}

func (s *apiSuite) Post(path string, body io.Reader) *httptest.ResponseRecorder {
	rs := req("POST", path, body)
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, rs)
	return w
}

func (s *apiSuite) ResponseEq(method, path string, body io.Reader, exp string, params ...any) {
	resp := s.request(method, path, body)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			s.T().Fatal(err)
		}
	}()

	s.jsonEq(resp.Body, exp, params...)
}

func (s *apiSuite) jsonEq(got io.Reader, exp string, params ...any) bool {
	return td.Cmp(s.T(), got, td.Smuggle(json.RawMessage{}, td.JSON(exp, params...)))
}

func (s *apiSuite) PostEq(path, body, exp string, params ...any) {
	s.ResponseEq("POST", path, strings.NewReader(body), exp, params...)
}

func (s *apiSuite) GetEq(path, exp string, params ...any) {
	s.ResponseEq("GET", path, nil, exp, params...)
}

func (s *apiSuite) Anchor(op td.TestDeep, model ...interface{}) interface{} {
	return s.dt().Anchor(op, model...)
}

func (s *apiSuite) DeepCmp(exp, got any, msg string) {
	s.dt().Cmp(got, exp, msg)
}

func (s *apiSuite) NonEmptyString() string {
	return makeString(s.dt().Anchor(td.NotEmpty(), ""))
}

func (s *apiSuite) toReader(obj any) io.Reader {
	jsonBytes, err := json.Marshal(obj)
	s.Nilf(err, "failed to marshal struct")
	return bytes.NewReader(jsonBytes)
}

func makeString(source any) string {
	str, ok := source.(string)
	if !ok {
		log.Fatalf("failed to convert %+v to string", source)
	}
	return str
}
