# JWT Application
Application that generate a JWT and use it in order to authorize for the endpoints exposed

# API Specs

### `POST /signup`
Endpoint to create an user row in postgres db. The payload should have the following fields:

```json
{
  "email": "test@axiomzen.co",
  "password": "axiomzen",
  "firstName": "Alex",
  "lastName": "Zimmerman"
}
```

where `email` is an unique key in the database.

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token" 
}
```

### `POST /login`
Endpoint to log an user in. The payload should have the following fields:

```json
{
  "email": "test@axiomzen.co",
  "password": "axiomzen"
}
```

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token"
}
```

### `GET /users`
Endpoint to retrieve a json of all users. This endpoint requires a valid `x-authentication-token` header to be passed in with the request.

The response body should look like:
```json
{
  "users": [
    {
      "email": "test@axiomzen.co",
      "firstName": "Alex",
      "lastName": "Zimmerman"
    }
  ]
}
```

### `PUT /users`
Endpoint to update the current user `firstName` or `lastName` only. This endpoint requires a valid `x-authentication-token` header to be passed in and it should only update the user of the JWT being passed in. The payload can have the following fields:

```json
{
  "firstName": "NewFirstName",
  "lastName": "NewLastName"
}
```

The response can body can be empty.


# Instructions

### Packages
The project packages are structured in this way:
```
docker: docker file for create postgreSQL instance
domain: Interactions with the domain entities in the database
entity: Entity structs
server: API functionality to expose
services: Business logic
tools: Useful internal functionalities
```
### Environment
In order to create a PostgreSQL instance execute the following commands (docker installed is required):
```
cd docker
docker-compose up 
or 
docker-compose up -d
```
### Running the API
From a terminal execute the following command to run the API:
```
go run main.go
```
### Testing the API
Import the collection Dapper.postman_collection.json using Postman.
Use the _Health(GET)_ request in order to be sure the API is running, you should to get this message:

```Welcome to Eduard Reyes backend API for Dapper!```

Use _Signup(POST), Login(POST), Get Users(GET) and Update(PUT)_ request according to the Assignment requeriments.

### Unit Tests
From a terminal execute the following command to run the unit test:
```
go test ./...
```

### Verify the coverage
In order to generate an HTML coverage report, from a terminal execute the following commands:
```
go test -coverprofile c.out ./...
go tool cover -html c.out
```

### Last notes
Please reach out if you have any questions regarding the solution for the assignment.
```
Email: reydward@gmail.com
Mobile: +573168284844
```