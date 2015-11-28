// Package dolstats provides programmatic access to the US Department of Labor
// Health & Safety Datasets API (http://developer.dol.gov/health-and-safety).
package dolstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	apiScheme = "http"
	apiHost   = "api.dol.gov"
	apiPath   = "V1"
)

var (
	APIErrorMsg          = errors.New("The API request returned an error")
	invalidResponseError = errors.New("The HTTP request failed.")
)

type API struct {
	SecretKey string
	Debug     bool
	Filters   map[string]string
	RawData   []byte
	endpoint  url.URL
}

type APIError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// OSHABool is a custom bool type used for properly unmarshaling JSON.
type OSHABool bool

// OSHADate is a custom time.Time type used for properly unmarshaling JSON.
type OSHADate time.Time

// OSHAFloat is a custom float64 type used for properly unmarshaling JSON.
type OSHAFloat float64

// OSHAInt is a custom int type used for properly unmarshaling JSON.
type OSHAInt int

func New(secretKey string) *API {
	return &API{
		SecretKey: secretKey,
		Filters: map[string]string{
			"top": "100",
		},
		endpoint: url.URL{
			Scheme: apiScheme,
			Host:   apiHost,
		},
	}
}

func (api *API) AddFilter(name string, value string) {
	api.Filters[name] = value
}

func (api *API) buildPath(queryEndpoint string) {
	queryVals := url.Values{}
	queryVals.Set("KEY", api.SecretKey)

	for k, v := range api.Filters {
		queryVals.Add(fmt.Sprintf("$%s", k), v)
	}

	api.endpoint.Path = fmt.Sprintf("%s/%s", apiPath, queryEndpoint)
	api.endpoint.RawQuery = queryVals.Encode()
}

func (api *API) doRequest() error {
	if api.Debug {
		log.Printf("API endpoint URL: %s", api.endpoint.String())
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", api.endpoint.String(), nil)
	if err != nil {
		return err
	}

	// Only accept JSON responses. The default response format is XML.
	req.Header.Add("ACCEPT", "application/json")
	if api.Debug {
		log.Printf("HTTP request headers: %v", req.Header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if api.Debug {
		log.Printf("Response body: %v", string(body))
	}

	apiErr := unmarshalErrorResponse(body)
	if apiErr != nil {
		return apiErr
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s HTTP status code returned was: %d.", invalidResponseError, resp.StatusCode)
	}

	api.RawData = body
	return nil
}

// unmarshalErrorResponse attempts to unmarshal an API error message if one
// exists. If an error is found, the details are returned.
func unmarshalErrorResponse(data []byte) error {
	var a APIError

	err := json.Unmarshal(data, &a)
	if err == nil && len(a.Error) > 0 {
		return fmt.Errorf("%s The error message was: %+s", APIErrorMsg, a.Error)
	}

	return nil
}

// An implementation of the UnmarshalJSONinterface for boolean strings.
func (ob *OSHABool) UnmarshalJSON(b []byte) error {
	boolStr := string(b[1 : len(b)-1])
	boolVal := boolStr == "Y" || boolStr == "Yes"

	*ob = OSHABool(boolVal)

	return nil
}

// An implementation of the UnmarhsalJSON interface for unix time stamps.
func (od *OSHADate) UnmarshalJSON(b []byte) error {
	timeStr := string(b[8 : len(b)-7])

	unixInt, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return err
	}

	unixTime := time.Unix(unixInt, 0)
	*od = OSHADate(unixTime.UTC())

	return nil
}

func (oi *OSHAInt) UnmarshalJSON(b []byte) error {
	// In order to prevent unmarshaling errors, If the field has a value of
	// "null", return 0.
	if string(b) == "null" {
		*oi = OSHAInt(0)

		return nil
	}

	istr := string(b[1 : len(b)-1])

	i, err := strconv.Atoi(istr)
	if err != nil {
		return err
	}

	*oi = OSHAInt(i)

	return nil
}
