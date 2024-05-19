
# User Management App

## Overview

This project is a user management application built with Go using the Fiber framework and Gorm for database interaction. It supports user registration, login, and CRUD operations on user data. The application uses JWT for authentication and middleware for authorization.

## Prerequisites

-   Go 1.16 or later
-   MySQL database
-   [Go Modules](https://blog.golang.org/using-go-modules) enabled
-   [Go Fiber](https://gofiber.io/) framework
-   [Gorm](https://gorm.io/) ORM

## Getting Started

### Clone the repository


```
git clone https://github.com/yourusername/go-user-management.git
cd go-user-management 
```

### Install dependencies

```
go mod tidy 
```
### Environment Variables

Create a `.env` file in the root directory and configure the following environment variables:

env

```
PORT=8000
DB_NAME=user_management
DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
SECRET_TOKEN=your_secret_token

```

### Database Setup

Ensure you have a MySQL database running and create a database named `user_management`. Update the `.env` file with your database credentials.

### Run the Application


```
go run main.go

```

The application will start on the port specified in the `.env` file (default is 8000).

## Project Structure

go

```
.
├── configs
│   ├── app.go
│   └── database.go
├── controllers
│   ├── auth_controller.go
│   └── user_controller.go
├── database
│   └── database.go
├── middleware
│   ├── auth_middleware.go
│   └── admin_middleware.go
├── models
│   └── user.go
├── request
│   ├── auth_request.go
│   └── user_request.go
├── routes
│   ├── routes.go
│   └── user_routes.go
├── utils
│   ├── hash_util.go
│   └── jwt_util.go
├── .env
├── go.mod
├── go.sum
└── main.go

```

### Directory Breakdown

-   `configs`: Configuration for the application and database initialization.
-   `controllers`: Handlers for incoming HTTP requests and interaction with models.
-   `database`: Contains the database instance.
-   `middleware`: Middleware for request validation, authentication, and authorization.
-   `models`: Database models and schema definitions.
-   `request`: Request structs and validation for incoming payloads.
-   `routes`: API route definitions.
-   `utils`: Utility functions for password hashing and JWT management.

## API Endpoints

### Auth

-   POST /users/register: Register a new user.
-   POST /users/login: Login a user.

### User

-   GET /users: Get all users (Admin only).
-   PUT /users: Update user details.
-   DELETE /users/:id: Delete user (Admin only).

## Middleware

-   `AuthMiddleware`: Verifies JWT and authenticates users.
-   `CheckAdminMiddleware`: Checks if the authenticated user has admin privileges.

## Running Tests

You can run tests using the following command:

sh

```
go test ./...

```
## Contributing

1.  Fork the repository
2.  Create your feature branch (`git checkout -b feature/new-feature`)
3.  Commit your changes (`git commit -am 'Add new feature'`)
4.  Push to the branch (`git push origin feature/new-feature`)
5.  Create a new Pull Request

## License

This project is licensed under the MIT License.
## Documentation API

```
https://documenter.getpostman.com/view/27917478/2sA3JT3yWo
```
