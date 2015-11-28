package dolstats

import "encoding/json"

const (
	fiAreaURI              = "Safety/FatalOccupationalInjuries/FI_AREA"
	fiCaseURI              = "Safety/FatalOccupationalInjuries/FI_CASE"
	fiCategoryURI          = "Safety/FatalOccupationalInjuries/FI_CATEGORY"
	fiCategory2URI         = "Safety/FatalOccupationalInjuries/FI_CATEGORY2"
	fiDataPubURI           = "Safety/FatalOccupationalInjuries/FI_DATA_PUB"
	fiDataTypeURI          = "Safety/FatalOccupationalInjuries/FI_DATA_TYPE"
	fiEventURI             = "Safety/FatalOccupationalInjuries/FI_EVENT"
	fiFootnoteURI          = "Safety/FatalOccupationalInjuries/FI_FOOTNOTE"
	fiGQTCaseURI           = "Safety/FatalOccupationalInjuries/FI_GQT_CASE"
	fiGQTCharURI           = "Safety/FatalOccupationalInjuries/FI_GQT_CHAR"
	fiGQTCharOwnershipURI  = "Safety/FatalOccupationalInjuries/FI_GQT_CHAR_OWNERSHIP"
	fiGQTOwnershipURI      = "Safety/FatalOccupationalInjuries/FI_GQT_OWNERSHIP"
	fiGQTStateURI          = "Safety/FatalOccupationalInjuries/FI_GQT_STATE"
	fiGQTStateOwnershipURI = "Safety/FatalOccupationalInjuries/FI_GQT_STATE_OWNERSHIP"
	fiIndustryURI          = "Safety/FatalOccupationalInjuries/FI_INDUSTRY"
	fiSeriesURI            = "Safety/FatalOccupationalInjuries/FI_SERIES"
)

type FIGQTCase GQTCase
type FIGQTChar GQTChar
type FIGQTCharOwnership GQTCharOwnership
type FIGQTOwnership GQTOwnership
type FIGQTState GQTState
type FIGQTStateOwnership GQTStateOwnership

type FIArea struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		AreaCode     string `json:"AREA_CODE,omitempty"`
		Name         string `json:"AREA_NAME,omitempty"`
		DisplayLevel string `json:"DISPLAY_LEVEL,omitempty"`
		Selectable   string `json:"SELECTABLE,omitempty"`
		SortSequence string `json:"SORT_SEQUENCE,omitempty"`
	} `json:"d,omitempty"`
}

type FICase struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		CaseType string `json:"CASE_TYPE,omitempty"`
		CaseText string `json:"CASE_TEXT,omitempty"`
	} `json:"d,omitempty"`
}

type FICategory struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		CaseCode     string `json:"CASE_CODE,omitempty"`
		CategoryCode string `json:"CATEGORY_CODE,omitempty"`
		CategoryText string `json:"CATEGORY_TEXT,omitempty"`
	} `json:"d,omitempty"`
}

type FICategory2 struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		CategoryCode string `json:"CATEGORY_CODE,omitempty"`
		CategoryText string `json:"CATEGORY_TEXT,omitempty"`
	} `json:"d,omitempty"`
}

type FIDataPub struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		SeriesID      string `json:"SERIES_ID,omitempty"`
		Year          string `json:"YEAR,omitempty"`
		Period        string `json:"PERIOD,omitempty"`
		Value         string `json:"VALUE,omitempty"`
		FootnoteCodes string `json:"FOOTNOTE_CODES,omitempty"`
	} `json:"d,omitempty"`
}

type FIDataType struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Code string `json:"DATA_TYPE_CODE,omitempty"`
		Text string `json:"DATA_TYPE_TEXT,omitempty"`
	} `json:"d,omitempty"`
}

type FIEvent struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Code         string `json:"EVENT_CODE,omitempty"`
		Text         string `json:"EVENT_TEXT,omitempty"`
		DisplayLevel string `json:"DISPLAY_LEVEL,omitempty"`
		Selectable   string `json:"SELECTABLE,omitempty"`
		SortSequence string `json:"SORT_SEQUENCE,omitempty"`
	} `json:"d,omitempty"`
}

type FIFootnote struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Code string `json:"FOOTNOTE_CODE,omitempty"`
		Text string `json:"FOOTNOTE_TEXT,omitempty"`
	} `json:"d,omitempty"`
}

type FIIndustry struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		Code         string `json:"INDUSTRY_CODE,omitempty"`
		Text         string `json:"INDUSTRY_TEXT,omitempty"`
		DisplayLevel string `json:"DISPLAY_LEVEL,omitempty"`
		Selectable   string `json:"SELECTABLE,omitempty"`
		SortSequence string `json:"SORT_SEQUENCE,omitempty"`
	} `json:"d,omitempty"`
}

type FISeries struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		ID             string `json:"SERIES_ID,omitempty"`
		CategoryCode   string `json:"CATEGORY_CODE,omitempty"`
		DataTypeCode   string `json:"DATA_TYPE_CODE,omitempty"`
		CaseCode       string `json:"CASE_CODE,omitempty"`
		IndustryCode   string `json:"INDUSTRY_CODE,omitempty"`
		EventCode      string `json:"EVENT_CODE,omitempty"`
		SourceCode     string `json:"SOURCE_CODE,omitempty"`
		OccupationCode string `json:"OCCUPATION_CODE,omitempty"`
		AreaCode       string `json:"AREA_CODE,omitempty"`
		FootnoteCodes  string `json:"FOOTNOTE_CODES,omitempty"`
		BeginYear      string `json:"BEGIN_YEAR,omitempty"`
		BeginPeriod    string `json:"BEGIN_PERIOD,omitempty"`
		EndYear        string `json:"END_YEAR,omitempty"`
		EndPeriod      string `json:"END_PERIOD,omitempty"`
	} `json:"d,omitempty"`
}

