FROM golang:1.16.4-alpine3.13 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o plantbook .

FROM golang:1.16.4-stretch

COPY --from=builder /app/plantbook /

EXPOSE 80

CMD ["/plantbook"]