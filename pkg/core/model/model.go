package model

import "time"

// TODO - ideally this shouldn't be specifying json attributes here as it should be purely domain
type Exposure struct {
	ID              string        `json:"id"`
	Equipment       EquipmentItem `json:"equipment"`
	DurationMinutes int           `json:"duration"`
	A8              float64       `json:"a8"`
	Points          float64       `json:"points"`
	User            User          `json:"user"`
}

type ExposureWithTime struct {
	Exposure     Exposure
	ExposureTime time.Time
}

type ExposureSummary struct {
	A8     float64 `json:"a8"`
	Points float64 `json:"points"`
	User   User    `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EquipmentItem struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	VibrationMagnitude float64 `json:"vibration_magnitude"`
}
