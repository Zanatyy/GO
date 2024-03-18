package main

import (
	"fmt"
	"reflect"
	)

func getType(myvar interface{}) string {
	return reflect.TypeOf(myvar).Name()
}


type Circle struct {
	r float64
}
type Square struct {
	l float64
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.r * c.r
}
func (s Square) Area() float64 {
	return s.l * s.l
}
func (c Circle) Perimeter() float64 {
	return 3.14 * 2 * c.r
}
func (s Square) Perimeter() float64 {
	return 4 * s.l
}
func measure(shape Shape) {
	fmt.Printf("The shape is: %s \n Area: %f , Perimeter: %f\n ", getType(shape), shape.Area(), shape.Perimeter())
}
func main() {
	circle1, square1 := Circle{5}, Square{6}
	measure(&circle1)
	measure(&square1)

}


