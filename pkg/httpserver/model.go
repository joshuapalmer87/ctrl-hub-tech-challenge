package httpserver

type ExposurePost struct {
	EquipmentID string `json:"equipment_id"`
	Duration    int    `json:"duration"`
	UserID      string `json:"user_id"`
}
