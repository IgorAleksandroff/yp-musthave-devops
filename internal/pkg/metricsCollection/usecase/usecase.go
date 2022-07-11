package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricsCollection"
)

type usecase struct {
	repository metricsCollection.Repository
}

func New(
	r metricsCollection.Repository,
) *usecase {
	return &usecase{
		repository: r,
	}
}
