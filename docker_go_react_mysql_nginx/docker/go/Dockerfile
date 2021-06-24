FROM golang:1.16

WORKDIR /go/src/github.com/edm20627/project

RUN apt-get update -qq && apt-get install -y vim git default-mysql-client

RUN go install github.com/pilu/fresh@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/kisielk/errcheck@latest

EXPOSE 8080

ENTRYPOINT [ "fresh" ]

# CMD ["go", "run", "main.go"]
