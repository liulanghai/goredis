package goredis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAof(t *testing.T) {
	aof := NewFileAof("gredis_test.aof")
	cmd := make([][]byte, 3, 3)
	cmd[0] = []byte("set")
	cmd[1] = []byte("name")
	cmd[2] = []byte("dhg")
	aof.ADD(cmd)

	cmd[0] = []byte("set")
	cmd[1] = []byte("age")
	cmd[2] = []byte("26")
	aof.ADD(cmd)

	cmd[0] = []byte("set")
	cmd[1] = []byte("mid")
	cmd[2] = []byte("170")
	aof.ADD(cmd)

	aof.Sync()
	aof.Close()

	var rs RString
	err := aof.Load()
	Convey("Subject: Load \n", t, func() {
		So(err, ShouldEqual, nil)
	})
	name, _ := rs.Get(StringDB, "name")
	age, _ := rs.Get(StringDB, "age")
	mid, _ := rs.Get(StringDB, "mid")
	Convey("Subject: SetGet \n", t, func() {
		So(err, ShouldEqual, nil)
		So(name, ShouldEqual, "dhg")
		So(age, ShouldEqual, "26")
		So(mid, ShouldEqual, "170")
	})
}
