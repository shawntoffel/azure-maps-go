package weather

import "testing"

func TestGetCurrentConditions(t *testing.T) {
	t.Skip()
	w := NewWeather("")

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
	w := NewWeather("")

	opts := &CurrentConditionsRequestOptions{
		Unit: "imperial",
	}

	resp, err := w.CurrentConditions("invalid-coords-test", opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", resp)
}
