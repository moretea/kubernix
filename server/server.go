package server

import (
	"github.com/moretea/kubernix/runtime"
	"sync"
)

type Server struct {
	lock      sync.Mutex
	sandboxes map[string]*runtime.Sandbox
}

func New() (*Server, error) {
	srvr := &Server{
		sandboxes: make(map[string]*runtime.Sandbox),
	}

	return srvr, nil
}
