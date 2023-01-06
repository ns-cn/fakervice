package proxy

// UdpProxy UDP的代理
type UdpProxy struct {
	Proxyable // 可代理的
}

func (proxy UdpProxy) Proxy(port int) error {
	// TODO UDP代理待实现
	return nil
}
