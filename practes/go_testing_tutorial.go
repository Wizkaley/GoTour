package main

import(
	"fmt"
)

func Calculate(x int)int{
	var res int
	res = x + 2
	return res
}

func main(){
	fmt.Println("Go testing tutorial")
	
	fmt.Println(Calculate(2))
}
