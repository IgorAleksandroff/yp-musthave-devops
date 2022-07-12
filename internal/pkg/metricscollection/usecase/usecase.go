package usecase

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection"
)

type usecase struct {
	repository metricscollection.Repository
}

func New(
	r metricscollection.Repository,
) *usecase {
	return &usecase{
		repository: r,
	}
}
