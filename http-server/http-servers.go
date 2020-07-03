package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func serveJSON(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "SERVING JSON\n")
    fmt.Fprintf(w, "data {\"name\": \"Whatever\"}\n")
}

func headers(w http.ResponseWriter, req * http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
          fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/serveJSON", serveJSON)
    http.HandleFunc("/headers", headers)

    fmt.Println("Listening on port 8090")
    http.ListenAndServe(":8090", nil)
}