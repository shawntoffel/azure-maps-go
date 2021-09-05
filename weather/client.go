package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const DefaultBaseUrl = "https://atlas.microsoft.com/weather"

type Weather struct {
	BaseUrl         string
	MSClientId      string
	client          *http.Client
	subscriptionKey string
}

func New(subscriptionKey string) Weather {
	return NewWithHttpClient(subscriptionKey, &http.Client{})
}

func NewWithHttpClient(subscriptionKey string, client *http.Client) Weather {
	return Weather{
		BaseUrl:         DefaultBaseUrl,
		client:          client,
		subscriptionKey: subscriptionKey,
	}
}

func (w Weather) CurrentConditions(query string, opts *CurrentConditionsRequestOptions) (*CurrentConditionsResponse, error) {
	resp := CurrentConditionsResponse{}
	err := w.doRequest(
		CurrentConditionsEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) DailyForecast(query string, opts *DailyForecastRequestOptions) (*DailyForecastResponse, error) {
	resp := DailyForecastResponse{}
	err := w.doRequest(
		DailyForecastEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) DailyIndicies(query string, opts *DailyIndicesRequestOptions) (*DailyIndiciesResponse, error) {
	resp := DailyIndiciesResponse{}
	err := w.doRequest(
		DailyIndiciesEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) HourlyForecast(query string, opts *HourlyForecastRequestOptions) (*HourlyForecastResponse, error) {
	resp := HourlyForecastResponse{}
	err := w.doRequest(
		HourlyForecastEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) MinuteForecast(query string, opts *MinuteForecastRequestOptions) (*MinuteForecastResponse, error) {
	resp := MinuteForecastResponse{}
	err := w.doRequest(
		MinuteForecastEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) QuarterDayForecast(query string, opts *QuarterDayForecastRequestOptions) (*QuarterDayForecastResponse, error) {
	resp := QuarterDayForecastResponse{}
	err := w.doRequest(
		QuarterDayForecastEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) SevereWeatherAlerts(query string, opts *SevereWeatherAlertsRequestOptions) (*SevereWeatherAlertsResponse, error) {
	resp := SevereWeatherAlertsResponse{}
	err := w.doRequest(
		SevereWeatherAlertsEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) WeatherAlongRoute(query string, opts *WeatherAlongRouteRequestOptions) (*WeatherAlongRouteResponse, error) {
	resp := WeatherAlongRouteResponse{}
	err := w.doRequest(
		WeatherAlongRouteEndpoint,
		query,
		opts,
		&resp,
	)

	return &resp, err
}

func (w Weather) buildHttpRequest(endpoint string, query string, op Optionable) (*http.Request, error) {
	options := InitializeOptions(op, w.subscriptionKey, w.MSClientId)

	uri := w.BaseUrl + endpoint + "/" + options.Format + "?" + options.Encode(query)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	if options.MSClientId != "" {
		req.Header.Set("x-ms-client-id", options.MSClientId)
	}

	return req, nil
}

func (w Weather) doRequest(endpoint string, query string, op Optionable, resp interface{}) error {
	req, err := w.buildHttpRequest(endpoint, query, op)
	if err != nil {
		return err
	}

	httpResp, err := w.client.Do(req)
	if err != nil {
		return err
	}

	if isErrorResponse(httpResp) {
		errResp := ODataErrorResponse{}

		err = decodeJson(httpResp.Body, &errResp)
		if err != nil {
			return err
		}

		return &errResp.Error
	}

	if httpResp.StatusCode != http.StatusOK {
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
