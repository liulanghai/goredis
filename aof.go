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

//Aofer AOF机制，用作持久化， 如果是远程AOF的话，则可能用作备份。
type Aofer interface {
	ADD(redcon.Command) error
	Sync() error
	Load() (*DB, error)
}

const (
	//Head 头部
	Head = "goredis"
	//Version 尾部
	Version = "V1.0"
	//HeadLen len
	HeadLen = len(Head + Version)
)

//FileAof aof到文件
type FileAof struct {
	FileName string
	f        *os.File
	log      Logger
	Open     bool
}

/*AOF 一条命令的格式为
\r\n paramLen  param1Len param1 \r\n
2 +
*/

//ADD add cmd to aof
func (aof *FileAof) ADD(cmd redcon.Command) {
	var err error
	if !aof.Open {
		return
	}
	if aof.f == nil {
		aof.f, err = os.Create(aof.FileName)
		if err != nil {
			fmt.Printf("open failed %v", err)
			return
		}
		n, err := aof.f.WriteString(Head + Version)
		if n != HeadLen {
			aof.log.Error("write aof file head failed %v", err)
			return
		}
	}
	aof.f.Write(coder.Encode(cmd))
}

//Sync 同步
func (aof *FileAof) Sync() error {
	if aof.f == nil || aof.Open == false {
		return nil
	}
	err := aof.f.Sync()
	if err != nil {
		aof.log.Error("sync failed %v", err)
	}
	return err
}
