package object

import (
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/core/service"
)

const (
	EXPOSE_NONE  = 0
	EXPOSE_OWNER = 1
	EXPOSE_OTHER = 2
	EXPOSE_ALL   = EXPOSE_OWNER | EXPOSE_OTHER
)

// 对象创建接口
type ObjectCreate interface {
	Ctor()              // 构造函数
	EntityType() string // Entity type
}

// 缓存接口
type Cacher interface {
	// 缓存kv
	Cache(key string, value interface{})
	// 获取value
	Value(key string) interface{}
	// 获取value并返回是否存在
	TryGetValue(key string) (interface{}, bool)
	// 是否存在key
	HasKey(key string) bool
	// 删除key
	DeleteCache(key string)
	// 删除所有key
	ClearAllCache()
}

type Object interface {
	// set id
	SetId(val int64)
	// db id
	DBId() int64
	// ObjId 唯一ID
	ObjId() rpc.Mailbox
	// SetObjId 设置唯一ID
	SetObjId(id rpc.Mailbox)
	// Archive 存档对象
	Archive() interface{}
	// Silence 沉默状态
	Silence() bool
	// SetSilence 设置沉默状态
	SetSilence(bool)
	// Dummy 是否是复制对象
	Dummy() bool
	// SetDummy 设置为复制对象
	SetDummy(c bool)
	// Sync 同步状态
	Sync() bool
	// SetSync 设置同步状态
	SetSync(bool)
	// Original 原始对象
	Original() *rpc.Mailbox
	// SetOriginal 设置原始对象
	SetOriginal(m *rpc.Mailbox)
	// Factory 所属的工厂
	Factory() *Factory
	// SetFactory 设置工厂，由工厂主动调用
	SetFactory(f *Factory)
	// Core 获取Core接口
	Core() service.CoreAPI
	// Type 类型(对应xml里面的type)
	Type() string
	// Entity 类型(对应xml里面的name)
	Entity() string
	// AttrType 获取某个属性的类型
	AttrType(name string) string
	// FindAttr 获取属性
	FindAttr(name string) interface{}
	// SetAttr 设置属性
	SetAttr(name string, value interface{}) error
	// Expose 导出状态
	Expose(name string) int
	// AllAttr 所有属性名
	AllAttr() []string
	// AttrIndex 属性的索引编号
	AttrIndex(name string) int
	// AddAttrObserver 增加一个属性观察者
	AddAttrObserver(name string, observer attrObserver) error
	// RemoveAttrObserver 删除属性观察者
	RemoveAttrObserver(name string)
	// AddTableObserver 增加表格观察者
	AddTableObserver(name string, observer tableObserver) error
	// RemoveTableObserver 删除表格观察者
	RemoveTableObserver(name string)
	// UpdateAttr 属性变动回调
	UpdateAttr(name string, val interface{}, old interface{})
	// UpdateTuple tuple变动回调
	UpdateTuple(name string, val interface{}, old interface{})
	// AddTableRow 表格增加行回调
	AddTableRow(name string, row int)
	// AddTableRowValue 表格增加行并设置值回调
	AddTableRowValue(name string, row int, val ...interface{})
	// SetTableRowValue 设置表格行
	SetTableRowValue(name string, row int, val ...interface{})
	// DelTableRow 删除表格行
	DelTableRow(name string, row int)
	// ClearTable 清除表格
	ClearTable(name string)
	// ChangeTable 表格单元格变动
	ChangeTable(name string, row, col int, val interface{})
	// ExistDummy 是否存在某个副本对象
	ExistDummy(dummy rpc.Mailbox) bool
	// AddDummy 关联一个副本对象
	AddDummy(dummy rpc.Mailbox, state int)
	// RemoveDummy 移除一个副本对象
	RemoveDummy(dummy rpc.Mailbox)
	// ChangeDummyState 更新副本对象的状态
	ChangeDummyState(dummy rpc.Mailbox, state int) error
}
