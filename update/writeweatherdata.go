package update

import (
	"fmt"
	"netatmo/utils"
)

const (
	outdoorModuleType = "NAModule1"
)

func WriteWeatherDataToFile(weatherData WeatherData, outputFile string) error {

	tempIndoor := weatherData.Body.Devices[0].DashboardData.Temperature

	var outdoorModule WeatherDataModule

	for _, module := range weatherData.Body.Devices[0].Modules {
		if module.Type == outdoorModuleType {
			outdoorModule = module
		}
	}
	tempOutdoor := outdoorModule.DashboardData.Temperature

	formattedString := fmt.Sprintf("🏠%.1f° 🌲%.1f°", tempIndoor, tempOutdoor)
	return utils.WriteToDestinationPath(outputFile, []byte(formattedString))
}
