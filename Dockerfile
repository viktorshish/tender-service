ARG GOLANG_VERSION=1.23.1

# устанавливаем зависимости
FROM golang:${GOLANG_VERSION} AS modules

COPY go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:${GOLANG_VERSION} AS builder
COPY --from=modules /go/pkg /go/pkg

COPY ./ /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/app ./cmd/app

FROM scratch

COPY --from=builder /app/bin/app /app

EXPOSE 8080

CMD ["/app"]
