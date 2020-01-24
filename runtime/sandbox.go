package runtime

import (
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

type Container struct {
}

type Sandbox struct {
	Config struct {
		Metadata    *cri.PodSandboxMetadata
		Annotations map[string]string
	}
	Name       string
	Uid        string
	Labels     map[string]string
	State      string
	Containers map[string]*Container
}

const SANDBOX_STATE_CREATED = "CREATED"
const SANDBOX_STATE_READY = "READY"

func NewSandbox(name string, uid string) (*Sandbox, error) {
	sb := Sandbox{
		Name:       name,
		Uid:        uid,
		State:      SANDBOX_STATE_CREATED,
		Containers: make(map[string]*Container),
	}

	return &sb, nil
}

func (s *Sandbox) IsReady() bool {
	return s.State == SANDBOX_STATE_READY
}
