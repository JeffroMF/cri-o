package v1

import (
	"context"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func (s *service) ListContainers(
	ctx context.Context, req *pb.ListContainersRequest,
) (*pb.ListContainersResponse, error) {
	return s.server.ListContainers(ctx, req)
}
