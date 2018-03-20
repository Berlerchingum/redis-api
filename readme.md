
  # Redis API

A simple Rest API Server made in Go and Redis.    

## Getting Started    

These instructions will get you a copy of the project up and running on your local machine.

### Dependencies

The project has been created using the Go programming language and Redis a database. Both of which are necessary to use the application.

* [Golang](https://golang.org/)
* [Redis](https://redis.io/)

### Installing


Retrieve the project files directly from github

```
git clone http://github.com/Berlerchingum/redis-api
```

```
cd redis-api
```

Install the dependencies
```
go get -u github.com/go-redis/redis
```
and

```
go get -u github.com/gorilla/mux
```

Build the project

```
go build
``` 

Finally launch the executable

```
./api
```

## Model   


The application notes follows this specific model (Go)

```go
Id      int    `json:"id"`
Content string `json:"content"`
Author  string `json:"author"`
```    

Upon creation, the model is updated with the createdAt row containing the date of the creation

```go
CreatedAt string `json:"created_at"`
```    

## Routes    

Here is the list of the available routes.    

### GET all notes    

Retrieve all the notes .

```http
http://localhost:12345/notes
```    

### POST a note    

Create a new note. The payload of the request should contain datas according to the model described above.   

```http
http://localhost:12345/notes/
```    

### GET a note    

Retrieve a note. Id is mandatory.    

```http
http://localhost:12345/notes/{id}
```    

### DELETE a note    

Delete a note. Id is mandatory.    

```http
http://localhost:12345/notes/{id}
```    

## Tests    

The Redis database should already  be populated with 3 different notes. All the following examples results can be CURL'ed or retrieved using Postman.    

### GET all notes    

Retrieve all the posts.    

```json
http://localhost:12345/notes
```    

Result:    

```json
200 0K
```    

```json  
[
	{
		"id":1,
		"content":"Ma première note",
		"author":"auteur",
		"created_at":"18/03/2018"
	},
	{
		"id":2,
		"content":"Ma seconde note","author":"auteur",
		"created_at":"18/03/2018"
	}
]  
```    

### POST a note    

Create a new note.    

```http
http://localhost:12345/notes/
```    

The request payload is the following  

```json
{
	"id": 3,
	"content": "Ma troisieme note",
	"author": "author"
}
```    

Result:    

```json
200 0K 
```    

```json
{
	"id": 3,
	"content": "Ma troisieme note",
	"author": "author",
	"created_at":"18/03/2018"
}
```    

### GET a note    

Retrieve the note with id equal to 1.    

```http
http://localhost:12345/notes/1
```    

Result:    

```json
200 0K
```    

```jso
{
	"id":1,
	"content":"Ma première note",
	"author":"auteur",
	"created_at":"18/03/2018"
}
```    

### DELETE a note    

Delete the note with id equal to 1.    

```http
http://localhost:12345/notes/1
```    

Result:    

```json
 200 0K
```    

```json 
{}
```    

## Deployment    

Add additional notes about how to deploy this on a live system    

## Authors    

* **Benjamin Eguimendia** - *dev* - [berlerchingum](https://github.com/berlerchingum)  
* **Corentin Donzelli** - *dev* - [chickenfoot360](https://github.com/chickenfoot360)
