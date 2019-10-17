package goredis

import (
	"fmt"
	"os"

	"github.com/tidwall/redcon"
)

/*
实现AOF持久化
*/

var coder binCode

//Aofer AOF持久化
type Aofer interface {
	ADD(redcon.Command) error
	Sync(*os.File) error
	Load(*os.File) (*DB, error)
}

const (
	//Head 头部
	Head = "goredisAOF"
	//Version 尾部
	Version = "V1.0"
)

//FileAof aof到文件
type FileAof struct {
	FileName string
	f        *os.File
}

/*AOF 一条命令的格式为
\r\n paramLen  param1Len param1 \r\n
2 +
*/

//ADD add cmd to aof
func (aof *FileAof) ADD(cmd redcon.Command) {
	var err error
	if aof.f == nil {
		aof.f, err = os.Create(aof.FileName)
		if err != nil {
			fmt.Printf("open failed %v", err)
			return
		}
	}
	aof.f.Write(coder.Encode(cmd))
}
