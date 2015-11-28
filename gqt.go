package dolstats

type GQTCase struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Copmonent   string  `json:"GQT_CASE_COMPONENT,omitempty"`
		Code        string  `json:"GQT_CASE_CODE,omitempty"`
		Text        string  `json:"GQT_CASE_TEXT,omitempty"`
		TableName   string  `json:"GQT_CASE_TABLE_NAME,omitempty"`
		ColumnNames string  `json:"GQT_CASE_COLUMN_NAMES,omitempty"`
		CaseTitle   string  `json:"GQT_CASE_TITLE,omitempty"`
		Order       string  `json:"GQT_ORDER,omitempty"`
		Display     OSHAInt `json:"GQT_DISPLAY,omitempty"`
		CodeDisplay OSHAInt `json:"GQT_CODE_DISPLAY,omitempty"`
	} `json:"d,omitempty"`
}

type GQTChar struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		GQTReportID      string `json:"GQT_REPORT_ID,omitempty"`
		GQTCaseCode      string `json:"GQT_CASE_CODE,omitempty"`
		GQTCaseComponent string `json:"GQT_CASE_COMPONENT,omitempty"`
	} `json:"d,omitempty"`
}

type GQTCharOwnership struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		GQTCaseCode   string `json:"GQT_CASE_CODE,omitempty"`
		OwnershipCode string `json:"OWNERSHIP_CODE,omitempty"`
		UniqueCode    string `json:"UNIQUE_CODE,omitempty"`
		PrivateCode   string `json:"PRIVATE_CODE,omitempty"`
	}
}

type GQTOwnership struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Code         string `json:"OWNERSHIP_CODE,omitempty"`
		Name         string `json:"OWNERSHIP_NAME,omitempty"`
		SortSequence string `json:"SORT_SEQUENCE,omitempty"`
	} `json:"d,omitempty"`
}

type GQTState struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Code         string `json:"STATE_CODE,omitempty"`
		Name         string `json:"STATE_NAME,omitempty"`
		SortSequence string `json:"SORT_SEQUENCE,omitempty"`
	} `json:"d,omitempty"`
}

type GQTStateOwnership struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Year          string `json:"YEAR,omitempty"`
		OwnershipCode string `json:"OWNERSHIP_CODE,omitempty"`
		StateCode     string `json:"STATE_CODE,omitempty"`
		StateName     string `json:"STATE_NAME,omitempty"`
		SortSequence  string `json:"SORT_SEQUENCE,omitempty"`
	} `json:"d,omitempty"`
}
