package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/j0rdan0/weather"
)

func main() {

	w := new(weather.Weather)
	w.Init()
	displayWeather(w)

}

func displayWeather(w *weather.Weather) {
	degree := '°'
	city, country := w.GetLocationDetails()
	temp, description := w.GetWeather()
	color.Set(color.FgBlue)
	fmt.Printf("Weather Fetcher v. 1\n")
	fmt.Printf("----\n")

	color.Set(color.FgRed)
	fmt.Printf("City:           ")
	color.Unset()
	fmt.Printf("%s\n", city)

	color.Set(color.FgRed)
	fmt.Printf("Country:        ")
	color.Unset()
	fmt.Printf("%s\n", country)

	fmt.Printf("\n")

	color.Set(color.FgRed)
	fmt.Printf("Temperature:     ")
	color.Unset()
	fmt.Printf("%d %cC\n", int(temp), degree)

	color.Set(color.FgRed)
	fmt.Printf("Description:     ")
	color.Unset()
	fmt.Printf("%s\n", description)

	fmt.Printf("\n")

}
