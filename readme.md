# Event Management REST API

A Go-based REST API for managing events, users, and event registrations.

## Features

- User authentication with JWT
- Create, read, update, and delete events
- Register for events and cancel registrations
- Secure password hashing
- SQLite database persistence

## Tech Stack

- [Go](https://golang.org) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [SQLite](https://www.sqlite.org) - Database
- [JWT](https://github.com/golang-jwt/jwt) - Authentication
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing

## Setup

1. Clone the repository
2. Make sure you have Go installed (version 1.24+)
3. Install dependencies:
   ```
   go mod download
   ```
4. Run the application:
   ```
   go run main.go
   ```
5. The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication

- `POST /signup` - Create a new user account
- `POST /login` - Authenticate and receive a JWT token

### Events

- `GET /events` - Get all events
- `GET /events/:id` - Get a specific event
- `POST /events` - Create a new event (authenticated)
- `PUT /events/:id` - Update an event (authenticated, owner only)
- `DELETE /events/:id` - Delete an event (authenticated, owner only)

### Event Registration

- `POST /events/:id/register` - Register for an event (authenticated)
- `DELETE /events/:id/register` - Cancel event registration (authenticated)

## API Usage

### Authentication

#### Create a User

```http
POST http://localhost:8080/signup
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "test"
}
```

#### Login

```http
POST http://localhost:8080/login
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "test"
}
```

### Events

#### Get All Events

```http
GET http://localhost:8080/events
```

#### Get Single Event

```http
GET http://localhost:8080/events/1
```

#### Create an Event

```http
POST http://localhost:8080/events
Content-Type: application/json
Authorization: <your-jwt-token>

{
  "name": "Test event",
  "description": "Test event description",
  "location": "Test location",
  "dateTime": "2025-01-01T15:30:00.000Z"
}
```

#### Update an Event

```http
PUT http://localhost:8080/events/1
Content-Type: application/json
Authorization: <your-jwt-token>

{
  "name": "Updated test event",
  "description": "Updated test event description",
  "location": "Updated test location",
  "dateTime": "2025-01-01T15:30:00.000Z"
}
```

#### Delete an Event

```http
DELETE http://localhost:8080/events/1
Authorization: <your-jwt-token>
```

### Event Registration

#### Register for an Event

```http
POST http://localhost:8080/events/1/register
Authorization: <your-jwt-token>
```

#### Cancel Event Registration

```http
DELETE http://localhost:8080/events/1/register
Authorization: <your-jwt-token>
```

## Project Structure

- `main.go` - Application entry point
- `db/` - Database connection and initialization
- `models/` - Data models and database operations
- `routes/` - HTTP route handlers
- `middlewares/` - Request middleware (authentication)
- `utils/` - Utility functions (JWT, password hashing)
- `api-test/` - HTTP test files

## Database

The application uses SQLite with the following tables:

- `users` - Store user accounts
- `events` - Store event information
- `registrations` - Store event registrations

## Authentication

API uses JWT tokens for authentication. Include the token in the `Authorization` header for protected routes.
