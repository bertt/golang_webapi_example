# golang_webapi_example

Example for building a testable Golang Web API

API has two requests:

- http get :8000/users

- http post :8000/users, body must have a user

Test coverage with:

```
gocov test | gocov report
````

todo:

- use database layer (postgres?)

- research virtual servers

