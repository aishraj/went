package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    i := 1
    for i <= 10 {
        z = findSqrt(z,x)
        i++
    }
    return z
}

func findSqrt(z,x float64) float64 {
    var seVal = ((z*z)-x)/(2*z)
    return z-seVal
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
