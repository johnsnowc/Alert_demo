package alert

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dto"
	"fmt"
	"testing"
)

func TestAddAlert(t *testing.T) {
	dal.InitMySQL()
	alert := NewAlertServiceImpl()
	info1 := dto.Info{
		RuleCode:       "test1",
		IndicatorId:    1,
		Op:             "?",
		ActualValue:    1.0,
		ThresholdValue: 1.0,
	}
	info2 := dto.Info{
		RuleCode:       "test2",
		IndicatorId:    1,
		Op:             "?",
		ActualValue:    2.0,
		ThresholdValue: 2.0,
	}
	info := make([]dto.Info, 0)
	info = append(info, info1)
	info = append(info, info2)
	alert.AddAlert(nil, 1, "2", info, 231245)

}

func TestSelectAlert(t *testing.T) {
	dal.InitMySQL()
	alert := NewAlertServiceImpl()
	alert1, _ := alert.SelectAlertByRoomId(nil, 1)
	fmt.Println(alert1)
	alert2, _ := alert.SelectAlertByRuleCode(nil, "1")
	fmt.Println(alert2)
}
