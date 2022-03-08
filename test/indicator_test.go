package main

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/indicator_dao"
	"Alert_demo/core/indicator"
	"fmt"
	"testing"
)

func TestInitMySQL(t *testing.T) {
	dal.InitMySQL()
	dal.DB.AutoMigrate(&indicator_dao.IndicatorEntity{})
}

func TestInsertIndicator(t *testing.T) {
	dal.InitMySQL()
	indicatorService := indicator.NewIndicatorServiceImpl()
	indicatorService.AddSimpleIndicator(
		nil,
		"test-2022-03-06",
		"test",
		"select amount from Transaction where trading_id = 1",
		100)
}

func TestSelectIndicator(t *testing.T) {
	dal.InitMySQL()
	indicator.FindAll(nil)
	indicatorService := indicator.NewIndicatorServiceImpl()
	result, _ := indicatorService.SelectIndicator(nil, "test-2022-03-06")
	fmt.Println(result)
	fmt.Println(result.LeftChild)
	fmt.Println(result.RightChild)

}

func TestDeleteIndicator(t *testing.T) {
	dal.InitMySQL()
	indicatorService := indicator.NewIndicatorServiceImpl()
	indicatorService.DeleteIndicator(nil, "test-2022-03-06")
}

func TestQueryData(t *testing.T) {
	dal.InitMySQL()
	indicatorService := indicator.NewIndicatorServiceImpl()
	fmt.Println(indicatorService.QueryData(nil, "test-2022-03-06"))

}

func TestFindAll(t *testing.T) {
	dal.InitMySQL()
	indicator.FindAll(nil)
}

func TestUpdateIndicator(t *testing.T) {
	dal.InitMySQL()
	indicatorService := indicator.NewIndicatorServiceImpl()
	indicatorService.UpdateIndicator(nil, "test-2022-03-06", "")
}
