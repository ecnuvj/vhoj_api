package adapter

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	ret := RpcProblemsToEntityProblems(nil)
	fmt.Println(ret)
}