func (api *API) QueryFIArea() (FIArea, error) {
	api.buildPath(fiAreaURI)
	api.doRequest()

	return unmarshalFIArea(api.RawData)
}

func unmarshalFIArea(rawData []byte) (FIArea, error) {
	var a FIArea

	err := json.Unmarshal(rawData, &a)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (api *API) QueryFICase() (FICase, error) {
	api.buildPath(fiCaseURI)
	api.doRequest()

	return unmarshalFICase(api.RawData)
}

func unmarshalFICase(rawData []byte) (FICase, error) {
	var c FICase

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (api *API) QueryFICategory() (FICategory, error) {
	api.buildPath(fiCategoryURI)
	api.doRequest()

	return unmarshalFICategory(api.RawData)
}

func unmarshalFICategory(rawData []byte) (FICategory, error) {
	var c FICategory

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (api *API) QueryFICategory2() (FICategory2, error) {
	api.buildPath(fiCategory2URI)
	api.doRequest()

	return unmarshalFICategory2(api.RawData)
}

func unmarshalFICategory2(rawData []byte) (FICategory2, error) {
	var c FICategory2

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (api *API) QueryFIDataPub() (FIDataPub, error) {
	api.buildPath(fiDataPubURI)
	api.doRequest()

	return unmarshalFIDataPub(api.RawData)
}

func unmarshalFIDataPub(rawData []byte) (FIDataPub, error) {
	var d FIDataPub

	err := json.Unmarshal(rawData, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (api *API) QueryFIDataType() (FIDataType, error) {
	api.buildPath(fiDataTypeURI)
	api.doRequest()

	return unmarshalFIDataType(api.RawData)
}

func unmarshalFIDataType(rawData []byte) (FIDataType, error) {
	var d FIDataType

	err := json.Unmarshal(rawData, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (api *API) QueryFIEvent() (FIEvent, error) {
	api.buildPath(fiEventURI)
	api.doRequest()

	return unmarshalFIEvent(api.RawData)
}

func unmarshalFIEvent(rawData []byte) (FIEvent, error) {
	var e FIEvent

	err := json.Unmarshal(rawData, &e)
	if err != nil {
		return e, err
	}

	return e, nil
}

func (api *API) QueryFIFootnote() (FIFootnote, error) {
	api.buildPath(fiFootnoteURI)
	api.doRequest()

	return unmarshalFIFootnote(api.RawData)
}

func unmarshalFIFootnote(rawData []byte) (FIFootnote, error) {
	var f FIFootnote

	err := json.Unmarshal(rawData, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

func (api *API) QueryFIGQTCase() (GQTCase, error) {
	api.buildPath(fiGQTCaseURI)
	api.doRequest()

	return unmarshalFIGQTCase(api.RawData)
}

func unmarshalFIGQTCase(rawData []byte) (GQTCase, error) {
	var c GQTCase

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (api *API) QueryFIGQTChar() (GQTChar, error) {
	api.buildPath(fiGQTCharURI)
	api.doRequest()

	return unmarshalFIGQTChar(api.RawData)
}

func unmarshalFIGQTChar(rawData []byte) (GQTChar, error) {
	var c GQTChar

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (api *API) QueryFIGQTCharOwnership() (GQTCharOwnership, error) {
	api.buildPath(fiGQTCharOwnershipURI)
	api.doRequest()

	return unmarshalFIGQTCharOwnership(api.RawData)
}

func unmarshalFIGQTCharOwnership(rawData []byte) (GQTCharOwnership, error) {
	var c GQTCharOwnership

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (api *API) QueryFIGQTOwnership() (GQTOwnership, error) {
	api.buildPath(fiGQTOwnershipURI)
	api.doRequest()

	return unmarshalFIGQTOwnership(api.RawData)
}

func unmarshalFIGQTOwnership(rawData []byte) (GQTOwnership, error) {
	var o GQTOwnership

	err := json.Unmarshal(rawData, &o)
	if err != nil {
		return o, err
	}

	return o, nil
}

func (api *API) QueryFIGQTState() (GQTState, error) {
	api.buildPath(fiGQTStateURI)
	api.doRequest()

	return unmarshalFIGQTState(api.RawData)
}

func unmarshalFIGQTState(rawData []byte) (GQTState, error) {
	var s GQTState

	err := json.Unmarshal(rawData, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (api *API) QueryFIStateOwnership() (GQTStateOwnership, error) {
	api.buildPath(fiGQTStateOwnershipURI)
	api.doRequest()

	return unmarshalFIGQTStateOwnership(api.RawData)
}

func unmarshalFIGQTStateOwnership(rawData []byte) (GQTStateOwnership, error) {
	var s GQTStateOwnership

	err := json.Unmarshal(rawData, &s)
	if err != nil {
		return s, err
	}

	return s, nil

}

func (api *API) QueryFIIndustry() (FIIndustry, error) {
	api.buildPath(fiIndustryURI)
	api.doRequest()

	return unmarshalFIIndustry(api.RawData)
}

func unmarshalFIIndustry(rawData []byte) (FIIndustry, error) {
	var i FIIndustry

	err := json.Unmarshal(rawData, &i)
	if err != nil {
		return i, err
	}

	return i, nil
}

func (api *API) QueryFISeries() (FISeries, error) {
	api.buildPath(fiSeriesURI)
	api.doRequest()

	return unmarshalFISeries(api.RawData)
}

func unmarshalFISeries(rawData []byte) (FISeries, error) {
	var s FISeries

	err := json.Unmarshal(rawData, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}
