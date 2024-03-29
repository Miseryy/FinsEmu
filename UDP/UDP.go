package udp

import (
	"errors"
	"net"
	"strconv"
)

type Udp_Sock struct {
	udpAddr *net.UDPAddr
	udpLn   *net.UDPConn
}

func New() *Udp_Sock {
	return &Udp_Sock{}
}

func (self *Udp_Sock) SetAddr(ip_addr string, port int) {
	self.udpAddr = &net.UDPAddr{
		IP:   net.ParseIP(ip_addr),
		Port: port,
	}
}

func (self *Udp_Sock) Listen() error {
	u, err := net.ListenUDP("udp", self.udpAddr)
	self.udpLn = u

	return err
}

func (self *Udp_Sock) ReadFrom(buf []byte) (int, *net.UDPAddr, error) {
	n, addr, err := self.udpLn.ReadFromUDP(buf)
	return n, addr, err

}

func (self *Udp_Sock) WriteTo(buf []byte, addr *net.UDPAddr) (int, error) {
	n, err := self.udpLn.WriteTo(buf, addr)
	return n, err

}

func (self *Udp_Sock) Close() error {
	if self.udpLn == nil {
		return errors.New("Closed")
	}

	e := self.udpLn.Close()
	return e
}

func (self *Udp_Sock) GetAddressAndPort() (string, string) {
	str_port := strconv.Itoa(self.udpAddr.Port)

	return self.udpAddr.IP.String(), str_port
}
