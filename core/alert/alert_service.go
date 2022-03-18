package alert

import (
	"Alert_demo/core/dal/alert_dao"
	"Alert_demo/core/dto"
	i "Alert_demo/core/interface"
	"context"
	"encoding/json"
	"log"
)

type AlertServiceImpl struct {
}

var impl alert_dao.AlertDaoImpl

func NewAlertServiceImpl() i.AlertService {
	return &AlertServiceImpl{}
}

func (a *AlertServiceImpl) AddAlert(ctx context.Context, roomId int64, ruleCode string, info []dto.Info, checkTime int64) (id int64, err error) {
	infos, _ := json.Marshal(info)
	params := alert_dao.AlertEntityParams{
		RuleCode:  ruleCode,
		RoomId:    roomId,
		Info:      string(infos),
		CheckTime: checkTime,
	}
	if id, err = impl.AddAlert(ctx, params); err != nil {
		log.Fatal(err)
	}
	return
}

func (a *AlertServiceImpl) SelectAlertByRoomId(ctx context.Context, roomId int64) (alerts []dto.Alert, err error) {
	var entities []alert_dao.AlertEntity
	var info []dto.Info
	if entities, err = impl.SelectAlertByRoomId(ctx, roomId); err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		if err = json.Unmarshal([]byte(entity.Info), &info); err != nil {
			log.Fatal(err)
		}
		alertTemp := dto.Alert{
			Id:        entity.Id,
			RuleCode:  entity.RuleCode,
			RoomID:    entity.RoomId,
			Info:      info,
			CheckTime: entity.CheckTime,
		}
		alerts = append(alerts, alertTemp)
	}
	return
}

func (a *AlertServiceImpl) SelectAlertByRuleCode(ctx context.Context, ruleCode string) (alerts []dto.Alert, err error) {
	var entities []alert_dao.AlertEntity
	var info []dto.Info
	if entities, err = impl.SelectAlertByRuleCode(ctx, ruleCode); err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		if err = json.Unmarshal([]byte(entity.Info), &info); err != nil {
			log.Fatal(err)
		}
		alertTemp := dto.Alert{
			Id:        entity.Id,
			RuleCode:  entity.RuleCode,
			RoomID:    entity.RoomId,
			Info:      info,
			CheckTime: entity.CheckTime,
		}
		alerts = append(alerts, alertTemp)
	}
	return
}
