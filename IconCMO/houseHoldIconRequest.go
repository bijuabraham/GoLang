package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HouseHolds struct {
	Title     string
	FirstName string
	LastName  string
	Address1  string
	Address2  string
	City      string
	State     string
	Zip       string
	Email     string
	Phone     string
	Phones    string
	Emails    string
	Picture   string
	MemberID  string
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

func houseHoldIconRequest() string {
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
			"Section": "households", 
			"Filters": {
				"startAt": 0,
				"limit": 10 
			},
			"Sort": {
				"last_name": "ascending",
				"first_name": "ascending"
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
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func parseHouseholdJSON(jsonString string) TmpHouseHold {

	var tmpH TmpHouseHold
	err := json.Unmarshal([]byte(jsonString), &tmpH)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", tmpH)
	// {Location:London Weather:[{Weather:Drizzle Description:light intensity drizzle}] Temperature:{Temperature:280.32 MinTemperature:279.15 MaxTemperature:281.15}}
	/*
		hh := HouseHolds{
			Title:     tmpH.HouseHolds[0].Title,
			FirstName: tmpH.HouseHolds[0].FirstName,
			LastName:  tmpH.HouseHolds[0].LastName,
			Address1:  tmpH.HouseHolds[0].Address1,
			Address2:  tmpH.HouseHolds[0].Address2,
			City:      tmpH.HouseHolds[0].City,
			State:     tmpH.HouseHolds[0].State,
			Zip:       tmpH.HouseHolds[0].Zip,
			Email:     tmpH.HouseHolds[0].Email,
			Phone:     tmpH.HouseHolds[0].Phone,
			Picture:   tmpH.HouseHolds[0].Picture,
		}
		fmt.Printf("%+v\n", hh)
	*/
	return tmpH
}

func (tmpH TmpHouseHold) convertCSV2VCF(f string) {
	csvfile := f + ".csv"
	dat, err := os.Create(csvfile)
	if err != nil {
		log.Fatal(err)
	}
	for i := range tmpH.HouseHolds {
		s := `"` + tmpH.HouseHolds[i].Title + `","` + tmpH.HouseHolds[i].FirstName + `","` + tmpH.HouseHolds[i].LastName + `","` + tmpH.HouseHolds[i].Address1 + `","` + tmpH.HouseHolds[i].Address2 + `","` + tmpH.HouseHolds[i].City + `","` + tmpH.HouseHolds[i].State + `","` + tmpH.HouseHolds[i].Zip + `","` + tmpH.HouseHolds[i].Email + `","` + tmpH.HouseHolds[i].Phone + `","` + tmpH.HouseHolds[i].Picture + `"`
		io.WriteString(dat, s)
		io.WriteString(dat, "\n")
		//fmt.Printf("String written %v", s)
	}
	//dat.Sync()
	vcffile := f + ".vcf"
	dat, err = os.Open(csvfile)
	if err != nil {
		log.Fatal(err)
	}
	csvHandle := csv.NewReader(dat)
	vdat, verr := os.Create(vcffile)
	if verr != nil {
		log.Fatal(verr)
	}
	for {
		record, rerr := csvHandle.Read()
		if rerr == io.EOF {
			fmt.Println("Broken")
			break
		}
		if rerr != nil {
			log.Fatal(rerr)
		}

		io.WriteString(vdat, "BEGIN:VCARD\n")
		io.WriteString(vdat, "VERSION:3.0\n")
		io.WriteString(vdat, "N:"+record[2]+";"+record[1]+";;;\n")
		io.WriteString(vdat, "FN:"+record[1]+" "+record[2]+"\n")
		io.WriteString(vdat, "ORG:MTC SFO;\n")
		io.WriteString(vdat, "BDAY;X-APPLE-OMIT-YEAR=2021;VALUE=date:"+record[3]+"\n")
		io.WriteString(vdat, "NOTE:"+record[4]+"\n")
		io.WriteString(vdat, "item1.ADR;TYPE=HOME;TYPE=pref:;;"+record[6]+"\\n"+record[7]+";"+record[8]+";"+record[9]+";United States\n")
		io.WriteString(vdat, "item1.X-ABADR:us\n")
		io.WriteString(vdat, "TEL;TYPE=CELL;TYPE=pref;TYPE=VOICE:"+record[10]+"\n")
		io.WriteString(vdat, "TEL;TYPE=IPHONE;TYPE=CELL;TYPE=VOICE:"+record[11]+"\n")
		io.WriteString(vdat, "TEL;TYPE=HOME;TYPE=VOICE:"+record[12]+"\n")
		io.WriteString(vdat, "EMAIL;TYPE=HOME;TYPE=pref;TYPE=INTERNET:"+record[13]+"\n")
		io.WriteString(vdat, "EMAIL;TYPE=WORK;TYPE=INTERNET:"+record[14]+"\n")
		io.WriteString(vdat, "PRODID:-//Mar Thoma Church Of San Francisco.//Registry dated 20210101//EN\n")
		io.WriteString(vdat, "END:VCARD\n")
	}
	//vdat.Sync()
}
