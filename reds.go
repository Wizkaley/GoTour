package main

import (
	"io"
	"fmt"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (read rot13Reader) Read(s [] byte)(int,error){
	var ascii []int  
	for i:= range s{
		ascii[i] = int(s[i])
		fmt.Println(ascii)
	}
	return 0,nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
