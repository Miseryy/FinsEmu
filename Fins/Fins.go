package fins

import (
	udp "FinsEmu/UDP"
	"fmt"
	"net"
)

type Fins struct {
}

func RecvHostData(sock *udp.Udp_Sock) ([]byte, *net.UDPAddr, error) {
	buff := make([]byte, 128)
	num, addr, err := sock.ReadFrom(buff)

	return buff[:num], addr, err
}

func CheckFinsCommand(buff []byte) ([]byte, error) {
	// C0 00 02 01 c8 00 01 01 fa 00 01 01 00 00 00 6d
	// Fins Header
	R_ICF := buff[0]
	// R_RSV := buff[1]
	// R_GCT := buff[2]
	// R_DNA := buff[3]
	// R_DA1 := buff[4]
	// R_DA2 := buff[5]
	// R_SNA := buff[6]
	// R_SA1 := buff[7]
	// R_SA2 := buff[8]
	// R_SID := buff[9]
	send_buff := make([]byte, 1024)

	send_buff[0] = 0b11000000 | R_ICF

	fmt.Printf("%X", send_buff[0])
	// fmt.Println(S_ICF)

	return send_buff, nil
}

func SendPLCData(sock udp.Udp_Sock) error {
	return nil
}
