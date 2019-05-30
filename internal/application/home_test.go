package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
)

func TestHomeHandler(t *testing.T) {
	logger, _ := test.NewNullLogger()

	req := httptest.NewRequest(
		"GET", "http://example.com/foo", nil,
	)
	w := httptest.NewRecorder()
	handler := HomeHandler(logger)
	handler(w, req)

	if http.StatusOK != w.Code {
		t.Errorf("fail")
	}
}
