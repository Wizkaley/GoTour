package main

import "fmt"

type Vertex struct{
	X,Y float64
}

func (v * Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v * Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	g := Vertex{2,5}
	g.Scale(2)

	ScaleFunc(&g,10)

	h := &Vertex{4,6}
	h.Scale(3)
	ScaleFunc(h,5)

	fmt.Println(g,h)
}

