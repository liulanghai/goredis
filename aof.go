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
var client localClient

//Aofer AOF机制，用作持久化， 如果是远程AOF的话，则可能用作备份。
type Aofer interface {
	ADD([][]byte) error
	Sync() error
	Load() (*DB, error)
	Close() error
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

//NewFileAof 新建一个AOF
func NewFileAof(fileName string) *FileAof {
	var aof FileAof
	aof.FileName = fileName
	aof.Open = true
	return &aof
}

/*AOF 一条命令的格式为
\r\n paramLen  param1Len param1 \r\n
*/

//ADD add cmd to aof
func (aof *FileAof) ADD(cmd [][]byte) {
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
	b := coder.Encode(cmd)
	cmdlen := uint32(len(b))
	aof.f.Write(uint2byte(cmdlen))
	aof.f.Write(b)
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

//Load 从文件加载AOF恢复数据
func (aof *FileAof) Load() error {

	f, err := os.Open(aof.FileName)
	if err != nil {
		panic(fmt.Sprintf("load from aof file %s failed %v", aof.FileName, err))
	}
	head := make([]byte, HeadLen, HeadLen)
	n, err := f.Read(head)
	if err != nil {
		panic("read aof file head failed ")
	}
	if n != HeadLen {
		panic("read file head failed")
	}

	if b2s(head) != (Head + Version) {
		panic("head vaild")
	}
	for {
		cmd, err := ReadCommand(f)
		if err != nil {
			break
		}
		//本地客户端
		Hander(client, redcon.Command{Args: cmd})
	}
	return nil
}

//Close close
func (aof *FileAof) Close() error {
	if aof.f != nil {
		return aof.f.Close()
	}
	return nil
}

//ReadCommand 读cmd
func ReadCommand(f *os.File) ([][]byte, error) {
	var cmd [][]byte
	head := make([]byte, 4, 4)

	n, err := f.Read(head)
	if err != nil || n != 4 {
		return cmd, err
	}
	comLen := (int)(byte2uint(head))
	c := make([]byte, comLen, comLen)
	n, err = f.Read(c)
	if n != comLen {
		return cmd, err
	}
	return coder.Decode(c)
}
