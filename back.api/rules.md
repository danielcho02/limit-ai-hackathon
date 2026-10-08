# Basics
- web framework : gin-gonic
- database : GORM PostgreSQL driver(docker image). 5332 port(localhost)

# Main Structure : Domain-Driven Development
- repository : database logic(basic CRUD, pagination, transaction)
- service : complicated & important logics(or helper function for specific domain)
- usecase : general service logic(not complicated). combine service & repository logic
- domain : define model beeing used in business logic
- handler/v1 : version 1 API structs(request/response)
- handler : handle http requests by calling usecase functions
- presenter : convert domain struct into response struct(ex: domain.User -> v1.UserInfo)

etc : utils(helper functions), config

# Every Request passes Middleware
- Error Handling
type & usage definition -> internal/errors
error handling middleware -> internal/middleware/error.go

- User Auth Handling : internal/middleware/auth.go

# Use Makefile for local-test
In Makefile in root directory, many instructions(frequently used) are defined.

