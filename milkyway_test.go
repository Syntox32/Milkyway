package milkyway

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

type TestObject struct {
	Msg	string	`json:"msg"`
}

var way *Milkyway = GenMilkyway()

func TestSimpleJsonResponse(t *testing.T) {
	doJsonRequest(t,
		"GET", "^/test$", "/test",
		&TestObject{Msg:"success"},
		`{"msg":"success"}`,
		http.StatusOK)

	doJsonRequest(t,
		"GET", "^/failure$", "/fail",
		&TestObject{Msg:"success"},
		`{"msg":"Not Found","status":404}`,
		http.StatusNotFound)
}


func doJsonRequest(t *testing.T,
		method string,
		pattern string,
		url string,
		object interface{},
		expected string,
		expectedCode int) {

	req, err := http.NewRequest(strings.ToUpper(method), "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	if way == nil {}
	way.Route(strings.ToUpper(method), pattern,
		func(w http.ResponseWriter, r *http.Request) JsonObject {
			return &object
		})
	http.HandlerFunc(way.Router).ServeHTTP(rr, req)
	way.Routes = make([]*Route, 0)

	if status := rr.Code; status != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedCode)
	}

	if strings.Replace(rr.Body.String(), "\n", "", -1) != expected {
		t.Errorf("handler returned unexpected body: got %v (%d) want %v (%d)",
			rr.Body.String(), len(rr.Body.String()), expected, len(expected))
	}
}
