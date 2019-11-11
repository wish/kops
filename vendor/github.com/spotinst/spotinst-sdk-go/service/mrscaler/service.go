package mrscaler

import (
	"context"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

// Service provides the API operation methods for making requests to
// endpoints of the Spotinst API. See this package's package overview docs
// for details on the service.
type Service interface {
	List(context.Context, *ListScalersInput) (*ListScalersOutput, error)
	Create(context.Context, *CreateScalerInput) (*CreateScalerOutput, error)
	Read(context.Context, *ReadScalerInput) (*ReadScalerOutput, error)
	ReadScalerCluster(context.Context, *ScalerClusterStatusInput) (*ScalerClusterStatusOutput, error)
	Update(context.Context, *UpdateScalerInput) (*UpdateScalerOutput, error)
	Delete(context.Context, *DeleteScalerInput) (*DeleteScalerOutput, error)
}

type ServiceOp struct {
	Client *client.Client
}

var _ Service = &ServiceOp{}

func New(sess *session.Session, cfgs ...*spotinst.Config) *ServiceOp {
	cfg := &spotinst.Config{}
	cfg.Merge(sess.Config)
	cfg.Merge(cfgs...)

	return &ServiceOp{
		Client: client.New(sess.Config),
	}
}
