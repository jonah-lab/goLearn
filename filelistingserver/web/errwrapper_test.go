package main

import (
	"github.com/jonah-lab/goLearn/filelistingserver/filter"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("123")
}

func TestErrwrapper(t *testing.T) {
	tests := []struct {
		h       filter.appHandler
		code    int
		message string
	}{
		{errPanic, 500, ""},
	}
	for _, tt := range tests {
		f := filter.errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expect: %s,%d got:%s,%d", tt.message, tt.code, body, response.Code)
		}
	}
}
