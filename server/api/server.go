package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"server/app"
)

type Server struct {
	*mux.Router
	calculator *app.Calculator
}

func NewServer() *Server {
	s := &Server{
		Router:     mux.NewRouter(),
		calculator: app.NewCalculator(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello, World!"))
		if err != nil {
			return
		}
	})
	s.HandleFunc("/calculate", s.calculate()).Methods("POST")
	s.HandleFunc("/past", s.getPastOperations()).Methods("GET")
}

func (s *Server) getPastOperations() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		operations := s.calculator.GetPastOperations()
		if err := json.NewEncoder(writer).Encode(operations); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) calculate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var operation app.Operation
		// Decode the request body into the Operation struct
		// If there is an error, return a 400 Bad Request
		// with an error message
		if err := json.NewDecoder(request.Body).Decode(&operation); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Use the Calculator to perform the operation
		// If there is an error, return a 400 Bad Request
		// with an error message
		result, err := s.calculator.Perform(operation)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Encode the result into the response body
		// If there is an error, return a 500 Internal Server Error
		// with an error message
		if err := json.NewEncoder(writer).Encode(*result); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

	}
}
