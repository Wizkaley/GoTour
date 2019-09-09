package main

import "fmt"

func main1() {
	names := [4]string{
		"Eshan",
		"Mehfi",
		"Gaurav",
		"Rajat",
	}

	fmt.Println(names)

	a := names[0:3]
	b := names[2:4]
	fmt.Println(a)
	fmt.Println(b)

	b[0] = "XXXX"
	fmt.Println(a, b)
}
