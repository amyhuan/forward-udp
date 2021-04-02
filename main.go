package main

import (
	"flag"
	"fmt"
	"net"
	// "log"
	// "time"

	// "github.com/google/gopacket"
	// "github.com/google/gopacket/layers"
)


var srcIP = flag.String("si", "127.0.0.1", "listen on this ip")
var srcPort = flag.Int("sp", 888, "listen on this port")
var newSrcIP = flag.String("nsi", "127.0.0.1", "send from this ip")
var newSrcPort = flag.Int("nsp", 777, "send from this port")
var dstIP = flag.String("di", "127.0.0.1", "send to this ip")
var dstPort = flag.Int("dp", 999, "send to this port")
var print = flag.Bool("p", true, "print received and sent packets")

func main() {
	flag.Parse()

	listenAddr,err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", *srcIP, *srcPort))
	if err != nil {
		fmt.Println("Error: ",err)
	} 
 
    /* Now listen at selected port */
    listenCon, err := net.ListenUDP("udp", listenAddr)
    defer listenCon.Close()

	sendFrom, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", *newSrcIP, *newSrcPort))
	if err != nil {
		fmt.Println("Error: ", err)
	} 
	sendTo, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", *dstIP, *dstPort))
	if err != nil {
		fmt.Println("Error: ", err)
	} 
    
    sendConn, err := net.DialUDP("udp", sendFrom, sendTo)
	if err != nil {
		fmt.Println("Error: ", err)
	} 
    defer sendConn.Close()
 
    buf := make([]byte, 1024)
    for {
        n, addr, err := listenCon.ReadFromUDP(buf)
		if err != nil {
            fmt.Println("Error: ", err)
        } 

		_, err = sendConn.Write(buf)
        if err != nil {
            fmt.Println("Error: ", err)
        } 

		if *print {
			fmt.Println("Received and sent ", string(buf[0:n]), " from ", addr)
		}
    }

	// protocol := "udp"
    // netaddr, _ := net.ResolveIPAddr("ip4", *srcIP)
    // listenCon, err := net.ListenIP("ip4:"+protocol, netaddr)
	// if err != nil {
	// 	log.Fatalf("ListenIP: %s\n", err)
	// }
	// if err = listenCon.SetDeadline(time.Now().Add(time.Millisecond)); err != nil {
    //     log.Fatalln("Can't set appropriate deadline!")
    // }
	// defer listenCon.Close()

	// sendCon, err := net.Dial("ip4:udp", *dstIP)
	// if err != nil {
	// 	log.Fatalf("Dial: %s\n", err)
	// }
	// defer sendCon.Close()
    
	// buf := make([]byte, 1024)

	// for {
	// 	numRead, _, err := listenCon.ReadFrom(buf)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	if numRead > 0 {
	// 		pkt := gopacket.NewPacket(buf[:numRead], layers.LayerTypeUDP, gopacket.Default)
	// 		if udpLayer := pkt.Layer(layers.LayerTypeUDP); udpLayer != nil {
	// 			udpPkt := udpLayer.(*layers.UDP)
	// 			if udpPkt.DstPort == layers.UDPPort(*srcPort) {
	// 				sendCon.Write(buf[:numRead])
	// 				fmt.Printf("Sent: % X\n", buf[:numRead])
	// 			}
	// 		}
	// 	}
	// }	
}
