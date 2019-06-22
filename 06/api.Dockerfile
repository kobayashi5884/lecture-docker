FROM golang:1.12.5-alpine
WORKDIR /go/src/app
RUN apk add --no-cache git \
  && go get \
  github.com/auth0/go-jwt-middleware \
  github.com/go-sql-driver/mysql \
  github.com/gorilla/mux \
  github.com/jinzhu/gorm \
  github.com/rs/cors \
  github.com/urfave/negroni \
  golang.org/x/crypto/bcrypt