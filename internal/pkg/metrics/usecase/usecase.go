package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metrics"
)

type usecase struct {
	repository metrics.Repository
}

func New(
	r metrics.Repository,
) *usecase {
	return &usecase{
		repository: r,
	}
}
