package entity

type Rank struct {
	Problems     []*ContestProblem `json:"problems"`
	Participants []*Participant    `json:"participants"`
}

type Participant struct {
	Username string                    `json:"username"`
	Rank     uint                      `json:"rank"`
	Score    uint                      `json:"score"`
	Penalty  uint                      `json:"penalty"`
	Problems map[string]*ProblemStatus `json:"problems"`
}

type ProblemStatus struct {
	ProblemId   uint   `json:"problem_id"`
	SubmitCount uint   `json:"submit_count"`
	SubmitTime  string `json:"submit_time"`
	Status      uint   `json:"status"`
}

type Participants []*Participant

func (p Participants) Len() int { return len(p) }

func (p Participants) Less(i, j int) bool {
	if p[i].Score == p[j].Score {
		return p[i].Penalty < p[j].Penalty
	}
	return p[i].Score > p[j].Score
}

func (p Participants) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
