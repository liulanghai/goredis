package goredis

import "testing"
import . "github.com/smartystreets/goconvey/convey"

var testDB = DB{
	Dict: make(map[string]ValueEntry),
}

func TestSet(t *testing.T) {
	var s RString
	err := s.Set(&testDB, "key", "123")
	val, ok := s.Get(&testDB, "key")
	val2, ok2 := s.Get(&testDB, "key2")
	Convey("Subject: SetGet \n", t, func() {
		So(err, ShouldEqual, nil)
		So(val, ShouldEqual, "123")
		So(ok, ShouldEqual, true)
		So(ok2, ShouldEqual, false)
		So(val2, ShouldEqual, "")
	})
}
