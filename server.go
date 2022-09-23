package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func addTasks(context *gin.Context) {
	var newTask todoElements
	//Binding json data
	if err := context.BindJSON(&newTask); err != nil {
		return
	}
	// Adding tasks to the slice
	for _, task := range todoTasks {
		taskId, err := strconv.Atoi(context.Param("Id"))
		handleError(err)
		if taskId == task.Id {
			context.IndentedJSON(http.StatusNotFound, "Tasks already exists, please change the ID of tasks")
		} else {
			todoTasks = append(todoTasks, newTask)
			context.IndentedJSON(http.StatusCreated, todoTasks)
		}
	}
}

func main() {
	// Run the server
	server := gin.Default()
	server.GET("/tasks", getTasks)
	server.POST("/add-task", addTasks)

	err := server.Run("localhost:5050")
	handleError(err)

}
