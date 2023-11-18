package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-maintenance-order-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-maintenance-order-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-maintenance-order-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var objectListItem *[]dpfm_api_output_formatter.ObjectListItem
	var operation *[]dpfm_api_output_formatter.Operation
	var operationComponent *[]dpfm_api_output_formatter.OperationComponent
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, errs, log)
		case "ObjectListItem":
			objectListItem = c.objectListItemCreateSql(nil, mtx, input, output, errs, log)
		case "Operation":
			operation = c.operationCreateSql(nil, mtx, input, output, errs, log)
		case "OperationComponent":
			operationComponent = c.operationComponentCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		ObjectListItem:     objectListItem,
		Operation:          operation,
		OperationComponent: operationComponent,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var objectListItem *[]dpfm_api_output_formatter.ObjectListItem
	var operation *[]dpfm_api_output_formatter.Operation
	var operationComponent *[]dpfm_api_output_formatter.OperationComponent
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "ObjectListItem":
			objectListItem = c.objectListItemUpdateSql(mtx, input, output, errs, log)
		case "Operation":
			operation = c.operationUpdateSql(mtx, input, output, errs, log)
		case "OperationComponent":
			operationComponent = c.operationComponentUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		ObjectListItem:     objectListItem,
		Operation:          operation,
		OperationComponent: operationComponent,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	headerData := input.Header
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "MaintenanceOrderHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) objectListItemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ObjectListItem {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.Header.ObjectListItem {
		input.Header.ObjectListItem[i].MaintenanceOrder = input.Header.MaintenanceOrder
		objectListItemData := input.Header.ObjectListItem[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": objectListItemData, "function": "MaintenanceOrderObjectListItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ObjectListItem Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToObjectListItemCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.Header.Operation {
		input.Header.Operation[i].MaintenanceOrder = input.Header.MaintenanceOrder
		operationData := input.Header.Operation[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationData, "function": "MaintenanceOrderOperation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Operation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationComponentCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.OperationComponent {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.Header.OperationComponent {
		input.Header.OperationComponent[i].MaintenanceOrder = input.Header.MaintenanceOrder
		operationComponentData := input.Header.OperationComponent[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationComponentData, "function": "MaintenanceOrderOperationComponent", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "OperationComponent Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationComponentCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "MaintenanceOrderHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "header Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) objectListItemUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ObjectListItem {
	req := make([]dpfm_api_processing_formatter.ObjectListItemUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, objectListItem := range header.ObjectListItem {
		objectListItemData := *dpfm_api_processing_formatter.ConvertToObjectListItemUpdates(header, objectListItem)

		if objectListItemIsUpdate(&objectListItemData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": objectListItemData, "function": "MaintenanceOrderObjectListItem", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "ObjectListItem Data cannot update"
				return nil
			}
		}
		req = append(req, objectListItemData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToObjectListItemUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	req := make([]dpfm_api_processing_formatter.OperationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, operation := range header.Operation {
		operationData := *dpfm_api_processing_formatter.ConvertToOperationUpdates(header, operation)

		if operationIsUpdate(&operationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationData, "function": "MaintenanceOrderOperation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Operation Data cannot update"
				return nil
			}
		}
		req = append(req, operationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationComponentUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.OperationComponent {
	req := make([]dpfm_api_processing_formatter.OperationComponentUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, operationComponent := range header.OperationComponent {
		operationComponentData := *dpfm_api_processing_formatter.ConvertToOperationComponentUpdates(header, operationComponent)

		if operationComponentIsUpdate(&operationComponentData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationComponentData, "function": "MaintenanceOrderOperationComponent", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "OperationComponent Data cannot update"
				return nil
			}
		}
		req = append(req, operationComponentData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationComponentUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	maintenanceOrder := header.MaintenanceOrder

	return !(maintenanceOrder == 0)
}

func objectListItemIsUpdate(MaintenanceOrderObjectListItem *dpfm_api_processing_formatter.ObjectListItemUpdates) bool {
	maintenanceOrder := MaintenanceOrderObjectListItem.MaintenanceOrder
	objectListItem := MaintenanceOrderObjectListItem.MaintenanceOrderObjectList
	maintenanceObjectListItem := MaintenanceOrderObjectListItem.MaintenanceObjectListItem

	return !(maintenanceOrder == "" || objectListItem == 0 || maintenanceObjectListItem == 0)
}

func operationIsUpdate(MaintenanceOrderOperation *dpfm_api_processing_formatter.OperationUpdates) bool {
	maintenanceOrder := MaintenanceOrderOperation.MaintenanceOrder
	maintenanceOrderOperation := MaintenanceOrderOperation.MaintenanceOrderOperation
	maintenanceOrderSubOperation := MaintenanceOrderOperation.MaintenanceOrderSubOperation

	return !(maintenanceOrder == "" || maintenanceOrderOperation == "" || maintenanceOrderSubOperation == "")
}

func operationComponentIsUpdate(maintenanceOrderOperationComponent *dpfm_api_processing_formatter.OperationComponentUpdates) bool {
	maintenanceOrder := maintenanceOrderOperationComponent.MaintenanceOrder
	maintenanceOrderComponent := maintenanceOrderOperationComponent.MaintenanceOrderComponent

	return !(maintenanceOrder == "" || maintenanceOrderComponent == "")
}
