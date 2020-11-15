package weather

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	apiCalls["IPLocation"] = "/locations/v1/cities/ipaddress"
	apiCalls["currentConditions"] = "/currentconditions/v1/"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (w *Weather) SetLocationIP() {
	const URL = "http://ifconfig.me"
	data, err := http.Get(URL)
	checkErr(err)
	r := bufio.NewReader(data.Body)
	w.IP, _ = r.ReadString('\n')

}

func (w *Weather) GetLocationIP() (ip string, err error) {
	return w.IP, nil
}

func (w *Weather) SetLocationInfo() {
	key := os.Getenv("ACCUWEATHER_KEY")
	ip, err := w.GetLocationIP()
	checkErr(err)
	url := endpoint + apiCalls["IPLocation"] + keyParam + strings.TrimSpace(key) + queryParam + strings.TrimSpace(ip)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	checkErr(err)
	checkCode(resp.StatusCode)

	json.NewDecoder(resp.Body).Decode(w)
}

func (l *Location) GetLocationKey() (key string) {
	return l.Key
}

func (l *Location) GetLocationDetails() (city, country string) {
	return l.LocalizedName, l.Country.LocalizedName

}

func (w *Weather) SetWeather() {
	locationKey := w.GetLocationKey()
	apiKey := os.Getenv("ACCUWEATHER_KEY")

	var c []Weather

	url := endpoint + apiCalls["currentConditions"] + locationKey + keyParam + strings.TrimSpace(apiKey)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	checkErr(err)
	checkCode(resp.StatusCode)

	json.NewDecoder(resp.Body).Decode(&c)

	w.WeatherText = c[0].WeatherText
	w.Temperature.Metric.Value = c[0].Temperature.Metric.Value
}

func (w *Weather) GetWeather() (temperature float64, description string) {
	return w.Temperature.Metric.Value, w.WeatherText
}

func (w *Weather) Init() {
	w.SetLocationIP()
	w.SetLocationInfo()
	w.SetWeather()

}

func checkCode(code int) {
	if code == UNAUTH {
		log.Fatal("[*] There was a problem with you API Key, most likely the limit of requests was reached. Please try again later")
		os.Exit(-1)
	}
	if code == FORBID {
		log.Fatal("[*] Check your API Key is set correctly and try again later")
		os.Exit(-1)
	}
}
