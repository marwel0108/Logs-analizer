package jsonlog

import (
	"encoding/json"
)

type Log struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Metadata  string `json:"metadata"`
}

func ParseJsonToLogStruct(jsonByteArray []byte) (log Log, err error) {
	var l Log
	err = json.Unmarshal(jsonByteArray, &l)
	if err != nil {
		return Log{}, err
	}

	return l, nil
}
