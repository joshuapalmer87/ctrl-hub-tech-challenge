package httpserver

import (
	"context"
	"ctrl-hub-technical-challenge/pkg/core/model"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const addr = ":8090" //TODO - make this more easily adaptable, move to config package
const exposureId = "exposure_id"

type ExposureService interface {
	CreateExposureRecord(userId, equipmentId string, durationMinutes int, exposureDateTime time.Time) (model.Exposure, error)
	GetAllExposureRecords() ([]model.Exposure, error)
	GetExposureRecord(ID string) (model.Exposure, error)
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
	http.HandleFunc("GET /ping", ping)
	http.HandleFunc("GET /exposure", s.getAllExposure)
	http.HandleFunc("POST /exposure", s.postExposure)
	http.HandleFunc("GET /exposure/{"+exposureId+"}", s.getExposure)

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
	fmt.Fprintf(w, "ping\n")
}
