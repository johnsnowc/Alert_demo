package handlers

import (
	"Alert_demo/core/dto"
	"Alert_demo/core/rule"
	"Alert_demo/kitex_gen/api"
	"context"
	"encoding/json"
	"log"
	"strconv"
)

var ruleService = rule.NewRuleServiceImpl()

func SelectRuleById(ctx context.Context, req *api.SelectRuleByIdRequest) (resp *api.Response, err error) {
	var temp *dto.Rule
	if temp, err = ruleService.SelectRuleById(ctx, req.Id); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	data, _ := json.Marshal(temp)
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    string(data),
	}
	return
}

func SelectRuleByCode(ctx context.Context, req *api.SelectRuleByCodeRequest) (resp *api.Response, err error) {
	var temp *dto.Rule
	if temp, err = ruleService.SelectRuleByCode(ctx, req.Code); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	data, _ := json.Marshal(temp)
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    string(data),
	}
	return
}

func SelectRuleByRoomId(ctx context.Context, req *api.SelectRuleByRoomIdRequest) (resp *api.Response, err error) {
	var temp []*dto.Rule
	if temp, err = ruleService.SelectRuleByRoomId(ctx, req.RoomId); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	data, _ := json.Marshal(temp)
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    string(data),
	}
	return
}

func AddRule(ctx context.Context, req *api.AddRuleRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = ruleService.AddRule(ctx, req.Code, req.Name, req.Expr, req.RoomId); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    strconv.FormatInt(id, 10),
	}
	return
}

func UpdateRule(ctx context.Context, req *api.UpdateRuleRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = ruleService.UpdateRule(ctx, req.Id, req.Expr); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    strconv.FormatInt(id, 10),
	}
	return
}

func DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = ruleService.DeleteRule(ctx, req.Id); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    strconv.FormatInt(id, 10),
	}
	return
}
