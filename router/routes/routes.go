package routes

import (
	"extensao-api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)
}
