package main

import (
	"agent/cmd/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	userRequest := models.UserRequest{
		Name:  "Lucas",
		Email: "lucas@gmail.com",
	}

	body, err := json.Marshal(userRequest)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		var responseApi models.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
			panic(err)
		}

		panic(responseApi.Reason)
	}

	var responseApi models.UserRequest
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println("New user created: ", responseApi)
}
