// controllers/task.controller.go
package controllers

import (
    "net/http"
    "task_manager_api/data"
    "task_manager_api/models"

    "github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    c.JSON(http.StatusOK, data.GetAllTasks())
}

func GetTask(c *gin.Context) {
    id := c.Param("id")
    task, found := data.GetTaskByID(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    var newTask models.Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }
    created := data.CreateTask(newTask)
    c.JSON(http.StatusCreated, created)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }
    task, ok := data.UpdateTask(id, updatedTask)
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    if !data.DeleteTask(id) {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
