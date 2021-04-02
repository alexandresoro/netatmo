package update

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/alexandresoro/netatmo/netatmoApi"
)

type WeatherDataDashboardData struct {
	TimeUtc          int     `json:"time_utc"`
	Temperature      float64 `json:"Temperature"`
	CO2              int     `json:"CO2"`
	Humidity         int     `json:"Humidity"`
	Noise            int     `json:"Noise"`
	Pressure         float64 `json:"Pressure"`
	AbsolutePressure float64 `json:"AbsolutePressure"`
	MinTemp          float64 `json:"min_temp"`
	MaxTemp          float64 `json:"max_temp"`
	DateMaxTemp      int     `json:"date_max_temp"`
	DateMinTemp      int     `json:"date_min_temp"`
	TempTrend        string  `json:"temp_trend"`
	PressureTrend    string  `json:"pressure_trend"`
}

type WeatherDataModule struct {
	Type          string                   `json:"type"`
	DashboardData WeatherDataDashboardData `json:"dashboard_data"`
}

type WeatherDataDevice struct {
	Type          string                   `json:"type"`
	DashboardData WeatherDataDashboardData `json:"dashboard_data"`
	Modules       []WeatherDataModule      `json:"modules"`
}

type WeatherDataUser struct {
	Mail string `json:"mail"`
}

type WeatherDataBody struct {
	Devices []WeatherDataDevice `json:"devices"`
	User    WeatherDataUser     `json:"user"`
}

type WeatherData struct {
	Body       WeatherDataBody `json:"body"`
	Status     string          `json:"status"`
	TimeExec   float64         `json:"time_exec"`
	TimeServer int             `json:"time_server"`
}

func GetWeatherData(accessToken string, deviceMacAddress string) (WeatherData, error) {

	req, err := http.NewRequest("GET", netatmoApi.NetatmoApiUrl+netatmoApi.ApiPath+netatmoApi.GetStationsEndpoint, nil)
	if err != nil {
		return WeatherData{}, nil
	}

	bearer := "Bearer " + accessToken
	req.Header.Add("Authorization", bearer)

	queryParams := req.URL.Query()
	queryParams.Add("device_id", deviceMacAddress)
	queryParams.Add("get_favorites", strconv.FormatBool(false))

	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return WeatherData{}, err
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return WeatherData{}, err
	}

	weatherData := WeatherData{}

	err = json.Unmarshal(respBody, &weatherData)
	if err != nil {
		return WeatherData{}, err
	}
	return weatherData, nil
}
