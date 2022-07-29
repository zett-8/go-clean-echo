package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	e := Echo()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := IndexHandler(c)

	if err != nil {
		t.Error("Failed to access /", err)
	}

	if rec.Code != http.StatusOK {
		t.Error("Status is not 200 but", rec.Code)
	}

	if rec.Body.String() != "Hello World" {
		t.Errorf("Expected %s but Got %s", "Hello World", rec.Body.String())
	}
}
