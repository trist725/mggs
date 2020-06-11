package server

import (
	mgate "github.com/trist725/mggs/gate"
	"github.com/trist725/mggs/msg"
)

func init() {
	msg.ServerProcessor.SetDefaultRouter(mgate.ChanRPC)
}
