package _interface

import (
	"Alert_demo/core/dto"
	"context"
)

type IndicatorService interface {
	QueryData(ctx context.Context, code int64) (data interface{}, err error)

	SelectIndicator(ctx context.Context, code int64) (indicator *dto.Indicator, err error)

	AddSimpleIndicator(ctx context.Context, code int64, name string, expr string, timeRange int64) (id int64, err error)

	AddCompleteIndicator(ctx context.Context, code int64, name string, expr string, timeRange int64) (id int64, err error)

	UpdateIndicator(ctx context.Context, code int64, expr string) (id int64, err error)

	DeleteIndicator(ctx context.Context, code int64) (id int64, err error)
}
