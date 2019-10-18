package goredis

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestEncodeDecode(t *testing.T) {
	var testcoder binCode
	c := make([][]byte, 3, 3)
	c[0] = []byte("hello")
	c[1] = []byte("world")
	c[2] = []byte("1")
	res := testcoder.Encode(c)
	r, err := testcoder.Decode(res)
	Convey("Subject: SetGet \n", t, func() {
		So(err, ShouldEqual, nil)
		So(len(r), ShouldEqual, 3)
		So(string(r[0]), ShouldEqual, "hello")
		So(string(r[1]), ShouldEqual, "world")
		So(string(r[2]), ShouldEqual, "1")
	})

}
