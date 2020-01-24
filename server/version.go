package server

import (
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// Version returns the runtime name, runtime version, and runtime API version.
func (s *Server) Version(ctx context.Context, rq *cri.VersionRequest) (*cri.VersionResponse, error) {
	log.Debugf("VersionRequest %+v", rq)

	return &cri.VersionResponse{
		Version:           "0.1.0",
		RuntimeName:       "kubernix",
		RuntimeVersion:    "0.1.0",
		RuntimeApiVersion: "0.1.0",
	}, nil
}
