# User Management System - Backend

Backend part of the User Managament System written in Go.

## Used Technologies
- Programming Language: Go
- API Documentation: Swagger
- Mocking: Mockery
- Unit Testing: Go testing package
- Docker: Containerization for deployment

## Getting Started

To run the backend, follow these steps:

1. Install Go packages:

    ```bash
    go get ./...
    ```

2. Run the project:

    ```bash
    go run .
    ```

    The backend will be available at http://127.0.0.1:3001

The unit test suite covers functions such as handling Getting All Users, Getting User By Id, Creating User, Updating User and Deleting User

## Swagger Documentation

For API documentation, visit [Swagger Documentation](http://127.0.0.1:3001/swagger). This documentation provides insights into available operations.

### API Endpoints

- **GET /api/users**
  - Get a list of all users.

- **POST /api/users**
  - Create a new user.

- **GET /api/users/{id}**
  - Get a user by ID.

- **PUT /api/users/{id}**
  - Update a user by ID.

- **DELETE /api/users/{id}**
  - Delete a user by ID.


## Mockery and Unit Testing

Mockery was utilized for mocking in unit tests, following the "given when then" principle. 

To run tests:

```bash
go test ./...
```

## Dockerization

To build and run the backend using Docker:

1. Build the Docker image:

```bash
docker build --tag fl-case-study-backend .
```

2. Run the docker container

```bash
docker run -p 3001:3001 fl-case-study-backend
```
Now, the backend is accessible at http://localhost:3001