package main

import(
 "fmt"
 "math"
)

type I interface{
	M()
}

type T struct{
	s string
}

type MyFloat float64

func (t *T) M(){
	fmt.Print(t.s)
}

func (f MyFloat) M(){
	fmt.Println(f)
}

func main(){
	var obj I
	obj = &T{"Eshan"}
	describe(obj)
	obj.M()

	obj = MyFloat(math.Pi)
	describe(obj)
	obj.M()
 }


 func describe(i I){
	 fmt.Printf("(%v,%T)\n",i,i)
 }

