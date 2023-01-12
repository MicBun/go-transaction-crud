FROM golang:latest

WORKDIR WORKDIR /go/src/github.com/MicBun/go-activity-tracking-api

COPY . .

RUN go get -d -v ./...

RUN go build -o go-activity-tracking-api .

EXPOSE 8080

ENTRYPOINT ["./go-activity-tracking-api"]