package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		Temperature_C    float64 `json:"temp_c"`
		Condition        struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func main() {

	q := "Coimbatore"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=3bceb2ff53f24856889182701252804&q=" + q + "&aqi=no")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Weather api is not available: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var weather Weather

	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	location, country, lastUpdatedEpoch, temperatureC, condition := weather.Location.Name, weather.Location.Country, weather.Current.LastUpdatedEpoch, weather.Current.Temperature_C, weather.Current.Condition.Text

	Location := fmt.Sprintf("%s, %s", location, country)
	color.New(color.BgHiBlue, color.FgBlack).Println(Location)

	date := time.Unix(lastUpdatedEpoch, 0)

	fmt.Print("Last Updated: ")
	color.Red(date.Format("2006-01-02 03:04:05 PM"))

	fmt.Print("Temperature: ")

	temp := fmt.Sprintf("%.2f", temperatureC)
	color.Green(temp + "c")

	fmt.Println("Condition:", condition)

}
