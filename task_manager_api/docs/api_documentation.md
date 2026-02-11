Base URL: http://localhost:8080

1️⃣ Get All Tasks
Endpoint: GET /tasks
curl -X GET http://localhost:8080/tasks

Response:
Status: 200 OK
Body (JSON):

[
  {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2026-02-11T00:00:00Z",
    "status": "Pending"
  },
  {
    "id": "2",
    "title": "Task 2",
    "description": "Second task",
    "due_date": "2026-02-12T00:00:00Z",
    "status": "In Progress"
  }
]

2️⃣ Get Task by ID
Endpoint: GET /tasks/:id

Curl Command:
curl -X GET http://localhost:8080/tasks/1

Expected Response (200 OK):
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2026-02-11T00:00:00Z",
  "status": "Pending"
}

If Task Not Found (404):
{
  "message": "Task not found"
}

3️⃣ Create a New Task
Endpoint: POST /tasks

Curl Command:
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "title": "New Task",
  "description": "Description of the task",
  "due_date": "2026-02-15T00:00:00Z",
  "status": "Pending"
}'
Expected Response (201 Created):
{
  "id": "3",
  "title": "New Task",
  "description": "Description of the task",
  "due_date": "2026-02-15T00:00:00Z",
  "status": "Pending"
}

4️⃣ Update a Task
Endpoint: PUT /tasks/:id

Curl Command:
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Task Title",
  "status": "Completed"
}'
Expected Response (200 OK):
{
  "id": "1",
  "title": "Updated Task Title",
  "description": "First task",
  "due_date": "2026-02-11T00:00:00Z",
  "status": "Completed"
}
If Task Not Found (404):
{
  "message": "Task not found"
}

5️⃣ Delete a Task
Endpoint: DELETE /tasks/:id

Curl Command:
curl -X DELETE http://localhost:8080/tasks/1
Expected Response (200 OK):
{
  "message": "Task deleted successfully"
}
If Task Not Found (404):
{
  "message": "Task not found"
}