FROM golang:1.8-alpine

# Set apps home directory.
ENV APP_DIR /go/src/github.com/MatsumotoHiroko/golang_api
RUN mkdir -p ${APP_DIR}

# Define current working directory.
WORKDIR ${APP_DIR}

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/BurntSushi/toml \
  && go get -u gopkg.in/mgo.v2 \
  && go get -u github.com/gorilla/mux \
  && go get -u gopkg.in/go-playground/validator.v9 \
  && go get -u github.com/auth0/go-jwt-middleware \
  && go get -u github.com/stretchr/testify \
  && go get -u github.com/dgrijalva/jwt-go

# Adds the application code to the image
ADD . ${APP_DIR}/
