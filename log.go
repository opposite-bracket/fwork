package fwork

import (
	"encoding/json"
	"fmt"
	"time"
)

type customLog struct {
	Timestamp time.Time   `json:"ts"`
	Data      interface{} `json:"json"`
}

func JsonInfoLog(data interface{}) {
	encoded, _ := json.Marshal(customLog{
		Timestamp: time.Now(),
		Data: data,
	})
	fmt.Println(string(encoded))
}
