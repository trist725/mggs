package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
)

type Gateway struct {
	servers map[gate.Agent]uint32
	clients map[gate.Agent]uint64

	//server_agent-clients_agent
	scMapper map[gate.Agent][]gate.Agent
}

func NewGateway() *Gateway {
	gateway := Gateway{
		servers:  make(map[gate.Agent]uint32),
		clients:  make(map[gate.Agent]uint64),
		scMapper: make(map[gate.Agent][]gate.Agent),
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
