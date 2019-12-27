package result

import (
	"fmt"
	"hot/utils"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}, w http.ResponseWriter) (n int, err error) {
	result := Result{0, "success", data}
	n, err = writeResponse(result, w)
	return
}

func Error(code int, message string, w http.ResponseWriter) (n int, err error) {
	result := Result{code, message, nil}
	n, err = writeResponse(result, w)
	return
}

func writeResponse(result Result, w http.ResponseWriter) (n int, err error) {
	jsonStr := utils.InterfaceToJsonString(result)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	n, err = fmt.Fprintf(w, "%s", jsonStr)
	return
}
