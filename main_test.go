package main

import (
	jsonutil "FinsEmu/JsonUtil"
	"testing"
)

func TestRead(t *testing.T) {
	a := jsonutil.New("./test.json")

	a.LoadJson()
	a.AddItemInt("daasdf", 1000).AddItemInt("ddddd", 10003)
	a.WriteJson()

}

func TestUdp(t *testing.T) {
	// soc := udp.New("192.168.56.102", 4003)
	// soc.Listen()
	// buf := make([]byte, 64)
	// n, addr, _ := soc.ReadFrom(buf)

	// fmt.Println(buf[:n])

	// soc.WriteTo([]byte("from go"), addr)

}
