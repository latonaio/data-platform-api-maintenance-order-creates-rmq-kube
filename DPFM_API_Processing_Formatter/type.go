package dpfm_api_processing_formatter

type HeaderUpdates struct {
	MaintenanceOrder int `json:"MaintenanceOrder"`
}

type ObjectListItemUpdates struct {
	MaintenanceOrder           string `json:"MaintenanceOrder"`
	MaintenanceOrderObjectList int    `json:"MaintenanceOrderObjectList"`
	MaintenanceObjectListItem  int    `json:"MaintenanceObjectListItem"`
}

type OperationUpdates struct {
	MaintenanceOrder             string `json:"MaintenanceOrder"`
	MaintenanceOrderOperation    string `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation string `json:"MaintenanceOrderSubOperation"`
}

type OperationComponentUpdates struct {
	MaintenanceOrder          string `json:"MaintenanceOrder"`
	MaintenanceOrderComponent string `json:"MaintenanceOrderComponent"`
}
