package rule

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	"Alert_demo/core/dto"
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

var ruleServiceImpl = NewRuleServiceImpl()

var perCustomerTransactionGreaterThan200 = &dto.Rule{
	Code:          "ATV>200",
	Rules:         nil,
	Name:          "perCustomerTransactionGreaterThan200",
	RoomId:        1,
	Logic:         "",
	RelationalOp:  ">",
	IndicatorCode: "per_customer_transaction",
	Value:         200,
}

var userNumsGreaterThan5 = &dto.Rule{
	Code:          "userNums>5",
	Rules:         nil,
	Name:          "userNumsGreaterThan5",
	RoomId:        1,
	Logic:         "",
	RelationalOp:  ">",
	IndicatorCode: "user_nums",
	Value:         5,
}

var completeRule1 = &dto.Rule{
	Code:          "completeRule1",
	Rules:         []*dto.Rule{perCustomerTransactionGreaterThan200, userNumsGreaterThan5},
	Name:          "completeRule1",
	RoomId:        1,
	Logic:         "||",
	RelationalOp:  "",
	IndicatorCode: "",
	Value:         0,
}

var amountGreaterThan500 = &dto.Rule{
	Code:          "amountGreaterThan500",
	Rules:         nil,
	Name:          "amountGreaterThan500",
	RoomId:        1,
	Logic:         "",
	RelationalOp:  ">",
	IndicatorCode: "amount",
	Value:         500,
}

var completeRule2 = &dto.Rule{
	Code:          "completeRule2",
	Rules:         []*dto.Rule{completeRule1, amountGreaterThan500},
	Name:          "completeRule2",
	RoomId:        1,
	Logic:         "&&",
	RelationalOp:  "",
	IndicatorCode: "",
	Value:         0,
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

func TestRuleServiceImpl_AddRule(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	expr, _ := json.Marshal(completeRule2)
	if ruleId, err := ruleServiceImpl.AddRule(ctx, completeRule2.Code, completeRule2.Name, string(expr), completeRule2.RoomId); ruleId <= 0 || err != nil {
		t.Errorf("add rule 'completeRule1' failed")
	}
}

//func TestRuleServiceImpl_SelectRuleById(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	ctx := context.Background()
//	if rule, err := ruleServiceImpl.SelectRuleById(ctx, 27); err != nil {
//		log.Println(err)
//		log.Println(rule)
//		log.Println(simpleRule5)
//		t.Errorf("select rule by id failed")
//	}
//	if rule, err := ruleServiceImpl.SelectRuleById(ctx, 30); err != gorm.ErrRecordNotFound {
//		log.Println(err)
//		log.Println(rule)
//		t.Errorf("select nonexistent rule by id failed")
//	}
//}
//
//func TestRuleServiceImpl_SelectRuleByCode(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	ctx := context.Background()
//	if rule, err := ruleServiceImpl.SelectRuleByCode(ctx, "simpleRule5__2"); err != nil {
//		log.Println(err)
//		log.Println(rule)
//		log.Println(simpleRule5)
//		t.Errorf("select rule by code failed")
//	}
//	if rule, err := ruleServiceImpl.SelectRuleByCode(ctx, "lalalaal"); err != gorm.ErrRecordNotFound {
//		log.Println(err)
//		log.Println(rule)
//		t.Errorf("select nonexistent rule by id failed")
//	}
//}
//
//func TestRuleServiceImpl_SelectRuleByRoomId(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	ctx := context.Background()
//	if rules, err := ruleServiceImpl.SelectRuleByRoomId(ctx, 1); err != nil || len(rules) != 2 {
//		log.Println(err)
//		log.Println(rules)
//		t.Errorf("select rules by roomid failed")
//	}
//}
//
//func TestRuleServiceImpl_UpdateRule(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	ctx := context.Background()
//	updateRule := simpleRule6
//	updateRule.RoomId = 3
//	updateRule.Name = "simpleRule6_updated"
//	expr, _ := json.Marshal(updateRule)
//	if ruleId, err := ruleServiceImpl.UpdateRule(ctx, 24, string(expr)); err != nil || ruleId != 24 {
//		t.Errorf("update rule failed")
//	}
//}
//
//func TestRuleServiceImpl_DeleteRule(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	ctx := context.Background()
//	_, err1 := ruleServiceImpl.DeleteRule(ctx, 22)
//	_, err2 := ruleServiceImpl.SelectRuleById(ctx, 22)
//	if err1 != nil || err2 != gorm.ErrRecordNotFound {
//		t.Errorf("delete rule failed")
//	}
//}
//
//func TestRuleServiceImpl_Validate(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	ctx := context.Background()
//	result, failedId, err := ruleServiceImpl.Validate(ctx, "simpleRule5__2")
//	if result != true || err != nil {
//		log.Println(result, failedId, err)
//		t.Errorf("validate rule simpleRule5__2 failed")
//	}
//	result, failedId, err = ruleServiceImpl.Validate(ctx, "completeRule7__2")
//	if result != true || err != nil {
//		log.Println(result, failedId, err)
//		t.Errorf("validate rule completeRule7__2 failed")
//	}
//}
