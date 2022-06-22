package main

import (
	"fmt"
	"math"
)

type IGeometry interface {
	luas() float64
	keliling() float64
}

type PERSEGI struct {
	panjang, lebar float64
}

//implementasi persegi
func (p PERSEGI) luas() float64 {
	return p.lebar * p.panjang
}

func (p PERSEGI) keliling() float64 {
	return 2*p.lebar + 2*p.panjang
}

type LINGKARAN struct {
	radius float64
}

//implementasi lingkaran
func (l LINGKARAN) luas() float64 {
	return math.Pi * l.radius * l.radius
}

func (l LINGKARAN) keliling() float64 {
	return 2 * math.Pi * l.radius
}

func measure(g IGeometry) {
	fmt.Println(g)
	fmt.Println("luas: ", g.luas())
	fmt.Println("keliling: ", g.keliling())
}

func main() {
	p := PERSEGI{panjang: 5, lebar: 10}
	c := LINGKARAN{radius: 10}
	measure(p)
	measure(c)
}
