package main

import (
	"flag"
	"fmt"
	"net"
)


var srcIP = flag.String("si", "0.0.0.0", "print received and sent packets")
var srcPort = flag.Int("sp", 8126, "print received and sent packets")
var dstIP = flag.String("di", "192.168.255.253", "print received and sent packets")
var dstPort = flag.Int("dp", 8126, "print received and sent packets")
var print = flag.Bool("p", true, "print received and sent packets")

func main() {
	listenAddr := net.UDPAddr{
		Port: *srcPort,
		IP:   net.ParseIP(*srcIP),
	}
	listenConn, err := net.ListenUDP("udp", &listenAddr)
	if err != nil {
		panic(err)
	}
	defer listenConn.Close()

	sendAddr := net.UDPAddr{
		Port: *dstPort,
		IP:   net.ParseIP(*dstIP),
	}
	sendConn, err := net.ListenUDP("udp", &sendAddr) 
	if err != nil {
		panic(err)
	}
	defer sendConn.Close()
	
	var buf [1024]byte
	for {
		rlen, _, err := sendConn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}
		if *print {
			fmt.Printf("Received bytes: %v", buf)
		}

		sendBuf := make([]byte, rlen)
		for idx := 0; idx < rlen; idx += 1 {
			sendBuf[idx] = buf[idx]
		}
		sendConn.Write(sendBuf)
	}	
}
