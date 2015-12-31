// Package main provides a RESTful Todo List Backend
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// My ToDo Entry structure
type Todo struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Created int64         `json:"created" bson:"created"`
	Name    string        `json:"name" bson:"name"`
	Done    bool          `json:"done" bson:"done"`
}

// Creates a new todo entry ready for db insertion
func NewTodo(name string) Todo {
	t := Todo{bson.NewObjectId(), time.Now().UnixNano(), name, false}
	return t
}

type TodoController struct {
	session *mgo.Session
}

func (tc TodoController) GetTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("GetTodo")
	// get parameter "id"
	id := p.ByName("id")
	log.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	log.Println(oid)
	te := Todo{}
	log.Println(te)

	if err := tc.session.DB("TodoList").C("Todos").Find(oid).One(&te); err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(oid, " not found")
		return
	}
	log.Println(te)
	tej, _ := json.Marshal(te)
	log.Println("tej")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tej)
}

func (tc TodoController) PostTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (tc TodoController) GetAllTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Get All Todos")

	// stub array of Todo
	todos := []Todo{}

	// Query MongoDB for whole TodoList Collection
	if err := tc.session.DB("TodoList").C("Todos").Find(nil).All(&todos); err != nil {
		// If not data respond with no content and exit
		w.WriteHeader(http.StatusNoContent)
		log.Println("No Content")
		return
	}

	// Marshal array of todo to JSON
	todosj, _ := json.Marshal(todos)

	// set header to JSON
	w.Header().Set("Content-Type", "application/json")
	// set status header to 200
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", todosj)
}

func NewTodoController(s *mgo.Session) *TodoController {
	return &TodoController{s}
}

// Get a session with the MongoDB
func getSession() *mgo.Session {
	s, err := mgo.Dial("omv.fritz.box")
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := httprouter.New()
	t1 := NewTodo("Gasflasche")
	t2 := NewTodo("Bierkasten")
	tc := NewTodoController(getSession())
	tc.session.DB("TodoList").C("Todos").Insert(t1, t2)
	r.GET("/todo/:id", tc.GetTodo)
	r.GET("/todo", tc.GetAllTodo)
	// r.POST("/todo", tc.PostTodo)
	// r.DELETE("/todo/:id", tc.DeleteTodo)

	// 	r.GET("/todo/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	t := Todo{
	// 		Created: time.Now().UnixNano(),
	// 		Name:    "Kuchen",
	// 		Done:    false,
	// 	}
	// 	tj, _ := json.Marshal(t)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(200)
	// 	fmt.Fprintf(w, "%s", tj)
	// })
	// r.POST("/todo", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 	t := Todo{}
	// 	json.NewDecoder(r.Body).Decode(&t)
	// 	// Marshal provided interface into JSON structure
	// 	tj, _ := json.Marshal(t)
	// 	// Write content-type, statuscode, payload
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(201)
	// 	fmt.Fprintf(w, "%s", tj)
	// })
	// r.DELETE("/todo/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	w.WriteHeader(201)
	// })

	log.Println("ListenAndServe localhost:8080")
	http.ListenAndServe("localhost:8080", r)

}
