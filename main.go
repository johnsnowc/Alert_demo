package main

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/indicator_dao"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	api "Alert_demo/kitex_gen/api/alert"
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
	dal.DB.AutoMigrate(&indicator_dao.IndicatorEntity{}, &rule_dao.RuleEntity{}, &task_dao.TaskEntity{})
	svr := api.NewServer(new(AlertImpl))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
