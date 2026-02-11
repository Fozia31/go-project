// main.go
package main

import (
    "fmt"
    "task_manager_api/router"
)

func main() {
    fmt.Println("Starting Task Manager API on port 8080...")
    r := router.SetupRouter()
    r.Run(":8080")
}
