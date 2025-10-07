FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o app ./backend/cmd/

FROM alpine:3.19
WORKDIR /root/
EXPOSE 8080
COPY --from=builder /app/app .
CMD ["./app"]