package proxy

// TcpProxy TCP的代理
type TcpProxy struct {
	Proxyable // 可代理的
}

func (proxy TcpProxy) Proxy(port int) error {
	// TODO TCP代理待实现
	return nil
}
