Here’s a sample `README.md` for a dummy REST API project in Go:

---

# Dummy REST API in Go

This is a simple REST API built with Go to demonstrate basic HTTP server functionality. It handles operations for managing `users` with endpoints for creating and retrieving users.

---

## Features

- **GET /users**: Retrieve all users.
- **POST /users**: Add a new user.
- Simple in-memory storage for user data.
- Error handling for invalid inputs.

---

## Prerequisites

- **Go (Golang)**: Version 1.17 or higher.
- Optional: Tools like `curl` or Postman for testing the API.

---

## Getting Started

### Clone the Repository
```bash
git clone https://github.com/your-repo-name/dummy-rest-go.git
cd dummy-rest-go
```

### Install Dependencies
This project uses Go modules. Run the following command to ensure all dependencies are installed:
```bash
go mod tidy
```

### Run the Server
Start the server by running:
```bash
go run main.go
```

The server will start on `http://localhost:8085`.

---

## API Endpoints

### **1. GET /users**

Retrieve a list of all users.

#### Request:
```bash
curl -X GET http://localhost:8085/users
```

#### Response:
- **Status 200**: Returns a JSON array of users.
```json
[
  {
    "email": "john.doe@example.com",
    "first_name": "John",
    "last_name": "Doe"
  }
]
```

---

### **2. POST /users**

Add a new user.

#### Request:
```bash
curl -X POST http://localhost:8085/users \
-H "Content-Type: application/json" \
-d '{
  "email": "jane.doe@example.com",
  "first_name": "Jane",
  "last_name": "Doe"
}'
```

#### Response:
- **Status 201**: User successfully created.
- **Status 400**: Invalid input (e.g., missing fields).

---

## Project Structure

```
dummy-rest-go/
├── main.go                # Entry point of the application
├── api/
│   ├── user_handler.go    # Handlers for user-related endpoints
│   ├── user.go            # User data structure and business logic
│   └── storage.go         # In-memory storage for users
└── go.mod                 # Go module dependencies
```

---

## Future Improvements

- Add database integration (e.g., PostgreSQL).
- Implement middleware for authentication and logging.
- Expand API with additional endpoints (e.g., DELETE, UPDATE).

---

## License

This project is open-source and licensed under the [MIT License](LICENSE).

---
