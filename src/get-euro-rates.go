package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// XML defined structures
type Envelope struct {
	Cube []Items `xml:"Cube>Cube>Cube"`
}

type Items struct {
	Curency string `xml:"currency,attr"`
	Rate    string `xml:"rate,attr"`
}

func main() {

	urlRates := "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	// no cache
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	urlRates = urlRates + "?" + timestamp
	fmt.Println("data from: " + urlRates + "\n")

	xmlDoc := getXML(urlRates)
	rssFeed := parseXML(xmlDoc)

	for _, item := range rssFeed.Cube {
		fmt.Println("Country Code: " + item.Curency + " \t Rate: " + item.Rate)
		// do whatever you want with the data, add to a table for example
	}
}

func getXML(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("unable to GET '%s': %s", url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("unable to read body '%s': %s", url, err)
	}

	return body
}

func parseXML(xmlDoc []byte) Envelope {

	envelope := Envelope{}

	reader := bytes.NewReader(xmlDoc)
	decoder := xml.NewDecoder(reader)

	if err := decoder.Decode(&envelope); err != nil {
		log.Printf("unable to parse XML '%s':\n%s", err, xmlDoc)
	}

	return envelope
}
