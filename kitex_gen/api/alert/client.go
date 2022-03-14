// Code generated by Kitex v0.0.8. DO NOT EDIT.

package alert

import (
	"Alert_demo/kitex_gen/api"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SelectIndicator(ctx context.Context, req *api.SelectIndicatorRequest, callOptions ...callopt.Option) (r *api.SelectIndicatorResponse, err error)
	AddSimpleIndicator(ctx context.Context, req *api.AddSimpleIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	AddComplexIndicator(ctx context.Context, req *api.AddComplexIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	UpdateIndicator(ctx context.Context, req *api.UpdateIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	SelectRuleById(ctx context.Context, req *api.SelectRuleIdRequest, callOptions ...callopt.Option) (r *api.SelectRuleResponse, err error)
	SelectRuleByRoomId(ctx context.Context, req *api.SelectRuleRoomIdRequest, callOptions ...callopt.Option) (r *api.SelectRuleResponse, err error)
	AddRule(ctx context.Context, req *api.AddRuleRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	UpdateRule(ctx context.Context, req *api.UpdateRuleRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	DeleteRule(ctx context.Context, req *api.DeleteRuleRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	SelectTaskById(ctx context.Context, req *api.SelectTaskIdRequest, callOptions ...callopt.Option) (r *api.SelectTaskResponse, err error)
	SelectTaskByRoomId(ctx context.Context, req *api.SelectTaskRoomIdRequest, callOptions ...callopt.Option) (r *api.SelectTasksResponse, err error)
	AddTask(ctx context.Context, req *api.AddTaskRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	UpdateTask(ctx context.Context, req *api.UpdateTaskRequest, callOptions ...callopt.Option) (r *api.Response, err error)
	DeleteTask(ctx context.Context, req *api.DeleteTaskRequest, callOptions ...callopt.Option) (r *api.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAlertClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAlertClient struct {
	*kClient
}

func (p *kAlertClient) SelectIndicator(ctx context.Context, req *api.SelectIndicatorRequest, callOptions ...callopt.Option) (r *api.SelectIndicatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectIndicator(ctx, req)
}

func (p *kAlertClient) AddSimpleIndicator(ctx context.Context, req *api.AddSimpleIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddSimpleIndicator(ctx, req)
}

func (p *kAlertClient) AddComplexIndicator(ctx context.Context, req *api.AddComplexIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddComplexIndicator(ctx, req)
}

func (p *kAlertClient) UpdateIndicator(ctx context.Context, req *api.UpdateIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateIndicator(ctx, req)
}

func (p *kAlertClient) DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteIndicator(ctx, req)
}

func (p *kAlertClient) SelectRuleById(ctx context.Context, req *api.SelectRuleIdRequest, callOptions ...callopt.Option) (r *api.SelectRuleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectRuleById(ctx, req)
}

func (p *kAlertClient) SelectRuleByRoomId(ctx context.Context, req *api.SelectRuleRoomIdRequest, callOptions ...callopt.Option) (r *api.SelectRuleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectRuleByRoomId(ctx, req)
}

func (p *kAlertClient) AddRule(ctx context.Context, req *api.AddRuleRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddRule(ctx, req)
}

func (p *kAlertClient) UpdateRule(ctx context.Context, req *api.UpdateRuleRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRule(ctx, req)
}

func (p *kAlertClient) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRule(ctx, req)
}

func (p *kAlertClient) SelectTaskById(ctx context.Context, req *api.SelectTaskIdRequest, callOptions ...callopt.Option) (r *api.SelectTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectTaskById(ctx, req)
}

func (p *kAlertClient) SelectTaskByRoomId(ctx context.Context, req *api.SelectTaskRoomIdRequest, callOptions ...callopt.Option) (r *api.SelectTasksResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectTaskByRoomId(ctx, req)
}

func (p *kAlertClient) AddTask(ctx context.Context, req *api.AddTaskRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTask(ctx, req)
}

func (p *kAlertClient) UpdateTask(ctx context.Context, req *api.UpdateTaskRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTask(ctx, req)
}

func (p *kAlertClient) DeleteTask(ctx context.Context, req *api.DeleteTaskRequest, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTask(ctx, req)
}
