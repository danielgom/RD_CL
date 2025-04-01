package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type healthSuite struct {
	apiSuite
}

func TestHealth(t *testing.T) {
	suite.Run(t, &healthSuite{apiSuite{
		dbName: "health_test",
	}})
}

func (c *healthSuite) TestHealth() {
	// This is a way to test a handler, we are explicitly calling the getHealth handler method
	// request is not really required unless we want to pass a modified request to the handler
	// for example a JSON body.
	w := c.Get("/v1/health")

	writerResult := w.Result()
	defer func() {
		err := writerResult.Body.Close()
		c.NoErrorf(err, "failed to close body")
	}()

	healthResponse := Health{
		Environment: "development",
		Healthy:     true,
		Database:    true,
	}

	JSONResponse, err := json.Marshal(healthResponse)
	c.NoErrorf(err, "failed to marshal struct")

	c.Equalf(http.StatusOK, writerResult.StatusCode, "Status code should be 200")
	c.jsonEq(writerResult.Body, string(JSONResponse))
}
