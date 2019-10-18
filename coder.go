package goredis

import (
	"bytes"
)

//CommandCoder 命令序列化与反序列化
type CommandCoder interface {
	Encode([][]byte) []byte
	Decode([]byte) ([][]byte, error)
}

type binCode struct{}

func (b *binCode) Encode(cmd [][]byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(uint2byte((uint32)(len(cmd))))
	for _, val := range cmd {
		buffer.Write(uint2byte((uint32)(len(val))))
		buffer.Write(val)
	}
	return buffer.Bytes()
}

func (b *binCode) Decode(in []byte) ([][]byte, error) {

	l := (int)(byte2uint(in[0:4])) //参数个数
	cmd := make([][]byte, l, l)
	index := 4
	for i := 0; i < l; i++ {
		t := (int)(byte2uint(in[index : index+4]))
		cmd[i] = in[index+4 : index+4+t]
		index = index + 4 + t
	}
	return cmd, nil
}
