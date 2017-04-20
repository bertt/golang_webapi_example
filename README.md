# golang_webapi_example

Example for building a testable Golang Web API

Used components: Mux router, MongoDB database, Docker

Install:

```
glide install

```

Startup database:

```
docker run -p 27017:27017 -d mongo

```

Startup server:

```
go run main.go

```

API has two requests:

- http get :8000/users

- http post :8000/users, body must have a user

Test coverage with:

```
gocov test | gocov report
````

todo:


- research virtual servers

