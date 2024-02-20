package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"shared/libs/lib-a"
)

func init() {
	log.Println("Starting service-b")
}

func Start() (err error) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/api/service-b", rootHandler)

	return http.ListenAndServe(":8010", r)
}

type HttpResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := HttpResponse{
		Status:  "success",
		Message: fmt.Sprintf("service-b answered: %s", lib_a.LibA()),
	}

	w.Header().Set("Content-Type", "application/json")

	responseJson, _ := json.Marshal(response)

	_, _ = w.Write(responseJson)
}
