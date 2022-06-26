package main

import (
	"github.com/cyberkhalid/gosec/goportmap/goportmap"
	"flag"
	"fmt"
	"sync"
)

func main()  {
	//cmd flags
	//cmd flags
	host:= flag.String("target","","Hostname or ip of the target.")
	portRange := flag.String("r", "0-1000", "port range to be scanned.")
	thread := flag.Int("t",5,"number of thread to use")
	flag.Parse()

	var goPortMap goportmap.GoPortMap
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})

	goPortMap.Init(*host, *portRange, *thread)
	goPortMap.Map(&wg, cond)
	wg.Wait()
	fmt.Println("[+] Done")
}