package server

import (
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	cri "k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

// Status returns the status of the runtime.
func (s *Server) Status(ctx context.Context, rq *cri.StatusRequest) (*cri.StatusResponse, error) {
	log.Debugf("StatusRequest %+v", rq)

	resp := &cri.StatusResponse{
		Status: &cri.RuntimeStatus{
			Conditions: []*cri.RuntimeCondition{
				{
					Type:   cri.RuntimeReady,
					Status: true,
				},
				{
					Type:   cri.NetworkReady,
					Status: true,
				},
			},
		},
	}

	return resp, nil
}
