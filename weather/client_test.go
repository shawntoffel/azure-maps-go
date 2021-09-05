package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const TestDataPath = "testdata"

func TestParseCurrentConditionsResponse(t *testing.T) {
	respData, err := loadFileData("current_conditions.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.CurrentConditions("", &CurrentConditionsRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestParseErrorResponse(t *testing.T) {
	respData, err := loadFileData("invalid_coordinates.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusBadRequest)
	defer srv.Close()

	resp, err := w.CurrentConditions("", &CurrentConditionsRequestOptions{})
	if err == nil {
		t.Errorf("expected error response, but was nil")
		t.Errorf("resp content: %+v", resp)
	}
}

func TestParseDailyForecastResponse(t *testing.T) {
	respData, err := loadFileData("daily_forecast.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.DailyForecast("", &DailyForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestParseDailyIndiciesResponse(t *testing.T) {
	respData, err := loadFileData("daily_indicies.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.DailyIndicies("", &DailyIndicesRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}
func TestParseHourlyForecastResponse(t *testing.T) {
	respData, err := loadFileData("hourly_forecast.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.HourlyForecast("", &HourlyForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestParseMinuteForecastResponse(t *testing.T) {
	respData, err := loadFileData("minute_forecast.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.MinuteForecast("", &MinuteForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestParseQuarterDayForecastResponse(t *testing.T) {
	respData, err := loadFileData("quarter_day_forecast.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.QuarterDayForecast("", &QuarterDayForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestParseSevereWeatherAlertsResponse(t *testing.T) {
	respData, err := loadFileData("severe_weather_alerts.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.SevereWeatherAlerts("", &SevereWeatherAlertsRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestParseWeatherAlongRouteResponse(t *testing.T) {
	respData, err := loadFileData("weather_along_route.json")
	if err != nil {
		t.Error(err)
	}

	w, srv := setupClient(respData, http.StatusOK)
	defer srv.Close()

	resp, err := w.WeatherAlongRoute("", &WeatherAlongRouteRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, resp, respData)
}

func TestBuildHttpRequest(t *testing.T) {
	w := New("key")
	w.MSClientId = "test"

	op := CurrentConditionsRequestOptions{
		Format:     "json",
		ApiVersion: "1.0",
		Duration:   intToPtr(2),
		Unit:       "imperial",
	}

	httpReq, err := w.buildHttpRequest(CurrentConditionsEndpoint, "test", op)
	if err != nil {
		t.Error(err)
	}

	urlHave := httpReq.URL.String()
	urlWant := w.BaseUrl + "/currentConditions/json?api-version=1.0&duration=2&query=test&subscription-key=key&unit=imperial"

	if urlHave != urlWant {
		t.Errorf("url mismatch, have: %s, want: %s", urlHave, urlWant)
	}

	headerHave := httpReq.Header.Get("x-ms-client-id")
	headerWant := "test"

	if headerHave != headerWant {
		t.Errorf("header mismatch, have: %s, want: %s", headerHave, headerWant)
	}
}

func loadFileData(filename string) ([]byte, error) {
	return ioutil.ReadFile(TestDataPath + "/" + filename)
}

func setupClient(respData []byte, responseCode int) (Weather, *httptest.Server) {
	server := getMockServer(respData, responseCode)

	w := New("")
	w.BaseUrl = server.URL

	return w, server
}

func getMockServer(data []byte, responseCode int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		fmt.Fprintln(w, string(data))
	}))

	return server
}

func intToPtr(i int) *int {
	return &i
}

func assertEqual(t *testing.T, resp interface{}, expected []byte) {
	data, err := json.Marshal(resp)
	if err != nil {
		t.Error(err)
	}

	buffer := bytes.Buffer{}
	err = json.Compact(&buffer, expected)
	if err != nil {
		t.Error(err)
	}
	compactExpected := buffer.Bytes()

	equal, err := areEqual(data, compactExpected)
	if !equal {
		t.Errorf("reserialized json did not match expected. \nhave: %s\n\nwant: %s", string(data), string(compactExpected))
	}
}

func areEqual(a, b []byte) (bool, error) {
	var aValue interface{}
	var bValue interface{}

	err := json.Unmarshal(a, &aValue)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(b, &bValue)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(aValue, bValue), nil
}
