package server

import (
	"github.com/moretea/kubernix/runtime"
	"sync"
)

type Server struct {
	lock        sync.Mutex
	sandboxes   map[string]*runtime.Sandbox
	derivations map[string]*runtime.Derivation
}

func New() (*Server, error) {
	srvr := &Server{
		sandboxes:   make(map[string]*runtime.Sandbox),
		derivations: make(map[string]*runtime.Derivation),
	}

	return srvr, nil
}
