package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	var timestamp = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	var urlRates = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml?" + timestamp

	fmt.Println("getting data from: " + urlRates)

	if xmlStr, err := getXML(urlRates); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println(xmlStr)

		// TODO: get rates
	}
}

func getXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, errr := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", errr)
	}

	return string(data), nil
}
