namespace go api

struct SelectIndicatorByCodeRequest {
    1: string code
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
    6: string expr
    7: i64 timeRange
}

struct DeleteIndicatorRequest {
    1: string code
}

struct SelectRuleByIdRequest {
    1: i64 id
}

struct SelectRuleByCodeRequest {
    1: string code
}

struct SelectRuleByRoomIdRequest {
    1: i64 roomId
}

struct AddRuleRequest{
    1: string code
    2: string name
    3: string expr
    4: i64 roomId
}

struct UpdateRuleRequest{
    1: i64 id
    2: string expr
}

struct DeleteRuleRequest{
    1: i64 id
}

struct SelectTaskByIdRequest {
    1: i64 id
}

struct SelectTaskByRoomIdRequest {
    1: i64 roomId
}

struct AddTaskRequest {
    1: string name
    2: i64 roomId
    3: string ruleCode
    4: i64 frequency
}

struct UpdateTaskRequest {
     1: i64 id
     2: i64 roomId
     3: string ruleCode
     4: i64 frequency
}

struct DeleteTaskRequest {
    1: i64  id
}

struct WorkRequest {
    1: i64 timeRange
}

struct Response {
     1: i16 code
     2: string message
     3: string data
}

service Alert {
    Response SelectIndicator(1: SelectIndicatorByCodeRequest req)
    Response AddSimpleIndicator(1: AddSimpleIndicatorRequest req)
    Response AddComplexIndicator(1: AddComplexIndicatorRequest req)
    Response UpdateIndicator(1: UpdateIndicatorRequest req)
    Response DeleteIndicator(1 : DeleteIndicatorRequest req)
    Response SelectRuleById(1:SelectRuleByIdRequest req)
    Response SelectRuleByCode(1:SelectRuleByCodeRequest req)
    Response SelectRuleByRoomId(1:SelectRuleByRoomIdRequest req)
    Response AddRule(1:AddRuleRequest req)
    Response UpdateRule(1:UpdateRuleRequest req)
    Response DeleteRule(1:DeleteRuleRequest req)
    Response SelectTaskById(1:SelectTaskByIdRequest req)
    Response SelectTaskByRoomId(1:SelectTaskByRoomIdRequest req)
    Response AddTask(1:AddTaskRequest req)
    Response UpdateTask(1:UpdateTaskRequest req)
    Response DeleteTask(1:DeleteTaskRequest req)
    Response Work(1: WorkRequest req)
}