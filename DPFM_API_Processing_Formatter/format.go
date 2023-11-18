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
		MaintenanceOrder:           data.MaintenanceOrder,
		MaintenanceOrderObjectList: data.MaintenanceOrderObjectList,
		MaintenanceObjectListItem:  data.MaintenanceObjectListItem,
	}
}

func ConvertToOperationUpdates(header dpfm_api_input_reader.Header, Operation dpfm_api_input_reader.Operation) *OperationUpdates {
	data := Operation

	return &OperationUpdates{
		MaintenanceOrder:             data.MaintenanceOrder,
		MaintenanceOrderOperation:    data.MaintenanceOrderOperation,
		MaintenanceOrderSubOperation: data.MaintenanceOrderSubOperation,
	}
}

func ConvertToOperationComponentUpdates(header dpfm_api_input_reader.Header, OperationComponent dpfm_api_input_reader.OperationComponent) *OperationComponentUpdates {
	data := OperationComponent

	return &OperationComponentUpdates{
		MaintenanceOrder:          data.MaintenanceOrder,
		MaintenanceOrderComponent: data.MaintenanceOrderComponent,
	}
}
