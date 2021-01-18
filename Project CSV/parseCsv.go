package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

func parseCsv(r *csv.Reader) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("BEGIN:VCARD")
		fmt.Println("VERSION:3.0")
		fmt.Println("N:" + record[2] + ";" + record[1] + ";;;")
		fmt.Println("FN:" + record[1] + " " + record[2])
		fmt.Println("ORG:MTC SFO;")
		fmt.Println("NICKNAME:" + record[3])
		fmt.Println("BDAY;X-APPLE-OMIT-YEAR=2021;VALUE=date:" + record[4])
		fmt.Println("NOTE:" + record[5])
		fmt.Println("item1.ADR;TYPE=HOME;TYPE=pref:;;" + record[6] + "\\n" + record[7] + ";" + record[8] + ";" + record[9] + ";United States")
		fmt.Println("item1.X-ABADR:us")
		fmt.Println("TEL;TYPE=CELL;TYPE=pref;TYPE=VOICE:" + record[10])
		fmt.Println("TEL;TYPE=IPHONE;TYPE=CELL;TYPE=VOICE:" + record[11])
		fmt.Println("TEL;TYPE=HOME;TYPE=VOICE:" + record[12])
		fmt.Println("EMAIL;TYPE=HOME;TYPE=pref;TYPE=INTERNET:" + record[13])
		fmt.Println("EMAIL;TYPE=WORK;TYPE=INTERNET:" + record[14])
		fmt.Println("PRODID:-//Mar Thoma Church Of San Francisco.//Registry dated 20210101//EN")
		fmt.Println("END:VCARD")
	}
}
