package server

import (
	"fmt"
	context "golang.org/x/net/context"
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// CreateContainer creates a new container in specified PodSandbox
func (s *Server) CreateContainer(ctx context.Context, rq *cri.CreateContainerRequest) (*cri.CreateContainerResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// StartContainer starts the container.
func (s *Server) StartContainer(ctx context.Context, rq *cri.StartContainerRequest) (*cri.StartContainerResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// StopContainer stops a running container with a grace period (i.e., timeout).
// This call is idempotent, and must not return an error if the container has
// already been stopped.
// TODO: what must the runtime do after the grace period is reached?
func (s *Server) StopContainer(ctx context.Context, rq *cri.StopContainerRequest) (*cri.StopContainerResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// RemoveContainer removes the container. If the container is running, the
// container must be forcibly removed.
// This call is idempotent, and must not return an error if the container has
// already been removed.
func (s *Server) RemoveContainer(ctx context.Context, rq *cri.RemoveContainerRequest) (*cri.RemoveContainerResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListContainers lists all containers by filters.
func (s *Server) ListContainers(ctx context.Context, rq *cri.ListContainersRequest) (*cri.ListContainersResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ContainerStatus returns status of the container. If the container is not
// present, returns an error.
func (s *Server) ContainerStatus(ctx context.Context, rq *cri.ContainerStatusRequest) (*cri.ContainerStatusResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ExecSync runs a command in a container synchronously.
func (s *Server) ExecSync(ctx context.Context, rq *cri.ExecSyncRequest) (*cri.ExecSyncResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// Exec prepares a streaming endpoint to execute a command in the container.
func (s *Server) Exec(ctx context.Context, rq *cri.ExecRequest) (*cri.ExecResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// Attach prepares a streaming endpoint to attach to a running container.
func (s *Server) Attach(ctx context.Context, rq *cri.AttachRequest) (*cri.AttachResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// PortForward prepares a streaming endpoint to forward ports from a PodSandbox.
func (s *Server) PortForward(ctx context.Context, rq *cri.PortForwardRequest) (*cri.PortForwardResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ContainerStats returns stats of the container. If the container does not
// exist, the call returns an error.
func (s *Server) ContainerStats(ctx context.Context, rq *cri.ContainerStatsRequest) (*cri.ContainerStatsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListContainerStats returns stats of all running containers.
func (s *Server) ListContainerStats(ctx context.Context, rq *cri.ListContainerStatsRequest) (*cri.ListContainerStatsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
