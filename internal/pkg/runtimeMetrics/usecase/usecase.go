package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/services/devopsServer"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics/entity"
)

type usecase struct {
	repository         runtimeMetrics.Repository
	devopsServerClient devopsServer.Client
}

func New(
	r runtimeMetrics.Repository,
	client devopsServer.Client,
) *usecase {
	r.SaveMetric("PollCount", entity.Counter(0))

	return &usecase{
		repository:         r,
		devopsServerClient: client,
	}
}
