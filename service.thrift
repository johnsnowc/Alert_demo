namespace go api

struct SelectIndicatorRequest {
    1: string code
}

struct SelectIndicatorResponse {
     1: string code
     2: string name
     3: string left
     4: string right
     5: string op
     6: i64 timeRange
     7: string expr
}

struct AddSimpleIndicatorRequest {
    1: string code
    2: string name
    3: string expr
    4: i64 timeRange
}

struct AddComplexIndicatorRequest {
    1: string code
    2: string name
    3: string left
    4: string right
    5: string op
    6: i64 timeRange
}

struct UpdateIndicatorRequest {
    1: string code
    2: string name
    3: string left
    4: string right
    5: string op
    6: i64 timeRange
}

struct DeleteIndicatorRequest {
    1: string code
}

struct SelectRuleIdRequest {
    1: i64 id
}

struct SelectRuleRoomIdRequest {
    1: i64 roomId
}

struct SelectRuleResponse {
     1: i64 id
     2: i64 roomId
     3: string name
     4: string expr
}

struct AddRuleRequest{
    1: i64 roomId
    2: string name
    3: string expr
}

struct UpdateRuleRequest{
    1: i64 id
    2: string expr
}

struct DeleteRuleRequest{
    1: i64 id
}

struct SelectTaskIdRequest {
    1: i64 id
}

struct SelectTaskRoomIdRequest {
    1: i64 roomId
}

struct SelectTaskResponse {
     1: i64 id
     2: i64 roomId
     3: string name
     4: i64 ruleId
}

struct SelectTasksResponse {
     1: set<SelectTaskResponse> tasks
}


struct AddTaskRequest {
    1: string name
    2: i64 roomId
    3: i64 ruleId
    4: i64 frequency
}

struct UpdateTaskRequest {
     1: i64 id
     2: i64 roomId
     3: i64  ruleId
     4: i64  frequency
}

struct DeleteTaskRequest {
    1: i64  id
}

struct Response {
     1: string state
}

service Alert {
    SelectIndicatorResponse SelectIndicator(1: SelectIndicatorRequest req)
    Response AddSimpleIndicator(1: AddSimpleIndicatorRequest req)
    Response AddComplexIndicator(1: AddComplexIndicatorRequest req)
    Response UpdateIndicator(1: UpdateIndicatorRequest req)
    Response DeleteIndicator(1 : DeleteIndicatorRequest req)
    SelectRuleResponse SelectRuleById(1:SelectRuleIdRequest req)
    SelectRuleResponse SelectRuleByRoomId(1:SelectRuleRoomIdRequest req)
    Response AddRule(1:AddRuleRequest req)
    Response UpdateRule(1:UpdateRuleRequest req)
    Response DeleteRule(1:DeleteRuleRequest req)
    SelectTaskResponse SelectTaskById(1:SelectTaskIdRequest req)
    SelectTasksResponse SelectTaskByRoomId(1:SelectTaskRoomIdRequest req)
    Response AddTask(1:AddTaskRequest req)
    Response UpdateTask(1:UpdateTaskRequest req)
    Response DeleteTask(1:DeleteTaskRequest req)
}


