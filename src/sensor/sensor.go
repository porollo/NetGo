package main

import (
	"fmt"
	"net"
	"os"
)

const (
	network    = "udp"
	address    = ":9995"
	readBuffer = 1048576
)

func main() {

	addr, _ := net.ResolveUDPAddr(network, address)
	sock, _ := net.ListenUDP(network, addr)
	sock.SetReadBuffer(readBuffer)

	i := 0
	for {
		i++
		buf := make([]byte, 1500)
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

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
