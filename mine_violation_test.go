package dolstats

import "testing"

func TestUnmarshalMineViolations(t *testing.T) {
	mock, err := loadMock("./testdata/mine_violation/msha_violations.json")
	if err != nil {
		t.Error("Error loading MineViolation mock. Error was:", err)
	}

	mv, err := unmarshalMineViolations(mock)
	if err != nil {
		t.Error("Error unmarshaling MineViolation. Error was:", err)
	}

	resOne := mv.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid MineViolation ID. Value was:", resOne.ID)
	}

	if resOne.CalendarYear != 2000 {
		t.Error("Invalid MineViolation CalendarYear. Value was:", resOne.CalendarYear)
	}

	if resOne.NumAffected != 2 {
		t.Error("Invalid MineViolation NumAffected. Value was:", resOne.NumAffected)
	}

	if resOne.AmountDue != 55 {
		t.Error("Invalid MineViolation AmountDue. Value was:", resOne.AmountDue)
	}
}
