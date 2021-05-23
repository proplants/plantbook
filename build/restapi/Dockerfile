FROM golang:alpine as modules
ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:latest as builder
COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /app
ADD . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -a -tags netgo -ldflags "-s -w -extldflags '-static'" -o main ./cmd/plantbook-server

FROM scratch
ENV PORT=8081
ENV HOST=0.0.0.0
COPY --from=builder /app/main /app/main
ENTRYPOINT ["/app/main"]
CMD ["--scheme=http"]
LABEL org.opencontainers.image.source https://github.com/kaatinga/plantbook
EXPOSE $PORT