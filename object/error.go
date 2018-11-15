package object

const (
	ERR_OBJECT_NONE          = 13000 + iota
	ERR_OBJECT_NOT_FOUND     // 对象没有找到
	ERR_ORIGIN_NOT_FOUND     // 对象没有找到
	ERR_OBJECT_RPC_CALL      // rpc 执行错误
	ERR_OBJECT_RPC_NOT_MATCH // rpc 对象不匹配
	ERR_OBJECT_REPLICATE     // 传送对象失败
)
