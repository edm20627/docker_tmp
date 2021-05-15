FROM golang:1.15.1

WORKDIR /go/src/github.com/edm20627/project

ENV GO111MODULE=on

RUN apt-get update -qq && apt-get install -y vim git sqlite3 libsqlite3-dev -y

RUN GO111MODULE=off go get\
      github.com/pilu/fresh\
      golang.org/x/tools/cmd/goimports\
      github.com/kisielk/errcheck

EXPOSE 8080

CMD ["fresh"]
# CMD ["go", "run", "main.go"]
