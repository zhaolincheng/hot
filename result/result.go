package result

import (
	"fmt"
	"hot/utils"
	"net/http"
)

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

func Success(data interface{}, w http.ResponseWriter) {
	result := Result{0, "success", data}
	writeResponse(result, w)
}

func Error(code int, message string, w http.ResponseWriter) {
	result := Result{code, message, nil}
	writeResponse(result, w)
}

func writeResponse(result Result, w http.ResponseWriter) {
	jsonStr := utils.InterfaceToJsonString(result)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	_, err := fmt.Fprintf(w, "%s", jsonStr)
	if err != nil {
		utils.Error.Println(err)
	}
}
