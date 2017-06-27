package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	cri "k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"

	"github.com/moretea/kubernix/runtime"
)

// RunPodSandbox creates and starts a pod-level sandbox. Runtimes must ensure
// the sandbox is in the ready state on success.
func (s Server) RunPodSandbox(ctx context.Context, rq *cri.RunPodSandboxRequest) (*cri.RunPodSandboxResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	sandboxName := rq.Config.Metadata.Name
	sandboxUid := rq.Config.Metadata.Uid

	logger := log.WithFields(log.Fields{"name": sandboxName, "uid": sandboxUid})

	if _, present := s.sandboxes[sandboxUid]; present {
		logger.Error("Already created")
		return nil, fmt.Errorf("Sandbox with Uid %v already created", sandboxUid)
	}

	sb, err := runtime.NewSandbox(sandboxName, sandboxUid)

	if err != nil {
		logger.Errorf("Could not create sandbox", err)
		return nil, fmt.Errorf("Could not create sandbox", err)
	}
	sb.Labels = rq.Config.Labels
	sb.Config.Metadata = rq.Config.Metadata
	sb.Config.Annotations = rq.Config.Annotations

	s.sandboxes[sandboxUid] = sb

	response := &cri.RunPodSandboxResponse{PodSandboxId: sandboxUid}
	if log.GetLevel() == log.DebugLevel {
		logger.Debug("Started sandbox")
	}
	return response, nil
}

// StopPodSandbox stops any running process that is part of the sandbox and
// reclaims network resources (e.g., IP addresses) allocated to the sandbox.
// If there are any running containers in the sandbox, they must be forcibly
// terminated.
// This call is idempotent, and must not return an error if all relevant
// resources have already been reclaimed. kubelet will call StopPodSandbox
// at least once before calling RemovePodSandbox. It will also attempt to
// reclaim resources eagerly, as soon as a sandbox is not needed. Hence,
// multiple StopPodSandbox calls are expected.
func (s Server) StopPodSandbox(ctx context.Context, rq *cri.StopPodSandboxRequest) (*cri.StopPodSandboxResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// RemovePodSandbox removes the sandbox. If there are any running containers
// in the sandbox, they must be forcibly terminated and removed.
// This call is idempotent, and must not return an error if the sandbox has
// already been removed.
func (s Server) RemovePodSandbox(ctx context.Context, rq *cri.RemovePodSandboxRequest) (*cri.RemovePodSandboxResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// PodSandboxStatus returns the status of the PodSandbox. If the PodSandbox is not
// present, returns an error.
func (s Server) PodSandboxStatus(ctx context.Context, rq *cri.PodSandboxStatusRequest) (*cri.PodSandboxStatusResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListPodSandbox returns a list of PodSandboxes.
func (s Server) ListPodSandbox(ctx context.Context, rq *cri.ListPodSandboxRequest) (*cri.ListPodSandboxResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	log.Debugf("ListPodSandboxRequest %+v", rq)

	var potentialSandboxes []*runtime.Sandbox

	filter := rq.Filter

	// First filter by ID.
	if filter != nil && filter.Id != "" {
		sandbox := s.sandboxes[filter.Id]
		if sandbox == nil {
			potentialSandboxes = []*runtime.Sandbox{}
		} else {
			potentialSandboxes = []*runtime.Sandbox{sandbox}
		}
	} else {
		for _, sb := range s.sandboxes {
			potentialSandboxes = append(potentialSandboxes, sb)
		}
	}

	log.Debugf("Potential sandboxes: %+v", potentialSandboxes)

	var pods []*cri.PodSandbox

	// Now check if the sandbox matches the labels && state
	for _, sb := range potentialSandboxes {
		if sandboxMatchFilter(sb, filter.State, &filter.LabelSelector) {
			var podState cri.PodSandboxState
			if sb.IsReady() {
				podState = cri.PodSandboxState_SANDBOX_READY
			} else {
				podState = cri.PodSandboxState_SANDBOX_NOTREADY
			}
			pod := &cri.PodSandbox{
				Id:          sb.Uid,
				Metadata:    sb.Config.Metadata,
				State:       podState,
				Labels:      sb.Labels,
				Annotations: sb.Config.Annotations,
			}
			pods = append(pods, pod)
		}
	}

	response := &cri.ListPodSandboxResponse{
		Items: pods,
	}

	log.Debugf("ListPodSandboxResponse %+v", response)
	return response, nil
}

// Helper function, used in for loop below.
func sandboxMatchFilter(sandbox *runtime.Sandbox, filterState *cri.PodSandboxStateValue, filterLabels *map[string]string) bool {
	if filterState != nil {
		if filterState.State == cri.PodSandboxState_SANDBOX_READY {
			if !sandbox.IsReady() {
				return false
			}
		}
	}

	if filterLabels != nil {
		for filterLabel, filterLabelValue := range *filterLabels {
			sandboxLabelValue, ok := sandbox.Labels[filterLabel]
			if !ok { // this label is not present in our pod
				return false
			} else {
				if sandboxLabelValue != filterLabelValue {
					return false
				}
			}
		}
	}

	return true
}
