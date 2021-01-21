package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func parseCsv(r *csv.Reader) {
	vCardFile := "output.vcf"
	dat, err := os.Create(vCardFile)
	if err != nil {
		log.Fatal(err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(dat, "BEGIN:VCARD\n")
		io.WriteString(dat, "VERSION:3.0\n")
		io.WriteString(dat, "N:"+record[2]+";"+record[1]+";;;\n")
		io.WriteString(dat, "FN:"+record[1]+" "+record[2]+"\n")
		io.WriteString(dat, "ORG:MTC SFO;\n")
		io.WriteString(dat, "NICKNAME:"+record[3]+"\n")
		io.WriteString(dat, "BDAY;X-APPLE-OMIT-YEAR=2021;VALUE=date:"+record[4]+"\n")
		io.WriteString(dat, "NOTE:"+record[5]+"\n")
		io.WriteString(dat, "item1.ADR;TYPE=HOME;TYPE=pref:;;"+record[6]+"\\n"+record[7]+";"+record[8]+";"+record[9]+";United States\n")
		io.WriteString(dat, "item1.X-ABADR:us\n")
		io.WriteString(dat, "TEL;TYPE=CELL;TYPE=pref;TYPE=VOICE:"+record[10]+"\n")
		io.WriteString(dat, "TEL;TYPE=IPHONE;TYPE=CELL;TYPE=VOICE:"+record[11]+"\n")
		io.WriteString(dat, "TEL;TYPE=HOME;TYPE=VOICE:"+record[12]+"\n")
		io.WriteString(dat, "EMAIL;TYPE=HOME;TYPE=pref;TYPE=INTERNET:"+record[13]+"\n")
		io.WriteString(dat, "EMAIL;TYPE=WORK;TYPE=INTERNET:"+record[14]+"\n")
		io.WriteString(dat, "PRODID:-//Mar Thoma Church Of San Francisco.//Registry dated 20210101//EN\n")
		io.WriteString(dat, "END:VCARD\n")

		/*
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
		*/
	}
	//dat.Sync()
}
