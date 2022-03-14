package rule

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	"Alert_demo/core/dto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
)

var ruleServiceImpl = NewRuleServiceImpl()

var simpleRule5 = &dto.Rule{
	Code:          "simpleRule5__2",
	Rules:         nil,
	Name:          "simpleRule5",
	RoomId:        2,
	Logic:         "",
	RelationalOp:  "==",
	IndicatorCode: "1",
	Value:         100,
}

var simpleRule6 = &dto.Rule{
	Code:          "simpleRule6__2",
	Rules:         nil,
	Name:          "simpleRule6",
	RoomId:        2,
	Logic:         "",
	RelationalOp:  ">=",
	IndicatorCode: "2",
	Value:         100,
}

var completeRule7 = &dto.Rule{
	Code:          "completeRule7__2",
	Rules:         []*dto.Rule{simpleRule5, simpleRule6},
	Name:          "completeRule7",
	RoomId:        2,
	Logic:         "||",
	RelationalOp:  "",
	IndicatorCode: "",
	Value:         nil,
}

func initMysql() {
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql succeed!")
	}
	dal.DB.AutoMigrate(&rule_dao.RuleEntity{}, &task_dao.TaskEntity{})
}

func TestCheck(t *testing.T) {
	var a, b, c int64 = 1, 1, 2
	var d, e float64 = 1.0, 2.0
	if ans, err := check(a, b, "=="); ans != true || err != nil {
		t.Errorf("1==1 expected be true,but something wrong")
	}
	if ans, err := check(a, d, "=="); ans != false || err == nil {
		t.Errorf("1==1.0 expected be false,but something wrong")
	}
	if ans, err := check(a, c, "=="); ans != false || err != nil {
		t.Errorf("1==2 expected be false,but something wrong")
	}
	if ans, err := check(a, c, "<="); ans != true || err != nil {
		t.Errorf("1<=2 expected be true,but something wrong")
	}
	if ans, err := check(a, e, "<="); ans != false || err == nil {
		t.Errorf("1<=1.0 expected be false,but something wrong")
	}
	if ans, err := check(c, a, "<="); ans != false || err != nil {
		t.Errorf("2<=1 expected be false,but something wrong")
	}
}

func TestRuleServiceImpl_AddRule(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	expr, _ := json.Marshal(simpleRule5)
	if ruleId, err := ruleServiceImpl.AddRule(ctx, simpleRule5.Code, simpleRule5.Name, string(expr), simpleRule5.RoomId); ruleId <= 0 || err != nil {
		t.Errorf("add rule 'simpleRule3' failed")
	}
	expr, _ = json.Marshal(simpleRule6)
	if ruleId, err := ruleServiceImpl.AddRule(ctx, simpleRule6.Code, simpleRule6.Name, string(expr), simpleRule6.RoomId); ruleId <= 0 || err != nil {
		t.Errorf("add rule 'simpleRule4' failed")
	}
	expr, _ = json.Marshal(completeRule7)
	if ruleId, err := ruleServiceImpl.AddRule(ctx, completeRule7.Code, completeRule7.Name, string(expr), completeRule7.RoomId); ruleId <= 0 || err != nil {
		t.Errorf("add rule 'completeRule1' failed")
	}
}

func TestRuleServiceImpl_SelectRuleById(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if rule, err := ruleServiceImpl.SelectRuleById(ctx, 27); err != nil {
		log.Println(err)
		log.Println(rule)
		log.Println(simpleRule5)
		t.Errorf("select rule by id failed")
	}
	if rule, err := ruleServiceImpl.SelectRuleById(ctx, 30); err != gorm.ErrRecordNotFound {
		log.Println(err)
		log.Println(rule)
		t.Errorf("select nonexistent rule by id failed")
	}
}

func TestRuleServiceImpl_SelectRuleByCode(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if rule, err := ruleServiceImpl.SelectRuleByCode(ctx, "simpleRule5__2"); err != nil {
		log.Println(err)
		log.Println(rule)
		log.Println(simpleRule5)
		t.Errorf("select rule by code failed")
	}
	if rule, err := ruleServiceImpl.SelectRuleByCode(ctx, "lalalaal"); err != gorm.ErrRecordNotFound {
		log.Println(err)
		log.Println(rule)
		t.Errorf("select nonexistent rule by id failed")
	}
}

func TestRuleServiceImpl_SelectRuleByRoomId(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if rules, err := ruleServiceImpl.SelectRuleByRoomId(ctx, 1); err != nil || len(rules) != 2 {
		log.Println(err)
		log.Println(rules)
		t.Errorf("select rules by roomid failed")
	}
}

func TestRuleServiceImpl_UpdateRule(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	updateRule := simpleRule6
	updateRule.RoomId = 3
	updateRule.Name = "simpleRule6_updated"
	expr, _ := json.Marshal(updateRule)
	if ruleId, err := ruleServiceImpl.UpdateRule(ctx, 24, string(expr)); err != nil || ruleId != 24 {
		t.Errorf("update rule failed")
	}
}

func TestRuleServiceImpl_DeleteRule(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	_, err1 := ruleServiceImpl.DeleteRule(ctx, 22)
	_, err2 := ruleServiceImpl.SelectRuleById(ctx, 22)
	if err1 != nil || err2 != gorm.ErrRecordNotFound {
		t.Errorf("delete rule failed")
	}
}

func TestRuleServiceImpl_Validate(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	result, failedId, err := ruleServiceImpl.Validate(ctx, "simpleRule5__2")
	if result != true || err != nil {
		log.Println(result, failedId, err)
		t.Errorf("validate rule simpleRule5__2 failed")
	}
	result, failedId, err = ruleServiceImpl.Validate(ctx, "completeRule7__2")
	if result != true || err != nil {
		log.Println(result, failedId, err)
		t.Errorf("validate rule completeRule7__2 failed")
	}
}
