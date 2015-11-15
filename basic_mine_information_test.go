package dolstats

import "testing"

func TestUnmarshalBasicMine(t *testing.T) {
	mock, err := loadMock("./testdata/basic_mine_information/msha_address_of_record.json")
	if err != nil {
		t.Error("Error loading BasicMine mock data. Error was:", err)
	}

	a, err := unmarshalAddressRecord(mock)
	if err != nil {
		t.Error("Error unmarshaling BasicMine. Error was:", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid Data length. Length is:", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid ID value returned. Value was:", resOne.ID)
	}

	if resOne.MineID != 100003 {
		t.Error("Invalid MineID value returned. value was:", resOne.MineID)
	}

	if resOne.Zip != 35040 {
		t.Error("Invalid Zip value returned. value was:", resOne.Zip)
	}
}
