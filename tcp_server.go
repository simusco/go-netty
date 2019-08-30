package go_netty

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"simusco.com/go-netty/channel"
	"simusco.com/go-netty/handler"
	"simusco.com/go-netty/protobuf"
)

func Accept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			_ = conn.Close()
			continue
		}

		cc := channel.NewChannel(conn)

		go func(ch *channel.Channel) {
			for {
				var bytes []byte
				_, err = ch.Read(bytes)
				if err != nil {
					ch.Close()
					fmt.Println("客户端断开了")
					break
				}

				tcpInboundHandler := handler.TcpInboundHandler{}

				params := new(LiveProto.ParamVO)
				err = proto.Unmarshal(bytes, params)
				if err != nil {
					tcpInboundHandler.ErrorCaught(ch, err)
					continue
				}

				tcpInboundHandler.Handle(ch, params)
			}
		}(cc)
	}

}
