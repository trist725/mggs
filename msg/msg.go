package msg

import "github.com/trist725/mggs/processor/protobuf"

var (
	//前后端包不同的操作
	ClientProcessor = protobuf.NewClientProcessor()
	ServerProcessor = protobuf.NewServerProcessor()
)
