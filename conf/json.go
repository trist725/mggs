package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Gate struct {
	LogLevel      string
	LogPath       string
	WSAddr        string
	CertFile      string
	KeyFile       string
	TCPClientAddr string
	TCPServerAddr string
	WSClientAddr  string
	WSServerAddr  string
	MaxConnNum    int
	ConsolePort   int
	ProfilePath   string
	MgoUrl        string
	MgoSessionNum int
	DBName        string
}

func init() {
	data, err := ioutil.ReadFile("gate.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &Gate)
	if err != nil {
		log.Fatal(err)
	}
}
