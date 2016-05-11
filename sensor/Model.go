package sensor

type SensorListReq struct {
	List []string `json:"list"`
}

type SensorListRes struct {
	Result []SensorDataRes `json:"result"`
}

type SensorDataRes struct {
	ID         string  `json:"id" bson:"_id"`
	Type       string  `json:"type"`
	TypeCode   int32   `json:"typecode"`
	Value      float64 `json:"value"`
	Location   string  `json:"location"`
	Color      string  `json:"color"`
	Status     bool    `json:"status"`
	Cost       float64 `json:"cost"`
	Coordinate struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinate"`
	History []float64 `json:"history"`
}

type SensorValueRes struct {
	Value   float64   `json:"value"`
	History []float64 `json:"history"`
}

type AddSensorReq struct {
	Location   string `json:"location"`
	Type       string `json:"type"`
	TypeCode   int32  `json:"typecode"`
	Coordinate struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinate"`
}

type OWAPIResponse struct {
	Message string `json:"message"`
	Cod     string `json:"cod"`
	Count   int    `json:"count"`
	List    []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
		} `json:"main"`
		Dt   int `json:"dt"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Country string `json:"country"`
		} `json:"sys"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"list"`
}

type BMAPIResponse struct {
	Datetime               string `json:"datetime"`
	CountryName            string `json:"country_name"`
	BreezometerAqi         int    `json:"breezometer_aqi"`
	BreezometerColor       string `json:"breezometer_color"`
	BreezometerDescription string `json:"breezometer_description"`
	CountryAqi             int    `json:"country_aqi"`
	CountryAqiPrefix       string `json:"country_aqi_prefix"`
	CountryColor           string `json:"country_color"`
	CountryDescription     string `json:"country_description"`
	DataValid              bool   `json:"data_valid"`
	KeyValid               bool   `json:"key_valid"`
	RandomRecommendations  struct {
		Children string `json:"children"`
		Sport    string `json:"sport"`
		Health   string `json:"health"`
		Inside   string `json:"inside"`
		Outside  string `json:"outside"`
	} `json:"random_recommendations"`
	DominantPollutantCanonicalName string `json:"dominant_pollutant_canonical_name"`
	DominantPollutantDescription   string `json:"dominant_pollutant_description"`
	DominantPollutantText          struct {
		Main    string `json:"main"`
		Effects string `json:"effects"`
		Causes  string `json:"causes"`
	} `json:"dominant_pollutant_text"`
}
