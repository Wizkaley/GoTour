package main

import "image"
import "image/color"


type Image struct{
	w,h int
}

func main() {
	m := Image{100,100}
	pic.ShowImage(m)
}


func (i *Image)Bounds() image.Rectangle{
	return image.Rect(0,0,i.w,i.h)
}


func (i * Image) ColorModel() color.Model{
	return color.RGBAModel
} 

func (i Image)At(x,y int)color.Color{
	v:= (x^y)
	return color.RGBA(v,v,0,255)
}