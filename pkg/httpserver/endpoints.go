package httpserver

import (
	"ctrl-hub-technical-challenge/pkg/core/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type ExposureService interface {
	CreateExposureRecord(userId, equipmentId string, durationMinutes int, exposureDateTime time.Time) (model.Exposure, error)
	GetAllExposureRecords() ([]model.Exposure, error)
	GetExposureRecord(ID string) (model.Exposure, error)
	GetUserExposureSummary(userId string, startDateTime, endDateTime time.Time) (model.ExposureSummary, error)
}

// TODO - add an error logging layer somewhere around here

// TODO #Important note - this endpoint is a bad idea in general, but it's in the spec
func (s *HttpServer) getAllExposure(w http.ResponseWriter, req *http.Request) {
	records, err := s.exposureService.GetAllExposureRecords()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}
	if result, ok := marshallJSON(w, records); ok {
		w.WriteHeader(http.StatusOK)
		// TODO - handle error
		w.Write(result)
	}
}

func (s *HttpServer) postExposure(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	var exposure ExposurePost
	err := json.NewDecoder(req.Body).Decode(&exposure)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}
	err = s.validatePostExposure(exposure)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}

	record, err := s.exposureService.CreateExposureRecord(exposure.UserID, exposure.EquipmentID, exposure.Duration, currentTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}

	if exposureResult, ok := marshallJSON(w, record); ok {
		w.WriteHeader(http.StatusCreated)
		// TODO - handle error
		w.Write(exposureResult)
	}
}

func (s *HttpServer) getExposure(w http.ResponseWriter, req *http.Request) {
	idString := req.PathValue(exposureId)
	err := uuid.Validate(idString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}

	record, err := s.exposureService.GetExposureRecord(idString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}
	if result, ok := marshallJSON(w, record); ok {
		w.WriteHeader(http.StatusOK)
		// TODO - handle error
		w.Write(result)
	}
}

func (s *HttpServer) getExposureSummary(w http.ResponseWriter, req *http.Request) {
	userIdString := req.PathValue(userId)
	err := uuid.Validate(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}

	startDateTimeString := req.URL.Query().Get(startingAt)
	startDateTime, err := time.Parse(time.RFC3339, startDateTimeString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	endDateTimeString := req.URL.Query().Get(endingAt)
	endDateTime, err := time.Parse(time.RFC3339, endDateTimeString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	record, err := s.exposureService.GetUserExposureSummary(userIdString, startDateTime, endDateTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return
	}
	if result, ok := marshallJSON(w, record); ok {
		w.WriteHeader(http.StatusOK)
		// TODO - handle error
		w.Write(result)
	}
}

func (s *HttpServer) validatePostExposure(exposure ExposurePost) error {
	err := uuid.Validate(exposure.EquipmentID)
	if err != nil {
		return fmt.Errorf("equipment ID is not a valid UUID: %v", err)
	}

	err = uuid.Validate(exposure.UserID)
	if err != nil {
		return fmt.Errorf("user ID is not a valid UUID: %v", err)
	}

	return nil
}

func marshallJSON(w http.ResponseWriter, response any) ([]byte, bool) {
	exposureResult, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO - return a more meaningful error message
		return nil, false
	}
	return exposureResult, true
}
