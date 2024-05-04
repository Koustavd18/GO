package main

import "fmt"

type triangle struct{
	height float64
	base float64
}

type square struct {
	sideLength float64
}

type shape interface{
	getArea() float64
}

func main() {
	t := triangle{
		height: 12,
		base: 10,
	}

	s := square{
		sideLength: 10,
	}
	
	PrintArea(t)
	PrintArea(s)
}

func (s square) getArea() float64{
	return s.sideLength*s.sideLength
}

func (t triangle) getArea() float64{
	return 0.5*t.height*t.base
}

func PrintArea(sh shape){
	fmt.Println(sh.getArea())
}
