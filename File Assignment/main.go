package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f := os.Args[1]
	d1, err := ioutil.ReadFile(f)
	check(err)
	fmt.Print(string(d1))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
