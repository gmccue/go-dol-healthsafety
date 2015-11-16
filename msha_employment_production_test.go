package dolstats

import "testing"

func TestUnmarshalMSHAEmploymentProduction(t *testing.T) {
	mock, err := loadMock("./testdata/msha_employment_production/msha_employment_production.json")
	if err != nil {
		t.Error("Error loading EmploymentProduction mock data. Error was: ", err)
	}

	a, err := unmarshalMSHAEmploymentProduction(mock)
	if err != nil {
		t.Error("Error unmarshaling MSHAEmploymentProduction. Error was: ", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid Data length. Length is: ", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid ID value returned. Value was:", resOne.ID)
	}

	if resOne.MineID != 2900418 {
		t.Error("Invalid MineID value returned. Value was:", resOne.MineID)
	}
}
