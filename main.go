package main

import (
	"github.com/api/app"
	"github.com/api/config"
	_ "github.com/api/docs"
)
func main(){
	config:= config.GetConfig()
	app:=&app.App{}
	app.Initialize(config)
	app.Run(":3000")
}