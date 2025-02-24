package response

import (
	"encoding/json"
	"go.uber.org/zap"
)

func marshall(v interface{}) []byte {
	val, err := json.Marshal(v)
	if err != nil {
		zap.L().Error("Error marshalling response", zap.Error(err))
		panic(err)
	}
	return val
}
