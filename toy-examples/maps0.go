package main

import (
    "fmt"
)

type Vertex struct {
    lat, long float32
}

func main() {
    m := make(map[string]Vertex)
    m["Test Location"] = Vertex {
        122.3232, -123.849329,
    }
    fmt.Println(m)
    fmt.Println(m["Test Location"])
}
