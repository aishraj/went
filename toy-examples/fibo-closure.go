package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a,b := 0, 1
    c := 0
    return func() int {
        c = a+b 
        a, b = b,c
        return a 
    }
    
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
