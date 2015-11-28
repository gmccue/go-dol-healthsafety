package dolstats

import "testing"

func TestUnmarshalFIArea(t *testing.T) {
	mock, err := loadMock("./testdata/fatal_occupational_injuries/fi_area.json")
	if err != nil {
		t.Error("Error loading FI_Area mock data. Error was:", err)
	}

	a, err := unmarshalFIArea(mock)
	if err != nil {
		t.Error("error unmarshaling FI_Area. Error was:", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid data length found for FI_Area. Length is:", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.AreaCode != "M70" {
		t.Error("Invalid AreaCode found for FI_Area.")
	}
}

func TestUnmarshalFICase(t *testing.T) {
	mock, err := loadMock("./testdata/fatal_occupational_injuries/fi_area.json")
	if err != nil {
		t.Error("Error loading FI_Area mock data. Error was:", err)
	}

	a, err := unmarshalFIArea(mock)
	if err != nil {
		t.Error("Error unmarshaling FI_Area. Error was:", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid data length found for FI_Area. Length is:", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.AreaCode != "M70" {
		t.Error("Invalid AreaCode found for FI_Area.")
	}
}

func TestUnmarshalFICategory(t *testing.T) {
	mock, err := loadMock("./testdata/fatal_occupational_injuries/fi_category.json")
	if err != nil {
		t.Error("Error loading FI_Category mock data. Error was:", err)
	}

	a, err := unmarshalFICategory(mock)
	if err != nil {
		t.Error("Error unmarshaling FI_Category. Error was:", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid data length found for FI_Category. Length is:", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.CategoryCode != "00X" {
		t.Error("Invalid CategoryCode found for FI_Category.")
	}
}
