package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const DefaultBaseUrl = "https://atlas.microsoft.com/weather"
const DefaultApiVersion = "1.0"

type Weather struct {
	BaseUrl         string
	ApiVersion      string
	client          *http.Client
	subscriptionKey string
	defaultParams   string
}

func NewWeather(subscriptionKey string) Weather {
	r := Request{
		ApiVersion:      DefaultApiVersion,
		SubscriptionKey: subscriptionKey,
	}

	return Weather{
		BaseUrl:         DefaultBaseUrl,
		ApiVersion:      DefaultApiVersion,
		client:          &http.Client{},
		subscriptionKey: subscriptionKey,
		defaultParams:   r.Encode(),
	}
}

func (w Weather) CurrentConditions(query string, opts *CurrentConditionsRequestOptions) (*CurrentConditionsResponse, error) {
	currentConditionsResponse := &CurrentConditionsResponse{}
	err := w.doRequest(
		CurrentConditionsEndpoint,
		query,
		opts,
		currentConditionsResponse,
	)
	if err != nil {
		return nil, err
	}

	return currentConditionsResponse, nil
}

func (w Weather) DailyForecast(query string, opts *DailyForecastRequestOptions) (*DailyForecastResponse, error) {
	dailyForecastResponse := &DailyForecastResponse{}
	err := w.doRequest(
		DailyForecastEndpoint,
		query,
		opts,
		dailyForecastResponse,
	)
	if err != nil {
		return nil, err
	}

	return dailyForecastResponse, nil
}

func (w Weather) buildBaseUrl(endpoint string, query string) string {
	return w.BaseUrl + endpoint + "/json?query=" + query + "&" + w.defaultParams
}

func (w Weather) doRequest(endpoint string, query string, opts Encodable, resp interface{}) error {
	url := w.buildBaseUrl(endpoint, query)
	if opts != nil {
		url += "&" + opts.Encode()
	}

	httpResp, err := w.client.Get(url)
	if err != nil {
		return err
	}

	if isErrorResponse(httpResp) {
		errResp := &ODataErrorResponse{}
		err = decodeJson(httpResp.Body, errResp)
		if err != nil {
			return err
		}

		return errResp.Error
	} else if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("code: %s", httpResp.Status)
	}

	err = decodeJson(httpResp.Body, &resp)
	if err != nil {
		return err
	}

	return nil
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

func isErrorResponse(r *http.Response) bool {
	return r.StatusCode == http.StatusBadRequest ||
		r.StatusCode == http.StatusUnauthorized ||
		r.StatusCode == http.StatusForbidden ||
		r.StatusCode == http.StatusNotFound ||
		r.StatusCode == http.StatusInternalServerError
}
