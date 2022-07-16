package main

import (
	"fmt"
	"mongo-golang/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getsession())
	r.GET("/users", uc.GetAllUsers)
	r.GET("/users/user/:id", uc.GetUser)
	r.POST("/users/user", uc.CreateUser)
	r.DELETE("/users/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)

}

func getsession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Not able to reach mongo instance")
		panic(err)

	}

	return s
}
