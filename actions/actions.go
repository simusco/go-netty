package actions

import (
	"simusco.com/go-netty/channel"
	"simusco.com/go-netty/protobuf"
)

type Action interface {
	Execute(channel *channel.Channel, params *LiveProto.ParamVO) *LiveProto.ResultVO
}

var actions map[int32]Action

func GetAction(eventCode int32) Action {
	return actions[eventCode]
}

func Register(eventCode int32, action Action) {
	actions[eventCode] = action
}
