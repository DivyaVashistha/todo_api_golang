package handler

import (
	"net/http"
	"github.com/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllDetails handler fetches details of a topic
func GetAllDetails(db *gorm.DB,w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	title:=vars["title"]
	topic:=getTopicOr404(db,title,w,r)
	if topic == nil{
		return
	}
	details:=[]model.Details{}
	db.Preload("Details").Find(&topic)
	respondJSON(w,http.StatusOK,details)
}

func getTopicOr404(db *gorm.DB,title string, w http.ResponseWriter, r *http.Request) *model.Topics{
	topic:=model.Topics{}
	if err:=db.First(&topic, model.Topics{Title:title}).Error;err!=nil{
		respondError(w,http.StatusNotFound,err.Error())
		return nil
	}
	return &topic
}