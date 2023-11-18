package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-maintenance-order-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-maintenance-order-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToObjectListItemCreates(sdc *dpfm_api_input_reader.SDC) (*[]ObjectListItem, error) {
	items := make([]ObjectListItem, 0)

	for _, data := range sdc.Header.ObjectListItem {
		item, err := TypeConverter[*ObjectListItem](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToOperationCreates(sdc *dpfm_api_input_reader.SDC) (*[]Operation, error) {
	items := make([]Operation, 0)

	for _, data := range sdc.Header.Operation {
		item, err := TypeConverter[*Operation](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToOperationComponentCreates(sdc *dpfm_api_input_reader.SDC) (*[]OperationComponent, error) {
	items := make([]OperationComponent, 0)

	for _, data := range *sdc.Header.OperationComponent {
		item, err := TypeConverter[*OperationComponent](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToObjectListItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ObjectListItemUpdates) (*[]ObjectListItem, error) {
	items := make([]ObjectListItem, 0)

	for _, data := range *itemUpdates {
		objectListItem, err := TypeConverter[*ObjectListItem](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *objectListItem)
	}

	return &items, nil
}

func ConvertToOperationUpdates(itemUpdates *[]dpfm_api_processing_formatter.OperationUpdates) (*[]Operation, error) {
	items := make([]Operation, 0)

	for _, data := range *itemUpdates {
		operation, err := TypeConverter[*Operation](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *operation)
	}

	return &items, nil
}

func ConvertToOperationComponentUpdates(itemUpdates *[]dpfm_api_processing_formatter.OperationComponentUpdates) (*[]OperationComponent, error) {
	items := make([]OperationComponent, 0)

	for _, data := range *itemUpdates {
		operationComponent, err := TypeConverter[*OperationComponent](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *operationComponent)
	}

	return &items, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
