// Package weather
package weather

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"weatherie/initializers"
)

const base = "http://api.weatherapi.com/v1"
const currentWeatherEndPoint = "current.json"

func endpoint(req string, place string) string {
	u, _ := url.Parse(base + "/" + req)

	q := u.Query()
	q.Set("key", initializers.WeatherAPIToken)
	q.Set("q", place)

	u.RawQuery = q.Encode()
	return u.String()
}

func weatherRequest(req string, place string) (string, error) {
	r, err := http.Get(endpoint(req, place))
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Current(place string) string {
	r, err := weatherRequest(currentWeatherEndPoint, place)
	if err != nil {
		log.Fatal("error during minecraft: %v", err)
	}

	return r
}
