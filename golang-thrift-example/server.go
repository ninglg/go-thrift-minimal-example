package main

import (
	"ThriftDemo/go-example"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"strings"
)

type FormatDataImpl struct{}

func (fdi *FormatDataImpl) DoFormat(data *example.Data) (r *example.Data, err error) {
	var rData example.Data
	rData.Text = strings.ToUpper(data.Text)

	return &rData, nil
}

const (
	ServerHost = "localhost"
	ServerPort = "8099"
)

func main() {
	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(ServerHost + ":" + ServerPort)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", ServerHost+":"+ServerPort)
	server.Serve()
}
