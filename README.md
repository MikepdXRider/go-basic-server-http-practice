# Goland Basic Server HTTP Practice

## Resources

- [Golang REST API With Mux - Travery Media](https://youtu.be/SonwZ6MF5BE)
- [Mux Github](https://github.com/gorilla/mux)
- [fmt package](https://pkg.go.dev/fmt#pkg-functions)

## How to use

- Clone repo to PC.
- `go mod tidy` in Comman Line to install dependencies.
- `go run .` to start up server.
- Postman, ThunderClient, or similar service recommended to make requests.
  - Browser will only work for GET methods.

## Available Routes

| method | operations                                                   | route                                   |
| ------ | ------------------------------------------------------------ | --------------------------------------- |
| GET    | returns all resources                                        | `http:/localhost/8000/api/v1/books`     |
| GET    | returns one resource by id                                   | `http:/localhost/8000/api/v1/books/:id` |
| DELETE | deletes one resource. returns deleted resource for reference | `http:/localhost/8000/api/v1/books/:id` |
| PUT    | replaces one resource. returns new resource for reference    | `http:/localhost/8000/api/v1/books/:id` |
| POST   | creates a new resource. returns new resource for reference   | `http:/localhost/8000/api/v1/books`     |
