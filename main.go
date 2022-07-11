package main

import (
	udp "FinsEmu/UDP"
	"fmt"
)

func main() {
	soc := udp.New("192.168.56.102", 4003)
	soc.Listen()
	buf := make([]byte, 64)
	n, addr, _ := soc.ReadFrom(buf)

	fmt.Println(buf[:n])

	soc.WriteTo([]byte("from go"), addr)

	// ui.Test()

}
