## Task Manager API

Simple CRUD API built with Go, Gin, and MongoDB for persistent storage.

**Base URL:** `http://localhost:8080`

### Storage / MongoDB configuration

- **Driver**: Official MongoDB Go Driver (`go.mongodb.org/mongo-driver`)
- **Connection string**:
  - **Env var**: `MONGOURL`
  - **Default** (if `MONGOURL` is not set): `mongodb://localhost:27017`
- **Database**: `task_manager_db`
- **Collection**: `tasks`

Example (PowerShell on Windows):

```powershell
$env:MONGOURL = "mongodb://localhost:27017"
go run .
```

The API remains backward compatible: request and response formats are unchanged, only the storage layer now uses MongoDB instead of an in-memory slice.

### Task Model

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "status": false
}
```

### Endpoints

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






