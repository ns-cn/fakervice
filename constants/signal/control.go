package signal

type Control int

const (
	QUIT     Control = 1
	PAUSE    Control = 1 << 1
	CONTINUE Control = 1 << 2
	RESUME   Control = 1 << 3
)
