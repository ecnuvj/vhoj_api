package entity

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestSubmissions_Sort(t *testing.T) {
	submissions := []*Submission{
		{
			SubmissionId: 1,
			SubmitTime:   time.Now().Add(3 * time.Minute),
		},
		{
			SubmissionId: 2,
			SubmitTime:   time.Now().Add(5 * time.Minute),
		},
		{
			SubmissionId: 3,
			SubmitTime:   time.Now(),
		},
	}
	sort.Sort(Submissions(submissions))
	str, _ := json.Marshal(submissions)
	fmt.Println(string(str))
}

func TestSubmissions_Len(t *testing.T) {
	userMap := make(map[uint][]*Submission)
	userMap[1] = append(userMap[1], &Submission{SubmissionId: 1})
	fmt.Println(userMap[1])

	userAcMap := make(map[uint]map[string]interface{})
	//userAcMap[1]["A"] = make(map[string]interface{})
	_, e := userAcMap[1]
	if !e {
		userAcMap[1] = make(map[string]interface{})
	}
	userAcMap[1]["a"] = 0
	userAcMap[1]["a"] = userAcMap[1]["a"].(int) + 1
	fmt.Println(userAcMap[1])
	//
	//tt := make(map[string]interface{})
	////tt["A"] = Submission{}
	//fmt.Println(tt["A"])
}

func TestParticipants_Len(t *testing.T) {
	t1 := time.Now()
	t2 := time.Now().Add(2*time.Hour + 30*time.Minute)
	var t3 = uint(t2.Sub(t1).Seconds())
	fmt.Println(fmt.Sprintf("%v:%v", t3/3600, (t3%3600)/60))
}
