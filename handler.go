package main

import (
	handler "Alert_demo/core/biz/handlers"
	"Alert_demo/kitex_gen/api"
	"context"
)

// AlertImpl implements the last service interface defined in the IDL.
type AlertImpl struct{}

// SelectIndicator implements the AlertImpl interface.
func (s *AlertImpl) SelectIndicator(ctx context.Context, req *api.SelectIndicatorByCodeRequest) (resp *api.Response, err error) {
	resp, err = handler.SelectIndicator(ctx, req)
	return
}

// AddSimpleIndicator implements the AlertImpl interface.
func (s *AlertImpl) AddSimpleIndicator(ctx context.Context, req *api.AddSimpleIndicatorRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.AddSimpleIndicator(ctx, req)
	return
}

// AddComplexIndicator implements the AlertImpl interface.
func (s *AlertImpl) AddComplexIndicator(ctx context.Context, req *api.AddComplexIndicatorRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.AddComplexIndicator(ctx, req)
	return
}

// UpdateIndicator implements the AlertImpl interface.
func (s *AlertImpl) UpdateIndicator(ctx context.Context, req *api.UpdateIndicatorRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.UpdateIndicator(ctx, req)
	return
}

// DeleteIndicator implements the AlertImpl interface.
func (s *AlertImpl) DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.DeleteIndicator(ctx, req)
	return
}

// SelectRuleByCode implements the AlertImpl interface.
func (s *AlertImpl) SelectRuleByCode(ctx context.Context, req *api.SelectRuleByCodeRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.SelectRuleByCode(ctx, req)
	return
}

// SelectRuleById implements the AlertImpl interface.
func (s *AlertImpl) SelectRuleById(ctx context.Context, req *api.SelectRuleByIdRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.SelectRuleById(ctx, req)
	return
}

// SelectRuleByRoomId implements the AlertImpl interface.
func (s *AlertImpl) SelectRuleByRoomId(ctx context.Context, req *api.SelectRuleByRoomIdRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.SelectRuleByRoomId(ctx, req)
	return
}

// AddRule implements the AlertImpl interface.
func (s *AlertImpl) AddRule(ctx context.Context, req *api.AddRuleRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.AddRule(ctx, req)
	return
}

// UpdateRule implements the AlertImpl interface.
func (s *AlertImpl) UpdateRule(ctx context.Context, req *api.UpdateRuleRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.UpdateRule(ctx, req)
	return
}

// DeleteRule implements the AlertImpl interface.
func (s *AlertImpl) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.DeleteRule(ctx, req)
	return
}

// SelectTaskById implements the AlertImpl interface.
func (s *AlertImpl) SelectTaskById(ctx context.Context, req *api.SelectTaskByIdRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.SelectTaskById(ctx, req)
	return
}

// SelectTaskByRoomId implements the AlertImpl interface.
func (s *AlertImpl) SelectTaskByRoomId(ctx context.Context, req *api.SelectTaskByRoomIdRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.SelectTaskByRoomId(ctx, req)
	return
}

// AddTask implements the AlertImpl interface.
func (s *AlertImpl) AddTask(ctx context.Context, req *api.AddTaskRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.AddTask(ctx, req)
	return
}

// UpdateTask implements the AlertImpl interface.
func (s *AlertImpl) UpdateTask(ctx context.Context, req *api.UpdateTaskRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.UpdateTask(ctx, req)
	return
}

// DeleteTask implements the AlertImpl interface.
func (s *AlertImpl) DeleteTask(ctx context.Context, req *api.DeleteTaskRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.DeleteTask(ctx, req)
	return
}

// Work implements the AlertImpl interface.
func (s *AlertImpl) Work(ctx context.Context, req *api.WorkRequest) (resp *api.Response, err error) {
	// TODO: Your code here...
	resp, err = handler.Work(ctx, req)
	return
}
