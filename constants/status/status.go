package status

type Status int

const (
	INIT       Status = 1
	RUNNING    Status = 1 << 1
	PAUSED     Status = 1 << 2
	TERMINATED Status = 1 << 3
)
