package goredis

import (
	"net"

	"github.com/tidwall/redcon"
)

/*
内部redis 客户端，实现 redcon.Conn
*/

type localClient struct{}

func (c localClient) RemoteAddr() string {
	return "127.0.0.1"
}
func (c localClient) Close() error {
	return nil
}
func (c localClient) WriteError(s string) {
	return
}

func (c localClient) WriteString(str string) {
	return
}
func (c localClient) WriteBulk(bulk []byte) {
	return
}
func (c localClient) WriteBulkString(bulk string) {
	return
}
func (c localClient) WriteInt(num int) {
	return
}
func (c localClient) WriteInt64(num int64) {
	return
}

func (c localClient) WriteNull() {
	return
}
func (c localClient) WriteArray(count int) {
	return
}

func (c localClient) WriteRaw(data []byte) {
	return
}
func (c localClient) Context() interface{} {
	return nil
}
func (c localClient) SetContext(v interface{}) {
	return
}
func (c localClient) SetReadBuffer(bytes int) {
	return
}

func (c localClient) ReadPipeline() []redcon.Command {
	return []redcon.Command{}
}

func (c localClient) PeekPipeline() []redcon.Command {
	return []redcon.Command{}
}
func (c localClient) NetConn() net.Conn {
	var conn net.Conn
	return conn
}
func (d localClient) Detach() redcon.DetachedConn {

	return localDetach{}
}

type localDetach struct {
	redcon.Conn
}

func (d localDetach) ReadCommand() (redcon.Command, error) {
	return redcon.Command{}, nil
}
func (d localDetach) Flush() error {
	return nil
}
