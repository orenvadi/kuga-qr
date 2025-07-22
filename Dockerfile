FROM golang:1.24.4-alpine
RUN apk add --no-cache git
WORKDIR /
COPY . .
RUN go mod download
RUN go build -o api cmd/lms/main.go
EXPOSE 8080
CMD ["./api", "-c", "/config/dev.yaml"]
