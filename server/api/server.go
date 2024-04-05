package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Operation struct {
	Operator string    `json:"operator"`
	Operands []float64 `json:"operands"`
}

type Server struct {
	*mux.Router

	pastOperations []Operation
}

func NewServer() *Server {
	s := &Server{
		Router:         mux.NewRouter(),
		pastOperations: []Operation{},
	}
	return s
}

func (s *Server) Calculate(o Operation) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
