package storage

import (
	"ctrl-hub-technical-challenge/pkg/core/model"
	"errors"
	"time"
)

type Service struct {
	// Terrible practise, but exposing for ease of use in tests as not using persistence
	ExposureMap     map[string]model.Exposure
	UserExposureMap map[string][]model.ExposureWithTime
}

func NewService() *Service {
	service := &Service{
		ExposureMap:     make(map[string]model.Exposure),
		UserExposureMap: make(map[string][]model.ExposureWithTime),
	}
	return service
}

func (s *Service) CreateExposure(exposure model.Exposure, exposureDateTime time.Time) error {
	s.ExposureMap[exposure.ID] = exposure
	userExposures, ok := s.UserExposureMap[exposure.User.ID]
	if !ok {
		userExposures = []model.ExposureWithTime{}
	}
	userExposures = append(userExposures, model.ExposureWithTime{
		Exposure:     exposure,
		ExposureTime: exposureDateTime,
	})
	s.UserExposureMap[exposure.User.ID] = userExposures
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
