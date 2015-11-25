package dolstats

import "encoding/json"

const mineAccidentsURI = "Safety/Accidents/mineAccidents"

type MineAccident struct {
	Data []struct {
		Metadata struct {
			URI  string `json:"uri,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"__metadata,omitempty"`
		ID                        OSHAInt `json:"id,omitempty"`
		MineID                    string  `json:"mineId,omitempty"`
		ControllerID              string  `json:"controllerId,omitempty"`
		ControllerName            string  `json:"controllerName,omitempty"`
		OperatorID                string  `json:"operatorId,omitempty"`
		OperatorName              string  `json:"operatorName,omitempty"`
		ContractID                string  `json:"contractId,omitempty"`
		DocumentNumber            string  `json:"documentNumber,omitempty"`
		SubUnitCode               string  `json:"subUnitCode,omitempty"`
		SubUnit                   string  `json:"subUnit,omitempty"`
		AccidentDate              string  `json:"accidentDate,omitempty"`
		CalendarYear              int     `json:"calendarYear,omitempty"`
		CalendarQuarter           int     `json:"calendarQuarter,omitempty"`
		FiscalQuarter             int     `json:"fiscalQuarter,omitempty"`
		AccidentTime              string  `json:"accidentTime,omitempty"`
		DegreeInjuryCode          string  `json:"degreeInjuryCode,omitempty"`
		DegreeInjury              string  `json:"degreeInjury,omitempty"`
		FIPSStateCode             string  `json:"fipsStateCode,omitempty"`
		UgLocationCode            string  `json:"ugLocationCode,omitempty"`
		UgLocation                string  `json:"ugLocation,omitempty"`
		UgMiningMethodCode        string  `json:"ugMiningMethodCode,omitempty"`
		UgMiningMethod            string  `json:"ugMiningMethod,omitempty"`
		MiningEquipmentCode       string  `json:"miningEquipmentCode,omitempty"`
		MiningEquipment           string  `json:"miningEquip,omitempty"`
		EquipmentManufacturerCode string  `json:"equipmentManufacturerCode,omitempty"`
		EquipmentManufacturerName string  `json:"equipmentManufacturerName,omitempty"`
		EquipmentModelNumber      string  `json:"equpmentModelNumber,omitempty"`
		ShiftBeginTime            string  `json:"shiftBeginTime,omitempty"`
		ClassificationCode        string  `json:"classificationCode,omitempty"`
		Classification            string  `json:"classification,omitempty"`
		AccidentTypeCode          string  `json:"accidentTypeCode,omitempty"`
		AccidentType              string  `json:"accidentType,omitempty"`
		NoInjuries                string  `json:"noInjuries,omitempty"`
		TotalExperience           string  `json:"totalExperience,omitempty"`
		MineExperience            string  `json:"mineExperience,omitempty"`
		JobExperience             string  `json:"jobExperience,omitempty"`
		OccupationCode            string  `json:"occupationCode,omitempty"`
		Occupation                string  `json:"occupation,omitempty"`
		ActivityCode              string  `json:"activityCode,omitempty"`
		Activity                  string  `json:"activity,omitempty"`
		InjurySourceCode          string  `json:"injurySourceCode,omitempty"`
		InjurySource              string  `json:"injurySource,omitempty"`
		NatureInjuryCode          string  `json:"natureInjuryCode,omitempty"`
		NatureInjury              string  `json:"natureInjury,omitempty"`
		InjuredBodyPartCode       string  `json:"injuredBodyPartCode,omitempty"`
		InjuredBodyPart           string  `json:"injuredBodyPart,omitempty"`
		ScheduleCharge            string  `json:"scheduleCharge,omitempty"`
		DaysRestrict              string  `json:"daysRestrict,omitempty"`
		DaysLost                  string  `json:"daysLost,omitempty"`
		TransTerm                 string  `json:"transTerm,omitempty"`
		ReturnToWorkDate          string  `json:"returnToWorkDate,omitempty"`
		ImmedNotifyCode           string  `json:"immedNotifyCode,omitempty"`
		ImmedNotify               string  `json:"immedNotify,omitempty"`
		InvestigationStartDate    string  `json:"investBeginDate,omitempty"`
		Narrative                 string  `json:"narrative,omitempty"`
		ClosedDocNumber           string  `json:"closedDocNumber,omitempty"`
		CoalMetalID               string  `json:"coalMetalInd,omitempty"`
	} `json:"d,omitempty"`
}

func (api *API) QueryMineAccidents() (MineAccident, error) {
	api.buildPath(mineAccidentsURI)
	api.doRequest()

	return unmarshalMineAccident(api.RawData)
}

func unmarshalMineAccident(rawData []byte) (MineAccident, error) {
	var m MineAccident

	err := json.Unmarshal(rawData, &m)
	if err != nil {
		return m, err
	}

	return m, err
}
