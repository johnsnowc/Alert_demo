package dto

type Result struct {
	RuleId    string  `json:"rule_id"`
	RuleName  string  `json:"rule_name"`
	RuleTree  *Rule   `json:"rule_tree"`
	Status    *Status `json:"status"`
	CheckTime int64   `json:"check_time"`
	Error     error   `json:"error"`
}
