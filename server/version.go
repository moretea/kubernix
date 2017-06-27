package server

import (
	context "golang.org/x/net/context"
	cri "k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

// Version returns the runtime name, runtime version, and runtime API version.
func (s Server) Version(ctx context.Context, rq *cri.VersionRequest) (*cri.VersionResponse, error) {

	return &cri.VersionResponse{
		Version:           "0.1.0",
		RuntimeName:       "kubernix",
		RuntimeVersion:    "0.1.0",
		RuntimeApiVersion: "0.1.0",
	}, nil
}
