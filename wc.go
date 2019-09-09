package main

import "strings"
import "fmt"

func wordCount(s string)map[string]int {

	m := map[string]int{}
	m[s] = cap(strings.Split(s,""))
	return m
	//fmt.Println(cap(strings.Split(s,"")))
	//maps
}

func main() {
	fmt.Println(wordCount("Eshan"))
}
