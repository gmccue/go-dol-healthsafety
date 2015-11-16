package dolstats

import "encoding/json"

const fullMineURI = "Mining/FullMineInfo/MSHA_mines"

type FullMineInfo struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ID                         OSHAInt  `json:"id"`
		MineID                     string   `json:"MINE_ID,omitempty"`
		CurrentMineName            string   `json:"CURRENT_MINE_NAME,omitempty"`
		CoalMetaInd                string   `json:"COAL_METAL_IND,omitempty"`
		CurrentMineType            string   `json:"CURRENT_MINE_TYPE,omitempty"`
		CurrentMineStatus          string   `json:"CURRENT_MINE_STATUS,omitempty"`
		CurrentStatusDate          OSHADate `json:"CURRENT_STATUS_DT,omitempty"`
		CurrentControllerID        string   `json:"CURRENT_CONTROLLER_ID,omitempty"`
		CurrentControllerName      string   `json:"CURRENT_CONTROLLER_NAME,omitempty"`
		CurrentOperatorID          string   `json:"CURRENT_OPERATOR_ID,omitempty"`
		CurrentOperatorName        string   `json:"CURRENT_OPERATOR_NAME,omitempty"`
		State                      string   `json:"STATE,omitempty"`
		BOMStateCode               string   `json:"BOM_STATE_CD,omitempty"`
		FIPSCountyCode             string   `json:"FIPS_CNTY_CD,omitempty"`
		FIPSCityName               string   `json:"FIPS_CNTY_NM,omitempty"`
		CongressionalDistrictCode  string   `json:"CONG_DIST_CD,omitempty"`
		CompanyType                string   `json:"COMPANY_TYPE,omitempty"`
		CurrentControllerStartDate OSHADate `json:"CURRENT_CONTROLLER_BEGIN_DT,omitempty"`
		District                   string   `json:"DISTRICT,omitempty"`
		OfficeCode                 string   `json:"OFFICE_CD,omitempty"`
		OfficeName                 string   `json:"OFFICE_NAME,omitempty"`
		AssessControlNum           string   `json:"ASSESS_CTRL_NO,omitempty"`
		PrimarySICCode             string   `json:"PRIMARY_SIC_CD,omitempty"`
		PrimarySIC                 string   `json:"PRIMARY_SIC,omitempty"`
		PrimarySICCode1            string   `json:"PRIMARY_SIC_CD_1,omitempty"`
		PrimarySICCodeSFX          string   `json:"PRIMARY_SIC_CD_SFX,omitempty"`
		SecondarySICCode           string   `json:"SECONDARY_SIC_CD,omitempty"`
		SecondarySIC               string   `json:"SECONDARY_SIC,omitempty"`
		SecondarySICCode1          string   `json:"SECONDARY_SIC_CD_1,omitempty"`
		SecondarySICCodeSFX        string   `json:"SECONDARY_SIC_CD_SFX,omitempty"`
		PrimaryCanvassCode         string   `json:"PRIMARY_CANVASS_CD,omitempty"`
		PrimaryCanvass             string   `json:"PRIMARY_CANVASS,omitempty"`
		SecondaryCanvassCode       string   `json:"SECONDARY_CANVASS_CD,omitempty"`
		SecondaryCanvass           string   `json:"SECONDARY_CANVASS,omitempty"`
		Current103I                string   `json:"CURRENT_103I,omitempty"`
		Current103IDT              OSHADate `json:"CURRENT_103I_DT,omitempty"`
		PortableOperation          OSHABool `json:"PORTABLE_OPERATION,omitempty"`
		PortableFIPSStateCode      string   `json:"PORTABLE_FIPS_ST_CD,omitempty"`
		DaysPerWeek                float64  `json:"DAYS_PER_WEEK,omitempty"`
		HoursPerShift              float64  `json:"HOURS_PER_SHIFT,omitempty"`
		ProductionShiftsPerDay     float64  `json:"PROD_SHIFTS_PER_DAY,omitempty"`
		MaintenanceShiftsPerDay    float64  `json:"MAINT_SHIFTS_PER_DAY,omitempty"`
		NumEmployees               float64  `json:"NO_EMPLOYEES,omitempty"`
		Part48Training             OSHABool `json:"PART48_TRAINING,omitempty"`
		Longitude                  float64  `json:"LONGITUDE,omitempty"`
		Latitude                   float64  `json:"LATITUDE,omitempty"`
		AvgMineHeight              int      `json:"AVG_MINE_HEIGHT,omitempty"`
		MineGasCategoryCode        string   `json:"MINE_GAS_CATEGORY_CD,omitempty"`
		MethaneLiberation          float64  `json:"METHANE_LIBERATION,omitempty"`
		NumProducingPits           float64  `json:"NO_PRODUCING_PITS,omitempty"`
		NumNonproducingPits        float64  `json:"NO_NONPRODUCING_PITS,omitempty"`
		NumTailingPonds            float64  `json:"NO_TAILING_PONDS,omitempty"`
		PillarRecoveryUsed         OSHABool `json:"PILLAR_RECOVERY_USED,omitempty"`
		HighwallMinerUsed          OSHABool `json:"HIGHWALL_MINER_USED,omitempty"`
		MultiplePits               OSHABool `json:"MULTIPLE_PITS,omitempty"`
		MinersRepIndicator         OSHABool `json:"MINERS_REP_IND,omitempty"`
		SafetyCommitteeIndicator   OSHABool `json:"SAFETY_COMMITTEE_IND,omitempty"`
		MilesFromOffice            float64  `json:"MILES_FROM_OFFICE,omitempty"`
		DirectionsToMine           string   `json:"DIRECTIONS_TO_MINE,omitempty"`
		NearestTown                string   `json:"NEAREST_TOWN,omitempty"`
	} `json:"d"`
}

func (api *API) QueryFullMineInfo() (FullMineInfo, error) {
	api.buildPath(fullMineURI)
	api.doRequest()

	return unmarshalFullMine(api.RawData)
}

func unmarshalFullMine(rawData []byte) (FullMineInfo, error) {
	var f FullMineInfo

	err := json.Unmarshal(rawData, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}
