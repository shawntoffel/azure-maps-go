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

type WeatherClient struct {
	BaseUrl         string
	ApiVersion      string
	client          *http.Client
	subscriptionKey string
	defaultParams   string
}

func NewWeatherClient(subscriptionKey string) WeatherClient {
	r := Request{
		ApiVersion:      DefaultApiVersion,
		SubscriptionKey: subscriptionKey,
	}

	return WeatherClient{
		BaseUrl:         DefaultBaseUrl,
		ApiVersion:      DefaultApiVersion,
		client:          &http.Client{},
		subscriptionKey: subscriptionKey,
		defaultParams:   r.Encode(),
	}
}

func (w WeatherClient) CurrentConditions(query string, opts *CurrentConditionsRequestOptions) (*CurrentConditionsResponse, error) {
	resp := &CurrentConditionsResponse{}
	err := w.doRequest(
		CurrentConditionsEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) DailyForecast(query string, opts *DailyForecastRequestOptions) (*DailyForecastResponse, error) {
	resp := &DailyForecastResponse{}
	err := w.doRequest(
		DailyForecastEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) DailyIndicies(query string, opts *DailyIndicesRequestOptions) (*DailyIndiciesResponse, error) {
	resp := &DailyIndiciesResponse{}
	err := w.doRequest(
		DailyIndiciesEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) HourlyForecast(query string, opts *HourlyForecastRequestOptions) (*HourlyForecastResponse, error) {
	resp := &HourlyForecastResponse{}
	err := w.doRequest(
		HourlyForecastEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) MinuteForecast(query string, opts *MinuteForecastRequestOptions) (*MinuteForecastResponse, error) {
	resp := &MinuteForecastResponse{}
	err := w.doRequest(
		MinuteForecastEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) QuarterDayForecast(query string, opts *QuarterDayForecastRequestOptions) (*QuarterDayForecastResponse, error) {
	resp := &QuarterDayForecastResponse{}
	err := w.doRequest(
		QuarterDayForecastEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) SevereWeatherAlerts(query string, opts *SevereWeatherAlertsRequestOptions) (*SevereWeatherAlertsResponse, error) {
	resp := &SevereWeatherAlertsResponse{}
	err := w.doRequest(
		SevereWeatherAlertsEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) WeatherAlongRoute(query string, opts *WeatherAlongRouteRequestOptions) (*WeatherAlongRouteResponse, error) {
	resp := &WeatherAlongRouteResponse{}
	err := w.doRequest(
		WeatherAlongRouteEndpoint,
		query,
		opts,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w WeatherClient) buildBaseUrl(endpoint string, query string) string {
	return w.BaseUrl + endpoint + "/json?query=" + query + "&" + w.defaultParams
}

func (w WeatherClient) doRequest(endpoint string, query string, opts Encodable, resp interface{}) error {
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
