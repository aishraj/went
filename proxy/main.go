package main
import {
    "encoding/hex"
    "flag"
    "fmt"
    "net"
    "os"
    "runtime"
    "strings"
    "time"
}
var ( 
    host *string = flag.String("host","", "target host or address")
    port *string = flag.String("port","0","target port")
    listen_port *string = flag.String("listen_port","0", "listen_port")
)
func

