package dolstats

import "testing"

func TestUnmarshalDirectRead(t *testing.T) {
	mock, err := loadMock("./testdata/gulf_oil_spill/osha_direct_read_sampling.json")
	if err != nil {
		t.Error("Error loading DirectRead data mock. Error was:", err)
	}

	dr, err := unmarshalDirectRead(mock)
	if err != nil {
		t.Error("Error unmarshaling DirectRead data. Error was:", err)
	}

	if len(dr.Data) != 2 {
		t.Error("Invalid data DirectRead data length. Length was:", len(dr.Data))
	}

	resOne := dr.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid DirectRead ID. Value was:", resOne.ID)
	}

	if resOne.SampleNumber != 10039 {
		t.Error("Invalid DirectRead SampleNumber. Value was:", resOne.SampleNumber)
	}
}

func TestUnmarshalNoiseReport(t *testing.T) {
	mock, err := loadMock("./testdata/gulf_oil_spill/osha_noise_report.json")
	if err != nil {
		t.Error("Error loading DirectRead data mock. Error was:", err)
	}

	nr, err := unmarshalNoiseReport(mock)
	if err != nil {
		t.Error("Error unmarshaling NoiseReport data. Error was:", err)
	}

	resOne := nr.Data[0]

	if resOne.ID != 1 {
		t.Error("Invalid NoiseReport ID. Value was:", resOne.ID)
	}

	if resOne.ExposureLevel != 88 {
		t.Error("Invalid NoiseReport ExposureLevel. Value was:", resOne.ExposureLevel)
	}
}
