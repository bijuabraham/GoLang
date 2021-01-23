package main

import (
	"encoding/csv"
	"fmt"
)

type csvHandle *csv.Reader

func main() {
	f := "household"
	householdIndexiconRequest()
	tmpbody := houseHoldIconRequest()
	fmt.Println(tmpbody)
	tmpH := parseHouseholdJSON(tmpbody)
	tmpH.convertCSV2VCF(f)
}
