package _interface

import (
	"Alert_demo/core/dto"
	"context"
)

type AlertService interface {
	AddAlert(ctx context.Context, roomId int64, ruleCode string, info []dto.Info, checkTime int64) (id int64, err error)

	SelectAlertByRoomId(ctx context.Context, roomId int64) (alert []dto.Alert, err error)

	SelectAlertByRuleCode(ctx context.Context, ruleCode string) (alert []dto.Alert, err error)
}
