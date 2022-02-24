package indicator_dao

import (
	"context"
)

type IndicatorDao interface {
	SelectIndicator(ctx context.Context, code int64) (indicator IndicatorEntity, err error)

	SelectAllIndicators(ctx context.Context) (indicators []IndicatorEntity, err error)

	AddIndicator(ctx context.Context, params IndicatorEntityParams) (id int64, err error)

	UpdateIndicator(ctx context.Context, code int64, params IndicatorEntityParams) (id int64, err error)

	DeleteIndicator(ctx context.Context, code int64) (id int64, err error)
}
