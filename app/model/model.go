package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User struct for defining user fields
type User struct{
	gorm.Model
	UserName string `gorm:"unique" json:"user_name"`
	Password string `gorm:"not null" json:"password"`
	Topics []Topics `gorm:"ForeignKey:UserID" json:"topics"`
}
// Topics struct for defining topic fields
type Topics struct{
	gorm.Model
	UserID uint `json:"user_id"`
	Title string `gorm:"unique;not null" json:"title"`
	Details []Details `gorm:"ForeignKey:TopicID" json:"details"`
}
// Details struct for defining details of a topic
type Details struct{
	gorm.Model
	TopicID uint `json:"topic_id"`
	Notes string `gorm:"not null" json:"notes"`
	Links string `json:"links"`
	IsComplete bool `json:"is_complete"`
}

// CompleteTask function is used to mark the task as completed
func (d *Details) CompleteTask(){
	d.IsComplete=true
}

// WorkInProgressTask function marks the task as in progress
func (d *Details) WorkInProgressTask(){
	d.IsComplete=false
}

func DBMigrate(db *gorm.DB) *gorm.DB{
	db.AutoMigrate(&User{},&Topics{},&Details{})
	db.Model(&Topics{}).AddForeignKey("user_id","users(id)","CASCADE","CASCADE")
	db.Model(&Details{}).AddForeignKey("topic_id","topics(id)","CASCADE","CASCADE")	
	return db
}
