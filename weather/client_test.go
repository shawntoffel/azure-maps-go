package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseCurrentConditionsResponse(t *testing.T) {
	w, srv := setupClient("current_conditions.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.CurrentConditions("", &CurrentConditionsRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	want := 1
	have := len(resp.Results)
	if want != have {
		t.Errorf("want %d, have %d", want, have)
	}

	t.Logf("%+v", resp)
}

func TestParseErrorResponse(t *testing.T) {
	w, srv := setupClient("invalid_coordinates.json", http.StatusBadRequest)
	defer srv.Close()

	resp, err := w.CurrentConditions("", &CurrentConditionsRequestOptions{})
	if err == nil {
		t.Errorf("expected error response, but was nil")
	}

	t.Logf("%+v", resp)
}

func TestParseDailyForecastResponse(t *testing.T) {
	w, srv := setupClient("daily_forecast.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.DailyForecast("", &DailyForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	want := 1
	have := len(resp.Forecasts)
	if want != have {
		t.Errorf("want %d, have %d", want, have)
	}

	t.Logf("%+v", resp)
}

func TestParseDailyIndiciesResponse(t *testing.T) {
	w, srv := setupClient("daily_indicies.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.DailyIndicies("", &DailyIndicesRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	want := 19
	have := len(resp.Results)
	if want != have {
		t.Errorf("want %d, have %d", want, have)
	}

	t.Logf("%+v", resp)
}
func TestParseHourlyForecastResponse(t *testing.T) {
	w, srv := setupClient("hourly_forecast.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.HourlyForecast("", &HourlyForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestParseMinuteForecastResponse(t *testing.T) {
	w, srv := setupClient("minute_forecast.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.MinuteForecast("", &MinuteForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestParseQuarterDayForecastResponse(t *testing.T) {
	w, srv := setupClient("quarter_day_forecast.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.QuarterDayForecast("", &QuarterDayForecastRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestParseSevereWeatherAlertsResponse(t *testing.T) {
	w, srv := setupClient("severe_weather_alerts.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.SevereWeatherAlerts("", &SevereWeatherAlertsRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestParseWeatherAlongRouteResponse(t *testing.T) {
	w, srv := setupClient("weather_along_route.json", http.StatusOK)
	defer srv.Close()

	resp, err := w.WeatherAlongRoute("", &WeatherAlongRouteRequestOptions{})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func setupClient(filename string, responseCode int) (WeatherClient, *httptest.Server) {
	server := getMockServerWithFileData(filename, responseCode)

	w := NewWeatherClient("")
	w.BaseUrl = server.URL

	return w, server
}

func getMockServerWithFileData(filename string, responseCode int) *httptest.Server {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)

	return getMockServer(string(bytes), responseCode)
}

func getMockServer(data string, responseCode int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		fmt.Fprintln(w, data)
	}))

	return server
}

func intToPtr(i int) *int {
	return &i
}
