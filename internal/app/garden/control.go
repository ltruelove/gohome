package garden

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
)

var Config config.Configuration

func Init() {
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for range ticker.C {
			t := time.Now()
			fmt.Println(t.Format("2006-01-02 15:04:05"), "Ticker ticked")
			GetSoilReading()
		}
	}()
}

func GetSoilReading() {
	address := fmt.Sprintf("http://%s/status", Config.GardenIp)
	resp, err := http.Get(address)
	if err != nil {
		// handle error
		fmt.Println("Timeout?")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	jsonString := string(body)

	//clear out that annoying line ending
	re := regexp.MustCompile(`\r?\n`)
	jsonString = re.ReplaceAllString(jsonString, " ")

	soilResponse := &data.Garden{}
	soilErr := json.Unmarshal(body, &soilResponse)
	if soilErr != nil {
		errorResponse := "Probably got a bad soil reading"
		fmt.Println(errorResponse)
		return
	}

	reading := fmt.Sprintf("Soil Reading: %d", soilResponse.SoilReading)
	fmt.Println(reading)

	if soilResponse.SoilReading < Config.SoilThreshold {
		StartWater()
	}
}

func StartWater() {
	waterStatusAddress := fmt.Sprintf("http://%s/status", Config.WaterIp)
	waterResp, waterErr := http.Get(waterStatusAddress)
	if waterErr != nil {
		// handle error
		fmt.Println("Timeout?")
		return
	}

	defer waterResp.Body.Close()
	waterBody, waterErr := ioutil.ReadAll(waterResp.Body)
	waterString := string(waterBody)

	//clear out that annoying line ending
	re := regexp.MustCompile(`\r?\n`)
	waterString = re.ReplaceAllString(waterString, " ")

	waterResponse := &data.Water{}
	if err := json.Unmarshal(waterBody, &waterResponse); err != nil {
		fmt.Println("Probably got a bad water reading")
		return
	}

	if waterResponse.Status == "on" {
		return
	}

	waterOnAddress := fmt.Sprintf("http://%s/on", Config.WaterIp)
	_, waterOnErr := http.Get(waterOnAddress)
	if waterOnErr != nil {
		// handle error
		fmt.Println("Water On Timeout?")
		return
	}

	timeChan := time.NewTimer(time.Minute * 5).C
	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
			waterOffAddress := fmt.Sprintf("http://%s/off", Config.WaterIp)
			_, waterOffErr := http.Get(waterOffAddress)
			if waterOffErr != nil {
				// handle error
				fmt.Println("Water Off Timeout?")
				return
			}

			waitChan := time.NewTimer(time.Minute * 5).C
			for {
				select {
				case <-waitChan:
					return
				}
			}
		}
	}
}
