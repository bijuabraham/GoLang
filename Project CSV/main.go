package main

func main(){
	iconRequest()
	csvHandle := readCsv()
	parseCsv(csvHandle)
}