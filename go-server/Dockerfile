# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# ECS or Lambda へのデプロイを想定しているので、GOOS=linux
RUN GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
