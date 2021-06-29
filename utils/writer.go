package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter,message string,error error, statusCode int)  {
	res:=ErrMsg{
		Message: message,
		Error: error,
	}
	v,err := json.Marshal(res)
	if err !=nil{
		fmt.Println("Error json Marshal")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	w.Write(v)
}
func WriteSuccess(w http.ResponseWriter,data interface{})  {
	res:=SuccessMsg{
		data,
	}
	v,err := json.Marshal(res)
	if err !=nil{
		fmt.Println("Error json Marshal")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(v)

}
