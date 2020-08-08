package app

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/api/app/handler"
	"github.com/api/app/model"
	"github.com/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App struct defines router and database
type App struct{
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) Initialize(config *config.Config){
	dbURI:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
config.DB.Username,
config.DB.Password,
config.DB.Host,
config.DB.Port,
config.DB.Name,
config.DB.Charset)
db,err:=gorm.Open(config.DB.Dialect,dbURI)
if err!= nil{
	log.Fatal("could not connect to database")
}
a.DB=model.DBMigrate(db)
a.Router=mux.NewRouter()
a.setRouters()
}

// setRouters sets all the required routes
func (a *App) setRouters(){
	a.Get("/topics",a.handleRequest(handler.GetAllUsers))
	a.Post("/users", a.handleRequest(handler.CreateUser))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)){
	a.Router.HandleFunc(path,f).Methods("GET")
}

func (a *App) Post(path string,f func(w http.ResponseWriter,r *http.Request)){
	a.Router.HandleFunc(path,f).Methods("POST")
}

type RequestHandlerFunction func(db *gorm.DB,w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		handler(a.DB,w,r)
	}
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
