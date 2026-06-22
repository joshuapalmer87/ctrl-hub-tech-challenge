package httpserver

import (
	"context"
	"ctrl-hub-technical-challenge/pkg/core/model"
	"errors"
	"fmt"
	"net/http"
)

const addr = ":8090" //TODO - make this more easily adaptable, move to config package

type ExposureService interface {
	CreateExposureRecord(userId, equipmentId string, durationMinutes int) (model.Exposure, error)
}
type HttpServer struct {
	server          *http.Server
	exposureService ExposureService
}

func NewHttpServer(service ExposureService) *HttpServer {
	return &HttpServer{
		server:          &http.Server{Addr: addr},
		exposureService: service,
	}
}

// Serve registers the routes and starts the server, blocking until it is shut
// down. A graceful shutdown returns a nil error; any other failure is returned.
func (s *HttpServer) Serve() error {

	// TODO - move to NewServeMux instead to only allowing one server instance
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/exposure", s.exposureEndpoint)

	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}

	return nil
}

// Shutdown gracefully stops the server, letting in-flight requests finish until
// they complete or the provided context is cancelled.
func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func ping(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		fmt.Fprintf(w, "ping\n")
	}
}
