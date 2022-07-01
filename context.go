package transformer

// Context 转换器上下文, 每一批数据会有一个独立的 Chain
// 该上下文在 Chain 的每个 Transformer 中传递, 可用来在同一 Chain 中的不同 Transformer 中传递数据
type Context interface {
	// Put 向 Context 写入数据
	Put(key string, value interface{})
	// Get 获取指定 key 的数据
	Get(key string) (interface{}, bool)
	// Remove 删除指定 key 的数据
	Remove(key string)

	// Destroy 销毁 Context
	Destroy() error
}

// ContextFactory 创建 Context
type ContextFactory func() (Context, error)
