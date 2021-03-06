package indicator

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/indicator_dao"
	"fmt"
	"log"
	"testing"
)

func TestInitMySQL(t *testing.T) {
	dal.InitMySQL()
	dal.DB.AutoMigrate(&indicator_dao.IndicatorEntity{})
}

func TestInsertIndicator(t *testing.T) {
	dal.InitMySQL()
	indicatorService := NewIndicatorServiceImpl()
	id, _ := indicatorService.AddSimpleIndicator(
		nil,
		"test-2022-03-15",
		"test1",
		"select amount from Transaction where trading_id = 2",
		100)
	log.Println(id)
}

func TestSelectIndicator(t *testing.T) {
	dal.InitMySQL()
	FindAll(nil)
	indicatorService := NewIndicatorServiceImpl()
	result, _ := indicatorService.SelectIndicator(nil, "test-2022-03-06")
	fmt.Println(result)
	fmt.Println(result.LeftChild)
	fmt.Println(result.RightChild)

}

func TestDeleteIndicator(t *testing.T) {
	dal.InitMySQL()
	indicatorService := NewIndicatorServiceImpl()
	indicatorService.DeleteIndicator(nil, "test-2022-03-06")
}

func TestQueryData(t *testing.T) {
	dal.InitMySQL()
	indicatorService := NewIndicatorServiceImpl()
	fmt.Println(indicatorService.QueryData(nil, "per_customer_transaction", 1))
}

func TestFindAll(t *testing.T) {
	dal.InitMySQL()
	FindAll(nil)
}

//func TestUpdateIndicator(t *testing.T) {
//	dal.InitMySQL()
//	indicatorService := NewIndicatorServiceImpl()
//	indicatorService.UpdateIndicator(nil, "test-2022-03-06", "")
//}
