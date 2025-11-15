# Task Manager API

Simple CRUD API built with Go and Gin. In-memory storage.

**Base URL:** `http://localhost:8080`

## Task Model

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "status": false
}
```

## Endpoints

### POST /tasks
Create a new task.

**Request:**
```json
{
  "title": "task title",
  "description": "Write documentation",
  "dueDate": "2024-12-31",
  "status": false
}
```

**Response:** `201 Created`
```json
{
  "id": 1,
  "title": "task title",
  "description": "Write documentation",
  "dueDate": "2024-12-31",
  "status": false
}
```

---

### GET /tasks
Get all tasks.

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "title": "Task 1",
    "description": "Description",
    "dueDate": "2024-12-31",
    "status": false
  }
]
```

Returns empty array `[]` if no tasks exist.

---

### GET /tasks/:id
Get task by ID.

**Response:** `200 OK`
```json
{
  "id": 1,
  "title": "Task 1",
  "description": "Description",
  "dueDate": "2024-12-31",
  "status": false
}
```

---

### PUT /tasks/:id
Update a task.

**Request:**
```json
{
  "title": "Updated title",
  "description": "Updated description",
  "dueDate": "2025-01-15",
  "status": true
}
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "title": "Updated title",
  "description": "Updated description",
  "dueDate": "2025-01-15",
  "status": true
}
```

---

### DELETE /tasks/:id
Delete a task.

**Response:** `200 OK`
```json
{
  "message": "task deleted successfully"
}
```






