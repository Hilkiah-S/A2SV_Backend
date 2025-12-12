# Task Manager API

Task Management API built with Go, Gin, and MongoDB with JWT-based authentication and authorization.

**Base URL:** `http://localhost:8080`

## Storage / MongoDB Configuration

- **Driver**: Official MongoDB Go Driver (`go.mongodb.org/mongo-driver`)
- **Connection string**:
  - **Env var**: `MONGO_URL`
  - **Default** (if `MONGO_URL` is not set): `mongodb://localhost:27017`
- **Database**: `task_manager_db`
- **Collections**: `tasks`, `users`


## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Most endpoints require a valid JWT token in the Authorization header.

### Authentication Header Format

```
Authorization: Bearer <token>
```

### User Roles

- **admin**: Can create, update, delete tasks, and promote other users to admin
- **user**: Can retrieve all tasks and retrieve tasks by ID


## Authentication Endpoints

### POST /auth/register

Create a new user account.

**Request:**
```json
{
  "username": "john_doe",
  "password": "securepassword123"
}
```

**Response:** `201 Created`
```json
{
  "id": 1,
  "username": "john_doe",
  "role": "admin"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `409 Conflict`: Username already exists

---

### POST /auth/login

Authenticate user and receive JWT token.

**Request:**
```json
{
  "username": "john_doe",
  "password": "securepassword123"
}
```

**Response:** `200 OK`
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "username": "john_doe",
  "role": "admin"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Invalid credentials

**Usage:** Include the returned token in the Authorization header for subsequent requests:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

### POST /admin/promote

Promote a user to admin role. **Admin only.**

**Headers:**
```
Authorization: Bearer <admin_token>
```

**Request:**
```json
{
  "username": "jane_doe"
}
```

**Response:** `200 OK`
```json
{
  "message": "user promoted to admin successfully"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Missing or invalid token
- `403 Forbidden`: Admin access required
- `404 Not Found`: User not found

---

## Task Endpoints

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

---

### POST /tasks

Create a new task. **Admin only.**

**Headers:**
```
Authorization: Bearer <admin_token>
```

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

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Missing or invalid token
- `403 Forbidden`: Admin access required

---

### GET /tasks

Get all tasks. **Authenticated users only.**

**Headers:**
```
Authorization: Bearer <token>
```

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

**Error Responses:**
- `401 Unauthorized`: Missing or invalid token

---

### GET /tasks/:id

Get task by ID. **Authenticated users only.**

**Headers:**
```
Authorization: Bearer <token>
```

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

**Error Responses:**
- `400 Bad Request`: Invalid task ID
- `401 Unauthorized`: Missing or invalid token
- `404 Not Found`: Task not found

---

### PUT /tasks/:id

Update a task. **Admin only.**

**Headers:**
```
Authorization: Bearer <admin_token>
```

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

**Error Responses:**
- `400 Bad Request`: Invalid request body or task ID
- `401 Unauthorized`: Missing or invalid token
- `403 Forbidden`: Admin access required
- `404 Not Found`: Task not found

---

### DELETE /tasks/:id

Delete a task. **Admin only.**

**Headers:**
```
Authorization: Bearer <admin_token>
```

**Response:** `200 OK`
```json
{
  "message": "task deleted successfully"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid task ID
- `401 Unauthorized`: Missing or invalid token
- `403 Forbidden`: Admin access required
- `404 Not Found`: Task not found



