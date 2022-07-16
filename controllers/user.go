package controllers

import (
	"encoding/json"
	"fmt"
	"mongo-golang/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{session: s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)

	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	_, err := json.Marshal(u) 

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%s\n", uj)
	json.NewEncoder(w).Encode(u)

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.ID = bson.NewObjectId()
	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)

	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)

	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", oid)

}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	coll := uc.session.DB("mongo-golang").C("users")

	var arr []models.User
	coll.Find(nil).All(&arr)

	json.NewEncoder(w).Encode(arr)

}
