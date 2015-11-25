package dolstats

import (
	"testing"
	"time"
)

func TestUnmarshalMineInspections(t *testing.T) {
	mock, err := loadMock("./testdata/mine_inspections/mine_inspections.json")
	if err != nil {
		t.Error("Error loading MineInspections mock. Error was:", err)
	}

	mi, err := unmarshalMineInspections(mock)
	if err != nil {
		t.Error("Error unmarshaling MineInspections. Error was:", err)
	}

	resOne := mi.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid MineInspection ID. Value was:", resOne.ID)
	}

	bDate := time.Time(resOne.InspectionBeginDate)
	if bDate.Year() != 2000 {
		t.Error("Invalid InspectionBeginDate year. Value was:", bDate.Year())
	}

	if bDate.Month() != 10 {
		t.Error("Invalid InspectionBeginDate month. Value was:", bDate.Month())
	}

	if bDate.Day() != 3 {
		t.Error("Invalid InspectionBeginDate day. Value was:", bDate.Day())
	}

	if resOne.ExplosiveStorage != false {
		t.Error("Invalid ExplosiveStroage value. Expected false.")
	}
}
