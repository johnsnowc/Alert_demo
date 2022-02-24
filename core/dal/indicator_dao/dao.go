package indicator_dao

import (
	"context"
)

type IndicatorDao interface {
	SelectIndicator(ctx context.Context, code int64) (indicator IndicatorEntity, err error)

	SelectAllIndicators(ctx context.Context) (indicators []IndicatorEntity, err error)

	AddSimpleIndicator(ctx context.Context, code int64, name string, expr string, timeRange int64) (id int64, err error)

	AddCompleteIndicator(ctx context.Context, code int64, name string, expr string, timeRange int64) (id int64, err error)

	UpdateIndicator(ctx context.Context, code int64, expr string) (id int64, err error)

	DeleteIndicator(ctx context.Context, code int64) (id int64, err error)
}
