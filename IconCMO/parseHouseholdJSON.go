package main

import (
	"encoding/json"
	"fmt"
)

func parseHouseholdJSON(jsonString string) {

	type HouseHolds struct {
		FirstName string
		LastName  string
		Address1  string
		Title     string
	}

	type TmpHouseHold struct {
		Statistics struct {
			Records int `json:"records"`
			Startat int `json:"startat"`
			Limit   int `json:"limit"`
		} `json:"statistics"`
		HouseHolds []struct {
			ID                 string `json:"id"`
			DateChanged        string `json:"date_changed"`
			Status             string `json:"status"`
			StatusDate         string `json:"status_date"`
			Title              string `json:"title"`
			FirstName          string `json:"first_name"`
			LastName           string `json:"last_name"`
			MailTo             string `json:"mail_to"`
			Address1           string `json:"address_1"`
			Address2           string `json:"address_2"`
			City               string `json:"city"`
			State              string `json:"state"`
			Zip                string `json:"zip"`
			Country            string `json:"country"`
			CarrierRoute       string `json:"carrier_route"`
			Email              string `json:"email"`
			EmailUnlisted      bool   `json:"email_unlisted"`
			Permission         bool   `json:"permission"`
			IncludeInDirectory bool   `json:"include_address_in_directory"`
			Phone              string `json:"phone"`
			PhoneUnlisted      bool   `json:"phone_unlisted"`
			Phones             string `json:"phones"`
			Emails             string `json:"emails"`
			Members            []struct {
				MemberID string `json:"id"`
			} `json:"members"`
			UserDefined1 string `json:"user_defined_1"`
			UserDefined2 string `json:"user_defined_2"`
			DDPatron     string `json:"dd_patron"`
			Picture      string `json:"picture"`
			Thumbnail    string `json:"thumbnail"`
		} `json:"households"`
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

	hh := HouseHolds{
		FirstName: tmpH.HouseHolds[0].FirstName,
		LastName:  tmpH.HouseHolds[0].LastName,
		Address1:  tmpH.HouseHolds[0].Address1,
		Title:     tmpH.HouseHolds[0].Title,
	}

	fmt.Printf("%+v\n", hh)
}
