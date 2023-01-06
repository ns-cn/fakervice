package proxy

type Protocol int

const (
	TCP Protocol = iota
	UDP
	HTTP
	HTTPS
)
