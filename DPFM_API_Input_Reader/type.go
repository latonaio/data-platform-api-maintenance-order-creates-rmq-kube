package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"MaintenanceOrder"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	MaintenanceOrder               string           `json:"MaintenanceOrder"`
	MaintOrderRoutingNumber        *string          `json:"MaintOrderRoutingNumber"`
	MaintenanceOrderType           *string          `json:"MaintenanceOrderType"`
	MaintenanceOrderDesc           *string          `json:"MaintenanceOrderDesc"`
	MaintOrdBasicStart             *string          `json:"MaintOrdBasicStart"`
	MaintOrdBasicEnd               *string          `json:"MaintOrdBasicEnd"`
	MaintOrdBasicStartDate         *string          `json:"MaintOrdBasicStartDate"`
	MaintOrdBasicStartTime         *string          `json:"MaintOrdBasicStartTime"`
	MaintOrdBasicEndDate           *string          `json:"MaintOrdBasicEndDate"`
	MaintOrdBasicEndTime           *string          `json:"MaintOrdBasicEndTime"`
	MaintOrdSchedldBscStrt         *string          `json:"MaintOrdSchedldBscStrt"`
	MaintOrdSchedldBscEnd          *string          `json:"MaintOrdSchedldBscEnd"`
	ScheduledBasicStartDate        *string          `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime        *string          `json:"ScheduledBasicStartTime"`
	ScheduledBasicEndDate          *string          `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime          *string          `json:"ScheduledBasicEndTime"`
	MaintenanceNotification        *string          `json:"MaintenanceNotification"`
	OrdIsNotSchedldAutomatically   *string          `json:"OrdIsNotSchedldAutomatically"`
	MainWorkCenterInternalID       *string          `json:"MainWorkCenterInternalID"`
	MainWorkCenterTypeCode         *string          `json:"MainWorkCenterTypeCode"`
	MainWorkCenter                 *string          `json:"MainWorkCenter"`
	MainWorkCenterPlant            *string          `json:"MainWorkCenterPlant"`
	WorkCenterInternalID           *string          `json:"WorkCenterInternalID"`
	WorkCenterTypeCode             *string          `json:"WorkCenterTypeCode"`
	WorkCenter                     *string          `json:"WorkCenter"`
	MaintenancePlanningPlant       *string          `json:"MaintenancePlanningPlant"`
	MaintenancePlant               *string          `json:"MaintenancePlant"`
	Assembly                       *string          `json:"Assembly"`
	MaintOrdProcessPhaseCode       *string          `json:"MaintOrdProcessPhaseCode"`
	MaintOrdProcessSubPhaseCode    *string          `json:"MaintOrdProcessSubPhaseCode"`
	BusinessArea                   *string          `json:"BusinessArea"`
	CompanyCode                    *string          `json:"CompanyCode"`
	CostCenter                     *string          `json:"CostCenter"`
	ReferenceElement               *string          `json:"ReferenceElement"`
	FunctionalArea                 *string          `json:"FunctionalArea"`
	AdditionalDeviceData           *string          `json:"AdditionalDeviceData"`
	Equipment                      *string          `json:"Equipment"`
	EquipmentName                  *string          `json:"EquipmentName"`
	FunctionalLocation             *string          `json:"FunctionalLocation"`
	MaintenanceOrderPlanningCode   *string          `json:"MaintenanceOrderPlanningCode"`
	MaintenancePlannerGroup        *string          `json:"MaintenancePlannerGroup"`
	MaintenanceActivityType        *string          `json:"MaintenanceActivityType"`
	MaintPriority                  *string          `json:"MaintPriority"`
	MaintPriorityType              *string          `json:"MaintPriorityType"`
	OrderProcessingGroup           *string          `json:"OrderProcessingGroup"`
	ProfitCenter                   *string          `json:"ProfitCenter"`
	ResponsibleCostCenter          *string          `json:"ResponsibleCostCenter"`
	MaintenanceRevision            *string          `json:"MaintenanceRevision"`
	SerialNumber                   *string          `json:"SerialNumber"`
	SuperiorProjectNetwork         *string          `json:"SuperiorProjectNetwork"`
	OperationSystemCondition       *string          `json:"OperationSystemCondition"`
	Project                        *string          `json:"Project"`
	ControllingObjectClass         *string          `json:"ControllingObjectClass"`
	MaintenanceOrderInternalID     *string          `json:"MaintenanceOrderInternalID"`
	MaintenanceObjectList          *string          `json:"MaintenanceObjectList"`
	MaintObjectLocAcctAssgmtNmbr   *string          `json:"MaintObjectLocAcctAssgmtNmbr"`
	AssetLocation                  *string          `json:"AssetLocation"`
	AssetRoom                      *string          `json:"AssetRoom"`
	PlantSection                   *string          `json:"PlantSection"`
	BasicSchedulingType            *string          `json:"BasicSchedulingType"`
	LatestAcceptableCompletionDate *string          `json:"LatestAcceptableCompletionDate"`
	MaintOrdPersonResponsible      *string          `json:"MaintOrdPersonResponsible"`
	LastChange                     *string          `json:"LastChange"`
	ControllingSettlementProfile   *string          `json:"ControllingSettlementProfile"`
	SystemStatusText               *string          `json:"SystemStatusText"`
	UserStatusText                 *string          `json:"UserStatusText"`
	TechnicalObject                *string          `json:"TechnicalObject"`
	TechnicalObjectLabel           *string          `json:"TechnicalObjectLabel"`
	TechObjIsEquipOrFuncnlLoc      *string          `json:"TechObjIsEquipOrFuncnlLoc"`
	ObjectListItem                 []ObjectListItem `json:"ObjectListItem"`
}

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
