package goredis

import "encoding/binary"

//TODO
func b2s(b []byte) string {
	return string(b)
}
func s2b(s string) []byte {
	return []byte(s)
}

//uint2byte 将
func uint2byte(a uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, a)
	return bs
}

func byte2uint(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
