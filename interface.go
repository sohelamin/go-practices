package main

import "fmt"

const Pi = 3.14

type Shape interface {
    Area() float64
}

type Rectangle struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func (rect Rectangle) Area() float64 {
    return rect.width * rect.height
}

func (circ Circle) Area() float64 {
    return Pi * circ.radius * circ.radius
}

func Calculate(sh Shape) {
    fmt.Println("Area: ", sh.Area())
}

func main() {
    rectangle := Rectangle{width: 20, height: 12}
    circle := Circle{radius: 5}

    Calculate(rectangle)
    Calculate(circle)
}
