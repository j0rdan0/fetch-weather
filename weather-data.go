package weather

const (
	endpoint   = "http://dataservice.accuweather.com"
	keyParam   = "?apikey="
	queryParam = "&q="
	UNAUTH     = 503
	FORBID     = 401
)

var apiCalls = make(map[string]string)

type Location struct {
	LocalizedName string `"json:LocalizedName"`
	Country       struct {
		LocalizedName string `"json:LocalizedName"`
	} `"json:Country"`
	Key string `"json:Key"`
	IP  string `"json:IP"`
}

type Weather struct {
	WeatherText string `json:"WeatherText"`
	Temperature struct {
		Metric struct {
			Value float64 `json:"Value"`
		} `json:"Metric"`
	} `json:"Temperature"`
	Location
}
