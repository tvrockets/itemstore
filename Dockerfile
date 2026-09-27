# ---------- Этап сборки ----------
# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.27.1-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod ./

COPY . .

# Docker BuildKit автоматически прокидывает эти переменные
ARG TARGETOS
ARG TARGETARCH
ARG APP_NAME=app
ARG BUILD_PATH=./cmd/server

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o /app/bin/${APP_NAME} ${BUILD_PATH}

# ---------- Финальный этап ----------
FROM scratch AS final

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
COPY --from=builder /app/bin/app ./app

USER 65534:65534

EXPOSE 8080

ENTRYPOINT ["./app"]