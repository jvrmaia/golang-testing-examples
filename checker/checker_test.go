package checker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	healthy = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	}))
	unhealthy = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "FAIL")
	}))
)

type CheckerTestSuite struct {
	suite.Suite

	healthyServer   string
	unhealthyServer string
}

func (suite *CheckerTestSuite) SetupSuite() {
	suite.healthyServer = "http://www.google.com"
	suite.unhealthyServer = unhealthy.URL
}

func (suite *CheckerTestSuite) TearDownSuite() {
	healthy.Close()
	unhealthy.Close()
}

func (suite *CheckerTestSuite) TestHealthyResource() {
	nums := []int{1,2,3,4,5,6}
	for _, t := range nums {
		fmt.Println(t)
		var c RestChecker
		c.Configure(suite.healthyServer, 1000)
		assert.NoError(suite.T(), c.Check())
	}
}

func (suite *CheckerTestSuite) TestUnhealthyResource() {
	nums := []int{1,2,3,4,5,6}
	for _, t := range nums {
		fmt.Println(t)
		var c RestChecker
		c.Configure(suite.unhealthyServer, 1000)
		assert.Error(suite.T(), c.Check())
	}
}

func (suite *CheckerTestSuite) TestMissingResource() {
	nums := []int{1,2,3,4,5,6}
	for _, t := range nums {
		fmt.Println(t)
		var c RestChecker
		c.Configure("http://localhost:12345", 1000)
		assert.Error(suite.T(), c.Check())
	}
}

func TestChecker(t *testing.T) {
	suite.Run(t, new(CheckerTestSuite))
}
