package indicator_dao

import (
	"context"
)

type IndicatorDao interface {
	SelectIndicator(ctx context.Context, code string) (indicator IndicatorEntity, err error)

	SelectAllIndicators(ctx context.Context) (indicators map[string]IndicatorEntity, err error)

	AddIndicator(ctx context.Context, params IndicatorEntityParams) (id int64, err error)

	UpdateIndicator(ctx context.Context, id int64, params IndicatorEntityParams) (indicatorId int64, err error)

	DeleteIndicator(ctx context.Context, code string) (id int64, err error)

	QueryData(ctx context.Context, code string, roomId int64) (result float64, err error)
}
