package handlers

import (
	"Alert_demo/core/dto"
	"Alert_demo/core/rule"
	"Alert_demo/kitex_gen/api"
	"context"
	"log"
)

func SelectRuleById(ctx context.Context, req *api.SelectRuleIdRequest) (resp *api.SelectRuleResponse, err error) {
	ruleService := rule.NewRuleServiceImpl()
	var temp *dto.Rule
	if temp, err = ruleService.SelectRuleById(ctx, req.Id); err != nil {
		log.Fatal(err)
	}
	resp = new(api.SelectRuleResponse)
	resp.Name = temp.Name
	resp.Id = temp.Id
	resp.RoomId = temp.RoomId
	resp.Expr = temp.Logic
	return
}

func SelectRuleByRoomId(ctx context.Context, req *api.SelectRuleRoomIdRequest) (resp *api.SelectRuleResponse, err error) {
	ruleService := rule.NewRuleServiceImpl()
	var temp *dto.Rule
	if temp, err = ruleService.SelectRuleById(ctx, req.RoomId); err != nil {
		log.Fatal(err)
	}
	resp = new(api.SelectRuleResponse)
	resp.Name = temp.Name
	resp.Id = temp.Id
	resp.RoomId = temp.RoomId
	resp.Expr = temp.Logic
	return
}

func AddRule(ctx context.Context, req *api.AddRuleRequest) (resp *api.Response, err error) {
	ruleService := rule.NewRuleServiceImpl()
	resp = new(api.Response)
	if _, err = ruleService.AddRule(ctx, req.Name, req.Expr, req.RoomId); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func UpdateRule(ctx context.Context, req *api.UpdateRuleRequest) (resp *api.Response, err error) {
	ruleService := rule.NewRuleServiceImpl()
	resp = new(api.Response)
	if _, err = ruleService.UpdateRule(ctx, req.Id, req.Expr); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.Response, err error) {
	ruleService := rule.NewRuleServiceImpl()
	resp = new(api.Response)
	if _, err = ruleService.DeleteRule(ctx, req.Id); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}
