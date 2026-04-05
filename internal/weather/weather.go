// Package weather
package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"weatherie/initializers"
)

const base = "http://api.weatherapi.com/v1"
const currentWeatherEndPoint = "current.json"
const ruCode = "ru"

func endpoint(req string, place string) string {
	u, _ := url.Parse(base + "/" + req)

	q := u.Query()
	q.Set("key", initializers.WeatherAPIToken)
	q.Set("q", place)
	q.Set("lang", ruCode)

	u.RawQuery = q.Encode()
	return u.String()
}

func request(req string, place string) ([]byte, error) {
	r, err := http.Get(endpoint(req, place))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Current(place string) (CurrentResp, error) {
	b, err := request(currentWeatherEndPoint, place)
	if err != nil {
		return CurrentResp{}, err
	}

	var current CurrentResp
	err = json.Unmarshal(b, &current)
	if err != nil {
		return CurrentResp{}, err
	}

	return current, nil
}
