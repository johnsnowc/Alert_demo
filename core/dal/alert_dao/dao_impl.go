package alert_dao

import (
	"Alert_demo/core/dal"
	"context"
	"log"
)

type AlertDaoImpl struct {
}

func NewAlertDaoImpl() AlertDao {
	return &AlertDaoImpl{}
}

func (a *AlertDaoImpl) AddAlert(ctx context.Context, params AlertEntityParams) (id int64, err error) {
	alertEntity := AlertEntity(params)
	if err = dal.DB.Debug().Create(&alertEntity).Error; err != nil {
		log.Println(err)
		return -1, err
	}
	return params.Id, nil
}

func (a *AlertDaoImpl) SelectAlertByRoomId(ctx context.Context, roomId int64) (alert []AlertEntity, err error) {
	if err = dal.DB.Debug().Where("room_id = ? ", roomId).Find(&alert).Error; err != nil {
		log.Println(err)
	}
	return
}

func (a *AlertDaoImpl) SelectAlertByRuleCode(ctx context.Context, ruleCode string) (alert []AlertEntity, err error) {
	if err = dal.DB.Debug().Where("rule_code = ? ", ruleCode).Find(&alert).Error; err != nil {
		log.Println(err)
	}
	return
}
