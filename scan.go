package goportmap

import (
	"time"
	"fmt"
	"net"
	"strings"
)

func scan(ip string, port int, timeout time.Duration){
	socket := fmt.Sprintf("%s:%d", ip,port)
	conn, err := net.DialTimeout("tcp", socket, timeout)
	
	if err != nil {
        if strings.Contains(err.Error(), "too many open files") {
            time.Sleep(timeout)
            scan(ip, port, timeout)
        }
		//fmt.Println(port, "closed")
        return
    }
	conn.Close()
	fmt.Println("[+] port",port, ": open")
}