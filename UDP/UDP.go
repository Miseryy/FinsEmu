package udp

import "net"

type Udp_Sock struct {
	udpAddr *net.UDPAddr
	udpLn   *net.UDPConn
}

func New(ip_addr string, port int) *Udp_Sock {
	t := &Udp_Sock{
		udpAddr: &net.UDPAddr{
			IP:   net.ParseIP(ip_addr),
			Port: port,
		},
	}

	return t
}

func (ud *Udp_Sock) Listen() error {
	u, err := net.ListenUDP("udp", ud.udpAddr)
	ud.udpLn = u

	return err
}

func (ud *Udp_Sock) ReadFrom(buf []byte) (int, *net.UDPAddr, error) {
	n, addr, err := ud.udpLn.ReadFromUDP(buf)
	return n, addr, err

}

func (ud *Udp_Sock) WriteTo(buf []byte, addr *net.UDPAddr) (int, error) {
	n, err := ud.udpLn.WriteTo(buf, addr)
	return n, err

}