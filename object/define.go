package object

// object类型定义
const (
	OBJECT_TYPE_NONE         = iota
	OBJECT_TYPE_OBJECT       // 对象
	OBJECT_TYPE_GHOST        // 对象副本
	OBJECT_TYPE_SHARE        // 共享数据
	OBJECT_TYPE_SCENE_OFFSET // 场景起始
)
