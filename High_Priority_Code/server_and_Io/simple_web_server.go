//Create a simple web server that responds with "Hello, World!" when accessed.

package serverandio

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func simple_web_server() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
}
