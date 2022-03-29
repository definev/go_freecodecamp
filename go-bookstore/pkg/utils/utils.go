package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(x any, r *http.Request) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err == nil {
			return
		}
	} else {
		panic(err)
	}
}

func StandardizedResponse(data any, errCode string, errMsg string) []byte {
	response := map[string]any{
		"data":    data,
		"errCode": errCode,
		"errMsg":  errMsg,
	}
	if decoded, err := json.Marshal(response); err != nil {
		panic(err)
	} else {
		return decoded
	}
}
