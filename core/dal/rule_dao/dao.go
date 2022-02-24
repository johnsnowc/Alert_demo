package rule_dao

import (
	"context"
)

type RuleDao interface {
	SelectRuleById(ctx context.Context, id int64) (rule RuleEntity, err error)

	SelectRuleByRoomId(ctx context.Context, roomId int64) (rules []RuleEntity, err error)

	AddRule(ctx context.Context, params RuleEntityParams) (ruleId int64, err error)

	UpdateRule(ctx context.Context, id int64, params RuleEntityParams) (ruleId int64, err error)

	DeleteRule(ctx context.Context, id int64) (ruleId int64, err error)
}
