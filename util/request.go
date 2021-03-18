package util

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func CheckRequest(request interface{}) {
	str, _ := json.Marshal(request)
	fmt.Println(string(str))
}

func StringToNumber(num string) (int, error) {
	if num == "" {
		return 0, nil
	}
	return strconv.Atoi(num)
}
