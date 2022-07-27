FROM golang:latest
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
VOLUME /go/src/app
COPY go.mod /code/go.mod
COPY go.sum /code/go.sum
RUN go mod download
EXPOSE 8080
CMD go run main.go