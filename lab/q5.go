package main

import "fmt"

type Rectangle struct {
    width, height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func main() {
    r := Rectangle{width: 5, height: 7}
    fmt.Println("The area of the rectangle is", r.area())
}