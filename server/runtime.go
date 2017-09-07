package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	cri "k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

// UpdateRuntimeConfig updates the runtime configuration based on the given request.
func (s *Server) UpdateRuntimeConfig(ctx context.Context, rq *cri.UpdateRuntimeConfigRequest) (*cri.UpdateRuntimeConfigResponse, error) {
	log.Debugf("UpdateRuntimeConfigRequest %+v", rq)
	return nil, fmt.Errorf("not implemented")
}
