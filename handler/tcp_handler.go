package handler

import (
	"github.com/golang/protobuf/proto"
	"simusco.com/go-netty/actions"
	"simusco.com/go-netty/channel"
	"simusco.com/go-netty/constant"
	"simusco.com/go-netty/protobuf"
)

type TcpInboundHandler struct {
}

func (h *TcpInboundHandler) Handle(channel *channel.Channel, params *LiveProto.ParamVO) {
	if params == nil {
		return
	}

	if params.Event <= 0 {
		res := LiveProto.ResultVO{
			Event: constant.InvalidEventCode,
			MsgId: params.MsgId,
			Msg:   "无效的EventCode",
		}

		b, _ := proto.Marshal(&res)

		_, _ = channel.Write(b)

		return
	}

	a := actions.GetAction(params.Event)
	if a == nil {
		res := LiveProto.ResultVO{
			Event: params.Event,
			MsgId: params.MsgId,
			Code:  constant.NotImplements,
			Msg:   "事件未实现",
		}

		b, _ := proto.Marshal(&res)
		_, _ = channel.Write(b)

		return
	}

	data := a.Execute(channel, params)
	if data != nil {
		data.Event = params.Event
		data.MsgId = params.MsgId

		b, _ := proto.Marshal(data)
		_, _ = channel.Write(b)
	}

}

func (h *TcpInboundHandler) ErrorCaught(channel *channel.Channel, err error) {

}
