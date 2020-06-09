package internal

import (
	"github.com/trist725/myleaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)

	skeleton.RegisterChanRPC("NewClient", rpcNewClient)
	skeleton.RegisterChanRPC("NewServer", rpcNewServer)
	skeleton.RegisterChanRPC("CloseClient", rpcCloseClient)
	skeleton.RegisterChanRPC("CloseServer", rpcCloseServer)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	//_ = a
	a.Close()
}

func rpcNewClient(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

// todo: 游服验证
func rpcNewServer(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseClient(args []interface{}) {
	a := args[0].(gate.Agent)
	a.Close()
}

func rpcCloseServer(args []interface{}) {
	a := args[0].(gate.Agent)
	a.Close()
}
