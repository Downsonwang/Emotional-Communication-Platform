FROM golang:1.17-alpine
MAINTAINER kelvinwang

WORKDIR $GOPATH/Gin
COPY . $GOPATH/Gin
RUN go build .
EXPOSE 8082

CMD ["./Gin"]



