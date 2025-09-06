package main

import (
	"fmt"
	"log"

	"github.com/affanhamid/goDFS/p2p"
)

func Onpeer(p2p.Peer) error {
	println("doing some logic with the peer outside of TCPTransport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.NOPDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
		OnPeer:        Onpeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
