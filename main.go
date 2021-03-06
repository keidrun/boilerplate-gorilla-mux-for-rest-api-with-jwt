package main // import "github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/config"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/controllers"
	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/driver"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

func init() {
	gotenv.Load()
}

func main() {
	conf := config.GetConfig()

	db := driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/api/users", controller.TokenVerifyMiddleWare(controller.GetUsers(db))).Methods("GET")
	router.HandleFunc("/api/users/{id}", controller.TokenVerifyMiddleWare(controller.GetUser(db))).Methods("GET")
	router.HandleFunc("/api/users", controller.TokenVerifyMiddleWare(controller.AddUser(db))).Methods("POST")
	router.HandleFunc("/api/users/{id}", controller.TokenVerifyMiddleWare(controller.UpdateUser(db))).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controller.TokenVerifyMiddleWare(controller.RemoveUser(db))).Methods("DELETE")

	log.Printf("Server up on port %v....\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Port), router))
}
