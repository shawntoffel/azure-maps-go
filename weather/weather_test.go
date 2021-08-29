package weather

import "testing"

func TestGetCurrentConditions(t *testing.T) {
	t.Skip()
	w := NewWeather()

	req := CurrentConditionsRequest{
		ApiVersion:      "1.0",
		Query:           "38.9575574,-104.5017414",
		Unit:            "imperial",
		SubscriptionKey: "",
	}

	resp, err := w.CurrentConditions(req)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}
