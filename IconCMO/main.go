package main

func main() {
	householdIndexiconRequest()
	tmpbody := houseHoldIconRequest()
	tmpH := parseHouseholdJSON(tmpbody)
	tmpH.saveCSV("household.csv")
}
