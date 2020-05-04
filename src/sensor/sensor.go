package main

import (
	"fmt"
	"net"
)

const (
	network    = "udp"
	address    = ":9995"
	readBuffer = 1500
)

func main() {

	addr, _ := net.ResolveUDPAddr(network, address)
	sock, _ := net.ListenUDP(network, addr)

	i := 0
	for {
		i++
		buf := make([]byte, readBuffer)
		rlen, _, err := sock.ReadFromUDP(buf)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(buf[0:rlen])
		fmt.Println(i)

		go savePackets(buf, rlen)
	}
}

func savePackets(buf []byte, rlen int) {
	print("Hello")
}
