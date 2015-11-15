package dolstats

import (
	"testing"
	"time"
)

func TestUnmarshalOSHACompliance(t *testing.T) {
	mock, err := loadMock("./testdata/dol_osha_compliance/osha_fatalities.json")
	if err != nil {
		t.Error("Error loading OSHA Fatalities mock data. Error was:", err)
	}

	a, err := unmarshalOSHAFatalities(mock)
	if err != nil {
		t.Error("Error unmarshaling OSHA Fatalities. Error was:", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid Data length. Length is:", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid ID value returned. Value was:", resOne.ID)
	}

	date := time.Time(resOne.Date)
	if date.Month() != 5 {
		t.Error("Invalid Month value returned. Value was:", date.Month())
	}

	if date.Day() != 3 {
		t.Error("Invalid Day value returned. Value was:", date.Day())
	}

	if date.Year() != 2001 {
		t.Error("Invalid Year value returned. Value was:", date.Year())
	}
}
