package handler

import(
	"encoding/json"
	"net/http"
	"github.com/api/app/model"
	"github.com/jinzhu/gorm"
)

// GetAllUsers handler fetches details of all users in the system
func GetAllUsers(db *gorm.DB,w http.ResponseWriter,r *http.Request){
	users:=[]model.User{}
	db.Preload("Topics").Preload("Topics.Details").Find(&users)
	respondJSON(w,http.StatusOK,users)
}

// CreateUser handler add details of new user in the system
func CreateUser(db *gorm.DB,w http.ResponseWriter,r * http.Request){
	user:=model.User{}
	decoder:=json.NewDecoder(r.Body)
	if err:= decoder.Decode(&user); err!= nil{
		respondError(w,http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()
	if err:=db.Save(&user).Error; err!=nil{
		respondError(w,http.StatusInternalServerError,err.Error())
		return
	}
	respondJSON(w,http.StatusCreated,user)
}