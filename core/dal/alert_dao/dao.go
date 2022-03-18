package alert_dao

import "context"

type AlertDao interface {
	AddAlert(ctx context.Context, params AlertEntityParams) (id int64, err error)
	SelectAlertByRoomId(ctx context.Context, roomId int64) (alert []AlertEntity, err error)
	SelectAlertByRuleCode(ctx context.Context, ruleCode string) (alert []AlertEntity, err error)
}
