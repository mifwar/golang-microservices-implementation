package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)

	LogError(err, "parse body")

	_ = json.Unmarshal([]byte(body), x)
}

func LogError(err error, sentence string) {
	if err != nil {
		log.Fatalf("%s error at : %v\n", sentence, err)
	}
}

func SendData(w http.ResponseWriter, x interface{}) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(x)
}

func GetParamID(r *http.Request, param string) int {
	params := mux.Vars(r)
	paramID := params[param]
	id, err := strconv.Atoi(paramID)

	LogError(err, "str conv")

	return id
}
