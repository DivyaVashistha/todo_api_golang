package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/api/app/model"
	"github.com/jinzhu/gorm"
)

// GetAllTopics handler to fetch details of all topics of a user
func GetAllTopics(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	
	vars:=mux.Vars(r) // get map of all variables from the request
	username:=vars["user_name"] // get value of username from the map
	user:=getUserOr404(db,username,w,r) // get details of user corresponding to username
	if user == nil{
		return
	}
	db.Preload("Topics").Preload("Details").Find(&user)
	respondJSON(w,http.StatusOK,user)
}

func getUserOr404(db *gorm.DB,username string,w http.ResponseWriter,r *http.Request) *model.User{
	user:=model.User{}
	if err:= db.First(&username,model.User{UserName:username}).Error; err!= nil{
		respondError(w, http.StatusNotFound,err.Error())
		return nil
	}
	return &user
}
