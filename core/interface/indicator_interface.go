package _interface

import (
	"Alert_demo/core/dto"
	"context"
)

type IndicatorService interface {
	QueryData(ctx context.Context, code string) (data int64, err error)

	SelectIndicator(ctx context.Context, code string) (indicator dto.Indicator, err error)

	AddSimpleIndicator(ctx context.Context, code string, name string, expr string, timeRange int64) (id int64, err error)

	AddCompleteIndicator(ctx context.Context, code string, name string, left string, right string, expr string, timeRange int64) (id int64, err error)

	UpdateIndicator(ctx context.Context, code string, expr string) (id int64, err error)

	DeleteIndicator(ctx context.Context, code string) (id int64, err error)
}
