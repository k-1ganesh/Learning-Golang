package main

import (
	// "fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func getTodos(context *gin.Context) {
	// Now client server communication happens through json .
	// We need to convert above data to json.
	// This can be done using context.IndentedJSON(status , object_to_convert)
	context.IndentedJSON(http.StatusOK, Todos)

}
func getTodosById(id string) (*Todo, error) {
	for i := range Todos {
		if Todos[i].Id == id {
			return &Todos[i], nil
		}
	}
	return nil, errors.New("value is not in database")
}
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)

}

func toggleTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo) // This returns the result to client
}
func addTodo(context *gin.Context) {
	// Now client will send data in the form of json.
	// server in this case dosen't understand json as it understand todo
	// So we need to convert that data to type todo.
	// Data will be in the post request. We need to extract that data.
	// This conversion can be done using BindJSON(&variable_name)
	var newTodo Todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		log.Fatal("Conversion Failed")
		return
	}
	Todos = append(Todos, newTodo)
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)    // To get the data from the server
	router.GET("/todos/:id", getTodo) // To get the data from the server
	router.PATCH("/todos/:id", toggleTodo)
	router.POST("/todos", addTodo) // To send the data to the server
	router.Run("localhost:8080")
}
