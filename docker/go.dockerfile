FROM golang:latest

RUN apt update && apt -y upgrade

COPY ../.env /go/src/app/

WORKDIR /go/src/app/

EXPOSE 8000