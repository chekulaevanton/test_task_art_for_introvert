package server

import (
	"log"
	"net/http"

	h "github.com/chekulaevanton/test_task_art_for_introvert/handlers"
)

type Server struct {
    addr string
    mux *http.ServeMux
    handler *h.CoursesHandler
}

func NewServer(addr string, handler *h.CoursesHandler) *Server {
    return &Server{
        addr: addr,
        mux: http.NewServeMux(),
        handler: handler,
    }
}

func (s *Server) Run() {
    log.Print("Starting server on " + s.addr)

    s.attachHandlers()
    err := http.ListenAndServe(s.addr, s.mux)
    if err != nil {
        log.Fatal("Server can not start:", err)
    }
}

func (s *Server) attachHandlers() {
    s.mux.HandleFunc("/courses", s.handler.HandleCourses)
    s.mux.HandleFunc("/course/", s.handler.HandleCourse)
}
