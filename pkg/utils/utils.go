package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), x)
}