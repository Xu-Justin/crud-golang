# crud-golang

This project is a basic example of performing CRUD and API using Go language.

## Usage

Run the following commands to download required package.

```bash
go mod download
```

Run the following commands to run the applications.

```bash
cd crud-golang/cmd/main
go run main.go
```

## API

| Method   | Route   | Request                                                   | Description                                     |
|:--------:|---------|-----------------------------------------------------------|-------------------------------------------------|
| `GET`    | /       | -                                                         | Display all users.                              |
| `GET`    | /{id}   | -                                                         | Display an user by {id}.                        |
| `POST`   | /       | <pre>{<br>  "nama": str<br>  "age" : int<br>}</pre>       | Create an user.                                 |
| `PUT`    | /{id}   | <pre>{<br>  "nama": str<br>  "age" : int<br>}</pre>       | Update an user by {id}.                         |
| `DELETE` | /{id}   | -                                                         | Delete an user by {id}.                         |
