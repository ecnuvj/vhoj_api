package util

import (
	"encoding/json"
	"fmt"
)

func CheckRequest(request interface{}) {
	str, _ := json.Marshal(request)
	fmt.Println(string(str))
}
