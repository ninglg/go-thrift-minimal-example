package main

import (
	"ThriftDemo/gen-go/example"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"net"
)

const (
	ClientHost = "localhost"
	ClientPort = "8099"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(ClientHost, ClientPort))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", ClientHost+":"+ClientPort)
	}
	defer transport.Close()

	data := example.Data{Text: "hello,world!"}
	d, err := client.DoFormat(&data)
	fmt.Println(d.Text)
}
