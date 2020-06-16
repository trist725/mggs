package base

import (
	"encoding/binary"

	"github.com/trist725/mggs/conf"
)

func ByteArrToInt32(src []byte) (dst int32) {
	if conf.LittleEndian {
		return int32(binary.LittleEndian.Uint32(src))
	}
	return int32(binary.BigEndian.Uint32(src))
}
