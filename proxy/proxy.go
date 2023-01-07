package proxy

// Proxyable 可代理的
type Proxyable interface {
	Proxy(int, string) error
}
