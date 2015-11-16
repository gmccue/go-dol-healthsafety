package dolstats

import "encoding/json"

const employmentProductionURI = "Employment/MSHA_employmentProduction"

type MSHAEmploymentProduction struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		ID                   OSHAInt `json:"id,omitempty"`
		MineID               OSHAInt `json:"mineId,omitempty"`
		MineName             string  `json:"mineName,omitempty"`
		State                string  `json:"state,omitempty"`
		SubUnitCode          int     `json:"subUnitCode,omitempty"`
		SubUnit              string  `json:"subUnit,omitempty"`
		CalendarYear         int     `json:"calendarYear,omitempty"`
		AnnualHours          int     `json:"annualHours,omitempty"`
		AnnualCostProduction int     `json:"annualCostProduction,omitempty"`
		AverageEmployeeCount int     `json:"averageEmployeeCount,omitempty"`
		AverageEmployeeHours int     `json:"averageEmployeeHours,omitempty"`
		CoalMetalInd         string  `json:"coalMetalInd,omitempty"`
	} `json:"d,omitempty"`
}

func (api *API) QueryMSHAEmploymentProduction() (MSHAEmploymentProduction, error) {
	api.buildPath(employmentProductionURI)
	api.doRequest()

	return unmarshalMSHAEmploymentProduction(api.RawData)
}

func unmarshalMSHAEmploymentProduction(rawData []byte) (MSHAEmploymentProduction, error) {
	var m MSHAEmploymentProduction

	err := json.Unmarshal(rawData, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}
