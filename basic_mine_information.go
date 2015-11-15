package dolstats

import "encoding/json"

const addressRecordURI = "Mining/BasicMineInfo/MSHA_addressOfRecord"

type AddressRecord struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID             OSHAInt  `json:"id"`
		MineID         OSHAInt  `json:"mineId"`
		MineName       string   `json:"mineName,omitempty"`
		ContactTitle   string   `json:"contactTitle,omitempty"`
		NearestTown    string   `json:"nearestTown,omitempty"`
		BusinessName   string   `json:"businessName,omitempty"`
		Street         string   `json:"street,omitempty"`
		POBox          int      `json:"poBox,omitempty"`
		City           string   `json:"city,omitempty"`
		StateCode      string   `json:"stateAbbr,omitempty"`
		FIPSStateCode  int      `json:"fipsStateCode,omitempty"`
		State          string   `json:"state,omitempty"`
		Zip            int      `json:"zipCode,omitempty"`
		Country        string   `json:"country,omitempty"`
		Province       string   `json:"province,omitempty"`
		PostalCode     string   `json:"postalCode,omitempty"`
		MineTypeCode   string   `json:"mineTypeCode,omitempty"`
		MineStatus     string   `json:"mineStatus,omitempty"`
		MineStatusDate OSHADate `json:"mineStatusDate,omitempty"`
		PrimarySICCode string   `json:"primarySicCode,omitempty"`
		CoalMetalInd   string   `json:"coalMetalInd,omitempty"`
	} `json:"d"`
}

func (api *API) QueryAddressOfRecord() (AddressRecord, error) {
	api.buildPath(addressRecordURI)

	err := api.doRequest()
	if err != nil {
		return AddressRecord{}, err
	}

	return unmarshalAddressRecord(api.RawData)
}

func unmarshalAddressRecord(rawData []byte) (AddressRecord, error) {
	var a AddressRecord

	err := json.Unmarshal(rawData, &a)
	if err != nil {
		return a, err
	}

	return a, nil
}
