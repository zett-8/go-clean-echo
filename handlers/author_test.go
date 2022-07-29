package handlers

//
//import (
//	"github.com/zett-8/go-clean-echo/db"
//	"github.com/zett-8/go-clean-echo/services"
//	"github.com/zett-8/go-clean-echo/stores"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestAuthorGetCase(t *testing.T) {
//	db, _ := db.Mock()
//
//	e := Echo()
//	s := stores.New(db)
//	ss := services.New(s)
//	New(e, ss)
//
//	req := httptest.NewRequest(http.MethodGet, "/aaa", nil)
//	rec := httptest.NewRecorder()
//
//	err := e.NewContext(req, rec)
//
//	if rec.Code != http.StatusOK {
//		t.Error("Status is not 200 but", rec.Code)
//	}
//
//	if rec.Body.String() != "Hello World" {
//		t.Errorf("err %s", rec.Body.String())
//	}
//}
