/*

	1. Temperature (#0080ff) (1)
	2. Pressure (#ff4000) (2)
	3. Humidity (#4000ff) (3)
	4. Air Quality Index (#39e600) (4)

*/
package sensor

import "net/http"
import "math/rand"
import "os"
import "fmt"
import "encoding/json"
import "strconv"

var openWeatherApiKey string = "e1ab28c30ac9fa4f6751261cd5d71c45"
var breezometerApiKey string = "64354dc2e46047b5a2ab0e069d6545b9"

func getSensorData(lat float64, lng float64, typeCode int32) float64 {

	//get current value from openWeatherMap
	response, err := http.Get(getAPIUrl(lat, lng, typeCode))
	if err != nil {
		fmt.Println("Error while getting response from API", err.Error())
		os.Exit(1)
	}

	if typeCode == 1 || typeCode == 2 || typeCode == 3 {
		result := OWAPIResponse{}
		json.NewDecoder(response.Body).Decode(&result)

		if typeCode == 1 {
			return result.List[0].Main.Temp
		} else if typeCode == 2 {
			return result.List[0].Main.Pressure
		} else if typeCode == 3 {
			return float64(result.List[0].Main.Humidity)
		}

	} else if typeCode == 4 {
		result := BMAPIResponse{}
		json.NewDecoder(response.Body).Decode(&result)
		return float64(result.CountryAqi)
	} else {
		panic("Invalid sensor type")
	}

	return 0.0
}

func getSensorColor(typeCode int32) string {

	switch typeCode {
	case 1:
		return "#0080ff"
	case 2:
		return "#ff4000"
	case 3:
		return "#4000ff"
	case 4:
		return "#39e600"
	default:
		panic("Invalid Type Code")
	}
}

func getAPIUrl(lat float64, lng float64, typeCode int32) string {

	var apiUrl string

	if typeCode == 1 || typeCode == 2 || typeCode == 3 {
		apiUrl = "http://api.openweathermap.org/data/2.5/find?lat=" + strconv.FormatFloat(lat, 'f', -1, 64) + "&lon=" + strconv.FormatFloat(lng, 'f', -1, 64) + "&units=imperial" + "&cnt=1&appid=" + openWeatherApiKey
		return apiUrl
	} else if typeCode == 4 {
		apiUrl = "https://api.breezometer.com/baqi/?lat=" + strconv.FormatFloat(lat, 'f', -1, 64) + "&lon=" + strconv.FormatFloat(lng, 'f', -1, 64) + "&key=" + breezometerApiKey
		return apiUrl
	} else {
		panic("Invalid type while forming api url")
	}
}

func getSensorCost(typeCode int32) float64 {
	if typeCode == 1 {
		return getRandomBetweenRange(40, 10) * 0.019
	} else if typeCode == 2 {
		return getRandomBetweenRange(60, 20) * 0.017
	} else if typeCode == 3 {
		return getRandomBetweenRange(20, 5) * 0.013
	} else if typeCode == 4 {
		return getRandomBetweenRange(80, 30) * 0.011
	} else {
		panic("Invalid type code")
	}
}

func getRandomBetweenRange(max int, min int) float64 {
	return float64(rand.Intn(max-min) + min)
}
