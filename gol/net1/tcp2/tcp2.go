package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"syscall"

	"golang.org/x/sys/unix"
)

/*

	c的例子，，地址和端口复用，默认都没有开启，所有，server和client通讯后，，server主动kill掉，，看到time-wait了
	，重启，会看到addres already use

	go，默认开启了 地址复用，没有开启端口复用，，所有，go可以在time-wait后，，再次重启立马listen

*/

func main() {

	addr := "127.0.0.1:12000"

	// go net.listen 开启地址复用，即， SO_REUSEADDR  开启，，但SO_REUSEPORT 端口复用没有开启

	// return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
	// func (fd *netFD) listenStream(laddr sockaddr, backlog int) error {
	//    if err := setDefaultListenerSockopts(fd.sysfd)

	nl, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen: %s, %v", addr, err)
		return
	}
	conn, err := nl.Accept()
	if err != nil {
		log.Printf("Accept: %s, %v", addr, err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Read: %s, %v", addr, err)
		return
	}
	fmt.Printf("read: %s", buf[:n])
	conn.Write([]byte("abc"))
	conn.Close()
	nl.Close()
}

// linux 配置 绑定端口的保护
func listen1(addr string) (nl net.Listener, err error) {
	var listenAddr *net.TCPAddr

	if addr == "" {
		addr = ":0"
		// 范围 /proc/sys/net/ipv4/ip_local_port_range
		// 随即端口有一定的概率提示端口已经被占用了
	}

	// 重启有一定的概率  address already in use
	if nl, err = net.Listen("tcp", addr); err != nil {
		return nil, fmt.Errorf("listen: %s, %v", addr, err)
	}

	listenAddr = nl.Addr().(*net.TCPAddr)
	log.Printf("listen: %s\n", listenAddr.String())
	return
}

// linux
// SO_REUSEADDR  SO_REUSEPORT
func listen2(addr string) (nl net.Listener, err error) {
	listenCfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			})
		},
	}
	nl, err = listenCfg.Listen(context.Background(), "tcp", addr)
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}
	return
}
