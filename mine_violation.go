package dolstats

import "encoding/json"

const mshaViolationsURI = "Mining/Violation/MSHA_violations"

type MineViolations struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID                         OSHAInt `json:"ID"`
		Event                      string  `json:"EVENT_NO,omitempty"`
		InspectionStartDate        string  `json:"INSPECTION_BEGIN_DT,omitempty"`
		InspectionEndDate          string  `json:"INSPECTION_END_DT,omitempty"`
		ViolationNumber            string  `json:"VIOLATION_NO,omitempty"`
		ControllerID               string  `json:"CONTROLLER_ID,omitempty"`
		ControllerName             string  `json:"CONTROLLER_NAME,omitempty"`
		ViolatorID                 string  `json:"VIOLATOR_ID,omitempty"`
		ViolatorName               string  `json:"VIOLATOR_NAME,omitempty"`
		ViolatorTypeCode           string  `json:"VIOLATOR_TYPE_CD,omitempty"`
		MineID                     string  `json:"MINE_ID,omitempty"`
		MineName                   string  `json:"MINE_NAME,omitempty"`
		MineType                   string  `json:"MINE_TYPE,omitempty"`
		CoalMetalInd               string  `json:"COAL_METAL_IND,omitempty"`
		ContractorID               string  `json:"CONTRACTOR_ID,omitempty"`
		ViolationIssueDate         string  `json:"VIOLATION_ISSUE_DT,omitempty"`
		ViolationDate              string  `json:"VIOLATION_OCCUR_DT,omitempty"`
		CalendarYear               OSHAInt `json:"CAL_YR,omitempty"`
		CalendarQuarter            OSHAInt `json:"CAL_QTR,omitempty"`
		FiscalYear                 OSHAInt `json:"FISCAL_YR,omitempty"`
		FiscalQuarter              OSHAInt `json:"FISCAL_QTR,omitempty"`
		ViolationIssueTime         string  `json:"VIOLATION_ISSUE_TIME,omitempty"`
		SigSub                     string  `json:"SIG_SUB,omitempty"`
		SectionOfAct               string  `json:"SECTION_OF_ACT,omitempty"`
		PartSection                string  `json:"PART_SECTION,omitempty"`
		SectionOfAct1              string  `json:"SECTION_OF_ACT_1,omitempty"`
		SectionOfAct2              string  `json:"SECTION_OF_ACT_2,omitempty"`
		CITOrdSafe                 string  `json:"CIT_ORD_SAFE,omitempty"`
		OrigTermDueDate            string  `json:"ORIG_TERM_DUE_DT,omitempty"`
		OrigTermDueTime            string  `json:"ORIG_TERM_DUE_TIME,omitempty"`
		LatestTermDueDate          string  `json:"LATEST_TERM_DUE_DT,omitempty"`
		LatestTermDueTime          string  `json:"LATEST_TERM_DUE_TIME,omitempty"`
		TerminationDate            string  `json:"TERMINATION_DT,omitempty"`
		TerminationTime            string  `json:"TERMINATION_TIME,omitempty"`
		TerminationType            string  `json:"TERMINATION_TYPE,omitempty"`
		VacateDate                 string  `json:"VACATE_DT,omitempty"`
		VacateTime                 string  `json:"VACATE_TIME,omitempty"`
		InitialViolationNumber     string  `json:"INITIAL_VIOL_NO,omitempty"`
		ReplacedByOrderNumber      string  `json:"REPLACED_BY_ORDER_NO,omitempty"`
		Likelihood                 string  `json:"LIKELIHOOD,omitempty"`
		InjIllness                 string  `json:"INJ_ILLNESS,omitempty"`
		NumAffected                OSHAInt `json:"NO_AFFECTED,omitempty"`
		Negligence                 string  `json:"NEGLIGENCE,omitempty"`
		WrittenNotice              string  `json:"WRITTEN_NOTICE,omitempty"`
		EnforcementArea            string  `json:"ENFORCEMENT_AREA,omitempty"`
		SpecialAssess              string  `json:"SPECIAL_ASSESS,omitempty"`
		PrimaryOrMill              string  `json:"PRIMARY_OR_MILL,omitempty"`
		RightToConfDate            string  `json:"RIGHT_TO_CONF_DT,omitempty"`
		ASMTGeneratedIND           string  `json:"ASMT_GENERATED_IND,omitempty"`
		FinalOrderIssueDate        string  `json:"FINAL_ORDER_ISSUE_DT,omitempty"`
		ProposedPenalty            OSHAInt `json:"PROPOSED_PENALTY,omitempty"`
		AmountDue                  OSHAInt `json:"AMOUNT_DUE,omitempty"`
		AmountPaid                 OSHAInt `json:"AMOUNT_PAID,omitempty"`
		BillPrintDate              string  `json:"BILL_PRINT_DT,omitempty"`
		LastActionCD               string  `json:"LAST_ACTION_CD,omitempty"`
		LastActionDate             string  `json:"LAST_ACTION_DT,omitempty"`
		DocketNumber               string  `json:"DOCKET_NO,omitempty"`
		DocketstatusCode           string  `json:"DOCKET_STATUS_CD,omitempty"`
		ContestedIND               string  `json:"CONTESTED_IND,omitempty"`
		ContestedDate              string  `json:"CONTESTED_DT,omitempty"`
		ViolatorViolationCount     OSHAInt `json:"VIOLATOR_VIOLATION_CNT,omitempty"`
		ViolatorInspectionDayCount OSHAInt `json:"VIOLATOR_INSPECTION_DAY_CNT,omitempty"`
	} `json:"d"`
}

func (api *API) QueryMineViolations() (MineViolations, error) {
	api.buildPath(mshaViolationsURI)
	api.doRequest()

	return unmarshalMineViolations(api.RawData)
}

func unmarshalMineViolations(rawData []byte) (MineViolations, error) {
	var m MineViolations

	err := json.Unmarshal(rawData, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}
