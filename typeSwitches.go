package main

import "fmt"

func do( i interface{}){
	switch v := i.(type){
	case int:
		fmt.Printf("TWICE %v IS %v\n",v,v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n",v,len(v))
	default:
		fmt.Printf("I dont know about type %T!\n",v)
	}
	
}

func main(){
	do(21)
	do("Eshan")
	do(true)
}

