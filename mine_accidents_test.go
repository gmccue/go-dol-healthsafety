package dolstats

import "testing"

func TestUnmarshalMineAccident(t *testing.T) {
	mock, err := loadMock("./testdata/mine_accidents/mine_accidents.json")
	if err != nil {
		t.Error("Error loading MineAccident mock. Error was:", err)
	}

	ma, err := unmarshalMineAccident(mock)
	if err != nil {
		t.Error("Error unmarshaling MineAccident data. Error was:", err)
	}

	resOne := ma.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid MineAccident ID. Value was:", resOne.ID)
	}

	if resOne.CalendarYear != 2007 {
		t.Error("Invalid MineAccident CalendarYear. Value was:", resOne.CalendarYear)
	}
}
