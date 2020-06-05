package internal

import (
	"github.com/trist725/mggs/base"
	"github.com/trist725/mggs/conf"
	"github.com/trist725/mggs/msg"
	"github.com/trist725/myleaf/gate"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Gate.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Gate.WSServerAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Gate.CertFile,
		KeyFile:         conf.Gate.KeyFile,
		TCPAddr:         conf.Gate.TCPServerAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		//AgentChanRPC:    game.ChanRPC,
	}
}

func (m *Module) OnDestroy() {

}
