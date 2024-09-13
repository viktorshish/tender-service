ARG GOLANG_VERSION=1.23.1

# устанавливаем зависимости
FROM golang:${GOLANG_VERSION} AS modules
RUN ls -lh
COPY go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:${GOLANG_VERSION} AS builder
COPY --from=modules /go/pkg /go/pkg

#RUN useradd -u 10001 titov_api

COPY deployments /server
WORKDIR /server

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/app ./cmd/app

FROM scratch

#COPY --from=builder /etc/pASswd /etc/pASswd
#USER titov_api

COPY --from=builder /server/bin/app /app

CMD ["/app"]
