package goredis

import (
	"bytes"

	"github.com/tidwall/redcon"
)

//CommandCoder 命令序列化与反序列化
type CommandCoder interface {
	Encode(redcon.Command) []byte
	Decode([]byte) (redcon.Command, error)
}

type binCode struct{}

func (b *binCode) Encode(cmd redcon.Command) []byte {
	var buffer bytes.Buffer
	buffer.Write(uint2byte((uint32)(len(cmd.Args))))
	for _, val := range cmd.Args {
		buffer.Write(uint2byte((uint32)(len(val))))
		buffer.Write(val)
	}
	return buffer.Bytes()
}

func (b *binCode) Decode(in []byte) redcon.Command {
	var cmd redcon.Command
	l := (int)(byte2uint(in[0:4])) //参数个数
	cmd.Args = make([][]byte, 0, l)
	index := 4
	for i := 0; i < l; i++ {
		t := (int)(byte2uint(in[index : index+4]))
		cmd.Args[i] = in[index+4 : index+4+t]
		index = index + 4 + t
	}
	return cmd
}
