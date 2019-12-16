package utils

import (
	"bytes"
	"encoding/binary"
)

func Pack(data [][]byte) []byte {
	var buf = new(bytes.Buffer)
	for _, v := range data {
		binary.Write(buf, binary.BigEndian, int32(len(v)))
		buf.Write(v)
	}
	return buf.Bytes()
}

func PackOne(data []byte) []byte {
	var buf = new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, int32(len(data)))
	buf.Write(data)
	return buf.Bytes()
}
