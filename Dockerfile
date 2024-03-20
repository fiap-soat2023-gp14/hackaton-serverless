FROM golang:1.21-alpine3.18 as build

COPY ./src /src

WORKDIR /src

RUN go build -o main main.go

FROM alpine:3.18

RUN echo "$PWD" && ls -la
COPY --from=build /src/main /main
ENTRYPOINT [ "/main" ]