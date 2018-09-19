package milkyway

import (
	"fmt"
	"log"
	"regexp"
	"net/http"
	"strings"
	"encoding/json"
)

// A very bad play on *The Milje Way*

type Milkyway struct {
	Server		*http.ServeMux
	Routes		[]*Route
	JsonErrors	bool
}

type Route struct {
	Method	string
	Pattern	string
	Regex	*regexp.Regexp
	Handle	JsonHandler //http.HandlerFunc
}

type JsonErrorResponse struct {
	Msg	string	`json:"msg"`
	Status int	`json:"status"`
}

type JsonObject interface{}
type JsonHandler func(w http.ResponseWriter, r *http.Request) JsonObject


func (r *Route) String() string {
	return fmt.Sprintf("Route(method=%s, pattern=%s)", r.Method, r.Pattern)
}

func GenMilkyway() *Milkyway {
	m := &Milkyway{}
	m.JsonErrors = true
	http.HandleFunc("/", m.Router)
	log.Println("Creating new milkyway router...")
	return m
}

func (m *Milkyway) httpError(code int, w http.ResponseWriter, path string) {
	log.Printf("%d - %s", code, path)
	if m.JsonErrors {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(&JsonErrorResponse{
			Msg: http.StatusText(code),
			Status: code,
		})
	} else {
		http.Error(w, http.StatusText(code), code)
	}
}

func (m *Milkyway) Router(w http.ResponseWriter, r *http.Request) {
	foundMatch := false
	foundMethod := false
	log.Printf("%s %s", strings.ToUpper(r.Method), r.URL.Path)
	for _, route := range m.Routes {
		match := route.Regex.FindStringSubmatch(r.URL.Path)
		if match != nil {
			foundMatch = true
			if strings.ToUpper(r.Method) == strings.ToUpper(route.Method) {
				foundMethod = true
				jsonObject := route.Handle(w, r)
				w.Header().Add("Content-Type", "application/json")
				json.NewEncoder(w).Encode(jsonObject)
				return
			}
		}
	}
	if foundMatch && !foundMethod {
		m.httpError(http.StatusMethodNotAllowed, w, r.URL.Path)
	} else {
		m.httpError(http.StatusNotFound, w, r.URL.Path)
	}
}

func (m *Milkyway) Route(method string, pattern string, fn JsonHandler) {
	log.Printf("Adding route: %s %s", method, pattern)
	regex := regexp.MustCompile(pattern)
	route := &Route{
		Method:  method,
		Pattern: pattern,
		Regex:   regex,
		Handle:  fn,
	}
	m.Routes = append(m.Routes, route)
}

func (m *Milkyway) Serve() {
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
