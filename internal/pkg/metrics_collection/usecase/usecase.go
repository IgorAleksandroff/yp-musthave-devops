package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metrics_collection"
)

type usecase struct {
	repository metrics_collection.Repository
}

func New(
	r metrics_collection.Repository,
) *usecase {
	return &usecase{
		repository: r,
	}
}
