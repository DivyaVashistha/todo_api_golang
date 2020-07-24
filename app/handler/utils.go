package handler

import (
	"encoding/json"
	"net/http"
)

// respondJSON generates the final response
func respondJSON(w http.ResponseWriter,status int,payload interface{}){
	response,err:=json.Marshal(payload)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError generates a response with custom message in case of error
func respondError(w http.ResponseWriter,code int,message string){
	respondJSON(w,code,map[string]string{"error":message})
}

