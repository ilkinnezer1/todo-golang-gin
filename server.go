package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Elements of Todo app
type todoElements struct {
	Id          int    `json:"Id"`
	Task        string `json:"Task"`
	isCompleted bool   `json:"Completed"`
}

var todoTasks = []todoElements{
	{Id: 1, Task: "Sleep", isCompleted: true},
	{Id: 2, Task: "Read", isCompleted: true},
	{Id: 3, Task: "Code", isCompleted: true},
}

// Handle Error
func handleError(err error) {
	if err != nil {
		fmt.Println("Oops, there is an error: ", err)
	}
}

// Get The tasks
func getTasks(context *gin.Context) {
	//Json conversion and return the tasks
	context.IndentedJSON(http.StatusOK, todoTasks)
}

func main() {
	// Run the server
	server := gin.Default()
	server.GET("/todo-tasks", getTasks)

	err := server.Run("localhost:5050")
	handleError(err)

}
