package goredis

import (
	"sync"
	"time"
)

//DB goredis db
type DB struct {
	Dict map[string]ValueEntry

	//Mu 保护Sync
	Mu sync.Mutex
}

//ValueEntry 单个value ,
//为了通用,所有的val采用一样的结构，会有一定的性能损失。
type ValueEntry struct {
	Val       interface{}
	Type      string
	LastVisit time.Time
}
