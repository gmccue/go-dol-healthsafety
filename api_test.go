package dolstats

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var (
	api        *API
	testSecret = "123abc"
)

func loadMock(path string) ([]byte, error) {
	mock, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return mock, nil
}

func TestMain(m *testing.M) {
	flag.Parse()

	api = New(testSecret)

	os.Exit(m.Run())
}

func TestAddFilter(t *testing.T) {
	api.AddFilter("top", "5")

	if api.Filters["top"] != "5" {
		t.Error("Invalid filter value. Value was:", api.Filters["top"])
	}

	api.Filters = map[string]string{}
}

func TestBuildPath(t *testing.T) {
	api.AddFilter("skip", "10")
	api.buildPath("test")

	if api.endpoint.Path != fmt.Sprintf("%s/test", apiPath) {
		t.Error("Invalid API endpoint built. Value was:", api.endpoint.Path)
	}

	if api.endpoint.RawQuery != fmt.Sprintf("%sskip=10&KEY=%s", "%24", testSecret) {
		t.Error("Invalid API RawQuery built. Value was:", api.endpoint.RawQuery)
	}
}

func TestUnmarshalOSHADate(t *testing.T) {
	dateJSON := `
	{
		"TestDate":"\/Date(285811200000)\/"
	}
	`

	OSHADateInstance := struct {
		TestDate OSHADate
	}{}

	err := json.Unmarshal([]byte(dateJSON), &OSHADateInstance)
	if err != nil {
		t.Error("Error unmarshaling OSHADate. Error was:", err)
	}

	date := time.Time(OSHADateInstance.TestDate)

	if date.Month() != 1 {
		t.Error("Invalid Month() value. Value was:", date.Month())
	}

	if date.Day() != 22 {
		t.Error("Invalid Day() value. Value was:", date.Day())
	}

	if date.Year() != 1979 {
		t.Error("Invalid Year() value. Value was:", date.Year())
	}
}
