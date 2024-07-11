package main

import "fmt"

type rect struct {
    width, height int
}

// Method with pointer receiver (*rect)
func (r *rect) doubleSize() {
    r.width *= 2
    r.height *= 2
}

// Method with value receiver (rect)
func (r rect) halfSize() {
    r.width /= 2
    r.height /= 2
}

func main() {
    r := rect{width: 10, height: 5}
    
    fmt.Println("Original struct:", r) // Original struct: {10 5}
    
    // Using pointer receiver method
    r.doubleSize()
    fmt.Println("After doubleSize() with pointer receiver:", r) // After doubleSize() with pointer receiver: {20 10}
    
    // Using value receiver method
    r.halfSize()
    fmt.Println("After halfSize() with value receiver:", r) // After halfSize() with value receiver: {20 10} (no change)
}
