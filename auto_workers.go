package dolstats

import "encoding/json"

const carClosedURI = "AutoWorkers/CAR_ClosedRepurposed"

type CarClosed struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID                        OSHAInt `json:"id"`
		NumberID                  OSHAInt `json:"Number_ID,omitempty"`
		LastAutomakerOwner        string  `json:"Last_Automaker_Owner,omitempty"`
		AutomakerPlantName        string  `json:"Automaker_Plant_Name,omitempty"`
		City                      string  `json:"City,omitempty"`
		State                     string  `json:"State,omitempty"`
		PlantProductCategory      string  `json:"Plant_Product_Category,omitempty"`
		YearClosed                string  `json:"Year_Closed_Scheduled_To_Close,omitempty"`
		CurrentStatus             string  `json:"Current_Conditions,omitempty"`
		AutomotiveOther           string  `json:"Automotive_Other,omitempty"`
		IndustrialAutomotiveMFG   string  `json:"Industrial_Automotive_MFG,omitempty"`
		Commercial                string  `json:"Commercial,omitempty"`
		Industrial                string  `json:"Industrial,omitempty"`
		LogisticsAndWarehousing   string  `json:"Logistics_And_Warehousing,omitempty"`
		Education                 string  `json:"Education,omitempty"`
		ResearchAndDevelopment    string  `json:"Research_and_Development,omitempty"`
		Recreational              string  `json:"Recreational,omitempty"`
		Government                string  `json:"Government,omitempty"`
		Residential               string  `json:"Residential,omitempty"`
		Vacant                    string  `json:"Vacant,omitempty"`
		Demolished                string  `json:"Demolished,omitempty"`
		SpecificReuseNotes        string  `json:"Specific_Reuse_Notes,omitempty"`
		FederalIncentivesProvided string  `json:"Federal_Incentives_Provided,omitempty"`
		StateIncentivesProvided   string  `json:"State_Incentives_Provided,omitempty"`
		LocalIncentivesProvided   string  `json:"Local_Incentives_Provided,omitempty"`
		LocalFoundationAssistance string  `json:"Local_or_Regional_Foundation_Assistance,omitempty"`
		Success5VS1NS             string  `json:"SuccessInPropertyValueRestoration5VS1NS,omitempty"`
		SuccessRestoring5VS1NS    string  `json:"SuccessInRestoringJobBase5VS1NS,omitempty"`
	} `json:"d"`
}

func (api *API) QueryCarClosedRepurposed() (CarClosed, error) {
	api.buildPath(carClosedURI)
	api.doRequest()

	return unmarshalCarClosed(api.RawData)
}

func unmarshalCarClosed(rawData []byte) (CarClosed, error) {
	var c CarClosed

	err := json.Unmarshal(rawData, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
