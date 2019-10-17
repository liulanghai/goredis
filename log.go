package goredis

//Logger log
type Logger interface {
	Debug(f string, args ...interface{})
	Info(f string, args ...interface{})
	Error(f string, args ...interface{})
	Warn(f string, args ...interface{})
}
