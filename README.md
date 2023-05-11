# CRUD API using Go

This program is a simple RESTful CRUD (Create, Read, Update, Delete) API using the Go programming language. It demonstrates how to build a basic API server with endpoints for managing a collection of books. The program uses the gorilla/mux package to handle HTTP routing and JSON encoding/decoding.

The tutorial for this program can be found at the following URL: https://www.youtube.com/watch?v=gIe543ufGmE

The source code can be found at: https://github.com/BorntoDev/CRUD-Golang/

## Features
* Create a new book
* Retrieve all books
* Retrieve a single book by ID
* Update a book by ID
* Delete a book by ID

## Dependencies

To run this program, you will need to install the Gorilla Mux router package:
```
go get -u github.com/gorilla/mux
```

## Data structures

The program uses two main data structures:

1. Book: Represents a book with the following fields:
    * ID: Unique identifier for the book
    Isbn: International Standard Book Number
    * Title: Title of the book
    * Author: Pointer to the author struct
2. Author: Represents an author with the following fields:
    * Firstname: Author's first name
    * Lastname: Author's last name

## API Endpoints
The API has the following endpoints:

1. `GET /api/books`: Retrieve all books
2. `GET /api/books/{id}`: Retrieve a single book by ID
3. `POST /api/books`: Create a new book
4. `PUT /api/books/{id}`: Update a book by ID
5. `DELETE /api/books/{id}`: Delete a book by ID

## Usage
To run the program, execute the following command:
```
go run main.go
```

The API will start listening on port 8000.
