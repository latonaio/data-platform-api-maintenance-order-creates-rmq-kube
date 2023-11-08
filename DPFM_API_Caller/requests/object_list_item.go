package requests

type ObjectListItem struct {
	MaintenanceOrder            string  `json:"MaintenanceOrder"`
	MaintenanceOrderObjectList  int     `json:"MaintenanceOrderObjectList"`
	MaintenanceObjectListItem   int     `json:"MaintenanceObjectListItem"`
	Equipment                   *string `json:"Equipment"`
	MaintenanceNotification     *string `json:"MaintenanceNotification"`
	Assembly                    *string `json:"Assembly"`
	Product                     *string `json:"Product"`
	SerialNumber                *string `json:"SerialNumber"`
	UniqueItemIdentifier        *string `json:"UniqueItemIdentifier"`
	FunctionalLocation          *string `json:"FunctionalLocation"`
	MaintObjectListItemSequence *string `json:"MaintObjectListItemSequence"`
}
