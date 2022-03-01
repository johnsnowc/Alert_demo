package _interface

import (
	"Alert_demo/core/dto"
	"context"
)

type RuleService interface {
	SelectRuleById(ctx context.Context, id int64) (rule *dto.Rule, err error)

	SelectRuleByRoomId(ctx context.Context, roomId int64) (rules []*dto.Rule, err error)

	AddRule(ctx context.Context, name string, expr string, roomId int64) (ruleId int64, err error)

	UpdateRule(ctx context.Context, id int64, expr string) (ruleId int64, err error)

	DeleteRule(ctx context.Context, id int64) (ruleId int64, err error)

	Validate(ctx context.Context, id int64) (result bool, err error)
}
