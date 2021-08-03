# Basic CRUD in Golang

## Run

Run the server by using:
```bash
go run main.go
````

## Server

Runs on local server on port 10000 at:
```link
http://localhost:10000
````

## Routes

### Create (POST)
````
/add
`````
There should be a json object in the body

### Read (GET)
````
/get/{id}
`````
To get an object by its ID

`````
/list
`````
To list all objects in the repository

### Update (PUT)
````
/update/{id}
`````
Should come with the json object in the body.

### Delete (DELETE)
````
/delete/{id}
`````
To delete an object using its ID.
