package main

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	"fmt"
	"log"
)

func main() {
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql succeed!")
	}
	defer dal.DB.Close()
	dal.DB.AutoMigrate(&rule_dao.RuleEntity{}, &task_dao.TaskEntity{})

	ruleDaoImpl := rule_dao.NewRuleDaoImpl()

	//rule1 := &dto.Rule{
	//	Rules:        nil,
	//	Name:         "rule3",
	//	RoomId:       3,
	//	Logic:        "",
	//	RelationalOp: "=",
	//	Indicator:    nil,
	//	Value:        100,
	//}
	//expr, _ := json.Marshal(rule1)
	//params := &rule_dao.RuleEntityParams{
	//	Name:   rule1.Name,
	//	RoomId: rule1.RoomId,
	//	Expr:   string(expr),
	//}
	//if ruleDaoImpl != nil {
	//	id, err := ruleDaoImpl.AddRule(nil, *params)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(id)
	//}
	id, err := ruleDaoImpl.SelectRuleById(nil, 4)
	log.Println(id)
}
