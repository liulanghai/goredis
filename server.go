package goredis

import (
	"strings"

	"github.com/tidwall/redcon"
)

//ServerConfig todo
type ServerConfig struct {
}

//Server todo
type Server struct {
	//端口
	Port int64
}

//CmdExecer 命令执行器
type CmdExecer interface {
	Do(redcon.Conn, redcon.Command, *DB)
}

//Hander 处理请求
func Hander(conn redcon.Conn, cmd redcon.Command) {
	c := strings.ToLower(b2s(cmd.Args[0]))
	execer, ok := SupportCMD[c]
	if !ok {
		conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
		return
	}
	execer.Do(conn, cmd, StringDB)
}

//StringDB string
var StringDB *DB

func init() {
	StringDB = &DB{
		Dict: make(map[string]ValueEntry),
	}

}

//StartServre 开始服务
func StartServre(bindAddr string) {

	err := redcon.ListenAndServe(bindAddr,
		Hander,

		func(conn redcon.Conn) bool {
			// use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// this is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		panic(err.Error())
	}

}
