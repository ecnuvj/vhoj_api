package adapter

import (
	"fmt"
	"testing"
	"time"
)

func TestNil(t *testing.T) {
	ret := RpcProblemsToEntityProblems(nil)
	fmt.Println(ret)
}

func TestTime(t *testing.T) {
	tm := time.Unix(1614960000, 0)
	fmt.Println(tm)
}
