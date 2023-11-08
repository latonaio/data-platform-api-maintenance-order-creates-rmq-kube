package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header         *Header           `json:"Header"`
	ObjectListItem *[]ObjectListItem `json:"ObjectListItem"`
}

type Header struct {
	MaintenanceOrder               string  `json:"MaintenanceOrder"`
	MaintOrderRoutingNumber        *string `json:"MaintOrderRoutingNumber"`
	MaintenanceOrderType           *string `json:"MaintenanceOrderType"`
	MaintenanceOrderDesc           *string `json:"MaintenanceOrderDesc"`
	MaintOrdBasicStart             *string `json:"MaintOrdBasicStart"`
	MaintOrdBasicEnd               *string `json:"MaintOrdBasicEnd"`
	MaintOrdBasicStartDate         *string `json:"MaintOrdBasicStartDate"`
	MaintOrdBasicStartTime         *string `json:"MaintOrdBasicStartTime"`
	MaintOrdBasicEndDate           *string `json:"MaintOrdBasicEndDate"`
	MaintOrdBasicEndTime           *string `json:"MaintOrdBasicEndTime"`
	MaintOrdSchedldBscStrt         *string `json:"MaintOrdSchedldBscStrt"`
	MaintOrdSchedldBscEnd          *string `json:"MaintOrdSchedldBscEnd"`
	ScheduledBasicStartDate        *string `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime        *string `json:"ScheduledBasicStartTime"`
	ScheduledBasicEndDate          *string `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime          *string `json:"ScheduledBasicEndTime"`
	MaintenanceNotification        *string `json:"MaintenanceNotification"`
	OrdIsNotSchedldAutomatically   *string `json:"OrdIsNotSchedldAutomatically"`
	MainWorkCenterInternalID       *string `json:"MainWorkCenterInternalID"`
	MainWorkCenterTypeCode         *string `json:"MainWorkCenterTypeCode"`
	MainWorkCenter                 *string `json:"MainWorkCenter"`
	MainWorkCenterPlant            *string `json:"MainWorkCenterPlant"`
	WorkCenterInternalID           *string `json:"WorkCenterInternalID"`
	WorkCenterTypeCode             *string `json:"WorkCenterTypeCode"`
	WorkCenter                     *string `json:"WorkCenter"`
	MaintenancePlanningPlant       *string `json:"MaintenancePlanningPlant"`
	MaintenancePlant               *string `json:"MaintenancePlant"`
	Assembly                       *string `json:"Assembly"`
	MaintOrdProcessPhaseCode       *string `json:"MaintOrdProcessPhaseCode"`
	MaintOrdProcessSubPhaseCode    *string `json:"MaintOrdProcessSubPhaseCode"`
	BusinessArea                   *string `json:"BusinessArea"`
	CompanyCode                    *string `json:"CompanyCode"`
	CostCenter                     *string `json:"CostCenter"`
	ReferenceElement               *string `json:"ReferenceElement"`
	FunctionalArea                 *string `json:"FunctionalArea"`
	AdditionalDeviceData           *string `json:"AdditionalDeviceData"`
	Equipment                      *string `json:"Equipment"`
	EquipmentName                  *string `json:"EquipmentName"`
	FunctionalLocation             *string `json:"FunctionalLocation"`
	MaintenanceOrderPlanningCode   *string `json:"MaintenanceOrderPlanningCode"`
	MaintenancePlannerGroup        *string `json:"MaintenancePlannerGroup"`
	MaintenanceActivityType        *string `json:"MaintenanceActivityType"`
	MaintPriority                  *string `json:"MaintPriority"`
	MaintPriorityType              *string `json:"MaintPriorityType"`
	OrderProcessingGroup           *string `json:"OrderProcessingGroup"`
	ProfitCenter                   *string `json:"ProfitCenter"`
	ResponsibleCostCenter          *string `json:"ResponsibleCostCenter"`
	MaintenanceRevision            *string `json:"MaintenanceRevision"`
	SerialNumber                   *string `json:"SerialNumber"`
	SuperiorProjectNetwork         *string `json:"SuperiorProjectNetwork"`
	OperationSystemCondition       *string `json:"OperationSystemCondition"`
	Project                        *string `json:"Project"`
	ControllingObjectClass         *string `json:"ControllingObjectClass"`
	MaintenanceOrderInternalID     *string `json:"MaintenanceOrderInternalID"`
	MaintenanceObjectList          *string `json:"MaintenanceObjectList"`
	MaintObjectLocAcctAssgmtNmbr   *string `json:"MaintObjectLocAcctAssgmtNmbr"`
	AssetLocation                  *string `json:"AssetLocation"`
	AssetRoom                      *string `json:"AssetRoom"`
	PlantSection                   *string `json:"PlantSection"`
	BasicSchedulingType            *string `json:"BasicSchedulingType"`
	LatestAcceptableCompletionDate *string `json:"LatestAcceptableCompletionDate"`
	MaintOrdPersonResponsible      *string `json:"MaintOrdPersonResponsible"`
	LastChange                     *string `json:"LastChange"`
	ControllingSettlementProfile   *string `json:"ControllingSettlementProfile"`
	SystemStatusText               *string `json:"SystemStatusText"`
	UserStatusText                 *string `json:"UserStatusText"`
	TechnicalObject                *string `json:"TechnicalObject"`
	TechnicalObjectLabel           *string `json:"TechnicalObjectLabel"`
	TechObjIsEquipOrFuncnlLoc      *string `json:"TechObjIsEquipOrFuncnlLoc"`
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
