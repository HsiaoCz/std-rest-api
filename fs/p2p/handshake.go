package p2p

// HandshakeFunc 
type HandShakeFunc func(any) error

func NOPHandShakeFunc(any) error { return nil }
