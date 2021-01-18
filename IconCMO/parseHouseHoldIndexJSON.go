package main

import (
	"encoding/json"
	"fmt"
)

func parseHouseHoldIndexJSON(jsonString string) {

	type HouseHoldID struct {
		Records   int
		ID        string
		FirstName string
	}

	type TmpHouseHold struct {
		Statistics struct {
			Records int `json:"records"`
			Startat int `json:"startat"`
			Limit   int `json:"limit"`
		} `json:"statistics"`
		HouseHoldIndex []struct {
			ID        string `json:"id"`
			Status    string `json:"status"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"householdindex"`
		Permissions struct {
			Create bool `json:"create"`
			Read   bool `json:"read"`
			Update bool `json:"update"`
			Delete bool `json:"delete"`
		} `json:"permissions"`
		Session string `json:"session"`
		Role    string `json:"role"`
	}

	var tmpH TmpHouseHold
	err := json.Unmarshal([]byte(jsonString), &tmpH)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", tmpH)
	// {Location:London Weather:[{Weather:Drizzle Description:light intensity drizzle}] Temperature:{Temperature:280.32 MinTemperature:279.15 MaxTemperature:281.15}}

	fmt.Println(tmpH.Statistics.Records)
	hhID := HouseHoldID{
		Records:   tmpH.Statistics.Records,
		ID:        tmpH.HouseHoldIndex[0].ID,
		FirstName: tmpH.HouseHoldIndex[0].FirstName,
	}

	fmt.Printf("%+v\n", hhID)
}
