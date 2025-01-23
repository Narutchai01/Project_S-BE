**# Project-S-BE

## Description

This is the backend for the Project-S project. It is a RESTful API that allows users to create, read, update, and delete users and posts.

## Installation

1. Clone the repository
   `git clone https://github.com/Narutchai01/Project_S-BE.git`
2. add `.env` file to the root directory follow by the `sample.env` file

```
PORT=
DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=
DB_PORT=



SUPA_API_URL=
SUPA_API_KEY=
SUPA_BUCKET_NAME=

JWT_SECRET_KEY=
```

3. Install the dependencies
   use `go install`

4. Run the server
   use `go run main.go` or `air` and `docker-compose up -d --build`

5. The server should be running on `localhost:8080`
go to `localhost:8080/swagger/index.html` to see the API documentation