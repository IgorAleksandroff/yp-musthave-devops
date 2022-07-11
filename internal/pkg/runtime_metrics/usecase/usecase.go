package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/services/devops_server"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"
)

type usecase struct {
	repository         runtime_metrics.Repository
	devopsServerClient devops_server.Client
}

func New(
	r runtime_metrics.Repository,
	client devops_server.Client,
) *usecase {
	r.SaveMetric("PollCount", entity.Counter(0))

	return &usecase{
		repository:         r,
		devopsServerClient: client,
	}
}
