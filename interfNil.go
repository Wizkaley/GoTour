package main

import "fmt"

//type i interface{}


func main(){
	var obj interface{}
	obj = 42
	describe(obj)
	
	obj="Hello"
	describe(obj)
}


func describe(i interface{}){
	fmt.Printf("(%v,%T)\n",i,i)
}
