package fins

import (
	jsonutil "FinsEmu/JsonUtil"
	udp "FinsEmu/UDP"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Fins struct {
}

type RecvParam struct {
	SID        byte
	COMMCODE1  byte
	COMMCODE2  byte
	IO_MEM     byte
	START_READ []byte
	READ_SIZE  []byte
}

func RecvHostData(sock *udp.Udp_Sock) ([]byte, *net.UDPAddr, error) {
	buff := make([]byte, 1024)
	num, addr, err := sock.ReadFrom(buff)

	return buff[:num], addr, err
}

func getIP4digit(ip string) int {
	split_ip := strings.Split(ip, ".")
	i, _ := strconv.Atoi(split_ip[3])

	return i

}

func MakeSendCommand(sock *udp.Udp_Sock, recv_param RecvParam, ip string, json_path string) ([]byte, error) {
	js := jsonutil.New(json_path)
	js.LoadJson()

	json_map := js.GetMap()
	read_size := (int(recv_param.READ_SIZE[0]) << 8) + int(recv_param.READ_SIZE[1])

	buff_len := 13 + read_size
	fmt.Println(buff_len)

	send_buff := make([]byte, buff_len)
	da1 := getIP4digit(ip)

	addr, _ := sock.GetAddressAndPort()

	digit4 := getIP4digit(addr)

	send_buff[0] = 0xC1           // ICF
	send_buff[1] = 0x00           // RSV
	send_buff[2] = 0x02           // GCT
	send_buff[3] = 0x00           // DNA
	send_buff[4] = byte(da1)      // DA1
	send_buff[5] = 0x00           // DA2
	send_buff[6] = 0x00           // SNA
	send_buff[7] = byte(digit4)   // SA1
	send_buff[8] = 0x00           // SA2
	send_buff[9] = recv_param.SID // SID
	send_buff[10] = recv_param.COMMCODE1
	send_buff[11] = recv_param.COMMCODE2
	send_buff[12] = 0x00 // end code
	send_buff[13] = 0x00 // end code

	// max 999
	const max = 999

	v := json_map
	_ = v
	for _, v := range json_map {
		if read_size > max {
			return nil, errors.New("Read Size Over")
		}
		_ = v
		// if recv_param.START_READ > k {
		// }
	}

	// recv_param.IO_MEM

	return send_buff, nil
}

func CheckFinsCommand(buff []byte) (RecvParam, error) {
	// C0 00 02 01 c8 00 01 01 fa 00 01 01 00 00 00 6d
	//  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
	recv_param := RecvParam{}

	if len(buff) < 13 {
		return recv_param, errors.New("Recv Data Error")
	}

	// Fins Header
	// R_ICF := buff[0]
	// R_RSV := buff[1]
	// R_GCT := buff[2]
	// R_DNA := buff[3]
	// R_DA1 := buff[4]
	// R_DA2 := buff[5]
	// R_SNA := buff[6]
	// R_SA1 := buff[7]
	// R_SA2 := buff[8]
	R_SID := buff[9]
	R_COMMCODE1 := buff[10]
	R_COMMCODE2 := buff[11]
	IO_MEM := buff[12]
	START_READ := []byte{buff[13], buff[14], buff[15]}
	READ_SIZE := []byte{buff[16], buff[17]}

	if IO_MEM != 0x82 {
		return recv_param, errors.New("Only DM Area")
	}

	recv_param.SID = R_SID
	recv_param.COMMCODE1 = R_COMMCODE1
	recv_param.COMMCODE2 = R_COMMCODE2
	recv_param.IO_MEM = IO_MEM
	recv_param.START_READ = START_READ
	recv_param.READ_SIZE = READ_SIZE

	return recv_param, nil
}

func SendPLCData(sock udp.Udp_Sock) error {
	return nil
}
