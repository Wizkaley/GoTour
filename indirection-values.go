package main
import (
	"fmt"
	"math"
)


type Vertex struct{
	X,Y float64
}

func (v Vertex) Abs()float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex)float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main(){
	g := Vertex{4,6}
	fmt.Println(g.Abs())
	fmt.Println(AbsFunc(g))


	f := &Vertex{3,5}
	fmt.Println(f.Abs())
	fmt.Println(AbsFunc(*f))

}