package main

import (
	"log"
	"encoding/csv"
	"os"
)

func readCsv() *csv.Reader{
	csvFileName := "test.csv"
	dat,err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	csvHandle := csv.NewReader(dat)
	return csvHandle
}