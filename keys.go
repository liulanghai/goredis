package goredis

import (
	"strings"

	"github.com/tidwall/redcon"
)

/*key相关的操作*/

type gkeys struct{}

//Del  if key exist  return true
func (g gkeys) Del(db *DB, key string) bool {
	exist := g.Exist(db, key)
	if !exist {
		return exist
	}
	db.Mu.Lock()
	defer db.Mu.Unlock() //TODO
	delete(db.Dict, key)
	return exist
}

//Exist key is exist
func (g gkeys) Exist(db *DB, key string) bool {
	_, ok := db.Dict[key]
	return ok
}

//Do 处理相关的命令及数据返回
func (g gkeys) Do(con redcon.Conn, cmd redcon.Command, d *DB) {
	key := b2s(cmd.Args[1])
	switch strings.ToLower(b2s(cmd.Args[0])) {
	case "del":
		exist := g.Del(d, key)
		if exist {
			con.WriteInt(1)
		} else {
			con.WriteInt(0)
		}
	case "get":
		exist := g.Exist(d, key)
		if exist {
			con.WriteInt(1)
		} else {
			con.WriteInt(0)
		}
	default:
		con.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
	}
	return
}
func init() {
	var k gkeys
	registerCommand("del", k)
	registerCommand("exists", k)
}
