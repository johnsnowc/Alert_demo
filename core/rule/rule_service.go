package rule

import (
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dto"
	"Alert_demo/core/indicator"
	i "Alert_demo/core/interface"
	"context"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"sync"
)

type RuleServiceImpl struct {
}

func NewRuleServiceImpl() i.RuleService {
	return &RuleServiceImpl{}
}

var rd = rule_dao.NewRuleDaoImpl()
var ind_service = indicator.NewIndicatorServiceImpl()

func (r RuleServiceImpl) SelectRuleById(ctx context.Context, id int64) (rule *dto.Rule, err error) {
	ruleEntity, err := rd.SelectRuleById(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal([]byte(ruleEntity.Expr), rule)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (r RuleServiceImpl) SelectRuleByRoomId(ctx context.Context, roomId int64) (rules []*dto.Rule, err error) {
	ruleEntitys, err := rd.SelectRuleByRoomId(ctx, roomId)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, entity := range ruleEntitys {
		var rule dto.Rule
		err = json.Unmarshal([]byte(entity.Expr), &rule)
		if err != nil {
			log.Fatal(err)
			return
		}
		rules = append(rules, &rule)
	}
	return
}

func (r RuleServiceImpl) AddRule(ctx context.Context, name string, expr string, roomId int64) (ruleId int64, err error) {
	params := &rule_dao.RuleEntityParams{
		Name:   name,
		RoomId: roomId,
		Expr:   expr,
	}
	ruleId, err = rd.AddRule(ctx, *params)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return
}

func (r RuleServiceImpl) UpdateRule(ctx context.Context, id int64, expr string) (ruleId int64, err error) {
	var rule dto.Rule
	err = json.Unmarshal([]byte(expr), &rule)
	if err != nil {
		log.Fatal(err)
		return id, err
	}
	params := &rule_dao.RuleEntityParams{
		Name:   rule.Name,
		RoomId: rule.RoomId,
		Expr:   expr,
	}
	ruleId, err = rd.UpdateRule(ctx, id, *params)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (r RuleServiceImpl) DeleteRule(ctx context.Context, id int64) (ruleId int64, err error) {
	ruleId, err = rd.DeleteRule(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (r RuleServiceImpl) Validate(ctx context.Context, id int64) (result bool, err error) {
	rootRule, err := r.SelectRuleById(ctx, id)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	//若rule数组为空，代表为叶子结点，直接查询指标值与阈值进行校验
	if len(rootRule.Rules) == 0 {
		sourceData, err1 := ind_service.QueryData(ctx, rootRule.Indicator.Code)
		if err1 != nil {
			log.Fatal(err1)
			return false, err1
		}
		isCorrect, err1 := check(sourceData, rootRule.Value, rootRule.RelationalOp)
		if err1 != nil {
			log.Fatal(err1)
			return false, err1
		}
		return isCorrect, nil
	} else if rootRule.Logic == "||" {
		//||逻辑，若有一个为真，结果为真，调用cancel通知其他协程退出
		ctx1, cancel := context.WithCancel(ctx)
		//利用WaitGroup等待所有子规则校验完毕
		var wg sync.WaitGroup
		isCorrect := false
		for j := 0; j < len(rootRule.Rules); j++ {
			wg.Add(1)
			go func(rule *dto.Rule) {
				isCorrect1, err1 := r.Validate(ctx1, rule.Id)
				if err1 != nil {
					log.Println(err)
					err = err1
					cancel()
				}
				if isCorrect1 == true {
					isCorrect = isCorrect1
					cancel()
				}
				wg.Done()
			}(rootRule.Rules[j])
		}
		wg.Wait()
		cancel()
		if isCorrect == true {
			return true, nil
		} else {
			return false, err
		}
	} else if rootRule.Logic == "&&" {
		//&&逻辑，若有一个为假，结果为假，调用cancel通知其他协程退出
		ctx1, cancel := context.WithCancel(ctx)
		//利用WaitGroup等待所有子规则校验完毕
		var wg sync.WaitGroup
		isCorrect := true
		for j := 0; j < len(rootRule.Rules); j++ {
			wg.Add(1)
			go func(rule *dto.Rule) {
				isCorrect1, err1 := r.Validate(ctx1, rule.Id)
				if err1 != nil {
					log.Println(err1)
					err = err1
					cancel()
				}
				if isCorrect1 == false {
					isCorrect = isCorrect1
					cancel()
				}
				wg.Done()
			}(rootRule.Rules[j])
		}
		wg.Wait()
		cancel()
		if err != nil {
			return false, err
		} else {
			return isCorrect, nil
		}
	} else {
		err = errors.New("该节点非叶子结点，也不是与或关系树")
		return false, err
	}
}

func check(data, value interface{}, op string) (bool, error) {
	if reflect.TypeOf(data) != reflect.TypeOf(value) {
		err := errors.New("两个值的类型不同")
		return false, err
	}
	switch data.(type) {
	case int64:
		switch op {
		case ">":
			return data.(int64) > value.(int64), nil
		case ">=":
			return data.(int64) >= value.(int64), nil
		case "<":
			return data.(int64) < value.(int64), nil
		case "<=":
			return data.(int64) <= value.(int64), nil
		case "==":
			return data.(int64) == value.(int64), nil
		case "!=":
			return data.(int64) != value.(int64), nil
		}
	case float64:
		switch op {
		case ">":
			return data.(float64) > value.(float64), nil
		case ">=":
			return data.(float64) >= value.(float64), nil
		case "<":
			return data.(float64) < value.(float64), nil
		case "<=":
			return data.(float64) <= value.(float64), nil
		case "==":
			return data.(float64) == value.(float64), nil
		case "!=":
			return data.(float64) != value.(float64), nil
		}
	default:
		err := errors.New("非int64和float64类型")
		return false, err
	}
	err := errors.New("无法比较")
	return false, err
}
