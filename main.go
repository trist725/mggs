package main

import (
	"github.com/trist725/mggs/client"
	"github.com/trist725/mggs/conf"
	"github.com/trist725/mggs/server"
	leaf "github.com/trist725/myleaf"
	lconf "github.com/trist725/myleaf/conf"
)

func main() {
	lconf.LogLevel = conf.Gate.LogLevel
	lconf.LogPath = conf.Gate.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Gate.ConsolePort
	lconf.ProfilePath = conf.Gate.ProfilePath

	leaf.Run(
		client.Module,
		server.Module,
	)
}
