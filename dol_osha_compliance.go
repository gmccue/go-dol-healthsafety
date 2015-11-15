package dolstats

import "encoding/json"

const OSHAFatalitiesURI = "Safety/Fatalities/OSHA_Fatalities"

type OSHAFatalities struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID          OSHAInt  `json:"id"`
		Date        OSHADate `json:"dateofincident"`
		Company     string   `json:"companyandlocation,omitempty"`
		Description string   `json:"description,omitempty"`
	} `json:"d"`
}

func (api *API) QueryOSHAFatalities() (OSHAFatalities, error) {
	api.buildPath(OSHAFatalitiesURI)
	api.doRequest()

	return unmarshalOSHAFatalities(api.RawData)
}

func unmarshalOSHAFatalities(rawData []byte) (OSHAFatalities, error) {
	var f OSHAFatalities

	err := json.Unmarshal(rawData, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}
