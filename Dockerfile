FROM golang:alpine
ADD . /go/src/github.com/bertt/golang_webapi_example
RUN apk add --update git
RUN go get github.com/gorilla/mux
RUN go install github.com/bertt/golang_webapi_example

CMD ["/go/bin/golang_webapi_example"]
EXPOSE 8000