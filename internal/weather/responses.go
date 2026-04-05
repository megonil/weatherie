package weather

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type CurrentBody struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`

	WindKph    float64 `json:"wind_kph"`
	WindDegree int     `json:"wind_degree"`
	WindDir    string  `json:"wind_dir"`

	PressureMb float64 `json:"pressure_mb"`
	PressureIn float64 `json:"pressure_in"`

	PrecipMm float64 `json:"precip_mm"`
	PrecipIn float64 `json:"precip_in"`

	Humidity int `json:"humidity"`
	Cloud    int `json:"cloud"`

	FeelslikeC float64 `json:"feelslike_c"`
	FeelslikeF float64 `json:"feelslike_f"`

	WindchillC float64 `json:"windchill_c"`
	WindchillF float64 `json:"windchill_f"`

	HeatIndexC float64 `json:"heatindex_c"`
	HeatIndexF float64 `json:"heatindex_f"`

	DewPointC float64 `json:"dewpoint_c"`
	DewPointF float64 `json:"dewpoint_f"`

	VisKm    float64 `json:"vis_km"`
	VisMiles float64 `json:"vis_miles"`

	UV float64 `json:"uv"`

	GustMph float64 `json:"gust_mph"`
	GustKph float64 `json:"gust_kph"`

	ShortRad float64 `json:"short_rad"`
	DiffRad  float64 `json:"diff_rad"`
	DNI      float64 `json:"dni"`
	GTI      float64 `json:"gti"`
}

type CurrentResp struct {
	Location Location    `json:"location"`
	Current  CurrentBody `json:"current"`
}

func (c *CurrentResp) String() string {
	var result strings.Builder
	tw := tabwriter.NewWriter(&result, 0, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "Локация\t%s, %s(%s) %.1f %.1f\n", c.Location.Name, c.Location.Country, c.Location.Region, c.Location.Lat, c.Location.Lon)
	fmt.Fprintf(tw, "Температура:\t%.1f°C, ощущается как %.1f°C\n", c.Current.TempC, c.Current.FeelslikeC)
	fmt.Fprintf(tw, "Влажность:\t%d%%\n", c.Current.Humidity)
	fmt.Fprintf(tw, "Погода:\t%s\n", c.Current.Condition.Text)
	fmt.Fprintf(tw, "Ветер:\t%.1f км/ч %s\n", c.Current.WindKph, c.Current.WindDir)
	fmt.Fprintf(tw, "Последний раз обновлено:\t%s\n", c.Current.LastUpdated)

	tw.Flush()
	return result.String()
}

func (c *CurrentResp) IconURL() string {
	return "http:" + c.Current.Condition.Icon
}
