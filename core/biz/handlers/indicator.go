package handlers

import (
	"Alert_demo/core/dto"
	"Alert_demo/core/indicator"
	"Alert_demo/kitex_gen/api"
	"context"
	"encoding/json"
	"log"
	"strconv"
)

var indicatorService = indicator.NewIndicatorServiceImpl()

func SelectIndicator(ctx context.Context, req *api.SelectIndicatorByCodeRequest) (resp *api.Response, err error) {
	var temp dto.Indicator
	temp, err = indicatorService.SelectIndicator(ctx, req.Code)
	if err != nil {
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

func AddSimpleIndicator(ctx context.Context, req *api.AddSimpleIndicatorRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = indicatorService.AddSimpleIndicator(ctx, req.Code, req.Name, req.Expr, req.TimeRange); err != nil {
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

func AddComplexIndicator(ctx context.Context, req *api.AddComplexIndicatorRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = indicatorService.AddCompleteIndicator(ctx, req.Code, req.Name, req.Left, req.Right, req.Op, req.TimeRange); err != nil {
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

func UpdateIndicator(ctx context.Context, req *api.UpdateIndicatorRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = indicatorService.UpdateIndicator(ctx, req.TimeRange, req.Code, req.Name, req.Left, req.Right, req.Op, req.Expr); err != nil {
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

func DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = indicatorService.DeleteIndicator(ctx, req.Code); err != nil {
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
