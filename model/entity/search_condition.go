package entity

type SearchCondition struct {
	Title     string `json:"title" form:"title"`
	ProblemId uint   `json:"problem_id" form:"problem_id"`
}
