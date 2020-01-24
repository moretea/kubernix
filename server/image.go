package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// ListImages lists existing images.
func (s *Server) ListImages(ctx context.Context, rq *cri.ListImagesRequest) (*cri.ListImagesResponse, error) {
	log.Debugf("ListImagesRequest %+v", rq)
	filter := ""

	if rq.Filter != nil {
		if rq.Filter.Image != nil {
			filter = rq.Filter.Image.Image
		}
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	response := cri.ListImagesResponse{}

	if filter != "" {
		derivation, ok := s.derivations[filter]
		if ok {
			response.Images = append(response.Images, &cri.Image{
				Id:       derivation.Id,
				RepoTags: []string{derivation.Id},
				Size_:    derivation.Size,
			})
		}
	} else {
		for id, derivation := range s.derivations {
			response.Images = append(response.Images, &cri.Image{
				Id:       id,
				RepoTags: []string{id},
				Size_:    derivation.Size,
			})
		}
	}

	return &response, nil
}

// ImageStatus returns the status of the image. If the image is not
// present, returns a response with ImageStatusResponse.
// Image set to nil
func (s *Server) ImageStatus(ctx context.Context, rq *cri.ImageStatusRequest) (*cri.ImageStatusResponse, error) {
	log.Debugf("ImageStatusRequest %+v", rq)
	return nil, fmt.Errorf("not implemented")
}

// PullImage pulls an image with authentication config.
func (s *Server) PullImage(ctx context.Context, rq *cri.PullImageRequest) (*cri.PullImageResponse, error) {
  image := rq.Image.Image
  s.pullDerivation(image)

	//	return nil, fmt.Errorf("not implemented")
	resp := &cri.PullImageResponse{
		ImageRef: image,
	}

	return resp, nil
}

// RemoveImage removes the image.
// This call is idempotent, and must not return an error if the image has
// already been removed.
func (s *Server) RemoveImage(ctx context.Context, rq *cri.RemoveImageRequest) (*cri.RemoveImageResponse, error) {
	log.Debugf("PullImageRequest %+v", rq)
	return nil, fmt.Errorf("not implemented")
}

// ImageFSInfo returns information of the filesystem that is used to store images.
func (s *Server) ImageFsInfo(ctx context.Context, rq *cri.ImageFsInfoRequest) (*cri.ImageFsInfoResponse, error) {
	log.Debugf("PullImageRequest %+v", rq)
	return nil, fmt.Errorf("not implemented")
}
