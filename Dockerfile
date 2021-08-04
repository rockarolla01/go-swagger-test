FROM golang:1.16-alpine

ADD go-swagger-test /bin/go-swagger-test
ADD swaggerui /bin/swaggerui

ENTRYPOINT /bin/go-swagger-test