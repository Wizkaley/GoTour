package main

import "fmt"

type Person struct{
	name string
	age int
}


func (p Person)String ()string{
	return fmt.Sprintf("The Age of %v is %v",p.name,p.age)
}


func main(){
	a := Person{"Eshan",400}
	b := Person{"Nitisha",18}
	fmt.Println(a,b)
}
