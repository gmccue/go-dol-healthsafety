package dolstats

import "encoding/json"

const mineInspectionURI = "Compliance/MineInspection/mineInspections"

type MineInspections struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		ID                             OSHAInt  `json:"id,omitempty"`
		EventNumber                    OSHAInt  `json:"eventNumber,omitempty"`
		MineID                         OSHAInt  `json:"mineId,omitempty"`
		InspectionBeginDate            OSHADate `json:"inspectionBeginDate,omitempty"`
		InspectionEndDate              OSHADate `json:"inspectionEndDate,omitempty"`
		ControllerID                   string   `json:"controllerId,omitempty"`
		ControllerName                 string   `json:"controllerName,omitempty"`
		OperatorID                     string   `json:"operatorId,omitempty"`
		OperatorName                   string   `json:"operatorName,omitempty"`
		CalendarYear                   OSHAInt  `json:"calendarYear,omitempty"`
		CalendarQuarter                OSHAInt  `json:"calendarQuarter,omitempty"`
		FiscalYear                     OSHAInt  `json:"fiscalYear,omitempty"`
		FiscalQuarter                  OSHAInt  `json:"fiscalQuarter,omitempty"`
		InspectionOfficeCode           string   `json:"inspectionOfficeCode,omitempty"`
		ActivityCode                   string   `json:"activityCode,omitempty"`
		Activity                       string   `json:"activity,omitempty"`
		ActivitySections               OSHAInt  `json:"activitySections,omitempty"`
		IdleSections                   OSHAInt  `json:"idleSections,omitempty"`
		ShaftSlopeSink                 OSHAInt  `json:"shaftSlopeSink,omitempty"`
		ImpoundConstr                  OSHAInt  `json:"impoundConstr,omitempty"`
		BuildingConstrSites            OSHAInt  `json:"buildingConstrSites,omitempty"`
		Draglines                      OSHAInt  `json:"draglines,omitempty"`
		UnclassifiedConstr             OSHAInt  `json:"unclassifiedConstr,omitempty"`
		CoRecords                      OSHABool `json:"coRecords,omitempty"`
		SurfUgMine                     OSHABool `json:"surfUgMine,omitempty"`
		SurfFacilityMine               OSHABool `json:"surfFacilityMine,omitempty"`
		RefusePile                     OSHABool `json:"refusePile,omitempty"`
		ExplosiveStorage               OSHABool `json:"explosiveStorage,omitempty"`
		OutbyArea                      OSHABool `json:"outbyArea,omitempty"`
		MajorConst                     OSHABool `json:"majorConst,omitempty"`
		ShaftsSlopes                   OSHABool `json:"shaftsSlopes,omitempty"`
		Impoundments                   OSHABool `json:"impoundments,omitempty"`
		MiscArea                       OSHABool `json:"miscArea,omitempty"`
		ProgramArea                    string   `json:"programArea,omitempty"`
		SumOfSampleCntAir              OSHAInt  `json:"sumOfSampleCntAir,omitempty"`
		SumOfSampleCntDstspot          OSHAInt  `json:"sumOfSampleCntDustspot,omitempty"`
		SumOfSampleCntDustSurvey       OSHAInt  `json:"sumOfSampleCntDustSurvey,omitempty"`
		SumOfSampleCntRespDust         OSHAInt  `json:"sumOfSampleCntRespDust,omitempty"`
		SumOfSampleCntNoise            OSHAInt  `json:"sumOfSampleCntNoise,omitempty"`
		SumOfSampleCntOther            OSHAInt  `json:"sumOfSampleCntOther,omitempty"`
		NbrInspector                   OSHAInt  `json:"nbrInspector,omitempty"`
		SumOfTotalInspHours            OSHAInt  `json:"sumOfTotalInspHours,omitempty"`
		SumOfTotalOnSiteHours          OSHAInt  `json:"sumOfTotalOnSiteHours,omitempty"`
		CoalMetalInd                   string   `json:"coalMetalInd,omitempty"`
		SumOfTotalInspHrsSpvrTrainee   OSHAInt  `json:"sumOfTotalInspHrsSpvrTrainee,omitempty"`
		SumOfTotalOnSiteHrsSpvrTrainee OSHAInt  `json:"sumOfTotalOnSiteHrsSpvrTrainee,omitempty"`
	} `json:"d,omitempty"`
}

func (api *API) QueryMineInspections() (MineInspections, error) {
	api.buildPath(mineInspectionURI)
	api.doRequest()

	return unmarshalMineInspections(api.RawData)
}

func unmarshalMineInspections(rawData []byte) (MineInspections, error) {
	var m MineInspections

	err := json.Unmarshal(rawData, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}
