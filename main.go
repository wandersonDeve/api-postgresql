package main

import (
	"api-postgresql/configs"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err!= nil {
        panic(err)
    }

	r := chi.NewRouter()

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
