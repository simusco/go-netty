package main

import (
	"fmt"
	"simusco.com/go-netty/actions"
	"simusco.com/go-netty/channel"
	"simusco.com/go-netty/constant"
	"simusco.com/go-netty/protobuf"
)

type TestAction struct {
}

func (action *TestAction) Execute(channel *channel.Channel, params *LiveProto.ParamVO) *LiveProto.ResultVO {
	if params != nil {
		fmt.Printf("%s,%d", params.Pkg, params.Event)
	}

	return &LiveProto.ResultVO{
		Code: constant.SUCCESS,
		Msg:  "成功",
	}
}

func init() {
	actions.Register(constant.TestEventCode, &TestAction{})
}
