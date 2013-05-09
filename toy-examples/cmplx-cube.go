package main

import (
    "fmt"
    "math/cmplx"
    )

func Cbrt(x complex128) complex128 {
    z := complex128(2)
    n := complex128(0)
    for {
        z = z - (cmplx.Pow(z,3) - x)/ (3 * z * z)
        delta := 1e-18
        if cmplx.Abs(n-z) < delta {
            break
        }
        n = z
    }
    return z
}

func main() {
    fmt.Println(Cbrt(2))
}
