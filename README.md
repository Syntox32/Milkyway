# Milkyway

Making a (for now) simple REST framework that I can use in my cloud computing class.

### Features
 - Simple regex based URL matching
 - Easy seperation of functionality for HTTP methods
 - Easier JSON response integration
 - Built-in JSON errors (with correct status codes)

### Example Usage

```
package main

import (
        "net/http"
        "github.com/Syntox32/Milkyway"
)

type Person struct {
        Name    string  `json:"name"`
}

func hello(w http.ResponseWriter, r *http.Request) milkyway.JsonObject {
        return &Person{Name: "Ola Nordmann"}
}

func main() {
        way := milkyway.GenMilkyway()
        way.Route("GET", "^/test$", hello)
        way.Serve()
}
```

This will give the following output from `curl -v localhost:8080/test`:

```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 19 Sep 2018 02:12:10 GMT
< Content-Length: 24
<
{"name":"Ola Nordmann"}
```

Errors look like this:

```
< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Wed, 19 Sep 2018 02:19:32 GMT
< Content-Length: 33
<
{"msg":"Not Found","status":404}
```

To turn off JSON errors, do `way.JsonErrors = false`.

### Running test
```
go test
```
