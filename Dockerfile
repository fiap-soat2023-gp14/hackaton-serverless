FROM public.ecr.aws/docker/library/golang:1.19 as build
COPY ./src /src
WORKDIR /src
RUN go build -o main
FROM public.ecr.aws/lambda/provided:al2
RUN echo "$PWD" && ls -la
COPY --from=build /src/main .
ENTRYPOINT ./main