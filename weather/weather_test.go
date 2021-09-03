package weather

import (
	"os"
	"testing"
)

var subscriptionKey = os.Getenv("AZ_MAPS_SUBSCRIPTION_KEY")

func TestGetCurrentConditions(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &CurrentConditionsRequestOptions{
		Unit: "imperial",
	}

	resp, err := w.CurrentConditions("38.9575574,-104.5017414", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestGetCurrentConditions2(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &CurrentConditionsRequestOptions{
		Unit: "imperial",
	}

	resp, err := w.CurrentConditions("invalid-coords-test", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestGetDailyForecast(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &DailyForecastRequestOptions{
		Unit: "imperial",
	}

	resp, err := w.DailyForecast("38.9575574,-104.5017414", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestGetDailyIndicies(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &DailyIndicesRequestOptions{
		IndexId: ptr(31),
	}

	resp, err := w.DailyIndicies("38.9575574,-104.5017414", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}
func TestGetHourlyForecast(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &HourlyForecastRequestOptions{
		Unit: "imperial",
	}

	resp, err := w.HourlyForecast("38.9575574,-104.5017414", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestGetMinuteForecast(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &MinuteForecastRequestOptions{
		Interval: ptr(1),
	}

	resp, err := w.MinuteForecast("38.9575574,-104.5017414", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestGetQuarterDayForecast(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &QuarterDayForecastRequestOptions{}

	resp, err := w.QuarterDayForecast("38.9575574,-104.5017414", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func TestGetSevereWeatherAlerts(t *testing.T) {
	t.Skip()
	w := NewWeather(subscriptionKey)

	opts := &SevereWeatherAlertsRequestOptions{}

	resp, err := w.SevereWeatherAlerts("39.18722,-106.81835", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}

func ptr(i int) *int {
	return &i
}
