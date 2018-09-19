package main

import (
	"net/http"
	"github.com/Syntox32/Milkyway"
)

type Person struct {
	Name	string	`json:"name"`
}

func hello(w http.ResponseWriter, r *http.Request) milkyway.JsonObject {
	return &Person{Name: "Ola Nordmann"}
}

func main() {
	way := milkyway.GenMilkyway()
	way.Route("GET", "^/test$", hello)
	way.Serve()
}
