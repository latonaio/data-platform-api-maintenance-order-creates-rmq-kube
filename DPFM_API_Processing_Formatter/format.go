package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-maintenance-order-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		MaintenanceOrder: data.MaintenanceOrder,
	}
}

func ConvertToObjectListItemUpdates(header dpfm_api_input_reader.Header, ObjectListItem dpfm_api_input_reader.ObjectListItem) *ObjectListItemUpdates {
	data := ObjectListItem

	return &ObjectListItemUpdates{
		MaintenanceOrderObjectList: data.MaintenanceOrderObjectList,
	}
}
