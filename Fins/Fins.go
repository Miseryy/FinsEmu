package fins

import (
	jsonutil "FinsEmu/JsonUtil"
	udp "FinsEmu/UDP"
	"errors"
	"net"
	"strconv"
	"strings"
)

type RecvParam struct {
	DNA        byte
	DA1        byte
	SNA        byte
	SA1        byte
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

func rangeCheck(v, min, max int) (bool, error) {
	if min > max {
		return false, errors.New("MIN > MAX")
	}

	if v < min {
		return false, nil
	}

	if v > max {
		return false, nil
	}

	return true, nil

}

func MakeSendCommand(sock *udp.Udp_Sock, recv_param RecvParam, ip string, json_path string) ([]byte, error) {
	js := jsonutil.New(json_path)
	js.LoadJson()

	json_map := js.GetMap()
	start_pos := (int(recv_param.START_READ[0]) << 8) + int(recv_param.START_READ[1])
	read_size := (int(recv_param.READ_SIZE[0]) << 8) + int(recv_param.READ_SIZE[1])

	const fins_command_len = 14
	buff_len := fins_command_len + (read_size * 2)

	send_buff := make([]byte, buff_len)

	send_buff[0] = 0xC0           // ICF
	send_buff[1] = 0x00           // RSV
	send_buff[2] = 0x02           // GCT
	send_buff[3] = recv_param.SNA // DNA
	send_buff[4] = recv_param.SA1 // DA1
	send_buff[5] = 0x00           // DA2
	send_buff[6] = recv_param.DNA // SNA
	send_buff[7] = recv_param.DA1 // SA1
	send_buff[8] = 0x00           // SA2
	send_buff[9] = recv_param.SID // SID
	send_buff[10] = recv_param.COMMCODE1
	send_buff[11] = recv_param.COMMCODE2
	send_buff[12] = 0x00 // end code
	send_buff[13] = 0x00 // end code

	// write command
	if send_buff[10] == 0x01 && send_buff[11] == 0x02 {
		return send_buff[:13], nil
	}

	// max 999
	const max = 999

	if read_size > max {
		return nil, errors.New("Read Size Over")
	}

	end_pos := start_pos + read_size

	for i := 0; i < end_pos-start_pos; i++ {
		read_num := strconv.Itoa(start_pos + i)
		high := fins_command_len + i*2
		low := fins_command_len + i*2 + 1

		v, e := json_map[read_num]

		if !e {
			// not regist data
			send_buff[high] = byte(0)
			send_buff[low] = byte(0)
			continue
		}

		value := int(v.(float64))

		high_hex := value >> 8
		low_hex := value & 0x00ff
		send_buff[high] = byte(high_hex)
		send_buff[low] = byte(low_hex)
	}

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
	R_DNA := buff[3]
	R_DA1 := buff[4]
	// R_DA2 := buff[5]
	R_SNA := buff[6]
	R_SA1 := buff[7]
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

	recv_param.DNA = R_DNA
	recv_param.DA1 = R_DA1
	recv_param.SNA = R_SNA
	recv_param.SA1 = R_SA1
	recv_param.SID = R_SID
	recv_param.COMMCODE1 = R_COMMCODE1
	recv_param.COMMCODE2 = R_COMMCODE2
	recv_param.IO_MEM = IO_MEM
	recv_param.START_READ = START_READ
	recv_param.READ_SIZE = READ_SIZE

	return recv_param, nil
}
