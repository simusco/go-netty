package main

import (
	"fmt"
	"net"
	"simusco.com/go-netty"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	go_netty.Accept(listener)
}
