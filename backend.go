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

func NewTodoController(s *mgo.Session) *TodoController {
	return &TodoController{s}
}

func (tc TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get parameter "id"
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	te := Todo{}
	json.NewDecoder(r.Body).Decode(&te)
	if err := tc.session.DB("TodoList").C("Todos").UpdateId(oid, te); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tej, _ := json.Marshal(te)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tej)
}

func (tc TodoController) GetTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get parameter "id"
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	te := Todo{}

	if err := tc.session.DB("TodoList").C("Todos").FindId(oid).One(&te); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tej, _ := json.Marshal(te)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tej)
}

func (tc TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get parameter "id"
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := tc.session.DB("TodoList").C("Todos").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
}

func (tc TodoController) CreateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Create empty Todo Object
	te := Todo{}
	json.NewDecoder(r.Body).Decode(&te)
	new_te := NewTodo(te.Name)
	tc.session.DB("TodoList").C("Todos").Insert(new_te)
	tej, _ := json.Marshal(new_te)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(201)
	// w.WriteHeader(500)
	fmt.Fprintf(w, "%s", tej)
}

func (tc TodoController) GetAllTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Get All Todos")

	// stub array of Todo
	todos := []Todo{}

	// Query MongoDB for whole TodoList Collection
	if err := tc.session.DB("TodoList").C("Todos").Find(nil).All(&todos); err != nil {
		// If not data respond with no content and exit
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Marshal array of todo to JSON
	todosj, _ := json.Marshal(todos)

	// set header to JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// set status header to 200
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", todosj)
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
	// t1 := NewTodo("Deo")
	// t2 := NewTodo("Mini Hand Creme")
	tc := NewTodoController(getSession())
	// tc.session.DB("TodoList").C("Todos").DropCollection()
	// tc.session.DB("TodoList").C("Todos").Insert(t1, t2)
	// r.ServeFiles("/web/*filepath", http.Dir("/home/oliver/coding/gocode/src/todo/web"))
	r.ServeFiles("/web/*filepath", http.Dir("/home/oliver/coding/gocode/src/todo/www"))
	r.GET("/todo/:id", tc.GetTodo)
	r.GET("/todo", tc.GetAllTodo)
	r.POST("/todo", tc.CreateTodo)
	r.DELETE("/todo/:id", tc.DeleteTodo)
	r.PUT("/todo/:id", tc.UpdateTodo)
	// r.PUT("/todo/:id", tc.Delete2Todo)
	log.Println("ListenAndServe localhost:8080")
	http.ListenAndServe(":8080", r)
	// Do this below for SSL Encrypted Backend
	// http.ListenAndServeTLS("localhost:8433", "server.pem", "server.key", r)
}
