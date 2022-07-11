package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type server struct {
	router *mux.Router
}

func New() *server {
	r := mux.NewRouter()

	r.HandleFunc("/update", func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "unknown handler", http.StatusNotImplemented)
	}).Methods(http.MethodPost)

	return &server{router: r}
}

func (s *server) AddHandler(method, path string, h Handler) {
	s.router.HandleFunc(path, h.Handle).Methods(method)
}

func (s *server) Run() error {
	return http.ListenAndServe(":8080", s.router)
}

type Server interface {
	Run() error
	AddHandler(method, path string, h Handler)
}

var _ Server = &server{}
