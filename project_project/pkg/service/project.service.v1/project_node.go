package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"test.com/project_common/errs"
	"test.com/project_grpc/project"
)

func (p *ProjectService) NodeList(c context.Context, msg *project.ProjectRpcMessage) (*project.ProjectNodeResponseMessage, error) {
	list, err := p.nodeDomain.TreeList()
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var nodes []*project.ProjectNodeMessage
	_ = copier.Copy(&nodes, list)
	return &project.ProjectNodeResponseMessage{Nodes: nodes}, nil
}
