package goportmap

import (
	"time"
	"sync"
	"fmt"
)

func (p *GoPortMap)Map(wg *sync.WaitGroup, cond *sync.Cond){
	counter := 0
	timeout := 5 * time.Second

	fmt.Println("[+] Attacking :", p.ip, )
	fmt.Println("[+] Port Range :", p.portRange, )

	for port := p.startRange; port <= p.endRange; port ++{
		cond.L.Lock()

		for counter == p.thread{ //control threads
			cond.Wait()
		}
		wg.Add(1)

		go func(port int){
			defer wg.Done()
			scan(p.ip,port, timeout)
			cond.L.Lock()
			counter--  //decrement
			cond.L.Unlock()
			cond.Signal()
		}(port)
		counter++
		cond.L.Unlock()
	}
}