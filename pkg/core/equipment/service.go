package equipment

import (
	"ctrl-hub-technical-challenge/pkg/core/model"
	"errors"
)

// TODO - actually make this a real service either connecting to data or to another supplying service

type Service struct {
	equipmentMap map[string]model.EquipmentItem
}

func NewService() *Service {
	service := &Service{
		equipmentMap: make(map[string]model.EquipmentItem),
	}

	// Statically add some data for examples sake
	service.equipmentMap["2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49"] = model.EquipmentItem{
		ID:                 "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
		Name:               "AirCat - Drill - 4337",
		VibrationMagnitude: 2.1,
	}
	service.equipmentMap["2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d50"] = model.EquipmentItem{
		ID:                 "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d50",
		Name:               "AirCat - Drill - 4338",
		VibrationMagnitude: 2.2,
	}
	return service
}

func (s *Service) GetEquipmentByID(id string) (model.EquipmentItem, error) {
	equipment, ok := s.equipmentMap[id]
	if !ok {
		return model.EquipmentItem{}, errors.New("equipment not found")
	}
	return equipment, nil
}
