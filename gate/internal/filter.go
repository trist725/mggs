package internal

import (
	"encoding/binary"

	"github.com/trist725/mggs/conf"
)

type filter interface {
	do() interface{}
}

type msgFilter struct {
	clientID int32
	data     []byte
}

func NewMsgFilter(data []byte, clientID int32) *msgFilter {
	return &msgFilter{
		clientID: clientID,
		data:     data,
	}
}

func (mf *msgFilter) do() interface{} {
	return mf.insertClientID()
}

func (mf *msgFilter) insertClientID() interface{} {
	id := make([]byte, 4)
	if conf.LittleEndian {
		binary.LittleEndian.PutUint32(id, uint32(mf.clientID))
	} else {
		binary.BigEndian.PutUint32(id, uint32(mf.clientID))
	}
	var newMsg []byte
	newMsg = append(newMsg, id...)
	newMsg = append(newMsg, mf.data...)
	return newMsg
}
