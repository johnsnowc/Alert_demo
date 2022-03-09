package rule_dao

import (
	"Alert_demo/core/dal"
	"context"
	"github.com/jinzhu/gorm"
	"log"
)

type RuleDaoImpl struct {
}

func NewRuleDaoImpl() RuleDao {
	return &RuleDaoImpl{}
}

func (r *RuleDaoImpl) SelectRuleById(ctx context.Context, id int64) (rule RuleEntity, err error) {
	if err = dal.DB.Debug().Where("id = ? AND is_deleted = ?", id, 0).First(&rule).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return
		}
		log.Println("rule dao select failed")
		log.Println(err)
		return
	}
	return
}

func (r *RuleDaoImpl) SelectRuleByCode(ctx context.Context, code string) (rule RuleEntity, err error) {
	if err = dal.DB.Debug().Where("code = ? AND is_deleted = ?", code, 0).First(&rule).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return
		}
		log.Println("rule dao select failed")
		log.Println(err)
		return
	}
	return
}

func (r *RuleDaoImpl) SelectRuleByRoomId(ctx context.Context, roomId int64) (rules []RuleEntity, err error) {
	if err = dal.DB.Debug().Where("room_id = ? AND is_deleted = ?", roomId, 0).Find(&rules).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (r *RuleDaoImpl) AddRule(ctx context.Context, params RuleEntityParams) (ruleId int64, err error) {
	ruleEntity := RuleEntity(params)
	if err = dal.DB.Debug().Create(&ruleEntity).Error; err != nil {
		log.Println(err)
		return -1, err
	}
	return ruleEntity.Id, nil
}

func (r *RuleDaoImpl) UpdateRule(ctx context.Context, id int64, params RuleEntityParams) (ruleId int64, err error) {
	ruleEntity := RuleEntity(params)
	if err = dal.DB.Debug().Model(&RuleEntity{}).Where("id = ? AND is_deleted = ?", id, 0).Updates(ruleEntity).Error; err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (r *RuleDaoImpl) DeleteRule(ctx context.Context, id int64) (ruleId int64, err error) {
	if err = dal.DB.Debug().Model(&RuleEntity{}).Where("id = ? AND is_deleted = ?", id, 0).Update("is_deleted", true).Error; err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}
