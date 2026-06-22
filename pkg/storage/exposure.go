package storage

import "ctrl-hub-technical-challenge/pkg/core/model"

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

func (s *Service) GetExposures() []model.Exposure {
	return nil
}
