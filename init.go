package goportmap

import (
	"strings"
	"strconv"
	"log"
)
func (p *GoPortMap)Init(ip, portRange string, thread int){
	p.ip = ip
	p.thread = thread
	p.portRange = portRange
	
	ips := strings.Split(portRange, "-")
	startRange, sErr := strconv.Atoi(ips[0])
	endRange, eErr := strconv.Atoi(ips[1])

	if (sErr != nil) || (eErr != nil){
		log.Fatal("[-] Invalid Range")
	}else if(startRange > endRange){
		log.Fatal("[-] Invalid Range")
	}

	p.startRange = startRange
	p.endRange = endRange


}