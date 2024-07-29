package util

import "encoding/json"

func MarshalJson(val interface{}) (string, error) {
	str, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func UnmarshalJson(buf string, val interface{}) error {
	return json.Unmarshal([]byte(buf), val)
}

func MustMarshalJson(val interface{}) string {
	str, err := MarshalJson(val)
	if err != nil {
		panic(err)
	}
	return str
}
