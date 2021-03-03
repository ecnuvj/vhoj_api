package entity

type Problem struct {
	ProblemId    uint   `json:"problem_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	SampleInput  string `json:"sample_input"`
	SampleOutput string `json:"sample_output"`
	Input        string `json:"input"`
	Output       string `json:"output"`
	TimeLimit    string `json:"time_limit"`
	MemoryLimit  string `json:"memory_limit"`
	Submitted    int64  `json:"submitted"`
	Accepted     int64  `json:"accepted"`
}
