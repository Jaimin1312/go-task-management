# Task Management REST API

## Overview

This is a RESTful API for managing tasks, allowing users to create, read, update, and delete their tasks. Users can register, log in, and manage their tasks securely.

## Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.15 or later)
- [MongoDB](https://www.mongodb.com/try/download/community) (either locally or using MongoDB Atlas)
- A REST client (like [Postman](https://www.postman.com/) or [cURL](https://curl.se/))

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Jaimin1312/go-task-management
   cd go-task-management

2. **Install dependencies: Use the following command to download and install the necessary Go packages:**
    ```bash
    go mod tidy

3. **Set up environment variables: Ensure your MongoDB connection string and other environment variables are set. You can default.yml file**
    ```yml
    mongodatabase:
        host: "mongodb://localhost:27017/?retryWrites=true&w=majority&appName=task-management"
        DBName: "task-management"
    server:
        port: 8082
    ```

Run the application: Start the server with:
go run main.go

## API Endpoints

The following are the URLs where you can access the API and the Swagger documentation:

```
    Serving API at http://127.0.0.1:8082/task-service/
    Swagger API at http://127.0.0.1:8082/task-service/swagger/index.html#/
```

## API Endpoints

1. Authentication
2. User Registration
# API Endpoints

## Authentication

### User Registration
- **Endpoint:** `POST /register`
- **Request Body:**
    ```json
    {
        "email": "your_email",
        "password": "your_password"
    }
    ```
- **Response:** Confirmation of registration.

### User Login
- **Endpoint:** `POST /login`
- **Request Body:**
    ```json
    {
        "email": "your_email",
        "password": "your_password"
    }
    ```
- **Response:** JWT token for authentication.

## Task Management

### Create a Task
- **Endpoint:** `POST /tasks`
- **Request Body:**
    ```json
    {
        "title": "Task Title",
        "description": "Task Description",
        "status": "todo"
    }
    ```
- **Response:** Created task object.

### Get All Tasks
- **Endpoint:** `GET /tasks`
- **Response:** List of all tasks for the authenticated user.

### Get a Task by ID
- **Endpoint:** `GET /tasks/{id}`
- **Response:** Task object with the specified ID.

### Update a Task
- **Endpoint:** `PUT /tasks/{id}`
- **Request Body:**
    ```json
    {
        "title": "Updated Task Title",
        "description": "Updated Task Description",
        "status": "in progress"
    }
    ```
- **Response:** Updated task object.

### Delete a Task
- **Endpoint:** `DELETE /tasks/{id}`
- **Response:** Confirmation of task deletion.

### Mark Multiple Tasks as Done
- **Endpoint:** `PUT /tasks/mark-done`
- **Request Body:**
    ```json
    {
        "task_ids": ["60d5ec49c6d8c06e1f20c5a8", "60d5ec49c6d8c06e1f20c5a9"]
    }
    ```
- **Response:** Confirmation of tasks marked as done.

## Error Handling

The API provides meaningful error responses. Common errors include:

- **400 Bad Request:** Invalid request body.
- **401 Unauthorized:** Missing or invalid authentication token.
- **404 Not Found:** Requested resource not found.





# Database Choice: MongoDB

## Justification

### 1. Schema Flexibility
NoSQL document storage allows for a dynamic schema, accommodating evolving data models without complex migrations.

### 2. High Performance
Optimized for large data volumes and high write loads, enabling rapid data access for applications like task management.

### 3. Scalability
Supports horizontal scaling through sharding, ensuring efficient handling of growing data and user requests.

### 4. Rich Query Language
Powerful query capabilities, including complex aggregations, facilitate efficient data retrieval based on various criteria.

### 5. Document-Oriented Storage
JSON-like document structure provides a natural mapping to application data models, simplifying development.

### 6. Strong Community and Ecosystem
Large community support and robust tools enhance development, troubleshooting, and feature access.

### 7. Ease of Use
Intuitive query syntax and data model reduce boilerplate code and accelerate development cycles.
