package handlers

import (
	"Alert_demo/core/dto"
	"Alert_demo/core/indicator"
	"Alert_demo/kitex_gen/api"
	"context"
	"log"
)

func SelectIndicator(ctx context.Context, req *api.SelectIndicatorRequest) (resp *api.SelectIndicatorResponse, err error) {
	indicatorService := indicator.NewIndicatorServiceImpl()
	var temp dto.Indicator
	temp, err = indicatorService.SelectIndicator(ctx, req.Code)
	if err != nil {
		log.Fatal(err)
	}
	resp = new(api.SelectIndicatorResponse)
	resp.Name = temp.Name
	resp.Code = temp.Code
	resp.Op = temp.ArithmeticOp
	resp.TimeRange = temp.TimeRange
	resp.Expr = temp.ObjectType.Expr
	return
}

func AddSimpleIndicator(ctx context.Context, req *api.AddSimpleIndicatorRequest) (resp *api.Response, err error) {
	indicatorService := indicator.NewIndicatorServiceImpl()
	resp = new(api.Response)
	if _, err = indicatorService.AddSimpleIndicator(ctx, req.Code, req.Name, req.Expr, req.TimeRange); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func AddComplexIndicator(ctx context.Context, req *api.AddComplexIndicatorRequest) (resp *api.Response, err error) {
	indicatorService := indicator.NewIndicatorServiceImpl()
	resp = new(api.Response)
	if _, err = indicatorService.AddCompleteIndicator(ctx, req.Code, req.Name, req.Left, req.Right, req.Op, req.TimeRange); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func UpdateIndicator(ctx context.Context, req *api.UpdateIndicatorRequest) (resp *api.Response, err error) {
	indicatorService := indicator.NewIndicatorServiceImpl()
	resp = new(api.Response)
	if _, err = indicatorService.UpdateIndicator(ctx, req.TimeRange, req.Code, req.Name, req.Left, req.Right, req.Op); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest) (resp *api.Response, err error) {
	indicatorService := indicator.NewIndicatorServiceImpl()
	resp = new(api.Response)
	if _, err = indicatorService.DeleteIndicator(ctx, req.Code); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}
