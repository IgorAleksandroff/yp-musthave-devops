package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics"
)

type usecase struct {
	repository runtime_metrics.Repository
}

func New(
	r runtime_metrics.Repository,
) *usecase {
	return &usecase{
		repository: r,
	}
}
