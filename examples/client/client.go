package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"simusco.com/go-netty/channel"
	"simusco.com/go-netty/constant"
	"simusco.com/go-netty/protobuf"
	"simusco.com/go-netty/util"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println("连接服务器错误")
		return
	}

	ch := channel.NewChannel(conn)

	b := make(chan bool, 1)
	go func(c *channel.Channel) {
		for {
			var bb []byte
			_, err := c.Read(bb)
			if err != nil {
				b <- true
				break
			}

			res := new(LiveProto.ResultVO)
			err = proto.Unmarshal(bb, res)
			if err != nil {
				fmt.Print("解析结果错误")
				continue
			}

			fmt.Printf("%d,%s", res.Code, res.Msg)
		}
	}(ch)

	WriteProtobuf(ch)

	<-b
}

func WriteProtobuf(ch *channel.Channel) {

	params := LiveProto.ParamVO{
		Token:     "1",
		Os:        "ios",
		Version:   100,
		Pkg:       "com.test",
		Event:     constant.TestEventCode,
		Channelid: "xiaomi",
		MsgId:     strconv.FormatInt(util.GetRand(100000, 200000), 10),
	}

	value, err := proto.Marshal(&params)
	if err != nil {
		fmt.Println("编码参数失败")
		return
	}

	_, _ = ch.Write(value)

}

func WriteString(ch *channel.Channel) {

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("%v", err)
			break
		}

		value := []byte(input)

		_, _ = ch.Write(value)
	}

}
