package main
import (
	"fmt"
	"math"
)


type Abser interface{
	Abs() float64
}

type MyFloat float64

type Vertex struct{
	X,Y float64
}
func (v * Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat ) Abs() float64{
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}


func main(){
	var obj Abser
	v := Vertex{3,5}
	f := MyFloat(-math.Sqrt2)

	obj = f //MyFloat implements Abser
	obj = &v //*Vertex implements Abser

	obj = &v

	fmt.Println(obj.Abs())
}