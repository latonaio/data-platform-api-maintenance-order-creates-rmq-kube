package dpfm_api_processing_formatter

type HeaderUpdates struct {
	MaintenanceOrder int `json:"MaintenanceOrder"`
}

type ObjectListItemUpdates struct {
	MaintenanceOrder           int `json:"MaintenanceOrder"`
	MaintenanceOrderObjectList int `json:"MaintenanceOrderObjectList"`
}
