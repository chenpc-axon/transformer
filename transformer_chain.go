package transformer

import dg "chenpc.com/axon/apis/datagram"

// Chain 转换器链
type Chain interface {
	Process(ctx Context, datagrams []dg.Datagram, err error)
}
