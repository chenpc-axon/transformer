package transformer

// TransformerErr 转换器错误
type TransformerErr interface {
	error
}

// IrrTransformerErr 不可恢复的转换器错误
type IrrTransformerErr struct {
	TransformerErr
}
