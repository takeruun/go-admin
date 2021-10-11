FROM golang:1.17.1-alpine

RUN apk update && apk --no-cache add git build-base nodejs npm

WORKDIR /go/src/app

RUN go get golang.org/x/tools/gopls@latest && \ 
  go get github.com/rubenv/sql-migrate/... && \
  go get github.com/pilu/fresh

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["fresh","-c", "runner.conf"]
