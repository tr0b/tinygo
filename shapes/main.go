package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	length float64
}

type triangle struct {
	baseLength float64
	height     float64
}

func main() {
	shs := []shape{square{length: 9.5}, triangle{baseLength: 8.5, height: 17.5}}
	for _, s := range shs {
		printArea(s)
	}

}

func (s square) getArea() float64 {
	return s.length * s.length
}

func (t triangle) getArea() float64 {
	return (t.baseLength * t.height) / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
