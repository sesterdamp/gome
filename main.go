package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type simpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", simpleRequest)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func simpleRequest(w http.ResponseWriter, r *http.Request) {
	res := simpleResponse{
		Status:  http.StatusOK,
		Message: "Success.",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		errRes := []byte(err.Error())
		w.Write(errRes)
	}
}
