package internal

import (
	"github.com/trist725/mggs/base"
	"github.com/trist725/myleaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("ServerCommand", ServerCommand)
}

func ServerCommand(args []interface{}) {
	//sa := args[1].(gate.Agent)
	cmdType := args[2].(uint16)
	//0-关闭客户端
	//todo: 1-握手
	switch cmdType {
	case 0:
		args[2] = base.ByteArrToInt32(args[0].([]byte))
		ServerCloseClient(args)
	case 1:
		//ServerHandShake()
	default:
	}
}

//close client
func ServerCloseClient(args []interface{}) {
	//msg := args[0]
	sa := args[1].(gate.Agent)
	clientID := args[2].(int32)
	if ca, ok := gateWay.scMapper[sa][clientID]; ok {
		args[0] = ca
		rpcCloseClient(args)
	}
}
