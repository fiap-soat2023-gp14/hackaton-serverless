FROM golang:1.20.2-alpine3.16 as build
COPY ./src /src
WORKDIR /src
RUN GOARCH=amd64 GOOS=linux go build -o main main.go
FROM alpine:3.16
RUN echo "$PWD" && ls -la
COPY --from=build /src/main /main
ENTRYPOINT [ "/main" ]
