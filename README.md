#  User Management API
## Introduction

This documentation provides an overview and guide to the User Management CRUD API built using Go Fiber, GORM, JWT, and MySQL. The API allows for creating, reading, updating, and deleting user data, with JWT-based authentication to secure endpoints.

## Prerequisites

-   Go (version 1.16 or higher)
-   MySQL (version 8.0 or higher)
-   Go Fiber
-   GORM
-   JWT

## Project Setup

### Clone the Repository

```
git clone <repository-url>
cd <repository-folder>
```
### Install Dependencies


```
go mod tidy
```

### Configuration

Create a `.env` file in the project root and add the following configuration variables:

makefile
```
DB_USER=<your-db-username>
DB_PASSWORD=<your-db-password>
DB_NAME=<your-db-name>
DB_HOST=<your-db-host>
DB_PORT=<your-db-port>
JWT_SECRET=<your-jwt-secret>
```

### Database Migration

Run the following command to migrate the database:

```
go run main.go migrate
```

### Running the Application


```
go run main.go
```

## API Endpoints

### Authentication

#### Register

-   **Endpoint:** `/api/register`
-   **Method:** `POST`
-   **Description:** Register a new user.
-   **Request Body:**
    
    json
    
```
  {
      "username": "string",
      "password": "string",
      "email": "string"
}
``` 
    
-   **Response:**
    
    json
    

-   `{
      "message": "User registered successfully"
    }` 
    

#### Login

-   **Endpoint:** `/api/login`
-   **Method:** `POST`
-   **Description:** Authenticate a user and generate a JWT token.
-   **Request Body:**
    
    json
    

-   `{
      "username": "string",
      "password": "string"
    }` 
    
-   **Response:**
    
    json
    

-   `{
      "token": "jwt-token"
    }` 
    

### User Management

#### Create User

-   **Endpoint:** `/api/users`
-   **Method:** `POST`
-   **Description:** Create a new user.
-   **Headers:**
    
    json
    

-   `{
      "Authorization": "Bearer jwt-token"
    }` 
    
-   **Request Body:**
    
    json
    
-   `{
      "username": "string",
      "password": "string",
      "email": "string"
    }` 
    
-   **Response:**
    
    json
    

-   `{
      "id": "int",
      "username": "string",
      "email": "string"
    }` 
    

#### Get All Users

-   **Endpoint:** `/api/users`
-   **Method:** `GET`
-   **Description:** Retrieve a list of all users.
-   **Headers:**
    
    json
    

-   `{
      "Authorization": "Bearer jwt-token"
    }` 
    
-   **Response:**
    
    json
    

-   `[
      {
        "id": "int",
        "username": "string",
        "email": "string"
      },
      ...
    ]` 
    

#### Get User by ID

-   **Endpoint:** `/api/users/:id`
-   **Method:** `GET`
-   **Description:** Retrieve a user by ID.
-   **Headers:**
    
    json
    

-   `{
      "Authorization": "Bearer jwt-token"
    }` 
    
-   **Response:**
    
    json
    

-   `{
      "id": "int",
      "username": "string",
      "email": "string"
    }` 
    

#### Update User

-   **Endpoint:** `/api/users/:id`
-   **Method:** `PUT`
-   **Description:** Update a user's information.
-   **Headers:**
    
    json
    

-   `{
      "Authorization": "Bearer jwt-token"
    }` 
    
-   **Request Body:**
    
    json
    
-   `{
      "username": "string",
      "email": "string"
    }` 
    
-   **Response:**
    
    json
    

-   `{
      "id": "int",
      "username": "string",
      "email": "string"
    }` 
    

#### Delete User

-   **Endpoint:** `/api/users/:id`
-   **Method:** `DELETE`
-   **Description:** Delete a user by ID.
-   **Headers:**
    
    json
    

-   `{
      "Authorization": "Bearer jwt-token"
    }` 
    
-   **Response:**
    
    json
    

-   `{
      "message": "User deleted successfully"
    }` 
    

## Error Handling

The API uses standard HTTP status codes to indicate the success or failure of an API request. Common status codes include:

-   `200 OK`: The request was successful.
-   `201 Created`: A new resource was successfully created.
-   `400 Bad Request`: The request was invalid or cannot be otherwise served.
-   `401 Unauthorized`: Authentication credentials were missing or incorrect.
-   `404 Not Found`: The requested resource could not be found.
-   `500 Internal Server Error`: An error occurred on the server.

