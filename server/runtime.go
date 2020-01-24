package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// UpdateRuntimeConfig updates the runtime configuration based on the given request.
func (s *Server) UpdateRuntimeConfig(ctx context.Context, rq *cri.UpdateRuntimeConfigRequest) (*cri.UpdateRuntimeConfigResponse, error) {
	log.Debugf("UpdateRuntimeConfigRequest %+v", rq)
	return nil, fmt.Errorf("not implemented")
}
