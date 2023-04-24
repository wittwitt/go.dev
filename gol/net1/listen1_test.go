package net1

import (
	"net"
	"testing"
)

func TestNet_Listen_random_port(t *testing.T) {
	listen := ":0"

	nl, err := net.Listen("tcp", listen)
	if err != nil {
		t.Log(err)
		return
	}

	switch addr := nl.Addr().(type) {
	case *net.UDPAddr:
		// p.SrcIP = addr.IP.String()
		// p.SrcPort = uint(addr.Port)
		// p.DstPort = uint(localAddr.(*net.UDPAddr).Port)

		t.Log(addr.Port)
	case *net.TCPAddr:
		// p.SrcIP = addr.IP.String()
		// p.SrcPort = uint(addr.Port)
		// p.DstPort = uint(localAddr.(*net.TCPAddr).Port)

		t.Log(addr.Port)
	}

	t.Log(nl.Addr().String())

}

func TestNet1_get_interface_IP(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		t.Log(err)
		return
	}
	IPs := make([]string, 0)
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				IPs = append(IPs, ipNet.IP.To4().String())
			}
		}
	}
	t.Log(IPs)
}

func TestNet1_external_lan_ip(t *testing.T) {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		t.Log(err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.TCPAddr)
	t.Log(localAddr.IP.String())
}
