package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

// Variables to delare
var (
	pcapFile string = "tcpdump.pcap00" // pcap filename
	//deviceName  string = "en0"  // network interface (this part is commented because we are not using OpenLive)
	snapshotLen uint32 = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = -1 * time.Second
	handle      *pcap.Handle
)

func main() {

	// Open tcpdump.pcap file and write header
	f, _ := os.Create("http_tcpdump.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
	defer f.Close()

	// Open pcapFile
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// This part is commented out because it is not used. If we want to capture packets using the
	// network interface, we can use this block of code. Remember to uncomment deviceName in the list of
	// variables at the top.
	/*
		// Open device (eth0) for capturing
		handle, err = pcap.OpenLive(deviceName, int32(snapshotLen), promiscuous, timeout)
		if err != nil {
			fmt.Printf("Error opening device %s: %v", deviceName, err)
			os.Exit(1)
		}
		defer handle.Close()
	*/

	// Set filter to only look for TCP Port 80
	var filter string = "tcp and port 80"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 80 packets.")

	// Loop through packets and write to file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	}
}
