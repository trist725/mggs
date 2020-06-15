package internal

import (
	"github.com/rs/xid"
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
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

//todo: 心跳
func rpcNewClient(args []interface{}) {
	a := args[0].(gate.Agent)
	uuid := xid.New()
	//保留0作为cmd
	if 0 == uuid.Counter() {
		a.Close()
		return
	}
	//3字节计数器可能有1/2^24概率溢出?
	gateWay.clients[a] = uuid.Counter()

	var sa gate.Agent
	if sa = gateWay.leastConnsServer(); sa == nil {
		rpcCloseClient(args)
		return
	}
	gateWay.scMapper[sa][uuid.Counter()] = a
	a.SetUserData(sa)
}

// todo: 游服验证
func rpcNewServer(args []interface{}) {
	a := args[0].(gate.Agent)
	uuid := xid.New()
	a.SetUserData(uuid)
	gateWay.servers[a] = uuid.String()
	gateWay.scMapper[a] = make(map[int32]gate.Agent)
	log.Release("server:[%s] already online.", uuid.String())
}

func rpcCloseClient(args []interface{}) {
	a := args[0].(gate.Agent)
	if a.UserData() != nil {
		sa := a.UserData().(gate.Agent)
		id := gateWay.clients[a]
		delete(gateWay.scMapper[sa], id)
	}
	delete(gateWay.clients, a)

	a.Close()
}

func rpcCloseServer(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Release("server:[%s] offline.", gateWay.servers[a])
	delete(gateWay.servers, a)

	clients := gateWay.scMapper[a]
	for _, ca := range clients {
		args[0] = ca
		rpcCloseClient(args)
	}
	delete(gateWay.scMapper, a)
	a.Close()

}
