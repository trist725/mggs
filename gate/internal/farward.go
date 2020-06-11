package internal

import (
	"github.com/trist725/mggs/conf"
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
)

func init() {
	skeleton.RegisterChanRPC("ClientForward", ClientForward)
	skeleton.RegisterChanRPC("ServerForward", ServerForward)
}

//to server
func ClientForward(args []interface{}) {
	msg := args[0].([]byte)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		log.Error("invalid server agent with client agent id:[%d]", gateWay.clients[a])
		return
	}

	var sa gate.Agent
	sa = a.UserData().(gate.Agent)
	sa.WriteMsg([][]byte{util.Int32ToByteArr(gateWay.clients[a], conf.LittleEndian), msg})
}

//to client
func ServerForward(args []interface{}) {
	msg := args[0].(gate.Agent)
	sa := args[1].(gate.Agent)
	clientID := args[2].(int32)
	if ca, ok := gateWay.scMapper[sa][clientID]; ok {
		ca.WriteMsg(msg)
	}

}
