package weather

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const DefaultBaseUrl = "https://atlas.microsoft.com/weather"

type Weather struct {
	BaseUrl string
	Client  *http.Client
}

func NewWeather() Weather {
	return Weather{
		BaseUrl: DefaultBaseUrl,
		Client:  &http.Client{},
	}
}

func (w Weather) CurrentConditions(r CurrentConditionsRequest) (*CurrentConditionsResponse, error) {
	url := w.BaseUrl + CurrentConditions + "/json?" + r.Encode()
	resp, err := w.Client.Get(url)
	if err != nil {
		return nil, err
	}

	ccResp := &CurrentConditionsResponse{}
	err = decodeJson(resp.Body, ccResp)
	if err != nil {
		return nil, err
	}

	return ccResp, nil
}

func decodeJson(body io.Reader, into interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if b == nil || len(b) < 1 {
		return nil
	}

	return json.Unmarshal(b, &into)
}
