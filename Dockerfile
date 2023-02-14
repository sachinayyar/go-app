FROM golang:1.17 AS build

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/app

FROM alpine:latest

COPY --from=build /go/bin/app /usr/local/bin/app

EXPOSE 8080

ENTRYPOINT ["app"]

CMD ["my-go-app"]
