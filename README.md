# Bookstore API
Book store API with CRUD and SQLite3

## Prerequisites
* (Golang)[https://go.dev/]
* (curl)[https://curl.se/docs/manpage.html]

## How to build
Run following commands to build app:
```bash
go mod tidy
go build -o tmp/bookstore cmd/main.go
```

## How to run
Run binary image with this command:
```bash
./tmp/bookstore
```

## How to send requests using curl
Create a new book:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "System Design Interview", "author": "Alex Xu", "publisher":"ByteByteGo"}' localhost:8080/book/
```

Read/Get a list of books from database:
```bash
curl -X GET localhost:8080/book/
```
Read/Get single book:
```bash
curl -X GET localhost:8080/book/{id}
```
**NOTE**: `{id}` is a positive integer.

Update information of the book:
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "System Design Interview", "author": "Alex Xu", "publisher":"ByteByteGo"}' localhost:8080/book/{id}
```
**NOTE**: `{id}` is a positive integer.

Delete/Remove book from database:
```bash
curl -X DELETE localhost:8080/book/{id}
```
**NOTE**: `{id}` is a positive integer.
