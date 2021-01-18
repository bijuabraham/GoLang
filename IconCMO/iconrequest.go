package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func iconRequest() {
	url := "https://secure1.iconcmo.com/api/"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{ 
		"Auth": { 
			"Phone": "4082421928", 
			"Username": "MarThoAPI", 
			"Password": "ruby1234" 
		}, 
		"Request": { 
			"Module": "membership", 
			"Section": "householdindex", 
			"Sort": {
				"last_name": "ascending",
				"first_name": "ascending"
			}
			,"Filters": {
				"startAt": 0,
				"limit": 10 
			}
		} 
	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	parseHouseHoldIndexJSON(string(body))
}
