package main

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{Id: "1", Item: "Clean Room", Completed: false},
	{Id: "2", Item: "Read Book", Completed: false},
	{Id: "3", Item: "Go for walk", Completed: false},
}

func getAllTodos(res http.ResponseWriter, req *http.Request) {
	response, err := json.Marshal(Todos)
	if err != nil {
		http.Error(res, "Found Error", http.StatusBadGateway)
	}
	res.Write(response)
}

func addTodo(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error Reading request body", http.StatusBadRequest)
	}
	var td Todo

	json.Unmarshal(body, &td)
	// fmt.Println(td)
	Todos = append(Todos, td)
	return

}

func getTodo(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	for _, value := range Todos {
		if value.Id == id {
			response, err := json.Marshal(value)
			if err != nil {
				http.Error(res, "Marshal Error", http.StatusBadRequest)
			}
			res.WriteHeader(http.StatusAccepted)
			res.Write(response)
			return
		}
	}
	http.Error(res, "Resource not found", 404)

}

func toggleTodo(res http.ResponseWriter , req *http.Request) {
      id := req.PathValue("id") 
	  for i:= range Todos {
		  if Todos[i].Id == id {
			  Todos[i].Completed = !Todos[i].Completed
			  return
		  }
	  }
	  http.Error(res , "Resource Not Found" , http.StatusBadRequest) 
}

func main() {

	// Lets create the routes.
	http.HandleFunc("GET /todos", getAllTodos)
	http.HandleFunc("GET /todos/{id}", getTodo)
	http.HandleFunc("POST /todos", addTodo)
	http.HandleFunc("PATCH /todos/{id}" , toggleTodo)
	fmt.Println("server started at 8080 port")
	// Lets start a server using http.ListenAndServe(port , mux)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
