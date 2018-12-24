package store

// 存储错误码
const (
	ERR_STORE_NONE              = 11000 + iota
	ERR_STORE_SQL               // sql 错误
	ERR_STORE_NOROW             // 没有查找到记录
	ERR_STORE_ROLE_INDEX        // 索引错误
	ERR_STORE_ROLE_NOT_FOUND    // 玩家没找到
	ERR_STORE_ROLE_STATUS_ERROR // 玩家状态错误
	ERR_STORE_ROLE_DELETED      // 角色已经删除
	ERR_STORE_ROLE_NAME         // 角色名重复
	ERR_STORE_SAVE_FAILED       // 角色存档失败
	ERR_STORE_ERROR             // 其它错误
)
