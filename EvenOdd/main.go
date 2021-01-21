package main

import "fmt"

func main() {
	letters := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range letters {
		if (n % 2) == 0 {
			fmt.Println(n, " is Even")
		} else {
			fmt.Println(n, " is Odd")
		}
	}
	//fmt.Println(letters)
}
