package transformer

import dg "chenpc.com/axon/apis/datagram"

// Transformer 转换器
type Transformer interface {
	// Key 唯一标识
	Key() string
	// Name 转换器的名称
	Name() string
	// Init 初始化转换器
	Init() error
	// Destroy 销毁转换器
	Destroy()
	// Handle 处理 Datagram
	Handle(ctx Context, datagrams []dg.Datagram, chain Chain)
}
