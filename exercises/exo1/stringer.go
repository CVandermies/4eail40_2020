package main

import (
	"fmt"
)

// Implement types Rectangle, Circle and Shape

type Rectangle struct {
	width, length int
}

type Circle struct {
	radius int
}

type Shape interface {
	String() string
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %v and length %v", r.width, r.length)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %v", c.radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
