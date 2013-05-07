package main

import (
    "fmt"
    "math"
)

var delta = 0.00

func Sqrt(x float64) float64 {
    z := 1.0
    var newZ = z
    for {
        z = findSqrt(z,x)
        newZ = findSqrt(z,x)
        if z-newZ <= delta {
            break
        }
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
