package response

import (
	"encoding/json"
	"fmt"
)

func marshall(v interface{}) []byte {
	val, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error marshalling response")
		panic(err)
	}
	return val
}
