package proxy

// HttpProxy HTTP代理
type HttpProxy struct {
	Proxyable // 可代理的
}

func (proxy HttpProxy) Proxy(port int) error {
	// TODO HTTP代理待实现
	return nil
}
