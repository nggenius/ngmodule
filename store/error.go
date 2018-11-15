package store

// 存储错误码
const (
	ERR_STORE_NONE           = 11000 + iota
	ERR_STORE_SQL            // sql 错误
	ERR_STORE_NOROW          // 没有查找到记录
	ERR_STORE_ROLE_INDEX     // 索引错误
	ERR_STORE_ROLE_NOT_FOUND // 玩家没找到
	ERR_STORE_ERROR          // 其它错误
)
