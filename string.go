package goredis

import (
	"strings"
	"time"

	"github.com/tidwall/redcon"
)

//RString goredis string
type RString string

const (
	//StringType string 类型
	StringType = "string"
)

//Set set
func (r RString) Set(db *DB, key, val string) error {
	db.Mu.Lock()
	defer db.Mu.Unlock() //后续优化

	k, ok := db.Dict[key]
	if ok {
		k.Val = val
		return nil
	}

	var v ValueEntry
	v.LastVisit = time.Now()
	v.Type = StringType
	v.Val = val
	db.Dict[key] = v
	return nil
}

//Get set
func (r RString) Get(db *DB, key string) (string, bool) {
	val, ok := db.Dict[key]
	if !ok {
		return "", false
	}
	return val.Val.(string), true
}

//Del  if key exist  return true
func (r RString) Del(db *DB, key string) bool {
	var exist bool
	db.Mu.Lock()
	defer db.Mu.Unlock() //TODO
	_, ok := db.Dict[key]
	if ok {
		exist = true
	}
	delete(db.Dict, key)

	return exist
}

//Do 处理相关的命令及数据返回
func (r RString) Do(con redcon.Conn, cmd redcon.Command, d *DB) {
	key := b2s(cmd.Args[1])
	switch strings.ToLower(b2s(cmd.Args[0])) {
	case "set":
		r.Set(d, key, b2s(cmd.Args[2]))
		con.WriteString(RedisOK)

	case "get":
		val, ok := r.Get(d, key)
		if ok {
			con.WriteString(val)
		} else {
			con.WriteNull()
		}
	case "del":
		exist := r.Del(d, key)
		if exist {
			con.WriteInt(1)
		} else {
			con.WriteInt(0)
		}
	}
}

func init() {
	var rs RString
	registerCommand("set", rs)
	registerCommand("get", rs)
	registerCommand("del", rs)
}
