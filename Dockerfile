# --- build stage ---
FROM golang:1.25-alpine AS build
WORKDIR /app

# 1) mod fajlovi (keš zavisnosti)
COPY go.mod go.sum ./
RUN go mod download

# 2) ceo projekat
COPY . .

# 3) build (entrypoint je ./cmd/app/main.go)
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-s -w" -o /bin/app ./cmd/app

# --- runtime stage ---
FROM alpine:3.20
# certovi za outbound HTTP(S), non-root user
RUN adduser -D -g '' app && apk add --no-cache ca-certificates
WORKDIR /home/app
USER app

# kopiraj statički binar
COPY --from=build /bin/app /usr/local/bin/app

# port i entrypoint
ENV SERVER_PORT=8080
EXPOSE 8080
CMD ["app"]
