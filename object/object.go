package object

import (
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/core/service"
)

// 对象创建接口
type ObjectCreate interface {
	Ctor()              // 构造函数
	ObjectType() string // Object type
}

type Object interface {
	// ObjectType 获取对象类型
	ObjectType() string
	// ObjId 唯一ID
	ObjId() rpc.Mailbox
	// SetObjId 设置唯一ID
	SetObjId(id rpc.Mailbox)
	// Factory 所属的工厂
	Factory() *Factory
	// SetFactory 设置工厂，由工厂主动调用
	SetFactory(f *Factory)
	// Core 获取Core接口
	Core() service.CoreAPI
	// SetCap 设置对象容量
	SetCap(int)
}
