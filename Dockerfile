FROM golang:alpine

WORKDIR /go/src

COPY ./Client/client.go ./api_client/
COPY ./Server/server.go ./api_server/

RUN apk update 
RUN apk add git

RUN go get -d -v ./api_server/
RUN go install -v ./api_server/
RUN go build ./api_client/client.go
EXPOSE 8080

CMD go run api_server/server.go
