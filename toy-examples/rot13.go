package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func rot13Cipher(b byte) byte {
    var a, z byte
     switch {
     case 'a' <= b && b <= 'z':
         a, z = 'a', 'z'
     case 'A' <= b && b <= 'Z':
         a, z = 'A', 'Z'
     default:
         return b
     }
     return (b-a+13)%(z-a+1) + a
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)
    for i := 0; i < n ; i++ {
        p[i] = rot13Cipher(p[i])
    }
    return
}
    
func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

