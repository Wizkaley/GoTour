package main
import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}


func rot13(x byte)byte{
	capital := x >='A' && x <= 'Z'

	if !capital && (x < 'a' || x > 'z'){
		fmt.Println("Not a letter")
	}
	x+=13
	if capital && x > 'Z'  || !capital &&  x > 'z'{
		x-=26
	}
	return x
}
func (read rot13Reader) Read(s [] byte)(int,error){
		n, err := read.r.Read(s)

		for i:=0;i<n;i++{
			s[i] = rot13(s[i])
		}
		return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}