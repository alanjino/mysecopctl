FROM docker.io/library/golang:latest
COPY main.go main.go
CMD go run main.go