package client

import (
	mgate "github.com/trist725/mggs/gate"
	"github.com/trist725/mggs/msg"
)

func init() {
	msg.ClientProcessor.SetRouter(mgate.ChanRPC)
}
