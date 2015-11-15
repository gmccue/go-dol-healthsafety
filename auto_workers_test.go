package dolstats

import "testing"

func TestUnmarshalAutoWorkers(t *testing.T) {
	mock, err := loadMock("./testdata/auto_workers/car_closed_repurposed.json")
	if err != nil {
		t.Error("Error loading AutoWorkers mock data. Error was: ", err)
	}

	a, err := unmarshalCarClosed(mock)
	if err != nil {
		t.Error("Error unmarshaling AutoWorkers. Error was: ", err)
	}

	if len(a.Data) != 2 {
		t.Error("Invalid Data length. Length is: ", len(a.Data))
	}

	resOne := a.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid ID value returned. Value was: ", resOne.ID)
	}
}
