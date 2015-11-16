package dolstats

import (
	"testing"
	"time"
)

func TestUnmarshalFullMine(t *testing.T) {
	mock, err := loadMock("./testdata/full_mine_information/msha_mines.json")
	if err != nil {
		t.Error("Error loading FullMine mock data. Error was:", err)
	}

	a, err := unmarshalFullMine(mock)
	if err != nil {
		t.Error("Error unmarshaling FullMine. Error was:", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid Data length. Length is:", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid ID value returned. Value was:", resOne.ID)
	}

	d := time.Time(resOne.CurrentStatusDate)

	if d.Day() != 25 {
		t.Error("Invalid day value returned. Value was:", d.Day())
	}

	if d.Month() != 6 {
		t.Error("Invalid month value returned. Value was:", d.Month())
	}

	if d.Year() != 1987 {
		t.Error("Invalid year value returned. Value was:", d.Year())
	}

	if resOne.PortableOperation != false {
		t.Error("Invalid PortableOperation value returned. Should have been false.")
	}
}
