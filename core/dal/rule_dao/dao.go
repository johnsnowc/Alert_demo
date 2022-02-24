package rule_dao

import (
	"context"
)

type RuleDao interface {
	SelectRuleById(ctx context.Context, id int64) (rule RuleEntity, err error)

	SelectRuleByRoomId(ctx context.Context, roomId int64) (rules []RuleEntity, err error)

	AddRule(ctx context.Context, name string, ruleType bool, expr string, roomId int64) (ruleId int64, err error)

	UpdateRule(ctx context.Context, id int64, expr string) (ruleId int64, err error)

	DeleteRule(ctx context.Context, id int64) (ruleId int64, err error)
}
