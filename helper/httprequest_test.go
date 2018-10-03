package helper

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockedServer struct {
	*httptest.Server
	Requests [][]byte
	Response string
	Status   string
}

func NewMockedServer(status string) *MockedServer {
	ser := &MockedServer{}
	ser.Server = httptest.NewServer(ser)
	ser.Requests = [][]byte{}
	ser.Response = "hello"
	ser.Status = status
	return ser
}

func (ser *MockedServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var err error
	lastRequest, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	ser.Requests = append(ser.Requests, lastRequest)

	if ser.Status == "404" {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	resp.Write([]byte(ser.Response))
}

func TestSpec(t *testing.T) {
	Convey("Given the http server which will return 200 status code", t, func() {
		mockedServer := NewMockedServer("200")

		Convey("Fetching the http server", func() {
			content := GetHttpContent(mockedServer.URL)

			Convey("Should get the content from the url", func() {
				So(len(content), ShouldBeGreaterThan, 0)
			})
		})
	})

	Convey("Given the http server which will return 404 status code", t, func() {
		mockedServer := NewMockedServer("404")

		Convey("Fetching the http server", func() {
			content := GetHttpContent(mockedServer.URL)

			Convey("Should not get any content since it is 404", func() {
				So(len(content), ShouldEqual, 0)
			})
		})
	})
}
