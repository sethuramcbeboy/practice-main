package serverandio

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func rest_api() {
	http.HandleFunc("/hello", helloHandler2)
	http.ListenAndServe(":8080", nil)
}

//How would you create and manage a RESTful API in Go?

//You can use the net/http package to create and manage a RESTful API:
