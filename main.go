package main

import (
	jsonutil "FinsEmu/JsonUtil"
	udp "FinsEmu/UDP"
	"fmt"
)

func udp_test() {
	soc := udp.New("192.168.56.102", 4003)
	soc.Listen()
	buf := make([]byte, 64)
	n, addr, _ := soc.ReadFrom(buf)

	fmt.Println(buf[:n])

	soc.WriteTo([]byte("from go"), addr)

}

func main() {
	a := jsonutil.New()
	a.LoadJson("./test.json")

	// ui.Test()

}
