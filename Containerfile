FROM golang:alpine

ENV CGO_ENABLED=0

WORKDIR /app

RUN set -eux; \
	apk update; \
	apk upgrade;

COPY . .

RUN set -eux; \
	go mod tidy; \
	go build -o bin/barband -ldflags="-s -w -extldflags='-static'" cmd/server/server.go;

CMD ["/app/bin/barband", "--help"]