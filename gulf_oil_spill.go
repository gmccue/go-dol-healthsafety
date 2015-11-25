package dolstats

import "encoding/json"

const (
	directReadURI  = "Safety/GulfOilSpill/OSHA_Direct_Read_Sampling"
	noiseReportURI = "Safety/GulfOilSpill/OSHA_NOISE_REPORT"
)

type DirectRead struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID             OSHAInt  `json:"id"`
		State          string   `json:"ESTAB_STATE,omitempty"`
		SiteName       string   `json:"ESTAB_SITE_NAME,omitempty"`
		SiteNameOther  string   `json:"ESTAB_SITE_NAME_OTHER,omitempty"`
		SampleDate     OSHADate `json:"SAMP_DATE,omitempty"`
		SampleNumber   int      `json:"SAMP_NO,omitempty"`
		OCCCode        string   `json:"OCC_CODE,omitempty"`
		Description    string   `json:"DES_SITU_ENCOUNTERED,omitempty"`
		Instrument     string   `json:"INSTRU_1,omitempty"`
		SubstanceAgent string   `json:"SUBSTANCE_AGENT,omitempty"`
		Reading        string   `json:"READING,omitempty"`
	} `json:"d"`
}

type NoiseReport struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID                       OSHAInt  `json:"id"`
		State                    string   `json:"ESTABLISHMENT_STATE,omitempty"`
		Name                     string   `json:"ESTABLISHMENT_NAME,omitempty"`
		SamplingDate             OSHADate `json:"SAMPLING_DATE,omitempty"`
		SamplingNumber           int      `json:"SAMPLING_NO,omitempty"`
		EmployeeOccupation       string   `json:"EMPLOYEE_OCCUPATION,omitempty"`
		EmployeeOccupationDetail string   `json:"EMPLOYEE_OCCUPATION_SPECIFY,omitempty"`
		Description              string   `json:"DES_SITU_ENCOUNTERED,omitempty"`
		SampleType               string   `json:"SAMPLE_TYPE,omitempty"`
		TotalTime                string   `json:"TOTAL_TIME,omitempty"`
		ExposureLevel            float64  `json:"EXPOSURE_LEVEL,omitempty"`
		PEL                      string   `json:"PEL,omitempty"`
		OSHAHCLAVG               float64  `json:"OSHA_HC_LAVG,omitempty"`
		OSHAHCTWA                float64  `json:"OSHA_HC_TWA,omitempty"`
		OSHAPELLAVG              float64  `json:"OSHA_PEL_LAVG,omitempty"`
		OSHAPELTWA               float64  `json:"OSHA_PEL_TWA,omitempty"`
	} `json:"d"`
}

func (api *API) QueryOSHADirectReadSampling() (DirectRead, error) {
	api.buildPath(directReadURI)
	api.doRequest()

	return unmarshalDirectRead(api.RawData)
}

func unmarshalDirectRead(rawData []byte) (DirectRead, error) {
	var d DirectRead

	err := json.Unmarshal(rawData, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (api *API) QueryOSHANoiseReport() (NoiseReport, error) {
	api.buildPath(noiseReportURI)
	api.doRequest()

	return unmarshalNoiseReport(api.RawData)
}

func unmarshalNoiseReport(rawData []byte) (NoiseReport, error) {
	var n NoiseReport

	err := json.Unmarshal(rawData, &n)
	if err != nil {
		return n, err
	}

	return n, nil
}
