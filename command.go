package goredis

//SupportCMD 支持的命令
var SupportCMD = map[string]CmdExecer{}

func registerCommand(command string, exec CmdExecer) {
	SupportCMD[command] = exec
}

const (
	//RedisOK ok
	RedisOK = "OK"
	//RedisNil nil
	RedisNil = "nil"
)
