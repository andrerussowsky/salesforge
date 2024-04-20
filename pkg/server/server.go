package server

import (
    "github.com/gorilla/mux"
)

type Server struct {
    Router *mux.Router
}

func NewServer() *Server {
    return &Server{
        Router: mux.NewRouter(),
    }
}
