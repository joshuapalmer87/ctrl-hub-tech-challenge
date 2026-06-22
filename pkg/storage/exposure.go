package storage

import (
	"ctrl-hub-technical-challenge/pkg/core/model"
	"errors"
)

// TODO - create a map, with getters and setters for it for now

type Service struct {
	// Terrible practise, but exposing for ease of use in tests as not using persistence
	ExposureMap map[string]model.Exposure
}

func NewService() *Service {
	service := &Service{
		ExposureMap: make(map[string]model.Exposure),
	}
	return service
}

func (s *Service) CreateExposure(exposure model.Exposure) error {
	s.ExposureMap[exposure.ID] = exposure
	return nil
}

func (s *Service) GetAllExposures() ([]model.Exposure, error) {
	records := make([]model.Exposure, 0, len(s.ExposureMap))
	for _, record := range s.ExposureMap {
		records = append(records, record)
	}
	return records, nil
}

func (s *Service) GetExposure(ID string) (model.Exposure, error) {
	exposure, ok := s.ExposureMap[ID]
	if !ok {
		return model.Exposure{}, errors.New("exposure not found with supplied id")
	}
	return exposure, nil
}
