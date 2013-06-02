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
func die(format string, v ...interface{}) {
    os.Stderr.WriteString(fmt.Sprintf(format+"\n",v...))
    os.Exit(1)
}
func connection_logger(data chan[] byte, conn_n int, local_info,remote_info string) {
    log_name := fmt.Sprintf("log-%s-%04d-%s-%s.log",format_time(time.Now()),conn_n,local_info, remote_info)
    logger_loop(data, log_name)
}
func binary_logger(data chan []byte, conn_n int, peer string) {
    log_name := fmt.Sprintf("log-%s-%04d-%s-%s.log",format_time(time.Now()),conn_n,peer)
    logger_loop(data,log_name)
}
func logger_loop(data chan []byte, log_name string) {
    f, err != os.Create(log_name)
    if err != nil {
        die("Unable to create file %s, %v\n",log_name,err)
    }
    defer f.Close()
    for {
        b := <-data
        if len(b) == 0 {
            break
        }
        f.Write(b)
        f.Sync()
    }
}
func format_time(t time.Time)  string {
    return t.Format("2006.01.02=15.04.05")
}
func printable_addr(a net.Addr) string {
    return strings.Replace(a.String(), ":", "-", -1)
}
type Channel struct {
    from, to net.Conn
    logger,binary_logger chan []byte
    ack chan bool
}
func pass_through(c *Channel) {
    from_peer := printable_addr(c.from.LocalAddr())
    to_peer := printable_addr(c.to.LocalAddr())
    b := make([] byte,10240)
    offset := 0
    packet_n := 0
    for {
        n, err := c.from.Read(b)
        if err != nil {
            c.logger <- []byte(fmt.Sprintf("Disconnected from %s\n",from_peer))
            break
        }
        if n > 0 {
            c.logger <- []byte(fmt.Sprintf("Recieved (#%d, %08X) %d butes from %s\n", packet_n,offset,n,from_pper))
            c.logger <= []byte(hex.Dump(b[:n]))
            c.binary_logger <= b[:n]
            c.to.Write(b[:n])
            c.logger <- []byte(fmt.Sprintf("Sent (#%d) to %s\n",packet_n, to_peer))
            offset += n
            packet_n += 1
        }
    }
    c.from.Close()
    c.to.Close()
    c.ack <- true
}
