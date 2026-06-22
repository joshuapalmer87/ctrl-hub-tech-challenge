package exposure

import (
	"ctrl-hub-technical-challenge/pkg/core/model"
	"fmt"
	"math"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserByID(id string) (model.User, error)
}

type EquipmentService interface {
	GetEquipmentByID(id string) (model.EquipmentItem, error)
}

type Storage interface {
	CreateExposure(exposure model.Exposure) error
	GetAllExposures() ([]model.Exposure, error)
	GetExposure(id string) (model.Exposure, error)
}

type Service struct {
	userService      UserService
	equipmentService EquipmentService
	exposureStorage  Storage
}

func NewService(userService UserService, equipmentService EquipmentService, storage Storage) *Service {
	return &Service{
		userService:      userService,
		equipmentService: equipmentService,
		exposureStorage:  storage,
	}
}

func (s *Service) CreateExposureRecord(userId, equipmentId string, durationMinutes int) (model.Exposure, error) {
	// Get user
	user, err := s.userService.GetUserByID(userId)
	if err != nil {
		return model.Exposure{}, fmt.Errorf("unable to find a matching user: %w", err)
	}

	// Get equipment
	equipmentItem, err := s.equipmentService.GetEquipmentByID(equipmentId)
	if err != nil {
		return model.Exposure{}, fmt.Errorf("unable to find a matching user: %w", err)
	}
	// Create UUID
	exposureID, err := uuid.NewRandom()
	if err != nil {
		return model.Exposure{}, fmt.Errorf("unable to generate exposure id: %w", err)
	}

	exposure := model.Exposure{
		ID:              exposureID.String(),
		Equipment:       equipmentItem,
		User:            user,
		DurationMinutes: durationMinutes,
		A8:              generateExposureA8(equipmentItem.VibrationMagnitude, durationMinutes),
		Points:          generateExposurePoints(equipmentItem.VibrationMagnitude, durationMinutes),
	}

	// Store via other
	err = s.exposureStorage.CreateExposure(exposure)
	if err != nil {
		return model.Exposure{}, fmt.Errorf("unable to store exposure reading: %w", err)
	}

	return exposure, nil
}

func (s *Service) GetAllExposureRecords() ([]model.Exposure, error) {
	records, err := s.exposureStorage.GetAllExposures()
	if err != nil {
		return []model.Exposure{}, fmt.Errorf("unable to read exposures: %w", err)
	}
	return records, nil
}

func (s *Service) GetExposureRecord(ID string) (model.Exposure, error) {
	record, err := s.exposureStorage.GetExposure(ID)
	if err != nil {
		return model.Exposure{}, fmt.Errorf("unable to read exposure: %w", err)
	}
	return record, nil
}

// NB - changed from README as that was invalid golang
func generateExposureA8(vibrationMagnitude float64, triggerTime int) float64 {
	return vibrationMagnitude * math.Sqrt(float64((triggerTime/60)/8))
}

func generateExposurePoints(vibrationMagnitude float64, triggerTime int) float64 {
	points := math.Pow((vibrationMagnitude/2.5), 2) * (((float64(triggerTime) / 60) / 8) * 100)
	return math.Round(points)
}
