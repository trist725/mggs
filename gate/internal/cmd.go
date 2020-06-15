package internal

import "github.com/trist725/myleaf/gate"

func init() {
	skeleton.RegisterChanRPC("ServerCloseClient", ServerCloseClient)
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
