package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
)

type Gateway struct {
	servers map[gate.Agent]string
	clients map[gate.Agent]int32

	scMapper map[gate.Agent]map[int32]gate.Agent
}

func NewGateway() *Gateway {
	gateway := Gateway{
		servers:  make(map[gate.Agent]string),
		clients:  make(map[gate.Agent]int32),
		scMapper: make(map[gate.Agent]map[int32]gate.Agent),
	}

	return &gateway
}

func (gw *Gateway) leastConnsServer() gate.Agent {
	if len(gw.servers) <= 0 {
		log.Error("No server are available.")
		return nil
	}

	var min int
	var agent gate.Agent
	for sa, cs := range gw.scMapper {
		if min == 0 || len(cs) <= min {
			min = len(cs)
			agent = sa
		}
	}

	return agent
}
