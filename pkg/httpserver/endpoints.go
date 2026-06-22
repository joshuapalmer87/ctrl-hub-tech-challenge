package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// TODO - add an error logging layer somewhere around here

func (s *HttpServer) exposureEndpoint(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if req.Method == "GET" {
		exposures, err := s.handleGetAllExposure(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO - return a more meaningful error message
			return
		}
		w.WriteHeader(http.StatusOK)
		// TODO - handle error
		w.Write(exposures)
	}

	if req.Method == "POST" {
		exposure, err := s.handlePostExposure(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO - return a more meaningful error message
			return
		}
		w.WriteHeader(http.StatusCreated)
		// TODO - handle error
		w.Write(exposure)
	}
}

// TODO #Important note - this endpoint is a bad idea in general, but it's in the spec
func (s *HttpServer) handleGetAllExposure(req *http.Request) ([]byte, error) {
	records, err := s.exposureService.GetAllRecords()
	if err != nil {
		return nil, fmt.Errorf("unable to get all records", err)
	}
	return json.Marshal(records)
}

func (s *HttpServer) handlePostExposure(req *http.Request) ([]byte, error) {
	// Decode and validate
	var exposure ExposurePost
	err := json.NewDecoder(req.Body).Decode(&exposure)
	if err != nil {
		return nil, fmt.Errorf("unable to decode exposure provided", err)
	}

	err = uuid.Validate(exposure.EquipmentID)
	if err != nil {
		return nil, fmt.Errorf("equipment ID is not a valid UUID: %v", err)
	}

	err = uuid.Validate(exposure.UserID)
	if err != nil {
		return nil, fmt.Errorf("user ID is not a valid UUID: %v", err)
	}

	record, err := s.exposureService.CreateExposureRecord(exposure.UserID, exposure.EquipmentID, exposure.Duration)
	if err != nil {
		return nil, fmt.Errorf("unable to persist exposure record", err)
	}

	return json.Marshal(record)
}
