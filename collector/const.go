package main

var (
	Version         = "v0"
	RequestIdCtxKey = &contextKey{"RequestId"}
)

const (
	Prog        = "collector"
	Description = "is a service that aggregates the information of starship"
	envPrefix   = "COLLECTOR_"
)

// copied and modified from net/http/http.go
// contextKey is a value for use with context.WithValue. It's used as
// a pointer, so it fits in an interface{} without allocation.
type contextKey struct {
	name string
}

func (k *contextKey) String() string { return Prog + " context value " + k.name }
