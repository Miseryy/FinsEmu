package main

import (
	jsonutil "FinsEmu/JsonUtil"
	udp "FinsEmu/UDP"
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	a := jsonutil.New()

	a.LoadJson("./test.json")
	a.AddItem("daasdf", 1000).AddItem("ddddd", 10003)
	a.WriteJson("./test.json")

}

func TestUdp(t *testing.T) {
	soc := udp.New("192.168.56.102", 4003)
	soc.Listen()
	buf := make([]byte, 64)
	n, addr, _ := soc.ReadFrom(buf)

	fmt.Println(buf[:n])

	soc.WriteTo([]byte("from go"), addr)

}
