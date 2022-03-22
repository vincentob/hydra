package httpclient_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/vincentob/hydra/httpclient"
)

const (
	URLPathGETJSON = "/api/v1/json"
)

// Case the testing case struct used for http client test
type Case struct {
	Name                string
	Url                 string
	Body                interface{}
	ExpectedErr         error
	ExpectedResponseNil bool
	ExpectedStatusCode  int
	ExpectedDataCount   int
}

// JSONResult the data for response
type JSONResult struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Data   []Data `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// HttpClientTestSuite extends base suite
type HttpClientTestSuite struct {
	suite.Suite

	mockServer           *httptest.Server
	mockGinServer        *gin.Engine
	mockGinServerAddress string
}

// SetupSuite
func (s *HttpClientTestSuite) SetupSuite() {
	logrus.Debug("http suite setup")

	s.NewMockServer()

	s.NewMockGinServer()
}

// TestSuite start suite test cases.
func TestHttpClient(t *testing.T) {
	suite.Run(t, new(HttpClientTestSuite))
}

// NewMockServer return a mock gin server that define my mock action.
func (s *HttpClientTestSuite) NewMockGinServer() {
	s.T().Helper()

	gin.SetMode(gin.ReleaseMode)

	s.mockGinServer = gin.New()

	// Register my mock route
	s.mockGinServer.Any(URLPathGETJSON, func(c *gin.Context) {
		c.JSON(http.StatusOK, &JSONResult{
			Msg:    "ok",
			Status: http.StatusOK,
			Data: []Data{
				{Name: "A", Age: 18},
				{Name: "B", Age: 2},
			},
		})
	})

	s.mockGinServerAddress = "http://localhost:8080"
	go func() {
		if err := s.mockGinServer.Run(); err != nil {
			logrus.Fatalf("Mock gin server start failed: %v", err.Error())
		}
	}()

	time.Sleep(time.Second * 1)
}

// NewMockServer return a mock http test server that define my mock action.
func (s *HttpClientTestSuite) NewMockServer() *httptest.Server {
	s.T().Helper()

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == URLPathGETJSON {
			logrus.Debug("Path: response json")
		}

		rw.WriteHeader(http.StatusOK)
	}))
}

// #############  Start  Test  Suite  ################
// TestGetWithJSONResult
func (s *HttpClientTestSuite) TestGetWithJsonResult() {
	cases := []Case{
		{
			Name:                "Request FAILED: Invalid url: unsupported protocol scheme",
			Url:                 "www.google.com",
			ExpectedErr:         errors.New("unsupported protocol scheme"),
			ExpectedResponseNil: true,
		},
		{
			Name:                "Request OK",
			Url:                 s.mockGinServerAddress + URLPathGETJSON,
			ExpectedErr:         nil,
			ExpectedResponseNil: false,
			ExpectedStatusCode:  http.StatusOK,
			ExpectedDataCount:   2,
		},
	}

	for _, c := range cases {
		var result JSONResult
		resp, err := httpclient.DoGetJson(c.Url, nil, &result)

		Convey(c.Name, s.T(), func() {
			Convey("check errors", func() {
				if c.ExpectedErr != nil {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}

			})

			Convey("check response", func() {
				if c.ExpectedResponseNil {
					So(resp, ShouldBeNil)
				} else {
					So(resp, ShouldNotBeNil)
				}
			})

			if !c.ExpectedResponseNil {
				Convey("check response status code", func() {
					So(resp.StatusCode, ShouldEqual, c.ExpectedStatusCode)
				})
				Convey("check response data count", func() {
					So(len(result.Data), ShouldEqual, c.ExpectedDataCount)
				})
			}
		})
	}
}

// TestPostJsonWithResult
func (s *HttpClientTestSuite) TestPostJsonWithResult() {
	cases := []Case{
		{
			Name:                "Request String Body",
			Url:                 s.mockGinServerAddress + URLPathGETJSON,
			Body:                "string body",
			ExpectedErr:         nil,
			ExpectedResponseNil: false,
			ExpectedStatusCode:  http.StatusOK,
			ExpectedDataCount:   2,
		},
		{
			Name:                "Request Struct Body",
			Url:                 s.mockGinServerAddress + URLPathGETJSON,
			Body:                Data{Name: "A", Age: 18},
			ExpectedErr:         nil,
			ExpectedResponseNil: false,
			ExpectedStatusCode:  http.StatusOK,
			ExpectedDataCount:   2,
		},
	}

	for _, c := range cases {
		var result JSONResult
		resp, err := httpclient.DoPostJsonWithResult(c.Url, c.Body, &result)

		Convey(c.Name, s.T(), func() {
			Convey("check errors", func() {
				if c.ExpectedErr != nil {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}

			})

			Convey("check response", func() {
				if c.ExpectedResponseNil {
					So(resp, ShouldBeNil)
				} else {
					So(resp, ShouldNotBeNil)
				}
			})

			if !c.ExpectedResponseNil {
				Convey("check response status code", func() {
					So(resp.StatusCode, ShouldEqual, c.ExpectedStatusCode)
				})
				Convey("check response data count", func() {
					So(len(result.Data), ShouldEqual, c.ExpectedDataCount)
				})
			}
		})
	}
}
