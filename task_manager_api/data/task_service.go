// data/task_service.go
package data

import (
    "time"
    "task_manager_api/models"
    "github.com/google/uuid"
)

var Tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []models.Task {
    return Tasks
}

func GetTaskByID(id string) (*models.Task, bool) {
    for _, task := range Tasks {
        if task.ID == id {
            return &task, true
        }
    }
    return nil, false
}

func CreateTask(task models.Task) models.Task {
    task.ID = uuid.New().String()
    Tasks = append(Tasks, task)
    return task
}

func UpdateTask(id string, updated models.Task) (*models.Task, bool) {
    for i, task := range Tasks {
        if task.ID == id {
            if updated.Title != "" {
                Tasks[i].Title = updated.Title
            }
            if updated.Description != "" {
                Tasks[i].Description = updated.Description
            }
            if !updated.DueDate.IsZero() {
                Tasks[i].DueDate = updated.DueDate
            }
            if updated.Status != "" {
                Tasks[i].Status = updated.Status
            }
            return &Tasks[i], true
        }
    }
    return nil, false
}

func DeleteTask(id string) bool {
    for i, task := range Tasks {
        if task.ID == id {
            Tasks = append(Tasks[:i], Tasks[i+1:]...)
            return true
        }
    }
    return false
}
